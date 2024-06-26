// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: post.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug         string                 `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Content      string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	MainImage    string                 `protobuf:"bytes,5,opt,name=main_image,json=mainImage,proto3" json:"main_image,omitempty"`
	CaptionImage string                 `protobuf:"bytes,6,opt,name=caption_image,json=captionImage,proto3" json:"caption_image,omitempty"`
	Type         string                 `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	IsFeatured   uint32                 `protobuf:"varint,8,opt,name=is_featured,json=isFeatured,proto3" json:"is_featured,omitempty"`
	Visitors     uint64                 `protobuf:"varint,9,opt,name=visitors,proto3" json:"visitors,omitempty"`
	CreatedBy    string                 `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy    string                 `protobuf:"bytes,11,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Tags         string                 `protobuf:"bytes,15,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetMainImage() string {
	if x != nil {
		return x.MainImage
	}
	return ""
}

func (x *Post) GetCaptionImage() string {
	if x != nil {
		return x.CaptionImage
	}
	return ""
}

func (x *Post) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Post) GetIsFeatured() uint32 {
	if x != nil {
		return x.IsFeatured
	}
	return 0
}

func (x *Post) GetVisitors() uint64 {
	if x != nil {
		return x.Visitors
	}
	return 0
}

func (x *Post) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Post) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Post) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Post) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type GetAllPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Post `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllPostsResponse) Reset() {
	*x = GetAllPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPostsResponse) ProtoMessage() {}

func (x *GetAllPostsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllPostsResponse.ProtoReflect.Descriptor instead.
func (*GetAllPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllPostsResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllPostsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllPostsResponse) GetData() []*Post {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPostByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostByIdRequest) Reset() {
	*x = GetPostByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIdRequest) ProtoMessage() {}

func (x *GetPostByIdRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPostByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPostByIdRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostByIdRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Post  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetPostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetPostResponse) GetData() *Post {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	ImageFilename string `protobuf:"bytes,4,opt,name=image_filename,json=imageFilename,proto3" json:"image_filename,omitempty"`
	ImageBuffer   []byte `protobuf:"bytes,5,opt,name=image_buffer,json=imageBuffer,proto3" json:"image_buffer,omitempty"`
	ImageCaption  string `protobuf:"bytes,6,opt,name=image_caption,json=imageCaption,proto3" json:"image_caption,omitempty"`
	Type          string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	IsFeatured    uint32 `protobuf:"varint,8,opt,name=is_featured,json=isFeatured,proto3" json:"is_featured,omitempty"`
	CreatedBy     string `protobuf:"bytes,9,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy     string `protobuf:"bytes,10,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Tags          string `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{4}
}

func (x *CreatePostRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreatePostRequest) GetImageFilename() string {
	if x != nil {
		return x.ImageFilename
	}
	return ""
}

func (x *CreatePostRequest) GetImageBuffer() []byte {
	if x != nil {
		return x.ImageBuffer
	}
	return nil
}

func (x *CreatePostRequest) GetImageCaption() string {
	if x != nil {
		return x.ImageCaption
	}
	return ""
}

func (x *CreatePostRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreatePostRequest) GetIsFeatured() uint32 {
	if x != nil {
		return x.IsFeatured
	}
	return 0
}

func (x *CreatePostRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *CreatePostRequest) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *CreatePostRequest) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type DeletePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeletePostResponse) Reset() {
	*x = DeletePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResponse) ProtoMessage() {}

func (x *DeletePostResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeletePostResponse.ProtoReflect.Descriptor instead.
func (*DeletePostResponse) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePostResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeletePostResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x03,
	0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x76, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0x70, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6c, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc9, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa7, 0x04, 0x0a, 0x0b, 0x50, 0x6f, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x26, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x58, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x56, 0x69,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
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

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_post_proto_goTypes = []interface{}{
	(*Post)(nil),                  // 0: tracer_study_grpc.Post
	(*GetAllPostsResponse)(nil),   // 1: tracer_study_grpc.GetAllPostsResponse
	(*GetPostByIdRequest)(nil),    // 2: tracer_study_grpc.GetPostByIdRequest
	(*GetPostResponse)(nil),       // 3: tracer_study_grpc.GetPostResponse
	(*CreatePostRequest)(nil),     // 4: tracer_study_grpc.CreatePostRequest
	(*DeletePostResponse)(nil),    // 5: tracer_study_grpc.DeletePostResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_post_proto_depIdxs = []int32{
	6,  // 0: tracer_study_grpc.Post.created_at:type_name -> google.protobuf.Timestamp
	6,  // 1: tracer_study_grpc.Post.updated_at:type_name -> google.protobuf.Timestamp
	6,  // 2: tracer_study_grpc.Post.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 3: tracer_study_grpc.GetAllPostsResponse.data:type_name -> tracer_study_grpc.Post
	0,  // 4: tracer_study_grpc.GetPostResponse.data:type_name -> tracer_study_grpc.Post
	7,  // 5: tracer_study_grpc.PostService.GetAllPosts:input_type -> google.protobuf.Empty
	2,  // 6: tracer_study_grpc.PostService.GetPostById:input_type -> tracer_study_grpc.GetPostByIdRequest
	4,  // 7: tracer_study_grpc.PostService.CreatePost:input_type -> tracer_study_grpc.CreatePostRequest
	4,  // 8: tracer_study_grpc.PostService.UpdatePost:input_type -> tracer_study_grpc.CreatePostRequest
	2,  // 9: tracer_study_grpc.PostService.DeletePost:input_type -> tracer_study_grpc.GetPostByIdRequest
	2,  // 10: tracer_study_grpc.PostService.AddVisitor:input_type -> tracer_study_grpc.GetPostByIdRequest
	1,  // 11: tracer_study_grpc.PostService.GetAllPosts:output_type -> tracer_study_grpc.GetAllPostsResponse
	3,  // 12: tracer_study_grpc.PostService.GetPostById:output_type -> tracer_study_grpc.GetPostResponse
	3,  // 13: tracer_study_grpc.PostService.CreatePost:output_type -> tracer_study_grpc.GetPostResponse
	3,  // 14: tracer_study_grpc.PostService.UpdatePost:output_type -> tracer_study_grpc.GetPostResponse
	5,  // 15: tracer_study_grpc.PostService.DeletePost:output_type -> tracer_study_grpc.DeletePostResponse
	3,  // 16: tracer_study_grpc.PostService.AddVisitor:output_type -> tracer_study_grpc.GetPostResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
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
		file_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPostsResponse); i {
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
			switch v := v.(*GetPostByIdRequest); i {
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
			switch v := v.(*GetPostResponse); i {
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
			switch v := v.(*CreatePostRequest); i {
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
			switch v := v.(*DeletePostResponse); i {
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
			NumMessages:   6,
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
