// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/messageitem"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/predicate"
)

// MessageItemUpdate is the builder for updating MessageItem entities.
type MessageItemUpdate struct {
	config
	hooks    []Hook
	mutation *MessageItemMutation
}

// Where appends a list predicates to the MessageItemUpdate builder.
func (miu *MessageItemUpdate) Where(ps ...predicate.MessageItem) *MessageItemUpdate {
	miu.mutation.Where(ps...)
	return miu
}

// SetType sets the "type" field.
func (miu *MessageItemUpdate) SetType(m messageitem.Type) *MessageItemUpdate {
	miu.mutation.SetType(m)
	return miu
}

// SetUsername sets the "username" field.
func (miu *MessageItemUpdate) SetUsername(s string) *MessageItemUpdate {
	miu.mutation.SetUsername(s)
	return miu
}

// SetMessage sets the "message" field.
func (miu *MessageItemUpdate) SetMessage(s string) *MessageItemUpdate {
	miu.mutation.SetMessage(s)
	return miu
}

// SetChannel sets the "channel" field.
func (miu *MessageItemUpdate) SetChannel(m messageitem.Channel) *MessageItemUpdate {
	miu.mutation.SetChannel(m)
	return miu
}

// SetDirection sets the "direction" field.
func (miu *MessageItemUpdate) SetDirection(m messageitem.Direction) *MessageItemUpdate {
	miu.mutation.SetDirection(m)
	return miu
}

// SetTime sets the "time" field.
func (miu *MessageItemUpdate) SetTime(t time.Time) *MessageItemUpdate {
	miu.mutation.SetTime(t)
	return miu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (miu *MessageItemUpdate) SetNillableTime(t *time.Time) *MessageItemUpdate {
	if t != nil {
		miu.SetTime(*t)
	}
	return miu
}

// ClearTime clears the value of the "time" field.
func (miu *MessageItemUpdate) ClearTime() *MessageItemUpdate {
	miu.mutation.ClearTime()
	return miu
}

// Mutation returns the MessageItemMutation object of the builder.
func (miu *MessageItemUpdate) Mutation() *MessageItemMutation {
	return miu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (miu *MessageItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(miu.hooks) == 0 {
		if err = miu.check(); err != nil {
			return 0, err
		}
		affected, err = miu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = miu.check(); err != nil {
				return 0, err
			}
			miu.mutation = mutation
			affected, err = miu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(miu.hooks) - 1; i >= 0; i-- {
			if miu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = miu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, miu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (miu *MessageItemUpdate) SaveX(ctx context.Context) int {
	affected, err := miu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (miu *MessageItemUpdate) Exec(ctx context.Context) error {
	_, err := miu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (miu *MessageItemUpdate) ExecX(ctx context.Context) {
	if err := miu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (miu *MessageItemUpdate) check() error {
	if v, ok := miu.mutation.GetType(); ok {
		if err := messageitem.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "MessageItem.type": %w`, err)}
		}
	}
	if v, ok := miu.mutation.Channel(); ok {
		if err := messageitem.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`ent: validator failed for field "MessageItem.channel": %w`, err)}
		}
	}
	if v, ok := miu.mutation.Direction(); ok {
		if err := messageitem.DirectionValidator(v); err != nil {
			return &ValidationError{Name: "direction", err: fmt.Errorf(`ent: validator failed for field "MessageItem.direction": %w`, err)}
		}
	}
	if _, ok := miu.mutation.MessageFeedID(); miu.mutation.MessageFeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MessageItem.message_feed"`)
	}
	return nil
}

func (miu *MessageItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageitem.Table,
			Columns: messageitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageitem.FieldID,
			},
		},
	}
	if ps := miu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := miu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messageitem.FieldType,
		})
	}
	if value, ok := miu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageitem.FieldUsername,
		})
	}
	if value, ok := miu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageitem.FieldMessage,
		})
	}
	if value, ok := miu.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messageitem.FieldChannel,
		})
	}
	if value, ok := miu.mutation.Direction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messageitem.FieldDirection,
		})
	}
	if value, ok := miu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messageitem.FieldTime,
		})
	}
	if miu.mutation.TimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: messageitem.FieldTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, miu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MessageItemUpdateOne is the builder for updating a single MessageItem entity.
type MessageItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageItemMutation
}

// SetType sets the "type" field.
func (miuo *MessageItemUpdateOne) SetType(m messageitem.Type) *MessageItemUpdateOne {
	miuo.mutation.SetType(m)
	return miuo
}

