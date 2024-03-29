package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Dbs struct {
	Neo4jDriver *neo4j.DriverWithContext
}

type Repositories struct {
	Dbs             Dbs
	Neo4jRepository Neo4jRepository
}

func InitRepositories(driver *neo4j.DriverWithContext) *Repositories {
	repositories := Repositories{
		Dbs: Dbs{
			Neo4jDriver: driver,
		},
		Neo4jRepository: NewNeo4jRepository(driver),
	}
	return &repositories
}
