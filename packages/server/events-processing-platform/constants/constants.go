package constants

const (
	AppSourceEventProcessingPlatform = "event-processing-platform"

	ComponentNeo4jRepository = "neo4jRepository"

	SourceOpenline  = "openline"
	SourceWebscrape = "webscrape"

	PromptType_EmailSummary         = "EmailSummary"
	PromptType_MapIndustry          = "MapIndustryToList"
	PromptType_ExtractIndustryValue = "ExtractIndustryValueFromAiResponse"

	Anthropic = "anthropic"

	NodeLabel_Organization     = "Organization"
	NodeLabel_InteractionEvent = "InteractionEvent"

	TenantKeyHeader = "X-OPENLINE-TENANT-KEY"

	Tcp = "tcp"

	GRPC     = "GRPC"
	SIZE     = "SIZE"
	URI      = "URI"
	STATUS   = "STATUS"
	HTTP     = "HTTP"
	ERROR    = "ERROR"
	METHOD   = "METHOD"
	METADATA = "METADATA"
	REQUEST  = "REQUEST"
	REPLY    = "REPLY"
	TIME     = "TIME"

	Topic        = "topic"
	Partition    = "partition"
	Message      = "message"
	WorkerID     = "workerID"
	Offset       = "offset"
	Time         = "time"
	GroupName    = "GroupName"
	StreamID     = "StreamID"
	EventID      = "EventID"
	EventType    = "EventType"
	EventNumber  = "EventNumber"
	CreatedDate  = "CreatedDate"
	UserMetadata = "UserMetadata"

	Validate        = "validate"
	FieldValidation = "field validation"
	RequiredHeaders = "required header"
	Base64          = "base64"
	Unmarshal       = "unmarshal"
	Uuid            = "uuid"
	Cookie          = "cookie"
	Token           = "token"
	Bcrypt          = "bcrypt"
	Redis           = "redis"

	EsAll = "$all"
)
