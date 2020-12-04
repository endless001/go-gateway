// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gateway/internal/ent/tenant"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TenantCreate is the builder for creating a Tenant entity.
type TenantCreate struct {
	config
	mutation *TenantMutation
	hooks    []Hook
}

// SetAppID sets the app_id field.
func (tc *TenantCreate) SetAppID(s string) *TenantCreate {
	tc.mutation.SetAppID(s)
	return tc
}

// SetName sets the name field.
func (tc *TenantCreate) SetName(s string) *TenantCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetSecret sets the secret field.
func (tc *TenantCreate) SetSecret(s string) *TenantCreate {
	tc.mutation.SetSecret(s)
	return tc
}

// SetWhiteIps sets the white_ips field.
func (tc *TenantCreate) SetWhiteIps(s string) *TenantCreate {
	tc.mutation.SetWhiteIps(s)
	return tc
}

// SetQpd sets the qpd field.
func (tc *TenantCreate) SetQpd(i int) *TenantCreate {
	tc.mutation.SetQpd(i)
	return tc
}

// SetQPS sets the qps field.
func (tc *TenantCreate) SetQPS(i int) *TenantCreate {
	tc.mutation.SetQPS(i)
	return tc
}

// SetCreateAt sets the create_at field.
func (tc *TenantCreate) SetCreateAt(t time.Time) *TenantCreate {
	tc.mutation.SetCreateAt(t)
	return tc
}

// SetUpdateAt sets the update_at field.
func (tc *TenantCreate) SetUpdateAt(t time.Time) *TenantCreate {
	tc.mutation.SetUpdateAt(t)
	return tc
}

// SetIsDelete sets the is_delete field.
func (tc *TenantCreate) SetIsDelete(i int8) *TenantCreate {
	tc.mutation.SetIsDelete(i)
	return tc
}

// SetID sets the id field.
func (tc *TenantCreate) SetID(i int64) *TenantCreate {
	tc.mutation.SetID(i)
	return tc
}

// Mutation returns the TenantMutation object of the builder.
func (tc *TenantCreate) Mutation() *TenantMutation {
	return tc.mutation
}

// Save creates the Tenant in the database.
func (tc *TenantCreate) Save(ctx context.Context) (*Tenant, error) {
	var (
		err  error
		node *Tenant
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TenantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TenantCreate) SaveX(ctx context.Context) *Tenant {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (tc *TenantCreate) check() error {
	if _, ok := tc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New("ent: missing required field \"app_id\"")}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := tc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New("ent: missing required field \"secret\"")}
	}
	if _, ok := tc.mutation.WhiteIps(); !ok {
		return &ValidationError{Name: "white_ips", err: errors.New("ent: missing required field \"white_ips\"")}
	}
	if _, ok := tc.mutation.Qpd(); !ok {
		return &ValidationError{Name: "qpd", err: errors.New("ent: missing required field \"qpd\"")}
	}
	if _, ok := tc.mutation.QPS(); !ok {
		return &ValidationError{Name: "qps", err: errors.New("ent: missing required field \"qps\"")}
	}
	if _, ok := tc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New("ent: missing required field \"create_at\"")}
	}
	if _, ok := tc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New("ent: missing required field \"update_at\"")}
	}
	if _, ok := tc.mutation.IsDelete(); !ok {
		return &ValidationError{Name: "is_delete", err: errors.New("ent: missing required field \"is_delete\"")}
	}
	return nil
}

func (tc *TenantCreate) sqlSave(ctx context.Context) (*Tenant, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (tc *TenantCreate) createSpec() (*Tenant, *sqlgraph.CreateSpec) {
	var (
		_node = &Tenant{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tenant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: tenant.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenant.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenant.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.Secret(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenant.FieldSecret,
		})
		_node.Secret = value
	}
	if value, ok := tc.mutation.WhiteIps(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tenant.FieldWhiteIps,
		})
		_node.WhiteIps = value
	}
	if value, ok := tc.mutation.Qpd(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tenant.FieldQpd,
		})
		_node.Qpd = value
	}
	if value, ok := tc.mutation.QPS(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tenant.FieldQPS,
		})
		_node.QPS = value
	}
	if value, ok := tc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tenant.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := tc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tenant.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := tc.mutation.IsDelete(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: tenant.FieldIsDelete,
		})
		_node.IsDelete = value
	}
	return _node, _spec
}

// TenantCreateBulk is the builder for creating a bulk of Tenant entities.
type TenantCreateBulk struct {
	config
	builders []*TenantCreate
}

// Save creates the Tenant entities in the database.
func (tcb *TenantCreateBulk) Save(ctx context.Context) ([]*Tenant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tenant, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TenantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (tcb *TenantCreateBulk) SaveX(ctx context.Context) []*Tenant {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
