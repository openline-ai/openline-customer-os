// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Action interface {
	IsAction()
}

type ExtensibleEntity interface {
	IsNode()
	IsExtensibleEntity()
	GetID() string
	GetTemplate() *EntityTemplate
}

type Node interface {
	IsNode()
	GetID() string
}

// Describes the number of pages and total elements included in a query response.
// **A `response` object.**
type Pages interface {
	IsPages()
	// The total number of pages included in the query response.
	// **Required.**
	GetTotalPages() int
	// The total number of elements included in the query response.
	// **Required.**
	GetTotalElements() int64
}

type Address struct {
	ID       string      `json:"id"`
	Country  *string     `json:"country"`
	State    *string     `json:"state"`
	City     *string     `json:"city"`
	Address  *string     `json:"address"`
	Address2 *string     `json:"address2"`
	Zip      *string     `json:"zip"`
	Phone    *string     `json:"phone"`
	Fax      *string     `json:"fax"`
	Source   *DataSource `json:"source"`
}

// A contact represents an individual in customerOS.
// **A `response` object.**
type Contact struct {
	// The unique ID associated with the contact in customerOS.
	// **Required**
	ID string `json:"id"`
	// The title associate with the contact in customerOS.
	Title *PersonTitle `json:"title"`
	// The first name of the contact in customerOS.
	FirstName *string `json:"firstName"`
	// The last name of the contact in customerOS.
	LastName *string `json:"lastName"`
	// An ISO8601 timestamp recording when the contact was created in customerOS.
	// **Required**
	CreatedAt time.Time `json:"createdAt"`
	// A user-defined label applied against a contact in customerOS.
	Label *string `json:"label"`
	// Readonly indicator for a contact
	Readonly bool       `json:"readonly"`
	Source   DataSource `json:"source"`
	// User-defined field that defines the relationship type the contact has with your business.  `Customer`, `Partner`, `Lead` are examples.
	ContactType *ContactType `json:"contactType"`
	// `organizationName` and `jobTitle` of the contact if it has been associated with an organization.
	// **Required.  If no values it returns an empty array.**
	Roles []*ContactRole `json:"roles"`
	// Identifies any contact groups the contact is associated with.
	//  **Required.  If no values it returns an empty array.**
	Groups []*ContactGroup `json:"groups"`
	// All phone numbers associated with a contact in customerOS.
	// **Required.  If no values it returns an empty array.**
	PhoneNumbers []*PhoneNumber `json:"phoneNumbers"`
	// All email addresses assocaited with a contact in customerOS.
	// **Required.  If no values it returns an empty array.**
	Emails []*Email `json:"emails"`
	// All addresses associated with a contact in customerOS.
	// **Required.  If no values it returns an empty array.**
	Addresses []*Address `json:"addresses"`
	// User defined metadata appended to the contact record in customerOS.
	// **Required.  If no values it returns an empty array.**
	CustomFields []*CustomField `json:"customFields"`
	FieldSets    []*FieldSet    `json:"fieldSets"`
	// Template of the contact in customerOS.
	Template *EntityTemplate `json:"template"`
	// Contact owner (user)
	Owner *User `json:"owner"`
	// Contact notes
	Notes         *NotePage         `json:"notes"`
	Conversations *ConversationPage `json:"conversations"`
	Actions       []Action          `json:"actions"`
}

func (Contact) IsExtensibleEntity()               {}
func (this Contact) GetID() string                { return this.ID }
func (this Contact) GetTemplate() *EntityTemplate { return this.Template }

func (Contact) IsNode() {}

// A collection of groups that a Contact belongs to.  Groups are user-defined entities.
// **A `return` object.**
type ContactGroup struct {
	// The unique ID associated with the `ContactGroup` in customerOS.
	// **Required**
	ID string `json:"id"`
	// The name of the `ContactGroup`.
	// **Required**
	Name     string        `json:"name"`
	Source   DataSource    `json:"source"`
	Contacts *ContactsPage `json:"contacts"`
}

// Create a groups that can be associated with a `Contact` in customerOS.
// **A `create` object.**
type ContactGroupInput struct {
	// The name of the `ContactGroup`.
	// **Required**
	Name string `json:"name"`
}

