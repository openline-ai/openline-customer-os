package service

import (
	"errors"
	"fmt"
	"github.com/machinebox/graphql"
	c "github.com/openline-ai/openline-customer-os/packages/server/comms-api/config"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
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

	GetInteractionEvent(interactionEventId *string, user *string) (*model.InteractionEventGetResponse, error)
	GetInteractionSession(sessionIdentifier *string, tenant *string, user *string) (*string, error)
}

func (s *customerOSService) GetInteractionEvent(interactionEventId *string, user *string) (*model.InteractionEventGetResponse, error) {
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

	err := s.addHeadersToGraphRequest(graphqlRequest, nil, user)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionEvent: %w", err)
	}

	ctx, cancel, err := s.ContextWithHeaders(nil, user)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionEvent: %w", err)
	}
	defer cancel()

	var graphqlResponse model.InteractionEventGetResponse
	if err := s.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}

	return &graphqlResponse, nil
}

func (s *customerOSService) addHeadersToGraphRequest(req *graphql.Request, tenant *string, user *string) error {
	req.Header.Add("X-Openline-API-KEY", s.conf.Service.CustomerOsAPIKey)
	if user != nil {
		req.Header.Add("X-Openline-USERNAME", *user)
	}
	if tenant != nil {
		req.Header.Add("X-Openline-TENANT", *tenant)
	}

	return nil
}

