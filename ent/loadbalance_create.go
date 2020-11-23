// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-gateway/ent/loadbalance"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// LoadBalanceCreate is the builder for creating a LoadBalance entity.
type LoadBalanceCreate struct {
	config
	mutation *LoadBalanceMutation
	hooks    []Hook
}

// SetServiceID sets the service_id field.
func (lbc *LoadBalanceCreate) SetServiceID(i int64) *LoadBalanceCreate {
	lbc.mutation.SetServiceID(i)
	return lbc
}

// SetCheckMethod sets the check_method field.
func (lbc *LoadBalanceCreate) SetCheckMethod(i int) *LoadBalanceCreate {
	lbc.mutation.SetCheckMethod(i)
	return lbc
}

// SetCheckTimeout sets the check_timeout field.
func (lbc *LoadBalanceCreate) SetCheckTimeout(i int) *LoadBalanceCreate {
	lbc.mutation.SetCheckTimeout(i)
	return lbc
}

// SetCheckInterval sets the check_interval field.
func (lbc *LoadBalanceCreate) SetCheckInterval(i int) *LoadBalanceCreate {
	lbc.mutation.SetCheckInterval(i)
	return lbc
}

// SetRoundType sets the round_type field.
func (lbc *LoadBalanceCreate) SetRoundType(i int) *LoadBalanceCreate {
	lbc.mutation.SetRoundType(i)
	return lbc
}

// SetIPList sets the ip_list field.
func (lbc *LoadBalanceCreate) SetIPList(s string) *LoadBalanceCreate {
	lbc.mutation.SetIPList(s)
	return lbc
}

// SetWeightList sets the weight_list field.
func (lbc *LoadBalanceCreate) SetWeightList(s string) *LoadBalanceCreate {
	lbc.mutation.SetWeightList(s)
	return lbc
}

// SetForbidList sets the forbid_list field.
func (lbc *LoadBalanceCreate) SetForbidList(s string) *LoadBalanceCreate {
	lbc.mutation.SetForbidList(s)
	return lbc
}

// SetUpstreamConnectTimeout sets the upstream_connect_timeout field.
func (lbc *LoadBalanceCreate) SetUpstreamConnectTimeout(i int) *LoadBalanceCreate {
	lbc.mutation.SetUpstreamConnectTimeout(i)
	return lbc
}

// SetUpstreamHeaderTimeout sets the upstream_header_timeout field.
func (lbc *LoadBalanceCreate) SetUpstreamHeaderTimeout(i int) *LoadBalanceCreate {
	lbc.mutation.SetUpstreamHeaderTimeout(i)
	return lbc
}

// SetUpstreamIdleTimeout sets the upstream_idle_timeout field.
func (lbc *LoadBalanceCreate) SetUpstreamIdleTimeout(i int) *LoadBalanceCreate {
	lbc.mutation.SetUpstreamIdleTimeout(i)
	return lbc
}

// SetUpstreamMaxIdle sets the upstream_max_idle field.
func (lbc *LoadBalanceCreate) SetUpstreamMaxIdle(i int) *LoadBalanceCreate {
	lbc.mutation.SetUpstreamMaxIdle(i)
	return lbc
}

// SetID sets the id field.
func (lbc *LoadBalanceCreate) SetID(i int64) *LoadBalanceCreate {
	lbc.mutation.SetID(i)
	return lbc
}

// Mutation returns the LoadBalanceMutation object of the builder.
func (lbc *LoadBalanceCreate) Mutation() *LoadBalanceMutation {
	return lbc.mutation
}