// Specifies how many pages of `ContactGroup` information has been returned in the query response.
// **A `response` object.**
type ContactGroupPage struct {
	// A collection of groups that a Contact belongs to.  Groups are user-defined entities.
	// **Required.  If no values it returns an empty array.**
	Content []*ContactGroup `json:"content"`
	// Total number of pages in the query response.
	// **Required.**
	TotalPages int `json:"totalPages"`
	// Total number of elements in the query response.
	// **Required.**
	TotalElements int64 `json:"totalElements"`
}

func (ContactGroupPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this ContactGroupPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this ContactGroupPage) GetTotalElements() int64 { return this.TotalElements }

// Update a group that can be associated with a `Contact` in customerOS.
// **A `update` object.**
type ContactGroupUpdateInput struct {
	// The unique ID associated with the `ContactGroup` in customerOS.
	// **Required**
	ID string `json:"id"`
	// The name of the `ContactGroup`.
	// **Required**
	Name string `json:"name"`
}

// Create an individual in customerOS.
// **A `create` object.**
type ContactInput struct {
	// The unique ID associated with the template of the contact in customerOS.
	TemplateID *string `json:"templateId"`
	// The title of the contact.
	Title *PersonTitle `json:"title"`
	// The first name of the contact.
	FirstName *string `json:"firstName"`
	// The last name of the contact.
	LastName *string `json:"lastName"`
	// A user-defined label attached to contact.
	Label *string `json:"label"`
	// User-defined field that defines the relationship type the contact has with your business.  `Customer`, `Partner`, `Lead` are examples.
	ContactTypeID *string `json:"contactTypeId"`
	// Readonly indicator for a contact
	Readonly *bool `json:"readonly"`
	// An ISO8601 timestamp recording when the contact was created in customerOS.
	CreatedAt *time.Time `json:"createdAt"`
	// User defined metadata appended to contact.
	// **Required.**
	CustomFields []*CustomFieldInput `json:"customFields"`
	FieldSets    []*FieldSetInput    `json:"fieldSets"`
	// An email addresses associted with the contact.
	Email *EmailInput `json:"email"`
	// A phone number associated with the contact.
	PhoneNumber *PhoneNumberInput `json:"phoneNumber"`
	// Id of the contact owner (user)
	OwnerID           *string                       `json:"ownerId"`
	ExternalReference *ExternalSystemReferenceInput `json:"externalReference"`
}

// Describes the relationship a Contact has with a Organization.
// **A `return` object**
type ContactRole struct {
	ID string `json:"id"`
	// Organization associated with a Contact.
	// **Required.**
	Organization *Organization `json:"organization"`
	// The Contact's job title.
	JobTitle *string    `json:"jobTitle"`
	Primary  bool       `json:"primary"`
	Source   DataSource `json:"source"`
}

// Describes the relationship a Contact has with an Organization.
// **A `create` object**
type ContactRoleInput struct {
	OrganizationID *string `json:"organizationId"`
	// The Contact's job title.
	JobTitle *string `json:"jobTitle"`
	Primary  *bool   `json:"primary"`
}

type ContactType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContactTypeInput struct {
	Name string `json:"name"`
}

type ContactTypeUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Updates data fields associated with an existing customer record in customerOS.
// **An `update` object.**
type ContactUpdateInput struct {
	// The unique ID associated with the contact in customerOS.
	// **Required.**
	ID string `json:"id"`
	// The title associate with the contact in customerOS.
	Title *PersonTitle `json:"title"`
	// The first name of the contact in customerOS.
	FirstName *string `json:"firstName"`
	// The last name of the contact in customerOS.
	LastName *string `json:"lastName"`
	// A user-defined label applied against a contact in customerOS.
	Label *string `json:"label"`
	// User-defined field that defines the relationship type the contact has with your business.  `Customer`, `Partner`, `Lead` are examples.
	ContactTypeID *string `json:"contactTypeId"`
	// Id of the contact owner (user)
	OwnerID *string `json:"ownerId"`
	// Readonly indicator for a contact
	Readonly *bool `json:"readonly"`
}

