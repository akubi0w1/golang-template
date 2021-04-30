// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/migrate"

	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/image"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/post"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/profile"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/tag"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Image is the client for interacting with the Image builders.
	Image *ImageClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// Profile is the client for interacting with the Profile builders.
	Profile *ProfileClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Image = NewImageClient(c.config)
	c.Post = NewPostClient(c.config)
	c.Profile = NewProfileClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Image:   NewImageClient(cfg),
		Post:    NewPostClient(cfg),
		Profile: NewProfileClient(cfg),
		Tag:     NewTagClient(cfg),
		User:    NewUserClient(cfg),
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
		config:  cfg,
		Image:   NewImageClient(cfg),
		Post:    NewPostClient(cfg),
		Profile: NewProfileClient(cfg),
		Tag:     NewTagClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Image.
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
	c.Image.Use(hooks...)
	c.Post.Use(hooks...)
	c.Profile.Use(hooks...)
	c.Tag.Use(hooks...)
	c.User.Use(hooks...)
}

// ImageClient is a client for the Image schema.
type ImageClient struct {
	config
}

// NewImageClient returns a client for the Image from the given config.
func NewImageClient(c config) *ImageClient {
	return &ImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `image.Hooks(f(g(h())))`.
func (c *ImageClient) Use(hooks ...Hook) {
	c.hooks.Image = append(c.hooks.Image, hooks...)
}

// Create returns a create builder for Image.
func (c *ImageClient) Create() *ImageCreate {
	mutation := newImageMutation(c.config, OpCreate)
	return &ImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Image entities.
func (c *ImageClient) CreateBulk(builders ...*ImageCreate) *ImageCreateBulk {
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Image.
func (c *ImageClient) Update() *ImageUpdate {
	mutation := newImageMutation(c.config, OpUpdate)
	return &ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageClient) UpdateOne(i *Image) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImage(i))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageClient) UpdateOneID(id int) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImageID(id))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Image.
func (c *ImageClient) Delete() *ImageDelete {
	mutation := newImageMutation(c.config, OpDelete)
	return &ImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ImageClient) DeleteOne(i *Image) *ImageDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ImageClient) DeleteOneID(id int) *ImageDeleteOne {
	builder := c.Delete().Where(image.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageDeleteOne{builder}
}

// Query returns a query builder for Image.
func (c *ImageClient) Query() *ImageQuery {
	return &ImageQuery{
		config: c.config,
	}
}

// Get returns a Image entity by its id.
func (c *ImageClient) Get(ctx context.Context, id int) (*Image, error) {
	return c.Query().Where(image.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageClient) GetX(ctx context.Context, id int) *Image {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPosts queries the posts edge of a Image.
func (c *ImageClient) QueryPosts(i *Image) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, image.PostsTable, image.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUploadedBy queries the uploadedBy edge of a Image.
func (c *ImageClient) QueryUploadedBy(i *Image) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, image.UploadedByTable, image.UploadedByColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageClient) Hooks() []Hook {
	return c.hooks.Image
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Create returns a create builder for Post.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTags queries the tags edge of a Post.
func (c *PostClient) QueryTags(po *Post) *TagQuery {
	query := &TagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, post.TagsTable, post.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryImages queries the images edge of a Post.
func (c *PostClient) QueryImages(po *Post) *ImageQuery {
	query := &ImageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, post.ImagesTable, post.ImagesColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Post.
func (c *PostClient) QueryAuthor(po *Post) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.AuthorTable, post.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}

// ProfileClient is a client for the Profile schema.
type ProfileClient struct {
	config
}

// NewProfileClient returns a client for the Profile from the given config.
func NewProfileClient(c config) *ProfileClient {
	return &ProfileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profile.Hooks(f(g(h())))`.
func (c *ProfileClient) Use(hooks ...Hook) {
	c.hooks.Profile = append(c.hooks.Profile, hooks...)
}

// Create returns a create builder for Profile.
func (c *ProfileClient) Create() *ProfileCreate {
	mutation := newProfileMutation(c.config, OpCreate)
	return &ProfileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Profile entities.
func (c *ProfileClient) CreateBulk(builders ...*ProfileCreate) *ProfileCreateBulk {
	return &ProfileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Profile.
func (c *ProfileClient) Update() *ProfileUpdate {
	mutation := newProfileMutation(c.config, OpUpdate)
	return &ProfileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfileClient) UpdateOne(pr *Profile) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfile(pr))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfileClient) UpdateOneID(id int) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfileID(id))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profile.
func (c *ProfileClient) Delete() *ProfileDelete {
	mutation := newProfileMutation(c.config, OpDelete)
	return &ProfileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProfileClient) DeleteOne(pr *Profile) *ProfileDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProfileClient) DeleteOneID(id int) *ProfileDeleteOne {
	builder := c.Delete().Where(profile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfileDeleteOne{builder}
}

// Query returns a query builder for Profile.
func (c *ProfileClient) Query() *ProfileQuery {
	return &ProfileQuery{
		config: c.config,
	}
}

// Get returns a Profile entity by its id.
func (c *ProfileClient) Get(ctx context.Context, id int) (*Profile, error) {
	return c.Query().Where(profile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfileClient) GetX(ctx context.Context, id int) *Profile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Profile.
func (c *ProfileClient) QueryUser(pr *Profile) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, profile.UserTable, profile.UserColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProfileClient) Hooks() []Hook {
	return c.hooks.Profile
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Create returns a create builder for Tag.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPosts queries the posts edge of a Tag.
func (c *TagClient) QueryPosts(t *Tag) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.PostsTable, tag.PostsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
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

// QueryProfile queries the profile edge of a User.
func (c *UserClient) QueryProfile(u *User) *ProfileQuery {
	query := &ProfileQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.ProfileTable, user.ProfileColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPosts queries the posts edge of a User.
func (c *UserClient) QueryPosts(u *User) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PostsTable, user.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUpload queries the upload edge of a User.
func (c *UserClient) QueryUpload(u *User) *ImageQuery {
	query := &ImageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.UploadTable, user.UploadColumn),
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