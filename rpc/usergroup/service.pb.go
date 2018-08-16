// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/usergroup/service.proto

package usergroup

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import track "user-api/rpc/track"
import user "user-api/rpc/user"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserGroup struct {
	Id                   string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DisplayName          string                     `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Description          string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	ShortBio             string                     `protobuf:"bytes,4,opt,name=short_bio,json=shortBio" json:"short_bio,omitempty"`
	Avatar               []byte                     `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Banner               []byte                     `protobuf:"bytes,6,opt,name=banner,proto3" json:"banner,omitempty"`
	OwnerId              string                     `protobuf:"bytes,7,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	Type                 *GroupTaxonomy             `protobuf:"bytes,8,opt,name=type" json:"type,omitempty"`
	Followers            []*user.User               `protobuf:"bytes,9,rep,name=followers" json:"followers,omitempty"`
	Members              []*UserGroup               `protobuf:"bytes,10,rep,name=members" json:"members,omitempty"`
	MemberOfGroups       []*UserGroup               `protobuf:"bytes,11,rep,name=memberOfGroups" json:"memberOfGroups,omitempty"`
	Links                []*Link                    `protobuf:"bytes,12,rep,name=links" json:"links,omitempty"`
	Tags                 []*track.Tag               `protobuf:"bytes,13,rep,name=tags" json:"tags,omitempty"`
	Address              *user.StreetAddress        `protobuf:"bytes,14,opt,name=address" json:"address,omitempty"`
	Privacy              *Privacy                   `protobuf:"bytes,16,opt,name=privacy" json:"privacy,omitempty"`
	RecommendedArtists   []*track.RelatedUserGroup  `protobuf:"bytes,17,rep,name=recommended_artists,json=recommendedArtists" json:"recommended_artists,omitempty"`
	FeaturedTrackGroup   *track.RelatedTrackGroup   `protobuf:"bytes,18,opt,name=featured_track_group,json=featuredTrackGroup" json:"featured_track_group,omitempty"`
	TrackGroups          []*track.RelatedTrackGroup `protobuf:"bytes,19,rep,name=track_groups,json=trackGroups" json:"track_groups,omitempty"`
	HighlightedTracks    []*track.Track             `protobuf:"bytes,20,rep,name=highlighted_tracks,json=highlightedTracks" json:"highlighted_tracks,omitempty"`
	GroupEmailAddress    string                     `protobuf:"bytes,22,opt,name=group_email_address,json=groupEmailAddress" json:"group_email_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UserGroup) Reset()         { *m = UserGroup{} }
func (m *UserGroup) String() string { return proto.CompactTextString(m) }
func (*UserGroup) ProtoMessage()    {}
func (*UserGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{0}
}
func (m *UserGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroup.Unmarshal(m, b)
}
func (m *UserGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroup.Marshal(b, m, deterministic)
}
func (dst *UserGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroup.Merge(dst, src)
}
func (m *UserGroup) XXX_Size() int {
	return xxx_messageInfo_UserGroup.Size(m)
}
func (m *UserGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroup.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroup proto.InternalMessageInfo

func (m *UserGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserGroup) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *UserGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UserGroup) GetShortBio() string {
	if m != nil {
		return m.ShortBio
	}
	return ""
}

func (m *UserGroup) GetAvatar() []byte {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *UserGroup) GetBanner() []byte {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *UserGroup) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *UserGroup) GetType() *GroupTaxonomy {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *UserGroup) GetFollowers() []*user.User {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *UserGroup) GetMembers() []*UserGroup {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *UserGroup) GetMemberOfGroups() []*UserGroup {
	if m != nil {
		return m.MemberOfGroups
	}
	return nil
}

func (m *UserGroup) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *UserGroup) GetTags() []*track.Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *UserGroup) GetAddress() *user.StreetAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserGroup) GetPrivacy() *Privacy {
	if m != nil {
		return m.Privacy
	}
	return nil
}

func (m *UserGroup) GetRecommendedArtists() []*track.RelatedUserGroup {
	if m != nil {
		return m.RecommendedArtists
	}
	return nil
}

func (m *UserGroup) GetFeaturedTrackGroup() *track.RelatedTrackGroup {
	if m != nil {
		return m.FeaturedTrackGroup
	}
	return nil
}

func (m *UserGroup) GetTrackGroups() []*track.RelatedTrackGroup {
	if m != nil {
		return m.TrackGroups
	}
	return nil
}

