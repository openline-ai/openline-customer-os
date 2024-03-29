package server

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/validator"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/common"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/resolver"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/grpc_client"
	cosHandler "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/handler"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/mapper"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/metrics"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/rest"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/service"
	commonAuthService "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-auth/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/caches"
	commonConfig "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/config"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	commonrepository "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository"
	commonrepositorypostgres "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/repository/postgres"
	commonservice "github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/service"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

type server struct {
	cfg    *config.Config
	log    logger.Logger
	doneCh chan struct{}
}

func NewServer(cfg *config.Config, log logger.Logger) *server {
	return &server{cfg: cfg, log: log, doneCh: make(chan struct{})}
}

func (server *server) Run(parentCtx context.Context) error {
	ctx, cancel := signal.NotifyContext(parentCtx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := validator.GetValidator().Struct(server.cfg); err != nil {
		return errors.Wrap(err, "cfg validate")
	}

	// Setting up tracing
	tracer, closer, err := tracing.NewJaegerTracer(&server.cfg.Jaeger, server.log)
	if err != nil {
		server.log.Fatalf("Could not initialize jaeger tracer: %s", err.Error())
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	registerPrometheusMetrics()

	// Initialize postgres db
	db, _ := InitDB(server.cfg, server.log)
	defer db.SqlDB.Close()

	// Setting up Neo4j
	neo4jDriver, err := commonConfig.NewNeo4jDriver(server.cfg.Neo4j)
	if err != nil {
		server.log.Fatalf("Could not establish connection with neo4j at: %v, error: %v", server.cfg.Neo4j.Target, err.Error())
	}
	defer neo4jDriver.Close(ctx)
	// check neo4j connectivity
	err = neo4jDriver.VerifyConnectivity(ctx)
	if err != nil {
		server.log.Fatalf("Could not verify connectivity with neo4j at: %v, error: %v", server.cfg.Neo4j.Target, err.Error())
	}

	// Setting up Postgres repositories
	commonServices := commonservice.InitServices(db.GormDB, &neo4jDriver)
	commonAuthServices := commonAuthService.InitServices(nil, commonServices, db.GormDB)

	// Setting up gRPC client
	df := grpc_client.NewDialFactory(server.cfg)
	gRPCconn, err := df.GetEventsProcessingPlatformConn()
	if err != nil {
		server.log.Fatalf("Failed to connect: %v", err)
	}
	defer df.Close(gRPCconn)
	grpcContainer := grpc_client.InitClients(gRPCconn)

	// Setting up Gin
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	adminApiHandler := cosHandler.NewAdminApiHandler(server.cfg, commonServices)

	serviceContainer := service.InitServices(server.log, &neo4jDriver, server.cfg, commonServices, commonAuthServices, grpcContainer)
	r.Use(cors.New(corsConfig))
	r.Use(ginzap.GinzapWithConfig(server.log.Logger(), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/metrics", "/health", "/readiness", "/"},
	}))
	r.Use(ginzap.RecoveryWithZap(server.log.Logger(), true))
	r.Use(prometheusMiddleware())
	r.Use(bodyLoggerMiddleware)

	commonCache := caches.NewCommonCache()

	r.POST("/query",
		cosHandler.TracingEnhancer(ctx, "/query"),
		apiKeyCheckerHTTPMiddleware(commonServices.CommonRepositories.TenantWebhookApiKeyRepository, commonServices.CommonRepositories.AppKeyRepository, commonservice.CUSTOMER_OS_API, commonservice.WithCache(commonCache)),
		tenantUserContextEnhancerMiddleware(commonservice.USERNAME_OR_TENANT, commonServices.CommonRepositories, commonservice.WithCache(commonCache)),
		server.graphqlHandler(grpcContainer, serviceContainer))
	if server.cfg.GraphQL.PlaygroundEnabled {
		r.GET("/", playgroundHandler())
	}
	r.GET("/whoami",
		cosHandler.TracingEnhancer(ctx, "/whoami"),
		commonservice.ApiKeyCheckerHTTP(commonServices.CommonRepositories.TenantWebhookApiKeyRepository, commonServices.CommonRepositories.AppKeyRepository, commonservice.CUSTOMER_OS_API, commonservice.WithCache(commonCache)),
		rest.WhoamiHandler(serviceContainer))
	r.POST("/admin/query",
		cosHandler.TracingEnhancer(ctx, "/admin/query"),
		adminApiHandler.GetAdminApiHandlerEnhancer(),
		server.graphqlHandler(grpcContainer, serviceContainer))

	if server.cfg.GraphQL.PlaygroundEnabled {
		r.GET("/admin/",
			cosHandler.TracingEnhancer(ctx, "/admin"),
			playgroundAdminHandler())
	}

	r.GET("/health", healthCheckHandler)
	r.GET("/readiness", healthCheckHandler)

	if server.cfg.ApiPort == server.cfg.MetricsPort {
		r.GET(server.cfg.Metrics.PrometheusPath, metricsHandler)
	} else {
		go func() {
			mr := gin.Default()
			mr.Use(prometheusMiddleware())
			mr.Use(bodyLoggerMiddleware)
			mr.GET(server.cfg.Metrics.PrometheusPath, metricsHandler)
			mr.Run(":" + server.cfg.MetricsPort)
		}()
	}

	r.Run(":" + server.cfg.ApiPort)

	<-server.doneCh
	server.log.Infof("Application %s exited properly", constants.ServiceName)
	return nil
}

// Define a custom middleware adapter for ApiKeyCheckerHTTP.
func apiKeyCheckerHTTPMiddleware(tenantApiKeyRepo commonrepositorypostgres.TenantWebhookApiKeyRepository, appKeyRepo commonrepositorypostgres.AppKeyRepository, app commonservice.App, opts ...commonservice.CommonServiceOption) func(c *gin.Context) {
	apiKeyChecker := commonservice.ApiKeyCheckerHTTP(tenantApiKeyRepo, appKeyRepo, app, opts...)
	return func(c *gin.Context) {
		if isIntrospectionQuery(c.Request) {
			c.Next() // Skip ApiKeyCheckerHTTP and continue to the next handler.
			return
		}
		apiKeyChecker(c)
	}
}

// Define a custom middleware adapter for TenantUserContextEnhancer.
func tenantUserContextEnhancerMiddleware(userContextType commonservice.HeaderAllowance, repos *commonrepository.Repositories, opts ...commonservice.CommonServiceOption) func(c *gin.Context) {
	tenantEnhancer := commonservice.TenantUserContextEnhancer(userContextType, repos, opts...)
	return func(c *gin.Context) {
		if isIntrospectionQuery(c.Request) {
			c.Next() // Skip TenantUserContextEnhancer and continue to the next handler.
			return
		}
		tenantEnhancer(c)
	}
}

func isIntrospectionQuery(req *http.Request) bool {
	var requestMap map[string]interface{}
	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		return false
	}

	// Create a new ReadCloser from the read bytes.
	req.Body = io.NopCloser(bytes.NewReader(requestBody))

	if err = json.Unmarshal(requestBody, &requestMap); err != nil {
		return false
	}

	if opName, ok := requestMap["operationName"].(string); ok && opName == "IntrospectionQuery" {
		// Check if "__schema" is present in the request
		if selectionSet, ok := requestMap["query"].(string); ok {
			if strings.Contains(selectionSet, "__schema {") && strings.Contains(selectionSet, "query IntrospectionQuery") {
				return true
			}
		}
	}
	return false
}

