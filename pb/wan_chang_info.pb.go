// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: wan_chang_info.proto

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

// C -> S 玩法场次信息
type C2SWanChangInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold int32    `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"`               // 前端透传 保留 前端要求，合理性前端提供
	Type WAN_TYPE `protobuf:"varint,2,opt,name=type,proto3,enum=WAN_TYPE" json:"type,omitempty"` // 主玩法类型
}

func (x *C2SWanChangInfo) Reset() {
	*x = C2SWanChangInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wan_chang_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2SWanChangInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2SWanChangInfo) ProtoMessage() {}

func (x *C2SWanChangInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wan_chang_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2SWanChangInfo.ProtoReflect.Descriptor instead.
func (*C2SWanChangInfo) Descriptor() ([]byte, []int) {
	return file_wan_chang_info_proto_rawDescGZIP(), []int{0}
}

func (x *C2SWanChangInfo) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *C2SWanChangInfo) GetType() WAN_TYPE {
	if x != nil {
		return x.Type
	}
	return WAN_TYPE_WAN_TYPE_UNDEFINED
}

// 玩法场次信息
type WanChangInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                            // 场次ID
	Name            string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                         // 场次名称
	CardId          string     `protobuf:"bytes,3,opt,name=cardId,proto3" json:"cardId,omitempty"`                     // 卡片样式 ID
	TagId           string     `protobuf:"bytes,4,opt,name=tagId,proto3" json:"tagId,omitempty"`                       // 展示标签
	DiFen           string     `protobuf:"bytes,5,opt,name=diFen,proto3" json:"diFen,omitempty"`                       // 底分
	Fee             string     `protobuf:"bytes,6,opt,name=fee,proto3" json:"fee,omitempty"`                           // 门票
	FengDingBeiShu  string     `protobuf:"bytes,7,opt,name=fengDingBeiShu,proto3" json:"fengDingBeiShu,omitempty"`     // 封顶倍数
	TuiJianJinChang string     `protobuf:"bytes,8,opt,name=tuiJianJinChang,proto3" json:"tuiJianJinChang,omitempty"`   // 推荐进场
	LimitMin        string     `protobuf:"bytes,9,opt,name=limitMin,proto3" json:"limitMin,omitempty"`                 // 场次下限
	LimitMax        string     `protobuf:"bytes,10,opt,name=limitMax,proto3" json:"limitMax,omitempty"`                // 场次上限
	PopGift         TRUE_FALSE `protobuf:"varint,11,opt,name=popGift,proto3,enum=TRUE_FALSE" json:"popGift,omitempty"` // 是否弹出破产礼包
	WanType         WAN_TYPE   `protobuf:"varint,12,opt,name=wanType,proto3,enum=WAN_TYPE" json:"wanType,omitempty"`   // 主玩法类型 (玩法ID)
}

func (x *WanChangInfo) Reset() {
	*x = WanChangInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wan_chang_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WanChangInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WanChangInfo) ProtoMessage() {}

func (x *WanChangInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wan_chang_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WanChangInfo.ProtoReflect.Descriptor instead.
func (*WanChangInfo) Descriptor() ([]byte, []int) {
	return file_wan_chang_info_proto_rawDescGZIP(), []int{1}
}

func (x *WanChangInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WanChangInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WanChangInfo) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *WanChangInfo) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *WanChangInfo) GetDiFen() string {
	if x != nil {
		return x.DiFen
	}
	return ""
}

func (x *WanChangInfo) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *WanChangInfo) GetFengDingBeiShu() string {
	if x != nil {
		return x.FengDingBeiShu
	}
	return ""
}

func (x *WanChangInfo) GetTuiJianJinChang() string {
	if x != nil {
		return x.TuiJianJinChang
	}
	return ""
}

func (x *WanChangInfo) GetLimitMin() string {
	if x != nil {
		return x.LimitMin
	}
	return ""
}

func (x *WanChangInfo) GetLimitMax() string {
	if x != nil {
		return x.LimitMax
	}
	return ""
}

func (x *WanChangInfo) GetPopGift() TRUE_FALSE {
	if x != nil {
		return x.PopGift
	}
	return TRUE_FALSE_FALSE
}

func (x *WanChangInfo) GetWanType() WAN_TYPE {
	if x != nil {
		return x.WanType
	}
	return WAN_TYPE_WAN_TYPE_UNDEFINED
}

// S -> C 玩法场次信息
type S2CWanChangInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hold           int32           `protobuf:"varint,1,opt,name=hold,proto3" json:"hold,omitempty"`                     // 前端透传 保留 前端要求，合理性前端提供
	MainInfoList   []*WanChangInfo `protobuf:"bytes,2,rep,name=mainInfoList,proto3" json:"mainInfoList,omitempty"`      // 主玩法场次信息列表
	SecondName     string          `protobuf:"bytes,3,opt,name=secondName,proto3" json:"secondName,omitempty"`          // 二级玩法标题名称
	SecondInfoList []*WanChangInfo `protobuf:"bytes,4,rep,name=secondInfoList,proto3" json:"secondInfoList,omitempty"`  // 二级玩法场次信息列表
	ThirdName      string          `protobuf:"bytes,5,opt,name=thirdName,proto3" json:"thirdName,omitempty"`            // 三级玩法标题名称
	ThirdInfoList  []*WanChangInfo `protobuf:"bytes,6,rep,name=thirdInfoList,proto3" json:"thirdInfoList,omitempty"`    // 三级玩法场次信息列表
	WanType        WAN_TYPE        `protobuf:"varint,7,opt,name=wanType,proto3,enum=WAN_TYPE" json:"wanType,omitempty"` // 主玩法类型 (玩法ID)
}

