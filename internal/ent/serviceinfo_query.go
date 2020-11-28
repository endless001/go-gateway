// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gateway/internal/ent/predicate"
	"go-gateway/internal/ent/serviceinfo"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ServiceInfoQuery is the builder for querying ServiceInfo entities.
type ServiceInfoQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.ServiceInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (siq *ServiceInfoQuery) Where(ps ...predicate.ServiceInfo) *ServiceInfoQuery {
	siq.predicates = append(siq.predicates, ps...)
	return siq
}

// Limit adds a limit step to the query.
func (siq *ServiceInfoQuery) Limit(limit int) *ServiceInfoQuery {
	siq.limit = &limit
	return siq
}

// Offset adds an offset step to the query.
func (siq *ServiceInfoQuery) Offset(offset int) *ServiceInfoQuery {
	siq.offset = &offset
	return siq
}

// Order adds an order step to the query.
func (siq *ServiceInfoQuery) Order(o ...OrderFunc) *ServiceInfoQuery {
	siq.order = append(siq.order, o...)
	return siq
}

// First returns the first ServiceInfo entity in the query. Returns *NotFoundError when no serviceinfo was found.
func (siq *ServiceInfoQuery) First(ctx context.Context) (*ServiceInfo, error) {
	nodes, err := siq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{serviceinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (siq *ServiceInfoQuery) FirstX(ctx context.Context) *ServiceInfo {
	node, err := siq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServiceInfo id in the query. Returns *NotFoundError when no id was found.
func (siq *ServiceInfoQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = siq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{serviceinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (siq *ServiceInfoQuery) FirstIDX(ctx context.Context) int64 {
	id, err := siq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ServiceInfo entity in the query, returns an error if not exactly one entity was returned.
func (siq *ServiceInfoQuery) Only(ctx context.Context) (*ServiceInfo, error) {
	nodes, err := siq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{serviceinfo.Label}
	default:
		return nil, &NotSingularError{serviceinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (siq *ServiceInfoQuery) OnlyX(ctx context.Context) *ServiceInfo {
	node, err := siq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only ServiceInfo id in the query, returns an error if not exactly one id was returned.
func (siq *ServiceInfoQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = siq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = &NotSingularError{serviceinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (siq *ServiceInfoQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := siq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServiceInfos.
func (siq *ServiceInfoQuery) All(ctx context.Context) ([]*ServiceInfo, error) {
	if err := siq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return siq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (siq *ServiceInfoQuery) AllX(ctx context.Context) []*ServiceInfo {
	nodes, err := siq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServiceInfo ids.
func (siq *ServiceInfoQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := siq.Select(serviceinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (siq *ServiceInfoQuery) IDsX(ctx context.Context) []int64 {
	ids, err := siq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (siq *ServiceInfoQuery) Count(ctx context.Context) (int, error) {
	if err := siq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return siq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (siq *ServiceInfoQuery) CountX(ctx context.Context) int {
	count, err := siq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (siq *ServiceInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := siq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return siq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (siq *ServiceInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := siq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (siq *ServiceInfoQuery) Clone() *ServiceInfoQuery {
	if siq == nil {
		return nil
	}
	return &ServiceInfoQuery{
		config:     siq.config,
		limit:      siq.limit,
		offset:     siq.offset,
		order:      append([]OrderFunc{}, siq.order...),
		unique:     append([]string{}, siq.unique...),
		predicates: append([]predicate.ServiceInfo{}, siq.predicates...),
		// clone intermediate query.
		sql:  siq.sql.Clone(),
		path: siq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LoadType int `json:"load_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServiceInfo.Query().
//		GroupBy(serviceinfo.FieldLoadType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (siq *ServiceInfoQuery) GroupBy(field string, fields ...string) *ServiceInfoGroupBy {
	group := &ServiceInfoGroupBy{config: siq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return siq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		LoadType int `json:"load_type,omitempty"`
//	}
//
//	client.ServiceInfo.Query().
//		Select(serviceinfo.FieldLoadType).
//		Scan(ctx, &v)
//
func (siq *ServiceInfoQuery) Select(field string, fields ...string) *ServiceInfoSelect {
	selector := &ServiceInfoSelect{config: siq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return siq.sqlQuery(), nil
	}
	return selector
}

func (siq *ServiceInfoQuery) prepareQuery(ctx context.Context) error {
	if siq.path != nil {
		prev, err := siq.path(ctx)
		if err != nil {
			return err
		}
		siq.sql = prev
	}
	return nil
}

func (siq *ServiceInfoQuery) sqlAll(ctx context.Context) ([]*ServiceInfo, error) {
	var (
		nodes = []*ServiceInfo{}
		_spec = siq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &ServiceInfo{config: siq.config}
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
	if err := sqlgraph.QueryNodes(ctx, siq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (siq *ServiceInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := siq.querySpec()
	return sqlgraph.CountNodes(ctx, siq.driver, _spec)
}

func (siq *ServiceInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := siq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (siq *ServiceInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceinfo.Table,
			Columns: serviceinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: serviceinfo.FieldID,
			},
		},
		From:   siq.sql,
		Unique: true,
	}
	if ps := siq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := siq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := siq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := siq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, serviceinfo.ValidColumn)
			}
		}
	}
	return _spec
}

func (siq *ServiceInfoQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(siq.driver.Dialect())
	t1 := builder.Table(serviceinfo.Table)
	selector := builder.Select(t1.Columns(serviceinfo.Columns...)...).From(t1)
	if siq.sql != nil {
		selector = siq.sql
		selector.Select(selector.Columns(serviceinfo.Columns...)...)
	}
	for _, p := range siq.predicates {
		p(selector)
	}
	for _, p := range siq.order {
		p(selector, serviceinfo.ValidColumn)
	}
	if offset := siq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := siq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ServiceInfoGroupBy is the builder for group-by ServiceInfo entities.
type ServiceInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sigb *ServiceInfoGroupBy) Aggregate(fns ...AggregateFunc) *ServiceInfoGroupBy {
	sigb.fns = append(sigb.fns, fns...)
	return sigb
}

// Scan applies the group-by query and scan the result into the given value.
func (sigb *ServiceInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sigb.path(ctx)
	if err != nil {
		return err
	}
	sigb.sql = query
	return sigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := sigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) StringX(ctx context.Context) string {
	v, err := sigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := sigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) IntX(ctx context.Context) int {
	v, err := sigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sigb.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (sigb *ServiceInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sigb *ServiceInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := sigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sigb *ServiceInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sigb.fields {
		if !serviceinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sigb *ServiceInfoGroupBy) sqlQuery() *sql.Selector {
	selector := sigb.sql
	columns := make([]string, 0, len(sigb.fields)+len(sigb.fns))
	columns = append(columns, sigb.fields...)
	for _, fn := range sigb.fns {
		columns = append(columns, fn(selector, serviceinfo.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(sigb.fields...)
}

// ServiceInfoSelect is the builder for select fields of ServiceInfo entities.
type ServiceInfoSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (sis *ServiceInfoSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := sis.path(ctx)
	if err != nil {
		return err
	}
	sis.sql = query
	return sis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sis *ServiceInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sis *ServiceInfoSelect) StringsX(ctx context.Context) []string {
	v, err := sis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sis *ServiceInfoSelect) StringX(ctx context.Context) string {
	v, err := sis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sis *ServiceInfoSelect) IntsX(ctx context.Context) []int {
	v, err := sis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sis *ServiceInfoSelect) IntX(ctx context.Context) int {
	v, err := sis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sis *ServiceInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sis *ServiceInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := sis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sis.fields) > 1 {
		return nil, errors.New("ent: ServiceInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sis *ServiceInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := sis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (sis *ServiceInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{serviceinfo.Label}
	default:
		err = fmt.Errorf("ent: ServiceInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sis *ServiceInfoSelect) BoolX(ctx context.Context) bool {
	v, err := sis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sis *ServiceInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sis.fields {
		if !serviceinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := sis.sqlQuery().Query()
	if err := sis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sis *ServiceInfoSelect) sqlQuery() sql.Querier {
	selector := sis.sql
	selector.Select(selector.Columns(sis.fields...)...)
	return selector
}