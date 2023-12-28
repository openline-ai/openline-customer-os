package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repositories struct {
	CommonReadRepository    CommonReadRepository
	ContractReadRepository  ContractReadRepository
	LogEntryWriteRepository LogEntryWriteRepository
	SocialWriteRepository   SocialWriteRepository
	UserReadRepository      UserReadRepository
	UserWriteRepository     UserWriteRepository
}

func InitNeo4jRepositories(driver *neo4j.DriverWithContext, neo4jDatabase string) *Repositories {
	repositories := Repositories{
		CommonReadRepository:    NewCommonReadRepository(driver, neo4jDatabase),
		ContractReadRepository:  NewContractReadRepository(driver, neo4jDatabase),
		LogEntryWriteRepository: NewLogEntryWriteRepository(driver, neo4jDatabase),
		SocialWriteRepository:   NewSocialWriteRepository(driver, neo4jDatabase),
		UserReadRepository:      NewUserReadRepository(driver, neo4jDatabase),
		UserWriteRepository:     NewUserWriteRepository(driver, neo4jDatabase),
	}
	return &repositories
}
