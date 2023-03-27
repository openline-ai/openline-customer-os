package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/comms-api/test/graph/model"
)

// Template is the resolver for the template field.
func (r *customFieldResolver) Template(ctx context.Context, obj *model.CustomField) (*model.CustomFieldTemplate, error) {
	panic(fmt.Errorf("not implemented: Template - template"))
}

// CustomFields is the resolver for the customFields field.
func (r *fieldSetResolver) CustomFields(ctx context.Context, obj *model.FieldSet) ([]*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFields - customFields"))
}

// Template is the resolver for the template field.
func (r *fieldSetResolver) Template(ctx context.Context, obj *model.FieldSet) (*model.FieldSetTemplate, error) {
	panic(fmt.Errorf("not implemented: Template - template"))
}

// CustomFieldsMergeAndUpdateInContact is the resolver for the customFieldsMergeAndUpdateInContact field.
func (r *mutationResolver) CustomFieldsMergeAndUpdateInContact(ctx context.Context, contactID string, customFields []*model.CustomFieldInput, fieldSets []*model.FieldSetInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: CustomFieldsMergeAndUpdateInContact - customFieldsMergeAndUpdateInContact"))
}

// CustomFieldMergeToContact is the resolver for the customFieldMergeToContact field.
func (r *mutationResolver) CustomFieldMergeToContact(ctx context.Context, contactID string, input model.CustomFieldInput) (*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFieldMergeToContact - customFieldMergeToContact"))
}

// CustomFieldUpdateInContact is the resolver for the customFieldUpdateInContact field.
func (r *mutationResolver) CustomFieldUpdateInContact(ctx context.Context, contactID string, input model.CustomFieldUpdateInput) (*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFieldUpdateInContact - customFieldUpdateInContact"))
}

// CustomFieldDeleteFromContactByName is the resolver for the customFieldDeleteFromContactByName field.
func (r *mutationResolver) CustomFieldDeleteFromContactByName(ctx context.Context, contactID string, fieldName string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: CustomFieldDeleteFromContactByName - customFieldDeleteFromContactByName"))
}

// CustomFieldDeleteFromContactByID is the resolver for the customFieldDeleteFromContactById field.
func (r *mutationResolver) CustomFieldDeleteFromContactByID(ctx context.Context, contactID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: CustomFieldDeleteFromContactByID - customFieldDeleteFromContactById"))
}

// CustomFieldMergeToFieldSet is the resolver for the customFieldMergeToFieldSet field.
func (r *mutationResolver) CustomFieldMergeToFieldSet(ctx context.Context, contactID string, fieldSetID string, input model.CustomFieldInput) (*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFieldMergeToFieldSet - customFieldMergeToFieldSet"))
}

// CustomFieldUpdateInFieldSet is the resolver for the customFieldUpdateInFieldSet field.
func (r *mutationResolver) CustomFieldUpdateInFieldSet(ctx context.Context, contactID string, fieldSetID string, input model.CustomFieldUpdateInput) (*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFieldUpdateInFieldSet - customFieldUpdateInFieldSet"))
}

// CustomFieldDeleteFromFieldSetByID is the resolver for the customFieldDeleteFromFieldSetById field.
func (r *mutationResolver) CustomFieldDeleteFromFieldSetByID(ctx context.Context, contactID string, fieldSetID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: CustomFieldDeleteFromFieldSetByID - customFieldDeleteFromFieldSetById"))
}

// FieldSetMergeToContact is the resolver for the fieldSetMergeToContact field.
func (r *mutationResolver) FieldSetMergeToContact(ctx context.Context, contactID string, input model.FieldSetInput) (*model.FieldSet, error) {
	panic(fmt.Errorf("not implemented: FieldSetMergeToContact - fieldSetMergeToContact"))
}

// FieldSetUpdateInContact is the resolver for the fieldSetUpdateInContact field.
func (r *mutationResolver) FieldSetUpdateInContact(ctx context.Context, contactID string, input model.FieldSetUpdateInput) (*model.FieldSet, error) {
	panic(fmt.Errorf("not implemented: FieldSetUpdateInContact - fieldSetUpdateInContact"))
}

// FieldSetDeleteFromContact is the resolver for the fieldSetDeleteFromContact field.
func (r *mutationResolver) FieldSetDeleteFromContact(ctx context.Context, contactID string, id string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: FieldSetDeleteFromContact - fieldSetDeleteFromContact"))
}

// CustomField returns generated.CustomFieldResolver implementation.
func (r *Resolver) CustomField() generated.CustomFieldResolver { return &customFieldResolver{r} }

// FieldSet returns generated.FieldSetResolver implementation.
func (r *Resolver) FieldSet() generated.FieldSetResolver { return &fieldSetResolver{r} }

type customFieldResolver struct{ *Resolver }
type fieldSetResolver struct{ *Resolver }
