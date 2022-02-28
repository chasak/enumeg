// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"
	"log"

	"example.com/enumeg/entgen/migrate"

	"example.com/enumeg/entgen/group"
	"example.com/enumeg/entgen/user"
	"example.com/enumeg/entgen/video"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Video is the client for interacting with the Video builders.
	Video *VideoClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Group = NewGroupClient(c.config)
	c.User = NewUserClient(c.config)
	c.Video = NewVideoClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("entgen: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entgen: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Group:  NewGroupClient(cfg),
		User:   NewUserClient(cfg),
		Video:  NewVideoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Group:  NewGroupClient(cfg),
		User:   NewUserClient(cfg),
		Video:  NewVideoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Group.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Group.Use(hooks...)
	c.User.Use(hooks...)
	c.Video.Use(hooks...)
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAdmins queries the admins edge of a Group.
func (c *GroupClient) QueryAdmins(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.AdminsTable, group.AdminsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMembers queries the members edge of a Group.
func (c *GroupClient) QueryMembers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.MembersTable, group.MembersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVideos queries the videos edge of a Group.
func (c *GroupClient) QueryVideos(gr *Group) *VideoQuery {
	query := &VideoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, group.VideosTable, group.VideosColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModerators queries the moderators edge of a Group.
func (c *GroupClient) QueryModerators(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.ModeratorsTable, group.ModeratorsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStreamers queries the streamers edge of a Group.
func (c *GroupClient) QueryStreamers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, group.StreamersTable, group.StreamersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryJoinedGroups queries the joined_groups edge of a User.
func (c *UserClient) QueryJoinedGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.JoinedGroupsTable, user.JoinedGroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModeratingGroups queries the moderating_groups edge of a User.
func (c *UserClient) QueryModeratingGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.ModeratingGroupsTable, user.ModeratingGroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStreamingGroups queries the streaming_groups edge of a User.
func (c *UserClient) QueryStreamingGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.StreamingGroupsTable, user.StreamingGroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAdminGroups queries the admin_groups edge of a User.
func (c *UserClient) QueryAdminGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.AdminGroupsTable, user.AdminGroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStreamVideos queries the stream_videos edge of a User.
func (c *UserClient) QueryStreamVideos(u *User) *VideoQuery {
	query := &VideoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.StreamVideosTable, user.StreamVideosPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikedVideos queries the liked_videos edge of a User.
func (c *UserClient) QueryLikedVideos(u *User) *VideoQuery {
	query := &VideoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.LikedVideosTable, user.LikedVideosPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryViewedVideos queries the viewed_videos edge of a User.
func (c *UserClient) QueryViewedVideos(u *User) *VideoQuery {
	query := &VideoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ViewedVideosTable, user.ViewedVideosPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModeratedVideos queries the moderated_videos edge of a User.
func (c *UserClient) QueryModeratedVideos(u *User) *VideoQuery {
	query := &VideoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ModeratedVideosTable, user.ModeratedVideosPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VideoClient is a client for the Video schema.
type VideoClient struct {
	config
}

// NewVideoClient returns a client for the Video from the given config.
func NewVideoClient(c config) *VideoClient {
	return &VideoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `video.Hooks(f(g(h())))`.
func (c *VideoClient) Use(hooks ...Hook) {
	c.hooks.Video = append(c.hooks.Video, hooks...)
}

// Create returns a create builder for Video.
func (c *VideoClient) Create() *VideoCreate {
	mutation := newVideoMutation(c.config, OpCreate)
	return &VideoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Video entities.
func (c *VideoClient) CreateBulk(builders ...*VideoCreate) *VideoCreateBulk {
	return &VideoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Video.
func (c *VideoClient) Update() *VideoUpdate {
	mutation := newVideoMutation(c.config, OpUpdate)
	return &VideoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VideoClient) UpdateOne(v *Video) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideo(v))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VideoClient) UpdateOneID(id int) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideoID(id))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Video.
func (c *VideoClient) Delete() *VideoDelete {
	mutation := newVideoMutation(c.config, OpDelete)
	return &VideoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VideoClient) DeleteOne(v *Video) *VideoDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VideoClient) DeleteOneID(id int) *VideoDeleteOne {
	builder := c.Delete().Where(video.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VideoDeleteOne{builder}
}

// Query returns a query builder for Video.
func (c *VideoClient) Query() *VideoQuery {
	return &VideoQuery{
		config: c.config,
	}
}

// Get returns a Video entity by its id.
func (c *VideoClient) Get(ctx context.Context, id int) (*Video, error) {
	return c.Query().Where(video.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VideoClient) GetX(ctx context.Context, id int) *Video {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroup queries the group edge of a Video.
func (c *VideoClient) QueryGroup(v *Video) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, video.GroupTable, video.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStreamers queries the streamers edge of a Video.
func (c *VideoClient) QueryStreamers(v *Video) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, video.StreamersTable, video.StreamersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModeraters queries the moderaters edge of a Video.
func (c *VideoClient) QueryModeraters(v *Video) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, video.ModeratersTable, video.ModeratersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLikes queries the likes edge of a Video.
func (c *VideoClient) QueryLikes(v *Video) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, video.LikesTable, video.LikesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryViewers queries the viewers edge of a Video.
func (c *VideoClient) QueryViewers(v *Video) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, video.ViewersTable, video.ViewersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VideoClient) Hooks() []Hook {
	return c.hooks.Video
}