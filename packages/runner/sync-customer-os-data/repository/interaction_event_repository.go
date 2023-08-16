package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/constants"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/tracing"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type InteractionEventRepository interface {
	GetMatchedInteractionEvent(ctx context.Context, tenant string, event entity.InteractionEventData) (string, error)
	MergeInteractionEvent(ctx context.Context, tenant string, syncDate time.Time, event entity.InteractionEventData) error
	MergeInteractionSessionForEvent(ctx context.Context, tenant, eventId, externalSource string, syncDate time.Time, session entity.InteractionSession) error

	MergeEmailInteractionSession(ctx context.Context, tenant string, date time.Time, message entity.EmailMessageData) (string, error)
	MergeEmailInteractionEvent(ctx context.Context, tenant, externalSystemId string, date time.Time, message entity.EmailMessageData) (string, error)
	LinkInteractionEventToSession(ctx context.Context, tenant, interactionEventId, interactionSessionId string) error

	InteractionEventSentByEmail(ctx context.Context, tenant, interactionEventId, emailId string) error
	InteractionEventSentToEmails(ctx context.Context, tenant, interactionEventId, sentType string, emails []string) error
	LinkInteractionEventAsPartOfByExternalId(ctx context.Context, tenant string, event entity.InteractionEventData) error

	FindParticipantByExternalId(ctx context.Context, tenant, externalId, externalSystem string) (*dbtype.Node, error)

	LinkInteractionEventWithSenderByExternalId(ctx context.Context, tenant, eventId, externalSystem string, sender entity.InteractionEventParticipant) error
	LinkInteractionEventWithRecipientByExternalId(ctx context.Context, tenant, eventId, externalSystem string, recipient entity.InteractionEventParticipant) error
	LinkInteractionEventWithRecipientByOpenlineId(ctx context.Context, tenant, eventId string, recipient entity.InteractionEventParticipant) error
	LinkInteractionEventWithSenderJobRole(ctx context.Context, tenant string, interactionEventId, organizationId, contactId string) error
	LinkInteractionEventWithRecipientJobRole(ctx context.Context, tenant string, interactionEventId, organizationId, contactId, relationType string) error
}

type interactionEventRepository struct {
	driver *neo4j.DriverWithContext
}

func NewInteractionEventRepository(driver *neo4j.DriverWithContext) InteractionEventRepository {
	return &interactionEventRepository{
		driver: driver,
	}
}

