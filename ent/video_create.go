// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"example.com/enumeg/ent/group"
	"example.com/enumeg/ent/user"
	"example.com/enumeg/ent/video"
)

// VideoCreate is the builder for creating a Video entity.
type VideoCreate struct {
	config
	mutation *VideoMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (vc *VideoCreate) SetTitle(s string) *VideoCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetVideotype sets the "videotype" field.
func (vc *VideoCreate) SetVideotype(v video.Videotype) *VideoCreate {
	vc.mutation.SetVideotype(v)
	return vc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (vc *VideoCreate) SetGroupID(id int) *VideoCreate {
	vc.mutation.SetGroupID(id)
	return vc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (vc *VideoCreate) SetNillableGroupID(id *int) *VideoCreate {
	if id != nil {
		vc = vc.SetGroupID(*id)
	}
	return vc
}

// SetGroup sets the "group" edge to the Group entity.
func (vc *VideoCreate) SetGroup(g *Group) *VideoCreate {
	return vc.SetGroupID(g.ID)
}

// AddStreamerIDs adds the "streamers" edge to the User entity by IDs.
func (vc *VideoCreate) AddStreamerIDs(ids ...int) *VideoCreate {
	vc.mutation.AddStreamerIDs(ids...)
	return vc
}

// AddStreamers adds the "streamers" edges to the User entity.
func (vc *VideoCreate) AddStreamers(u ...*User) *VideoCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return vc.AddStreamerIDs(ids...)
}

// AddModeraterIDs adds the "moderaters" edge to the User entity by IDs.
func (vc *VideoCreate) AddModeraterIDs(ids ...int) *VideoCreate {
	vc.mutation.AddModeraterIDs(ids...)
	return vc
}

// AddModeraters adds the "moderaters" edges to the User entity.
func (vc *VideoCreate) AddModeraters(u ...*User) *VideoCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return vc.AddModeraterIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the User entity by IDs.
func (vc *VideoCreate) AddLikeIDs(ids ...int) *VideoCreate {
	vc.mutation.AddLikeIDs(ids...)
	return vc
}

// AddLikes adds the "likes" edges to the User entity.
func (vc *VideoCreate) AddLikes(u ...*User) *VideoCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return vc.AddLikeIDs(ids...)
}

// AddViewerIDs adds the "viewers" edge to the User entity by IDs.
func (vc *VideoCreate) AddViewerIDs(ids ...int) *VideoCreate {
	vc.mutation.AddViewerIDs(ids...)
	return vc
}

// AddViewers adds the "viewers" edges to the User entity.
func (vc *VideoCreate) AddViewers(u ...*User) *VideoCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return vc.AddViewerIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vc *VideoCreate) Mutation() *VideoMutation {
	return vc.mutation
}

// Save creates the Video in the database.
func (vc *VideoCreate) Save(ctx context.Context) (*Video, error) {
	var (
		err  error
		node *Video
	)
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			if node, err = vc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			if vc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideoCreate) SaveX(ctx context.Context) *Video {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideoCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideoCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideoCreate) check() error {
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Video.title"`)}
	}
	if _, ok := vc.mutation.Videotype(); !ok {
		return &ValidationError{Name: "videotype", err: errors.New(`ent: missing required field "Video.videotype"`)}
	}
	if v, ok := vc.mutation.Videotype(); ok {
		if err := video.VideotypeValidator(v); err != nil {
			return &ValidationError{Name: "videotype", err: fmt.Errorf(`ent: validator failed for field "Video.videotype": %w`, err)}
		}
	}
	return nil
}

func (vc *VideoCreate) sqlSave(ctx context.Context) (*Video, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (vc *VideoCreate) createSpec() (*Video, *sqlgraph.CreateSpec) {
	var (
		_node = &Video{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: video.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: video.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: video.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := vc.mutation.Videotype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: video.FieldVideotype,
		})
		_node.Videotype = value
	}
	if nodes := vc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   video.GroupTable,
			Columns: []string{video.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_videos = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.StreamersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   video.StreamersTable,
			Columns: video.StreamersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ModeratersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   video.ModeratersTable,
			Columns: video.ModeratersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   video.LikesTable,
			Columns: video.LikesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.ViewersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   video.ViewersTable,
			Columns: video.ViewersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideoCreateBulk is the builder for creating many Video entities in bulk.
type VideoCreateBulk struct {
	config
	builders []*VideoCreate
}

// Save creates the Video entities in the database.
func (vcb *VideoCreateBulk) Save(ctx context.Context) ([]*Video, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Video, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideoCreateBulk) SaveX(ctx context.Context) []*Video {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideoCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideoCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
