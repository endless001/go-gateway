// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-gateway/internal/ent/httprule"
	"go-gateway/internal/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// HttpRuleUpdate is the builder for updating HttpRule entities.
type HttpRuleUpdate struct {
	config
	hooks    []Hook
	mutation *HttpRuleMutation
}

// Where adds a new predicate for the builder.
func (hru *HttpRuleUpdate) Where(ps ...predicate.HttpRule) *HttpRuleUpdate {
	hru.mutation.predicates = append(hru.mutation.predicates, ps...)
	return hru
}

// SetServiceID sets the service_id field.
func (hru *HttpRuleUpdate) SetServiceID(i int64) *HttpRuleUpdate {
	hru.mutation.ResetServiceID()
	hru.mutation.SetServiceID(i)
	return hru
}

// AddServiceID adds i to service_id.
func (hru *HttpRuleUpdate) AddServiceID(i int64) *HttpRuleUpdate {
	hru.mutation.AddServiceID(i)
	return hru
}

// SetRuleType sets the rule_type field.
func (hru *HttpRuleUpdate) SetRuleType(i int) *HttpRuleUpdate {
	hru.mutation.ResetRuleType()
	hru.mutation.SetRuleType(i)
	return hru
}

// AddRuleType adds i to rule_type.
func (hru *HttpRuleUpdate) AddRuleType(i int) *HttpRuleUpdate {
	hru.mutation.AddRuleType(i)
	return hru
}

// SetRule sets the rule field.
func (hru *HttpRuleUpdate) SetRule(s string) *HttpRuleUpdate {
	hru.mutation.SetRule(s)
	return hru
}

// SetNeedHTTPS sets the need_https field.
func (hru *HttpRuleUpdate) SetNeedHTTPS(i int) *HttpRuleUpdate {
	hru.mutation.ResetNeedHTTPS()
	hru.mutation.SetNeedHTTPS(i)
	return hru
}

// AddNeedHTTPS adds i to need_https.
func (hru *HttpRuleUpdate) AddNeedHTTPS(i int) *HttpRuleUpdate {
	hru.mutation.AddNeedHTTPS(i)
	return hru
}

// SetNeedWebsocket sets the need_websocket field.
func (hru *HttpRuleUpdate) SetNeedWebsocket(i int) *HttpRuleUpdate {
	hru.mutation.ResetNeedWebsocket()
	hru.mutation.SetNeedWebsocket(i)
	return hru
}

// AddNeedWebsocket adds i to need_websocket.
func (hru *HttpRuleUpdate) AddNeedWebsocket(i int) *HttpRuleUpdate {
	hru.mutation.AddNeedWebsocket(i)
	return hru
}

// SetNeedStripURI sets the need_strip_uri field.
func (hru *HttpRuleUpdate) SetNeedStripURI(i int) *HttpRuleUpdate {
	hru.mutation.ResetNeedStripURI()
	hru.mutation.SetNeedStripURI(i)
	return hru
}

// AddNeedStripURI adds i to need_strip_uri.
func (hru *HttpRuleUpdate) AddNeedStripURI(i int) *HttpRuleUpdate {
	hru.mutation.AddNeedStripURI(i)
	return hru
}

// SetURLRewrite sets the url_rewrite field.
func (hru *HttpRuleUpdate) SetURLRewrite(s string) *HttpRuleUpdate {
	hru.mutation.SetURLRewrite(s)
	return hru
}

// SetHeaderTransfor sets the header_transfor field.
func (hru *HttpRuleUpdate) SetHeaderTransfor(s string) *HttpRuleUpdate {
	hru.mutation.SetHeaderTransfor(s)
	return hru
}

