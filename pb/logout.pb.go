// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: logout.proto

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

// C -> S 登出
type C2SLogout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SLogout) Reset() {
	*x = C2SLogout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SLogout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SLogout) ProtoMessage() {}

func (x *C2SLogout) ProtoReflect() protoreflect.Message {
	mi := &file_logout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SLogout.ProtoReflect.Descriptor instead.
func (*C2SLogout) Descriptor() ([]byte, []int) {
	return file_logout_proto_rawDescGZIP(), []int{0}
}

// S -> C 登出
type S2CLogout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code RESULT `protobuf:"varint,1,opt,name=code,proto3,enum=RESULT" json:"code,omitempty"` // 响应码
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`                // 响应信息
}

func (x *S2CLogout) Reset() {
	*x = S2CLogout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CLogout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CLogout) ProtoMessage() {}

func (x *S2CLogout) ProtoReflect() protoreflect.Message {
	mi := &file_logout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CLogout.ProtoReflect.Descriptor instead.
func (*S2CLogout) Descriptor() ([]byte, []int) {
	return file_logout_proto_rawDescGZIP(), []int{1}
}

func (x *S2CLogout) GetCode() RESULT {
	if x != nil {
		return x.Code
	}
	return RESULT_SUCCESS
}

func (x *S2CLogout) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_logout_proto protoreflect.FileDescriptor

var file_logout_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09,
	0x43, 0x32, 0x53, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x22, 0x3a, 0x0a, 0x09, 0x53, 0x32, 0x43,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logout_proto_rawDescOnce sync.Once
	file_logout_proto_rawDescData = file_logout_proto_rawDesc
)

func file_logout_proto_rawDescGZIP() []byte {
	file_logout_proto_rawDescOnce.Do(func() {
		file_logout_proto_rawDescData = protoimpl.X.CompressGZIP(file_logout_proto_rawDescData)
	})
	return file_logout_proto_rawDescData
}

var file_logout_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logout_proto_goTypes = []interface{}{
	(*C2SLogout)(nil), // 0: C2SLogout
	(*S2CLogout)(nil), // 1: S2CLogout
	(RESULT)(0),       // 2: RESULT
}
var file_logout_proto_depIdxs = []int32{
	2, // 0: S2CLogout.code:type_name -> RESULT
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_logout_proto_init() }
func file_logout_proto_init() {
	if File_logout_proto != nil {
		return
	}
	file_consts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_logout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SLogout); i {
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
		file_logout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CLogout); i {
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
			RawDescriptor: file_logout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_logout_proto_goTypes,
		DependencyIndexes: file_logout_proto_depIdxs,
		MessageInfos:      file_logout_proto_msgTypes,
	}.Build()
	File_logout_proto = out.File
	file_logout_proto_rawDesc = nil
	file_logout_proto_goTypes = nil
	file_logout_proto_depIdxs = nil
}
