// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"example.com/enumeg/entgen/group"
	"example.com/enumeg/entgen/predicate"
	"example.com/enumeg/entgen/user"
	"example.com/enumeg/entgen/video"
)

// GroupQuery is the builder for querying Group entities.
type GroupQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Group
	// eager-loading edges.
	withAdmins     *UserQuery
	withMembers    *UserQuery
	withVideos     *VideoQuery
	withModerators *UserQuery
	withStreamers  *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupQuery builder.
func (gq *GroupQuery) Where(ps ...predicate.Group) *GroupQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GroupQuery) Limit(limit int) *GroupQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GroupQuery) Offset(offset int) *GroupQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GroupQuery) Unique(unique bool) *GroupQuery {
	gq.unique = &unique
	return gq
}

// Order adds an order step to the query.
func (gq *GroupQuery) Order(o ...OrderFunc) *GroupQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryAdmins chains the current query on the "admins" edge.
func (gq *GroupQuery) QueryAdmins() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.AdminsTable, group.AdminsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMembers chains the current query on the "members" edge.
func (gq *GroupQuery) QueryMembers() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.MembersTable, group.MembersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVideos chains the current query on the "videos" edge.
func (gq *GroupQuery) QueryVideos() *VideoQuery {
	query := &VideoQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.VideosTable, group.VideosColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryModerators chains the current query on the "moderators" edge.
func (gq *GroupQuery) QueryModerators() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.ModeratorsTable, group.ModeratorsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStreamers chains the current query on the "streamers" edge.
func (gq *GroupQuery) QueryStreamers() *UserQuery {
	query := &UserQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.StreamersTable, group.StreamersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Group entity from the query.
// Returns a *NotFoundError when no Group was found.
func (gq *GroupQuery) First(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{group.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GroupQuery) FirstX(ctx context.Context) *Group {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Group ID from the query.
// Returns a *NotFoundError when no Group ID was found.
func (gq *GroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{group.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GroupQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Group entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Group entity is not found.
// Returns a *NotFoundError when no Group entities are found.
func (gq *GroupQuery) Only(ctx context.Context) (*Group, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{group.Label}
	default:
		return nil, &NotSingularError{group.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GroupQuery) OnlyX(ctx context.Context) *Group {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Group ID in the query.
// Returns a *NotSingularError when exactly one Group ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gq *GroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = &NotSingularError{group.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Groups.
func (gq *GroupQuery) All(ctx context.Context) ([]*Group, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GroupQuery) AllX(ctx context.Context) []*Group {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Group IDs.
func (gq *GroupQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gq.Select(group.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GroupQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GroupQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GroupQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GroupQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GroupQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GroupQuery) Clone() *GroupQuery {
	if gq == nil {
		return nil
	}
	return &GroupQuery{
		config:         gq.config,
		limit:          gq.limit,
		offset:         gq.offset,
		order:          append([]OrderFunc{}, gq.order...),
		predicates:     append([]predicate.Group{}, gq.predicates...),
		withAdmins:     gq.withAdmins.Clone(),
		withMembers:    gq.withMembers.Clone(),
		withVideos:     gq.withVideos.Clone(),
		withModerators: gq.withModerators.Clone(),
		withStreamers:  gq.withStreamers.Clone(),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// WithAdmins tells the query-builder to eager-load the nodes that are connected to
// the "admins" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithAdmins(opts ...func(*UserQuery)) *GroupQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withAdmins = query
	return gq
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithMembers(opts ...func(*UserQuery)) *GroupQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withMembers = query
	return gq
}

// WithVideos tells the query-builder to eager-load the nodes that are connected to
// the "videos" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithVideos(opts ...func(*VideoQuery)) *GroupQuery {
	query := &VideoQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withVideos = query
	return gq
}

// WithModerators tells the query-builder to eager-load the nodes that are connected to
// the "moderators" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithModerators(opts ...func(*UserQuery)) *GroupQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withModerators = query
	return gq
}

// WithStreamers tells the query-builder to eager-load the nodes that are connected to
// the "streamers" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GroupQuery) WithStreamers(opts ...func(*UserQuery)) *GroupQuery {
	query := &UserQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withStreamers = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Group.Query().
//		GroupBy(group.FieldName).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
//
func (gq *GroupQuery) GroupBy(field string, fields ...string) *GroupGroupBy {
	group := &GroupGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Group.Query().
//		Select(group.FieldName).
//		Scan(ctx, &v)
//
func (gq *GroupQuery) Select(fields ...string) *GroupSelect {
	gq.fields = append(gq.fields, fields...)
	return &GroupSelect{GroupQuery: gq}
}

func (gq *GroupQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gq.fields {
		if !group.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GroupQuery) sqlAll(ctx context.Context) ([]*Group, error) {
	var (
		nodes       = []*Group{}
		_spec       = gq.querySpec()
		loadedTypes = [5]bool{
			gq.withAdmins != nil,
			gq.withMembers != nil,
			gq.withVideos != nil,
			gq.withModerators != nil,
			gq.withStreamers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Group{config: gq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("entgen: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gq.withAdmins; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Group, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Admins = []*User{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Group)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   group.AdminsTable,
				Columns: group.AdminsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(group.AdminsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, gq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "admins": %w`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "admins" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Admins = append(nodes[i].Edges.Admins, n)
			}
		}
	}

	if query := gq.withMembers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Group, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Members = []*User{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Group)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   group.MembersTable,
				Columns: group.MembersPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(group.MembersPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, gq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "members": %w`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "members" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Members = append(nodes[i].Edges.Members, n)
			}
		}
	}

	if query := gq.withVideos; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Group)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Videos = []*Video{}
		}
		query.withFKs = true
		query.Where(predicate.Video(func(s *sql.Selector) {
			s.Where(sql.InValues(group.VideosColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.group_videos
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "group_videos" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "group_videos" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Videos = append(node.Edges.Videos, n)
		}
	}

	if query := gq.withModerators; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Group, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Moderators = []*User{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Group)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   group.ModeratorsTable,
				Columns: group.ModeratorsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(group.ModeratorsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, gq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "moderators": %w`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "moderators" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Moderators = append(nodes[i].Edges.Moderators, n)
			}
		}
	}

	if query := gq.withStreamers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Group, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Streamers = []*User{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Group)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   group.StreamersTable,
				Columns: group.StreamersPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(group.StreamersPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, gq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "streamers": %w`, err)
		}
		query.Where(user.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "streamers" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Streamers = append(nodes[i].Edges.Streamers, n)
			}
		}
	}

	return nodes, nil
}

func (gq *GroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	_spec.Node.Columns = gq.fields
	if len(gq.fields) > 0 {
		_spec.Unique = gq.unique != nil && *gq.unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GroupQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("entgen: check existence: %w", err)
	}
	return n > 0, nil
}

func (gq *GroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if unique := gq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for i := range fields {
			if fields[i] != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(group.Table)
	columns := gq.fields
	if len(columns) == 0 {
		columns = group.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.unique != nil && *gq.unique {
		selector.Distinct()
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupGroupBy is the group-by builder for Group entities.
type GroupGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GroupGroupBy) Aggregate(fns ...AggregateFunc) *GroupGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ggb *GroupGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.sql = query
	return ggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GroupGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("entgen: GroupGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GroupGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ggb *GroupGroupBy) StringX(ctx context.Context) string {
	v, err := ggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("entgen: GroupGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GroupGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ggb *GroupGroupBy) IntX(ctx context.Context) int {
	v, err := ggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("entgen: GroupGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GroupGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ggb *GroupGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("entgen: GroupGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GroupGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GroupGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ggb *GroupGroupBy) BoolX(ctx context.Context) bool {
	v, err := ggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GroupGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ggb.fields {
		if !group.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GroupGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql.Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
		for _, f := range ggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ggb.fields...)...)
}

// GroupSelect is the builder for selecting fields of Group entities.
type GroupSelect struct {
	*GroupQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GroupSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.sql = gs.GroupQuery.sqlQuery(ctx)
	return gs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GroupSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("entgen: GroupSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GroupSelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gs *GroupSelect) StringX(ctx context.Context) string {
	v, err := gs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("entgen: GroupSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GroupSelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gs *GroupSelect) IntX(ctx context.Context) int {
	v, err := gs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("entgen: GroupSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GroupSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gs *GroupSelect) Float64X(ctx context.Context) float64 {
	v, err := gs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("entgen: GroupSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GroupSelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gs *GroupSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{group.Label}
	default:
		err = fmt.Errorf("entgen: GroupSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gs *GroupSelect) BoolX(ctx context.Context) bool {
	v, err := gs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GroupSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gs.sql.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
