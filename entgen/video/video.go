// Code generated by entc, DO NOT EDIT.

package video

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the video type in the database.
	Label = "video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldVideotype holds the string denoting the videotype field in the database.
	FieldVideotype = "videotype"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeStreamers holds the string denoting the streamers edge name in mutations.
	EdgeStreamers = "streamers"
	// EdgeModeraters holds the string denoting the moderaters edge name in mutations.
	EdgeModeraters = "moderaters"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgeViewers holds the string denoting the viewers edge name in mutations.
	EdgeViewers = "viewers"
	// Table holds the table name of the video in the database.
	Table = "videos"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "videos"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_videos"
	// StreamersTable is the table that holds the streamers relation/edge. The primary key declared below.
	StreamersTable = "user_stream_videos"
	// StreamersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	StreamersInverseTable = "users"
	// ModeratersTable is the table that holds the moderaters relation/edge. The primary key declared below.
	ModeratersTable = "user_moderated_videos"
	// ModeratersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ModeratersInverseTable = "users"
	// LikesTable is the table that holds the likes relation/edge. The primary key declared below.
	LikesTable = "user_liked_videos"
	// LikesInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	LikesInverseTable = "users"
	// ViewersTable is the table that holds the viewers relation/edge. The primary key declared below.
	ViewersTable = "user_viewed_videos"
	// ViewersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ViewersInverseTable = "users"
)

// Columns holds all SQL columns for video fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldTitle,
	FieldDescription,
	FieldVideotype,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "videos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_videos",
}

var (
	// StreamersPrimaryKey and StreamersColumn2 are the table columns denoting the
	// primary key for the streamers relation (M2M).
	StreamersPrimaryKey = []string{"user_id", "video_id"}
	// ModeratersPrimaryKey and ModeratersColumn2 are the table columns denoting the
	// primary key for the moderaters relation (M2M).
	ModeratersPrimaryKey = []string{"user_id", "video_id"}
	// LikesPrimaryKey and LikesColumn2 are the table columns denoting the
	// primary key for the likes relation (M2M).
	LikesPrimaryKey = []string{"user_id", "video_id"}
	// ViewersPrimaryKey and ViewersColumn2 are the table columns denoting the
	// primary key for the viewers relation (M2M).
	ViewersPrimaryKey = []string{"user_id", "video_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Videotype defines the type for the "videotype" enum field.
type Videotype string

// Videotype values.
const (
	VideotypeLive     Videotype = "live"
	VideotypeVideo    Videotype = "video"
	VideotypePlaylist Videotype = "playlist"
)

func (v Videotype) String() string {
	return string(v)
}

// VideotypeValidator is a validator for the "videotype" field enum values. It is called by the builders before save.
func VideotypeValidator(v Videotype) error {
	switch v {
	case VideotypeLive, VideotypeVideo, VideotypePlaylist:
		return nil
	default:
		return fmt.Errorf("video: invalid enum value for videotype field: %q", v)
	}
}
