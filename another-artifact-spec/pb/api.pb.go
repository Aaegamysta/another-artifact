// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: another-artifact-spec/proto/api.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Nature int32

const (
	Nature_NATURE_UNSPECIFIED Nature = 0
	Nature_NATURE_EVENT       Nature = 1
	Nature_NATURE_BIRTH       Nature = 2
	Nature_NATURE_DEATH       Nature = 3
	Nature_NATURE_HOLIDAY     Nature = 4
)

// Enum value maps for Nature.
var (
	Nature_name = map[int32]string{
		0: "NATURE_UNSPECIFIED",
		1: "NATURE_EVENT",
		2: "NATURE_BIRTH",
		3: "NATURE_DEATH",
		4: "NATURE_HOLIDAY",
	}
	Nature_value = map[string]int32{
		"NATURE_UNSPECIFIED": 0,
		"NATURE_EVENT":       1,
		"NATURE_BIRTH":       2,
		"NATURE_DEATH":       3,
		"NATURE_HOLIDAY":     4,
	}
)

func (x Nature) Enum() *Nature {
	p := new(Nature)
	*p = x
	return p
}

func (x Nature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Nature) Descriptor() protoreflect.EnumDescriptor {
	return file_another_artifact_spec_proto_api_proto_enumTypes[0].Descriptor()
}

func (Nature) Type() protoreflect.EnumType {
	return &file_another_artifact_spec_proto_api_proto_enumTypes[0]
}

func (x Nature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Nature.Descriptor instead.
func (Nature) EnumDescriptor() ([]byte, []int) {
	return file_another_artifact_spec_proto_api_proto_rawDescGZIP(), []int{0}
}

type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link   string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Widith int64  `protobuf:"varint,2,opt,name=widith,proto3" json:"widith,omitempty"`
	Height int64  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_another_artifact_spec_proto_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_another_artifact_spec_proto_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_another_artifact_spec_proto_api_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *Media) GetWidith() int64 {
	if x != nil {
		return x.Widith
	}
	return 0
}

func (x *Media) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Story struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Day     int64  `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Month   int64  `protobuf:"varint,3,opt,name=month,proto3" json:"month,omitempty"`
	Nature  Nature `protobuf:"varint,4,opt,name=nature,proto3,enum=pb.Nature" json:"nature,omitempty"`
	Title   string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Year    int64  `protobuf:"varint,7,opt,name=year,proto3" json:"year,omitempty"`
	Media   *Media `protobuf:"bytes,8,opt,name=media,proto3" json:"media,omitempty"`
}

func (x *Story) Reset() {
	*x = Story{}
	if protoimpl.UnsafeEnabled {
		mi := &file_another_artifact_spec_proto_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Story) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Story) ProtoMessage() {}

func (x *Story) ProtoReflect() protoreflect.Message {
	mi := &file_another_artifact_spec_proto_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Story.ProtoReflect.Descriptor instead.
func (*Story) Descriptor() ([]byte, []int) {
	return file_another_artifact_spec_proto_api_proto_rawDescGZIP(), []int{1}
}

func (x *Story) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Story) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Story) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Story) GetNature() Nature {
	if x != nil {
		return x.Nature
	}
	return Nature_NATURE_UNSPECIFIED
}

func (x *Story) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Story) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Story) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Story) GetMedia() *Media {
	if x != nil {
		return x.Media
	}
	return nil
}

type ListStoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month  int64  `protobuf:"varint,1,opt,name=month,proto3" json:"month,omitempty"`
	Day    int64  `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Nature Nature `protobuf:"varint,3,opt,name=nature,proto3,enum=pb.Nature" json:"nature,omitempty"`
}

func (x *ListStoriesRequest) Reset() {
	*x = ListStoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_another_artifact_spec_proto_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStoriesRequest) ProtoMessage() {}

func (x *ListStoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_another_artifact_spec_proto_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStoriesRequest.ProtoReflect.Descriptor instead.
func (*ListStoriesRequest) Descriptor() ([]byte, []int) {
	return file_another_artifact_spec_proto_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListStoriesRequest) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *ListStoriesRequest) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *ListStoriesRequest) GetNature() Nature {
	if x != nil {
		return x.Nature
	}
	return Nature_NATURE_UNSPECIFIED
}

type ListStoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stories   []*Story `protobuf:"bytes,1,rep,name=stories,proto3" json:"stories,omitempty"`
	FetchedAt int64    `protobuf:"varint,2,opt,name=fetched_at,json=fetchedAt,proto3" json:"fetched_at,omitempty"`
}

func (x *ListStoriesResponse) Reset() {
	*x = ListStoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_another_artifact_spec_proto_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStoriesResponse) ProtoMessage() {}

