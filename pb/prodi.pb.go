// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: prodi.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Prodi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kode          string                 `protobuf:"bytes,1,opt,name=kode,proto3" json:"kode,omitempty"`
	KodeDikti     string                 `protobuf:"bytes,2,opt,name=kode_dikti,json=kodeDikti,proto3" json:"kode_dikti,omitempty"`
	KodeFak       string                 `protobuf:"bytes,3,opt,name=kode_fak,json=kodeFak,proto3" json:"kode_fak,omitempty"`
	KodeIntegrasi string                 `protobuf:"bytes,4,opt,name=kode_integrasi,json=kodeIntegrasi,proto3" json:"kode_integrasi,omitempty"`
	Nama          string                 `protobuf:"bytes,5,opt,name=nama,proto3" json:"nama,omitempty"`
	Jenjang       string                 `protobuf:"bytes,6,opt,name=jenjang,proto3" json:"jenjang,omitempty"`
	NamaFak       string                 `protobuf:"bytes,7,opt,name=nama_fak,json=namaFak,proto3" json:"nama_fak,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Prodi) Reset() {
	*x = Prodi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prodi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prodi) ProtoMessage() {}

func (x *Prodi) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prodi.ProtoReflect.Descriptor instead.
func (*Prodi) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{0}
}

func (x *Prodi) GetKode() string {
	if x != nil {
		return x.Kode
	}
	return ""
}

func (x *Prodi) GetKodeDikti() string {
	if x != nil {
		return x.KodeDikti
	}
	return ""
}

func (x *Prodi) GetKodeFak() string {
	if x != nil {
		return x.KodeFak
	}
	return ""
}

func (x *Prodi) GetKodeIntegrasi() string {
	if x != nil {
		return x.KodeIntegrasi
	}
	return ""
}

func (x *Prodi) GetNama() string {
	if x != nil {
		return x.Nama
	}
	return ""
}

func (x *Prodi) GetJenjang() string {
	if x != nil {
		return x.Jenjang
	}
	return ""
}

func (x *Prodi) GetNamaFak() string {
	if x != nil {
		return x.NamaFak
	}
	return ""
}

func (x *Prodi) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Prodi) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Prodi) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type EmptyProdiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyProdiRequest) Reset() {
	*x = EmptyProdiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyProdiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyProdiRequest) ProtoMessage() {}

func (x *EmptyProdiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyProdiRequest.ProtoReflect.Descriptor instead.
func (*EmptyProdiRequest) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{1}
}

type GetAllProdiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*Prodi `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAllProdiResponse) Reset() {
	*x = GetAllProdiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prodi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProdiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProdiResponse) ProtoMessage() {}

func (x *GetAllProdiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prodi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProdiResponse.ProtoReflect.Descriptor instead.
func (*GetAllProdiResponse) Descriptor() ([]byte, []int) {
	return file_prodi_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllProdiResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetAllProdiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllProdiResponse) GetData() []*Prodi {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_prodi_proto protoreflect.FileDescriptor

var file_prodi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6b, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x69, 0x6b, 0x74, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x6b, 0x74, 0x69, 0x12, 0x19,
	0x0a, 0x08, 0x6b, 0x6f, 0x64, 0x65, 0x5f, 0x66, 0x61, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6b, 0x6f, 0x64, 0x65, 0x46, 0x61, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x6b, 0x6f, 0x64,
	0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x73, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6b, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x73, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x65, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x65, 0x6e, 0x6a, 0x61, 0x6e, 0x67, 0x12, 0x19,
	0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x61, 0x5f, 0x66, 0x61, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x61, 0x6d, 0x61, 0x46, 0x61, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x71, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x32, 0x6d, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64,
	0x69, 0x12, 0x24, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_prodi_proto_rawDescOnce sync.Once
	file_prodi_proto_rawDescData = file_prodi_proto_rawDesc
)

func file_prodi_proto_rawDescGZIP() []byte {
	file_prodi_proto_rawDescOnce.Do(func() {
		file_prodi_proto_rawDescData = protoimpl.X.CompressGZIP(file_prodi_proto_rawDescData)
	})
	return file_prodi_proto_rawDescData
}

var file_prodi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prodi_proto_goTypes = []interface{}{
	(*Prodi)(nil),                 // 0: tracer_study_grpc.Prodi
	(*EmptyProdiRequest)(nil),     // 1: tracer_study_grpc.EmptyProdiRequest
	(*GetAllProdiResponse)(nil),   // 2: tracer_study_grpc.GetAllProdiResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_prodi_proto_depIdxs = []int32{
	3, // 0: tracer_study_grpc.Prodi.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: tracer_study_grpc.Prodi.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: tracer_study_grpc.Prodi.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 3: tracer_study_grpc.GetAllProdiResponse.data:type_name -> tracer_study_grpc.Prodi
	1, // 4: tracer_study_grpc.ProdiService.GetAllProdi:input_type -> tracer_study_grpc.EmptyProdiRequest
	2, // 5: tracer_study_grpc.ProdiService.GetAllProdi:output_type -> tracer_study_grpc.GetAllProdiResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_prodi_proto_init() }
func file_prodi_proto_init() {
	if File_prodi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prodi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prodi); i {
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
		file_prodi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyProdiRequest); i {
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
		file_prodi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProdiResponse); i {
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
			RawDescriptor: file_prodi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prodi_proto_goTypes,
		DependencyIndexes: file_prodi_proto_depIdxs,
		MessageInfos:      file_prodi_proto_msgTypes,
	}.Build()
	File_prodi_proto = out.File
	file_prodi_proto_rawDesc = nil
	file_prodi_proto_goTypes = nil
	file_prodi_proto_depIdxs = nil
}