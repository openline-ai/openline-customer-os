package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/constants"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/domain/email/events"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type EmailRepository interface {
	FailEmailValidation(ctx context.Context, emailId string, event events.EmailFailedValidationEvent) error
	EmailValidated(ctx context.Context, emailId string, event events.EmailValidatedEvent) error
	LinkWithContact(ctx context.Context, tenant, contactId, emailId, label string, primary bool, updatedAt time.Time) error
	LinkWithOrganization(ctx context.Context, tenant, organizationId, emailId, label string, primary bool, updatedAt time.Time) error
	LinkWithUser(ctx context.Context, tenant, userId, emailId, label string, primary bool, updatedAt time.Time) error
}

type emailRepository struct {
	driver *neo4j.DriverWithContext
}

func NewEmailRepository(driver *neo4j.DriverWithContext) EmailRepository {
	return &emailRepository{
		driver: driver,
	}
}

func (r *emailRepository) FailEmailValidation(ctx context.Context, emailId string, event events.EmailFailedValidationEvent) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailRepository.FailEmailValidation")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, event.Tenant)
	span.LogFields(log.String("emailId", emailId))

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (t:Tenant {name:$tenant})<-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]-(e:Email:Email_%s {id:$id})
		 		SET e.validationError = $validationError,
		     		e.validated = false,
					e.updatedAt = $validatedAt`

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, event.Tenant),
			map[string]any{
				"id":              emailId,
				"tenant":          event.Tenant,
				"validationError": event.ValidationError,
				"validatedAt":     event.ValidatedAt,
			})
		return nil, err
	})
	return err
}

func (r *emailRepository) EmailValidated(ctx context.Context, emailId string, event events.EmailValidatedEvent) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailRepository.EmailValidated")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, event.Tenant)
	span.LogFields(log.String("emailId", emailId))

	cypher := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})<-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]-(e:Email:Email_%s {id:$id})
		 		SET e.validationError = $validationError,
					e.email = $email,
		     		e.validated = true,
					e.acceptsMail = $acceptsMail,
					e.canConnectSmtp = $canConnectSmtp,
					e.hasFullInbox = $hasFullInbox,
					e.isCatchAll = $isCatchAll,
					e.isDeliverable = $isDeliverable,
					e.isDisabled = $isDisabled,
					e.isValidSyntax = $isValidSyntax,
					e.username = $username,
					e.updatedAt = $validatedAt,
					e.isReachable = $isReachable
				WITH e, CASE WHEN $domain <> '' THEN true ELSE false END AS shouldMergeDomain
				WHERE shouldMergeDomain
				MERGE (d:Domain {domain:$domain})
				ON CREATE SET 	d.id=randomUUID(), 
								d.createdAt=$now, 
								d.updatedAt=$now,
								d.appSource=$source,
								d.source=$appSource
				WITH d, e
				MERGE (e)-[:HAS_DOMAIN]->(d)`, event.Tenant)
	params := map[string]any{
		"id":              emailId,
		"tenant":          event.Tenant,
		"validationError": event.ValidationError,
		"email":           event.EmailAddress,
		"domain":          strings.ToLower(event.Domain),
		"acceptsMail":     event.AcceptsMail,
		"canConnectSmtp":  event.CanConnectSmtp,
		"hasFullInbox":    event.HasFullInbox,
		"isCatchAll":      event.IsCatchAll,
		"isDeliverable":   event.IsDeliverable,
		"isDisabled":      event.IsDisabled,
		"isValidSyntax":   event.IsValidSyntax,
		"username":        event.Username,
		"validatedAt":     event.ValidatedAt,
		"isReachable":     event.IsReachable,
		"now":             utils.Now(),
		"source":          constants.SourceOpenline,
		"appSource":       constants.AppSourceEventProcessingPlatform,
	}
	span.LogFields(log.String("cypher", cypher))
	tracing.LogObjectAsJson(span, "params", params)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, cypher, params)
		return nil, err
	})
	return err
}

func (r *emailRepository) LinkWithContact(ctx context.Context, tenant, contactId, emailId, label string, primary bool, updatedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailRepository.LinkWithContact")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("emailId", emailId), log.String("contactId", contactId))

	query := `MATCH (t:Tenant {name:$tenant})<-[:CONTACT_BELONGS_TO_TENANT]-(c:Contact {id:$contactId}),
				(t)<-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]-(e:Email {id:$emailId})
		MERGE (c)-[rel:HAS]->(e)
		SET	rel.primary = $primary,
			rel.label = $label,	
			c.updatedAt = $updatedAt,
			rel.syncedWithEventStore = true`
	params := map[string]any{
		"tenant":    tenant,
		"contactId": contactId,
		"emailId":   emailId,
		"label":     label,
		"primary":   primary,
		"updatedAt": updatedAt,
	}
	span.LogFields(log.String("query", query), log.Object("params", params))

	return r.executeQuery(ctx, query, params)
}

func (r *emailRepository) LinkWithOrganization(ctx context.Context, tenant, organizationId, emailId, label string, primary bool, updatedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailRepository.LinkWithOrganization")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("emailId", emailId), log.String("organizationId", organizationId))

	query := `
		MATCH (t:Tenant {name:$tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(org:Organization {id:$organizationId}),
				(t)<-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]-(e:Email {id:$emailId})
		MERGE (org)-[rel:HAS]->(e)
		SET	rel.primary = $primary,
			rel.label = $label,	
			org.updatedAt = $updatedAt,
			rel.syncedWithEventStore = true`
	params := map[string]any{
		"tenant":         tenant,
		"organizationId": organizationId,
		"emailId":        emailId,
		"label":          label,
		"primary":        primary,
		"updatedAt":      updatedAt,
	}
	span.LogFields(log.String("query", query), log.Object("params", params))

	return r.executeQuery(ctx, query, params)
}

func (r *emailRepository) LinkWithUser(ctx context.Context, tenant, userId, emailId, label string, primary bool, updatedAt time.Time) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailRepository.LinkWithUser")
	defer span.Finish()
	tracing.SetNeo4jRepositorySpanTags(ctx, span, tenant)
	span.LogFields(log.String("emailId", emailId), log.String("userId", userId))

	query := `
		MATCH (t:Tenant {name:$tenant})<-[:USER_BELONGS_TO_TENANT]-(u:User {id:$userId}),
				(t)<-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]-(e:Email {id:$emailId})
		MERGE (u)-[rel:HAS]->(e)
		SET	rel.primary = $primary,
			rel.label = $label,	
			u.updatedAt = $updatedAt,
			rel.syncedWithEventStore = true`
	params := map[string]any{
		"tenant":    tenant,
		"userId":    userId,
		"emailId":   emailId,
		"label":     label,
		"primary":   primary,
		"updatedAt": updatedAt,
	}
	span.LogFields(log.String("query", query), log.Object("params", params))
	return r.executeQuery(ctx, query, params)
}

func (r *emailRepository) executeQuery(ctx context.Context, query string, params map[string]any) error {
	return utils.ExecuteWriteQuery(ctx, *r.driver, query, params)
}
