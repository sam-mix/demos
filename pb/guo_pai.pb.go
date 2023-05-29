// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: guo_pai.proto

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

// C -> S 玩家过牌操作
type C2SGuoPai struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold int32 `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"` // 前端透传 保留 前端要求，合理性前端提供
}

func (x *C2SGuoPai) Reset() {
	*x = C2SGuoPai{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guo_pai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SGuoPai) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SGuoPai) ProtoMessage() {}

func (x *C2SGuoPai) ProtoReflect() protoreflect.Message {
	mi := &file_guo_pai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SGuoPai.ProtoReflect.Descriptor instead.
func (*C2SGuoPai) Descriptor() ([]byte, []int) {
	return file_guo_pai_proto_rawDescGZIP(), []int{0}
}

func (x *C2SGuoPai) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

// S -> C 玩家过牌操作
type S2CGuoPai struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold int32 `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"` // 前端透传 保留 前端要求，合理性前端提供
}

func (x *S2CGuoPai) Reset() {
	*x = S2CGuoPai{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guo_pai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CGuoPai) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CGuoPai) ProtoMessage() {}

func (x *S2CGuoPai) ProtoReflect() protoreflect.Message {
	mi := &file_guo_pai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CGuoPai.ProtoReflect.Descriptor instead.
func (*S2CGuoPai) Descriptor() ([]byte, []int) {
	return file_guo_pai_proto_rawDescGZIP(), []int{1}
}

func (x *S2CGuoPai) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

var File_guo_pai_proto protoreflect.FileDescriptor

var file_guo_pai_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x75, 0x6f, 0x5f, 0x70, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1f, 0x0a, 0x09, 0x43, 0x32, 0x53, 0x47, 0x75, 0x6f, 0x50, 0x61, 0x69, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x6c, 0x64,
	0x22, 0x1f, 0x0a, 0x09, 0x53, 0x32, 0x43, 0x47, 0x75, 0x6f, 0x50, 0x61, 0x69, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x6c,
	0x64, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_guo_pai_proto_rawDescOnce sync.Once
	file_guo_pai_proto_rawDescData = file_guo_pai_proto_rawDesc
)

func file_guo_pai_proto_rawDescGZIP() []byte {
	file_guo_pai_proto_rawDescOnce.Do(func() {
		file_guo_pai_proto_rawDescData = protoimpl.X.CompressGZIP(file_guo_pai_proto_rawDescData)
	})
	return file_guo_pai_proto_rawDescData
}

var file_guo_pai_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_guo_pai_proto_goTypes = []interface{}{
	(*C2SGuoPai)(nil), // 0: C2SGuoPai
	(*S2CGuoPai)(nil), // 1: S2CGuoPai
}
var file_guo_pai_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_guo_pai_proto_init() }
func file_guo_pai_proto_init() {
	if File_guo_pai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_guo_pai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SGuoPai); i {
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
		file_guo_pai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CGuoPai); i {
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
			RawDescriptor: file_guo_pai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_guo_pai_proto_goTypes,
		DependencyIndexes: file_guo_pai_proto_depIdxs,
		MessageInfos:      file_guo_pai_proto_msgTypes,
	}.Build()
	File_guo_pai_proto = out.File
	file_guo_pai_proto_rawDesc = nil
	file_guo_pai_proto_goTypes = nil
	file_guo_pai_proto_depIdxs = nil
}
