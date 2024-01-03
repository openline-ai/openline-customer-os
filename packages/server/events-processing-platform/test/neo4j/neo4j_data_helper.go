package neo4j

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	neo4jtest "github.com/openline-ai/customer-os-neo4j-repository/test"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/graph_db/entity"
)

func CleanupAllData(ctx context.Context, driver *neo4j.DriverWithContext) {
	neo4jtest.ExecuteWriteQuery(ctx, driver, `MATCH (n) DETACH DELETE n`, map[string]any{})
}

func CreateTenant(ctx context.Context, driver *neo4j.DriverWithContext, tenant string) {
	query := `MERGE (t:Tenant {name:$tenant})`
	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant": tenant,
	})
}

func CreateOrganization(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, organization entity.OrganizationEntity) string {
	orgId := organization.ID
	if orgId == "" {
		orgId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization {id:$id})
				SET o:Organization_%s,
					o.name=$name,
					o.hide=$hide
				`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant": tenant,
		"name":   organization.Name,
		"hide":   organization.Hide,
		"id":     orgId,
	})
	return orgId
}

func CreateUser(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, user entity.UserEntity) string {
	userId := user.Id
	if userId == "" {
		userId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:USER_BELONGS_TO_TENANT]-(u:User {id:$id})
				SET u:User_%s,
					u.firstName=$firstName,
					u.lastName=$lastName,
					u.roles=$roles
				`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":    tenant,
		"id":        userId,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"roles":     user.Roles,
	})
	return userId
}

func CreateSocial(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, social entity.SocialEntity) string {
	socialId := utils.NewUUIDIfEmpty(social.Id)
	query := fmt.Sprintf(`MERGE (s:Social:Social_%s {id: $id})
				SET s.url=$url,
					s.platformName=$platformName
				`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":           socialId,
		"url":          social.Url,
		"platformName": social.PlatformName,
	})
	return socialId
}

func CreateContact(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, contact entity.ContactEntity) string {
	contactId := contact.Id
	if contactId == "" {
		contactId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:CONTACT_BELONGS_TO_TENANT]-(c:Contact {id:$id})
				SET c:Contact_%s,
					c.firstName=$firstName,
					c.lastName=$lastName
				`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":    tenant,
		"id":        contactId,
		"firstName": contact.FirstName,
		"lastName":  contact.LastName,
	})
	return contactId
}

func CreateJobRole(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, jobRole entity.JobRoleEntity) string {
	jobRoleId := jobRole.Id
	if jobRoleId == "" {
		jobRoleId = uuid.New().String()
	}
	query := fmt.Sprintf(`CREATE (jobRole:JobRole:JobRole_%s {id:$jobRoleId})`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"jobRoleId": jobRoleId,
	})
	return jobRoleId

}

func CreateLogEntryForOrg(ctx context.Context, driver *neo4j.DriverWithContext, tenant, orgId string, logEntry entity.LogEntryEntity) string {
	logEntryId := logEntry.Id
	if logEntryId == "" {
		logEntryId = uuid.New().String()
	}
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})<-[:ORGANIZATION_BELONGS_TO_TENANT]-(o:Organization {id:$orgId})
			  MERGE (o)-[:LOGGED]->(l:LogEntry {id:$id})
				SET l:LogEntry_%s,
					l:TimelineEvent,
					l:TimelineEvent_%s,
					l.content=$content,
					l.contentType=$contentType,
					l.startedAt=$startedAt
				`, tenant, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":      tenant,
		"orgId":       orgId,
		"id":          logEntryId,
		"content":     logEntry.Content,
		"contentType": logEntry.ContentType,
		"startedAt":   logEntry.StartedAt,
	})
	return logEntryId
}