func (r *interactionEventRepository) GetMatchedInteractionEvent(ctx context.Context, tenant string, event entity.InteractionEventData) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.GetMatchedInteractionEvent")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (t:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(ext:ExternalSystem {id:$externalSystem})
				OPTIONAL MATCH (ext)<-[:IS_LINKED_WITH {externalId:$issueExternalId}]-(ie:InteractionEvent_%s)
				WITH ie WHERE ie is not null
				return ie.id limit 1`

	dbRecords, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]interface{}{
				"tenant":          tenant,
				"externalSystem":  event.ExternalSystem,
				"issueExternalId": event.ExternalId,
			})
		if err != nil {
			return nil, err
		}
		return queryResult.Collect(ctx)
	})
	if err != nil {
		return "", err
	}
	issueIDs := dbRecords.([]*db.Record)
	if len(issueIDs) > 0 {
		return issueIDs[0].Values[0].(string), nil
	}
	return "", nil
}

func (r *interactionEventRepository) MergeInteractionEvent(ctx context.Context, tenant string, syncDate time.Time, event entity.InteractionEventData) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.MergeInteractionEvent")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystemId}) 
		 MERGE (ie:InteractionEvent_%s {id:$id})-[rel:IS_LINKED_WITH {externalId:$externalId}]->(e) 
		 ON CREATE SET 
		  	ie:InteractionEvent,
		  	ie:TimelineEvent, 
		  	ie:TimelineEvent_%s, 
		  	rel.syncDate=$syncDate, 
		  	ie.createdAt=$createdAt,
			ie.channel=$channel,
			ie.eventType=$type, 
		  	ie.identifier=$identifier, 
		  	ie.content=$content, 
		  	ie.contentType=$contentType,
			ie.hide=$hide,
		  	ie.source=$source, 
		  	ie.sourceOfTruth=$sourceOfTruth,
		  	ie.appSource=$appSource
		 ON MATCH SET 
			ie.content=$content, 
		  	ie.contentType=$contentType
		 RETURN ie.id`

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		params := map[string]interface{}{
			"tenant":           tenant,
			"id":               event.Id,
			"content":          event.Content,
			"contentType":      event.ContentType,
			"externalSystemId": event.ExternalSystem,
			"createdAt":        utils.TimePtrFirstNonNilNillableAsAny(event.CreatedAt),
			"type":             event.Type,
			"identifier":       utils.FirstNotEmpty(event.Identifier, event.ExternalId),
			"source":           event.ExternalSystem,
			"sourceOfTruth":    event.ExternalSystem,
			"appSource":        constants.AppSourceSyncCustomerOsData,
			"externalId":       event.ExternalId,
			"syncDate":         syncDate,
			"channel":          event.Channel,
			"hide":             event.Hide,
		}

		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			params)
		if err != nil {
			return nil, err
		}
		_, err = queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (r *interactionEventRepository) MergeEmailInteractionSession(ctx context.Context, tenant string, syncDate time.Time, message entity.EmailMessageData) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.MergeEmailInteractionSession")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MERGE (is:InteractionSession_%s {identifier:$identifier, source:$source, channel:$channel}) " +
		" ON CREATE SET " +
		"  is:InteractionSession, " +
		"  is.id=randomUUID(), " +
		"  is.syncDate=$syncDate, " +
		"  is.createdAt=$createdAt, " +
		"  is.name=$name, " +
		"  is.status=$status," +
		"  is.type=$type," +
		"  is.sourceOfTruth=$sourceOfTruth, " +
		"  is.appSource=$appSource " +
		" WITH is " +
		" RETURN is.id"

	dbRecord, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]interface{}{
				"tenant":        tenant,
				"source":        message.ExternalSystem,
				"sourceOfTruth": message.ExternalSystem,
				"appSource":     constants.AppSourceSyncCustomerOsData,
				"identifier":    message.EmailThreadId,
				"name":          message.Subject,
				"syncDate":      syncDate,
				"createdAt":     utils.TimePtrFirstNonNilNillableAsAny(message.CreatedAt),
				"status":        "ACTIVE",
				"type":          "THREAD",
				"channel":       "EMAIL",
			})
		if err != nil {
			return nil, err
		}
		record, err := queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record, nil
	})
	if err != nil {
		return "", err
	}
	return dbRecord.(*db.Record).Values[0].(string), nil
}

func (r *interactionEventRepository) MergeEmailInteractionEvent(ctx context.Context, tenant, externalSystemId string, syncDate time.Time, message entity.EmailMessageData) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.MergeEmailInteractionEvent")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystemId}) " +
		" MERGE (ie:InteractionEvent_%s {source:$source, channel:$channel})-[rel:IS_LINKED_WITH {externalId:$externalId}]->(e) " +
		" ON CREATE SET " +
		"  ie:InteractionEvent, " +
		"  ie:TimelineEvent, " +
		"  ie:TimelineEvent_%s, " +
		"  rel.syncDate=$syncDate, " +
		"  ie.createdAt=$createdAt, " +
		"  ie.id=randomUUID(), " +
		"  ie.identifier=$identifier, " +
		"  ie.content=$content, " +
		"  ie.contentType=$contentType, " +
		"  ie.sourceOfTruth=$sourceOfTruth, " +
		"  ie.appSource=$appSource " +
		" WITH ie " +
		" RETURN ie.id"

	dbRecord, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		params := map[string]interface{}{
			"tenant":           tenant,
			"externalSystemId": externalSystemId,
			"identifier":       message.EmailMessageId,
			"source":           message.ExternalSystem,
			"sourceOfTruth":    message.ExternalSystem,
			"appSource":        constants.AppSourceSyncCustomerOsData,
			"externalId":       message.ExternalId,
			"syncDate":         syncDate,
			"createdAt":        utils.TimePtrFirstNonNilNillableAsAny(message.CreatedAt),
			"channel":          "EMAIL",
		}

		if message.Html != "" {
			params["content"] = message.Html
			params["contentType"] = "text/html"
		} else {
			params["content"] = message.Text
			params["contentType"] = "text/plain"
		}

		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			params)
		if err != nil {
			return nil, err
		}
		record, err := queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return record, nil
	})
	if err != nil {
		return "", err
	}

	return dbRecord.(*db.Record).Values[0].(string), nil
}

