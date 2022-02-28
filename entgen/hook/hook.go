// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"example.com/enumeg/entgen"
)

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *entgen.GroupMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.GroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.GroupMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *entgen.UserMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.UserMutation", m)
	}
	return f(ctx, mv)
}

// The VideoFunc type is an adapter to allow the use of ordinary
// function as Video mutator.
type VideoFunc func(context.Context, *entgen.VideoMutation) (entgen.Value, error)

// Mutate calls f(ctx, m).
func (f VideoFunc) Mutate(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
	mv, ok := m.(*entgen.VideoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entgen.VideoMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, entgen.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entgen.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entgen.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m entgen.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op entgen.Op) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entgen.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk entgen.Hook, cond Condition) entgen.Hook {
	return func(next entgen.Mutator) entgen.Mutator {
		return entgen.MutateFunc(func(ctx context.Context, m entgen.Mutation) (entgen.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, entgen.Delete|entgen.Create)
//
func On(hk entgen.Hook, op entgen.Op) entgen.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, entgen.Update|entgen.UpdateOne)
//
func Unless(hk entgen.Hook, op entgen.Op) entgen.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) entgen.Hook {
	return func(entgen.Mutator) entgen.Mutator {
		return entgen.MutateFunc(func(context.Context, entgen.Mutation) (entgen.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []entgen.Hook {
//		return []entgen.Hook{
//			Reject(entgen.Delete|entgen.Update),
//		}
//	}
//
func Reject(op entgen.Op) entgen.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []entgen.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...entgen.Hook) Chain {
	return Chain{append([]entgen.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() entgen.Hook {
	return func(mutator entgen.Mutator) entgen.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...entgen.Hook) Chain {
	newHooks := make([]entgen.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
