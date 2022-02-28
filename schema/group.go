package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("url"),
		field.String("thumbnail"),
		field.Int("views"),
		field.Int("subscribers"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admins", User.Type),
		edge.To("members", User.Type),
		edge.To("videos", Video.Type),
		edge.To("moderators", User.Type),
		edge.To("streamers", User.Type),
	}
}
