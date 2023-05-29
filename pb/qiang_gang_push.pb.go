// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: qiang_gang_push.proto

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

// C -> S 玩家被抢杠 推送
type C2SQiangGangPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SQiangGangPush) Reset() {
	*x = C2SQiangGangPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qiang_gang_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SQiangGangPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SQiangGangPush) ProtoMessage() {}

func (x *C2SQiangGangPush) ProtoReflect() protoreflect.Message {
	mi := &file_qiang_gang_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SQiangGangPush.ProtoReflect.Descriptor instead.
func (*C2SQiangGangPush) Descriptor() ([]byte, []int) {
	return file_qiang_gang_push_proto_rawDescGZIP(), []int{0}
}

// S -> C 玩家被抢杠 推送
type S2CQiangGangPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatId int32 `protobuf:"varint,1,opt,name=SeatId,proto3" json:"SeatId,omitempty"` // 被抢杠的玩家座位方位
	Card   int32 `protobuf:"varint,2,opt,name=Card,proto3" json:"Card,omitempty"`     // 被抢杠的牌
}

func (x *S2CQiangGangPush) Reset() {
	*x = S2CQiangGangPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qiang_gang_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CQiangGangPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CQiangGangPush) ProtoMessage() {}

func (x *S2CQiangGangPush) ProtoReflect() protoreflect.Message {
	mi := &file_qiang_gang_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CQiangGangPush.ProtoReflect.Descriptor instead.
func (*S2CQiangGangPush) Descriptor() ([]byte, []int) {
	return file_qiang_gang_push_proto_rawDescGZIP(), []int{1}
}

func (x *S2CQiangGangPush) GetSeatId() int32 {
	if x != nil {
		return x.SeatId
	}
	return 0
}

func (x *S2CQiangGangPush) GetCard() int32 {
	if x != nil {
		return x.Card
	}
	return 0
}

var File_qiang_gang_push_proto protoreflect.FileDescriptor

var file_qiang_gang_push_proto_rawDesc = []byte{
	0x0a, 0x15, 0x71, 0x69, 0x61, 0x6e, 0x67, 0x5f, 0x67, 0x61, 0x6e, 0x67, 0x5f, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x32, 0x53, 0x51, 0x69,
	0x61, 0x6e, 0x67, 0x47, 0x61, 0x6e, 0x67, 0x50, 0x75, 0x73, 0x68, 0x22, 0x3e, 0x0a, 0x10, 0x53,
	0x32, 0x43, 0x51, 0x69, 0x61, 0x6e, 0x67, 0x47, 0x61, 0x6e, 0x67, 0x50, 0x75, 0x73, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x65, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x61, 0x72, 0x64, 0x42, 0x0b, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qiang_gang_push_proto_rawDescOnce sync.Once
	file_qiang_gang_push_proto_rawDescData = file_qiang_gang_push_proto_rawDesc
)

func file_qiang_gang_push_proto_rawDescGZIP() []byte {
	file_qiang_gang_push_proto_rawDescOnce.Do(func() {
		file_qiang_gang_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_qiang_gang_push_proto_rawDescData)
	})
	return file_qiang_gang_push_proto_rawDescData
}

var file_qiang_gang_push_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_qiang_gang_push_proto_goTypes = []interface{}{
	(*C2SQiangGangPush)(nil), // 0: C2SQiangGangPush
	(*S2CQiangGangPush)(nil), // 1: S2CQiangGangPush
}
var file_qiang_gang_push_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_qiang_gang_push_proto_init() }
func file_qiang_gang_push_proto_init() {
	if File_qiang_gang_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qiang_gang_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SQiangGangPush); i {
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
		file_qiang_gang_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CQiangGangPush); i {
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
			RawDescriptor: file_qiang_gang_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qiang_gang_push_proto_goTypes,
		DependencyIndexes: file_qiang_gang_push_proto_depIdxs,
		MessageInfos:      file_qiang_gang_push_proto_msgTypes,
	}.Build()
	File_qiang_gang_push_proto = out.File
	file_qiang_gang_push_proto_rawDesc = nil
	file_qiang_gang_push_proto_goTypes = nil
	file_qiang_gang_push_proto_depIdxs = nil
}