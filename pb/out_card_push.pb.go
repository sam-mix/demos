// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: out_card_push.proto

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

// C -> S 用户出牌 推送
type C2SOutCardPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SOutCardPush) Reset() {
	*x = C2SOutCardPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_out_card_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SOutCardPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SOutCardPush) ProtoMessage() {}

func (x *C2SOutCardPush) ProtoReflect() protoreflect.Message {
	mi := &file_out_card_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SOutCardPush.ProtoReflect.Descriptor instead.
func (*C2SOutCardPush) Descriptor() ([]byte, []int) {
	return file_out_card_push_proto_rawDescGZIP(), []int{0}
}

// S -> C 用户出牌 推送
type S2COutCardPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatSeq int32 `protobuf:"varint,1,opt,name=SeatSeq,proto3" json:"SeatSeq,omitempty"` //出牌的玩家座位方位
	Card    int32 `protobuf:"varint,2,opt,name=Card,proto3" json:"Card,omitempty"`       //出的牌
}

func (x *S2COutCardPush) Reset() {
	*x = S2COutCardPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_out_card_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2COutCardPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2COutCardPush) ProtoMessage() {}

func (x *S2COutCardPush) ProtoReflect() protoreflect.Message {
	mi := &file_out_card_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2COutCardPush.ProtoReflect.Descriptor instead.
func (*S2COutCardPush) Descriptor() ([]byte, []int) {
	return file_out_card_push_proto_rawDescGZIP(), []int{1}
}

func (x *S2COutCardPush) GetSeatSeq() int32 {
	if x != nil {
		return x.SeatSeq
	}
	return 0
}

func (x *S2COutCardPush) GetCard() int32 {
	if x != nil {
		return x.Card
	}
	return 0
}

var File_out_card_push_proto protoreflect.FileDescriptor

var file_out_card_push_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x75, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x32, 0x53, 0x4f, 0x75, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x50, 0x75, 0x73, 0x68, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x32, 0x43, 0x4f, 0x75,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x50, 0x75, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x61,
	0x74, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x65, 0x61, 0x74,
	0x53, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x43, 0x61, 0x72, 0x64, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xaa,
	0x02, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_out_card_push_proto_rawDescOnce sync.Once
	file_out_card_push_proto_rawDescData = file_out_card_push_proto_rawDesc
)

func file_out_card_push_proto_rawDescGZIP() []byte {
	file_out_card_push_proto_rawDescOnce.Do(func() {
		file_out_card_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_out_card_push_proto_rawDescData)
	})
	return file_out_card_push_proto_rawDescData
}

var file_out_card_push_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_out_card_push_proto_goTypes = []interface{}{
	(*C2SOutCardPush)(nil), // 0: C2SOutCardPush
	(*S2COutCardPush)(nil), // 1: S2COutCardPush
}
var file_out_card_push_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_out_card_push_proto_init() }
func file_out_card_push_proto_init() {
	if File_out_card_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_out_card_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SOutCardPush); i {
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
		file_out_card_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2COutCardPush); i {
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
			RawDescriptor: file_out_card_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_out_card_push_proto_goTypes,
		DependencyIndexes: file_out_card_push_proto_depIdxs,
		MessageInfos:      file_out_card_push_proto_msgTypes,
	}.Build()
	File_out_card_push_proto = out.File
	file_out_card_push_proto_rawDesc = nil
	file_out_card_push_proto_goTypes = nil
	file_out_card_push_proto_depIdxs = nil
}