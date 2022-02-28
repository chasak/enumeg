// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

type CreateGroupReq struct {
	Name       string "json:\"name\""
	Admins     []int  "json:\"admins\""
	Members    []int  "json:\"members\""
	Videos     []int  "json:\"videos\""
	Moderators []int  "json:\"moderators\""
	Streamers  []int  "json:\"streamers\""
}

type CreateUserReq struct {
	Name             string "json:\"name\""
	Password         string "json:\"password\""
	JoinedGroups     []int  "json:\"joined_groups\""
	ModeratingGroups []int  "json:\"moderating_groups\""
	StreamingGroups  []int  "json:\"streaming_groups\""
	AdminGroups      []int  "json:\"admin_groups\""
	StreamVideos     []int  "json:\"stream_videos\""
	LikedVideos      []int  "json:\"liked_videos\""
	ViewedVideos     []int  "json:\"viewed_videos\""
	ModeratedVideos  []int  "json:\"moderated_videos\""
}

type CreateVideoReq struct {
	Title      string                  "json:\"title\""
	Videotype  CreateVideoReqVideotype "json:\"videotype\""
	Group      OptInt                  "json:\"group\""
	Streamers  []int                   "json:\"streamers\""
	Moderaters []int                   "json:\"moderaters\""
	Likes      []int                   "json:\"likes\""
	Viewers    []int                   "json:\"viewers\""
}

type CreateVideoReqVideotype string

const (
	CreateVideoReqVideotypeLive     CreateVideoReqVideotype = "live"
	CreateVideoReqVideotypeVideo    CreateVideoReqVideotype = "video"
	CreateVideoReqVideotypePlaylist CreateVideoReqVideotype = "playlist"
)

// DeleteGroupNoContent is response for DeleteGroup operation.
type DeleteGroupNoContent struct{}

func (*DeleteGroupNoContent) deleteGroupRes() {}

// DeleteUserNoContent is response for DeleteUser operation.
type DeleteUserNoContent struct{}

func (*DeleteUserNoContent) deleteUserRes() {}

// DeleteVideoNoContent is response for DeleteVideo operation.
type DeleteVideoNoContent struct{}

func (*DeleteVideoNoContent) deleteVideoRes() {}

// Ref: #/components/schemas/Group_AdminsList
type GroupAdminsList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/GroupCreate
type GroupCreate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*GroupCreate) createGroupRes() {}

// Ref: #/components/schemas/GroupList
type GroupList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/Group_MembersList
type GroupMembersList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/Group_ModeratorsList
type GroupModeratorsList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/GroupRead
type GroupRead struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*GroupRead) readGroupRes() {}

// Ref: #/components/schemas/Group_StreamersList
type GroupStreamersList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/GroupUpdate
type GroupUpdate struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*GroupUpdate) updateGroupRes() {}

// Ref: #/components/schemas/Group_VideosList
type GroupVideosList struct {
	ID        int                      "json:\"id\""
	Title     string                   "json:\"title\""
	Videotype GroupVideosListVideotype "json:\"videotype\""
}

type GroupVideosListVideotype string

const (
	GroupVideosListVideotypeLive     GroupVideosListVideotype = "live"
	GroupVideosListVideotypeVideo    GroupVideosListVideotype = "video"
	GroupVideosListVideotypePlaylist GroupVideosListVideotype = "playlist"
)

type ListGroupAdminsOKApplicationJSON []GroupAdminsList

func (ListGroupAdminsOKApplicationJSON) listGroupAdminsRes() {}

type ListGroupMembersOKApplicationJSON []GroupMembersList

func (ListGroupMembersOKApplicationJSON) listGroupMembersRes() {}

type ListGroupModeratorsOKApplicationJSON []GroupModeratorsList

func (ListGroupModeratorsOKApplicationJSON) listGroupModeratorsRes() {}

type ListGroupOKApplicationJSON []GroupList

func (ListGroupOKApplicationJSON) listGroupRes() {}

type ListGroupStreamersOKApplicationJSON []GroupStreamersList

func (ListGroupStreamersOKApplicationJSON) listGroupStreamersRes() {}

type ListGroupVideosOKApplicationJSON []GroupVideosList

func (ListGroupVideosOKApplicationJSON) listGroupVideosRes() {}

type ListUserAdminGroupsOKApplicationJSON []UserAdminGroupsList