func InitDB(cfg *config.Config, log logger.Logger) (db *commonConfig.StorageDB, err error) {
	if db, err = commonConfig.NewPostgresDBConn(cfg.Postgres); err != nil {
		log.Fatalf("Could not open db connection: %s", err.Error())
	}
	return
}

func (server *server) graphqlHandler(grpcContainer *grpc_client.Clients, serviceContainer *service.Services) gin.HandlerFunc {
	// instantiate graph resolver
	graphResolver := resolver.NewResolver(server.log, serviceContainer, grpcContainer)
	// make a data loader
	loader := dataloader.NewDataLoader(serviceContainer)
	schemaConfig := generated.Config{Resolvers: graphResolver}
	schemaConfig.Directives.HasRole = cosHandler.GetRoleChecker()
	schemaConfig.Directives.HasTenant = cosHandler.GetTenantChecker()
	schemaConfig.Directives.HasIdentityId = cosHandler.GetIdentityIdChecker()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(schemaConfig))
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		var e error

		if er, ok := err.(error); ok {
			e = er
		}
		if e == nil {
			e = ctx.Err()
		}

		if errors.Is(e, context.Canceled) {
			return gqlerror.Errorf("Request canceled!")
		}
		if errors.Is(e, context.DeadlineExceeded) {
			return gqlerror.Errorf("Request timeout!")
		}

		buf := make([]byte, 4096)
		stackSize := runtime.Stack(buf, false)
		server.log.Errorf("panic occurred: %v\nBacktrace:\n%s", e, string(buf[:stackSize]))
		return gqlerror.Errorf("Internal server error!")
	})
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		// Error hook place, Returned error can be customized. Check https://gqlgen.com/reference/errors/
		return err
	})
	srv.Use(extension.FixedComplexityLimit(server.cfg.GraphQL.FixedComplexityLimit))

	return func(c *gin.Context) {
		customCtx := &common.CustomContext{}
		if c.Keys[commonservice.KEY_TENANT_NAME] != nil {
			customCtx.Tenant = c.Keys[commonservice.KEY_TENANT_NAME].(string)
		}
		if c.Keys[commonservice.KEY_USER_ROLES] != nil {
			customCtx.Roles = mapper.MapRolesToModel(c.Keys[commonservice.KEY_USER_ROLES].([]string))
		}
		if c.Keys[commonservice.KEY_USER_ID] != nil {
			customCtx.UserId = c.Keys[commonservice.KEY_USER_ID].(string)
		}
		if c.Keys[commonservice.KEY_USER_EMAIL] != nil {
			customCtx.UserEmail = c.Keys[commonservice.KEY_USER_EMAIL].(string)
		}
		if c.Keys[commonservice.KEY_IDENTITY_ID] != nil {
			customCtx.IdentityId = c.Keys[commonservice.KEY_IDENTITY_ID].(string)
		}

		graphqlOperationName := extractGraphQLMethodName(c.Request)
		c.Request.Header.Set("X-GraphQL-Operation-Name", graphqlOperationName)
		customCtx.GraphqlRootOperationName = graphqlOperationName

		logMiddleware := loggerMiddleware(customCtx, graphqlOperationName)
		logMiddleware(c)

		dataloaderMiddleware := dataloader.Middleware(loader, srv)
		h := common.WithContext(customCtx, dataloaderMiddleware)

		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundAdminHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/admin/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "OK"})
}

func metricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func registerPrometheusMetrics() {
	prometheus.MustRegister(metrics.MetricsGraphqlRequestCount)
	prometheus.MustRegister(metrics.MetricsGraphqlRequestDuration)
	prometheus.MustRegister(metrics.MetricsGraphqlRequestErrorCount)
}

func loggerMiddleware(ctx *common.CustomContext, graphqlOperationName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		zap.L().With(
			zap.String("tenant", ctx.Tenant),
			zap.String("userId", ctx.UserId),
			zap.String("identityId", ctx.IdentityId),
		).Sugar().Infof("GraphQL Method: %s", graphqlOperationName)

		// Execute the GraphQL handler
		c.Next()
	}
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(start time.Time) {
			duration := time.Since(start)

			operationName := c.Request.Header.Get("X-GraphQL-Operation-Name")
			if operationName == "" {
				operationName = c.Request.URL.Path
			}
			metrics.MetricsGraphqlRequestDuration.WithLabelValues(operationName).Observe(duration.Seconds())
			metrics.MetricsGraphqlRequestCount.WithLabelValues(operationName, strconv.Itoa(c.Writer.Status())).Inc()
			// Increment the error count if the GraphQL response has errors
			if len(c.Errors) > 0 || (c.Writer.Size() > 0 && c.Writer.Written()) {
				var response struct {
					Errors []struct{} `json:"errors"`
				}
				bodyBytes := c.MustGet("bodyBytes").([]byte)
				if err := json.Unmarshal(bodyBytes, &response); err == nil && len(response.Errors) > 0 {
					metrics.MetricsGraphqlRequestErrorCount.WithLabelValues(operationName).Inc()
				}
			}
		}(time.Now())
		c.Next()
	}
}

func bodyLoggerMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	c.Set("bodyBytes", blw.body.Bytes())
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func extractGraphQLMethodName(req *http.Request) string {
	// Read the request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		// Handle error
		return ""
	}

	// Restore the request body
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	// Parse the request body as JSON
	var requestBody map[string]interface{}
	if err := json.Unmarshal(body, &requestBody); err != nil {
		// Handle error
		return ""
	}

	// Extract the method name from the GraphQL request
	if operationName, ok := requestBody["operationName"].(string); ok {
		return operationName
	}

	// If the method name is not found, you can add additional logic here to extract it from the request body or headers if applicable
	// ...
	return ""
}