func CreateIssue(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, issue entity.IssueEntity) string {
	issueId := utils.NewUUIDIfEmpty(issue.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:ISSUE_BELONGS_TO_TENANT]-(i:Issue {id:$id})
				SET i:Issue_%s,
					i:TimelineEvent,
					i:TimelineEvent_%s,
					i.subject=$subject,
					i.status=$status,
					i.priority=$priority,
					i.description=$description,
					i.source=$source,
					i.sourceOfTruth=$sourceOfTruth
				`, tenant, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":        tenant,
		"id":            issueId,
		"subject":       issue.Subject,
		"status":        issue.Status,
		"priority":      issue.Priority,
		"description":   issue.Description,
		"source":        issue.Source,
		"sourceOfTruth": issue.SourceOfTruth,
	})
	return issueId
}

func CreateComment(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, comment entity.CommentEntity) string {
	commentId := utils.NewUUIDIfEmpty(comment.Id)
	query := fmt.Sprintf(`MERGE (c:Comment:Comment_%s {id:$id})
				SET c.content=$content,
					c.contentType=$contentType,
					c.source=$source,
					c.sourceOfTruth=$sourceOfTruth
				`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":            commentId,
		"content":       comment.Content,
		"contentType":   comment.ContentType,
		"source":        comment.Source,
		"sourceOfTruth": comment.SourceOfTruth,
	})
	return commentId
}

func CreateContract(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, contract entity.ContractEntity) string {
	contractId := utils.NewUUIDIfEmpty(contract.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})
				MERGE (t)<-[:CONTRACT_BELONGS_TO_TENANT]-(c:Contract {id:$id})
				SET 
					c:Contract_%s,
					c.name=$name,
					c.contractUrl=$contractUrl,
					c.source=$source,
					c.sourceOfTruth=$sourceOfTruth,
					c.status=$status,
					c.renewalCycle=$renewalCycle,
					c.renewalPeriods=$renewalPeriods,
					c.signedAt=$signedAt,
					c.serviceStartedAt=$serviceStartedAt,
					c.endedAt=$endedAt
				`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":               contractId,
		"tenant":           tenant,
		"name":             contract.Name,
		"contractUrl":      contract.ContractUrl,
		"source":           contract.Source,
		"sourceOfTruth":    contract.SourceOfTruth,
		"status":           contract.Status,
		"renewalCycle":     contract.RenewalCycle,
		"renewalPeriods":   contract.RenewalPeriods,
		"signedAt":         utils.TimePtrFirstNonNilNillableAsAny(contract.SignedAt),
		"serviceStartedAt": utils.TimePtrFirstNonNilNillableAsAny(contract.ServiceStartedAt),
		"endedAt":          utils.TimePtrFirstNonNilNillableAsAny(contract.EndedAt),
	})
	return contractId
}

func CreateContractForOrganization(ctx context.Context, driver *neo4j.DriverWithContext, tenant, organizationId string, contract entity.ContractEntity) string {
	contractId := CreateContract(ctx, driver, tenant, contract)
	LinkContractWithOrganization(ctx, driver, contractId, organizationId)
	return contractId
}

func LinkContractWithOrganization(ctx context.Context, driver *neo4j.DriverWithContext, contractId, organizationId string) {
	query := `MATCH (c:Contract {id:$contractId}), (o:Organization {id:$organizationId})
				MERGE (o)-[:HAS_CONTRACT]->(c) `
	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"contractId":     contractId,
		"organizationId": organizationId,
	})
}