func (m *UserGroup) GetHighlightedTracks() []*track.Track {
	if m != nil {
		return m.HighlightedTracks
	}
	return nil
}

func (m *UserGroup) GetGroupEmailAddress() string {
	if m != nil {
		return m.GroupEmailAddress
	}
	return ""
}

type TrackAnalytics struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	TotalPlays           int32    `protobuf:"varint,3,opt,name=total_plays,json=totalPlays" json:"total_plays,omitempty"`
	PaidPlays            int32    `protobuf:"varint,4,opt,name=paid_plays,json=paidPlays" json:"paid_plays,omitempty"`
	FreePlays            int32    `protobuf:"varint,5,opt,name=free_plays,json=freePlays" json:"free_plays,omitempty"`
	TotalCredits         float32  `protobuf:"fixed32,6,opt,name=total_credits,json=totalCredits" json:"total_credits,omitempty"`
	UserGroupCredits     float32  `protobuf:"fixed32,7,opt,name=user_group_credits,json=userGroupCredits" json:"user_group_credits,omitempty"`
	ResonateCredits      float32  `protobuf:"fixed32,8,opt,name=resonate_credits,json=resonateCredits" json:"resonate_credits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrackAnalytics) Reset()         { *m = TrackAnalytics{} }
func (m *TrackAnalytics) String() string { return proto.CompactTextString(m) }
func (*TrackAnalytics) ProtoMessage()    {}
func (*TrackAnalytics) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{1}
}
func (m *TrackAnalytics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrackAnalytics.Unmarshal(m, b)
}
func (m *TrackAnalytics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrackAnalytics.Marshal(b, m, deterministic)
}
func (dst *TrackAnalytics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackAnalytics.Merge(dst, src)
}
func (m *TrackAnalytics) XXX_Size() int {
	return xxx_messageInfo_TrackAnalytics.Size(m)
}
func (m *TrackAnalytics) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackAnalytics.DiscardUnknown(m)
}

var xxx_messageInfo_TrackAnalytics proto.InternalMessageInfo

func (m *TrackAnalytics) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TrackAnalytics) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TrackAnalytics) GetTotalPlays() int32 {
	if m != nil {
		return m.TotalPlays
	}
	return 0
}

func (m *TrackAnalytics) GetPaidPlays() int32 {
	if m != nil {
		return m.PaidPlays
	}
	return 0
}

func (m *TrackAnalytics) GetFreePlays() int32 {
	if m != nil {
		return m.FreePlays
	}
	return 0
}

func (m *TrackAnalytics) GetTotalCredits() float32 {
	if m != nil {
		return m.TotalCredits
	}
	return 0
}

func (m *TrackAnalytics) GetUserGroupCredits() float32 {
	if m != nil {
		return m.UserGroupCredits
	}
	return 0
}

func (m *TrackAnalytics) GetResonateCredits() float32 {
	if m != nil {
		return m.ResonateCredits
	}
	return 0
}

type LabelTrackAnalytics struct {
	UserGroup            *track.RelatedUserGroup `protobuf:"bytes,1,opt,name=user_group,json=userGroup" json:"user_group,omitempty"`
	Tracks               []*TrackAnalytics       `protobuf:"bytes,2,rep,name=tracks" json:"tracks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LabelTrackAnalytics) Reset()         { *m = LabelTrackAnalytics{} }
func (m *LabelTrackAnalytics) String() string { return proto.CompactTextString(m) }
func (*LabelTrackAnalytics) ProtoMessage()    {}
func (*LabelTrackAnalytics) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{2}
}
func (m *LabelTrackAnalytics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelTrackAnalytics.Unmarshal(m, b)
}
func (m *LabelTrackAnalytics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelTrackAnalytics.Marshal(b, m, deterministic)
}
func (dst *LabelTrackAnalytics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelTrackAnalytics.Merge(dst, src)
}
func (m *LabelTrackAnalytics) XXX_Size() int {
	return xxx_messageInfo_LabelTrackAnalytics.Size(m)
}
func (m *LabelTrackAnalytics) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelTrackAnalytics.DiscardUnknown(m)
}

var xxx_messageInfo_LabelTrackAnalytics proto.InternalMessageInfo

func (m *LabelTrackAnalytics) GetUserGroup() *track.RelatedUserGroup {
	if m != nil {
		return m.UserGroup
	}
	return nil
}

func (m *LabelTrackAnalytics) GetTracks() []*TrackAnalytics {
	if m != nil {
		return m.Tracks
	}
	return nil
}