// SetUsername sets the "username" field.
func (miuo *MessageItemUpdateOne) SetUsername(s string) *MessageItemUpdateOne {
	miuo.mutation.SetUsername(s)
	return miuo
}

// SetMessage sets the "message" field.
func (miuo *MessageItemUpdateOne) SetMessage(s string) *MessageItemUpdateOne {
	miuo.mutation.SetMessage(s)
	return miuo
}

// SetChannel sets the "channel" field.
func (miuo *MessageItemUpdateOne) SetChannel(m messageitem.Channel) *MessageItemUpdateOne {
	miuo.mutation.SetChannel(m)
	return miuo
}

// SetDirection sets the "direction" field.
func (miuo *MessageItemUpdateOne) SetDirection(m messageitem.Direction) *MessageItemUpdateOne {
	miuo.mutation.SetDirection(m)
	return miuo
}

// SetTime sets the "time" field.
func (miuo *MessageItemUpdateOne) SetTime(t time.Time) *MessageItemUpdateOne {
	miuo.mutation.SetTime(t)
	return miuo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (miuo *MessageItemUpdateOne) SetNillableTime(t *time.Time) *MessageItemUpdateOne {
	if t != nil {
		miuo.SetTime(*t)
	}
	return miuo
}

// ClearTime clears the value of the "time" field.
func (miuo *MessageItemUpdateOne) ClearTime() *MessageItemUpdateOne {
	miuo.mutation.ClearTime()
	return miuo
}

// Mutation returns the MessageItemMutation object of the builder.
func (miuo *MessageItemUpdateOne) Mutation() *MessageItemMutation {
	return miuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (miuo *MessageItemUpdateOne) Select(field string, fields ...string) *MessageItemUpdateOne {
	miuo.fields = append([]string{field}, fields...)
	return miuo
}

// Save executes the query and returns the updated MessageItem entity.
func (miuo *MessageItemUpdateOne) Save(ctx context.Context) (*MessageItem, error) {
	var (
		err  error
		node *MessageItem
	)
	if len(miuo.hooks) == 0 {
		if err = miuo.check(); err != nil {
			return nil, err
		}
		node, err = miuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = miuo.check(); err != nil {
				return nil, err
			}
			miuo.mutation = mutation
			node, err = miuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(miuo.hooks) - 1; i >= 0; i-- {
			if miuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = miuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, miuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MessageItem)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (miuo *MessageItemUpdateOne) SaveX(ctx context.Context) *MessageItem {
	node, err := miuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (miuo *MessageItemUpdateOne) Exec(ctx context.Context) error {
	_, err := miuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (miuo *MessageItemUpdateOne) ExecX(ctx context.Context) {
	if err := miuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (miuo *MessageItemUpdateOne) check() error {
	if v, ok := miuo.mutation.GetType(); ok {
		if err := messageitem.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "MessageItem.type": %w`, err)}
		}
	}
	if v, ok := miuo.mutation.Channel(); ok {
		if err := messageitem.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`ent: validator failed for field "MessageItem.channel": %w`, err)}
		}
	}
	if v, ok := miuo.mutation.Direction(); ok {
		if err := messageitem.DirectionValidator(v); err != nil {
			return &ValidationError{Name: "direction", err: fmt.Errorf(`ent: validator failed for field "MessageItem.direction": %w`, err)}
		}
	}
	if _, ok := miuo.mutation.MessageFeedID(); miuo.mutation.MessageFeedCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "MessageItem.message_feed"`)
	}
	return nil
}

func (miuo *MessageItemUpdateOne) sqlSave(ctx context.Context) (_node *MessageItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageitem.Table,
			Columns: messageitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageitem.FieldID,
			},
		},
	}
	id, ok := miuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MessageItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := miuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageitem.FieldID)
		for _, f := range fields {
			if !messageitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messageitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := miuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := miuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messageitem.FieldType,
		})
	}
	if value, ok := miuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageitem.FieldUsername,
		})
	}
	if value, ok := miuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messageitem.FieldMessage,
		})
	}
	if value, ok := miuo.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messageitem.FieldChannel,
		})
	}
	if value, ok := miuo.mutation.Direction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: messageitem.FieldDirection,
		})
	}
	if value, ok := miuo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messageitem.FieldTime,
		})
	}
	if miuo.mutation.TimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: messageitem.FieldTime,
		})
	}
	_node = &MessageItem{config: miuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, miuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messageitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
