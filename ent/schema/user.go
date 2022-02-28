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
		field.String("name"),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("joined_groups", Group.Type).Ref("members"),
		edge.From("moderating_groups", Group.Type).Ref("moderators"),
		edge.From("streaming_groups", Group.Type).Ref("streamers"),
		edge.From("admin_groups", Group.Type).Ref("admins"),
		edge.To("stream_videos", Video.Type),
		edge.To("liked_videos", Video.Type),
		edge.To("viewed_videos", Video.Type),
		edge.To("moderated_videos", Video.Type),
	}
}