func (s *customerOSService) CreateInteractionEvent(options ...EventOption) (*model.InteractionEventCreateResponse, error) {
	graphqlRequest := graphql.NewRequest(
		`mutation CreateInteractionEvent(
				$sessionId: ID!, 
				$eventIdentifier: String,
				$channel: String,
				$channelData: String,
				$sentBy: [InteractionEventParticipantInput!]!, 
				$sentTo: [InteractionEventParticipantInput!]!, 
				$appSource: String!, 
				$repliesTo: ID, 
				$content: String, 
				$contentType: String) {
  					interactionEvent_Create(
    					event: {interactionSession: $sessionId, 
								eventIdentifier: $eventIdentifier,
								channel: $channel, 
								channelData: $channelData,
								sentBy: $sentBy, 
								sentTo: $sentTo, 
								appSource: $appSource, 
								repliesTo: $repliesTo, 
								content: $content, 
								contentType: $contentType}
  					) {
						id
						content
						contentType
						createdAt
						channel
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

	params := EventOptions{}
	for _, opt := range options {
		opt(&params)
	}

	graphqlRequest.Var("sessionId", params.sessionId)
	graphqlRequest.Var("eventIdentifier", params.eventIdentifier)
	graphqlRequest.Var("content", params.content)
	graphqlRequest.Var("contentType", params.contentType)
	graphqlRequest.Var("channelData", params.channelData)
	graphqlRequest.Var("channel", params.channel)
	graphqlRequest.Var("sentBy", params.sentBy)
	graphqlRequest.Var("sentTo", params.sentTo)
	graphqlRequest.Var("appSource", params.appSource)

	err := s.addHeadersToGraphRequest(graphqlRequest, params.tenant, params.username)

	if err != nil {
		return nil, fmt.Errorf("error while adding headers to graph request: %w", err)
	}
	ctx, cancel, err := s.ContextWithHeaders(params.tenant, params.username)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionEvent: %w", err)
	}
	defer cancel()

	var graphqlResponse model.InteractionEventCreateResponse
	if err := s.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("CreateInteractionEvent: %w", err)
	}

	return &graphqlResponse, nil
}

func (s *customerOSService) GetInteractionSession(sessionIdentifier *string, tenant *string, user *string) (*string, error) {
	graphqlRequest := graphql.NewRequest(
		`query GetInteractionSession($sessionIdentifier: String!) {
  					interactionSession_BySessionIdentifier(sessionIdentifier: $sessionIdentifier) {
       					id
				}
			}`)

	graphqlRequest.Var("sessionIdentifier", sessionIdentifier)

	err := s.addHeadersToGraphRequest(graphqlRequest, tenant, user)

	if err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}
	ctx, cancel, err := s.ContextWithHeaders(tenant, user)
	if err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := s.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("GetInteractionSession: %w", err)
	}
	id := graphqlResponse["interactionSession_BySessionIdentifier"]["id"]
	return &id, nil
}

func (s *customerOSService) CreateInteractionSession(options ...SessionOption) (*string, error) {
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

	err := s.addHeadersToGraphRequest(graphqlRequest, params.tenant, params.username)

	if err != nil {
		return nil, fmt.Errorf("CreateContactWithPhone: %w", err)
	}

	ctx, cancel, err := s.ContextWithHeaders(params.tenant, params.username)
	if err != nil {
		return nil, fmt.Errorf("CreateInteractionSession: %v", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := s.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("CreateInteractionSession: %w", err)
	}
	id := graphqlResponse["interactionSession_Create"]["id"]
	return &id, nil

}

func (s *customerOSService) CreateAnalysis(options ...AnalysisOption) (*string, error) {
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

	err := s.addHeadersToGraphRequest(graphqlRequest, params.tenant, params.username)

	if err != nil {
		return nil, fmt.Errorf("CreateAnalysis: error while while adding headers to graph request: %w", err)
	}

	ctx, cancel, err := s.ContextWithHeaders(params.tenant, params.username)
	if err != nil {
		return nil, fmt.Errorf("CreateAnalysis: %v", err)
	}
	defer cancel()

	var graphqlResponse map[string]map[string]string
	if err := s.graphqlClient.Run(ctx, graphqlRequest, &graphqlResponse); err != nil {
		return nil, fmt.Errorf("CreateAnalysis: %w", err)
	}
	id := graphqlResponse["analysis_Create"]["id"]
	return &id, nil

}

type EventOptions struct {
	tenant          *string
	username        *string
	sessionId       *string
	eventIdentifier *string
	repliesTo       *string
	content         *string
	contentType     *string
	channel         *string
	channelData     *string
	sentBy          []model.InteractionEventParticipantInput
	sentTo          []model.InteractionEventParticipantInput
	appSource       *string
}

type SessionOptions struct {
	channel           *string
	name              *string
	status            *string
	appSource         *string
	tenant            *string
	username          *string
	sessionIdentifier *string
	sessionType       *string
	attendedBy        []model.InteractionSessionParticipantInput
}

type AnalysisOptions struct {
	analysisType *string
	content      *string
	contentType  *string
	appSource    *string
	tenant       *string
	username     *string
	describes    *model.AnalysisDescriptionInput
}

type EventOption func(*EventOptions)
type SessionOption func(*SessionOptions)
type AnalysisOption func(*AnalysisOptions)

func WithTenant(value *string) EventOption {
	return func(options *EventOptions) {
		options.tenant = value
	}
}

func WithUsername(value *string) EventOption {
	return func(options *EventOptions) {
		options.username = value
	}
}

func WithSessionId(value *string) EventOption {
	return func(options *EventOptions) {
		options.sessionId = value
	}
}

func WithRepliesTo(value *string) EventOption {
	return func(options *EventOptions) {
		options.repliesTo = value
	}
}

func WithContent(value *string) EventOption {
	return func(options *EventOptions) {
		options.content = value
	}
}

func WithContentType(value *string) EventOption {
	return func(options *EventOptions) {
		options.contentType = value
	}
}

func WithChannel(value *string) EventOption {
	return func(options *EventOptions) {
		options.channel = value
	}
}

func WithSentBy(value []model.InteractionEventParticipantInput) EventOption {
	return func(options *EventOptions) {
		options.sentBy = value
	}
}

func WithSentTo(value []model.InteractionEventParticipantInput) EventOption {
	return func(options *EventOptions) {
		options.sentTo = value
	}
}

func WithAppSource(value *string) EventOption {
	return func(options *EventOptions) {
		options.appSource = value
	}
}

func WithSessionIdentifier(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.sessionIdentifier = value
	}
}

func WithSessionChannel(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.channel = value
	}
}

func WithSessionName(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.name = value
	}
}

func WithSessionStatus(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.status = value
	}
}

func WithSessionAttendedBy(value []model.InteractionSessionParticipantInput) SessionOption {
	return func(options *SessionOptions) {
		options.attendedBy = value
	}
}

func WithSessionAppSource(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.appSource = value
	}
}

func WithSessionTenant(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.tenant = value
	}
}

func WithSessionUsername(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.username = value
	}
}

func WithSessionType(value *string) SessionOption {
	return func(options *SessionOptions) {
		options.sessionType = value
	}
}

func WithAnalysisType(value *string) AnalysisOption {
	return func(options *AnalysisOptions) {
		options.analysisType = value
	}
}

func WithAnalysisContent(value *string) AnalysisOption {
	return func(options *AnalysisOptions) {
		options.content = value
	}
}

func WithAnalysisContentType(value *string) AnalysisOption {
	return func(options *AnalysisOptions) {
		options.contentType = value
	}
}

func WithAnalysisAppSource(value *string) AnalysisOption {
	return func(options *AnalysisOptions) {
		options.appSource = value
	}
}

func WithAnalysisTenant(value *string) AnalysisOption {
	return func(options *AnalysisOptions) {
		options.tenant = value
	}
}

func WithAnalysisUsername(value *string) AnalysisOption {
	return func(options *AnalysisOptions) {
		options.username = value
	}
}

func WithAnalysisDescribes(value *model.AnalysisDescriptionInput) AnalysisOption {
	return func(options *AnalysisOptions) {
		options.describes = value
	}
}

func WithEventIdentifier(eventIdentifier string) EventOption {
	return func(options *EventOptions) {
		options.eventIdentifier = &eventIdentifier
	}
}

func WithChannelData(ChannelData *string) EventOption {
	return func(options *EventOptions) {
		options.channelData = ChannelData
	}
}

func (s *customerOSService) ContextWithHeaders(tenant *string, username *string) (context.Context, context.CancelFunc, error) {
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
