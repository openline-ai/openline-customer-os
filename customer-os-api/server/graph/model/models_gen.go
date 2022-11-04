// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type PagedResult interface {
	IsPagedResult()
	GetTotalPages() int
	GetTotalElements() int64
}

type BooleanResult struct {
	Result bool `json:"result"`
}

type CompanyPosition struct {
	CompanyName string  `json:"companyName"`
	JobTitle    *string `json:"jobTitle"`
}

type CompanyPositionInput struct {
	CompanyName string  `json:"companyName"`
	JobTitle    *string `json:"jobTitle"`
}

// Contact - represents one person that can be contacted. In B2C
type Contact struct {
	ID               string             `json:"id"`
	Title            *PersonTitle       `json:"title"`
	FirstName        string             `json:"firstName"`
	LastName         string             `json:"lastName"`
	CreatedAt        time.Time          `json:"createdAt"`
	Label            *string            `json:"label"`
	Notes            *string            `json:"notes"`
	ContactType      *string            `json:"contactType"`
	CompanyPositions []*CompanyPosition `json:"companyPositions"`
	Groups           []*ContactGroup    `json:"groups"`
	TextCustomFields []*TextCustomField `json:"textCustomFields"`
	PhoneNumbers     []*PhoneNumberInfo `json:"phoneNumbers"`
	Emails           []*EmailInfo       `json:"emails"`
}

type ContactGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContactGroupInput struct {
	Name string `json:"name"`
}

type ContactGroupUpdateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContactGroupsPage struct {
	Content       []*ContactGroup `json:"content"`
	TotalPages    int             `json:"totalPages"`
	TotalElements int64           `json:"totalElements"`
}

func (ContactGroupsPage) IsPagedResult()               {}
func (this ContactGroupsPage) GetTotalPages() int      { return this.TotalPages }
func (this ContactGroupsPage) GetTotalElements() int64 { return this.TotalElements }

type ContactInput struct {
	Title            *PersonTitle            `json:"title"`
	FirstName        string                  `json:"firstName"`
	LastName         string                  `json:"lastName"`
	Label            *string                 `json:"label"`
	Notes            *string                 `json:"notes"`
	ContactType      *string                 `json:"contactType"`
	TextCustomFields []*TextCustomFieldInput `json:"textCustomFields"`
	CompanyPosition  *CompanyPositionInput   `json:"companyPosition"`
	Email            *EmailInput             `json:"email"`
	PhoneNumber      *PhoneNumberInput       `json:"phoneNumber"`
}

type ContactUpdateInput struct {
	ID          string       `json:"id"`
	Title       *PersonTitle `json:"title"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	Label       *string      `json:"label"`
	Notes       *string      `json:"notes"`
	ContactType *string      `json:"contactType"`
}

type ContactsPage struct {
	Content       []*Contact `json:"content"`
	TotalPages    int        `json:"totalPages"`
	TotalElements int64      `json:"totalElements"`
}

func (ContactsPage) IsPagedResult()               {}
func (this ContactsPage) GetTotalPages() int      { return this.TotalPages }
func (this ContactsPage) GetTotalElements() int64 { return this.TotalElements }

type EmailInfo struct {
	ID      string     `json:"id"`
	Email   string     `json:"email"`
	Label   EmailLabel `json:"label"`
	Primary bool       `json:"primary"`
}

type EmailInput struct {
	Email   string     `json:"email"`
	Label   EmailLabel `json:"label"`
	Primary *bool      `json:"primary"`
}

type EmailUpdateInput struct {
	ID           string      `json:"id"`
	EmailDetails *EmailInput `json:"emailDetails"`
}

type PaginationFilter struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type PhoneNumberInfo struct {
	ID      string     `json:"id"`
	Number  string     `json:"number"`
	Label   PhoneLabel `json:"label"`
	Primary bool       `json:"primary"`
}

type PhoneNumberInput struct {
	Number  string     `json:"number"`
	Label   PhoneLabel `json:"label"`
	Primary *bool      `json:"primary"`
}

type PhoneNumberUpdateInput struct {
	ID                 string            `json:"id"`
	PhoneNumberDetails *PhoneNumberInput `json:"phoneNumberDetails"`
}

type TenantUser struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type TenantUserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type TenantUsersPage struct {
	Content       []*TenantUser `json:"content"`
	TotalPages    int           `json:"totalPages"`
	TotalElements int64         `json:"totalElements"`
}

func (TenantUsersPage) IsPagedResult()               {}
func (this TenantUsersPage) GetTotalPages() int      { return this.TotalPages }
func (this TenantUsersPage) GetTotalElements() int64 { return this.TotalElements }

type TextCustomField struct {
	ID    string  `json:"id"`
	Group *string `json:"group"`
	Name  string  `json:"name"`
	Value string  `json:"value"`
}

type TextCustomFieldInput struct {
	Group *string `json:"group"`
	Name  string  `json:"name"`
	Value string  `json:"value"`
}

type TextCustomFieldUpdateInput struct {
	ID                     string                `json:"id"`
	TextCustomFieldDetails *TextCustomFieldInput `json:"textCustomFieldDetails"`
}

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

type PersonTitle string

const (
	PersonTitleMr   PersonTitle = "MR"
	PersonTitleMrs  PersonTitle = "MRS"
	PersonTitleMiss PersonTitle = "MISS"
	PersonTitleMs   PersonTitle = "MS"
	PersonTitleDr   PersonTitle = "DR"
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

type PhoneLabel string

const (
	PhoneLabelMain   PhoneLabel = "MAIN"
	PhoneLabelWork   PhoneLabel = "WORK"
	PhoneLabelHome   PhoneLabel = "HOME"
	PhoneLabelMobile PhoneLabel = "MOBILE"
	PhoneLabelOther  PhoneLabel = "OTHER"
)

var AllPhoneLabel = []PhoneLabel{
	PhoneLabelMain,
	PhoneLabelWork,
	PhoneLabelHome,
	PhoneLabelMobile,
	PhoneLabelOther,
}

func (e PhoneLabel) IsValid() bool {
	switch e {
	case PhoneLabelMain, PhoneLabelWork, PhoneLabelHome, PhoneLabelMobile, PhoneLabelOther:
		return true
	}
	return false
}

func (e PhoneLabel) String() string {
	return string(e)
}

func (e *PhoneLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PhoneLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PhoneLabel", str)
	}
	return nil
}

func (e PhoneLabel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