func (r *interactionEventRepository) LinkInteractionEventToSession(ctx context.Context, tenant, interactionEventId, interactionSessionId string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkInteractionEventToSession")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (is:InteractionSession_%s {id:$interactionSessionId}) " +
		" MATCH (ie:InteractionEvent {id:$interactionEventId})" +
		" MERGE (ie)-[:PART_OF]->(is) "
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]interface{}{
				"tenant":               tenant,
				"interactionSessionId": interactionSessionId,
				"interactionEventId":   interactionEventId,
			})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) InteractionEventSentByEmail(ctx context.Context, tenant, interactionEventId, emailId string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.InteractionEventSentByEmail")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (is:InteractionEvent_%s {id:$interactionEventId}) " +
		" MATCH (e:Email_%s {id: $emailId}) " +
		" MERGE (is)-[:SENT_BY]->(e) "
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": interactionEventId,
				"emailId":            emailId,
			})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) InteractionEventSentToEmails(ctx context.Context, tenant, interactionEventId, sentType string, emails []string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.InteractionEventSentToEmails")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := "MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) " +
		" MATCH (e:Email_%s) WHERE e.rawEmail in $emails " +
		" MERGE (ie)-[:SENT_TO {type: $sentType}]->(e) "
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, tenant, tenant),
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": interactionEventId,
				"sentType":           sentType,
				"emails":             emails,
			})
		return nil, err
	})

	return err
}