func CreateOpportunity(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, opportunity entity.OpportunityEntity) string {
	opportunityId := utils.NewUUIDIfEmpty(opportunity.Id)
	query := fmt.Sprintf(`
				MERGE (op:Opportunity {id:$id})
				SET 
					op:Opportunity_%s,
					op.name=$name,
					op.source=$source,
					op.sourceOfTruth=$sourceOfTruth,
					op.internalStage=$internalStage,
					op.internalType=$internalType,
					op.renewedAt=$renewedAt,
					op.amount=$amount,
					op.maxAmount=$maxAmount,
					op.renewalLikelihood=$renewalLikelihood,
					op.renewalUpdatedByUserId=$renewalUpdatedByUserId,
					op.comments=$comments
				`, tenant)

	if opportunity.InternalType == "RENEWAL" {
		query += `, op:RenewalOpportunity`
	}

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":                     opportunityId,
		"name":                   opportunity.Name,
		"source":                 opportunity.Source,
		"sourceOfTruth":          opportunity.SourceOfTruth,
		"internalStage":          opportunity.InternalStage,
		"internalType":           opportunity.InternalType,
		"amount":                 opportunity.Amount,
		"maxAmount":              opportunity.MaxAmount,
		"renewedAt":              utils.TimePtrFirstNonNilNillableAsAny(opportunity.RenewalDetails.RenewedAt),
		"renewalLikelihood":      opportunity.RenewalDetails.RenewalLikelihood,
		"renewalUpdatedByUserId": opportunity.RenewalDetails.RenewalUpdatedByUserId,
		"comments":               opportunity.Comments,
	})
	return opportunityId
}

func LinkContractWithOpportunity(ctx context.Context, driver *neo4j.DriverWithContext, contractId, opportunityId string, renewal bool) {
	query := `MATCH (c:Contract {id:$contractId}), (o:Opportunity {id:$opportunityId})
				MERGE (c)-[:HAS_OPPORTUNITY]->(o) `
	if renewal {
		query += `MERGE (c)-[:ACTIVE_RENEWAL]->(o)`
	}
	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"contractId":    contractId,
		"opportunityId": opportunityId,
	})
}

func CreateServiceLineItemForContract(ctx context.Context, driver *neo4j.DriverWithContext, tenant, contractId string, serviceLineItem entity.ServiceLineItemEntity) string {
	serviceLineItemId := utils.NewUUIDIfEmpty(serviceLineItem.Id)
	query := fmt.Sprintf(`MATCH (c:Contract {id:$contractId})
				MERGE (c)-[:HAS_SERVICE]->(sli:ServiceLineItem {id:$id})
				ON CREATE SET 
					sli:ServiceLineItem_%s,
					sli.name=$name,
					sli.price=$price,
					sli.quantity=$quantity,
					sli.billed=$billed,
					sli.comments=$comments
				`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"contractId": contractId,
		"id":         serviceLineItemId,
		"price":      serviceLineItem.Price,
		"name":       serviceLineItem.Name,
		"quantity":   serviceLineItem.Quantity,
		"billed":     serviceLineItem.Billed,
		"comments":   serviceLineItem.Comments,
	})
	return serviceLineItemId
}

func CreatePhoneNumber(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, phoneNumber entity.PhoneNumberEntity) string {
	phoneNumberId := utils.NewUUIDIfEmpty(phoneNumber.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:PHONE_NUMBER_BELONGS_TO_TENANT]-(i:PhoneNumber {id:$id})
				SET i:PhoneNumber_%s,
					i.e164=$e164,
					i.validated=$validated,
					i.rawPhoneNumber=$rawPhoneNumber,
					i.source=$source,
					i.sourceOfTruth=$sourceOfTruth,
					i.appSource=$appSource,
					i.createdAt=$createdAt,
					i.updatedAt=$updatedAt`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":         tenant,
		"id":             phoneNumberId,
		"e164":           phoneNumber.E164,
		"validated":      phoneNumber.Validated,
		"rawPhoneNumber": phoneNumber.RawPhoneNumber,
		"source":         phoneNumber.Source,
		"sourceOfTruth":  phoneNumber.SourceOfTruth,
		"appSource":      phoneNumber.AppSource,
		"createdAt":      phoneNumber.CreatedAt,
		"updatedAt":      phoneNumber.UpdatedAt,
	})
	return phoneNumberId
}

func CreateEmail(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, entity entity.EmailEntity) string {
	emailId := utils.NewUUIDIfEmpty(entity.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name:$tenant})
								MERGE (e:Email {id:$emailId})
								MERGE (e)-[:EMAIL_ADDRESS_BELONGS_TO_TENANT]->(t)
								ON CREATE SET e:Email_%s,
									e.email=$email,
									e.rawEmail=$rawEmail,
									e.isReachable=$isReachable,
									e.createdAt=$createdAt,
									e.updatedAt=$updatedAt
							`, tenant)
	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":      tenant,
		"emailId":     emailId,
		"email":       entity.Email,
		"rawEmail":    entity.RawEmail,
		"isReachable": entity.IsReachable,
		"createdAt":   entity.CreatedAt,
		"updatedAt":   entity.UpdatedAt,
	})
	return emailId
}

