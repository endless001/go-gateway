// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-gateway/internal/ent/accesscontrol"
	"go-gateway/internal/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AccessControlUpdate is the builder for updating AccessControl entities.
type AccessControlUpdate struct {
	config
	hooks    []Hook
	mutation *AccessControlMutation
}

// Where adds a new predicate for the builder.
func (acu *AccessControlUpdate) Where(ps ...predicate.AccessControl) *AccessControlUpdate {
	acu.mutation.predicates = append(acu.mutation.predicates, ps...)
	return acu
}

// SetServiceID sets the service_id field.
func (acu *AccessControlUpdate) SetServiceID(i int64) *AccessControlUpdate {
	acu.mutation.ResetServiceID()
	acu.mutation.SetServiceID(i)
	return acu
}

// AddServiceID adds i to service_id.
func (acu *AccessControlUpdate) AddServiceID(i int64) *AccessControlUpdate {
	acu.mutation.AddServiceID(i)
	return acu
}

// SetOpenAuth sets the open_auth field.
func (acu *AccessControlUpdate) SetOpenAuth(i int) *AccessControlUpdate {
	acu.mutation.ResetOpenAuth()
	acu.mutation.SetOpenAuth(i)
	return acu
}

// AddOpenAuth adds i to open_auth.
func (acu *AccessControlUpdate) AddOpenAuth(i int) *AccessControlUpdate {
	acu.mutation.AddOpenAuth(i)
	return acu
}

// SetBlackList sets the black_list field.
func (acu *AccessControlUpdate) SetBlackList(s string) *AccessControlUpdate {
	acu.mutation.SetBlackList(s)
	return acu
}

// SetWhiteList sets the white_list field.
func (acu *AccessControlUpdate) SetWhiteList(s string) *AccessControlUpdate {
	acu.mutation.SetWhiteList(s)
	return acu
}

// SetWhiteHostName sets the white_host_name field.
func (acu *AccessControlUpdate) SetWhiteHostName(s string) *AccessControlUpdate {
	acu.mutation.SetWhiteHostName(s)
	return acu
}

// SetClientipFlowLimit sets the clientip_flow_limit field.
func (acu *AccessControlUpdate) SetClientipFlowLimit(i int) *AccessControlUpdate {
	acu.mutation.ResetClientipFlowLimit()
	acu.mutation.SetClientipFlowLimit(i)
	return acu
}

// AddClientipFlowLimit adds i to clientip_flow_limit.
func (acu *AccessControlUpdate) AddClientipFlowLimit(i int) *AccessControlUpdate {
	acu.mutation.AddClientipFlowLimit(i)
	return acu
}

// SetServiceFlowLimit sets the service_flow_limit field.
func (acu *AccessControlUpdate) SetServiceFlowLimit(i int) *AccessControlUpdate {
	acu.mutation.ResetServiceFlowLimit()
	acu.mutation.SetServiceFlowLimit(i)
	return acu
}

// AddServiceFlowLimit adds i to service_flow_limit.
func (acu *AccessControlUpdate) AddServiceFlowLimit(i int) *AccessControlUpdate {
	acu.mutation.AddServiceFlowLimit(i)
	return acu
}

// Mutation returns the AccessControlMutation object of the builder.
func (acu *AccessControlUpdate) Mutation() *AccessControlMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AccessControlUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(acu.hooks) == 0 {
		affected, err = acu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccessControlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acu.mutation = mutation
			affected, err = acu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acu.hooks) - 1; i >= 0; i-- {
			mut = acu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AccessControlUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AccessControlUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AccessControlUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acu *AccessControlUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accesscontrol.Table,
			Columns: accesscontrol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: accesscontrol.FieldID,
			},
		},
	}
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.ServiceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: accesscontrol.FieldServiceID,
		})
	}
	if value, ok := acu.mutation.AddedServiceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: accesscontrol.FieldServiceID,
		})
	}
	if value, ok := acu.mutation.OpenAuth(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldOpenAuth,
		})
	}
	if value, ok := acu.mutation.AddedOpenAuth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldOpenAuth,
		})
	}
	if value, ok := acu.mutation.BlackList(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesscontrol.FieldBlackList,
		})
	}
	if value, ok := acu.mutation.WhiteList(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesscontrol.FieldWhiteList,
		})
	}
	if value, ok := acu.mutation.WhiteHostName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesscontrol.FieldWhiteHostName,
		})
	}
	if value, ok := acu.mutation.ClientipFlowLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldClientipFlowLimit,
		})
	}
	if value, ok := acu.mutation.AddedClientipFlowLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldClientipFlowLimit,
		})
	}
	if value, ok := acu.mutation.ServiceFlowLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldServiceFlowLimit,
		})
	}
	if value, ok := acu.mutation.AddedServiceFlowLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldServiceFlowLimit,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesscontrol.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AccessControlUpdateOne is the builder for updating a single AccessControl entity.
type AccessControlUpdateOne struct {
	config
	hooks    []Hook
	mutation *AccessControlMutation
}