func (x *S2CWanChangInfo) Reset() {
	*x = S2CWanChangInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wan_chang_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2CWanChangInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2CWanChangInfo) ProtoMessage() {}

func (x *S2CWanChangInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wan_chang_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2CWanChangInfo.ProtoReflect.Descriptor instead.
func (*S2CWanChangInfo) Descriptor() ([]byte, []int) {
	return file_wan_chang_info_proto_rawDescGZIP(), []int{2}
}

func (x *S2CWanChangInfo) GetHold() int32 {
	if x != nil {
		return x.Hold
	}
	return 0
}

func (x *S2CWanChangInfo) GetMainInfoList() []*WanChangInfo {
	if x != nil {
		return x.MainInfoList
	}
	return nil
}

func (x *S2CWanChangInfo) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *S2CWanChangInfo) GetSecondInfoList() []*WanChangInfo {
	if x != nil {
		return x.SecondInfoList
	}
	return nil
}

func (x *S2CWanChangInfo) GetThirdName() string {
	if x != nil {
		return x.ThirdName
	}
	return ""
}

func (x *S2CWanChangInfo) GetThirdInfoList() []*WanChangInfo {
	if x != nil {
		return x.ThirdInfoList
	}
	return nil
}

func (x *S2CWanChangInfo) GetWanType() WAN_TYPE {
	if x != nil {
		return x.WanType
	}
	return WAN_TYPE_WAN_TYPE_UNDEFINED
}

var File_wan_chang_info_proto protoreflect.FileDescriptor

var file_wan_chang_info_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x61, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0f, 0x43, 0x32, 0x53, 0x57, 0x61, 0x6e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x57, 0x41, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xde, 0x02, 0x0a, 0x0c, 0x57,
	0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x69, 0x46, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x69,
	0x46, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x66, 0x65, 0x6e, 0x67, 0x44, 0x69, 0x6e,
	0x67, 0x42, 0x65, 0x69, 0x53, 0x68, 0x75, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66,
	0x65, 0x6e, 0x67, 0x44, 0x69, 0x6e, 0x67, 0x42, 0x65, 0x69, 0x53, 0x68, 0x75, 0x12, 0x28, 0x0a,
	0x0f, 0x74, 0x75, 0x69, 0x4a, 0x69, 0x61, 0x6e, 0x4a, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x75, 0x69, 0x4a, 0x69, 0x61, 0x6e, 0x4a,
	0x69, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x4d, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x4d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61, 0x78, 0x12,
	0x25, 0x0a, 0x07, 0x70, 0x6f, 0x70, 0x47, 0x69, 0x66, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x54, 0x52, 0x55, 0x45, 0x5f, 0x46, 0x41, 0x4c, 0x53, 0x45, 0x52, 0x07, 0x70,
	0x6f, 0x70, 0x47, 0x69, 0x66, 0x74, 0x12, 0x23, 0x0a, 0x07, 0x77, 0x61, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x57, 0x41, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x52, 0x07, 0x77, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa7, 0x02, 0x0a, 0x0f,
	0x53, 0x32, 0x43, 0x57, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68,
	0x6f, 0x6c, 0x64, 0x12, 0x31, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x57, 0x61, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x57, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x68, 0x69, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x68, 0x69, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0d, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x57, 0x61, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x74, 0x68, 0x69, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x07, 0x77, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x09, 0x2e, 0x57, 0x41, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x07, 0x77, 0x61,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0xaa, 0x02, 0x02,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wan_chang_info_proto_rawDescOnce sync.Once
	file_wan_chang_info_proto_rawDescData = file_wan_chang_info_proto_rawDesc
)

func file_wan_chang_info_proto_rawDescGZIP() []byte {
	file_wan_chang_info_proto_rawDescOnce.Do(func() {
		file_wan_chang_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_wan_chang_info_proto_rawDescData)
	})
	return file_wan_chang_info_proto_rawDescData
}

var file_wan_chang_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wan_chang_info_proto_goTypes = []interface{}{
	(*C2SWanChangInfo)(nil), // 0: C2SWanChangInfo
	(*WanChangInfo)(nil),    // 1: WanChangInfo
	(*S2CWanChangInfo)(nil), // 2: S2CWanChangInfo
	(WAN_TYPE)(0),           // 3: WAN_TYPE
	(TRUE_FALSE)(0),         // 4: TRUE_FALSE
}
var file_wan_chang_info_proto_depIdxs = []int32{
	3, // 0: C2SWanChangInfo.type:type_name -> WAN_TYPE
	4, // 1: WanChangInfo.popGift:type_name -> TRUE_FALSE
	3, // 2: WanChangInfo.wanType:type_name -> WAN_TYPE
	1, // 3: S2CWanChangInfo.mainInfoList:type_name -> WanChangInfo
	1, // 4: S2CWanChangInfo.secondInfoList:type_name -> WanChangInfo
	1, // 5: S2CWanChangInfo.thirdInfoList:type_name -> WanChangInfo
	3, // 6: S2CWanChangInfo.wanType:type_name -> WAN_TYPE
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_wan_chang_info_proto_init() }
func file_wan_chang_info_proto_init() {
	if File_wan_chang_info_proto != nil {
		return
	}
	file_consts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wan_chang_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2SWanChangInfo); i {
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
		file_wan_chang_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WanChangInfo); i {
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
		file_wan_chang_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2CWanChangInfo); i {
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
			RawDescriptor: file_wan_chang_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wan_chang_info_proto_goTypes,
		DependencyIndexes: file_wan_chang_info_proto_depIdxs,
		MessageInfos:      file_wan_chang_info_proto_msgTypes,
	}.Build()
	File_wan_chang_info_proto = out.File
	file_wan_chang_info_proto_rawDesc = nil
	file_wan_chang_info_proto_goTypes = nil
	file_wan_chang_info_proto_depIdxs = nil
}
