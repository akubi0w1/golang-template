package mysql

import (
	"context"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/predicate"
	enttag "github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/tag"
)

type TagImpl struct {
	cli *ent.Client
}

func NewTag(cli *ent.Client) *TagImpl {
	return &TagImpl{
		cli: cli,
	}
}

func (t *TagImpl) FindByIDs(ctx context.Context, ids []int) (entity.TagList, error) {
	if len(ids) == 0 {
		return entity.TagList{}, nil
	}
	conds := make([]predicate.Tag, 0, len(ids))
	for i := range ids {
		conds = append(conds, enttag.IDEQ(ids[i]))
	}
	tags, err := t.cli.Tag.Query().Where(enttag.Or(conds...)).All(ctx)
	if err != nil {
		return entity.TagList{}, code.Errorf(code.Database, "failed to find tags by ids: %v", err)
	}
	return toEntityTagList(tags), nil
}

func (t *TagImpl) FindByTags(ctx context.Context, tags []string) (entity.TagList, error) {
	if len(tags) == 0 {
		return entity.TagList{}, nil
	}
	conds := make([]predicate.Tag, 0, len(tags))
	for i := range tags {
		conds = append(conds, enttag.TagEQ(tags[i]))
	}
	_tags, err := t.cli.Tag.Query().Where(enttag.Or(conds...)).All(ctx)
	if err != nil {
		return entity.TagList{}, code.Errorf(code.Database, "failed to find tags by tags: %v", err)
	}
	return toEntityTagList(_tags), nil
}

func (t *TagImpl) InsertBulk(ctx context.Context, tags entity.TagList) (entity.TagList, error) {
	if len(tags) == 0 {
		return entity.TagList{}, nil
	}
	queries := make([]*ent.TagCreate, 0, len(tags))
	for i := range tags {
		queries = append(queries, t.cli.Tag.Create().
			SetTag(tags[i].Tag).
			SetCreatedAt(tags[i].CreatedAt),
		)
	}
	_tags, err := t.cli.Tag.CreateBulk(queries...).Save(ctx)
	if err != nil {
		return entity.TagList{}, code.Errorf(code.Database, "failed to insert tags: %v", err)
	}
	return toEntityTagList(_tags), nil
}

func (t *TagImpl) IsExist(ctx context.Context, tag string) (bool, error) {
	isExist, err := t.cli.Tag.Query().Where(enttag.TagEQ(tag)).Exist(ctx)
	if err != nil {
		return false, code.Errorf(code.Database, "failed to check tag=%s is exist: %v", tag, err)
	}
	return isExist, nil
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
