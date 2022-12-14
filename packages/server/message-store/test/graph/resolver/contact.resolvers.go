package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/openline-ai/openline-customer-os/packages/server/message-store/test/graph/generated"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/test/graph/model"
)

// ContactType is the resolver for the contactType field.
func (r *contactResolver) ContactType(ctx context.Context, obj *model.Contact) (*model.ContactType, error) {
	panic(fmt.Errorf("not implemented: ContactType - contactType"))
}

// Roles is the resolver for the roles field.
func (r *contactResolver) Roles(ctx context.Context, obj *model.Contact) ([]*model.ContactRole, error) {
	panic(fmt.Errorf("not implemented: Roles - roles"))
}

// Groups is the resolver for the groups field.
func (r *contactResolver) Groups(ctx context.Context, obj *model.Contact) ([]*model.ContactGroup, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// PhoneNumbers is the resolver for the phoneNumbers field.
func (r *contactResolver) PhoneNumbers(ctx context.Context, obj *model.Contact) ([]*model.PhoneNumber, error) {
	if r.Resolver.PhoneNumbersByContact != nil {
		return r.Resolver.PhoneNumbersByContact(ctx, obj)
	}
	panic(fmt.Errorf("not implemented: PhoneNumbers - phoneNumbers"))
}

// Emails is the resolver for the emails field.
func (r *contactResolver) Emails(ctx context.Context, obj *model.Contact) ([]*model.Email, error) {
	if r.Resolver.EmailsByContact != nil {
		return r.Resolver.EmailsByContact(ctx, obj)
	}
	panic(fmt.Errorf("not implemented: Emails - emails"))
}

// Addresses is the resolver for the addresses field.
func (r *contactResolver) Addresses(ctx context.Context, obj *model.Contact) ([]*model.Address, error) {
	panic(fmt.Errorf("not implemented: Addresses - addresses"))
}

// CustomFields is the resolver for the customFields field.
func (r *contactResolver) CustomFields(ctx context.Context, obj *model.Contact) ([]*model.CustomField, error) {
	panic(fmt.Errorf("not implemented: CustomFields - customFields"))
}

// FieldSets is the resolver for the fieldSets field.
func (r *contactResolver) FieldSets(ctx context.Context, obj *model.Contact) ([]*model.FieldSet, error) {
	panic(fmt.Errorf("not implemented: FieldSets - fieldSets"))
}

// Template is the resolver for the template field.
func (r *contactResolver) Template(ctx context.Context, obj *model.Contact) (*model.EntityTemplate, error) {
	panic(fmt.Errorf("not implemented: Template - template"))
}

// Owner is the resolver for the owner field.
func (r *contactResolver) Owner(ctx context.Context, obj *model.Contact) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// Notes is the resolver for the notes field.
func (r *contactResolver) Notes(ctx context.Context, obj *model.Contact, pagination *model.Pagination) (*model.NotePage, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}

// Conversations is the resolver for the conversations field.
func (r *contactResolver) Conversations(ctx context.Context, obj *model.Contact, pagination *model.Pagination, sort []*model.SortBy) (*model.ConversationPage, error) {
	panic(fmt.Errorf("not implemented: Conversations - conversations"))
}

// Actions is the resolver for the actions field.
func (r *contactResolver) Actions(ctx context.Context, obj *model.Contact, from time.Time, to time.Time, actionTypes []model.ActionType) ([]model.Action, error) {
	panic(fmt.Errorf("not implemented: Actions - actions"))
}

// ContactCreate is the resolver for the contact_Create field.
func (r *mutationResolver) ContactCreate(ctx context.Context, input model.ContactInput) (*model.Contact, error) {
	if r.Resolver.ContactCreate != nil {
		return r.Resolver.ContactCreate(ctx, input)
	}
	panic(fmt.Errorf("not implemented: ContactCreate - contact_Create"))
}

// ContactUpdate is the resolver for the contact_Update field.
func (r *mutationResolver) ContactUpdate(ctx context.Context, input model.ContactUpdateInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented: ContactUpdate - contact_Update"))
}

// ContactHardDelete is the resolver for the contact_HardDelete field.
func (r *mutationResolver) ContactHardDelete(ctx context.Context, contactID string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: ContactHardDelete - contact_HardDelete"))
}

// ContactSoftDelete is the resolver for the contact_SoftDelete field.
func (r *mutationResolver) ContactSoftDelete(ctx context.Context, contactID string) (*model.Result, error) {
	panic(fmt.Errorf("not implemented: ContactSoftDelete - contact_SoftDelete"))
}

// Contact is the resolver for the contact field.
func (r *queryResolver) Contact(ctx context.Context, id string) (*model.Contact, error) {
	if r.Resolver.GetContactById != nil {
		return r.Resolver.GetContactById(ctx, id)
	}
	panic(fmt.Errorf("not implemented: Contact - contact"))
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context, pagination *model.Pagination, where *model.Filter, sort []*model.SortBy) (*model.ContactsPage, error) {
	panic(fmt.Errorf("not implemented: Contacts - contacts"))
}

// ContactByEmail is the resolver for the contact_ByEmail field.
func (r *queryResolver) ContactByEmail(ctx context.Context, email string) (*model.Contact, error) {
	if r.Resolver.GetContactByEmail != nil {
		return r.Resolver.GetContactByEmail(ctx, email)
	}
	panic(fmt.Errorf("not implemented: ContactByEmail - contact_ByEmail"))
}

// ContactByPhone is the resolver for the contact_ByPhone field.
func (r *queryResolver) ContactByPhone(ctx context.Context, e164 string) (*model.Contact, error) {
	if r.Resolver.GetContactByPhone != nil {
		return r.Resolver.GetContactByPhone(ctx, e164)
	}

	panic(fmt.Errorf("not implemented: ContactByPhone - contact_ByPhone"))
}

// Contact returns generated.ContactResolver implementation.
func (r *Resolver) Contact() generated.ContactResolver { return &contactResolver{r} }

type contactResolver struct{ *Resolver }