func (ListUserAdminGroupsOKApplicationJSON) listUserAdminGroupsRes() {}

type ListUserJoinedGroupsOKApplicationJSON []UserJoinedGroupsList

func (ListUserJoinedGroupsOKApplicationJSON) listUserJoinedGroupsRes() {}

type ListUserLikedVideosOKApplicationJSON []UserLikedVideosList

func (ListUserLikedVideosOKApplicationJSON) listUserLikedVideosRes() {}

type ListUserModeratedVideosOKApplicationJSON []UserModeratedVideosList

func (ListUserModeratedVideosOKApplicationJSON) listUserModeratedVideosRes() {}

type ListUserModeratingGroupsOKApplicationJSON []UserModeratingGroupsList

func (ListUserModeratingGroupsOKApplicationJSON) listUserModeratingGroupsRes() {}

type ListUserOKApplicationJSON []UserList

func (ListUserOKApplicationJSON) listUserRes() {}

type ListUserStreamVideosOKApplicationJSON []UserStreamVideosList

func (ListUserStreamVideosOKApplicationJSON) listUserStreamVideosRes() {}

type ListUserStreamingGroupsOKApplicationJSON []UserStreamingGroupsList

func (ListUserStreamingGroupsOKApplicationJSON) listUserStreamingGroupsRes() {}

type ListUserViewedVideosOKApplicationJSON []UserViewedVideosList

func (ListUserViewedVideosOKApplicationJSON) listUserViewedVideosRes() {}

type ListVideoLikesOKApplicationJSON []VideoLikesList

func (ListVideoLikesOKApplicationJSON) listVideoLikesRes() {}

type ListVideoModeratersOKApplicationJSON []VideoModeratersList

func (ListVideoModeratersOKApplicationJSON) listVideoModeratersRes() {}

type ListVideoOKApplicationJSON []VideoList

func (ListVideoOKApplicationJSON) listVideoRes() {}

type ListVideoStreamersOKApplicationJSON []VideoStreamersList

func (ListVideoStreamersOKApplicationJSON) listVideoStreamersRes() {}

type ListVideoViewersOKApplicationJSON []VideoViewersList

func (ListVideoViewersOKApplicationJSON) listVideoViewersRes() {}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUpdateVideoReqVideotype returns new OptUpdateVideoReqVideotype with value set to v.
func NewOptUpdateVideoReqVideotype(v UpdateVideoReqVideotype) OptUpdateVideoReqVideotype {
	return OptUpdateVideoReqVideotype{
		Value: v,
		Set:   true,
	}
}

// OptUpdateVideoReqVideotype is optional UpdateVideoReqVideotype.
type OptUpdateVideoReqVideotype struct {
	Value UpdateVideoReqVideotype
	Set   bool
}

// IsSet returns true if OptUpdateVideoReqVideotype was set.
func (o OptUpdateVideoReqVideotype) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUpdateVideoReqVideotype) Reset() {
	var v UpdateVideoReqVideotype
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUpdateVideoReqVideotype) SetTo(v UpdateVideoReqVideotype) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUpdateVideoReqVideotype) Get() (v UpdateVideoReqVideotype, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUpdateVideoReqVideotype) Or(d UpdateVideoReqVideotype) UpdateVideoReqVideotype {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type R400 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R400) createGroupRes()              {}
func (*R400) createUserRes()               {}
func (*R400) createVideoRes()              {}
func (*R400) deleteGroupRes()              {}
func (*R400) deleteUserRes()               {}
func (*R400) deleteVideoRes()              {}
func (*R400) listGroupAdminsRes()          {}
func (*R400) listGroupMembersRes()         {}
func (*R400) listGroupModeratorsRes()      {}
func (*R400) listGroupRes()                {}
func (*R400) listGroupStreamersRes()       {}
func (*R400) listGroupVideosRes()          {}
func (*R400) listUserAdminGroupsRes()      {}
func (*R400) listUserJoinedGroupsRes()     {}
func (*R400) listUserLikedVideosRes()      {}
func (*R400) listUserModeratedVideosRes()  {}
func (*R400) listUserModeratingGroupsRes() {}
func (*R400) listUserRes()                 {}
func (*R400) listUserStreamVideosRes()     {}
func (*R400) listUserStreamingGroupsRes()  {}
func (*R400) listUserViewedVideosRes()     {}
func (*R400) listVideoLikesRes()           {}
func (*R400) listVideoModeratersRes()      {}
func (*R400) listVideoRes()                {}
func (*R400) listVideoStreamersRes()       {}
func (*R400) listVideoViewersRes()         {}
func (*R400) readGroupRes()                {}
func (*R400) readUserRes()                 {}
func (*R400) readVideoGroupRes()           {}
func (*R400) readVideoRes()                {}
func (*R400) updateGroupRes()              {}
func (*R400) updateUserRes()               {}
func (*R400) updateVideoRes()              {}

