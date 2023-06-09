// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: hu_push.proto

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

// 胡牌输赢消息体
type HuWinLoss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeatSeq         int32   `protobuf:"varint,1,opt,name=SeatSeq,proto3" json:"SeatSeq,omitempty"`                 // 座位方位
	HuCard          int32   `protobuf:"varint,2,opt,name=HuCard,proto3" json:"HuCard,omitempty"`                   // 胡的牌
	SeatWinLoss     string  `protobuf:"bytes,3,opt,name=SeatWinLoss,proto3" json:"SeatWinLoss,omitempty"`          // 座位输赢
	WinMulti        int64   `protobuf:"varint,4,opt,name=WinMulti,proto3" json:"WinMulti,omitempty"`               // 赢得倍数
	GiveCardSeatSeq int32   `protobuf:"varint,5,opt,name=GiveCardSeatSeq,proto3" json:"GiveCardSeatSeq,omitempty"` // 打出这张牌的玩家座位方位
	IsQiangGang     bool    `protobuf:"varint,6,opt,name=IsQiangGang,proto3" json:"IsQiangGang,omitempty"`         // 是否为抢杠胡
	IsAfterGang     bool    `protobuf:"varint,7,opt,name=IsAfterGang,proto3" json:"IsAfterGang,omitempty"`         // 是否为杠上开花
	IsSkyHu         bool    `protobuf:"varint,8,opt,name=IsSkyHu,proto3" json:"IsSkyHu,omitempty"`                 // 是否为天胡
	IsDiHu          bool    `protobuf:"varint,9,opt,name=IsDiHu,proto3" json:"IsDiHu,omitempty"`                   // 是否为地胡
	Extend          string  `protobuf:"bytes,10,opt,name=Extend,proto3" json:"Extend,omitempty"`                   // {"type":"自摸，点炮","multi":[1,2,4,8,16,32...],"desc":"杠上开花—自摸,点炮，呼叫转移"}
	HandCards       []int32 `protobuf:"varint,17,rep,packed,name=HandCards,proto3" json:"HandCards,omitempty"`     // 手牌
}

func (x *HuWinLoss) Reset() {
	*x = HuWinLoss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hu_push_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuWinLoss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuWinLoss) ProtoMessage() {}

func (x *HuWinLoss) ProtoReflect() protoreflect.Message {
	mi := &file_hu_push_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuWinLoss.ProtoReflect.Descriptor instead.
func (*HuWinLoss) Descriptor() ([]byte, []int) {
	return file_hu_push_proto_rawDescGZIP(), []int{0}
}

func (x *HuWinLoss) GetSeatSeq() int32 {
	if x != nil {
		return x.SeatSeq
	}
	return 0
}

func (x *HuWinLoss) GetHuCard() int32 {
	if x != nil {
		return x.HuCard
	}
	return 0
}

func (x *HuWinLoss) GetSeatWinLoss() string {
	if x != nil {
		return x.SeatWinLoss
	}
	return ""
}

func (x *HuWinLoss) GetWinMulti() int64 {
	if x != nil {
		return x.WinMulti
	}
	return 0
}

func (x *HuWinLoss) GetGiveCardSeatSeq() int32 {
	if x != nil {
		return x.GiveCardSeatSeq
	}
	return 0
}

func (x *HuWinLoss) GetIsQiangGang() bool {
	if x != nil {
		return x.IsQiangGang
	}
	return false
}

func (x *HuWinLoss) GetIsAfterGang() bool {
	if x != nil {
		return x.IsAfterGang
	}
	return false
}

func (x *HuWinLoss) GetIsSkyHu() bool {
	if x != nil {
		return x.IsSkyHu
	}
	return false
}

func (x *HuWinLoss) GetIsDiHu() bool {
	if x != nil {
		return x.IsDiHu
	}
	return false
}

func (x *HuWinLoss) GetExtend() string {
	if x != nil {
		return x.Extend
	}
	return ""
}

func (x *HuWinLoss) GetHandCards() []int32 {
	if x != nil {
		return x.HandCards
	}
	return nil
}

// C -> S 玩家胡牌 推送
type C2SHuPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2SHuPush) Reset() {
	*x = C2SHuPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hu_push_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SHuPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SHuPush) ProtoMessage() {}

