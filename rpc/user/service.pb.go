// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/user/service.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import track "user-api/rpc/track"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                     string                    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username               string                    `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email                  string                    `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	DisplayName            string                    `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	FullName               string                    `protobuf:"bytes,5,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	FirstName              string                    `protobuf:"bytes,6,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName               string                    `protobuf:"bytes,7,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Member                 bool                      `protobuf:"varint,8,opt,name=member" json:"member,omitempty"`
	Avatar                 []byte                    `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	NewsletterNotification bool                      `protobuf:"varint,10,opt,name=newsletter_notification,json=newsletterNotification" json:"newsletter_notification,omitempty"`
	FavoriteTracks         []string                  `protobuf:"bytes,11,rep,name=favorite_tracks,json=favoriteTracks" json:"favorite_tracks,omitempty"`
	FollowedGroups         []string                  `protobuf:"bytes,12,rep,name=followed_groups,json=followedGroups" json:"followed_groups,omitempty"`
	OwnerOfGroups          []*track.RelatedUserGroup `protobuf:"bytes,13,rep,name=owner_of_groups,json=ownerOfGroups" json:"owner_of_groups,omitempty"`
	ResidenceAddress       *StreetAddress            `protobuf:"bytes,14,opt,name=residence_address,json=residenceAddress" json:"residence_address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                  `json:"-"`
	XXX_unrecognized       []byte                    `json:"-"`
	XXX_sizecache          int32                     `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetMember() bool {
	if m != nil {
		return m.Member
	}
	return false
}

func (m *User) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *User) GetNewsletterNotification() bool {
	if m != nil {
		return m.NewsletterNotification
	}
	return false
}

func (m *User) GetFavoriteTracks() []string {
	if m != nil {
		return m.FavoriteTracks
	}
	return nil
}

func (m *User) GetFollowedGroups() []string {
	if m != nil {
		return m.FollowedGroups
	}
	return nil
}

func (m *User) GetOwnerOfGroups() []*track.RelatedUserGroup {
	if m != nil {
		return m.OwnerOfGroups
	}
	return nil
}

func (m *User) GetResidenceAddress() *StreetAddress {
	if m != nil {
		return m.ResidenceAddress
	}
	return nil
}

type Playlists struct {
	Playlists            []*track.RelatedTrackGroup `protobuf:"bytes,1,rep,name=playlists" json:"playlists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Playlists) Reset()         { *m = Playlists{} }
func (m *Playlists) String() string { return proto.CompactTextString(m) }
func (*Playlists) ProtoMessage()    {}
func (*Playlists) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{1}
}
func (m *Playlists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Playlists.Unmarshal(m, b)
}
func (m *Playlists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Playlists.Marshal(b, m, deterministic)
}
func (dst *Playlists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Playlists.Merge(dst, src)
}
func (m *Playlists) XXX_Size() int {
	return xxx_messageInfo_Playlists.Size(m)
}
func (m *Playlists) XXX_DiscardUnknown() {
	xxx_messageInfo_Playlists.DiscardUnknown(m)
}

var xxx_messageInfo_Playlists proto.InternalMessageInfo

func (m *Playlists) GetPlaylists() []*track.RelatedTrackGroup {
	if m != nil {
		return m.Playlists
	}
	return nil
}

