package model

type Output struct {
	Skip                     bool     `json:"skip,omitempty"`
	SkipReason               string   `json:"skipReason,omitempty"`
	Id                       string   `json:"id,omitempty"`
	Name                     string   `json:"name,omitempty"`
	FirstName                string   `json:"firstName,omitempty"`
	LastName                 string   `json:"lastName,omitempty"`
	Prefix                   string   `json:"prefix,omitempty"`
	Email                    string   `json:"email,omitempty"`
	AdditionalEmails         []string `json:"additionalEmails,omitempty"`
	PhoneNumber              string   `json:"phoneNumber,omitempty"`
	AdditionalPhoneNumbers   []string `json:"additionalPhoneNumbers,omitempty"`
	CreatedAt                string   `json:"createdAt,omitempty"`
	UpdatedAt                string   `json:"updatedAt,omitempty"`
	ExternalId               string   `json:"externalId,omitempty"`
	ExternalOwnerId          string   `json:"externalOwnerId,omitempty"`
	ExternalUserId           string   `json:"externalUserId,omitempty"`
	ExternalCreatorId        string   `json:"externalCreatorId,omitempty"`
	ExternalOrganizationId   string   `json:"externalOrganizationId,omitempty"`
	ExternalUrl              string   `json:"externalUrl,omitempty"`
	ExternalSourceTable      string   `json:"externalSourceTable,omitempty"`
	ExternalSystem           string   `json:"externalSystem,omitempty"`
	ExternalSyncId           string   `json:"externalSyncId,omitempty"`
	ExternalContactsIds      []string `json:"externalContactIds,omitempty"`
	ExternalOrganizationsIds []string `json:"externalOrganizationsIds,omitempty"`
	Description              string   `json:"description,omitempty"`
	Domains                  []string `json:"domains,omitempty"`
	Notes                    []struct {
		FieldSource string `json:"fieldSource,omitempty"`
		Note        string `json:"note,omitempty"`
	} `json:"notes,omitempty"`
	Website            string `json:"website,omitempty"`
	Industry           string `json:"industry,omitempty"`
	IsPublic           bool   `json:"isPublic,omitempty"`
	Employees          int    `json:"employees,omitempty"`
	IndustryGroup      string `json:"industryGroup,omitempty"`
	TargetAudience     string `json:"targetAudience,omitempty"`
	ValueProposition   string `json:"valueProposition,omitempty"`
	Market             string `json:"market,omitempty"`
	LastFundingRound   string `json:"lastFundingRound,omitempty"`
	LastFundingAmount  string `json:"lastFundingAmount,omitempty"`
	LocationName       string `json:"locationName,omitempty"`
	Country            string `json:"country,omitempty"`
	Region             string `json:"region,omitempty"`
	Locality           string `json:"locality,omitempty"`
	Address            string `json:"address,omitempty"`
	Address2           string `json:"address2,omitempty"`
	Zip                string `json:"zip,omitempty"`
	RelationshipName   string `json:"relationshipName,omitempty"`
	RelationshipStage  string `json:"relationshipStage,omitempty"`
	ParentOrganization struct {
		ExternalId           string `json:"externalId,omitempty"`
		OrganizationRelation string `json:"organizationRelation,omitempty"`
		Type                 string `json:"type,omitempty"`
	} `json:"parentOrganization,omitempty"`
	Html             string   `json:"html,omitempty"`
	Text             string   `json:"text,omitempty"`
	Subject          string   `json:"subject,omitempty"`
	MentionedTags    []string `json:"mentionedTags,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	StartedAt        string   `json:"startedAt,omitempty"`
	EndedAt          string   `json:"endedAt,omitempty"`
	Agenda           string   `json:"agenda,omitempty"`
	ContentType      string   `json:"contentType,omitempty"`
	Location         string   `json:"location,omitempty"`
	ConferenceUrl    string   `json:"conferenceUrl,omitempty"`
	MeetingUrl       string   `json:"meetingUrl,omitempty"`
	FromEmail        string   `json:"fromEmail,omitempty"`
	ToEmail          []string `json:"toEmail,omitempty"`
	CcEmail          []string `json:"ccEmail,omitempty"`
	BccEmail         []string `json:"bccEmail,omitempty"`
	Direction        string   `json:"direction,omitempty"`
	MessageId        string   `json:"messageId,omitempty"`
	ThreadId         string   `json:"threadId,omitempty"`
	Label            string   `json:"label,omitempty"`
	JobTitle         string   `json:"jobTitle,omitempty"`
	TextCustomFields []struct {
		Name           string `json:"name,omitempty"`
		Value          string `json:"value,omitempty"`
		ExternalSystem string `json:"externalSystem,omitempty"`
		CreatedAt      string `json:"createdAt,omitempty"`
	} `json:"textCustomFields,omitempty"`
	Status                         string   `json:"status,omitempty"`
	Priority                       string   `json:"priority,omitempty"`
	CollaboratorUserExternalIds    []string `json:"collaboratorUserExternalIds,omitempty"`
	FollowerUserExternalIds        []string `json:"followerUserExternalIds,omitempty"`
	ReporterOrganizationExternalId string   `json:"reporterOrganizationExternalId,omitempty"`
	AssigneeUserExternalId         string   `json:"assigneeUserExternalId,omitempty"`
	MentionedIssueExternalId       string   `json:"mentionedIssueExternalId,omitempty"`
	Content                        string   `json:"content,omitempty"`
	Type                           string   `json:"type,omitempty"`
	Channel                        string   `json:"channel,omitempty"`
	PartOfExternalId               string   `json:"partOfExternalId,omitempty"`
	PartOfSession                  struct {
		ExternalId string `json:"externalId,omitempty"`
		Channel    string `json:"channel,omitempty"`
		Type       string `json:"type,omitempty"`
		CreatedAt  string `json:"createdAt,omitempty"`
		Status     string `json:"status,omitempty"`
		Identifier string `json:"identifier,omitempty"`
	} `json:"partOfSession,omitempty"`
	SentBy struct {
		ExternalId      string `json:"externalId,omitempty"`
		ParticipantType string `json:"participantType,omitempty"`
		RelationType    string `json:"relationType,omitempty"`
	} `json:"sentBy,omitempty"`

	SentTo []struct {
		ExternalId      string `json:"externalId,omitempty"`
		ParticipantType string `json:"participantType,omitempty"`
		RelationType    string `json:"relationType,omitempty"`
	} `json:"sentTo,omitempty"`
	Timezone               string `json:"timezone,omitempty"`
	OpenlineOrganizationId string `json:"openlineOrganizationId,omitempty"`
	Identifier             string `json:"identifier,omitempty"`
}