// Specifies how many pages of contact information has been returned in the query response.
// **A `response` object.**
type ContactsPage struct {
	// A contact entity in customerOS.
	// **Required.  If no values it returns an empty array.**
	Content []*Contact `json:"content"`
	// Total number of pages in the query response.
	// **Required.**
	TotalPages int `json:"totalPages"`
	// Total number of elements in the query response.
	// **Required.**
	TotalElements int64 `json:"totalElements"`
}

func (ContactsPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this ContactsPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this ContactsPage) GetTotalElements() int64 { return this.TotalElements }

type Conversation struct {
	ID           string             `json:"id"`
	StartedAt    time.Time          `json:"startedAt"`
	EndedAt      *time.Time         `json:"endedAt"`
	Status       ConversationStatus `json:"status"`
	Channel      *string            `json:"channel"`
	MessageCount int64              `json:"messageCount"`
	Contacts     []*Contact         `json:"contacts"`
	Users        []*User            `json:"users"`
}

func (Conversation) IsNode()            {}
func (this Conversation) GetID() string { return this.ID }

type ConversationInput struct {
	ID         *string            `json:"id"`
	StartedAt  *time.Time         `json:"startedAt"`
	ContactIds []string           `json:"contactIds"`
	UserIds    []string           `json:"userIds"`
	Status     ConversationStatus `json:"status"`
	Channel    *string            `json:"channel"`
}

type ConversationPage struct {
	Content       []*Conversation `json:"content"`
	TotalPages    int             `json:"totalPages"`
	TotalElements int64           `json:"totalElements"`
}

func (ConversationPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this ConversationPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this ConversationPage) GetTotalElements() int64 { return this.TotalElements }

type ConversationUpdateInput struct {
	ID                        string              `json:"id"`
	ContactIds                []string            `json:"contactIds"`
	UserIds                   []string            `json:"userIds"`
	Status                    *ConversationStatus `json:"status"`
	Channel                   *string             `json:"channel"`
	SkipMessageCountIncrement bool                `json:"skipMessageCountIncrement"`
}

// Describes a custom, user-defined field associated with a `Contact`.
// **A `return` object.**
type CustomField struct {
	// The unique ID associated with the custom field.
	// **Required**
	ID string `json:"id"`
	// The name of the custom field.
	// **Required**
	Name string `json:"name"`
	// Datatype of the custom field.
	// **Required**
	Datatype CustomFieldDataType `json:"datatype"`
	// The value of the custom field.
	// **Required**
	Value AnyTypeValue `json:"value"`
	// The source of the custom field value
	Source    DataSource           `json:"source"`
	CreatedAt time.Time            `json:"createdAt"`
	Template  *CustomFieldTemplate `json:"template"`
}

func (CustomField) IsNode()            {}
func (this CustomField) GetID() string { return this.ID }

// Describes a custom, user-defined field associated with a `Contact` of type String.
// **A `create` object.**
type CustomFieldInput struct {
	// The unique ID associated with the custom field.
	ID *string `json:"id"`
	// The name of the custom field.
	// **Required**
	Name string `json:"name"`
	// Datatype of the custom field.
	// **Required**
	Datatype CustomFieldDataType `json:"datatype"`
	// The value of the custom field.
	// **Required**
	Value      AnyTypeValue `json:"value"`
	TemplateID *string      `json:"templateId"`
}

type CustomFieldTemplate struct {
	ID        string                  `json:"id"`
	Name      string                  `json:"name"`
	Type      CustomFieldTemplateType `json:"type"`
	Order     int                     `json:"order"`
	Mandatory bool                    `json:"mandatory"`
	Length    *int                    `json:"length"`
	Min       *int                    `json:"min"`
	Max       *int                    `json:"max"`
}

func (CustomFieldTemplate) IsNode()            {}
func (this CustomFieldTemplate) GetID() string { return this.ID }

type CustomFieldTemplateInput struct {
	Name      string                  `json:"name"`
	Type      CustomFieldTemplateType `json:"type"`
	Order     int                     `json:"order"`
	Mandatory bool                    `json:"mandatory"`
	Length    *int                    `json:"length"`
	Min       *int                    `json:"min"`
	Max       *int                    `json:"max"`
}

