package service

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/source"
	"time"
)

type result struct {
	completed int
	failed    int
	skipped   int
}

type SyncService interface {
	Sync(ctx context.Context, sourceService source.SourceDataService, syncDate time.Time, tenant, runId string, batchSize int) (int, int, int)
}
