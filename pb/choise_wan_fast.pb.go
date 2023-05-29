// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: choise_wan_fast.proto

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

// C -> S 快速选玩法
type C2SChoiseWanFast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold int32 `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"` // 前端透传 保留 前端要求，合理性前端提供
}

func (x *C2SChoiseWanFast) Reset() {
	*x = C2SChoiseWanFast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_choise_wan_fast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SChoiseWanFast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SChoiseWanFast) ProtoMessage() {}

func (x *C2SChoiseWanFast) ProtoReflect() protoreflect.Message {
	mi := &file_choise_wan_fast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SChoiseWanFast.ProtoReflect.Descriptor instead.
func (*C2SChoiseWanFast) Descriptor() ([]byte, []int) {
	return file_choise_wan_fast_proto_rawDescGZIP(), []int{0}
}

func (x *C2SChoiseWanFast) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

// S -> C 快速选玩法
type S2CChoiseWanFast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold int32 `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"` // 前端透传 保留 前端要求，合理性前端提供
	// string ip = 1;// 玩法服务的IP或域名
	// uint32 port = 2; // 玩法服务的端口
	Code RESULT `protobuf:"varint,2,opt,name=code,proto3,enum=RESULT" json:"code,omitempty"` // 响应码
	Msg  string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`                // 响应信息
}

func (x *S2CChoiseWanFast) Reset() {
	*x = S2CChoiseWanFast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_choise_wan_fast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CChoiseWanFast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CChoiseWanFast) ProtoMessage() {}

func (x *S2CChoiseWanFast) ProtoReflect() protoreflect.Message {
	mi := &file_choise_wan_fast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CChoiseWanFast.ProtoReflect.Descriptor instead.
func (*S2CChoiseWanFast) Descriptor() ([]byte, []int) {
	return file_choise_wan_fast_proto_rawDescGZIP(), []int{1}
}

func (x *S2CChoiseWanFast) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *S2CChoiseWanFast) GetCode() RESULT {
	if x != nil {
		return x.Code
	}
	return RESULT_SUCCESS
}

func (x *S2CChoiseWanFast) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_choise_wan_fast_proto protoreflect.FileDescriptor

var file_choise_wan_fast_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x5f, 0x77, 0x61, 0x6e, 0x5f, 0x66, 0x61, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x32, 0x53, 0x43, 0x68, 0x6f, 0x69,
	0x73, 0x65, 0x57, 0x61, 0x6e, 0x46, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0x55, 0x0a,
	0x10, 0x53, 0x32, 0x43, 0x43, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x57, 0x61, 0x6e, 0x46, 0x61, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_choise_wan_fast_proto_rawDescOnce sync.Once
	file_choise_wan_fast_proto_rawDescData = file_choise_wan_fast_proto_rawDesc
)

func file_choise_wan_fast_proto_rawDescGZIP() []byte {
	file_choise_wan_fast_proto_rawDescOnce.Do(func() {
		file_choise_wan_fast_proto_rawDescData = protoimpl.X.CompressGZIP(file_choise_wan_fast_proto_rawDescData)
	})
	return file_choise_wan_fast_proto_rawDescData
}

var file_choise_wan_fast_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_choise_wan_fast_proto_goTypes = []interface{}{
	(*C2SChoiseWanFast)(nil), // 0: C2SChoiseWanFast
	(*S2CChoiseWanFast)(nil), // 1: S2CChoiseWanFast
	(RESULT)(0),              // 2: RESULT
}
var file_choise_wan_fast_proto_depIdxs = []int32{
	2, // 0: S2CChoiseWanFast.code:type_name -> RESULT
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_choise_wan_fast_proto_init() }
func file_choise_wan_fast_proto_init() {
	if File_choise_wan_fast_proto != nil {
		return
	}
	file_consts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_choise_wan_fast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SChoiseWanFast); i {
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
		file_choise_wan_fast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CChoiseWanFast); i {
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
			RawDescriptor: file_choise_wan_fast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_choise_wan_fast_proto_goTypes,
		DependencyIndexes: file_choise_wan_fast_proto_depIdxs,
		MessageInfos:      file_choise_wan_fast_proto_msgTypes,
	}.Build()
	File_choise_wan_fast_proto = out.File
	file_choise_wan_fast_proto_rawDesc = nil
	file_choise_wan_fast_proto_goTypes = nil
	file_choise_wan_fast_proto_depIdxs = nil
}