// Describes a custom, user-defined field associated with a `Contact`.
// **An `update` object.**
type CustomFieldUpdateInput struct {
	// The unique ID associated with the custom field.
	// **Required**
	ID string `json:"id"`
	// The name of the custom field.
	// **Required**
	Name string `json:"name"`
	// Datatype of the custom field.
	// **Required**
	Datatype CustomFieldDataType `json:"datatype"`
	// The value of the custom field.
	// **Required**
	Value AnyTypeValue `json:"value"`
}

// Describes an email address associated with a `Contact` in customerOS.
// **A `return` object.**
type Email struct {
	// The unique ID associated with the contact in customerOS.
	// **Required**
	ID string `json:"id"`
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	Email string `json:"email"`
	// Describes the type of email address (WORK, PERSONAL, etc).
	Label *EmailLabel `json:"label"`
	// Identifies whether the email address is primary or not.
	// **Required.**
	Primary bool       `json:"primary"`
	Source  DataSource `json:"source"`
}

// Describes an email address associated with a `Contact` in customerOS.
// **A `create` object.**
type EmailInput struct {
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	Email string `json:"email"`
	// Describes the type of email address (WORK, PERSONAL, etc).
	Label *EmailLabel `json:"label"`
	// Identifies whether the email address is primary or not.
	// **Required.**
	Primary *bool `json:"primary"`
}

// Describes an email address associated with a `Contact` in customerOS.
// **An `update` object.**
type EmailUpdateInput struct {
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	ID string `json:"id"`
	// An email address assocaited with the contact in customerOS.
	// **Required.**
	Email string `json:"email"`
	// Describes the type of email address (WORK, PERSONAL, etc).
	Label *EmailLabel `json:"label"`
	// Identifies whether the email address is primary or not.
	// **Required.**
	Primary *bool `json:"primary"`
}

type EntityTemplate struct {
	ID           string                   `json:"id"`
	Version      int                      `json:"version"`
	Name         string                   `json:"name"`
	Extends      *EntityTemplateExtension `json:"extends"`
	FieldSets    []*FieldSetTemplate      `json:"fieldSets"`
	CustomFields []*CustomFieldTemplate   `json:"customFields"`
	CreatedAt    time.Time                `json:"createdAt"`
}

func (EntityTemplate) IsNode()            {}
func (this EntityTemplate) GetID() string { return this.ID }

type EntityTemplateInput struct {
	Name         string                      `json:"name"`
	Extends      *EntityTemplateExtension    `json:"extends"`
	FieldSets    []*FieldSetTemplateInput    `json:"fieldSets"`
	CustomFields []*CustomFieldTemplateInput `json:"customFields"`
}

type ExternalSystemReferenceInput struct {
	ID       string             `json:"id"`
	SyncDate *time.Time         `json:"syncDate"`
	Type     ExternalSystemType `json:"type"`
}

type FieldSet struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	CreatedAt    time.Time         `json:"createdAt"`
	CustomFields []*CustomField    `json:"customFields"`
	Template     *FieldSetTemplate `json:"template"`
	Source       DataSource        `json:"source"`
}

type FieldSetInput struct {
	ID           *string             `json:"id"`
	Name         string              `json:"name"`
	CustomFields []*CustomFieldInput `json:"customFields"`
	TemplateID   *string             `json:"templateId"`
}

type FieldSetTemplate struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Order        int                    `json:"order"`
	CustomFields []*CustomFieldTemplate `json:"customFields"`
}

func (FieldSetTemplate) IsNode()            {}
func (this FieldSetTemplate) GetID() string { return this.ID }

type FieldSetTemplateInput struct {
	Name         string                      `json:"name"`
	Order        int                         `json:"order"`
	CustomFields []*CustomFieldTemplateInput `json:"customFields"`
}

type FieldSetUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Filter struct {
	Not    *Filter     `json:"NOT"`
	And    []*Filter   `json:"AND"`
	Or     []*Filter   `json:"OR"`
	Filter *FilterItem `json:"filter"`
}

