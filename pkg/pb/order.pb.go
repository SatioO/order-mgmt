// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: pkg/pb/order.proto

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

type PaymentMethod int32

const (
	PaymentMethod_WALLET      PaymentMethod = 0
	PaymentMethod_CREDIT_CARD PaymentMethod = 1
)

// Enum value maps for PaymentMethod.
var (
	PaymentMethod_name = map[int32]string{
		0: "WALLET",
		1: "CREDIT_CARD",
	}
	PaymentMethod_value = map[string]int32{
		"WALLET":      0,
		"CREDIT_CARD": 1,
	}
)

func (x PaymentMethod) Enum() *PaymentMethod {
	p := new(PaymentMethod)
	*p = x
	return p
}

func (x PaymentMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_order_proto_enumTypes[0].Descriptor()
}

func (PaymentMethod) Type() protoreflect.EnumType {
	return &file_pkg_pb_order_proto_enumTypes[0]
}

func (x PaymentMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethod.Descriptor instead.
func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_order_proto_rawDescGZIP(), []int{0}
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId       int32         `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	SellerId         int32         `protobuf:"varint,2,opt,name=sellerId,proto3" json:"sellerId,omitempty"`
	Product          *Product      `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	PaymentMethod    PaymentMethod `protobuf:"varint,4,opt,name=paymentMethod,proto3,enum=order.PaymentMethod" json:"paymentMethod,omitempty"`
	DeliveryLocation string        `protobuf:"bytes,5,opt,name=deliveryLocation,proto3" json:"deliveryLocation,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CreateOrderRequest) GetSellerId() int32 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *CreateOrderRequest) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *CreateOrderRequest) GetPaymentMethod() PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return PaymentMethod_WALLET
}

func (x *CreateOrderRequest) GetDeliveryLocation() string {
	if x != nil {
		return x.DeliveryLocation
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   int32  `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=productName,proto3" json:"productName,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Product) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId       int32    `protobuf:"varint,1,opt,name=customerId,proto3" json:"customerId,omitempty"`
	SellerId         int32    `protobuf:"varint,2,opt,name=sellerId,proto3" json:"sellerId,omitempty"`
	Product          *Product `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	PaymentMethod    string   `protobuf:"bytes,4,opt,name=paymentMethod,proto3" json:"paymentMethod,omitempty"`
	DeliveryLocation string   `protobuf:"bytes,5,opt,name=deliveryLocation,proto3" json:"deliveryLocation,omitempty"`
	OrderId          string   `protobuf:"bytes,6,opt,name=orderId,proto3" json:"orderId,omitempty"`
	OrderStatus      string   `protobuf:"bytes,7,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
	CreatedTime      string   `protobuf:"bytes,8,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime      string   `protobuf:"bytes,9,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderResponse) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *OrderResponse) GetSellerId() int32 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *OrderResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *OrderResponse) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *OrderResponse) GetDeliveryLocation() string {
	if x != nil {
		return x.DeliveryLocation
	}
	return ""
}

func (x *OrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderResponse) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *OrderResponse) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *OrderResponse) GetUpdatedTime() string {
	if x != nil {
		return x.UpdatedTime
	}
	return ""
}

var File_pkg_pb_order_proto protoreflect.FileDescriptor

var file_pkg_pb_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0xe2, 0x01, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x3a, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x49, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc7, 0x02, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x2a, 0x2c, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x41, 0x4c, 0x4c, 0x45, 0x54,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x43, 0x41, 0x52,
	0x44, 0x10, 0x01, 0x32, 0x50, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_order_proto_rawDescOnce sync.Once
	file_pkg_pb_order_proto_rawDescData = file_pkg_pb_order_proto_rawDesc
)

func file_pkg_pb_order_proto_rawDescGZIP() []byte {
	file_pkg_pb_order_proto_rawDescOnce.Do(func() {
		file_pkg_pb_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_order_proto_rawDescData)
	})
	return file_pkg_pb_order_proto_rawDescData
}

var file_pkg_pb_order_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_pb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_pb_order_proto_goTypes = []interface{}{
	(PaymentMethod)(0),         // 0: order.PaymentMethod
	(*CreateOrderRequest)(nil), // 1: order.CreateOrderRequest
	(*Product)(nil),            // 2: order.Product
	(*OrderResponse)(nil),      // 3: order.OrderResponse
}
var file_pkg_pb_order_proto_depIdxs = []int32{
	2, // 0: order.CreateOrderRequest.product:type_name -> order.Product
	0, // 1: order.CreateOrderRequest.paymentMethod:type_name -> order.PaymentMethod
	2, // 2: order.OrderResponse.product:type_name -> order.Product
	1, // 3: order.OrderService.CreateOrder:input_type -> order.CreateOrderRequest
	3, // 4: order.OrderService.CreateOrder:output_type -> order.OrderResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_pb_order_proto_init() }
func file_pkg_pb_order_proto_init() {
	if File_pkg_pb_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_pkg_pb_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_pkg_pb_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
			RawDescriptor: file_pkg_pb_order_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_order_proto_goTypes,
		DependencyIndexes: file_pkg_pb_order_proto_depIdxs,
		EnumInfos:         file_pkg_pb_order_proto_enumTypes,
		MessageInfos:      file_pkg_pb_order_proto_msgTypes,
	}.Build()
	File_pkg_pb_order_proto = out.File
	file_pkg_pb_order_proto_rawDesc = nil
	file_pkg_pb_order_proto_goTypes = nil
	file_pkg_pb_order_proto_depIdxs = nil
}