type UserGroupTrackAnalytics struct {
	ArtistTrackAnalytics []*TrackAnalytics      `protobuf:"bytes,1,rep,name=artist_track_analytics,json=artistTrackAnalytics" json:"artist_track_analytics,omitempty"`
	LabelTrackAnalytics  []*LabelTrackAnalytics `protobuf:"bytes,2,rep,name=label_track_analytics,json=labelTrackAnalytics" json:"label_track_analytics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UserGroupTrackAnalytics) Reset()         { *m = UserGroupTrackAnalytics{} }
func (m *UserGroupTrackAnalytics) String() string { return proto.CompactTextString(m) }
func (*UserGroupTrackAnalytics) ProtoMessage()    {}
func (*UserGroupTrackAnalytics) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{3}
}
func (m *UserGroupTrackAnalytics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroupTrackAnalytics.Unmarshal(m, b)
}
func (m *UserGroupTrackAnalytics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroupTrackAnalytics.Marshal(b, m, deterministic)
}
func (dst *UserGroupTrackAnalytics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupTrackAnalytics.Merge(dst, src)
}
func (m *UserGroupTrackAnalytics) XXX_Size() int {
	return xxx_messageInfo_UserGroupTrackAnalytics.Size(m)
}
func (m *UserGroupTrackAnalytics) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupTrackAnalytics.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupTrackAnalytics proto.InternalMessageInfo

func (m *UserGroupTrackAnalytics) GetArtistTrackAnalytics() []*TrackAnalytics {
	if m != nil {
		return m.ArtistTrackAnalytics
	}
	return nil
}

func (m *UserGroupTrackAnalytics) GetLabelTrackAnalytics() []*LabelTrackAnalytics {
	if m != nil {
		return m.LabelTrackAnalytics
	}
	return nil
}

type UserGroupRecommended struct {
	UserGroupId          string   `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	RecommendedId        string   `protobuf:"bytes,2,opt,name=recommended_id,json=recommendedId" json:"recommended_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGroupRecommended) Reset()         { *m = UserGroupRecommended{} }
func (m *UserGroupRecommended) String() string { return proto.CompactTextString(m) }
func (*UserGroupRecommended) ProtoMessage()    {}
func (*UserGroupRecommended) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{4}
}
func (m *UserGroupRecommended) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroupRecommended.Unmarshal(m, b)
}
func (m *UserGroupRecommended) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroupRecommended.Marshal(b, m, deterministic)
}
func (dst *UserGroupRecommended) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupRecommended.Merge(dst, src)
}
func (m *UserGroupRecommended) XXX_Size() int {
	return xxx_messageInfo_UserGroupRecommended.Size(m)
}
func (m *UserGroupRecommended) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupRecommended.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupRecommended proto.InternalMessageInfo

func (m *UserGroupRecommended) GetUserGroupId() string {
	if m != nil {
		return m.UserGroupId
	}
	return ""
}

func (m *UserGroupRecommended) GetRecommendedId() string {
	if m != nil {
		return m.RecommendedId
	}
	return ""
}

type UserGroupMembers struct {
	UserGroupId          string       `protobuf:"bytes,1,opt,name=user_group_id,json=userGroupId" json:"user_group_id,omitempty"`
	Members              []*UserGroup `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserGroupMembers) Reset()         { *m = UserGroupMembers{} }
func (m *UserGroupMembers) String() string { return proto.CompactTextString(m) }
func (*UserGroupMembers) ProtoMessage()    {}
func (*UserGroupMembers) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{5}
}
func (m *UserGroupMembers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGroupMembers.Unmarshal(m, b)
}
func (m *UserGroupMembers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGroupMembers.Marshal(b, m, deterministic)
}
func (dst *UserGroupMembers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupMembers.Merge(dst, src)
}
func (m *UserGroupMembers) XXX_Size() int {
	return xxx_messageInfo_UserGroupMembers.Size(m)
}
func (m *UserGroupMembers) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupMembers.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupMembers proto.InternalMessageInfo

func (m *UserGroupMembers) GetUserGroupId() string {
	if m != nil {
		return m.UserGroupId
	}
	return ""
}

func (m *UserGroupMembers) GetMembers() []*UserGroup {
	if m != nil {
		return m.Members
	}
	return nil
}