func (r *interactionEventRepository) LinkInteractionEventAsPartOfByExternalId(ctx context.Context, tenant string, event entity.InteractionEventData) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkInteractionEventAsPartOfByExternalId")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) 
		MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(ext:ExternalSystem {id:$externalSystemId})<-[:IS_LINKED_WITH {externalId:$partOfExternalId}]-(n) 
		WHERE n:Issue
		MERGE (ie)-[result:PART_OF]->(n) 
		return result`
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": event.Id,
				"partOfExternalId":   event.PartOfExternalId,
				"externalSystemId":   event.ExternalSystem,
			})
		if err != nil {
			return nil, err
		}
		_, err = queryResult.Single(ctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (r *interactionEventRepository) LinkInteractionEventWithSenderByExternalId(ctx context.Context, tenant, eventId, externalSystem string, sender entity.InteractionEventParticipant) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkInteractionEventWithSenderByExternalId")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := fmt.Sprintf(`MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) 
		MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(ext:ExternalSystem {id:$externalSystemId})<-[:IS_LINKED_WITH {externalId:$sentByExternalId}]-(n) 
		WHERE ($nodeLabel = '' OR $nodeLabel in labels(n))
		MERGE (ie)-[result:SENT_BY]->(n)`, tenant)
	span.LogFields(log.String("query", query))

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"tenant":             tenant,
			"interactionEventId": eventId,
			"nodeLabel":          sender.GetNodeLabel(),
			"sentByExternalId":   sender.ExternalId,
			"externalSystemId":   externalSystem,
		})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) LinkInteractionEventWithRecipientByExternalId(ctx context.Context, tenant, eventId, externalSystem string, recipient entity.InteractionEventParticipant) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkInteractionEventWithRecipientByExternalId")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)

	query := `MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) 
		MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(ext:ExternalSystem {id:$externalSystemId})<-[:IS_LINKED_WITH {externalId:$sentByExternalId}]-(n) 
		WHERE ($nodeLabel = '' OR $nodeLabel in labels(n))
		MERGE (ie)-[result:SENT_TO]->(n)
		ON CREATE SET result.type=$relationType
		return result`
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, fmt.Sprintf(query, tenant),
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": eventId,
				"sentByExternalId":   recipient.ExternalId,
				"nodeLabel":          recipient.GetNodeLabel(),
				"relationType":       recipient.RelationType,
				"externalSystemId":   externalSystem,
			})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) LinkInteractionEventWithRecipientByOpenlineId(ctx context.Context, tenant, eventId string, recipient entity.InteractionEventParticipant) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkInteractionEventWithRecipientByOpenlineId")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := fmt.Sprintf(`MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) 
		MATCH (n:%s {id:$id}) 
		MERGE (ie)-[result:SENT_TO]->(n)
		ON CREATE SET result.type=$relationType
		return result`, tenant, recipient.GetNodeLabel()+"_"+tenant)
	span.LogFields(log.String("query", query))

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query,
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": eventId,
				"id":                 recipient.OpenlineId,
				"relationType":       recipient.RelationType,
			})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) MergeInteractionSessionForEvent(ctx context.Context, tenant, eventId, externalSource string, syncDate time.Time, session entity.InteractionSession) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.MergeInteractionSessionForEvent")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := fmt.Sprintf(`MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(e:ExternalSystem {id:$externalSystemId}) 
								MATCH (ie:InteractionEvent_%s {id:$interactionEventId})
		 						MERGE (is:InteractionSession_%s)-[rel:IS_LINKED_WITH {externalId:$externalId}]->(e)
								ON CREATE SET
									is:InteractionSession,
									is.id=randomUUID(),
									rel.syncDate=$syncDate,
									is.createdAt=$createdAt,
									is.updatedAt=$createdAt,
									is.source=$source,
									is.sourceOfTruth=$sourceOfTruth,
									is.appSource=$appSource,
									is.identifier=$identifier,
									is.status=$status,
									is.type=$type,
									is.channel=$channel
		WITH is, ie
		MERGE (ie)-[r:PART_OF]->(is)`, tenant, tenant)

	return utils.ExecuteQuery(ctx, *r.driver, query, map[string]interface{}{
		"tenant":             tenant,
		"interactionEventId": eventId,
		"externalId":         session.ExternalId,
		"identifier":         utils.FirstNotEmpty(session.Identifier, session.ExternalId),
		"externalSystemId":   externalSource,
		"source":             externalSource,
		"sourceOfTruth":      externalSource,
		"appSource":          constants.AppSourceSyncCustomerOsData,
		"syncDate":           syncDate,
		"createdAt":          utils.TimePtrFirstNonNilNillableAsAny(session.CreatedAt),
		"channel":            session.Channel,
		"type":               session.Type,
		"status":             session.Status,
	})
}

func (r *interactionEventRepository) FindParticipantByExternalId(ctx context.Context, tenant, externalId, externalSystem string) (*dbtype.Node, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.FindParticipantByExternalId")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := `MATCH (:Tenant {name:$tenant})<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(ext:ExternalSystem {id:$externalSystemId})<-[:IS_LINKED_WITH {externalId:$externalId}]-(n) 
		return n`
	span.LogFields(log.String("query", query))

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)
	result, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, query,
			map[string]interface{}{
				"tenant":           tenant,
				"externalId":       externalId,
				"externalSystemId": externalSystem,
			})
		return utils.ExtractFirstRecordFirstValueAsDbNodePtr(ctx, queryResult, err)
	})
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result.(*dbtype.Node), err
}

func (r *interactionEventRepository) LinkInteractionEventWithSenderJobRole(ctx context.Context, tenant string, interactionEventId, organizationId, contactId string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkInteractionEventWithSenderJobRole")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := fmt.Sprintf(`MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) 
		MATCH (:Organization {id:$organizationId})<-[:ROLE_IN]-(j:JobRole)<-[:WORKS_AS]-(:Contact {id:$contactId}) 
		MERGE (ie)-[result:SENT_BY]->(j)
		return result`, tenant)
	span.LogFields(log.String("query", query))

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query,
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": interactionEventId,
				"organizationId":     organizationId,
				"contactId":          contactId,
			})
		return nil, err
	})
	return err
}

func (r *interactionEventRepository) LinkInteractionEventWithRecipientJobRole(ctx context.Context, tenant string, interactionEventId, organizationId, contactId, relationType string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "InteractionEventRepository.LinkInteractionEventWithRecipientJobRole")
	defer span.Finish()
	tracing.SetDefaultNeo4jRepositorySpanTags(ctx, span)

	query := fmt.Sprintf(`MATCH (ie:InteractionEvent_%s {id:$interactionEventId}) 
		MATCH (:Organization {id:$organizationId})<-[:ROLE_IN]-(j:JobRole)<-[:WORKS_AS]-(:Contact {id:$contactId}) 
		MERGE (ie)-[result:SENT_TO]->(j)
		ON CREATE SET result.type=$relationType
		return result`, tenant)
	span.LogFields(log.String("query", query))

	session := utils.NewNeo4jWriteSession(ctx, *r.driver)
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query,
			map[string]interface{}{
				"tenant":             tenant,
				"interactionEventId": interactionEventId,
				"organizationId":     organizationId,
				"contactId":          contactId,
				"relationType":       relationType,
			})
		return nil, err
	})
	return err
}
