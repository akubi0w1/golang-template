package mysql

import (
	"context"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	entpost "github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/post"
)

type PostImpl struct {
	cli *ent.Client
}

func NewPost(cli *ent.Client) *PostImpl {
	return &PostImpl{
		cli: cli,
	}
}

func (p *PostImpl) Insert(ctx context.Context, post entity.Post) (entity.Post, error) {
	newPost, err := p.cli.Post.Create().
		SetTitle(post.Title).
		SetBody(post.Body).
		SetCreatedAt(post.CreatedAt).
		SetUpdatedAt(post.UpdatedAt).
		AddTagIDs(post.Tags.GetIDs()...).
		AddImageIDs(post.Images.GetIDs()...).
		SetAuthorID(post.AuthorID.Int()).
		Save(ctx)
	if err != nil {
		return entity.Post{}, code.Errorf(code.Database, "failed to insert: %v", err)
	}
	newPost, err = p.cli.Post.Query().
		Where(entpost.IDEQ(newPost.ID)).
		WithAuthor().
		WithTags().
		WithImages().
		Only(ctx)
	if err != nil {
		return entity.Post{}, code.Errorf(code.NotFound, "failed to find new post: %v", err)
	}
	return toEntityPost(newPost), nil
}

func (p *PostImpl) FindAll(ctx context.Context) (entity.PostList, error) {
	posts, err := p.cli.Post.Query().
		Where(entpost.DeletedAtIsNil()).
		WithAuthor().
		WithTags().
		WithImages().
		All(ctx)
	if err != nil {
		return entity.PostList{}, code.Errorf(code.Database, "failed to get all posts: %v", err)
	}
	return toEntityPostList(posts), nil
}

func (p *PostImpl) FindByID(ctx context.Context, id int) (entity.Post, error) {
	post, err := p.cli.Post.Query().
		Where(entpost.DeletedAtIsNil()).
		WithAuthor().
		WithTags().
		WithImages().
		Only(ctx)
	if err != nil {
		return entity.Post{}, code.Errorf(code.NotFound, "failed to find postID=%d: %v", id, err)
	}
	return toEntityPost(post), nil
}

// TODO: add test
func (p *PostImpl) Count(ctx context.Context) (int, error) {
	total, err := p.cli.Post.Query().
		Where(entpost.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return 0, code.Errorf(code.Database, "failed to count all posts: %v", err)
	}
	return total, nil
}

func toEntityPost(p *ent.Post) entity.Post {
	return entity.Post{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Tags:      toEntityTagList(p.Edges.Tags),
		Images:    toEntityImageList(p.Edges.Images),
		AuthorID:  entity.UserID(p.Edges.Author.ID),
	}
}

func toEntityPostList(ps []*ent.Post) entity.PostList {
	list := make(entity.PostList, 0, len(ps))
	for i := range ps {
		list = append(list, toEntityPost(ps[i]))
	}
	return list
}