type FilterItem struct {
	Property      string             `json:"property"`
	Operation     ComparisonOperator `json:"operation"`
	Value         AnyTypeValue       `json:"value"`
	CaseSensitive *bool              `json:"caseSensitive"`
}

type Note struct {
	ID        string     `json:"id"`
	HTML      string     `json:"html"`
	CreatedAt time.Time  `json:"createdAt"`
	CreatedBy *User      `json:"createdBy"`
	Source    DataSource `json:"source"`
}

type NoteInput struct {
	HTML string `json:"html"`
}

type NotePage struct {
	Content       []*Note `json:"content"`
	TotalPages    int     `json:"totalPages"`
	TotalElements int64   `json:"totalElements"`
}

func (NotePage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this NotePage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this NotePage) GetTotalElements() int64 { return this.TotalElements }

type NoteUpdateInput struct {
	ID   string `json:"id"`
	HTML string `json:"html"`
}

type Organization struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Description      *string           `json:"description"`
	Domain           *string           `json:"domain"`
	Website          *string           `json:"website"`
	Industry         *string           `json:"industry"`
	IsPublic         *bool             `json:"isPublic"`
	CreatedAt        time.Time         `json:"createdAt"`
	Readonly         *bool             `json:"readonly"`
	OrganizationType *OrganizationType `json:"organizationType"`
	// All addresses associated with an organization in customerOS.
	// **Required.  If no values it returns an empty array.**
	Addresses []*Address `json:"addresses"`
	Source    DataSource `json:"source"`
}

func (Organization) IsNode()            {}
func (this Organization) GetID() string { return this.ID }

type OrganizationInput struct {
	// The name of the organization.
	// **Required.**
	Name               string  `json:"name"`
	Description        *string `json:"description"`
	Domain             *string `json:"domain"`
	Website            *string `json:"website"`
	Industry           *string `json:"industry"`
	IsPublic           *bool   `json:"isPublic"`
	OrganizationTypeID *string `json:"organizationTypeId"`
}

type OrganizationPage struct {
	Content       []*Organization `json:"content"`
	TotalPages    int             `json:"totalPages"`
	TotalElements int64           `json:"totalElements"`
}

func (OrganizationPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this OrganizationPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this OrganizationPage) GetTotalElements() int64 { return this.TotalElements }

type OrganizationType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OrganizationTypeInput struct {
	Name string `json:"name"`
}

type OrganizationTypeUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PageViewAction struct {
	ID             string    `json:"id"`
	StartedAt      time.Time `json:"startedAt"`
	EndedAt        time.Time `json:"endedAt"`
	PageTitle      string    `json:"pageTitle"`
	PageURL        string    `json:"pageUrl"`
	Application    string    `json:"application"`
	SessionID      string    `json:"sessionId"`
	OrderInSession int64     `json:"orderInSession"`
	EngagedTime    int64     `json:"engagedTime"`
}

func (PageViewAction) IsAction() {}

func (PageViewAction) IsNode()            {}
func (this PageViewAction) GetID() string { return this.ID }

// If provided as part of the request, results will be filtered down to the `page` and `limit` specified.
type Pagination struct {
	// The results page to return in the response.
	// **Required.**
	Page int `json:"page"`
	// The maximum number of results in the response.
	// **Required.**
	Limit int `json:"limit"`
}

// Describes a phone number associated with a `Contact` in customerOS.
// **A `return` object.**
type PhoneNumber struct {
	// The unique ID associated with the phone number.
	// **Required**
	ID string `json:"id"`
	// The phone number in e164 format.
	// **Required**
	E164 string `json:"e164"`
	// Defines the type of phone number.
	Label *PhoneNumberLabel `json:"label"`
	// Determines if the phone number is primary or not.
	// **Required**
	Primary bool       `json:"primary"`
	Source  DataSource `json:"source"`
}

// Describes a phone number associated with a `Contact` in customerOS.
// **A `create` object.**
type PhoneNumberInput struct {
	// The phone number in e164 format.
	// **Required**
	E164 string `json:"e164"`
	// Defines the type of phone number.
	Label *PhoneNumberLabel `json:"label"`
	// Determines if the phone number is primary or not.
	// **Required**
	Primary *bool `json:"primary"`
}

