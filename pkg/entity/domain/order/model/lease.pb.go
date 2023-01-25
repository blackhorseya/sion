// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: pb/domain/order/lease.proto

package model

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Lease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	No           string                 `protobuf:"bytes,1,opt,name=no,proto3" json:"no,omitempty"`
	CarId        string                 `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	CarLatitude  float64                `protobuf:"fixed64,3,opt,name=car_latitude,json=carLatitude,proto3" json:"car_latitude,omitempty"`
	CarLongitude float64                `protobuf:"fixed64,4,opt,name=car_longitude,json=carLongitude,proto3" json:"car_longitude,omitempty"`
	StartAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	LastPickAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=last_pick_at,json=lastPickAt,proto3" json:"last_pick_at,omitempty"`
}

func (x *Lease) Reset() {
	*x = Lease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_domain_order_lease_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lease) ProtoMessage() {}

func (x *Lease) ProtoReflect() protoreflect.Message {
	mi := &file_pb_domain_order_lease_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lease.ProtoReflect.Descriptor instead.
func (*Lease) Descriptor() ([]byte, []int) {
	return file_pb_domain_order_lease_proto_rawDescGZIP(), []int{0}
}

func (x *Lease) GetNo() string {
	if x != nil {
		return x.No
	}
	return ""
}

func (x *Lease) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *Lease) GetCarLatitude() float64 {
	if x != nil {
		return x.CarLatitude
	}
	return 0
}

func (x *Lease) GetCarLongitude() float64 {
	if x != nil {
		return x.CarLongitude
	}
	return 0
}

func (x *Lease) GetStartAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartAt
	}
	return nil
}

func (x *Lease) GetEndAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *Lease) GetLastPickAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastPickAt
	}
	return nil
}

var File_pb_domain_order_lease_proto protoreflect.FileDescriptor

var file_pb_domain_order_lease_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x6f, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x5f, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x61,
	0x72, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x72,
	0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x63, 0x61, 0x72, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x35,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74,
	0x50, 0x69, 0x63, 0x6b, 0x41, 0x74, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x79,
	0x61, 0x2f, 0x69, 0x72, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_domain_order_lease_proto_rawDescOnce sync.Once
	file_pb_domain_order_lease_proto_rawDescData = file_pb_domain_order_lease_proto_rawDesc
)

func file_pb_domain_order_lease_proto_rawDescGZIP() []byte {
	file_pb_domain_order_lease_proto_rawDescOnce.Do(func() {
		file_pb_domain_order_lease_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_domain_order_lease_proto_rawDescData)
	})
	return file_pb_domain_order_lease_proto_rawDescData
}

var file_pb_domain_order_lease_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_domain_order_lease_proto_goTypes = []interface{}{
	(*Lease)(nil),                 // 0: order.Lease
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_pb_domain_order_lease_proto_depIdxs = []int32{
	1, // 0: order.Lease.start_at:type_name -> google.protobuf.Timestamp
	1, // 1: order.Lease.end_at:type_name -> google.protobuf.Timestamp
	1, // 2: order.Lease.last_pick_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pb_domain_order_lease_proto_init() }
func file_pb_domain_order_lease_proto_init() {
	if File_pb_domain_order_lease_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_domain_order_lease_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lease); i {
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
			RawDescriptor: file_pb_domain_order_lease_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_domain_order_lease_proto_goTypes,
		DependencyIndexes: file_pb_domain_order_lease_proto_depIdxs,
		MessageInfos:      file_pb_domain_order_lease_proto_msgTypes,
	}.Build()
	File_pb_domain_order_lease_proto = out.File
	file_pb_domain_order_lease_proto_rawDesc = nil
	file_pb_domain_order_lease_proto_goTypes = nil
	file_pb_domain_order_lease_proto_depIdxs = nil
}