func CreateEmailForUser(ctx context.Context, driver *neo4j.DriverWithContext, tenant, userId string, entity entity.EmailEntity) string {
	emailId := CreateEmail(ctx, driver, tenant, entity)
	LinkEmailWithUser(ctx, driver, emailId, userId)
	return emailId
}

func LinkEmailWithUser(ctx context.Context, driver *neo4j.DriverWithContext, emailId, userId string) {
	query := `MATCH (e:Email {id:$emailId}), (u:User {id:$userId})
				MERGE (u)-[:HAS]->(e) `
	ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"emailId": emailId,
		"userId":  userId,
	})
}

func CreateLocation(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, location entity.LocationEntity) string {
	locationId := utils.NewUUIDIfEmpty(location.Id)
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:LOCATION_BELONGS_TO_TENANT]-(i:Location {id:$id})
				SET i:Location_%s,
					i.name=$name,
					i.createdAt=$createdAt,
					i.updatedAt=$updatedAt,
					i.country=$country,
					i.region=$region,    
					i.locality=$locality,    
					i.address=$address,    
					i.address2=$address2,    
					i.zip=$zip,    
					i.addressType=$addressType,    
					i.houseNumber=$houseNumber,    
					i.postalCode=$postalCode,    
					i.plusFour=$plusFour,    
					i.commercial=$commercial,    
					i.predirection=$predirection,    
					i.district=$district,    
					i.street=$street,    
					i.rawAddress=$rawAddress,    
					i.latitude=$latitude,    
					i.longitude=$longitude,    
					i.timeZone=$timeZone,    
					i.utcOffset=$utcOffset,    
					i.sourceOfTruth=$sourceOfTruth,
					i.source=$source,
					i.appSource=$appSource`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":        tenant,
		"id":            locationId,
		"name":          location.Name,
		"createdAt":     location.CreatedAt,
		"updatedAt":     location.UpdatedAt,
		"country":       location.Country,
		"region":        location.Region,
		"locality":      location.Locality,
		"address":       location.Address,
		"address2":      location.Address2,
		"zip":           location.Zip,
		"addressType":   location.AddressType,
		"houseNumber":   location.HouseNumber,
		"postalCode":    location.PostalCode,
		"plusFour":      location.PlusFour,
		"commercial":    location.Commercial,
		"predirection":  location.Predirection,
		"district":      location.District,
		"street":        location.Street,
		"rawAddress":    location.RawAddress,
		"latitude":      location.Latitude,
		"longitude":     location.Longitude,
		"timeZone":      location.TimeZone,
		"utcOffset":     location.UtcOffset,
		"sourceOfTruth": location.SourceOfTruth,
		"source":        location.Source,
		"appSource":     location.AppSource,
	})
	return locationId

}

func LinkIssueReportedBy(ctx context.Context, driver *neo4j.DriverWithContext, issueId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (i:Issue {id:$issueId})
				MERGE (i)-[:REPORTED_BY]->(e)`

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"issueId":  issueId,
		"entityId": entityId,
	})
}

func LinkIssueAssignedTo(ctx context.Context, driver *neo4j.DriverWithContext, issueId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (i:Issue {id:$issueId})
				MERGE (i)-[:ASSIGNED_TO]->(e)`

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"issueId":  issueId,
		"entityId": entityId,
	})
}

func LinkIssueFollowedBy(ctx context.Context, driver *neo4j.DriverWithContext, issueId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (i:Issue {id:$issueId})
				MERGE (i)-[:FOLLOWED_BY]->(e)`

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"issueId":  issueId,
		"entityId": entityId,
	})
}

