// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

// Contact - represents one person that can be contacted for a Customer. In B2C
type Contact struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	CreatedAt   time.Time `json:"createdAt"`
	Label       *string   `json:"label"`
	ContactType *string   `json:"contactType"`
}

type ContactGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ContactGroupInput struct {
	Name string `json:"name"`
}

type ContactInput struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Label       *string `json:"label"`
	ContactType *string `json:"contactType"`
}