type R404 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R404) deleteGroupRes()              {}
func (*R404) deleteUserRes()               {}
func (*R404) deleteVideoRes()              {}
func (*R404) listGroupAdminsRes()          {}
func (*R404) listGroupMembersRes()         {}
func (*R404) listGroupModeratorsRes()      {}
func (*R404) listGroupRes()                {}
func (*R404) listGroupStreamersRes()       {}
func (*R404) listGroupVideosRes()          {}
func (*R404) listUserAdminGroupsRes()      {}
func (*R404) listUserJoinedGroupsRes()     {}
func (*R404) listUserLikedVideosRes()      {}
func (*R404) listUserModeratedVideosRes()  {}
func (*R404) listUserModeratingGroupsRes() {}
func (*R404) listUserRes()                 {}
func (*R404) listUserStreamVideosRes()     {}
func (*R404) listUserStreamingGroupsRes()  {}
func (*R404) listUserViewedVideosRes()     {}
func (*R404) listVideoLikesRes()           {}
func (*R404) listVideoModeratersRes()      {}
func (*R404) listVideoRes()                {}
func (*R404) listVideoStreamersRes()       {}
func (*R404) listVideoViewersRes()         {}
func (*R404) readGroupRes()                {}
func (*R404) readUserRes()                 {}
func (*R404) readVideoGroupRes()           {}
func (*R404) readVideoRes()                {}
func (*R404) updateGroupRes()              {}
func (*R404) updateUserRes()               {}
func (*R404) updateVideoRes()              {}

type R409 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R409) createGroupRes()              {}
func (*R409) createUserRes()               {}
func (*R409) createVideoRes()              {}
func (*R409) deleteGroupRes()              {}
func (*R409) deleteUserRes()               {}
func (*R409) deleteVideoRes()              {}
func (*R409) listGroupAdminsRes()          {}
func (*R409) listGroupMembersRes()         {}
func (*R409) listGroupModeratorsRes()      {}
func (*R409) listGroupRes()                {}
func (*R409) listGroupStreamersRes()       {}
func (*R409) listGroupVideosRes()          {}
func (*R409) listUserAdminGroupsRes()      {}
func (*R409) listUserJoinedGroupsRes()     {}
func (*R409) listUserLikedVideosRes()      {}
func (*R409) listUserModeratedVideosRes()  {}
func (*R409) listUserModeratingGroupsRes() {}
func (*R409) listUserRes()                 {}
func (*R409) listUserStreamVideosRes()     {}
func (*R409) listUserStreamingGroupsRes()  {}
func (*R409) listUserViewedVideosRes()     {}
func (*R409) listVideoLikesRes()           {}
func (*R409) listVideoModeratersRes()      {}
func (*R409) listVideoRes()                {}
func (*R409) listVideoStreamersRes()       {}
func (*R409) listVideoViewersRes()         {}
func (*R409) readGroupRes()                {}
func (*R409) readUserRes()                 {}
func (*R409) readVideoGroupRes()           {}
func (*R409) readVideoRes()                {}
func (*R409) updateGroupRes()              {}
func (*R409) updateUserRes()               {}
func (*R409) updateVideoRes()              {}

type R500 struct {
	Code   int    "json:\"code\""
	Status string "json:\"status\""
	Errors jx.Raw "json:\"errors\""
}

