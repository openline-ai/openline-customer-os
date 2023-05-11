package repository

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
)

type SearchRepository interface {
	GCliSearch(ctx context.Context, tenant, keyword string, limit int) ([]*db.Record, error)
}

type searchRepository struct {
	driver *neo4j.DriverWithContext
}

func NewSearchRepository(driver *neo4j.DriverWithContext) SearchRepository {
	return &searchRepository{
		driver: driver,
	}
}

func (r *searchRepository) GCliSearch(ctx context.Context, tenant, keyword string, limit int) ([]*db.Record, error) {
	session := utils.NewNeo4jReadSession(ctx, *r.driver)
	defer session.Close(ctx)

	params := map[string]any{
		"tenant":        tenant,
		"fuzzyKeyword":  fmt.Sprintf("%s~", keyword),
		"keyword":       keyword,
		"indexStandard": "basicSearchStandard_location_terms",
		"limit":         limit,
	}
	query := "CALL { " +
		" CALL db.index.fulltext.queryNodes($indexStandard, $keyword) YIELD node, score RETURN score, node, labels(node) as labels limit $limit " +
		" union" +
		" CALL db.index.fulltext.queryNodes($indexStandard, $fuzzyKeyword) YIELD node, score RETURN score, node, labels(node) as labels limit $limit " +
		"} " +
		" with labels, node, score order by score desc " +
		" with labels, node, collect(score) as scores " +
		" return labels, head(scores) as score, node order by score desc limit $limit"

	records, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		queryResult, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}
		return queryResult.Collect(ctx)
	})
	return records.([]*db.Record), err
}