// Describes a phone number associated with a `Contact` in customerOS.
// **An `update` object.**
type PhoneNumberUpdateInput struct {
	// The unique ID associated with the phone number.
	// **Required**
	ID string `json:"id"`
	// The phone number in e164 format.
	// **Required**
	E164 string `json:"e164"`
	// Defines the type of phone number.
	Label *PhoneNumberLabel `json:"label"`
	// Determines if the phone number is primary or not.
	// **Required**
	Primary *bool `json:"primary"`
}

// Describes the success or failure of the GraphQL call.
// **A `return` object**
type Result struct {
	// The result of the GraphQL call.
	// **Required.**
	Result bool `json:"result"`
}

type SortBy struct {
	By            string           `json:"by"`
	Direction     SortingDirection `json:"direction"`
	CaseSensitive *bool            `json:"caseSensitive"`
}

// Describes the User of customerOS.  A user is the person who logs into the Openline platform.
// **A `return` object**
type User struct {
	// The unique ID associated with the customerOS user.
	// **Required**
	ID string `json:"id"`
	// The first name of the customerOS user.
	// **Required**
	FirstName string `json:"firstName"`
	// The last name of the customerOS user.
	// **Required**
	LastName string `json:"lastName"`
	// The email address of the customerOS user.
	// **Required**
	Email string `json:"email"`
	// Timestamp of user creation.
	// **Required**
	CreatedAt     time.Time         `json:"createdAt"`
	Source        DataSource        `json:"source"`
	Conversations *ConversationPage `json:"conversations"`
}

// Describes the User of customerOS.  A user is the person who logs into the Openline platform.
// **A `create` object.**
type UserInput struct {
	// The first name of the customerOS user.
	// **Required**
	FirstName string `json:"firstName"`
	// The last name of the customerOS user.
	// **Required**
	LastName string `json:"lastName"`
	// The email address of the customerOS user.
	// **Required**
	Email string `json:"email"`
}

// Specifies how many pages of `User` information has been returned in the query response.
// **A `return` object.**
type UserPage struct {
	// A `User` entity in customerOS.
	// **Required.  If no values it returns an empty array.**
	Content []*User `json:"content"`
	// Total number of pages in the query response.
	// **Required.**
	TotalPages int `json:"totalPages"`
	// Total number of elements in the query response.
	// **Required.**
	TotalElements int64 `json:"totalElements"`
}

func (UserPage) IsPages() {}

// The total number of pages included in the query response.
// **Required.**
func (this UserPage) GetTotalPages() int { return this.TotalPages }

// The total number of elements included in the query response.
// **Required.**
func (this UserPage) GetTotalElements() int64 { return this.TotalElements }

type ActionType string

const (
	ActionTypePageView ActionType = "PAGE_VIEW"
)

var AllActionType = []ActionType{
	ActionTypePageView,
}

func (e ActionType) IsValid() bool {
	switch e {
	case ActionTypePageView:
		return true
	}
	return false
}

func (e ActionType) String() string {
	return string(e)
}

func (e *ActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionType", str)
	}
	return nil
}

func (e ActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ComparisonOperator string

const (
	ComparisonOperatorEq       ComparisonOperator = "EQ"
	ComparisonOperatorContains ComparisonOperator = "CONTAINS"
)

var AllComparisonOperator = []ComparisonOperator{
	ComparisonOperatorEq,
	ComparisonOperatorContains,
}

func (e ComparisonOperator) IsValid() bool {
	switch e {
	case ComparisonOperatorEq, ComparisonOperatorContains:
		return true
	}
	return false
}

func (e ComparisonOperator) String() string {
	return string(e)
}

func (e *ComparisonOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ComparisonOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ComparisonOperator", str)
	}
	return nil
}

func (e ComparisonOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ConversationStatus string

const (
	ConversationStatusActive ConversationStatus = "ACTIVE"
	ConversationStatusClosed ConversationStatus = "CLOSED"
)

var AllConversationStatus = []ConversationStatus{
	ConversationStatusActive,
	ConversationStatusClosed,
}

func (e ConversationStatus) IsValid() bool {
	switch e {
	case ConversationStatusActive, ConversationStatusClosed:
		return true
	}
	return false
}

func (e ConversationStatus) String() string {
	return string(e)
}

func (e *ConversationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConversationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConversationStatus", str)
	}
	return nil
}