type GroupTaxonomy struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupTaxonomy) Reset()         { *m = GroupTaxonomy{} }
func (m *GroupTaxonomy) String() string { return proto.CompactTextString(m) }
func (*GroupTaxonomy) ProtoMessage()    {}
func (*GroupTaxonomy) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{6}
}
func (m *GroupTaxonomy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupTaxonomy.Unmarshal(m, b)
}
func (m *GroupTaxonomy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupTaxonomy.Marshal(b, m, deterministic)
}
func (dst *GroupTaxonomy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupTaxonomy.Merge(dst, src)
}
func (m *GroupTaxonomy) XXX_Size() int {
	return xxx_messageInfo_GroupTaxonomy.Size(m)
}
func (m *GroupTaxonomy) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupTaxonomy.DiscardUnknown(m)
}

var xxx_messageInfo_GroupTaxonomy proto.InternalMessageInfo

func (m *GroupTaxonomy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GroupTaxonomy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GroupTaxonomy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GroupTaxonomies struct {
	Types                []*GroupTaxonomy `protobuf:"bytes,1,rep,name=types" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GroupTaxonomies) Reset()         { *m = GroupTaxonomies{} }
func (m *GroupTaxonomies) String() string { return proto.CompactTextString(m) }
func (*GroupTaxonomies) ProtoMessage()    {}
func (*GroupTaxonomies) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{7}
}
func (m *GroupTaxonomies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupTaxonomies.Unmarshal(m, b)
}
func (m *GroupTaxonomies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupTaxonomies.Marshal(b, m, deterministic)
}
func (dst *GroupTaxonomies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupTaxonomies.Merge(dst, src)
}
func (m *GroupTaxonomies) XXX_Size() int {
	return xxx_messageInfo_GroupTaxonomies.Size(m)
}
func (m *GroupTaxonomies) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupTaxonomies.DiscardUnknown(m)
}

var xxx_messageInfo_GroupTaxonomies proto.InternalMessageInfo

func (m *GroupTaxonomies) GetTypes() []*GroupTaxonomy {
	if m != nil {
		return m.Types
	}
	return nil
}

type Link struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Platform             string   `protobuf:"bytes,2,opt,name=platform" json:"platform,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Uri                  string   `protobuf:"bytes,4,opt,name=uri" json:"uri,omitempty"`
	PersonalData         bool     `protobuf:"varint,5,opt,name=personal_data,json=personalData" json:"personal_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{8}
}
func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (dst *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(dst, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Link) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Link) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Link) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Link) GetPersonalData() bool {
	if m != nil {
		return m.PersonalData
	}
	return false
}

type Privacy struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Private              bool     `protobuf:"varint,2,opt,name=private" json:"private,omitempty"`
	OwnedTracks          bool     `protobuf:"varint,3,opt,name=owned_tracks,json=ownedTracks" json:"owned_tracks,omitempty"`
	SupportedArtists     bool     `protobuf:"varint,4,opt,name=supported_artists,json=supportedArtists" json:"supported_artists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Privacy) Reset()         { *m = Privacy{} }
func (m *Privacy) String() string { return proto.CompactTextString(m) }
func (*Privacy) ProtoMessage()    {}
func (*Privacy) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{9}
}
func (m *Privacy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Privacy.Unmarshal(m, b)
}
func (m *Privacy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Privacy.Marshal(b, m, deterministic)
}
func (dst *Privacy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Privacy.Merge(dst, src)
}
func (m *Privacy) XXX_Size() int {
	return xxx_messageInfo_Privacy.Size(m)
}
func (m *Privacy) XXX_DiscardUnknown() {
	xxx_messageInfo_Privacy.DiscardUnknown(m)
}

var xxx_messageInfo_Privacy proto.InternalMessageInfo

func (m *Privacy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Privacy) GetPrivate() bool {
	if m != nil {
		return m.Private
	}
	return false
}

func (m *Privacy) GetOwnedTracks() bool {
	if m != nil {
		return m.OwnedTracks
	}
	return false
}

func (m *Privacy) GetSupportedArtists() bool {
	if m != nil {
		return m.SupportedArtists
	}
	return false
}

type GroupedUserGroups struct {
	Labels               []*UserGroup `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
	Artists              []*UserGroup `protobuf:"bytes,2,rep,name=artists" json:"artists,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GroupedUserGroups) Reset()         { *m = GroupedUserGroups{} }
func (m *GroupedUserGroups) String() string { return proto.CompactTextString(m) }
func (*GroupedUserGroups) ProtoMessage()    {}
func (*GroupedUserGroups) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_3f1927c4248e5009, []int{10}
}
func (m *GroupedUserGroups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupedUserGroups.Unmarshal(m, b)
}
func (m *GroupedUserGroups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupedUserGroups.Marshal(b, m, deterministic)
}
func (dst *GroupedUserGroups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupedUserGroups.Merge(dst, src)
}
func (m *GroupedUserGroups) XXX_Size() int {
	return xxx_messageInfo_GroupedUserGroups.Size(m)
}
func (m *GroupedUserGroups) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupedUserGroups.DiscardUnknown(m)
}

