package entity

import (
	"fmt"
)

type EmailEntity struct {
	Id      string
	Email   string
	Label   string
	Primary bool
}

func (email EmailEntity) ToString() string {
	return fmt.Sprintf("id: %s\nemail: %s\nlabel: %s", email.Id, email.Email, email.Label)
}

type EmailEntities []EmailEntity