type Tracks struct {
	Tracks               []*track.Track `protobuf:"bytes,1,rep,name=tracks" json:"tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Tracks) Reset()         { *m = Tracks{} }
func (m *Tracks) String() string { return proto.CompactTextString(m) }
func (*Tracks) ProtoMessage()    {}
func (*Tracks) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{2}
}
func (m *Tracks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tracks.Unmarshal(m, b)
}
func (m *Tracks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tracks.Marshal(b, m, deterministic)
}
func (dst *Tracks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tracks.Merge(dst, src)
}
func (m *Tracks) XXX_Size() int {
	return xxx_messageInfo_Tracks.Size(m)
}
func (m *Tracks) XXX_DiscardUnknown() {
	xxx_messageInfo_Tracks.DiscardUnknown(m)
}

var xxx_messageInfo_Tracks proto.InternalMessageInfo

func (m *Tracks) GetTracks() []*track.Track {
	if m != nil {
		return m.Tracks
	}
	return nil
}

type StreetAddress struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PersonalData         bool              `protobuf:"varint,3,opt,name=personal_data,json=personalData" json:"personal_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StreetAddress) Reset()         { *m = StreetAddress{} }
func (m *StreetAddress) String() string { return proto.CompactTextString(m) }
func (*StreetAddress) ProtoMessage()    {}
func (*StreetAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{3}
}
func (m *StreetAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreetAddress.Unmarshal(m, b)
}
func (m *StreetAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreetAddress.Marshal(b, m, deterministic)
}
func (dst *StreetAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreetAddress.Merge(dst, src)
}
func (m *StreetAddress) XXX_Size() int {
	return xxx_messageInfo_StreetAddress.Size(m)
}
func (m *StreetAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_StreetAddress.DiscardUnknown(m)
}

var xxx_messageInfo_StreetAddress proto.InternalMessageInfo

func (m *StreetAddress) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreetAddress) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *StreetAddress) GetPersonalData() bool {
	if m != nil {
		return m.PersonalData
	}
	return false
}

type Artists struct {
	Artists              []*track.RelatedUserGroup `protobuf:"bytes,1,rep,name=artists" json:"artists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Artists) Reset()         { *m = Artists{} }
func (m *Artists) String() string { return proto.CompactTextString(m) }
func (*Artists) ProtoMessage()    {}
func (*Artists) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{4}
}
func (m *Artists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artists.Unmarshal(m, b)
}
func (m *Artists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artists.Marshal(b, m, deterministic)
}
func (dst *Artists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artists.Merge(dst, src)
}
func (m *Artists) XXX_Size() int {
	return xxx_messageInfo_Artists.Size(m)
}
func (m *Artists) XXX_DiscardUnknown() {
	xxx_messageInfo_Artists.DiscardUnknown(m)
}

var xxx_messageInfo_Artists proto.InternalMessageInfo

func (m *Artists) GetArtists() []*track.RelatedUserGroup {
	if m != nil {
		return m.Artists
	}
	return nil
}

// Used for:
// - connecting to user group
// - follow/unfollow group
type UserToUserGroup struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserGroupId          string   `protobuf:"bytes,2,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserToUserGroup) Reset()         { *m = UserToUserGroup{} }
func (m *UserToUserGroup) String() string { return proto.CompactTextString(m) }
func (*UserToUserGroup) ProtoMessage()    {}
func (*UserToUserGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{5}
}
func (m *UserToUserGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToUserGroup.Unmarshal(m, b)
}
func (m *UserToUserGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToUserGroup.Marshal(b, m, deterministic)
}
func (dst *UserToUserGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToUserGroup.Merge(dst, src)
}
func (m *UserToUserGroup) XXX_Size() int {
	return xxx_messageInfo_UserToUserGroup.Size(m)
}
func (m *UserToUserGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToUserGroup.DiscardUnknown(m)
}

var xxx_messageInfo_UserToUserGroup proto.InternalMessageInfo

func (m *UserToUserGroup) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserToUserGroup) GetUserGroupId() string {
	if m != nil {
		return m.UserGroupId
	}
	return ""
}

// Used for:
// - adding/removing favorite tracks
type UserToTrack struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	TrackId              string   `protobuf:"bytes,2,opt,name=track_id,json=trackId" json:"track_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserToTrack) Reset()         { *m = UserToTrack{} }
func (m *UserToTrack) String() string { return proto.CompactTextString(m) }
func (*UserToTrack) ProtoMessage()    {}
func (*UserToTrack) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{6}
}
func (m *UserToTrack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToTrack.Unmarshal(m, b)
}
func (m *UserToTrack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToTrack.Marshal(b, m, deterministic)
}
func (dst *UserToTrack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToTrack.Merge(dst, src)
}
func (m *UserToTrack) XXX_Size() int {
	return xxx_messageInfo_UserToTrack.Size(m)
}
func (m *UserToTrack) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToTrack.DiscardUnknown(m)
}

var xxx_messageInfo_UserToTrack proto.InternalMessageInfo

func (m *UserToTrack) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserToTrack) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

// Used for:
// - adding a track play (i.e. track has been streamed >=45s)
type Play struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	TrackId              string   `protobuf:"bytes,2,opt,name=track_id,json=trackId" json:"track_id,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Credits              float32  `protobuf:"fixed32,4,opt,name=credits" json:"credits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Play) Reset()         { *m = Play{} }
func (m *Play) String() string { return proto.CompactTextString(m) }
func (*Play) ProtoMessage()    {}
func (*Play) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{7}
}
func (m *Play) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Play.Unmarshal(m, b)
}
func (m *Play) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Play.Marshal(b, m, deterministic)
}
func (dst *Play) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Play.Merge(dst, src)
}
func (m *Play) XXX_Size() int {
	return xxx_messageInfo_Play.Size(m)
}
func (m *Play) XXX_DiscardUnknown() {
	xxx_messageInfo_Play.DiscardUnknown(m)
}