var xxx_messageInfo_GroupedUserGroups proto.InternalMessageInfo

func (m *GroupedUserGroups) GetLabels() []*UserGroup {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *GroupedUserGroups) GetArtists() []*UserGroup {
	if m != nil {
		return m.Artists
	}
	return nil
}

func init() {
	proto.RegisterType((*UserGroup)(nil), "resonate.api.user.UserGroup")
	proto.RegisterType((*TrackAnalytics)(nil), "resonate.api.user.TrackAnalytics")
	proto.RegisterType((*LabelTrackAnalytics)(nil), "resonate.api.user.LabelTrackAnalytics")
	proto.RegisterType((*UserGroupTrackAnalytics)(nil), "resonate.api.user.UserGroupTrackAnalytics")
	proto.RegisterType((*UserGroupRecommended)(nil), "resonate.api.user.UserGroupRecommended")
	proto.RegisterType((*UserGroupMembers)(nil), "resonate.api.user.UserGroupMembers")
	proto.RegisterType((*GroupTaxonomy)(nil), "resonate.api.user.GroupTaxonomy")
	proto.RegisterType((*GroupTaxonomies)(nil), "resonate.api.user.GroupTaxonomies")
	proto.RegisterType((*Link)(nil), "resonate.api.user.Link")
	proto.RegisterType((*Privacy)(nil), "resonate.api.user.Privacy")
	proto.RegisterType((*GroupedUserGroups)(nil), "resonate.api.user.GroupedUserGroups")
}

func init() {
	proto.RegisterFile("rpc/usergroup/service.proto", fileDescriptor_service_3f1927c4248e5009)
}