// Mutation returns the HttpRuleMutation object of the builder.
func (hru *HttpRuleUpdate) Mutation() *HttpRuleMutation {
	return hru.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (hru *HttpRuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hru.hooks) == 0 {
		affected, err = hru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HttpRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hru.mutation = mutation
			affected, err = hru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hru.hooks) - 1; i >= 0; i-- {
			mut = hru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (hru *HttpRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := hru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hru *HttpRuleUpdate) Exec(ctx context.Context) error {
	_, err := hru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hru *HttpRuleUpdate) ExecX(ctx context.Context) {
	if err := hru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hru *HttpRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   httprule.Table,
			Columns: httprule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: httprule.FieldID,
			},
		},
	}
	if ps := hru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hru.mutation.ServiceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: httprule.FieldServiceID,
		})
	}
	if value, ok := hru.mutation.AddedServiceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: httprule.FieldServiceID,
		})
	}
	if value, ok := hru.mutation.RuleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldRuleType,
		})
	}
	if value, ok := hru.mutation.AddedRuleType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldRuleType,
		})
	}
	if value, ok := hru.mutation.Rule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprule.FieldRule,
		})
	}
	if value, ok := hru.mutation.NeedHTTPS(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedHTTPS,
		})
	}
	if value, ok := hru.mutation.AddedNeedHTTPS(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedHTTPS,
		})
	}
	if value, ok := hru.mutation.NeedWebsocket(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedWebsocket,
		})
	}
	if value, ok := hru.mutation.AddedNeedWebsocket(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedWebsocket,
		})
	}
	if value, ok := hru.mutation.NeedStripURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedStripURI,
		})
	}
	if value, ok := hru.mutation.AddedNeedStripURI(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedStripURI,
		})
	}
	if value, ok := hru.mutation.URLRewrite(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprule.FieldURLRewrite,
		})
	}
	if value, ok := hru.mutation.HeaderTransfor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprule.FieldHeaderTransfor,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httprule.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// HttpRuleUpdateOne is the builder for updating a single HttpRule entity.
type HttpRuleUpdateOne struct {
	config
	hooks    []Hook
	mutation *HttpRuleMutation
}

// SetServiceID sets the service_id field.
func (hruo *HttpRuleUpdateOne) SetServiceID(i int64) *HttpRuleUpdateOne {
	hruo.mutation.ResetServiceID()
	hruo.mutation.SetServiceID(i)
	return hruo
}

// AddServiceID adds i to service_id.
func (hruo *HttpRuleUpdateOne) AddServiceID(i int64) *HttpRuleUpdateOne {
	hruo.mutation.AddServiceID(i)
	return hruo
}

// SetRuleType sets the rule_type field.
func (hruo *HttpRuleUpdateOne) SetRuleType(i int) *HttpRuleUpdateOne {
	hruo.mutation.ResetRuleType()
	hruo.mutation.SetRuleType(i)
	return hruo
}

// AddRuleType adds i to rule_type.
func (hruo *HttpRuleUpdateOne) AddRuleType(i int) *HttpRuleUpdateOne {
	hruo.mutation.AddRuleType(i)
	return hruo
}

// SetRule sets the rule field.
func (hruo *HttpRuleUpdateOne) SetRule(s string) *HttpRuleUpdateOne {
	hruo.mutation.SetRule(s)
	return hruo
}

// SetNeedHTTPS sets the need_https field.
func (hruo *HttpRuleUpdateOne) SetNeedHTTPS(i int) *HttpRuleUpdateOne {
	hruo.mutation.ResetNeedHTTPS()
	hruo.mutation.SetNeedHTTPS(i)
	return hruo
}

// AddNeedHTTPS adds i to need_https.
func (hruo *HttpRuleUpdateOne) AddNeedHTTPS(i int) *HttpRuleUpdateOne {
	hruo.mutation.AddNeedHTTPS(i)
	return hruo
}

// SetNeedWebsocket sets the need_websocket field.
func (hruo *HttpRuleUpdateOne) SetNeedWebsocket(i int) *HttpRuleUpdateOne {
	hruo.mutation.ResetNeedWebsocket()
	hruo.mutation.SetNeedWebsocket(i)
	return hruo
}

