// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/messagefeed"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/messageitem"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/predicate"
)

// MessageItemQuery is the builder for querying MessageItem entities.
type MessageItemQuery struct {
	config
	limit           *int
	offset          *int
	unique          *bool
	order           []OrderFunc
	fields          []string
	predicates      []predicate.MessageItem
	withMessageFeed *MessageFeedQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MessageItemQuery builder.
func (miq *MessageItemQuery) Where(ps ...predicate.MessageItem) *MessageItemQuery {
	miq.predicates = append(miq.predicates, ps...)
	return miq
}

// Limit adds a limit step to the query.
func (miq *MessageItemQuery) Limit(limit int) *MessageItemQuery {
	miq.limit = &limit
	return miq
}

// Offset adds an offset step to the query.
func (miq *MessageItemQuery) Offset(offset int) *MessageItemQuery {
	miq.offset = &offset
	return miq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (miq *MessageItemQuery) Unique(unique bool) *MessageItemQuery {
	miq.unique = &unique
	return miq
}

// Order adds an order step to the query.
func (miq *MessageItemQuery) Order(o ...OrderFunc) *MessageItemQuery {
	miq.order = append(miq.order, o...)
	return miq
}

// QueryMessageFeed chains the current query on the "message_feed" edge.
func (miq *MessageItemQuery) QueryMessageFeed() *MessageFeedQuery {
	query := &MessageFeedQuery{config: miq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(messageitem.Table, messageitem.FieldID, selector),
			sqlgraph.To(messagefeed.Table, messagefeed.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, messageitem.MessageFeedTable, messageitem.MessageFeedColumn),
		)
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MessageItem entity from the query.
// Returns a *NotFoundError when no MessageItem was found.
func (miq *MessageItemQuery) First(ctx context.Context) (*MessageItem, error) {
	nodes, err := miq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{messageitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (miq *MessageItemQuery) FirstX(ctx context.Context) *MessageItem {
	node, err := miq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MessageItem ID from the query.
// Returns a *NotFoundError when no MessageItem ID was found.
func (miq *MessageItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{messageitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (miq *MessageItemQuery) FirstIDX(ctx context.Context) int {
	id, err := miq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MessageItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MessageItem entity is found.
// Returns a *NotFoundError when no MessageItem entities are found.
func (miq *MessageItemQuery) Only(ctx context.Context) (*MessageItem, error) {
	nodes, err := miq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{messageitem.Label}
	default:
		return nil, &NotSingularError{messageitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (miq *MessageItemQuery) OnlyX(ctx context.Context) *MessageItem {
	node, err := miq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MessageItem ID in the query.
// Returns a *NotSingularError when more than one MessageItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (miq *MessageItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = miq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{messageitem.Label}
	default:
		err = &NotSingularError{messageitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (miq *MessageItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := miq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MessageItems.
func (miq *MessageItemQuery) All(ctx context.Context) ([]*MessageItem, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return miq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (miq *MessageItemQuery) AllX(ctx context.Context) []*MessageItem {
	nodes, err := miq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MessageItem IDs.
func (miq *MessageItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := miq.Select(messageitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (miq *MessageItemQuery) IDsX(ctx context.Context) []int {
	ids, err := miq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (miq *MessageItemQuery) Count(ctx context.Context) (int, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return miq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (miq *MessageItemQuery) CountX(ctx context.Context) int {
	count, err := miq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (miq *MessageItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return miq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (miq *MessageItemQuery) ExistX(ctx context.Context) bool {
	exist, err := miq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MessageItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (miq *MessageItemQuery) Clone() *MessageItemQuery {
	if miq == nil {
		return nil
	}
	return &MessageItemQuery{
		config:          miq.config,
		limit:           miq.limit,
		offset:          miq.offset,
		order:           append([]OrderFunc{}, miq.order...),
		predicates:      append([]predicate.MessageItem{}, miq.predicates...),
		withMessageFeed: miq.withMessageFeed.Clone(),
		// clone intermediate query.
		sql:    miq.sql.Clone(),
		path:   miq.path,
		unique: miq.unique,
	}
}

// WithMessageFeed tells the query-builder to eager-load the nodes that are connected to
// the "message_feed" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MessageItemQuery) WithMessageFeed(opts ...func(*MessageFeedQuery)) *MessageItemQuery {
	query := &MessageFeedQuery{config: miq.config}
	for _, opt := range opts {
		opt(query)
	}
	miq.withMessageFeed = query
	return miq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type messageitem.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MessageItem.Query().
//		GroupBy(messageitem.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (miq *MessageItemQuery) GroupBy(field string, fields ...string) *MessageItemGroupBy {
	grbuild := &MessageItemGroupBy{config: miq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return miq.sqlQuery(ctx), nil
	}
	grbuild.label = messageitem.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type messageitem.Type `json:"type,omitempty"`
//	}
//
//	client.MessageItem.Query().
//		Select(messageitem.FieldType).
//		Scan(ctx, &v)
func (miq *MessageItemQuery) Select(fields ...string) *MessageItemSelect {
	miq.fields = append(miq.fields, fields...)
	selbuild := &MessageItemSelect{MessageItemQuery: miq}
	selbuild.label = messageitem.Label
	selbuild.flds, selbuild.scan = &miq.fields, selbuild.Scan
	return selbuild
}

func (miq *MessageItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range miq.fields {
		if !messageitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if miq.path != nil {
		prev, err := miq.path(ctx)
		if err != nil {
			return err
		}
		miq.sql = prev
	}
	return nil
}

func (miq *MessageItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MessageItem, error) {
	var (
		nodes       = []*MessageItem{}
		withFKs     = miq.withFKs
		_spec       = miq.querySpec()
		loadedTypes = [1]bool{
			miq.withMessageFeed != nil,
		}
	)
	if miq.withMessageFeed != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, messageitem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MessageItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MessageItem{config: miq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, miq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := miq.withMessageFeed; query != nil {
		if err := miq.loadMessageFeed(ctx, query, nodes, nil,
			func(n *MessageItem, e *MessageFeed) { n.Edges.MessageFeed = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (miq *MessageItemQuery) loadMessageFeed(ctx context.Context, query *MessageFeedQuery, nodes []*MessageItem, init func(*MessageItem), assign func(*MessageItem, *MessageFeed)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MessageItem)
	for i := range nodes {
		if nodes[i].message_feed_message_item == nil {
			continue
		}
		fk := *nodes[i].message_feed_message_item
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(messagefeed.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "message_feed_message_item" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (miq *MessageItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := miq.querySpec()
	_spec.Node.Columns = miq.fields
	if len(miq.fields) > 0 {
		_spec.Unique = miq.unique != nil && *miq.unique
	}
	return sqlgraph.CountNodes(ctx, miq.driver, _spec)
}

func (miq *MessageItemQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := miq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (miq *MessageItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messageitem.Table,
			Columns: messageitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messageitem.FieldID,
			},
		},
		From:   miq.sql,
		Unique: true,
	}
	if unique := miq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := miq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messageitem.FieldID)
		for i := range fields {
			if fields[i] != messageitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := miq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := miq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := miq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := miq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (miq *MessageItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(miq.driver.Dialect())
	t1 := builder.Table(messageitem.Table)
	columns := miq.fields
	if len(columns) == 0 {
		columns = messageitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if miq.sql != nil {
		selector = miq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if miq.unique != nil && *miq.unique {
		selector.Distinct()
	}
	for _, p := range miq.predicates {
		p(selector)
	}
	for _, p := range miq.order {
		p(selector)
	}
	if offset := miq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := miq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MessageItemGroupBy is the group-by builder for MessageItem entities.
type MessageItemGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (migb *MessageItemGroupBy) Aggregate(fns ...AggregateFunc) *MessageItemGroupBy {
	migb.fns = append(migb.fns, fns...)
	return migb
}

// Scan applies the group-by query and scans the result into the given value.
func (migb *MessageItemGroupBy) Scan(ctx context.Context, v any) error {
	query, err := migb.path(ctx)
	if err != nil {
		return err
	}
	migb.sql = query
	return migb.sqlScan(ctx, v)
}

func (migb *MessageItemGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range migb.fields {
		if !messageitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := migb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := migb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (migb *MessageItemGroupBy) sqlQuery() *sql.Selector {
	selector := migb.sql.Select()
	aggregation := make([]string, 0, len(migb.fns))
	for _, fn := range migb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(migb.fields)+len(migb.fns))
		for _, f := range migb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(migb.fields...)...)
}

// MessageItemSelect is the builder for selecting fields of MessageItem entities.
type MessageItemSelect struct {
	*MessageItemQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mis *MessageItemSelect) Scan(ctx context.Context, v any) error {
	if err := mis.prepareQuery(ctx); err != nil {
		return err
	}
	mis.sql = mis.MessageItemQuery.sqlQuery(ctx)
	return mis.sqlScan(ctx, v)
}

func (mis *MessageItemSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := mis.sql.Query()
	if err := mis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
