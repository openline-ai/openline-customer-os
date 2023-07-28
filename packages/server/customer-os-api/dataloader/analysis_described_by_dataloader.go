package dataloader

import (
	"context"
	"errors"
	"github.com/graph-gophers/dataloader"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-api/repository"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"reflect"
)

func (i *Loaders) GetDescribedByFor(ctx context.Context, linkedWith repository.LinkedWith, entityId string) (*entity.AnalysisEntities, error) {
	thunk := i.DescribedByFor.Load(context.WithValue(ctx, "LinkedWith", linkedWith), dataloader.StringKey(entityId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	resultObj := result.(entity.AnalysisEntities)
	return &resultObj, nil
}

func (b *analysisBatcher) getDescribedByFor(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	ids, keyOrder := sortKeys(keys)

	ctx, cancel := utils.GetLongLivedContext(ctx)
	defer cancel()

	analysisEntitiesPtr, err := b.analysisService.GetDescribedByForXX(ctx, ids, ctx.Value("LinkedWith").(repository.LinkedWith))
	if err != nil {
		// check if context deadline exceeded error occurred
		if ctx.Err() == context.DeadlineExceeded {
			return []*dataloader.Result{{Data: nil, Error: errors.New("deadline exceeded to get analysis for meeting")}}
		}
		return []*dataloader.Result{{Data: nil, Error: err}}
	}

	analysisGrouped := make(map[string]entity.AnalysisEntities)
	for _, val := range *analysisEntitiesPtr {
		if list, ok := analysisGrouped[val.GetDataloaderKey()]; ok {
			analysisGrouped[val.GetDataloaderKey()] = append(list, val)
		} else {
			analysisGrouped[val.GetDataloaderKey()] = entity.AnalysisEntities{val}
		}
	}

	// construct an output array of dataloader results
	results := make([]*dataloader.Result, len(keys))
	for contactId, record := range analysisGrouped {
		ix, ok := keyOrder[contactId]
		if ok {
			results[ix] = &dataloader.Result{Data: record, Error: nil}
			delete(keyOrder, contactId)
		}
	}
	for _, ix := range keyOrder {
		results[ix] = &dataloader.Result{Data: entity.AnalysisEntities{}, Error: nil}
	}

	if err = assertEntitiesType(results, reflect.TypeOf(entity.AnalysisEntities{})); err != nil {
		return []*dataloader.Result{{nil, err}}
	}

	return results
}