// AddNeedWebsocket adds i to need_websocket.
func (hruo *HttpRuleUpdateOne) AddNeedWebsocket(i int) *HttpRuleUpdateOne {
	hruo.mutation.AddNeedWebsocket(i)
	return hruo
}

// SetNeedStripURI sets the need_strip_uri field.
func (hruo *HttpRuleUpdateOne) SetNeedStripURI(i int) *HttpRuleUpdateOne {
	hruo.mutation.ResetNeedStripURI()
	hruo.mutation.SetNeedStripURI(i)
	return hruo
}

// AddNeedStripURI adds i to need_strip_uri.
func (hruo *HttpRuleUpdateOne) AddNeedStripURI(i int) *HttpRuleUpdateOne {
	hruo.mutation.AddNeedStripURI(i)
	return hruo
}

// SetURLRewrite sets the url_rewrite field.
func (hruo *HttpRuleUpdateOne) SetURLRewrite(s string) *HttpRuleUpdateOne {
	hruo.mutation.SetURLRewrite(s)
	return hruo
}

// SetHeaderTransfor sets the header_transfor field.
func (hruo *HttpRuleUpdateOne) SetHeaderTransfor(s string) *HttpRuleUpdateOne {
	hruo.mutation.SetHeaderTransfor(s)
	return hruo
}

// Mutation returns the HttpRuleMutation object of the builder.
func (hruo *HttpRuleUpdateOne) Mutation() *HttpRuleMutation {
	return hruo.mutation
}

// Save executes the query and returns the updated entity.
func (hruo *HttpRuleUpdateOne) Save(ctx context.Context) (*HttpRule, error) {
	var (
		err  error
		node *HttpRule
	)
	if len(hruo.hooks) == 0 {
		node, err = hruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HttpRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hruo.mutation = mutation
			node, err = hruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hruo.hooks) - 1; i >= 0; i-- {
			mut = hruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (hruo *HttpRuleUpdateOne) SaveX(ctx context.Context) *HttpRule {
	node, err := hruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hruo *HttpRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := hruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hruo *HttpRuleUpdateOne) ExecX(ctx context.Context) {
	if err := hruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hruo *HttpRuleUpdateOne) sqlSave(ctx context.Context) (_node *HttpRule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   httprule.Table,
			Columns: httprule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: httprule.FieldID,
			},
		},
	}
	id, ok := hruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing HttpRule.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := hruo.mutation.ServiceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: httprule.FieldServiceID,
		})
	}
	if value, ok := hruo.mutation.AddedServiceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: httprule.FieldServiceID,
		})
	}
	if value, ok := hruo.mutation.RuleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldRuleType,
		})
	}
	if value, ok := hruo.mutation.AddedRuleType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldRuleType,
		})
	}
	if value, ok := hruo.mutation.Rule(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprule.FieldRule,
		})
	}
	if value, ok := hruo.mutation.NeedHTTPS(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedHTTPS,
		})
	}
	if value, ok := hruo.mutation.AddedNeedHTTPS(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedHTTPS,
		})
	}
	if value, ok := hruo.mutation.NeedWebsocket(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedWebsocket,
		})
	}
	if value, ok := hruo.mutation.AddedNeedWebsocket(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedWebsocket,
		})
	}
	if value, ok := hruo.mutation.NeedStripURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedStripURI,
		})
	}
	if value, ok := hruo.mutation.AddedNeedStripURI(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: httprule.FieldNeedStripURI,
		})
	}
	if value, ok := hruo.mutation.URLRewrite(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprule.FieldURLRewrite,
		})
	}
	if value, ok := hruo.mutation.HeaderTransfor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: httprule.FieldHeaderTransfor,
		})
	}
	_node = &HttpRule{config: hruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, hruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{httprule.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