func (*R500) createGroupRes()              {}
func (*R500) createUserRes()               {}
func (*R500) createVideoRes()              {}
func (*R500) deleteGroupRes()              {}
func (*R500) deleteUserRes()               {}
func (*R500) deleteVideoRes()              {}
func (*R500) listGroupAdminsRes()          {}
func (*R500) listGroupMembersRes()         {}
func (*R500) listGroupModeratorsRes()      {}
func (*R500) listGroupRes()                {}
func (*R500) listGroupStreamersRes()       {}
func (*R500) listGroupVideosRes()          {}
func (*R500) listUserAdminGroupsRes()      {}
func (*R500) listUserJoinedGroupsRes()     {}
func (*R500) listUserLikedVideosRes()      {}
func (*R500) listUserModeratedVideosRes()  {}
func (*R500) listUserModeratingGroupsRes() {}
func (*R500) listUserRes()                 {}
func (*R500) listUserStreamVideosRes()     {}
func (*R500) listUserStreamingGroupsRes()  {}
func (*R500) listUserViewedVideosRes()     {}
func (*R500) listVideoLikesRes()           {}
func (*R500) listVideoModeratersRes()      {}
func (*R500) listVideoRes()                {}
func (*R500) listVideoStreamersRes()       {}
func (*R500) listVideoViewersRes()         {}
func (*R500) readGroupRes()                {}
func (*R500) readUserRes()                 {}
func (*R500) readVideoGroupRes()           {}
func (*R500) readVideoRes()                {}
func (*R500) updateGroupRes()              {}
func (*R500) updateUserRes()               {}
func (*R500) updateVideoRes()              {}

type UpdateGroupReq struct {
	Name       OptString "json:\"name\""
	Admins     []int     "json:\"admins\""
	Members    []int     "json:\"members\""
	Videos     []int     "json:\"videos\""
	Moderators []int     "json:\"moderators\""
	Streamers  []int     "json:\"streamers\""
}

type UpdateUserReq struct {
	Name             OptString "json:\"name\""
	Password         OptString "json:\"password\""
	JoinedGroups     []int     "json:\"joined_groups\""
	ModeratingGroups []int     "json:\"moderating_groups\""
	StreamingGroups  []int     "json:\"streaming_groups\""
	AdminGroups      []int     "json:\"admin_groups\""
	StreamVideos     []int     "json:\"stream_videos\""
	LikedVideos      []int     "json:\"liked_videos\""
	ViewedVideos     []int     "json:\"viewed_videos\""
	ModeratedVideos  []int     "json:\"moderated_videos\""
}

type UpdateVideoReq struct {
	Title      OptString                  "json:\"title\""
	Videotype  OptUpdateVideoReqVideotype "json:\"videotype\""
	Group      OptInt                     "json:\"group\""
	Streamers  []int                      "json:\"streamers\""
	Moderaters []int                      "json:\"moderaters\""
	Likes      []int                      "json:\"likes\""
	Viewers    []int                      "json:\"viewers\""
}

type UpdateVideoReqVideotype string

const (
	UpdateVideoReqVideotypeLive     UpdateVideoReqVideotype = "live"
	UpdateVideoReqVideotypeVideo    UpdateVideoReqVideotype = "video"
	UpdateVideoReqVideotypePlaylist UpdateVideoReqVideotype = "playlist"
)

// Ref: #/components/schemas/User_AdminGroupsList
type UserAdminGroupsList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/UserCreate
type UserCreate struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

func (*UserCreate) createUserRes() {}

// Ref: #/components/schemas/User_JoinedGroupsList
type UserJoinedGroupsList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/User_LikedVideosList
type UserLikedVideosList struct {
	ID        int                          "json:\"id\""
	Title     string                       "json:\"title\""
	Videotype UserLikedVideosListVideotype "json:\"videotype\""
}

type UserLikedVideosListVideotype string

const (
	UserLikedVideosListVideotypeLive     UserLikedVideosListVideotype = "live"
	UserLikedVideosListVideotypeVideo    UserLikedVideosListVideotype = "video"
	UserLikedVideosListVideotypePlaylist UserLikedVideosListVideotype = "playlist"
)

// Ref: #/components/schemas/UserList
type UserList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/User_ModeratedVideosList
type UserModeratedVideosList struct {
	ID        int                              "json:\"id\""
	Title     string                           "json:\"title\""
	Videotype UserModeratedVideosListVideotype "json:\"videotype\""
}

type UserModeratedVideosListVideotype string

const (
	UserModeratedVideosListVideotypeLive     UserModeratedVideosListVideotype = "live"
	UserModeratedVideosListVideotypeVideo    UserModeratedVideosListVideotype = "video"
	UserModeratedVideosListVideotypePlaylist UserModeratedVideosListVideotype = "playlist"
)