// Save creates the LoadBalance in the database.
func (lbc *LoadBalanceCreate) Save(ctx context.Context) (*LoadBalance, error) {
	var (
		err  error
		node *LoadBalance
	)
	if len(lbc.hooks) == 0 {
		if err = lbc.check(); err != nil {
			return nil, err
		}
		node, err = lbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LoadBalanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lbc.check(); err != nil {
				return nil, err
			}
			lbc.mutation = mutation
			node, err = lbc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lbc.hooks) - 1; i >= 0; i-- {
			mut = lbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lbc *LoadBalanceCreate) SaveX(ctx context.Context) *LoadBalance {
	v, err := lbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (lbc *LoadBalanceCreate) check() error {
	if _, ok := lbc.mutation.ServiceID(); !ok {
		return &ValidationError{Name: "service_id", err: errors.New("ent: missing required field \"service_id\"")}
	}
	if _, ok := lbc.mutation.CheckMethod(); !ok {
		return &ValidationError{Name: "check_method", err: errors.New("ent: missing required field \"check_method\"")}
	}
	if _, ok := lbc.mutation.CheckTimeout(); !ok {
		return &ValidationError{Name: "check_timeout", err: errors.New("ent: missing required field \"check_timeout\"")}
	}
	if _, ok := lbc.mutation.CheckInterval(); !ok {
		return &ValidationError{Name: "check_interval", err: errors.New("ent: missing required field \"check_interval\"")}
	}
	if _, ok := lbc.mutation.RoundType(); !ok {
		return &ValidationError{Name: "round_type", err: errors.New("ent: missing required field \"round_type\"")}
	}
	if _, ok := lbc.mutation.IPList(); !ok {
		return &ValidationError{Name: "ip_list", err: errors.New("ent: missing required field \"ip_list\"")}
	}
	if _, ok := lbc.mutation.WeightList(); !ok {
		return &ValidationError{Name: "weight_list", err: errors.New("ent: missing required field \"weight_list\"")}
	}
	if _, ok := lbc.mutation.ForbidList(); !ok {
		return &ValidationError{Name: "forbid_list", err: errors.New("ent: missing required field \"forbid_list\"")}
	}
	if _, ok := lbc.mutation.UpstreamConnectTimeout(); !ok {
		return &ValidationError{Name: "upstream_connect_timeout", err: errors.New("ent: missing required field \"upstream_connect_timeout\"")}
	}
	if _, ok := lbc.mutation.UpstreamHeaderTimeout(); !ok {
		return &ValidationError{Name: "upstream_header_timeout", err: errors.New("ent: missing required field \"upstream_header_timeout\"")}
	}
	if _, ok := lbc.mutation.UpstreamIdleTimeout(); !ok {
		return &ValidationError{Name: "upstream_idle_timeout", err: errors.New("ent: missing required field \"upstream_idle_timeout\"")}
	}
	if _, ok := lbc.mutation.UpstreamMaxIdle(); !ok {
		return &ValidationError{Name: "upstream_max_idle", err: errors.New("ent: missing required field \"upstream_max_idle\"")}
	}
	return nil
}

func (lbc *LoadBalanceCreate) sqlSave(ctx context.Context) (*LoadBalance, error) {
	_node, _spec := lbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lbc.driver, _spec); err != nil {
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

func (lbc *LoadBalanceCreate) createSpec() (*LoadBalance, *sqlgraph.CreateSpec) {
	var (
		_node = &LoadBalance{config: lbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: loadbalance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: loadbalance.FieldID,
			},
		}
	)
	if id, ok := lbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lbc.mutation.ServiceID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: loadbalance.FieldServiceID,
		})
		_node.ServiceID = value
	}
	if value, ok := lbc.mutation.CheckMethod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldCheckMethod,
		})
		_node.CheckMethod = value
	}
	if value, ok := lbc.mutation.CheckTimeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldCheckTimeout,
		})
		_node.CheckTimeout = value
	}
	if value, ok := lbc.mutation.CheckInterval(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldCheckInterval,
		})
		_node.CheckInterval = value
	}
	if value, ok := lbc.mutation.RoundType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldRoundType,
		})
		_node.RoundType = value
	}
	if value, ok := lbc.mutation.IPList(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loadbalance.FieldIPList,
		})
		_node.IPList = value
	}
	if value, ok := lbc.mutation.WeightList(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loadbalance.FieldWeightList,
		})
		_node.WeightList = value
	}
	if value, ok := lbc.mutation.ForbidList(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: loadbalance.FieldForbidList,
		})
		_node.ForbidList = value
	}
	if value, ok := lbc.mutation.UpstreamConnectTimeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldUpstreamConnectTimeout,
		})
		_node.UpstreamConnectTimeout = value
	}
	if value, ok := lbc.mutation.UpstreamHeaderTimeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldUpstreamHeaderTimeout,
		})
		_node.UpstreamHeaderTimeout = value
	}
	if value, ok := lbc.mutation.UpstreamIdleTimeout(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldUpstreamIdleTimeout,
		})
		_node.UpstreamIdleTimeout = value
	}
	if value, ok := lbc.mutation.UpstreamMaxIdle(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: loadbalance.FieldUpstreamMaxIdle,
		})
		_node.UpstreamMaxIdle = value
	}
	return _node, _spec
}

// LoadBalanceCreateBulk is the builder for creating a bulk of LoadBalance entities.
type LoadBalanceCreateBulk struct {
	config
	builders []*LoadBalanceCreate
}

// Save creates the LoadBalance entities in the database.
func (lbcb *LoadBalanceCreateBulk) Save(ctx context.Context) ([]*LoadBalance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lbcb.builders))
	nodes := make([]*LoadBalance, len(lbcb.builders))
	mutators := make([]Mutator, len(lbcb.builders))
	for i := range lbcb.builders {
		func(i int, root context.Context) {
			builder := lbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoadBalanceMutation)
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
					_, err = mutators[i+1].Mutate(root, lbcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lbcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (lbcb *LoadBalanceCreateBulk) SaveX(ctx context.Context) []*LoadBalance {
	v, err := lbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