func (x *C2SHuPush) ProtoReflect() protoreflect.Message {
	mi := &file_hu_push_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SHuPush.ProtoReflect.Descriptor instead.
func (*C2SHuPush) Descriptor() ([]byte, []int) {
	return file_hu_push_proto_rawDescGZIP(), []int{1}
}

// S -> C 玩家胡牌 推送
type S2CHuPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WinLossList []*HuWinLoss `protobuf:"bytes,1,rep,name=WinLossList,proto3" json:"WinLossList,omitempty"` // 输赢列表
}

func (x *S2CHuPush) Reset() {
	*x = S2CHuPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hu_push_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CHuPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CHuPush) ProtoMessage() {}

func (x *S2CHuPush) ProtoReflect() protoreflect.Message {
	mi := &file_hu_push_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CHuPush.ProtoReflect.Descriptor instead.
func (*S2CHuPush) Descriptor() ([]byte, []int) {
	return file_hu_push_proto_rawDescGZIP(), []int{2}
}

func (x *S2CHuPush) GetWinLossList() []*HuWinLoss {
	if x != nil {
		return x.WinLossList
	}
	return nil
}

var File_hu_push_proto protoreflect.FileDescriptor

var file_hu_push_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x75, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd1, 0x02, 0x0a, 0x09, 0x48, 0x75, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x65, 0x61, 0x74, 0x53, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x53, 0x65, 0x61, 0x74, 0x53, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x75, 0x43, 0x61, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48, 0x75, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x74, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x61, 0x74, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x69, 0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x57, 0x69, 0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x28, 0x0a,
	0x0f, 0x47, 0x69, 0x76, 0x65, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x61, 0x74, 0x53, 0x65, 0x71,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x47, 0x69, 0x76, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x53, 0x65, 0x61, 0x74, 0x53, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x51, 0x69, 0x61,
	0x6e, 0x67, 0x47, 0x61, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x73,
	0x51, 0x69, 0x61, 0x6e, 0x67, 0x47, 0x61, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x49, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x49,
	0x73, 0x53, 0x6b, 0x79, 0x48, 0x75, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73,
	0x53, 0x6b, 0x79, 0x48, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x44, 0x69, 0x48, 0x75, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x44, 0x69, 0x48, 0x75, 0x12, 0x16, 0x0a,
	0x06, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x32, 0x53, 0x48, 0x75, 0x50, 0x75, 0x73, 0x68,
	0x22, 0x39, 0x0a, 0x09, 0x53, 0x32, 0x43, 0x48, 0x75, 0x50, 0x75, 0x73, 0x68, 0x12, 0x2c, 0x0a,
	0x0b, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x48, 0x75, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x52, 0x0b,
	0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0b, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hu_push_proto_rawDescOnce sync.Once
	file_hu_push_proto_rawDescData = file_hu_push_proto_rawDesc
)

func file_hu_push_proto_rawDescGZIP() []byte {
	file_hu_push_proto_rawDescOnce.Do(func() {
		file_hu_push_proto_rawDescData = protoimpl.X.CompressGZIP(file_hu_push_proto_rawDescData)
	})
	return file_hu_push_proto_rawDescData
}

var file_hu_push_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hu_push_proto_goTypes = []interface{}{
	(*HuWinLoss)(nil), // 0: HuWinLoss
	(*C2SHuPush)(nil), // 1: C2SHuPush
	(*S2CHuPush)(nil), // 2: S2CHuPush
}
var file_hu_push_proto_depIdxs = []int32{
	0, // 0: S2CHuPush.WinLossList:type_name -> HuWinLoss
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hu_push_proto_init() }
func file_hu_push_proto_init() {
	if File_hu_push_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hu_push_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuWinLoss); i {
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
		file_hu_push_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SHuPush); i {
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
		file_hu_push_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CHuPush); i {
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
			RawDescriptor: file_hu_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hu_push_proto_goTypes,
		DependencyIndexes: file_hu_push_proto_depIdxs,
		MessageInfos:      file_hu_push_proto_msgTypes,
	}.Build()
	File_hu_push_proto = out.File
	file_hu_push_proto_rawDesc = nil
	file_hu_push_proto_goTypes = nil
	file_hu_push_proto_depIdxs = nil
}
