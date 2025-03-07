// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: brevesvault.proto

package brevesvaultrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Language int32

const (
	Language_LANGUAGE_UNKNOWN Language = 0
	Language_ENGLISH          Language = 1
	Language_FRENCH           Language = 2
	Language_KOREAN           Language = 3
	Language_SPANISH          Language = 4
	Language_JAPANESE         Language = 5
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "LANGUAGE_UNKNOWN",
		1: "ENGLISH",
		2: "FRENCH",
		3: "KOREAN",
		4: "SPANISH",
		5: "JAPANESE",
	}
	Language_value = map[string]int32{
		"LANGUAGE_UNKNOWN": 0,
		"ENGLISH":          1,
		"FRENCH":           2,
		"KOREAN":           3,
		"SPANISH":          4,
		"JAPANESE":         5,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_brevesvault_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_brevesvault_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{0}
}

type LanguageLevel int32

const (
	LanguageLevel_LEVEL_UNKNOWN LanguageLevel = 0
	LanguageLevel_A1            LanguageLevel = 1
	LanguageLevel_A2            LanguageLevel = 2
	LanguageLevel_B1            LanguageLevel = 3
	LanguageLevel_B2            LanguageLevel = 4
	LanguageLevel_C1            LanguageLevel = 5
	LanguageLevel_C2            LanguageLevel = 6
)

// Enum value maps for LanguageLevel.
var (
	LanguageLevel_name = map[int32]string{
		0: "LEVEL_UNKNOWN",
		1: "A1",
		2: "A2",
		3: "B1",
		4: "B2",
		5: "C1",
		6: "C2",
	}
	LanguageLevel_value = map[string]int32{
		"LEVEL_UNKNOWN": 0,
		"A1":            1,
		"A2":            2,
		"B1":            3,
		"B2":            4,
		"C1":            5,
		"C2":            6,
	}
)

func (x LanguageLevel) Enum() *LanguageLevel {
	p := new(LanguageLevel)
	*p = x
	return p
}

func (x LanguageLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LanguageLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_brevesvault_proto_enumTypes[1].Descriptor()
}

func (LanguageLevel) Type() protoreflect.EnumType {
	return &file_brevesvault_proto_enumTypes[1]
}

func (x LanguageLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LanguageLevel.Descriptor instead.
func (LanguageLevel) EnumDescriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{1}
}

