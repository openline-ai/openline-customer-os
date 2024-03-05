// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Role string

const (
	RoleAdmin                   Role = "ADMIN"
	RoleCustomerOsPlatformOwner Role = "CUSTOMER_OS_PLATFORM_OWNER"
	RoleOwner                   Role = "OWNER"
	RoleUser                    Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleCustomerOsPlatformOwner,
	RoleOwner,
	RoleUser,
}

func (e Role) String() string {
	return string(e)
}