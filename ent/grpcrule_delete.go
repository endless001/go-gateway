// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-gateway/ent/grpcrule"
	"go-gateway/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// GrpcRuleDelete is the builder for deleting a GrpcRule entity.
type GrpcRuleDelete struct {
	config
	hooks    []Hook
	mutation *GrpcRuleMutation
}

// Where adds a new predicate to the delete builder.
func (grd *GrpcRuleDelete) Where(ps ...predicate.GrpcRule) *GrpcRuleDelete {
	grd.mutation.predicates = append(grd.mutation.predicates, ps...)
	return grd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (grd *GrpcRuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(grd.hooks) == 0 {
		affected, err = grd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GrpcRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			grd.mutation = mutation
			affected, err = grd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(grd.hooks) - 1; i >= 0; i-- {
			mut = grd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, grd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (grd *GrpcRuleDelete) ExecX(ctx context.Context) int {
	n, err := grd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (grd *GrpcRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: grpcrule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: grpcrule.FieldID,
			},
		},
	}
	if ps := grd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, grd.driver, _spec)
}

// GrpcRuleDeleteOne is the builder for deleting a single GrpcRule entity.
type GrpcRuleDeleteOne struct {
	grd *GrpcRuleDelete
}

// Exec executes the deletion query.
func (grdo *GrpcRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := grdo.grd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{grpcrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (grdo *GrpcRuleDeleteOne) ExecX(ctx context.Context) {
	grdo.grd.ExecX(ctx)
}
