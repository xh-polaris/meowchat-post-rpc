// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: post.proto

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

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt    int64  `protobuf:"varint,2,opt,name=createAt,proto3" json:"createAt,omitempty"`
	IsAnonymous bool   `protobuf:"varint,3,opt,name=isAnonymous,proto3" json:"isAnonymous,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Text        string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	CoverUrl    string `protobuf:"bytes,6,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`
	Tags        []*Tag `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	UserId      int64  `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status      int64  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	IsDelete    bool   `protobuf:"varint,10,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	DeleteAt    int64  `protobuf:"varint,11,opt,name=deleteAt,proto3" json:"deleteAt,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Post) GetIsAnonymous() bool {
	if x != nil {
		return x.IsAnonymous
	}
	return false
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Post) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Post) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Post) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Post) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Post) GetIsDelete() bool {
	if x != nil {
		return x.IsDelete
	}
	return false
}

func (x *Post) GetDeleteAt() int64 {
	if x != nil {
		return x.DeleteAt
	}
	return 0
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skip 		int64  `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Count       int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"` // 获取的数量
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *ListReq) GetSKip() int64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *ListReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResp.ProtoReflect.Descriptor instead.
func (*ListResp) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{3}
}

func (x *ListResp) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type RetrieveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *RetrieveReq) Reset() {
	*x = RetrieveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveReq) ProtoMessage() {}

func (x *RetrieveReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveReq.ProtoReflect.Descriptor instead.
func (*RetrieveReq) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{4}
}

func (x *RetrieveReq) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type RetrieveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *RetrieveResp) Reset() {
	*x = RetrieveResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveResp) ProtoMessage() {}

func (x *RetrieveResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveResp.ProtoReflect.Descriptor instead.
func (*RetrieveResp) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{5}
}

func (x *RetrieveResp) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type CreatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *CreatReq) Reset() {
	*x = CreatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatReq) ProtoMessage() {}

func (x *CreatReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatReq.ProtoReflect.Descriptor instead.
func (*CreatReq) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{6}
}

func (x *CreatReq) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type CreatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *CreatResp) Reset() {
	*x = CreatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatResp) ProtoMessage() {}

func (x *CreatResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatResp.ProtoReflect.Descriptor instead.
func (*CreatResp) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{7}
}

func (x *CreatResp) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateReq) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{9}
}

type DestroyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *DestroyReq) Reset() {
	*x = DestroyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyReq) ProtoMessage() {}

func (x *DestroyReq) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyReq.ProtoReflect.Descriptor instead.
func (*DestroyReq) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{10}
}

func (x *DestroyReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type DestroyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DestroyResp) Reset() {
	*x = DestroyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyResp) ProtoMessage() {}

func (x *DestroyResp) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyResp.ProtoReflect.Descriptor instead.
func (*DestroyResp) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{11}
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x22, 0x29, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa3, 0x02,
	0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x22, 0x2f, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73,
	0x74, 0x22, 0x23, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x24, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x52, 0x65, 0x73, 0x70, 0x32, 0xea, 0x01, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x52,
	0x50, 0x43, 0x12, 0x25, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x05,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0f, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x12, 0x10,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData = file_post_proto_rawDesc
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_proto_rawDescData)
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_post_proto_goTypes = []interface{}{
	(*Tag)(nil),          // 0: post.Tag
	(*Post)(nil),         // 1: post.Post
	(*ListReq)(nil),      // 2: post.ListReq
	(*ListResp)(nil),     // 3: post.ListResp
	(*RetrieveReq)(nil),  // 4: post.RetrieveReq
	(*RetrieveResp)(nil), // 5: post.RetrieveResp
	(*CreatReq)(nil),     // 6: post.CreatReq
	(*CreatResp)(nil),    // 7: post.CreatResp
	(*UpdateReq)(nil),    // 8: post.UpdateReq
	(*UpdateResp)(nil),   // 9: post.UpdateResp
	(*DestroyReq)(nil),   // 10: post.DestroyReq
	(*DestroyResp)(nil),  // 11: post.DestroyResp
}
var file_post_proto_depIdxs = []int32{
	0,  // 0: post.Post.tags:type_name -> post.Tag
	1,  // 1: post.ListResp.posts:type_name -> post.Post
	1,  // 2: post.RetrieveResp.post:type_name -> post.Post
	1,  // 3: post.CreatReq.post:type_name -> post.Post
	1,  // 4: post.UpdateReq.post:type_name -> post.Post
	2,  // 5: post.postRPC.List:input_type -> post.ListReq
	4,  // 6: post.postRPC.Retrieve:input_type -> post.RetrieveReq
	6,  // 7: post.postRPC.Creat:input_type -> post.CreatReq
	8,  // 8: post.postRPC.Update:input_type -> post.UpdateReq
	10, // 9: post.postRPC.Destroy:input_type -> post.DestroyReq
	3,  // 10: post.postRPC.List:output_type -> post.ListResp
	5,  // 11: post.postRPC.Retrieve:output_type -> post.RetrieveResp
	7,  // 12: post.postRPC.Creat:output_type -> post.CreatResp
	9,  // 13: post.postRPC.Update:output_type -> post.UpdateResp
	11, // 14: post.postRPC.Destroy:output_type -> post.DestroyResp
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResp); i {
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
		file_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveReq); i {
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
		file_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveResp); i {
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
		file_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatReq); i {
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
		file_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatResp); i {
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
		file_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_post_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResp); i {
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
		file_post_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyReq); i {
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
		file_post_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestroyResp); i {
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
			RawDescriptor: file_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_rawDesc = nil
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
