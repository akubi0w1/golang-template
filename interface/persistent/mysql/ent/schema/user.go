package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").Unique(),
		field.String("email").Unique(),
		field.String("password"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Unique(),
		edge.To("posts", Post.Type),
		edge.To("upload", Image.Type),
	}
}
