// Code generated by ent, DO NOT EDIT.

package gen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/conversationitem"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/gen/predicate"
)

// ConversationItemUpdate is the builder for updating ConversationItem entities.
type ConversationItemUpdate struct {
	config
	hooks    []Hook
	mutation *ConversationItemMutation
}

// Where appends a list predicates to the ConversationItemUpdate builder.
func (ciu *ConversationItemUpdate) Where(ps ...predicate.ConversationItem) *ConversationItemUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetType sets the "type" field.
func (ciu *ConversationItemUpdate) SetType(c conversationitem.Type) *ConversationItemUpdate {
	ciu.mutation.SetType(c)
	return ciu
}

// SetSenderId sets the "senderId" field.
func (ciu *ConversationItemUpdate) SetSenderId(s string) *ConversationItemUpdate {
	ciu.mutation.SetSenderId(s)
	return ciu
}

// SetSenderType sets the "senderType" field.
func (ciu *ConversationItemUpdate) SetSenderType(ct conversationitem.SenderType) *ConversationItemUpdate {
	ciu.mutation.SetSenderType(ct)
	return ciu
}

// SetMessage sets the "message" field.
func (ciu *ConversationItemUpdate) SetMessage(s string) *ConversationItemUpdate {
	ciu.mutation.SetMessage(s)
	return ciu
}

// SetChannel sets the "channel" field.
func (ciu *ConversationItemUpdate) SetChannel(c conversationitem.Channel) *ConversationItemUpdate {
	ciu.mutation.SetChannel(c)
	return ciu
}

// SetDirection sets the "direction" field.
func (ciu *ConversationItemUpdate) SetDirection(c conversationitem.Direction) *ConversationItemUpdate {
	ciu.mutation.SetDirection(c)
	return ciu
}

// SetTime sets the "time" field.
func (ciu *ConversationItemUpdate) SetTime(t time.Time) *ConversationItemUpdate {
	ciu.mutation.SetTime(t)
	return ciu
}

// Mutation returns the ConversationItemMutation object of the builder.
func (ciu *ConversationItemUpdate) Mutation() *ConversationItemMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *ConversationItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ciu.hooks) == 0 {
		if err = ciu.check(); err != nil {
			return 0, err
		}
		affected, err = ciu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConversationItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ciu.check(); err != nil {
				return 0, err
			}
			ciu.mutation = mutation
			affected, err = ciu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ciu.hooks) - 1; i >= 0; i-- {
			if ciu.hooks[i] == nil {
				return 0, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = ciu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ciu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *ConversationItemUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *ConversationItemUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *ConversationItemUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciu *ConversationItemUpdate) check() error {
	if v, ok := ciu.mutation.GetType(); ok {
		if err := conversationitem.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.type": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.SenderType(); ok {
		if err := conversationitem.SenderTypeValidator(v); err != nil {
			return &ValidationError{Name: "senderType", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.senderType": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Channel(); ok {
		if err := conversationitem.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.channel": %w`, err)}
		}
	}
	if v, ok := ciu.mutation.Direction(); ok {
		if err := conversationitem.DirectionValidator(v); err != nil {
			return &ValidationError{Name: "direction", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.direction": %w`, err)}
		}
	}
	if _, ok := ciu.mutation.ConversationID(); ciu.mutation.ConversationCleared() && !ok {
		return errors.New(`gen: clearing a required unique edge "ConversationItem.conversation"`)
	}
	return nil
}

func (ciu *ConversationItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   conversationitem.Table,
			Columns: conversationitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: conversationitem.FieldID,
			},
		},
	}
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldType,
		})
	}
	if value, ok := ciu.mutation.SenderId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversationitem.FieldSenderId,
		})
	}
	if value, ok := ciu.mutation.SenderType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldSenderType,
		})
	}
	if value, ok := ciu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversationitem.FieldMessage,
		})
	}
	if value, ok := ciu.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldChannel,
		})
	}
	if value, ok := ciu.mutation.Direction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldDirection,
		})
	}
	if value, ok := ciu.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: conversationitem.FieldTime,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversationitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ConversationItemUpdateOne is the builder for updating a single ConversationItem entity.
type ConversationItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConversationItemMutation
}

// SetType sets the "type" field.
func (ciuo *ConversationItemUpdateOne) SetType(c conversationitem.Type) *ConversationItemUpdateOne {
	ciuo.mutation.SetType(c)
	return ciuo
}

