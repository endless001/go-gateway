// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gateway/internal/ent/grpcrule"
	"go-gateway/internal/ent/predicate"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// GrpcRuleQuery is the builder for querying GrpcRule entities.
type GrpcRuleQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.GrpcRule
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (grq *GrpcRuleQuery) Where(ps ...predicate.GrpcRule) *GrpcRuleQuery {
	grq.predicates = append(grq.predicates, ps...)
	return grq
}

// Limit adds a limit step to the query.
func (grq *GrpcRuleQuery) Limit(limit int) *GrpcRuleQuery {
	grq.limit = &limit
	return grq
}

// Offset adds an offset step to the query.
func (grq *GrpcRuleQuery) Offset(offset int) *GrpcRuleQuery {
	grq.offset = &offset
	return grq
}

// Order adds an order step to the query.
func (grq *GrpcRuleQuery) Order(o ...OrderFunc) *GrpcRuleQuery {
	grq.order = append(grq.order, o...)
	return grq
}

// First returns the first GrpcRule entity in the query. Returns *NotFoundError when no grpcrule was found.
func (grq *GrpcRuleQuery) First(ctx context.Context) (*GrpcRule, error) {
	nodes, err := grq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grpcrule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (grq *GrpcRuleQuery) FirstX(ctx context.Context) *GrpcRule {
	node, err := grq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GrpcRule id in the query. Returns *NotFoundError when no id was found.
func (grq *GrpcRuleQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = grq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grpcrule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (grq *GrpcRuleQuery) FirstIDX(ctx context.Context) int64 {
	id, err := grq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only GrpcRule entity in the query, returns an error if not exactly one entity was returned.
func (grq *GrpcRuleQuery) Only(ctx context.Context) (*GrpcRule, error) {
	nodes, err := grq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grpcrule.Label}
	default:
		return nil, &NotSingularError{grpcrule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (grq *GrpcRuleQuery) OnlyX(ctx context.Context) *GrpcRule {
	node, err := grq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only GrpcRule id in the query, returns an error if not exactly one id was returned.
func (grq *GrpcRuleQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = grq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = &NotSingularError{grpcrule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (grq *GrpcRuleQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := grq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GrpcRules.
func (grq *GrpcRuleQuery) All(ctx context.Context) ([]*GrpcRule, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return grq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (grq *GrpcRuleQuery) AllX(ctx context.Context) []*GrpcRule {
	nodes, err := grq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GrpcRule ids.
func (grq *GrpcRuleQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := grq.Select(grpcrule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (grq *GrpcRuleQuery) IDsX(ctx context.Context) []int64 {
	ids, err := grq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (grq *GrpcRuleQuery) Count(ctx context.Context) (int, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return grq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (grq *GrpcRuleQuery) CountX(ctx context.Context) int {
	count, err := grq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (grq *GrpcRuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := grq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return grq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (grq *GrpcRuleQuery) ExistX(ctx context.Context) bool {
	exist, err := grq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (grq *GrpcRuleQuery) Clone() *GrpcRuleQuery {
	if grq == nil {
		return nil
	}
	return &GrpcRuleQuery{
		config:     grq.config,
		limit:      grq.limit,
		offset:     grq.offset,
		order:      append([]OrderFunc{}, grq.order...),
		unique:     append([]string{}, grq.unique...),
		predicates: append([]predicate.GrpcRule{}, grq.predicates...),
		// clone intermediate query.
		sql:  grq.sql.Clone(),
		path: grq.path,
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
//	client.GrpcRule.Query().
//		GroupBy(grpcrule.FieldServiceID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (grq *GrpcRuleQuery) GroupBy(field string, fields ...string) *GrpcRuleGroupBy {
	group := &GrpcRuleGroupBy{config: grq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return grq.sqlQuery(), nil
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
//	client.GrpcRule.Query().
//		Select(grpcrule.FieldServiceID).
//		Scan(ctx, &v)
//
func (grq *GrpcRuleQuery) Select(field string, fields ...string) *GrpcRuleSelect {
	selector := &GrpcRuleSelect{config: grq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := grq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return grq.sqlQuery(), nil
	}
	return selector
}

func (grq *GrpcRuleQuery) prepareQuery(ctx context.Context) error {
	if grq.path != nil {
		prev, err := grq.path(ctx)
		if err != nil {
			return err
		}
		grq.sql = prev
	}
	return nil
}

func (grq *GrpcRuleQuery) sqlAll(ctx context.Context) ([]*GrpcRule, error) {
	var (
		nodes = []*GrpcRule{}
		_spec = grq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &GrpcRule{config: grq.config}
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
	if err := sqlgraph.QueryNodes(ctx, grq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (grq *GrpcRuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := grq.querySpec()
	return sqlgraph.CountNodes(ctx, grq.driver, _spec)
}

func (grq *GrpcRuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := grq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (grq *GrpcRuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grpcrule.Table,
			Columns: grpcrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: grpcrule.FieldID,
			},
		},
		From:   grq.sql,
		Unique: true,
	}
	if ps := grq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := grq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := grq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := grq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, grpcrule.ValidColumn)
			}
		}
	}
	return _spec
}

func (grq *GrpcRuleQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(grq.driver.Dialect())
	t1 := builder.Table(grpcrule.Table)
	selector := builder.Select(t1.Columns(grpcrule.Columns...)...).From(t1)
	if grq.sql != nil {
		selector = grq.sql
		selector.Select(selector.Columns(grpcrule.Columns...)...)
	}
	for _, p := range grq.predicates {
		p(selector)
	}
	for _, p := range grq.order {
		p(selector, grpcrule.ValidColumn)
	}
	if offset := grq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := grq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GrpcRuleGroupBy is the builder for group-by GrpcRule entities.
type GrpcRuleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (grgb *GrpcRuleGroupBy) Aggregate(fns ...AggregateFunc) *GrpcRuleGroupBy {
	grgb.fns = append(grgb.fns, fns...)
	return grgb
}

// Scan applies the group-by query and scan the result into the given value.
func (grgb *GrpcRuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := grgb.path(ctx)
	if err != nil {
		return err
	}
	grgb.sql = query
	return grgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := grgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) StringsX(ctx context.Context) []string {
	v, err := grgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = grgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) StringX(ctx context.Context) string {
	v, err := grgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) IntsX(ctx context.Context) []int {
	v, err := grgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = grgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) IntX(ctx context.Context) int {
	v, err := grgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := grgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = grgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := grgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(grgb.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := grgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := grgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (grgb *GrpcRuleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = grgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (grgb *GrpcRuleGroupBy) BoolX(ctx context.Context) bool {
	v, err := grgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (grgb *GrpcRuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range grgb.fields {
		if !grpcrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := grgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := grgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (grgb *GrpcRuleGroupBy) sqlQuery() *sql.Selector {
	selector := grgb.sql
	columns := make([]string, 0, len(grgb.fields)+len(grgb.fns))
	columns = append(columns, grgb.fields...)
	for _, fn := range grgb.fns {
		columns = append(columns, fn(selector, grpcrule.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(grgb.fields...)
}

// GrpcRuleSelect is the builder for select fields of GrpcRule entities.
type GrpcRuleSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (grs *GrpcRuleSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := grs.path(ctx)
	if err != nil {
		return err
	}
	grs.sql = query
	return grs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (grs *GrpcRuleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := grs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (grs *GrpcRuleSelect) StringsX(ctx context.Context) []string {
	v, err := grs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = grs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (grs *GrpcRuleSelect) StringX(ctx context.Context) string {
	v, err := grs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (grs *GrpcRuleSelect) IntsX(ctx context.Context) []int {
	v, err := grs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = grs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (grs *GrpcRuleSelect) IntX(ctx context.Context) int {
	v, err := grs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (grs *GrpcRuleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := grs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = grs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (grs *GrpcRuleSelect) Float64X(ctx context.Context) float64 {
	v, err := grs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(grs.fields) > 1 {
		return nil, errors.New("ent: GrpcRuleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := grs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (grs *GrpcRuleSelect) BoolsX(ctx context.Context) []bool {
	v, err := grs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (grs *GrpcRuleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = grs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{grpcrule.Label}
	default:
		err = fmt.Errorf("ent: GrpcRuleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (grs *GrpcRuleSelect) BoolX(ctx context.Context) bool {
	v, err := grs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (grs *GrpcRuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range grs.fields {
		if !grpcrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := grs.sqlQuery().Query()
	if err := grs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (grs *GrpcRuleSelect) sqlQuery() sql.Querier {
	selector := grs.sql
	selector.Select(selector.Columns(grs.fields...)...)
	return selector
}