func (e ConversationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CustomFieldDataType string

const (
	CustomFieldDataTypeText     CustomFieldDataType = "TEXT"
	CustomFieldDataTypeBool     CustomFieldDataType = "BOOL"
	CustomFieldDataTypeDatetime CustomFieldDataType = "DATETIME"
	CustomFieldDataTypeInteger  CustomFieldDataType = "INTEGER"
	CustomFieldDataTypeDecimal  CustomFieldDataType = "DECIMAL"
)

var AllCustomFieldDataType = []CustomFieldDataType{
	CustomFieldDataTypeText,
	CustomFieldDataTypeBool,
	CustomFieldDataTypeDatetime,
	CustomFieldDataTypeInteger,
	CustomFieldDataTypeDecimal,
}

func (e CustomFieldDataType) IsValid() bool {
	switch e {
	case CustomFieldDataTypeText, CustomFieldDataTypeBool, CustomFieldDataTypeDatetime, CustomFieldDataTypeInteger, CustomFieldDataTypeDecimal:
		return true
	}
	return false
}

func (e CustomFieldDataType) String() string {
	return string(e)
}

func (e *CustomFieldDataType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CustomFieldDataType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CustomFieldDataType", str)
	}
	return nil
}

func (e CustomFieldDataType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CustomFieldTemplateType string

const (
	CustomFieldTemplateTypeText CustomFieldTemplateType = "TEXT"
)

var AllCustomFieldTemplateType = []CustomFieldTemplateType{
	CustomFieldTemplateTypeText,
}

func (e CustomFieldTemplateType) IsValid() bool {
	switch e {
	case CustomFieldTemplateTypeText:
		return true
	}
	return false
}

func (e CustomFieldTemplateType) String() string {
	return string(e)
}

func (e *CustomFieldTemplateType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CustomFieldTemplateType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CustomFieldTemplateType", str)
	}
	return nil
}

func (e CustomFieldTemplateType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DataSource string

const (
	DataSourceNa       DataSource = "NA"
	DataSourceOpenline DataSource = "OPENLINE"
	DataSourceHubspot  DataSource = "HUBSPOT"
	DataSourceZendesk  DataSource = "ZENDESK"
)

var AllDataSource = []DataSource{
	DataSourceNa,
	DataSourceOpenline,
	DataSourceHubspot,
	DataSourceZendesk,
}

func (e DataSource) IsValid() bool {
	switch e {
	case DataSourceNa, DataSourceOpenline, DataSourceHubspot, DataSourceZendesk:
		return true
	}
	return false
}

func (e DataSource) String() string {
	return string(e)
}

func (e *DataSource) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataSource(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DataSource", str)
	}
	return nil
}

func (e DataSource) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Describes the type of email address (WORK, PERSONAL, etc).
// **A `return` object.
type EmailLabel string

const (
	EmailLabelMain  EmailLabel = "MAIN"
	EmailLabelWork  EmailLabel = "WORK"
	EmailLabelHome  EmailLabel = "HOME"
	EmailLabelOther EmailLabel = "OTHER"
)

var AllEmailLabel = []EmailLabel{
	EmailLabelMain,
	EmailLabelWork,
	EmailLabelHome,
	EmailLabelOther,
}

func (e EmailLabel) IsValid() bool {
	switch e {
	case EmailLabelMain, EmailLabelWork, EmailLabelHome, EmailLabelOther:
		return true
	}
	return false
}

func (e EmailLabel) String() string {
	return string(e)
}

func (e *EmailLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EmailLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EmailLabel", str)
	}
	return nil
}