// SetSenderId sets the "senderId" field.
func (ciuo *ConversationItemUpdateOne) SetSenderId(s string) *ConversationItemUpdateOne {
	ciuo.mutation.SetSenderId(s)
	return ciuo
}

// SetSenderType sets the "senderType" field.
func (ciuo *ConversationItemUpdateOne) SetSenderType(ct conversationitem.SenderType) *ConversationItemUpdateOne {
	ciuo.mutation.SetSenderType(ct)
	return ciuo
}

// SetMessage sets the "message" field.
func (ciuo *ConversationItemUpdateOne) SetMessage(s string) *ConversationItemUpdateOne {
	ciuo.mutation.SetMessage(s)
	return ciuo
}

// SetChannel sets the "channel" field.
func (ciuo *ConversationItemUpdateOne) SetChannel(c conversationitem.Channel) *ConversationItemUpdateOne {
	ciuo.mutation.SetChannel(c)
	return ciuo
}

// SetDirection sets the "direction" field.
func (ciuo *ConversationItemUpdateOne) SetDirection(c conversationitem.Direction) *ConversationItemUpdateOne {
	ciuo.mutation.SetDirection(c)
	return ciuo
}

// SetTime sets the "time" field.
func (ciuo *ConversationItemUpdateOne) SetTime(t time.Time) *ConversationItemUpdateOne {
	ciuo.mutation.SetTime(t)
	return ciuo
}

// Mutation returns the ConversationItemMutation object of the builder.
func (ciuo *ConversationItemUpdateOne) Mutation() *ConversationItemMutation {
	return ciuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *ConversationItemUpdateOne) Select(field string, fields ...string) *ConversationItemUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated ConversationItem entity.
func (ciuo *ConversationItemUpdateOne) Save(ctx context.Context) (*ConversationItem, error) {
	var (
		err  error
		node *ConversationItem
	)
	if len(ciuo.hooks) == 0 {
		if err = ciuo.check(); err != nil {
			return nil, err
		}
		node, err = ciuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConversationItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ciuo.check(); err != nil {
				return nil, err
			}
			ciuo.mutation = mutation
			node, err = ciuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ciuo.hooks) - 1; i >= 0; i-- {
			if ciuo.hooks[i] == nil {
				return nil, fmt.Errorf("gen: uninitialized hook (forgotten import gen/runtime?)")
			}
			mut = ciuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ciuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ConversationItem)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ConversationItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *ConversationItemUpdateOne) SaveX(ctx context.Context) *ConversationItem {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *ConversationItemUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *ConversationItemUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ciuo *ConversationItemUpdateOne) check() error {
	if v, ok := ciuo.mutation.GetType(); ok {
		if err := conversationitem.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.type": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.SenderType(); ok {
		if err := conversationitem.SenderTypeValidator(v); err != nil {
			return &ValidationError{Name: "senderType", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.senderType": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Channel(); ok {
		if err := conversationitem.ChannelValidator(v); err != nil {
			return &ValidationError{Name: "channel", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.channel": %w`, err)}
		}
	}
	if v, ok := ciuo.mutation.Direction(); ok {
		if err := conversationitem.DirectionValidator(v); err != nil {
			return &ValidationError{Name: "direction", err: fmt.Errorf(`gen: validator failed for field "ConversationItem.direction": %w`, err)}
		}
	}
	if _, ok := ciuo.mutation.ConversationID(); ciuo.mutation.ConversationCleared() && !ok {
		return errors.New(`gen: clearing a required unique edge "ConversationItem.conversation"`)
	}
	return nil
}

func (ciuo *ConversationItemUpdateOne) sqlSave(ctx context.Context) (_node *ConversationItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   conversationitem.Table,
			Columns: conversationitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: conversationitem.FieldID,
			},
		},
	}
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`gen: missing "ConversationItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conversationitem.FieldID)
		for _, f := range fields {
			if !conversationitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("gen: invalid field %q for query", f)}
			}
			if f != conversationitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldType,
		})
	}
	if value, ok := ciuo.mutation.SenderId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversationitem.FieldSenderId,
		})
	}
	if value, ok := ciuo.mutation.SenderType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldSenderType,
		})
	}
	if value, ok := ciuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: conversationitem.FieldMessage,
		})
	}
	if value, ok := ciuo.mutation.Channel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldChannel,
		})
	}
	if value, ok := ciuo.mutation.Direction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: conversationitem.FieldDirection,
		})
	}
	if value, ok := ciuo.mutation.Time(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: conversationitem.FieldTime,
		})
	}
	_node = &ConversationItem{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversationitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