var xxx_messageInfo_Play proto.InternalMessageInfo

func (m *Play) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Play) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

func (m *Play) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Play) GetCredits() float32 {
	if m != nil {
		return m.Credits
	}
	return 0
}

type CreatePlayRequest struct {
	Play                 *Play    `protobuf:"bytes,1,opt,name=play" json:"play,omitempty"`
	UpdatedCredits       float32  `protobuf:"fixed32,2,opt,name=updated_credits,json=updatedCredits" json:"updated_credits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayRequest) Reset()         { *m = CreatePlayRequest{} }
func (m *CreatePlayRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePlayRequest) ProtoMessage()    {}
func (*CreatePlayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{8}
}
func (m *CreatePlayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayRequest.Unmarshal(m, b)
}
func (m *CreatePlayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePlayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayRequest.Merge(dst, src)
}
func (m *CreatePlayRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePlayRequest.Size(m)
}
func (m *CreatePlayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayRequest proto.InternalMessageInfo

func (m *CreatePlayRequest) GetPlay() *Play {
	if m != nil {
		return m.Play
	}
	return nil
}

func (m *CreatePlayRequest) GetUpdatedCredits() float32 {
	if m != nil {
		return m.UpdatedCredits
	}
	return 0
}

type CreatePlayResponse struct {
	UpdatedPlayCount     int32    `protobuf:"varint,1,opt,name=updated_play_count,json=updatedPlayCount" json:"updated_play_count,omitempty"`
	UpdatedCredits       float32  `protobuf:"fixed32,2,opt,name=updated_credits,json=updatedCredits" json:"updated_credits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayResponse) Reset()         { *m = CreatePlayResponse{} }
func (m *CreatePlayResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePlayResponse) ProtoMessage()    {}
func (*CreatePlayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_486dc82330f52085, []int{9}
}
func (m *CreatePlayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayResponse.Unmarshal(m, b)
}
func (m *CreatePlayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayResponse.Marshal(b, m, deterministic)
}
func (dst *CreatePlayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayResponse.Merge(dst, src)
}
func (m *CreatePlayResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePlayResponse.Size(m)
}
func (m *CreatePlayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayResponse proto.InternalMessageInfo

func (m *CreatePlayResponse) GetUpdatedPlayCount() int32 {
	if m != nil {
		return m.UpdatedPlayCount
	}
	return 0
}

func (m *CreatePlayResponse) GetUpdatedCredits() float32 {
	if m != nil {
		return m.UpdatedCredits
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "resonate.api.user.User")
	proto.RegisterType((*Playlists)(nil), "resonate.api.user.Playlists")
	proto.RegisterType((*Tracks)(nil), "resonate.api.user.Tracks")
	proto.RegisterType((*StreetAddress)(nil), "resonate.api.user.StreetAddress")
	proto.RegisterMapType((map[string]string)(nil), "resonate.api.user.StreetAddress.DataEntry")
	proto.RegisterType((*Artists)(nil), "resonate.api.user.Artists")
	proto.RegisterType((*UserToUserGroup)(nil), "resonate.api.user.UserToUserGroup")
	proto.RegisterType((*UserToTrack)(nil), "resonate.api.user.UserToTrack")
	proto.RegisterType((*Play)(nil), "resonate.api.user.Play")
	proto.RegisterType((*CreatePlayRequest)(nil), "resonate.api.user.CreatePlayRequest")
	proto.RegisterType((*CreatePlayResponse)(nil), "resonate.api.user.CreatePlayResponse")
}

func init() { proto.RegisterFile("rpc/user/service.proto", fileDescriptor_service_486dc82330f52085) }

var fileDescriptor_service_486dc82330f52085 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x85, 0x2e, 0xd6, 0x65, 0x74, 0xb1, 0xbd, 0x29, 0x6c, 0x46, 0xbd, 0x40, 0x65, 0x5a, 0x54,
	0xe8, 0x45, 0x2e, 0xdc, 0x87, 0x14, 0x29, 0x9a, 0xc2, 0xf1, 0x45, 0x71, 0xd3, 0xd8, 0x05, 0x1d,
	0x3f, 0xb4, 0x2f, 0xc4, 0x46, 0x1c, 0xb5, 0x84, 0x29, 0x92, 0xdd, 0x5d, 0xca, 0xd0, 0x37, 0xf4,
	0x13, 0xfa, 0x2d, 0xfd, 0xb7, 0x60, 0x67, 0x97, 0xb2, 0x0c, 0x8b, 0x4e, 0x20, 0xf8, 0x45, 0xe0,
	0x9c, 0x33, 0xe7, 0xcc, 0x72, 0x76, 0x77, 0x28, 0xd8, 0x11, 0xe9, 0x78, 0x2f, 0x93, 0x28, 0xf6,
	0x24, 0x8a, 0x59, 0x38, 0xc6, 0x61, 0x2a, 0x12, 0x95, 0xb0, 0x6d, 0x81, 0x32, 0x89, 0xb9, 0xc2,
	0x21, 0x4f, 0xc3, 0xa1, 0x4e, 0xe8, 0xf5, 0xf5, 0xef, 0x77, 0x3c, 0x0d, 0xf7, 0xb4, 0x46, 0x09,
	0x3e, 0xbe, 0xba, 0x2d, 0x72, 0xff, 0xab, 0x42, 0xf5, 0x52, 0xa2, 0x60, 0x5d, 0x28, 0x87, 0x81,
	0x53, 0xea, 0x97, 0x06, 0x4d, 0xaf, 0x1c, 0x06, 0xac, 0x07, 0x0d, 0x2d, 0x8e, 0xf9, 0x14, 0x9d,
	0x32, 0xa1, 0x8b, 0x98, 0x7d, 0x04, 0x1b, 0x38, 0xe5, 0x61, 0xe4, 0x54, 0x88, 0x30, 0x01, 0xfb,
	0x1c, 0xda, 0x41, 0x28, 0xd3, 0x88, 0xcf, 0x7d, 0x52, 0x55, 0x89, 0x6c, 0x59, 0xec, 0x4c, 0x0b,
	0x3f, 0x86, 0xe6, 0x24, 0x8b, 0x22, 0xc3, 0x6f, 0x18, 0x57, 0x0d, 0x10, 0xf9, 0x29, 0xc0, 0x24,
	0x14, 0x52, 0x19, 0xb6, 0x46, 0x6c, 0x93, 0x90, 0x5c, 0x1b, 0xf1, 0x9c, 0xad, 0x1b, 0xad, 0x06,
	0x88, 0xdc, 0x81, 0xda, 0x14, 0xa7, 0x6f, 0x51, 0x38, 0x8d, 0x7e, 0x69, 0xd0, 0xf0, 0x6c, 0xa4,
	0x71, 0x3e, 0xe3, 0x8a, 0x0b, 0xa7, 0xd9, 0x2f, 0x0d, 0xda, 0x9e, 0x8d, 0xd8, 0x53, 0xd8, 0x8d,
	0xf1, 0x5a, 0x46, 0xa8, 0x14, 0x0a, 0x3f, 0x4e, 0x54, 0x38, 0x09, 0xc7, 0x5c, 0x85, 0x49, 0xec,
	0x00, 0x19, 0xec, 0xdc, 0xd0, 0x67, 0x4b, 0x2c, 0xfb, 0x0a, 0x36, 0x27, 0x7c, 0x96, 0x88, 0x50,
	0xa1, 0x4f, 0xfd, 0x94, 0x4e, 0xab, 0x5f, 0x19, 0x34, 0xbd, 0x6e, 0x0e, 0xbf, 0x21, 0x94, 0x12,
	0x93, 0x28, 0x4a, 0xae, 0x31, 0xf0, 0xff, 0x12, 0x49, 0x96, 0x4a, 0xa7, 0x6d, 0x13, 0x2d, 0x3c,
	0x22, 0x94, 0xbd, 0x82, 0xcd, 0xe4, 0x3a, 0x46, 0xe1, 0x27, 0x93, 0x3c, 0xb1, 0xd3, 0xaf, 0x0c,
	0x5a, 0xfb, 0x4f, 0x86, 0x77, 0x36, 0x74, 0xe8, 0x61, 0xc4, 0x15, 0x06, 0x7a, 0xc7, 0x48, 0xee,
	0x75, 0x48, 0x7b, 0x3e, 0xb1, 0x66, 0xaf, 0x41, 0x9f, 0x82, 0x30, 0xc0, 0x78, 0x8c, 0x3e, 0x0f,
	0x02, 0x81, 0x52, 0x3a, 0xdd, 0x7e, 0x69, 0xd0, 0xda, 0xef, 0xaf, 0xb0, 0xbb, 0x50, 0x02, 0x51,
	0x1d, 0x98, 0x3c, 0x6f, 0x6b, 0x21, 0xb5, 0x88, 0x7b, 0x0e, 0xcd, 0xdf, 0x23, 0x3e, 0x8f, 0x42,
	0xa9, 0x24, 0x7b, 0x01, 0xcd, 0x34, 0x0f, 0x9c, 0x12, 0x2d, 0xf1, 0x8b, 0xe2, 0x25, 0x52, 0x1b,
	0xcc, 0x1a, 0x6f, 0x64, 0xee, 0x33, 0xa8, 0xd9, 0xfe, 0x7c, 0x0f, 0x35, 0xdb, 0x3f, 0x63, 0xe5,
	0xac, 0xb0, 0xa2, 0x54, 0xcf, 0xe6, 0xb9, 0xff, 0x97, 0xa0, 0x73, 0x6b, 0xc1, 0x77, 0xce, 0xec,
	0x73, 0xa8, 0x06, 0x5c, 0x71, 0xa7, 0x4c, 0x8e, 0x5f, 0xbf, 0xef, 0x85, 0x87, 0x47, 0x5c, 0xf1,
	0xe3, 0x58, 0x89, 0xb9, 0x47, 0x3a, 0xf6, 0x04, 0x3a, 0x29, 0x0a, 0xad, 0x89, 0x7c, 0x32, 0xaa,
	0xd0, 0x59, 0x68, 0xe7, 0xa0, 0xce, 0xef, 0x3d, 0x85, 0xe6, 0x42, 0xc7, 0xb6, 0xa0, 0x72, 0x85,
	0x73, 0xbb, 0x04, 0xfd, 0xa8, 0xef, 0xc6, 0x8c, 0x47, 0x59, 0x7e, 0x69, 0x4c, 0xf0, 0xac, 0xfc,
	0x63, 0xc9, 0x7d, 0x09, 0xf5, 0x03, 0xa1, 0xa8, 0x95, 0x3f, 0x43, 0x9d, 0x9b, 0x47, 0xfb, 0xf6,
	0x1f, 0xb4, 0xd7, 0xb9, 0xc6, 0x3d, 0x83, 0x4d, 0x8d, 0xbe, 0x49, 0x16, 0x1c, 0xdb, 0x85, 0xba,
	0x16, 0xf9, 0x8b, 0x7e, 0xd4, 0x74, 0x78, 0x1a, 0x30, 0x17, 0x3a, 0x44, 0xd0, 0xd1, 0xd2, 0xb4,
	0x59, 0x57, 0x2b, 0xcb, 0xa5, 0xa7, 0x81, 0x7b, 0x00, 0x2d, 0xe3, 0x47, 0x0d, 0x2f, 0xf6, 0x7a,
	0x0c, 0x0d, 0xda, 0x8b, 0x1b, 0x9b, 0x3a, 0xc5, 0xa7, 0x81, 0xfb, 0x37, 0x54, 0xf5, 0x49, 0x59,
	0x47, 0xcb, 0x18, 0x54, 0xd5, 0x3c, 0x45, 0x3b, 0x4d, 0xe8, 0x99, 0x39, 0x50, 0x1f, 0x0b, 0x0c,
	0x42, 0x25, 0x69, 0x8e, 0x94, 0xbd, 0x3c, 0x74, 0x43, 0xd8, 0x3e, 0x14, 0xc8, 0x15, 0xea, 0x7a,
	0x1e, 0xfe, 0x93, 0xa1, 0x54, 0xec, 0x1b, 0xa8, 0xea, 0x43, 0x46, 0x35, 0x5b, 0xfb, 0xbb, 0x2b,
	0xba, 0x49, 0xd9, 0x94, 0xa4, 0xaf, 0x66, 0x96, 0x06, 0xba, 0xb7, 0x7e, 0x5e, 0xa3, 0x4c, 0x35,
	0xba, 0x16, 0x3e, 0xb4, 0xa5, 0xae, 0x80, 0x2d, 0x97, 0x92, 0x69, 0x12, 0x4b, 0x64, 0xdf, 0x02,
	0xcb, 0xe5, 0x34, 0xec, 0xc6, 0x49, 0x16, 0x2b, 0xaa, 0xbc, 0xe1, 0x6d, 0x59, 0x46, 0x0b, 0x0e,
	0x35, 0xfe, 0xc1, 0xc5, 0xf6, 0xff, 0x6d, 0x98, 0x5d, 0xb8, 0x30, 0xf3, 0x99, 0x3d, 0x07, 0x30,
	0xc5, 0x69, 0x3c, 0xaf, 0x7a, 0x25, 0x4d, 0xf4, 0x8a, 0x08, 0xf6, 0x13, 0xd4, 0x47, 0xa8, 0xd6,
	0x14, 0xff, 0x02, 0x70, 0x49, 0xcb, 0xbb, 0x5f, 0xbf, 0xea, 0xd2, 0x1e, 0x4f, 0x53, 0x35, 0xd7,
	0x06, 0x47, 0x18, 0xe1, 0xfa, 0x06, 0xaf, 0xa0, 0x75, 0x42, 0x83, 0xd2, 0x9c, 0x6f, 0xb7, 0xc0,
	0x61, 0xe9, 0x0e, 0xdc, 0x63, 0xf6, 0x1a, 0x3a, 0x97, 0xf1, 0xe4, 0xc1, 0xec, 0x7e, 0x83, 0xad,
	0x83, 0x20, 0x38, 0x59, 0x1e, 0xf8, 0xec, 0xb3, 0x42, 0x47, 0xe2, 0xef, 0x71, 0x3b, 0x87, 0x47,
	0x1e, 0x4e, 0x93, 0x19, 0x3e, 0x94, 0xe1, 0x31, 0xb4, 0x47, 0xa8, 0x6e, 0x06, 0x77, 0x61, 0xf7,
	0x3f, 0x29, 0xb8, 0x27, 0x46, 0x36, 0x82, 0xed, 0x11, 0xaa, 0x93, 0xdb, 0x9f, 0xb5, 0x42, 0xaf,
	0xc7, 0x45, 0xf3, 0x5b, 0xb2, 0x23, 0xe8, 0x8e, 0x50, 0x9d, 0x5f, 0xc7, 0xf6, 0xab, 0xb0, 0x9e,
	0xcb, 0x31, 0x6c, 0x8e, 0x50, 0x51, 0xf0, 0x32, 0x94, 0x2a, 0x11, 0xf3, 0xb5, 0x6c, 0x7e, 0x85,
	0x47, 0x23, 0x54, 0x17, 0x59, 0x9a, 0x26, 0x42, 0x61, 0x90, 0x4f, 0xe4, 0x42, 0xab, 0xde, 0x0a,
	0x22, 0x17, 0xfd, 0x91, 0x5f, 0x51, 0x1a, 0x7d, 0xab, 0x3e, 0x86, 0x77, 0x26, 0x55, 0xef, 0xcb,
	0xf7, 0x64, 0x99, 0x21, 0xf3, 0xa2, 0xf6, 0x67, 0x55, 0x53, 0x6f, 0x6b, 0xf4, 0x37, 0xed, 0x87,
	0x77, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x94, 0x1b, 0x14, 0xf5, 0x09, 0x00, 0x00,
}