func (e EmailLabel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EntityTemplateExtension string

const (
	EntityTemplateExtensionContact EntityTemplateExtension = "CONTACT"
)

var AllEntityTemplateExtension = []EntityTemplateExtension{
	EntityTemplateExtensionContact,
}

func (e EntityTemplateExtension) IsValid() bool {
	switch e {
	case EntityTemplateExtensionContact:
		return true
	}
	return false
}

func (e EntityTemplateExtension) String() string {
	return string(e)
}

func (e *EntityTemplateExtension) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntityTemplateExtension(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EntityTemplateExtension", str)
	}
	return nil
}

func (e EntityTemplateExtension) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ExternalSystemType string

const (
	ExternalSystemTypeHubspot ExternalSystemType = "HUBSPOT"
	ExternalSystemTypeZendesk ExternalSystemType = "ZENDESK"
)

var AllExternalSystemType = []ExternalSystemType{
	ExternalSystemTypeHubspot,
	ExternalSystemTypeZendesk,
}

func (e ExternalSystemType) IsValid() bool {
	switch e {
	case ExternalSystemTypeHubspot, ExternalSystemTypeZendesk:
		return true
	}
	return false
}

func (e ExternalSystemType) String() string {
	return string(e)
}

func (e *ExternalSystemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ExternalSystemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ExternalSystemType", str)
	}
	return nil
}

func (e ExternalSystemType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The honorific title of an individual.
// **A `response` object.**
type PersonTitle string

const (
	// For men, regardless of marital status.
	PersonTitleMr PersonTitle = "MR"
	// For married women.
	PersonTitleMrs PersonTitle = "MRS"
	// For girls, unmarried women, and married women who continue to use their maiden name.
	PersonTitleMiss PersonTitle = "MISS"
	// For women, regardless of marital status, or when marital status is unknown.
	PersonTitleMs PersonTitle = "MS"
	// For the holder of a doctoral degree.
	PersonTitleDr PersonTitle = "DR"
)

var AllPersonTitle = []PersonTitle{
	PersonTitleMr,
	PersonTitleMrs,
	PersonTitleMiss,
	PersonTitleMs,
	PersonTitleDr,
}

func (e PersonTitle) IsValid() bool {
	switch e {
	case PersonTitleMr, PersonTitleMrs, PersonTitleMiss, PersonTitleMs, PersonTitleDr:
		return true
	}
	return false
}

func (e PersonTitle) String() string {
	return string(e)
}

func (e *PersonTitle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PersonTitle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PersonTitle", str)
	}
	return nil
}

func (e PersonTitle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Defines the type of phone number.
// **A `response` object. **
type PhoneNumberLabel string

const (
	PhoneNumberLabelMain   PhoneNumberLabel = "MAIN"
	PhoneNumberLabelWork   PhoneNumberLabel = "WORK"
	PhoneNumberLabelHome   PhoneNumberLabel = "HOME"
	PhoneNumberLabelMobile PhoneNumberLabel = "MOBILE"
	PhoneNumberLabelOther  PhoneNumberLabel = "OTHER"
)

var AllPhoneNumberLabel = []PhoneNumberLabel{
	PhoneNumberLabelMain,
	PhoneNumberLabelWork,
	PhoneNumberLabelHome,
	PhoneNumberLabelMobile,
	PhoneNumberLabelOther,
}

func (e PhoneNumberLabel) IsValid() bool {
	switch e {
	case PhoneNumberLabelMain, PhoneNumberLabelWork, PhoneNumberLabelHome, PhoneNumberLabelMobile, PhoneNumberLabelOther:
		return true
	}
	return false
}

func (e PhoneNumberLabel) String() string {
	return string(e)
}

func (e *PhoneNumberLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PhoneNumberLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PhoneNumberLabel", str)
	}
	return nil
}

func (e PhoneNumberLabel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortingDirection string

const (
	SortingDirectionAsc  SortingDirection = "ASC"
	SortingDirectionDesc SortingDirection = "DESC"
)

var AllSortingDirection = []SortingDirection{
	SortingDirectionAsc,
	SortingDirectionDesc,
}

func (e SortingDirection) IsValid() bool {
	switch e {
	case SortingDirectionAsc, SortingDirectionDesc:
		return true
	}
	return false
}

func (e SortingDirection) String() string {
	return string(e)
}

func (e *SortingDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortingDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortingDirection", str)
	}
	return nil
}

func (e SortingDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
