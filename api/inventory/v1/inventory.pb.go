// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: inventory/v1/inventory.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GoodsInvInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Stocks int32 `protobuf:"varint,2,opt,name=stocks,proto3" json:"stocks,omitempty"`
}

func (x *GoodsInvInfo) Reset() {
	*x = GoodsInvInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_v1_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInvInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInvInfo) ProtoMessage() {}

func (x *GoodsInvInfo) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_v1_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInvInfo.ProtoReflect.Descriptor instead.
func (*GoodsInvInfo) Descriptor() ([]byte, []int) {
	return file_inventory_v1_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsInvInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoodsInvInfo) GetStocks() int32 {
	if x != nil {
		return x.Stocks
	}
	return 0
}

type SellInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsInfo []*GoodsInvInfo `protobuf:"bytes,1,rep,name=goodsInfo,proto3" json:"goodsInfo,omitempty"`
}

func (x *SellInfo) Reset() {
	*x = SellInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_v1_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SellInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SellInfo) ProtoMessage() {}

func (x *SellInfo) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_v1_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SellInfo.ProtoReflect.Descriptor instead.
func (*SellInfo) Descriptor() ([]byte, []int) {
	return file_inventory_v1_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *SellInfo) GetGoodsInfo() []*GoodsInvInfo {
	if x != nil {
		return x.GoodsInfo
	}
	return nil
}

var File_inventory_v1_inventory_proto protoreflect.FileDescriptor

var file_inventory_v1_inventory_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x48, 0x0a, 0x08, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3c, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0x94, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x40, 0x0a,
	0x06, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x4b, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x76, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x04,
	0x53, 0x65, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x06, 0x52, 0x65, 0x42, 0x61,
	0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1a, 0x5a, 0x18, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_v1_inventory_proto_rawDescOnce sync.Once
	file_inventory_v1_inventory_proto_rawDescData = file_inventory_v1_inventory_proto_rawDesc
)

func file_inventory_v1_inventory_proto_rawDescGZIP() []byte {
	file_inventory_v1_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_v1_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_v1_inventory_proto_rawDescData)
	})
	return file_inventory_v1_inventory_proto_rawDescData
}

var file_inventory_v1_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_inventory_v1_inventory_proto_goTypes = []interface{}{
	(*GoodsInvInfo)(nil),  // 0: api.inventory.v1.GoodsInvInfo
	(*SellInfo)(nil),      // 1: api.inventory.v1.SellInfo
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_inventory_v1_inventory_proto_depIdxs = []int32{
	0, // 0: api.inventory.v1.SellInfo.goodsInfo:type_name -> api.inventory.v1.GoodsInvInfo
	0, // 1: api.inventory.v1.Inventory.SetInv:input_type -> api.inventory.v1.GoodsInvInfo
	0, // 2: api.inventory.v1.Inventory.InvDetail:input_type -> api.inventory.v1.GoodsInvInfo
	1, // 3: api.inventory.v1.Inventory.Sell:input_type -> api.inventory.v1.SellInfo
	1, // 4: api.inventory.v1.Inventory.ReBack:input_type -> api.inventory.v1.SellInfo
	2, // 5: api.inventory.v1.Inventory.SetInv:output_type -> google.protobuf.Empty
	0, // 6: api.inventory.v1.Inventory.InvDetail:output_type -> api.inventory.v1.GoodsInvInfo
	2, // 7: api.inventory.v1.Inventory.Sell:output_type -> google.protobuf.Empty
	2, // 8: api.inventory.v1.Inventory.ReBack:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inventory_v1_inventory_proto_init() }
func file_inventory_v1_inventory_proto_init() {
	if File_inventory_v1_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_v1_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInvInfo); i {
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
		file_inventory_v1_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SellInfo); i {
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
			RawDescriptor: file_inventory_v1_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_v1_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_v1_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_v1_inventory_proto_msgTypes,
	}.Build()
	File_inventory_v1_inventory_proto = out.File
	file_inventory_v1_inventory_proto_rawDesc = nil
	file_inventory_v1_inventory_proto_goTypes = nil
	file_inventory_v1_inventory_proto_depIdxs = nil
}
