package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/machinebox/graphql"
	c "github.com/openline-ai/openline-customer-os/packages/server/comms-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/model"
	cosModel "github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/graph/model"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

type customerOSService struct {
	graphqlClient *graphql.Client
	conf          *c.Config
}

type CustomerOSService interface {
	CreateAnalysis(options ...AnalysisOption) (*string, error)
	CreateInteractionEvent(options ...EventOption) (*model.InteractionEventCreateResponse, error)
	CreateInteractionSession(options ...SessionOption) (*string, error)
	ForwardQuery(tenant, query *string) ([]byte, error)
	CreateMeeting(options ...MeetingOption) (*string, error)
	GetUserByEmail(email *string) (*string, error)
	GetContactByEmail(user *string, email *string) (*string, error)

	GetTenant(user *string) (*model.TenantResponse, error)
	GetInteractionEvent(interactionEventId *string, user *string) (*model.InteractionEventGetResponse, error)
	GetInteractionSession(sessionIdentifier *string, tenant *string, user *string) (*string, error)
	AddAttachmentToInteractionSession(sessionId string, attachmentId string, tenant *string, user *string) (*string, error)
	AddAttachmentToInteractionEvent(eventId string, attachmentId string, tenant *string, user *string) (*string, error)
}

func (cosService *customerOSService) AddAttachmentToInteractionSession(sessionId string, attachmentId string, tenant *string, user *string) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`mutation AddAttachmentInteractionSession($sessionId: ID!, $attachmentId: ID!) {
				interactionSession_LinkAttachment(
						sessionId: $sessionId,
						attachmentId: $attachmentId
				) {
						id
				}
			}`)

	graphqlRequest.Var("sessionId", sessionId)
	graphqlRequest.Var("attachmentId", attachmentId)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, tenant, user)

	if err != nil {
		return nil, fmt.Errorf("AddAttachmentToInteractionSession: %w", err)
	}
	ctx, cancel, err := cosService.ContextWithHeaders(tenant, user)
	if err != nil {
		return nil, fmt.Errorf("AddAttachmentToInteractionSession: %w", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("AddAttachmentToInteractionSession: %w", err)
	}
	id := graphqlResponse["interactionSession_LinkAttachment"]["id"]
	return &id, nil
}

func (cosService *customerOSService) AddAttachmentToInteractionEvent(eventId string, attachmentId string, tenant *string, user *string) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`mutation AddAttachmentInteractionSession($eventId: ID!, $attachmentId: ID!) {
				interactionEvent_LinkAttachment(
						eventId: $eventId,
						attachmentId: $attachmentId
				) {
						id
				}
			}`)

	graphqlRequest.Var("eventId", eventId)
	graphqlRequest.Var("attachmentId", attachmentId)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, tenant, user)

	if err != nil {
		return nil, fmt.Errorf("AddAttachmentToInteractionEvent: %w", err)
	}
	ctx, cancel, err := cosService.ContextWithHeaders(tenant, user)
	if err != nil {
		return nil, fmt.Errorf("AddAttachmentToInteractionEvent: %w", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("AddAttachmentToInteractionEvent: %w", err)
	}
	id := graphqlResponse["interactionEvent_LinkAttachment"]["id"]
	return &id, nil
}

func (cosService *customerOSService) GetTenant(user *string) (*model.TenantResponse, error) {
	graphqlRequest := graphql.NewRequest(
		`query GetTenant {
			tenant
		}`)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, nil, user)
	if err != nil {
		return nil, fmt.Errorf("GetTenant: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(nil, user)
	if err != nil {
		return nil, fmt.Errorf("GetTenant: %w", err)
	}
	defer cancel()

	var graphqlResponse model.TenantResponse
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetTenant: %w", err)
	}

	return &graphqlResponse, nil
}

func (cosService *customerOSService) GetInteractionEvent(interactionEventId *string, user *string) (*model.InteractionEventGetResponse, error) {
	graphqlRequest := graphql.NewRequest(
		`query GetInteractionEvent($id: ID!) {
			interactionEvent(id: $id) {
				eventIdentifier,
				channelData,
				interactionSession{
					id
					name
				}
			}
		}`)
	graphqlRequest.Var("id", interactionEventId)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, nil, user)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionEvent: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(nil, user)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionEvent: %w", err)
	}
	defer cancel()

	var graphqlResponse model.InteractionEventGetResponse
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}

	return &graphqlResponse, nil
}

