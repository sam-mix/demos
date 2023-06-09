// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: login.proto

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

// C -> S 登录信息
type C2SLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold int32  `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"` // 前端透传 保留 前端要求，合理性前端提供
	Acc  string `protobuf:"bytes,2,opt,name=acc,proto3" json:"acc,omitempty"`    // 玩家账号
}

func (x *C2SLogin) Reset() {
	*x = C2SLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SLogin) ProtoMessage() {}

func (x *C2SLogin) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SLogin.ProtoReflect.Descriptor instead.
func (*C2SLogin) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *C2SLogin) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *C2SLogin) GetAcc() string {
	if x != nil {
		return x.Acc
	}
	return ""
}

// S -> C 登录信息
type S2CLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold       int32       `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"`            // 前端透传 保留 前端要求，合理性前端提供
	PlayerInfo *PlayerInfo `protobuf:"bytes,2,opt,name=playerInfo,proto3" json:"playerInfo,omitempty"` // 玩家信息
}

func (x *S2CLogin) Reset() {
	*x = S2CLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CLogin) ProtoMessage() {}

func (x *S2CLogin) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CLogin.ProtoReflect.Descriptor instead.
func (*S2CLogin) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *S2CLogin) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *S2CLogin) GetPlayerInfo() *PlayerInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x08, 0x43,
	0x32, 0x53, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x63, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x63, 0x22, 0x4b, 0x0a,
	0x08, 0x53, 0x32, 0x43, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x2b, 0x0a,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0xaa, 0x02, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_login_proto_goTypes = []interface{}{
	(*C2SLogin)(nil),   // 0: C2SLogin
	(*S2CLogin)(nil),   // 1: S2CLogin
	(*PlayerInfo)(nil), // 2: PlayerInfo
}
var file_login_proto_depIdxs = []int32{
	2, // 0: S2CLogin.playerInfo:type_name -> PlayerInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SLogin); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CLogin); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