type UserVideoInList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle      string                 `protobuf:"bytes,2,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Likes         int64                  `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
	Language      Language               `protobuf:"varint,5,opt,name=language,proto3,enum=brevesvaultrpc.Language" json:"language,omitempty"`
	Level         LanguageLevel          `protobuf:"varint,6,opt,name=level,proto3,enum=brevesvaultrpc.LanguageLevel" json:"level,omitempty"`
	Action        string                 `protobuf:"bytes,7,opt,name=action,proto3" json:"action,omitempty"`
	VideoUrl      string                 `protobuf:"bytes,8,opt,name=videoUrl,proto3" json:"videoUrl,omitempty"`
	ActionLevel   string                 `protobuf:"bytes,9,opt,name=actionLevel,proto3" json:"actionLevel,omitempty"`
	VideoId       string                 `protobuf:"bytes,10,opt,name=videoId,proto3" json:"videoId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserVideoInList) Reset() {
	*x = UserVideoInList{}
	mi := &file_brevesvault_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserVideoInList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVideoInList) ProtoMessage() {}

func (x *UserVideoInList) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVideoInList.ProtoReflect.Descriptor instead.
func (*UserVideoInList) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{0}
}

func (x *UserVideoInList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UserVideoInList) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *UserVideoInList) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *UserVideoInList) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_LANGUAGE_UNKNOWN
}

func (x *UserVideoInList) GetLevel() LanguageLevel {
	if x != nil {
		return x.Level
	}
	return LanguageLevel_LEVEL_UNKNOWN
}

func (x *UserVideoInList) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *UserVideoInList) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *UserVideoInList) GetActionLevel() string {
	if x != nil {
		return x.ActionLevel
	}
	return ""
}

func (x *UserVideoInList) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

type UserVideoList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserVideos    []*UserVideoInList     `protobuf:"bytes,1,rep,name=userVideos,proto3" json:"userVideos,omitempty"`
	Total         int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserVideoList) Reset() {
	*x = UserVideoList{}
	mi := &file_brevesvault_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserVideoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVideoList) ProtoMessage() {}

func (x *UserVideoList) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVideoList.ProtoReflect.Descriptor instead.
func (*UserVideoList) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{1}
}

func (x *UserVideoList) GetUserVideos() []*UserVideoInList {
	if x != nil {
		return x.UserVideos
	}
	return nil
}

func (x *UserVideoList) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UserVideoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserVideoRequest) Reset() {
	*x = UserVideoRequest{}
	mi := &file_brevesvault_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVideoRequest) ProtoMessage() {}

func (x *UserVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVideoRequest.ProtoReflect.Descriptor instead.
func (*UserVideoRequest) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{2}
}

func (x *UserVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateVideoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle      string                 `protobuf:"bytes,3,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	Likes         int64                  `protobuf:"varint,4,opt,name=likes,proto3" json:"likes,omitempty"`
	Language      Language               `protobuf:"varint,5,opt,name=language,proto3,enum=brevesvaultrpc.Language" json:"language,omitempty"`
	Level         LanguageLevel          `protobuf:"varint,6,opt,name=level,proto3,enum=brevesvaultrpc.LanguageLevel" json:"level,omitempty"`
	ActionId      int64                  `protobuf:"varint,7,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateVideoRequest) Reset() {
	*x = CreateVideoRequest{}
	mi := &file_brevesvault_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVideoRequest) ProtoMessage() {}

func (x *CreateVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVideoRequest.ProtoReflect.Descriptor instead.
func (*CreateVideoRequest) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVideoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateVideoRequest) GetSubtitle() string {
	if x != nil {
		return x.Subtitle
	}
	return ""
}

func (x *CreateVideoRequest) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *CreateVideoRequest) GetLanguage() Language {
	if x != nil {
		return x.Language
	}
	return Language_LANGUAGE_UNKNOWN
}

func (x *CreateVideoRequest) GetLevel() LanguageLevel {
	if x != nil {
		return x.Level
	}
	return LanguageLevel_LEVEL_UNKNOWN
}

func (x *CreateVideoRequest) GetActionId() int64 {
	if x != nil {
		return x.ActionId
	}
	return 0
}

type VideoLikes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VideoLikes) Reset() {
	*x = VideoLikes{}
	mi := &file_brevesvault_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VideoLikes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoLikes) ProtoMessage() {}

func (x *VideoLikes) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoLikes.ProtoReflect.Descriptor instead.
func (*VideoLikes) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{4}
}

func (x *VideoLikes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LikeVideoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoLikes    []*VideoLikes          `protobuf:"bytes,2,rep,name=videoLikes,proto3" json:"videoLikes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LikeVideoRequest) Reset() {
	*x = LikeVideoRequest{}
	mi := &file_brevesvault_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LikeVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeVideoRequest) ProtoMessage() {}

func (x *LikeVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeVideoRequest.ProtoReflect.Descriptor instead.
func (*LikeVideoRequest) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{5}
}

func (x *LikeVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeVideoRequest) GetVideoLikes() []*VideoLikes {
	if x != nil {
		return x.VideoLikes
	}
	return nil
}

type UpdateUserVideoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoId       int64                  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	IsDifficult   bool                   `protobuf:"varint,3,opt,name=isDifficult,proto3" json:"isDifficult,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserVideoRequest) Reset() {
	*x = UpdateUserVideoRequest{}
	mi := &file_brevesvault_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserVideoRequest) ProtoMessage() {}

func (x *UpdateUserVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserVideoRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserVideoRequest) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserVideoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateUserVideoRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *UpdateUserVideoRequest) GetIsDifficult() bool {
	if x != nil {
		return x.IsDifficult
	}
	return false
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CreatedId     int64                  `protobuf:"varint,1,opt,name=created_id,json=createdId,proto3" json:"created_id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_brevesvault_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_brevesvault_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_brevesvault_proto_rawDescGZIP(), []int{7}
}

func (x *CreateResponse) GetCreatedId() int64 {
	if x != nil {
		return x.CreatedId
	}
	return 0
}

func (x *CreateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_brevesvault_proto protoreflect.FileDescriptor

var file_brevesvault_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x72, 0x70, 0x63, 0x22, 0xb4, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12,
	0x34, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x0d, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xf4, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x34,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x10, 0x4c, 0x69, 0x6b, 0x65, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x6b, 0x65,
	0x73, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x22, 0x6e, 0x0a,
	0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x73, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x60, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e,
	0x47, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x52, 0x45, 0x4e, 0x43,
	0x48, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x4f, 0x52, 0x45, 0x41, 0x4e, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x50, 0x41, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x4a, 0x41, 0x50, 0x41, 0x4e, 0x45, 0x53, 0x45, 0x10, 0x05, 0x2a, 0x52, 0x0a, 0x0d, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x11, 0x0a, 0x0d, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x41, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x41, 0x32, 0x10, 0x02, 0x12, 0x06,
	0x0a, 0x02, 0x42, 0x31, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x42, 0x32, 0x10, 0x04, 0x12, 0x06,
	0x0a, 0x02, 0x43, 0x31, 0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x32, 0x10, 0x06, 0x32, 0xe2,
	0x02, 0x0a, 0x12, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x20, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65,
	0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x22, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x72, 0x65,
	0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x09, 0x4c, 0x69,
	0x6b, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x20, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x72, 0x65, 0x76,
	0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x26, 0x2e, 0x62,
	0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x5a, 0x6f, 0x6c, 0x61, 0x72, 0x61, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x2d, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x72, 0x65, 0x76, 0x65, 0x73, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_brevesvault_proto_rawDescOnce sync.Once
	file_brevesvault_proto_rawDescData []byte
)