func (cosService *customerOSService) addHeadersToGraphRequest(req *graphql.Request, tenant *string, user *string) error {
	req.Header.Add("X-Openline-API-KEY", cosService.conf.Service.CustomerOsAPIKey)
	if user != nil {
		req.Header.Add("X-Openline-USERNAME", *user)
	}
	if tenant != nil {
		req.Header.Add("X-Openline-TENANT", *tenant)
	}

	return nil
}

func (cosService *customerOSService) CreateInteractionEvent(options ...EventOption) (*model.InteractionEventCreateResponse, error) {
	graphqlRequest := graphql.NewRequest(
		`mutation CreateInteractionEvent(
				$sessionId: ID, 
				$meetingId: ID,
				$eventIdentifier: String,
				$channel: String,
				$channelData: String,
				$sentBy: [InteractionEventParticipantInput!]!, 
				$sentTo: [InteractionEventParticipantInput!]!, 
				$appSource: String!, 
				$repliesTo: ID, 
				$content: String, 
				$contentType: String
				$eventType: String,
				$createdAt: Time) {
  					interactionEvent_Create(
    					event: {interactionSession: $sessionId, 
								meetingId: $meetingId,
								eventIdentifier: $eventIdentifier,
								channel: $channel, 
								channelData: $channelData,
								sentBy: $sentBy, 
								sentTo: $sentTo, 
								appSource: $appSource, 
								repliesTo: $repliesTo, 
								content: $content, 
								contentType: $contentType
								eventType: $eventType,
								createdAt: $createdAt}
  					) {
						id
						content
						contentType
						createdAt
						channel
						eventType
						interactionSession {
							name
						}
						sentBy {
						  __typename
						  ... on EmailParticipant {
							emailParticipant {
							  id
							  email
							  contacts {
								id
	                          }
							}
							type
						  }
						  ... on UserParticipant {
							userParticipant {
							  id
							  firstName
							}
							type
						  }
						  ... on PhoneNumberParticipant {
							phoneNumberParticipant {
							  id
							  rawPhoneNumber
							  contacts {
								id
	                          }
							}
							type
						  }
						  ... on ContactParticipant {
							contactParticipant {
							  id
							  firstName
							}
							type
						  }
						}
						sentTo {
						  __typename
						  ... on EmailParticipant {
							emailParticipant {
							  id
							  email
							  contacts {
								id
	                          }
							}
							type
						  }
						  ... on UserParticipant {
							userParticipant {
							  id
							  firstName
							}
							type
						  }
						  ... on PhoneNumberParticipant {
							phoneNumberParticipant {
							  id
							  rawPhoneNumber
							  contacts {
								id
	                          }
							}
							type
						  }
						  ... on ContactParticipant {
							contactParticipant {
							  id
							  firstName
							}
							type
						  }
						}
					  }
					}`)

	params := EventOptions{
		sentTo: []cosModel.InteractionEventParticipantInput{},
		sentBy: []cosModel.InteractionEventParticipantInput{},
	}
	for _, opt := range options {
		opt(&params)
	}

	graphqlRequest.Var("sessionId", params.sessionId)
	graphqlRequest.Var("eventIdentifier", params.eventIdentifier)
	graphqlRequest.Var("content", params.content)
	graphqlRequest.Var("contentType", params.contentType)
	graphqlRequest.Var("channelData", params.channelData)
	graphqlRequest.Var("channel", params.channel)
	graphqlRequest.Var("eventType", params.eventType)
	graphqlRequest.Var("sentBy", params.sentBy)
	graphqlRequest.Var("sentTo", params.sentTo)
	graphqlRequest.Var("appSource", params.appSource)
	graphqlRequest.Var("meetingId", params.meetingId)
	graphqlRequest.Var("createdAt", params.createdAt)

	log.Printf("CreateInteractionEvent: %v", graphqlRequest.Header)
	err := cosService.addHeadersToGraphRequest(graphqlRequest, params.tenant, params.username)

	if err != nil {
		return nil, fmt.Errorf("error while adding headers to graph request: %w", err)
	}
	ctx, cancel, err := cosService.ContextWithHeaders(params.tenant, params.username)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionEvent: %w", err)
	}
	defer cancel()

	var graphqlResponse model.InteractionEventCreateResponse
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("CreateInteractionEvent: %w", err)
	}

	return &graphqlResponse, nil
}

