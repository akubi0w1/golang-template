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

// TODO: add test
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

// TODO: add test
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

// TODO: add test
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

func toEntityTag(t *ent.Tag) entity.Tag {
	return entity.Tag{
		ID:        t.ID,
		Tag:       t.Tag,
		CreatedAt: t.CreatedAt,
	}
}

func toEntityTagList(tags []*ent.Tag) entity.TagList {
	list := make(entity.TagList, 0, len(tags))
	for i := range tags {
		list = append(list, toEntityTag(tags[i]))
	}
	return list
}

func toEntityImage(image *ent.Image) entity.Image {
	return entity.Image{
		ID:        image.ID,
		URL:       image.URL,
		CreatedBy: entity.UserID(image.Edges.UploadedBy.ID),
	}
}

func toEntityImageList(images []*ent.Image) entity.ImageList {
	list := make(entity.ImageList, 0, len(images))
	for i := range images {
		list = append(list, toEntityImage(images[i]))
	}
	return list
}