func (x *ListStoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_another_artifact_spec_proto_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStoriesResponse.ProtoReflect.Descriptor instead.
func (*ListStoriesResponse) Descriptor() ([]byte, []int) {
	return file_another_artifact_spec_proto_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListStoriesResponse) GetStories() []*Story {
	if x != nil {
		return x.Stories
	}
	return nil
}

func (x *ListStoriesResponse) GetFetchedAt() int64 {
	if x != nil {
		return x.FetchedAt
	}
	return 0
}

type SynchronizeStoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Month  int64  `protobuf:"varint,1,opt,name=month,proto3" json:"month,omitempty"`
	Day    int64  `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Nature Nature `protobuf:"varint,3,opt,name=nature,proto3,enum=pb.Nature" json:"nature,omitempty"`
}

func (x *SynchronizeStoriesRequest) Reset() {
	*x = SynchronizeStoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_another_artifact_spec_proto_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynchronizeStoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizeStoriesRequest) ProtoMessage() {}

func (x *SynchronizeStoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_another_artifact_spec_proto_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizeStoriesRequest.ProtoReflect.Descriptor instead.
func (*SynchronizeStoriesRequest) Descriptor() ([]byte, []int) {
	return file_another_artifact_spec_proto_api_proto_rawDescGZIP(), []int{4}
}

func (x *SynchronizeStoriesRequest) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *SynchronizeStoriesRequest) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *SynchronizeStoriesRequest) GetNature() Nature {
	if x != nil {
		return x.Nature
	}
	return Nature_NATURE_UNSPECIFIED
}

type SynchronizeStoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSynchronized bool `protobuf:"varint,1,opt,name=is_synchronized,json=isSynchronized,proto3" json:"is_synchronized,omitempty"`
}

func (x *SynchronizeStoriesResponse) Reset() {
	*x = SynchronizeStoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_another_artifact_spec_proto_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynchronizeStoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynchronizeStoriesResponse) ProtoMessage() {}

func (x *SynchronizeStoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_another_artifact_spec_proto_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynchronizeStoriesResponse.ProtoReflect.Descriptor instead.
func (*SynchronizeStoriesResponse) Descriptor() ([]byte, []int) {
	return file_another_artifact_spec_proto_api_proto_rawDescGZIP(), []int{5}
}

func (x *SynchronizeStoriesResponse) GetIsSynchronized() bool {
	if x != nil {
		return x.IsSynchronized
	}
	return false
}

var File_another_artifact_spec_proto_api_proto protoreflect.FileDescriptor

var file_another_artifact_spec_proto_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2d, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x4b, 0x0a, 0x05, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x64, 0x69,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x69, 0x64, 0x69, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xc8, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x64, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x06, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x1f, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x05, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x22, 0x60, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x12, 0x22, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x06, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x59, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x67, 0x0a, 0x19, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x06, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x45, 0x0a, 0x1a, 0x53, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x69, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x64,
	0x2a, 0x6a, 0x0a, 0x06, 0x4e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x41,
	0x54, 0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x42,
	0x49, 0x52, 0x54, 0x48, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45,
	0x5f, 0x44, 0x45, 0x41, 0x54, 0x48, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x48, 0x4f, 0x4c, 0x49, 0x44, 0x41, 0x59, 0x10, 0x04, 0x32, 0xac, 0x01, 0x0a,
	0x16, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x53, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x4f, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_another_artifact_spec_proto_api_proto_rawDescOnce sync.Once
	file_another_artifact_spec_proto_api_proto_rawDescData = file_another_artifact_spec_proto_api_proto_rawDesc
)

func file_another_artifact_spec_proto_api_proto_rawDescGZIP() []byte {
	file_another_artifact_spec_proto_api_proto_rawDescOnce.Do(func() {
		file_another_artifact_spec_proto_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_another_artifact_spec_proto_api_proto_rawDescData)
	})
	return file_another_artifact_spec_proto_api_proto_rawDescData
}

var file_another_artifact_spec_proto_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_another_artifact_spec_proto_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_another_artifact_spec_proto_api_proto_goTypes = []interface{}{
	(Nature)(0),                        // 0: pb.Nature
	(*Media)(nil),                      // 1: pb.Media
	(*Story)(nil),                      // 2: pb.Story
	(*ListStoriesRequest)(nil),         // 3: pb.ListStoriesRequest
	(*ListStoriesResponse)(nil),        // 4: pb.ListStoriesResponse
	(*SynchronizeStoriesRequest)(nil),  // 5: pb.SynchronizeStoriesRequest
	(*SynchronizeStoriesResponse)(nil), // 6: pb.SynchronizeStoriesResponse
}
var file_another_artifact_spec_proto_api_proto_depIdxs = []int32{
	0, // 0: pb.Story.nature:type_name -> pb.Nature
	1, // 1: pb.Story.media:type_name -> pb.Media
	0, // 2: pb.ListStoriesRequest.nature:type_name -> pb.Nature
	2, // 3: pb.ListStoriesResponse.stories:type_name -> pb.Story
	0, // 4: pb.SynchronizeStoriesRequest.nature:type_name -> pb.Nature
	3, // 5: pb.AnotherArtifactService.ListStories:input_type -> pb.ListStoriesRequest
	5, // 6: pb.AnotherArtifactService.SynchronizeOn:input_type -> pb.SynchronizeStoriesRequest
	4, // 7: pb.AnotherArtifactService.ListStories:output_type -> pb.ListStoriesResponse
	6, // 8: pb.AnotherArtifactService.SynchronizeOn:output_type -> pb.SynchronizeStoriesResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_another_artifact_spec_proto_api_proto_init() }
func file_another_artifact_spec_proto_api_proto_init() {
	if File_another_artifact_spec_proto_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_another_artifact_spec_proto_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_another_artifact_spec_proto_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Story); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_another_artifact_spec_proto_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStoriesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_another_artifact_spec_proto_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStoriesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_another_artifact_spec_proto_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynchronizeStoriesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_another_artifact_spec_proto_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynchronizeStoriesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_another_artifact_spec_proto_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_another_artifact_spec_proto_api_proto_goTypes,
		DependencyIndexes: file_another_artifact_spec_proto_api_proto_depIdxs,
		EnumInfos:         file_another_artifact_spec_proto_api_proto_enumTypes,
		MessageInfos:      file_another_artifact_spec_proto_api_proto_msgTypes,
	}.Build()
	File_another_artifact_spec_proto_api_proto = out.File
	file_another_artifact_spec_proto_api_proto_rawDesc = nil
	file_another_artifact_spec_proto_api_proto_goTypes = nil
	file_another_artifact_spec_proto_api_proto_depIdxs = nil
}
