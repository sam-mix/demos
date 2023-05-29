// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: choise_wan.proto

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

// C -> S 选玩法
type C2SChoiseWan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold    int32    `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"`                      // 前端透传 保留 前端要求，合理性前端提供
	Type    WAN_TYPE `protobuf:"varint,2,opt,name=type,proto3,enum=WAN_TYPE" json:"type,omitempty"`        // 玩法类型
	ChangCi uint64   `protobuf:"varint,3,opt,name=chang_ci,json=changCi,proto3" json:"chang_ci,omitempty"` // 场次ID 传0表示快速开始
}

func (x *C2SChoiseWan) Reset() {
	*x = C2SChoiseWan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_choise_wan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SChoiseWan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SChoiseWan) ProtoMessage() {}

func (x *C2SChoiseWan) ProtoReflect() protoreflect.Message {
	mi := &file_choise_wan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SChoiseWan.ProtoReflect.Descriptor instead.
func (*C2SChoiseWan) Descriptor() ([]byte, []int) {
	return file_choise_wan_proto_rawDescGZIP(), []int{0}
}

func (x *C2SChoiseWan) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *C2SChoiseWan) GetType() WAN_TYPE {
	if x != nil {
		return x.Type
	}
	return WAN_TYPE_WAN_TYPE_UNDEFINED
}

func (x *C2SChoiseWan) GetChangCi() uint64 {
	if x != nil {
		return x.ChangCi
	}
	return 0
}

// S -> C 选玩法
type S2CChoiseWan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold int32  `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"`             // 前端透传 保留 前端要求，合理性前端提供
	Code RESULT `protobuf:"varint,2,opt,name=code,proto3,enum=RESULT" json:"code,omitempty"` // 响应码
	Msg  string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`                // 响应信息
}

func (x *S2CChoiseWan) Reset() {
	*x = S2CChoiseWan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_choise_wan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CChoiseWan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CChoiseWan) ProtoMessage() {}

func (x *S2CChoiseWan) ProtoReflect() protoreflect.Message {
	mi := &file_choise_wan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CChoiseWan.ProtoReflect.Descriptor instead.
func (*S2CChoiseWan) Descriptor() ([]byte, []int) {
	return file_choise_wan_proto_rawDescGZIP(), []int{1}
}

func (x *S2CChoiseWan) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *S2CChoiseWan) GetCode() RESULT {
	if x != nil {
		return x.Code
	}
	return RESULT_SUCCESS
}

func (x *S2CChoiseWan) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_choise_wan_proto protoreflect.FileDescriptor

var file_choise_wan_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x5f, 0x77, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5c, 0x0a, 0x0c, 0x43, 0x32, 0x53, 0x43, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x57, 0x61, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x09, 0x2e, 0x57, 0x41, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x5f, 0x63, 0x69, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x43, 0x69, 0x22, 0x51,
	0x0a, 0x0c, 0x53, 0x32, 0x43, 0x43, 0x68, 0x6f, 0x69, 0x73, 0x65, 0x57, 0x61, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f,
	0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x07, 0x2e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_choise_wan_proto_rawDescOnce sync.Once
	file_choise_wan_proto_rawDescData = file_choise_wan_proto_rawDesc
)

func file_choise_wan_proto_rawDescGZIP() []byte {
	file_choise_wan_proto_rawDescOnce.Do(func() {
		file_choise_wan_proto_rawDescData = protoimpl.X.CompressGZIP(file_choise_wan_proto_rawDescData)
	})
	return file_choise_wan_proto_rawDescData
}

var file_choise_wan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_choise_wan_proto_goTypes = []interface{}{
	(*C2SChoiseWan)(nil), // 0: C2SChoiseWan
	(*S2CChoiseWan)(nil), // 1: S2CChoiseWan
	(WAN_TYPE)(0),        // 2: WAN_TYPE
	(RESULT)(0),          // 3: RESULT
}
var file_choise_wan_proto_depIdxs = []int32{
	2, // 0: C2SChoiseWan.type:type_name -> WAN_TYPE
	3, // 1: S2CChoiseWan.code:type_name -> RESULT
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_choise_wan_proto_init() }
func file_choise_wan_proto_init() {
	if File_choise_wan_proto != nil {
		return
	}
	file_consts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_choise_wan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SChoiseWan); i {
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
		file_choise_wan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CChoiseWan); i {
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
			RawDescriptor: file_choise_wan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_choise_wan_proto_goTypes,
		DependencyIndexes: file_choise_wan_proto_depIdxs,
		MessageInfos:      file_choise_wan_proto_msgTypes,
	}.Build()
	File_choise_wan_proto = out.File
	file_choise_wan_proto_rawDesc = nil
	file_choise_wan_proto_goTypes = nil
	file_choise_wan_proto_depIdxs = nil
}
