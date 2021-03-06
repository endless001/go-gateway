// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gateway/internal/ent/loadbalance"
	"go-gateway/internal/ent/predicate"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// LoadBalanceQuery is the builder for querying LoadBalance entities.
type LoadBalanceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.LoadBalance
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (lbq *LoadBalanceQuery) Where(ps ...predicate.LoadBalance) *LoadBalanceQuery {
	lbq.predicates = append(lbq.predicates, ps...)
	return lbq
}

// Limit adds a limit step to the query.
func (lbq *LoadBalanceQuery) Limit(limit int) *LoadBalanceQuery {
	lbq.limit = &limit
	return lbq
}

// Offset adds an offset step to the query.
func (lbq *LoadBalanceQuery) Offset(offset int) *LoadBalanceQuery {
	lbq.offset = &offset
	return lbq
}

// Order adds an order step to the query.
func (lbq *LoadBalanceQuery) Order(o ...OrderFunc) *LoadBalanceQuery {
	lbq.order = append(lbq.order, o...)
	return lbq
}

// First returns the first LoadBalance entity in the query. Returns *NotFoundError when no loadbalance was found.
func (lbq *LoadBalanceQuery) First(ctx context.Context) (*LoadBalance, error) {
	nodes, err := lbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{loadbalance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lbq *LoadBalanceQuery) FirstX(ctx context.Context) *LoadBalance {
	node, err := lbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LoadBalance id in the query. Returns *NotFoundError when no id was found.
func (lbq *LoadBalanceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{loadbalance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lbq *LoadBalanceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only LoadBalance entity in the query, returns an error if not exactly one entity was returned.
func (lbq *LoadBalanceQuery) Only(ctx context.Context) (*LoadBalance, error) {
	nodes, err := lbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{loadbalance.Label}
	default:
		return nil, &NotSingularError{loadbalance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lbq *LoadBalanceQuery) OnlyX(ctx context.Context) *LoadBalance {
	node, err := lbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only LoadBalance id in the query, returns an error if not exactly one id was returned.
func (lbq *LoadBalanceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = &NotSingularError{loadbalance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lbq *LoadBalanceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LoadBalances.
func (lbq *LoadBalanceQuery) All(ctx context.Context) ([]*LoadBalance, error) {
	if err := lbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lbq *LoadBalanceQuery) AllX(ctx context.Context) []*LoadBalance {
	nodes, err := lbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LoadBalance ids.
func (lbq *LoadBalanceQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := lbq.Select(loadbalance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lbq *LoadBalanceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lbq *LoadBalanceQuery) Count(ctx context.Context) (int, error) {
	if err := lbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lbq *LoadBalanceQuery) CountX(ctx context.Context) int {
	count, err := lbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lbq *LoadBalanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := lbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lbq *LoadBalanceQuery) ExistX(ctx context.Context) bool {
	exist, err := lbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lbq *LoadBalanceQuery) Clone() *LoadBalanceQuery {
	if lbq == nil {
		return nil
	}
	return &LoadBalanceQuery{
		config:     lbq.config,
		limit:      lbq.limit,
		offset:     lbq.offset,
		order:      append([]OrderFunc{}, lbq.order...),
		unique:     append([]string{}, lbq.unique...),
		predicates: append([]predicate.LoadBalance{}, lbq.predicates...),
		// clone intermediate query.
		sql:  lbq.sql.Clone(),
		path: lbq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ServiceID int64 `json:"service_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LoadBalance.Query().
//		GroupBy(loadbalance.FieldServiceID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (lbq *LoadBalanceQuery) GroupBy(field string, fields ...string) *LoadBalanceGroupBy {
	group := &LoadBalanceGroupBy{config: lbq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lbq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		ServiceID int64 `json:"service_id,omitempty"`
//	}
//
//	client.LoadBalance.Query().
//		Select(loadbalance.FieldServiceID).
//		Scan(ctx, &v)
//
func (lbq *LoadBalanceQuery) Select(field string, fields ...string) *LoadBalanceSelect {
	selector := &LoadBalanceSelect{config: lbq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lbq.sqlQuery(), nil
	}
	return selector
}

func (lbq *LoadBalanceQuery) prepareQuery(ctx context.Context) error {
	if lbq.path != nil {
		prev, err := lbq.path(ctx)
		if err != nil {
			return err
		}
		lbq.sql = prev
	}
	return nil
}

func (lbq *LoadBalanceQuery) sqlAll(ctx context.Context) ([]*LoadBalance, error) {
	var (
		nodes = []*LoadBalance{}
		_spec = lbq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &LoadBalance{config: lbq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, lbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lbq *LoadBalanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lbq.querySpec()
	return sqlgraph.CountNodes(ctx, lbq.driver, _spec)
}

func (lbq *LoadBalanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := lbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (lbq *LoadBalanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   loadbalance.Table,
			Columns: loadbalance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: loadbalance.FieldID,
			},
		},
		From:   lbq.sql,
		Unique: true,
	}
	if ps := lbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, loadbalance.ValidColumn)
			}
		}
	}
	return _spec
}

func (lbq *LoadBalanceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(lbq.driver.Dialect())
	t1 := builder.Table(loadbalance.Table)
	selector := builder.Select(t1.Columns(loadbalance.Columns...)...).From(t1)
	if lbq.sql != nil {
		selector = lbq.sql
		selector.Select(selector.Columns(loadbalance.Columns...)...)
	}
	for _, p := range lbq.predicates {
		p(selector)
	}
	for _, p := range lbq.order {
		p(selector, loadbalance.ValidColumn)
	}
	if offset := lbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LoadBalanceGroupBy is the builder for group-by LoadBalance entities.
type LoadBalanceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lbgb *LoadBalanceGroupBy) Aggregate(fns ...AggregateFunc) *LoadBalanceGroupBy {
	lbgb.fns = append(lbgb.fns, fns...)
	return lbgb
}

// Scan applies the group-by query and scan the result into the given value.
func (lbgb *LoadBalanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := lbgb.path(ctx)
	if err != nil {
		return err
	}
	lbgb.sql = query
	return lbgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := lbgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(lbgb.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := lbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := lbgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lbgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) StringX(ctx context.Context) string {
	v, err := lbgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(lbgb.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := lbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := lbgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lbgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) IntX(ctx context.Context) int {
	v, err := lbgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(lbgb.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := lbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := lbgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lbgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := lbgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(lbgb.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := lbgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := lbgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (lbgb *LoadBalanceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lbgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lbgb *LoadBalanceGroupBy) BoolX(ctx context.Context) bool {
	v, err := lbgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lbgb *LoadBalanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lbgb.fields {
		if !loadbalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lbgb *LoadBalanceGroupBy) sqlQuery() *sql.Selector {
	selector := lbgb.sql
	columns := make([]string, 0, len(lbgb.fields)+len(lbgb.fns))
	columns = append(columns, lbgb.fields...)
	for _, fn := range lbgb.fns {
		columns = append(columns, fn(selector, loadbalance.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(lbgb.fields...)
}

// LoadBalanceSelect is the builder for select fields of LoadBalance entities.
type LoadBalanceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (lbs *LoadBalanceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := lbs.path(ctx)
	if err != nil {
		return err
	}
	lbs.sql = query
	return lbs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (lbs *LoadBalanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := lbs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(lbs.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := lbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (lbs *LoadBalanceSelect) StringsX(ctx context.Context) []string {
	v, err := lbs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = lbs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (lbs *LoadBalanceSelect) StringX(ctx context.Context) string {
	v, err := lbs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(lbs.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := lbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (lbs *LoadBalanceSelect) IntsX(ctx context.Context) []int {
	v, err := lbs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = lbs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (lbs *LoadBalanceSelect) IntX(ctx context.Context) int {
	v, err := lbs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(lbs.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := lbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (lbs *LoadBalanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := lbs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = lbs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (lbs *LoadBalanceSelect) Float64X(ctx context.Context) float64 {
	v, err := lbs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(lbs.fields) > 1 {
		return nil, errors.New("ent: LoadBalanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := lbs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (lbs *LoadBalanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := lbs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (lbs *LoadBalanceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = lbs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{loadbalance.Label}
	default:
		err = fmt.Errorf("ent: LoadBalanceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (lbs *LoadBalanceSelect) BoolX(ctx context.Context) bool {
	v, err := lbs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lbs *LoadBalanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range lbs.fields {
		if !loadbalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := lbs.sqlQuery().Query()
	if err := lbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lbs *LoadBalanceSelect) sqlQuery() sql.Querier {
	selector := lbs.sql
	selector.Select(selector.Columns(lbs.fields...)...)
	return selector
}
