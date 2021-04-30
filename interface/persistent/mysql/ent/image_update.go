// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/image"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/post"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/predicate"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/user"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where adds a new predicate for the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetURL sets the "url" field.
func (iu *ImageUpdate) SetURL(s string) *ImageUpdate {
	iu.mutation.SetURL(s)
	return iu
}

// SetCreatedAt sets the "created_at" field.
func (iu *ImageUpdate) SetCreatedAt(t time.Time) *ImageUpdate {
	iu.mutation.SetCreatedAt(t)
	return iu
}

// SetPostsID sets the "posts" edge to the Post entity by ID.
func (iu *ImageUpdate) SetPostsID(id int) *ImageUpdate {
	iu.mutation.SetPostsID(id)
	return iu
}

// SetNillablePostsID sets the "posts" edge to the Post entity by ID if the given value is not nil.
func (iu *ImageUpdate) SetNillablePostsID(id *int) *ImageUpdate {
	if id != nil {
		iu = iu.SetPostsID(*id)
	}
	return iu
}

// SetPosts sets the "posts" edge to the Post entity.
func (iu *ImageUpdate) SetPosts(p *Post) *ImageUpdate {
	return iu.SetPostsID(p.ID)
}

// SetUploadedByID sets the "uploadedBy" edge to the User entity by ID.
func (iu *ImageUpdate) SetUploadedByID(id int) *ImageUpdate {
	iu.mutation.SetUploadedByID(id)
	return iu
}

// SetNillableUploadedByID sets the "uploadedBy" edge to the User entity by ID if the given value is not nil.
func (iu *ImageUpdate) SetNillableUploadedByID(id *int) *ImageUpdate {
	if id != nil {
		iu = iu.SetUploadedByID(*id)
	}
	return iu
}

// SetUploadedBy sets the "uploadedBy" edge to the User entity.
func (iu *ImageUpdate) SetUploadedBy(u *User) *ImageUpdate {
	return iu.SetUploadedByID(u.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// ClearPosts clears the "posts" edge to the Post entity.
func (iu *ImageUpdate) ClearPosts() *ImageUpdate {
	iu.mutation.ClearPosts()
	return iu
}

// ClearUploadedBy clears the "uploadedBy" edge to the User entity.
func (iu *ImageUpdate) ClearUploadedBy() *ImageUpdate {
	iu.mutation.ClearUploadedBy()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: image.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: image.FieldURL,
		})
	}
	if value, ok := iu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: image.FieldCreatedAt,
		})
	}
	if iu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.PostsTable,
			Columns: []string{image.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.PostsTable,
			Columns: []string{image.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.UploadedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.UploadedByTable,
			Columns: []string{image.UploadedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.UploadedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.UploadedByTable,
			Columns: []string{image.UploadedByColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetURL sets the "url" field.
func (iuo *ImageUpdateOne) SetURL(s string) *ImageUpdateOne {
	iuo.mutation.SetURL(s)
	return iuo
}

// SetCreatedAt sets the "created_at" field.
func (iuo *ImageUpdateOne) SetCreatedAt(t time.Time) *ImageUpdateOne {
	iuo.mutation.SetCreatedAt(t)
	return iuo
}

// SetPostsID sets the "posts" edge to the Post entity by ID.
func (iuo *ImageUpdateOne) SetPostsID(id int) *ImageUpdateOne {
	iuo.mutation.SetPostsID(id)
	return iuo
}

// SetNillablePostsID sets the "posts" edge to the Post entity by ID if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillablePostsID(id *int) *ImageUpdateOne {
	if id != nil {
		iuo = iuo.SetPostsID(*id)
	}
	return iuo
}

// SetPosts sets the "posts" edge to the Post entity.
func (iuo *ImageUpdateOne) SetPosts(p *Post) *ImageUpdateOne {
	return iuo.SetPostsID(p.ID)
}

// SetUploadedByID sets the "uploadedBy" edge to the User entity by ID.
func (iuo *ImageUpdateOne) SetUploadedByID(id int) *ImageUpdateOne {
	iuo.mutation.SetUploadedByID(id)
	return iuo
}

// SetNillableUploadedByID sets the "uploadedBy" edge to the User entity by ID if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableUploadedByID(id *int) *ImageUpdateOne {
	if id != nil {
		iuo = iuo.SetUploadedByID(*id)
	}
	return iuo
}

// SetUploadedBy sets the "uploadedBy" edge to the User entity.
func (iuo *ImageUpdateOne) SetUploadedBy(u *User) *ImageUpdateOne {
	return iuo.SetUploadedByID(u.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// ClearPosts clears the "posts" edge to the Post entity.
func (iuo *ImageUpdateOne) ClearPosts() *ImageUpdateOne {
	iuo.mutation.ClearPosts()
	return iuo
}

// ClearUploadedBy clears the "uploadedBy" edge to the User entity.
func (iuo *ImageUpdateOne) ClearUploadedBy() *ImageUpdateOne {
	iuo.mutation.ClearUploadedBy()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	var (
		err  error
		node *Image
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: image.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Image.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: image.FieldURL,
		})
	}
	if value, ok := iuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: image.FieldCreatedAt,
		})
	}
	if iuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.PostsTable,
			Columns: []string{image.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.PostsTable,
			Columns: []string{image.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.UploadedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.UploadedByTable,
			Columns: []string{image.UploadedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.UploadedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.UploadedByTable,
			Columns: []string{image.UploadedByColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}