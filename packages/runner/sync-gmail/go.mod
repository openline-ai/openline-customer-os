module github.com/openline-ai/openline-customer-os/packages/runner/sync-gmail

go 1.20

//replace github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module => ./../../server/customer-os-common-module

require (
	github.com/caarlos0/env/v6 v6.10.1
	github.com/google/uuid v1.3.1
	github.com/joho/godotenv v1.5.1
	github.com/neo4j/neo4j-go-driver/v5 v5.13.0
	github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module v0.0.0-20230908062116-37b5a96799dd
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.9.3
	gorm.io/driver/postgres v1.5.3
	gorm.io/gorm v1.25.5
)

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.25.0 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
