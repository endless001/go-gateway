// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gateway/internal/ent/tcprule"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TcpRuleCreate is the builder for creating a TcpRule entity.
type TcpRuleCreate struct {
	config
	mutation *TcpRuleMutation
	hooks    []Hook
}

// SetServiceID sets the service_id field.
func (trc *TcpRuleCreate) SetServiceID(i int64) *TcpRuleCreate {
	trc.mutation.SetServiceID(i)
	return trc
}

// SetPort sets the Port field.
func (trc *TcpRuleCreate) SetPort(i int) *TcpRuleCreate {
	trc.mutation.SetPort(i)
	return trc
}

// SetID sets the id field.
func (trc *TcpRuleCreate) SetID(i int64) *TcpRuleCreate {
	trc.mutation.SetID(i)
	return trc
}

// Mutation returns the TcpRuleMutation object of the builder.
func (trc *TcpRuleCreate) Mutation() *TcpRuleMutation {
	return trc.mutation
}

// Save creates the TcpRule in the database.
func (trc *TcpRuleCreate) Save(ctx context.Context) (*TcpRule, error) {
	var (
		err  error
		node *TcpRule
	)
	if len(trc.hooks) == 0 {
		if err = trc.check(); err != nil {
			return nil, err
		}
		node, err = trc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TcpRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = trc.check(); err != nil {
				return nil, err
			}
			trc.mutation = mutation
			node, err = trc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(trc.hooks) - 1; i >= 0; i-- {
			mut = trc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, trc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (trc *TcpRuleCreate) SaveX(ctx context.Context) *TcpRule {
	v, err := trc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (trc *TcpRuleCreate) check() error {
	if _, ok := trc.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "service_id", err: errors.New("ent: missing required field \"service_id\"")}
	}
	if _, ok := trc.mutation.Port(); !ok {
		return &ValidationError{Name: "Port", err: errors.New("ent: missing required field \"Port\"")}
	}
	return nil
}

func (trc *TcpRuleCreate) sqlSave(ctx context.Context) (*TcpRule, error) {
	_node, _spec := trc.createSpec()
	if err := sqlgraph.CreateNode(ctx, trc.driver, _spec); err != nil {
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

func (trc *TcpRuleCreate) createSpec() (*TcpRule, *sqlgraph.CreateSpec) {
	var (
		_node = &TcpRule{config: trc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tcprule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: tcprule.FieldID,
			},
		}
	)
	if id, ok := trc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := trc.mutation.ServiceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: tcprule.FieldServiceID,
		})
		_node.ServiceID = value
	}
	if value, ok := trc.mutation.Port(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tcprule.FieldPort,
		})
		_node.Port = value
	}
	return _node, _spec
}

// TcpRuleCreateBulk is the builder for creating a bulk of TcpRule entities.
type TcpRuleCreateBulk struct {
	config
	builders []*TcpRuleCreate
}

// Save creates the TcpRule entities in the database.
func (trcb *TcpRuleCreateBulk) Save(ctx context.Context) ([]*TcpRule, error) {
	specs := make([]*sqlgraph.CreateSpec, len(trcb.builders))
	nodes := make([]*TcpRule, len(trcb.builders))
	mutators := make([]Mutator, len(trcb.builders))
	for i := range trcb.builders {
		func(i int, root context.Context) {
			builder := trcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TcpRuleMutation)
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
					_, err = mutators[i+1].Mutate(root, trcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, trcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, trcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (trcb *TcpRuleCreateBulk) SaveX(ctx context.Context) []*TcpRule {
	v, err := trcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