func file_brevesvault_proto_rawDescGZIP() []byte {
	file_brevesvault_proto_rawDescOnce.Do(func() {
		file_brevesvault_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_brevesvault_proto_rawDesc), len(file_brevesvault_proto_rawDesc)))
	})
	return file_brevesvault_proto_rawDescData
}

var file_brevesvault_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_brevesvault_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_brevesvault_proto_goTypes = []any{
	(Language)(0),                  // 0: brevesvaultrpc.Language
	(LanguageLevel)(0),             // 1: brevesvaultrpc.LanguageLevel
	(*UserVideoInList)(nil),        // 2: brevesvaultrpc.UserVideoInList
	(*UserVideoList)(nil),          // 3: brevesvaultrpc.UserVideoList
	(*UserVideoRequest)(nil),       // 4: brevesvaultrpc.UserVideoRequest
	(*CreateVideoRequest)(nil),     // 5: brevesvaultrpc.CreateVideoRequest
	(*VideoLikes)(nil),             // 6: brevesvaultrpc.VideoLikes
	(*LikeVideoRequest)(nil),       // 7: brevesvaultrpc.LikeVideoRequest
	(*UpdateUserVideoRequest)(nil), // 8: brevesvaultrpc.UpdateUserVideoRequest
	(*CreateResponse)(nil),         // 9: brevesvaultrpc.CreateResponse
}
var file_brevesvault_proto_depIdxs = []int32{
	0,  // 0: brevesvaultrpc.UserVideoInList.language:type_name -> brevesvaultrpc.Language
	1,  // 1: brevesvaultrpc.UserVideoInList.level:type_name -> brevesvaultrpc.LanguageLevel
	2,  // 2: brevesvaultrpc.UserVideoList.userVideos:type_name -> brevesvaultrpc.UserVideoInList
	0,  // 3: brevesvaultrpc.CreateVideoRequest.language:type_name -> brevesvaultrpc.Language
	1,  // 4: brevesvaultrpc.CreateVideoRequest.level:type_name -> brevesvaultrpc.LanguageLevel
	6,  // 5: brevesvaultrpc.LikeVideoRequest.videoLikes:type_name -> brevesvaultrpc.VideoLikes
	4,  // 6: brevesvaultrpc.brevesVaultService.GetUserVideos:input_type -> brevesvaultrpc.UserVideoRequest
	5,  // 7: brevesvaultrpc.brevesVaultService.CreateVideo:input_type -> brevesvaultrpc.CreateVideoRequest
	7,  // 8: brevesvaultrpc.brevesVaultService.LikeVideo:input_type -> brevesvaultrpc.LikeVideoRequest
	8,  // 9: brevesvaultrpc.brevesVaultService.UpdateUserVideo:input_type -> brevesvaultrpc.UpdateUserVideoRequest
	3,  // 10: brevesvaultrpc.brevesVaultService.GetUserVideos:output_type -> brevesvaultrpc.UserVideoList
	9,  // 11: brevesvaultrpc.brevesVaultService.CreateVideo:output_type -> brevesvaultrpc.CreateResponse
	9,  // 12: brevesvaultrpc.brevesVaultService.LikeVideo:output_type -> brevesvaultrpc.CreateResponse
	3,  // 13: brevesvaultrpc.brevesVaultService.UpdateUserVideo:output_type -> brevesvaultrpc.UserVideoList
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_brevesvault_proto_init() }
func file_brevesvault_proto_init() {
	if File_brevesvault_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_brevesvault_proto_rawDesc), len(file_brevesvault_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brevesvault_proto_goTypes,
		DependencyIndexes: file_brevesvault_proto_depIdxs,
		EnumInfos:         file_brevesvault_proto_enumTypes,
		MessageInfos:      file_brevesvault_proto_msgTypes,
	}.Build()
	File_brevesvault_proto = out.File
	file_brevesvault_proto_goTypes = nil
	file_brevesvault_proto_depIdxs = nil
}