// Ref: #/components/schemas/User_ModeratingGroupsList
type UserModeratingGroupsList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/UserRead
type UserRead struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

func (*UserRead) readUserRes() {}

// Ref: #/components/schemas/User_StreamVideosList
type UserStreamVideosList struct {
	ID        int                           "json:\"id\""
	Title     string                        "json:\"title\""
	Videotype UserStreamVideosListVideotype "json:\"videotype\""
}

type UserStreamVideosListVideotype string

const (
	UserStreamVideosListVideotypeLive     UserStreamVideosListVideotype = "live"
	UserStreamVideosListVideotypeVideo    UserStreamVideosListVideotype = "video"
	UserStreamVideosListVideotypePlaylist UserStreamVideosListVideotype = "playlist"
)

// Ref: #/components/schemas/User_StreamingGroupsList
type UserStreamingGroupsList struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

// Ref: #/components/schemas/UserUpdate
type UserUpdate struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

func (*UserUpdate) updateUserRes() {}

// Ref: #/components/schemas/User_ViewedVideosList
type UserViewedVideosList struct {
	ID        int                           "json:\"id\""
	Title     string                        "json:\"title\""
	Videotype UserViewedVideosListVideotype "json:\"videotype\""
}

type UserViewedVideosListVideotype string

const (
	UserViewedVideosListVideotypeLive     UserViewedVideosListVideotype = "live"
	UserViewedVideosListVideotypeVideo    UserViewedVideosListVideotype = "video"
	UserViewedVideosListVideotypePlaylist UserViewedVideosListVideotype = "playlist"
)

// Ref: #/components/schemas/VideoCreate
type VideoCreate struct {
	ID        int                  "json:\"id\""
	Title     string               "json:\"title\""
	Videotype VideoCreateVideotype "json:\"videotype\""
}

func (*VideoCreate) createVideoRes() {}

type VideoCreateVideotype string

const (
	VideoCreateVideotypeLive     VideoCreateVideotype = "live"
	VideoCreateVideotypeVideo    VideoCreateVideotype = "video"
	VideoCreateVideotypePlaylist VideoCreateVideotype = "playlist"
)

// Ref: #/components/schemas/Video_GroupRead
type VideoGroupRead struct {
	ID   int    "json:\"id\""
	Name string "json:\"name\""
}

func (*VideoGroupRead) readVideoGroupRes() {}

// Ref: #/components/schemas/Video_LikesList
type VideoLikesList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/VideoList
type VideoList struct {
	ID        int                "json:\"id\""
	Title     string             "json:\"title\""
	Videotype VideoListVideotype "json:\"videotype\""
}

type VideoListVideotype string

const (
	VideoListVideotypeLive     VideoListVideotype = "live"
	VideoListVideotypeVideo    VideoListVideotype = "video"
	VideoListVideotypePlaylist VideoListVideotype = "playlist"
)

// Ref: #/components/schemas/Video_ModeratersList
type VideoModeratersList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/VideoRead
type VideoRead struct {
	ID        int                "json:\"id\""
	Title     string             "json:\"title\""
	Videotype VideoReadVideotype "json:\"videotype\""
}

func (*VideoRead) readVideoRes() {}

type VideoReadVideotype string

const (
	VideoReadVideotypeLive     VideoReadVideotype = "live"
	VideoReadVideotypeVideo    VideoReadVideotype = "video"
	VideoReadVideotypePlaylist VideoReadVideotype = "playlist"
)

// Ref: #/components/schemas/Video_StreamersList
type VideoStreamersList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}

// Ref: #/components/schemas/VideoUpdate
type VideoUpdate struct {
	ID        int                  "json:\"id\""
	Title     string               "json:\"title\""
	Videotype VideoUpdateVideotype "json:\"videotype\""
}

func (*VideoUpdate) updateVideoRes() {}

type VideoUpdateVideotype string

const (
	VideoUpdateVideotypeLive     VideoUpdateVideotype = "live"
	VideoUpdateVideotypeVideo    VideoUpdateVideotype = "video"
	VideoUpdateVideotypePlaylist VideoUpdateVideotype = "playlist"
)

// Ref: #/components/schemas/Video_ViewersList
type VideoViewersList struct {
	ID       int    "json:\"id\""
	Name     string "json:\"name\""
	Password string "json:\"password\""
}