var fileDescriptor_service_3f1927c4248e5009 = []byte{
	// 1140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xeb, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0x1d, 0x3b, 0xb6, 0x8f, 0x2f, 0xb1, 0x27, 0x21, 0x5d, 0x5c, 0x50, 0xdd, 0x4d, 0x81,
	0x50, 0xa8, 0x23, 0x85, 0x82, 0x04, 0xff, 0x92, 0xa6, 0xb2, 0x02, 0x6d, 0x29, 0x1b, 0xb7, 0x48,
	0x95, 0x90, 0x35, 0xf1, 0x4c, 0x92, 0x51, 0xf6, 0xa6, 0x99, 0x71, 0x8a, 0x7f, 0xf0, 0x03, 0xf1,
	0x08, 0xc0, 0x6b, 0xf1, 0x2c, 0x3c, 0x42, 0x35, 0x97, 0x5d, 0xaf, 0xaf, 0x75, 0xa5, 0xfc, 0xb1,
	0x76, 0xce, 0xf9, 0xce, 0x37, 0xc7, 0xe7, 0xba, 0x0b, 0x77, 0x79, 0x3c, 0x3c, 0x18, 0x09, 0xca,
	0x2f, 0x79, 0x34, 0x8a, 0x0f, 0x04, 0xe5, 0x37, 0x6c, 0x48, 0xbb, 0x31, 0x8f, 0x64, 0x84, 0x5a,
	0x9c, 0x8a, 0x28, 0xc4, 0x92, 0x76, 0x71, 0xcc, 0xba, 0x0a, 0xd5, 0xbe, 0xa7, 0x7e, 0x1f, 0xe1,
	0x98, 0x1d, 0x24, 0x86, 0xd3, 0x36, 0xed, 0xce, 0x14, 0x40, 0x72, 0x3c, 0xbc, 0x9e, 0x46, 0xb8,
	0xff, 0x97, 0xa0, 0xf2, 0x4a, 0x50, 0xde, 0x53, 0x37, 0xa2, 0x06, 0xe4, 0x19, 0x71, 0x72, 0x9d,
	0xdc, 0x7e, 0xc5, 0xcb, 0x33, 0x82, 0xee, 0x43, 0x8d, 0x30, 0x11, 0xfb, 0x78, 0x3c, 0x08, 0x71,
	0x40, 0x9d, 0xbc, 0xd6, 0x54, 0xad, 0xec, 0x05, 0x0e, 0x28, 0xea, 0x40, 0x95, 0x50, 0x31, 0xe4,
	0x2c, 0x96, 0x2c, 0x0a, 0x9d, 0x0d, 0x8b, 0x98, 0x88, 0xd0, 0x5d, 0xa8, 0x88, 0xab, 0x88, 0xcb,
	0xc1, 0x39, 0x8b, 0x9c, 0x82, 0xd6, 0x97, 0xb5, 0xe0, 0x98, 0x45, 0x68, 0x17, 0x36, 0xf1, 0x0d,
	0x96, 0x98, 0x3b, 0xc5, 0x4e, 0x6e, 0xbf, 0xe6, 0xd9, 0x93, 0x92, 0x9f, 0xe3, 0x30, 0xa4, 0xdc,
	0xd9, 0x34, 0x72, 0x73, 0x42, 0x1f, 0x43, 0x39, 0x7a, 0x1b, 0x52, 0x3e, 0x60, 0xc4, 0x29, 0x69,
	0xae, 0x92, 0x3e, 0x9f, 0x12, 0xf4, 0x18, 0x0a, 0x72, 0x1c, 0x53, 0xa7, 0xdc, 0xc9, 0xed, 0x57,
	0x0f, 0x3b, 0xdd, 0xb9, 0x78, 0x75, 0xf5, 0x9f, 0xec, 0xe3, 0xdf, 0xa3, 0x30, 0x0a, 0xc6, 0x9e,
	0x46, 0xa3, 0x6f, 0xa1, 0x72, 0x11, 0xf9, 0x7e, 0xf4, 0x96, 0x72, 0xe1, 0x54, 0x3a, 0x1b, 0xfb,
	0xd5, 0xc3, 0x3b, 0x0b, 0x4c, 0x55, 0x8c, 0xbc, 0x09, 0x12, 0x7d, 0x07, 0xa5, 0x80, 0x06, 0xe7,
	0xca, 0x08, 0xb4, 0xd1, 0x27, 0x4b, 0x8c, 0xf4, 0x9d, 0x5e, 0x02, 0x46, 0x27, 0xd0, 0x30, 0x8f,
	0x3f, 0x5f, 0x68, 0x8d, 0x70, 0xaa, 0x6b, 0x98, 0xcf, 0xd8, 0xa0, 0x47, 0x50, 0xf4, 0x59, 0x78,
	0x2d, 0x9c, 0xda, 0x52, 0x87, 0x9f, 0xb1, 0xf0, 0xda, 0x33, 0x28, 0xf4, 0x10, 0x0a, 0x12, 0x5f,
	0x0a, 0xa7, 0xae, 0xd1, 0xbb, 0x0b, 0xd0, 0x7d, 0x7c, 0xe9, 0x69, 0x0c, 0xfa, 0x01, 0x4a, 0x98,
	0x10, 0x4e, 0x85, 0x70, 0x1a, 0x4b, 0x03, 0x79, 0x26, 0x39, 0xa5, 0xf2, 0xc8, 0xe0, 0xbc, 0xc4,
	0x00, 0x3d, 0x86, 0x52, 0xcc, 0xd9, 0x0d, 0x1e, 0x8e, 0x9d, 0xa6, 0xb6, 0x6d, 0x2f, 0xb0, 0x7d,
	0x69, 0x10, 0x5e, 0x02, 0x45, 0x7d, 0xd8, 0xe6, 0x74, 0x18, 0x05, 0x01, 0x0d, 0x09, 0x25, 0x03,
	0xcc, 0x25, 0x13, 0x52, 0x38, 0x2d, 0xed, 0xec, 0xde, 0x02, 0x06, 0x8f, 0xfa, 0x58, 0x52, 0x32,
	0x09, 0x0f, 0xca, 0xd8, 0x1f, 0x19, 0x73, 0xf4, 0x1a, 0x76, 0x2e, 0x28, 0x96, 0x23, 0x4e, 0xc9,
	0x40, 0x17, 0xfe, 0x40, 0x37, 0x95, 0x83, 0xb4, 0x63, 0x0f, 0x96, 0xd3, 0xf6, 0x15, 0xd8, 0xf2,
	0x26, 0x0c, 0x13, 0x19, 0xea, 0x41, 0x2d, 0x43, 0x27, 0x9c, 0x6d, 0xed, 0xe6, 0x7a, 0x7c, 0x55,
	0x99, 0x3e, 0x0b, 0xd4, 0x03, 0x74, 0xc5, 0x2e, 0xaf, 0x7c, 0x76, 0x79, 0x25, 0x13, 0x1f, 0x85,
	0xb3, 0xa3, 0xe9, 0x9c, 0x45, 0x29, 0x52, 0x00, 0xaf, 0x95, 0xb1, 0xd1, 0x12, 0x81, 0xba, 0xb0,
	0xad, 0x7d, 0x19, 0xd0, 0x00, 0x33, 0x7f, 0x90, 0x64, 0x6f, 0x57, 0x77, 0x47, 0x4b, 0xab, 0x9e,
	0x2a, 0x8d, 0x4d, 0x97, 0xfb, 0x77, 0x1e, 0x1a, 0xda, 0xf4, 0x28, 0xc4, 0xfe, 0x58, 0xb2, 0xa1,
	0x98, 0xeb, 0xfb, 0x1d, 0x28, 0x4a, 0x26, 0xfd, 0xa4, 0xe1, 0xcd, 0x01, 0xdd, 0x83, 0xaa, 0x8c,
	0x24, 0xf6, 0x07, 0xaa, 0xf9, 0x85, 0x6e, 0xf5, 0xa2, 0x07, 0x5a, 0xf4, 0x52, 0x49, 0xd0, 0xa7,
	0x00, 0x31, 0x66, 0xc4, 0xea, 0x0b, 0x5a, 0x5f, 0x51, 0x92, 0x54, 0x7d, 0xc1, 0x29, 0xb5, 0xea,
	0xa2, 0x51, 0x2b, 0x89, 0x51, 0xef, 0x41, 0xdd, 0xd0, 0x0f, 0x39, 0x25, 0x4c, 0x0a, 0xdd, 0xf9,
	0x79, 0xaf, 0xa6, 0x85, 0x4f, 0x8c, 0x0c, 0x7d, 0x0d, 0x48, 0x45, 0xc3, 0x44, 0x3f, 0x45, 0x96,
	0x34, 0xb2, 0x39, 0x4a, 0x2a, 0x22, 0x41, 0x7f, 0x09, 0xcd, 0x24, 0x90, 0x29, 0xb6, 0xac, 0xb1,
	0x5b, 0x89, 0xdc, 0x42, 0xdd, 0x7f, 0x72, 0xb0, 0xfd, 0x0c, 0x9f, 0x53, 0x7f, 0x26, 0x34, 0xc7,
	0x00, 0x93, 0x0b, 0x75, 0x88, 0xd6, 0x2c, 0xca, 0x4a, 0xea, 0x0d, 0xfa, 0x1e, 0x36, 0x6d, 0x7a,
	0xf3, 0x3a, 0xbd, 0xf7, 0x97, 0xa5, 0x37, 0xbd, 0xd6, 0xb3, 0x06, 0xee, 0x7f, 0x39, 0xb8, 0x93,
	0x72, 0xce, 0xb8, 0xf6, 0x2b, 0xec, 0x9a, 0x66, 0xb1, 0x05, 0x8e, 0x13, 0x8d, 0x93, 0x5b, 0xf7,
	0x9a, 0x1d, 0x43, 0x30, 0x43, 0xfc, 0x06, 0x3e, 0xf2, 0x55, 0x28, 0xe6, 0x78, 0x8d, 0xfb, 0x9f,
	0x2f, 0x1a, 0x37, 0xf3, 0xa1, 0xf3, 0xb6, 0xfd, 0x79, 0xa1, 0x8b, 0x61, 0x67, 0x12, 0xa3, 0x49,
	0xdb, 0x22, 0x17, 0xea, 0x99, 0xc4, 0xa6, 0xd5, 0x58, 0x4d, 0xa3, 0x78, 0x4a, 0xd0, 0x67, 0xd0,
	0xc8, 0x4e, 0x0a, 0x46, 0x6c, 0x7d, 0xd6, 0x33, 0xd2, 0x53, 0xe2, 0x86, 0xd0, 0x4c, 0xaf, 0x78,
	0x6e, 0xe7, 0xee, 0x3a, 0xf4, 0x99, 0x99, 0x9e, 0xff, 0x80, 0x99, 0xee, 0xf6, 0xa0, 0x3e, 0xb5,
	0x59, 0xe6, 0xda, 0x09, 0xd9, 0xcd, 0x64, 0xbc, 0x35, 0x7b, 0x07, 0x41, 0x41, 0xaf, 0x54, 0xb3,
	0x30, 0xf5, 0xb3, 0x7b, 0x0a, 0x5b, 0x59, 0x22, 0x46, 0xd5, 0x9e, 0x29, 0x2a, 0x78, 0x92, 0xd2,
	0xf7, 0x6f, 0x35, 0x03, 0x77, 0xff, 0x80, 0x82, 0xda, 0x00, 0x73, 0xae, 0xb4, 0xa1, 0x1c, 0xfb,
	0x58, 0x5e, 0x44, 0x3c, 0xb0, 0xee, 0xa4, 0xe7, 0xd4, 0xcd, 0x8d, 0x8c, 0x9b, 0x4d, 0xd8, 0x18,
	0x71, 0x66, 0xd7, 0xb6, 0x7a, 0x54, 0x6d, 0x1a, 0x53, 0xae, 0x9c, 0xf0, 0x07, 0x04, 0x4b, 0xac,
	0x1b, 0xb9, 0xec, 0xd5, 0x12, 0xe1, 0x09, 0x96, 0xd8, 0xfd, 0x2b, 0x07, 0x25, 0x3b, 0xe8, 0xe7,
	0x5c, 0x70, 0xec, 0x96, 0x90, 0x26, 0x20, 0x65, 0x2f, 0x39, 0xaa, 0xd7, 0x0d, 0xb5, 0xcc, 0xd3,
	0x61, 0xb8, 0xa1, 0xd5, 0x55, 0x2d, 0xb3, 0xc3, 0xee, 0x2b, 0x68, 0x89, 0x51, 0x1c, 0x47, 0x5c,
	0x66, 0x56, 0x45, 0x41, 0xe3, 0x9a, 0xa9, 0xc2, 0xee, 0x00, 0xf7, 0xcf, 0x1c, 0xb4, 0x74, 0x74,
	0x32, 0x7d, 0xa9, 0xb6, 0xd4, 0xa6, 0x2e, 0xcc, 0x24, 0xa6, 0xab, 0xb3, 0x6c, 0xb1, 0xaa, 0x38,
	0x92, 0xeb, 0xd6, 0x2a, 0x0e, 0x0b, 0x3e, 0xfc, 0xb7, 0x94, 0xa9, 0xc6, 0x33, 0xf3, 0xee, 0x85,
	0x9e, 0xc3, 0xd6, 0x13, 0x4e, 0xb1, 0xa4, 0x93, 0x57, 0xaf, 0x95, 0x74, 0xed, 0x95, 0x5a, 0xf4,
	0x23, 0xd4, 0x7a, 0x54, 0xde, 0x0e, 0xd7, 0x29, 0x6c, 0xbd, 0x8a, 0xc9, 0x07, 0xb8, 0xb6, 0x68,
	0x57, 0x3d, 0x0d, 0x62, 0x39, 0x56, 0x54, 0x27, 0xd4, 0xa7, 0xb7, 0x41, 0xd5, 0x07, 0xd4, 0xa3,
	0x52, 0x0f, 0x99, 0x4c, 0x26, 0x97, 0xe2, 0xdb, 0x0f, 0x96, 0xf5, 0xc9, 0x54, 0x25, 0xfc, 0x02,
	0xad, 0x6c, 0xdc, 0xfa, 0xaa, 0x73, 0x56, 0x90, 0xba, 0xef, 0x69, 0x3e, 0xd5, 0xaf, 0x67, 0xd0,
	0x38, 0x22, 0x24, 0x3b, 0xd8, 0xbe, 0x58, 0x59, 0x27, 0x13, 0xe0, 0x8a, 0x7f, 0xff, 0x1a, 0x5a,
	0x1e, 0x0d, 0xa2, 0x1b, 0x7a, 0xcb, 0xbc, 0x3f, 0x01, 0x1c, 0x11, 0x92, 0x8c, 0xc8, 0xbd, 0x55,
	0x84, 0x16, 0xb4, 0x82, 0xec, 0x05, 0xd4, 0x4d, 0xb6, 0x6f, 0x89, 0xef, 0x37, 0x9d, 0x9c, 0x99,
	0xcd, 0xb4, 0xba, 0x7e, 0x1e, 0xae, 0xd2, 0x4e, 0x33, 0x1d, 0x57, 0xdf, 0x54, 0xd2, 0x2f, 0xad,
	0xf3, 0x4d, 0xfd, 0x31, 0xf4, 0xcd, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x61, 0x86, 0x69,
	0x81, 0x0d, 0x00, 0x00,
}