func CreateInteractionEvent(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, interactionEvent entity.InteractionEventEntity) string {
	interactionEventId := utils.NewUUIDIfEmpty(interactionEvent.Id)
	query := fmt.Sprintf(`MERGE (i:InteractionEvent {id:$id})
				SET i:InteractionEvent_%s,
					i:TimelineEvent,
					i:TimelineEvent_%s,
					i.content=$content,
					i.contentType=$contentType,
					i.channel=$channel,
					i.channelData=$channelData,
					i.identifier=$identifier,
					i.eventType=$eventType,
					i.source=$source,
					i.sourceOfTruth=$sourceOfTruth
				`, tenant, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"id":            interactionEventId,
		"content":       interactionEvent.Content,
		"contentType":   interactionEvent.ContentType,
		"channel":       interactionEvent.Channel,
		"channelData":   interactionEvent.ChannelData,
		"identifier":    interactionEvent.Identifier,
		"eventType":     interactionEvent.EventType,
		"source":        interactionEvent.Source,
		"sourceOfTruth": interactionEvent.SourceOfTruth,
	})
	return interactionEventId
}

func CreateTag(ctx context.Context, driver *neo4j.DriverWithContext, tenant string, tag entity.TagEntity) string {
	tagId := tag.Id
	if tagId == "" {
		tagId = uuid.New().String()
	}

	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:TAG_BELONGS_TO_TENANT]-(tag:Tag {id:$tagId})
				SET tag:Tag_%s,
					tag.name=$name`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant": tenant,
		"tagId":  tagId,
		"name":   tag.Name,
	})
	return tagId
}

func LinkTag(ctx context.Context, driver *neo4j.DriverWithContext, tagId, entityId string) {

	query := `MATCH (e {id:$entityId})
				MATCH (t:Tag {id:$tagId})
				MERGE (e)-[rel:TAGGED]->(t)
				SET rel.taggedAt=$now`

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tagId":    tagId,
		"entityId": entityId,
		"now":      utils.Now(),
	})
}

func LinkSocial(ctx context.Context, driver *neo4j.DriverWithContext, socialId, entityId string) {
	query := `MATCH (e {id:$entityId})
				MATCH (s:Social {id:$socialId})
				MERGE (e)-[:HAS]->(s)`

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"socialId": socialId,
		"entityId": entityId,
	})
}

func CreateExternalSystem(ctx context.Context, driver *neo4j.DriverWithContext, tenant, externalSystem string) {
	query := fmt.Sprintf(`MATCH (t:Tenant {name: $tenant})
			  MERGE (t)<-[:EXTERNAL_SYSTEM_BELONGS_TO_TENANT]-(ext:ExternalSystem {id:$externalSystemId})
				ON CREATE SET ext:ExternalSystem_%s`, tenant)

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":           tenant,
		"externalSystemId": externalSystem,
	})
}

func CreateWorkspace(ctx context.Context, driver *neo4j.DriverWithContext, workspace string, provider string, tenant string) {
	query := `MATCH (t:Tenant {name: $tenant})
			  MERGE (t)-[:HAS_WORKSPACE]->(w:Workspace {name:$workspace, provider:$provider})`

	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"tenant":    tenant,
		"provider":  provider,
		"workspace": workspace,
	})
}

func CreateCountry(ctx context.Context, driver *neo4j.DriverWithContext, codeA2, codeA3, name, phoneCode string) {
	query := `MERGE (c:Country{codeA3: $codeA3}) 
				ON CREATE SET 
					c.phoneCode = $phoneCode,
					c.codeA2 = $codeA2,
					c.name = $name, 
					c.createdAt = $now, 
					c.updatedAt = $now`
	neo4jtest.ExecuteWriteQuery(ctx, driver, query, map[string]any{
		"codeA2":    codeA2,
		"codeA3":    codeA3,
		"phoneCode": phoneCode,
		"name":      name,
		"now":       utils.Now(),
	})
}