func (cosService *customerOSService) GetInteractionSession(sessionIdentifier *string, tenant *string, user *string) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`query GetInteractionSession($sessionIdentifier: String!) {
  					interactionSession_BySessionIdentifier(sessionIdentifier: $sessionIdentifier) {
       					id
				}
			}`)

	graphqlRequest.Var("sessionIdentifier", sessionIdentifier)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, tenant, user)

	if err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}
	ctx, cancel, err := cosService.ContextWithHeaders(tenant, user)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}
	id := graphqlResponse["interactionSession_BySessionIdentifier"]["id"]
	return &id, nil
}

func (cosService *customerOSService) CreateInteractionSession(options ...SessionOption) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`mutation CreateInteractionSession($sessionIdentifier: String, $channel: String, $name: String!, $type: String, $status: String!, $appSource: String!, $attendedBy: [InteractionSessionParticipantInput!]) {
				interactionSession_Create(
				session: {
					sessionIdentifier: $sessionIdentifier
        			channel: $channel
        			name: $name
        			status: $status
					type: $type
        			appSource: $appSource
                    attendedBy: $attendedBy
    			}
  			) {
				id
   			}
		}
	`)

	params := SessionOptions{}
	for _, opt := range options {
		opt(&params)
	}

	graphqlRequest.Var("sessionIdentifier", params.sessionIdentifier)
	graphqlRequest.Var("channel", params.channel)
	graphqlRequest.Var("name", params.name)
	graphqlRequest.Var("status", params.status)
	graphqlRequest.Var("appSource", params.appSource)
	graphqlRequest.Var("attendedBy", params.attendedBy)
	graphqlRequest.Var("type", params.sessionType)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, params.tenant, params.username)

	if err != nil {
		return nil, fmt.Errorf("CreateContactWithPhone: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(params.tenant, params.username)
	if err != nil {
		return nil, fmt.Errorf("CreateInteractionSession: %v", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("CreateMeeting: %w", err)
	}
	id := graphqlResponse["meeting_Create"]["id"]
	return &id, nil

}

func (cosService *customerOSService) GetUserByEmail(email *string) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`query GetUserByEmail($email: String!){ user_ByEmail(email: $email) { id } }`)

	graphqlRequest.Var("email", *email)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, nil, email)

	if err != nil {
		return nil, fmt.Errorf("user_ByEmail: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(nil, email)
	if err != nil {
		return nil, fmt.Errorf("user_ByEmail: %v", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("user_ByEmail: %w", err)
	}
	id := graphqlResponse["user_ByEmail"]["id"]
	return &id, nil
}

func (cosService *customerOSService) GetContactByEmail(user *string, email *string) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`query GetUserByEmail($email: String!){ contact_ByEmail(email: $email) { id } }`)

	graphqlRequest.Var("email", *email)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, nil, user)

	if err != nil {
		return nil, fmt.Errorf("add headers contact_ByEmail: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(nil, user)
	if err != nil {
		return nil, fmt.Errorf("context contact_ByEmail: %v", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("contact_ByEmail: %w", err)
	}
	id := graphqlResponse["contact_ByEmail"]["id"]
	return &id, nil
}

func (cosService *customerOSService) ForwardQuery(tenant, query *string) ([]byte, error) {
	graphqlRequest := graphql.NewRequest(*query)

	err := cosService.addHeadersToGraphRequest(graphqlRequest, tenant, nil)

	if err != nil {
		return nil, fmt.Errorf("ForwardQuery: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(tenant, nil)
	if err != nil {
		return nil, fmt.Errorf("ForwardQuery: %v", err)
	}
	defer cancel()

	var graphqlResponse interface{}
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("ForwardQuery: %w", err)
	}

	// Encode to JSON first to escape special characters
	normalized, _ := json.Marshal(graphqlResponse)

	// Decode again to convert escaped chars back to original bytes
	var result interface{}
	json.Unmarshal(normalized, &result)

	// Convert result back to JSON
	jsonBytes, _ := json.Marshal(result)

	return jsonBytes, nil
}

func (cosService *customerOSService) CreateMeeting(options ...MeetingOption) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`mutation CreateMeeting($name: String!, $startedAt: Time!, $endedAt: Time!, $appSource: String!, $createdBy: [MeetingParticipantInput!], $attendedBy: [MeetingParticipantInput!], $noteInput: NoteInput!, $externalSystem: ExternalSystemReferenceInput!) {
  				meeting_Create( meeting: {name: $name, startedAt: $startedAt, endedAt: $endedAt, appSource: $appSource, createdBy: $createdBy, attendedBy: $attendedBy, note: $noteInput, externalSystem: $externalSystem}
			) {
				id
       			name
       			source
				startedAt
				endedAt
				attendedBy {
					__typename
					... on UserParticipant {
						userParticipant {
							id
							firstName
						}
					}
					... on ContactParticipant {
						contactParticipant {
							id
							firstName
						}
					}
					... on OrganizationParticipant {
						organizationParticipant {
							id
							name
						}
					}
				}
				createdBy {
					__typename
						... on UserParticipant {
								userParticipant {
									id
									firstName
							   	}
						   	}
						   	... on ContactParticipant {
								   contactParticipant {
									   	id
										firstName
							   		}
						   	}
					   	}
				note {
					id
					html
					createdAt
					updatedAt
					appSource
					sourceOfTruth
				}
				createdAt
				updatedAt
				appSource
				sourceOfTruth
			}
		}`)

	params := MeetingOptions{}
	for _, opt := range options {
		opt(&params)
	}

	graphqlRequest.Var("name", params.name)
	graphqlRequest.Var("startedAt", params.startedAt)
	graphqlRequest.Var("endedAt", params.endedAt)
	graphqlRequest.Var("appSource", params.appSource)
	graphqlRequest.Var("attendedBy", params.attendedBy)
	graphqlRequest.Var("createdBy", params.createdBy)
	graphqlRequest.Var("noteInput", params.noteInput)
	graphqlRequest.Var("externalSystem", params.externalSystem)
	err := cosService.addHeadersToGraphRequest(graphqlRequest, params.tenant, params.username)

	if err != nil {
		return nil, fmt.Errorf("addHeadersToGraphRequest: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(params.tenant, params.username)
	if err != nil {
		return nil, fmt.Errorf("ContextWithHeaders: %v", err)
	}
	defer cancel()
	var graphqlResponse model.CreateMeetingResponse
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {

		log.Printf("graphqlResponse: %v", err)
		return nil, fmt.Errorf("meeting_Create: %w", err)
	}

	return &graphqlResponse.MeetingCreate.Id, nil
}

func (cosService *customerOSService) CreateAnalysis(options ...AnalysisOption) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`mutation CreateAnalysis($content: String, $contentType: String, $analysisType: String, $appSource: String!, $describes: [AnalysisDescriptionInput!]!) {
				analysis_Create(
					analysis: {
						content: $content
						contentType: $contentType
						analysisType: $analysisType
						describes: $describes
						appSource: $appSource
					}
				  ) {
					  id
				}
			}
	`)

	params := AnalysisOptions{}
	for _, opt := range options {
		opt(&params)
	}

	graphqlRequest.Var("content", params.content)
	graphqlRequest.Var("contentType", params.contentType)
	graphqlRequest.Var("analysisType", params.analysisType)
	graphqlRequest.Var("appSource", params.appSource)

	if params.describes != nil {
		graphqlRequest.Var("describes", params.describes)
	}

	err := cosService.addHeadersToGraphRequest(graphqlRequest, params.tenant, params.username)

	if err != nil {
		return nil, fmt.Errorf("CreateAnalysis: error while while adding headers to graph request: %w", err)
	}

	ctx, cancel, err := cosService.ContextWithHeaders(params.tenant, params.username)
	if err != nil {
		return nil, fmt.Errorf("CreateAnalysis: %v", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := cosService.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("CreateAnalysis: %w", err)
	}
	id := graphqlResponse["analysis_Create"]["id"]
	return &id, nil

}

func (cosService *customerOSService) ContextWithHeaders(tenant *string, username *string) (context.Context, context.CancelFunc, error) {
	if tenant == nil && username == nil {
		return nil, nil, errors.New("no username and no tenant specified")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if tenant != nil {
		ctx = metadata.AppendToOutgoingContext(ctx, "X-Openline-TENANT`", *tenant)
	}

	if username != nil {
		ctx = metadata.AppendToOutgoingContext(ctx, "X-Openline-USERNAME`", *username)
	}
	return ctx, cancel, nil
}

func NewCustomerOSService(graphqlClient *graphql.Client, config *c.Config) CustomerOSService {
	return &customerOSService{
		graphqlClient: graphqlClient,
		conf:          config,
	}
}
