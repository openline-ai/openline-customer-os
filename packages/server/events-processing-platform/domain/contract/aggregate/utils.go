package aggregate

import (
	"context"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/eventstore"
	"github.com/openline-ai/openline-customer-os/packages/server/events-processing-platform/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"strings"
)

func GetContractObjectID(aggregateID string, tenant string) string {
	if tenant == "" {
		return getContractObjectUUID(aggregateID)
	}
	return strings.ReplaceAll(aggregateID, string(ContractAggregateType)+"-"+tenant+"-", "")
}

// Use this method when tenant is not known
func getContractObjectUUID(aggregateID string) string {
	parts := strings.Split(aggregateID, "-")
	fullUUID := parts[len(parts)-5] + "-" + parts[len(parts)-4] + "-" + parts[len(parts)-3] + "-" + parts[len(parts)-2] + "-" + parts[len(parts)-1]
	return fullUUID
}

func LoadContractAggregate(ctx context.Context, eventStore eventstore.AggregateStore, tenant, objectID string) (*ContractAggregate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoadContractAggregate")
	defer span.Finish()
	span.SetTag(tracing.SpanTagTenant, tenant)
	span.LogFields(log.String("ObjectID", objectID))

	contractAggregate := NewContractAggregateWithTenantAndID(tenant, objectID)

	err := eventStore.Exists(ctx, contractAggregate.GetID())
	if err != nil {
		if !errors.Is(err, eventstore.ErrAggregateNotFound) {
			tracing.TraceErr(span, err)
			return nil, err
		} else {
			return contractAggregate, nil
		}
	}

	if err = eventStore.Load(ctx, contractAggregate); err != nil {
		tracing.TraceErr(span, err)
		return nil, err
	}

	return contractAggregate, nil
}
