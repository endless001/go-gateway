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

// HttpRuleDelete is the builder for deleting a HttpRule entity.
type HttpRuleDelete struct {
	config
	hooks    []Hook
	mutation *HttpRuleMutation
}

// Where adds a new predicate to the delete builder.
func (hrd *HttpRuleDelete) Where(ps ...predicate.HttpRule) *HttpRuleDelete {
	hrd.mutation.predicates = append(hrd.mutation.predicates, ps...)
	return hrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hrd *HttpRuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hrd.hooks) == 0 {
		affected, err = hrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HttpRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hrd.mutation = mutation
			affected, err = hrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hrd.hooks) - 1; i >= 0; i-- {
			mut = hrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hrd *HttpRuleDelete) ExecX(ctx context.Context) int {
	n, err := hrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hrd *HttpRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: httprule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: httprule.FieldID,
			},
		},
	}
	if ps := hrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, hrd.driver, _spec)
}

// HttpRuleDeleteOne is the builder for deleting a single HttpRule entity.
type HttpRuleDeleteOne struct {
	hrd *HttpRuleDelete
}

// Exec executes the deletion query.
func (hrdo *HttpRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := hrdo.hrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{httprule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hrdo *HttpRuleDeleteOne) ExecX(ctx context.Context) {
	hrdo.hrd.ExecX(ctx)
}