// SetServiceID sets the service_id field.
func (acuo *AccessControlUpdateOne) SetServiceID(i int64) *AccessControlUpdateOne {
	acuo.mutation.ResetServiceID()
	acuo.mutation.SetServiceID(i)
	return acuo
}

// AddServiceID adds i to service_id.
func (acuo *AccessControlUpdateOne) AddServiceID(i int64) *AccessControlUpdateOne {
	acuo.mutation.AddServiceID(i)
	return acuo
}

// SetOpenAuth sets the open_auth field.
func (acuo *AccessControlUpdateOne) SetOpenAuth(i int) *AccessControlUpdateOne {
	acuo.mutation.ResetOpenAuth()
	acuo.mutation.SetOpenAuth(i)
	return acuo
}

// AddOpenAuth adds i to open_auth.
func (acuo *AccessControlUpdateOne) AddOpenAuth(i int) *AccessControlUpdateOne {
	acuo.mutation.AddOpenAuth(i)
	return acuo
}

// SetBlackList sets the black_list field.
func (acuo *AccessControlUpdateOne) SetBlackList(s string) *AccessControlUpdateOne {
	acuo.mutation.SetBlackList(s)
	return acuo
}

// SetWhiteList sets the white_list field.
func (acuo *AccessControlUpdateOne) SetWhiteList(s string) *AccessControlUpdateOne {
	acuo.mutation.SetWhiteList(s)
	return acuo
}

// SetWhiteHostName sets the white_host_name field.
func (acuo *AccessControlUpdateOne) SetWhiteHostName(s string) *AccessControlUpdateOne {
	acuo.mutation.SetWhiteHostName(s)
	return acuo
}

// SetClientipFlowLimit sets the clientip_flow_limit field.
func (acuo *AccessControlUpdateOne) SetClientipFlowLimit(i int) *AccessControlUpdateOne {
	acuo.mutation.ResetClientipFlowLimit()
	acuo.mutation.SetClientipFlowLimit(i)
	return acuo
}

// AddClientipFlowLimit adds i to clientip_flow_limit.
func (acuo *AccessControlUpdateOne) AddClientipFlowLimit(i int) *AccessControlUpdateOne {
	acuo.mutation.AddClientipFlowLimit(i)
	return acuo
}

// SetServiceFlowLimit sets the service_flow_limit field.
func (acuo *AccessControlUpdateOne) SetServiceFlowLimit(i int) *AccessControlUpdateOne {
	acuo.mutation.ResetServiceFlowLimit()
	acuo.mutation.SetServiceFlowLimit(i)
	return acuo
}

// AddServiceFlowLimit adds i to service_flow_limit.
func (acuo *AccessControlUpdateOne) AddServiceFlowLimit(i int) *AccessControlUpdateOne {
	acuo.mutation.AddServiceFlowLimit(i)
	return acuo
}

// Mutation returns the AccessControlMutation object of the builder.
func (acuo *AccessControlUpdateOne) Mutation() *AccessControlMutation {
	return acuo.mutation
}

// Save executes the query and returns the updated entity.
func (acuo *AccessControlUpdateOne) Save(ctx context.Context) (*AccessControl, error) {
	var (
		err  error
		node *AccessControl
	)
	if len(acuo.hooks) == 0 {
		node, err = acuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AccessControlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			acuo.mutation = mutation
			node, err = acuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acuo.hooks) - 1; i >= 0; i-- {
			mut = acuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AccessControlUpdateOne) SaveX(ctx context.Context) *AccessControl {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AccessControlUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AccessControlUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acuo *AccessControlUpdateOne) sqlSave(ctx context.Context) (_node *AccessControl, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accesscontrol.Table,
			Columns: accesscontrol.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: accesscontrol.FieldID,
			},
		},
	}
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AccessControl.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := acuo.mutation.ServiceID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: accesscontrol.FieldServiceID,
		})
	}
	if value, ok := acuo.mutation.AddedServiceID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: accesscontrol.FieldServiceID,
		})
	}
	if value, ok := acuo.mutation.OpenAuth(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldOpenAuth,
		})
	}
	if value, ok := acuo.mutation.AddedOpenAuth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldOpenAuth,
		})
	}
	if value, ok := acuo.mutation.BlackList(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesscontrol.FieldBlackList,
		})
	}
	if value, ok := acuo.mutation.WhiteList(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesscontrol.FieldWhiteList,
		})
	}
	if value, ok := acuo.mutation.WhiteHostName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: accesscontrol.FieldWhiteHostName,
		})
	}
	if value, ok := acuo.mutation.ClientipFlowLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldClientipFlowLimit,
		})
	}
	if value, ok := acuo.mutation.AddedClientipFlowLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldClientipFlowLimit,
		})
	}
	if value, ok := acuo.mutation.ServiceFlowLimit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldServiceFlowLimit,
		})
	}
	if value, ok := acuo.mutation.AddedServiceFlowLimit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: accesscontrol.FieldServiceFlowLimit,
		})
	}
	_node = &AccessControl{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesscontrol.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
