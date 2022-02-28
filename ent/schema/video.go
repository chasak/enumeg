package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Enum("videotype").Values("live", "video", "playlist"),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("videos").Unique(),
		edge.From("streamers", User.Type).Ref("stream_videos"),
		edge.From("moderaters", User.Type).Ref("moderated_videos"),
		edge.From("likes", User.Type).Ref("liked_videos"),
		edge.From("viewers", User.Type).Ref("viewed_videos"),
	}
}
