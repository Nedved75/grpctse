// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: relay/processing/v1/models.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation int32

const (
	Operation_AUTHORIZE Operation = 0
	Operation_CANCEL    Operation = 1
	Operation_CHARGE    Operation = 2
	Operation_REFUND    Operation = 3
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "AUTHORIZE",
		1: "CANCEL",
		2: "CHARGE",
		3: "REFUND",
	}
	Operation_value = map[string]int32{
		"AUTHORIZE": 0,
		"CANCEL":    1,
		"CHARGE":    2,
		"REFUND":    3,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_relay_processing_v1_models_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_relay_processing_v1_models_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_relay_processing_v1_models_proto_rawDescGZIP(), []int{0}
}

type OperationStatus int32

const (
	OperationStatus_PENDING   OperationStatus = 0
	OperationStatus_COMPLETED OperationStatus = 1
	OperationStatus_FAILED    OperationStatus = 2
)

// Enum value maps for OperationStatus.
var (
	OperationStatus_name = map[int32]string{
		0: "PENDING",
		1: "COMPLETED",
		2: "FAILED",
	}
	OperationStatus_value = map[string]int32{
		"PENDING":   0,
		"COMPLETED": 1,
		"FAILED":    2,
	}
)

func (x OperationStatus) Enum() *OperationStatus {
	p := new(OperationStatus)
	*p = x
	return p
}

func (x OperationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_relay_processing_v1_models_proto_enumTypes[1].Descriptor()
}

func (OperationStatus) Type() protoreflect.EnumType {
	return &file_relay_processing_v1_models_proto_enumTypes[1]
}

func (x OperationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationStatus.Descriptor instead.
func (OperationStatus) EnumDescriptor() ([]byte, []int) {
	return file_relay_processing_v1_models_proto_rawDescGZIP(), []int{1}
}

type OperationSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Operation      Operation              `protobuf:"varint,2,opt,name=operation,proto3,enum=relay.processing.v1.Operation" json:"operation,omitempty"`
	Amount         int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	LastUpdated    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	Reference      string                 `protobuf:"bytes,5,opt,name=reference,proto3" json:"reference,omitempty"`
	ReconReference string                 `protobuf:"bytes,6,opt,name=recon_reference,json=reconReference,proto3" json:"recon_reference,omitempty"`
	Currency       string                 `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	Status         OperationStatus        `protobuf:"varint,100,opt,name=status,proto3,enum=relay.processing.v1.OperationStatus" json:"status,omitempty"`
	Error          *OperationError        `protobuf:"bytes,101,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *OperationSummary) Reset() {
	*x = OperationSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationSummary) ProtoMessage() {}

func (x *OperationSummary) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationSummary.ProtoReflect.Descriptor instead.
func (*OperationSummary) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_models_proto_rawDescGZIP(), []int{0}
}

func (x *OperationSummary) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OperationSummary) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_AUTHORIZE
}

func (x *OperationSummary) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *OperationSummary) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *OperationSummary) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *OperationSummary) GetReconReference() string {
	if x != nil {
		return x.ReconReference
	}
	return ""
}

func (x *OperationSummary) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *OperationSummary) GetStatus() OperationStatus {
	if x != nil {
		return x.Status
	}
	return OperationStatus_PENDING
}

func (x *OperationSummary) GetError() *OperationError {
	if x != nil {
		return x.Error
	}
	return nil
}

type OperationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OperationError) Reset() {
	*x = OperationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationError) ProtoMessage() {}

func (x *OperationError) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationError.ProtoReflect.Descriptor instead.
func (*OperationError) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_models_proto_rawDescGZIP(), []int{1}
}

func (x *OperationError) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OperationError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: ISO-4217, Minor units of currency, e.g. Amount=100, Currency=EUR -> 1.00 EUR
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// Format: ISO-4217 alpha code, e.g. EUR, USD, GBP
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	// Merchant specific reference, e.g. invoice number
	Reference *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_models_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Order) GetReference() *wrapperspb.StringValue {
	if x != nil {
		return x.Reference
	}
	return nil
}

var File_relay_processing_v1_models_proto protoreflect.FileDescriptor

var file_relay_processing_v1_models_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x03, 0x0a, 0x10, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3e,
	0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x77,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2a, 0x3e, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x45, 0x46, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x2a, 0x39, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x42, 0x4e, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x65, 0x64, 0x76, 0x65, 0x64, 0x37, 0x35, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x74, 0x73,
	0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x13, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_processing_v1_models_proto_rawDescOnce sync.Once
	file_relay_processing_v1_models_proto_rawDescData = file_relay_processing_v1_models_proto_rawDesc
)

func file_relay_processing_v1_models_proto_rawDescGZIP() []byte {
	file_relay_processing_v1_models_proto_rawDescOnce.Do(func() {
		file_relay_processing_v1_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_processing_v1_models_proto_rawDescData)
	})
	return file_relay_processing_v1_models_proto_rawDescData
}

var file_relay_processing_v1_models_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_relay_processing_v1_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_relay_processing_v1_models_proto_goTypes = []interface{}{
	(Operation)(0),                 // 0: relay.processing.v1.Operation
	(OperationStatus)(0),           // 1: relay.processing.v1.OperationStatus
	(*OperationSummary)(nil),       // 2: relay.processing.v1.OperationSummary
	(*OperationError)(nil),         // 3: relay.processing.v1.OperationError
	(*Order)(nil),                  // 4: relay.processing.v1.Order
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 6: google.protobuf.StringValue
}
var file_relay_processing_v1_models_proto_depIdxs = []int32{
	0, // 0: relay.processing.v1.OperationSummary.operation:type_name -> relay.processing.v1.Operation
	5, // 1: relay.processing.v1.OperationSummary.last_updated:type_name -> google.protobuf.Timestamp
	1, // 2: relay.processing.v1.OperationSummary.status:type_name -> relay.processing.v1.OperationStatus
	3, // 3: relay.processing.v1.OperationSummary.error:type_name -> relay.processing.v1.OperationError
	6, // 4: relay.processing.v1.Order.reference:type_name -> google.protobuf.StringValue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_relay_processing_v1_models_proto_init() }
func file_relay_processing_v1_models_proto_init() {
	if File_relay_processing_v1_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_processing_v1_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationSummary); i {
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
		file_relay_processing_v1_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationError); i {
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
		file_relay_processing_v1_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
			RawDescriptor: file_relay_processing_v1_models_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relay_processing_v1_models_proto_goTypes,
		DependencyIndexes: file_relay_processing_v1_models_proto_depIdxs,
		EnumInfos:         file_relay_processing_v1_models_proto_enumTypes,
		MessageInfos:      file_relay_processing_v1_models_proto_msgTypes,
	}.Build()
	File_relay_processing_v1_models_proto = out.File
	file_relay_processing_v1_models_proto_rawDesc = nil
	file_relay_processing_v1_models_proto_goTypes = nil
	file_relay_processing_v1_models_proto_depIdxs = nil
}
