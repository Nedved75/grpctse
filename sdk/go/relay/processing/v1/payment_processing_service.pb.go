// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: relay/processing/v1/payment_processing_service.proto

package v1

import (
	relay "github.com/Nedved75/grpctse/sdk/go/relay"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChargeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentMethodConfigurationId string              `protobuf:"bytes,1,opt,name=payment_method_configuration_id,json=paymentMethodConfigurationId,proto3" json:"payment_method_configuration_id,omitempty"`
	RequestId                    string              `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Reference                    string              `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	PaymentMethod                relay.PaymentMethod `protobuf:"varint,4,opt,name=payment_method,json=paymentMethod,proto3,enum=relay.PaymentMethod" json:"payment_method,omitempty"`
	Order                        *Order              `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
	// Payment method specific information
	Data *anypb.Any `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ChargeRequest) Reset() {
	*x = ChargeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeRequest) ProtoMessage() {}

func (x *ChargeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeRequest.ProtoReflect.Descriptor instead.
func (*ChargeRequest) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChargeRequest) GetPaymentMethodConfigurationId() string {
	if x != nil {
		return x.PaymentMethodConfigurationId
	}
	return ""
}

func (x *ChargeRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ChargeRequest) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *ChargeRequest) GetPaymentMethod() relay.PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return relay.PaymentMethod(0)
}

func (x *ChargeRequest) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *ChargeRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type ChargeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	ChargeId  string `protobuf:"bytes,2,opt,name=charge_id,json=chargeId,proto3" json:"charge_id,omitempty"`
	// Payment method specific information
	Data   *anypb.Any    `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	Status *relay.Status `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ChargeResponse) Reset() {
	*x = ChargeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChargeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChargeResponse) ProtoMessage() {}

func (x *ChargeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChargeResponse.ProtoReflect.Descriptor instead.
func (*ChargeResponse) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{1}
}

func (x *ChargeResponse) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *ChargeResponse) GetChargeId() string {
	if x != nil {
		return x.ChargeId
	}
	return ""
}

func (x *ChargeResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChargeResponse) GetStatus() *relay.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type RefundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Reference string `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	Amount    int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// Payment method specific information
	Data *anypb.Any `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{2}
}

func (x *RefundRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *RefundRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *RefundRequest) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *RefundRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RefundRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type RefundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefundId string `protobuf:"bytes,1,opt,name=refund_id,json=refundId,proto3" json:"refund_id,omitempty"`
	// Payment method specific information
	Data   *anypb.Any    `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	Status *relay.Status `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RefundResponse) Reset() {
	*x = RefundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundResponse) ProtoMessage() {}

func (x *RefundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundResponse.ProtoReflect.Descriptor instead.
func (*RefundResponse) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{3}
}

func (x *RefundResponse) GetRefundId() string {
	if x != nil {
		return x.RefundId
	}
	return ""
}

func (x *RefundResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RefundResponse) GetStatus() *relay.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *GetSummaryRequest) Reset() {
	*x = GetSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryRequest) ProtoMessage() {}

func (x *GetSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetSummaryRequest) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetSummaryRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type GetSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorizedAmount int64         `protobuf:"varint,1,opt,name=authorized_amount,json=authorizedAmount,proto3" json:"authorized_amount,omitempty"`
	ChargedAmount    int64         `protobuf:"varint,2,opt,name=charged_amount,json=chargedAmount,proto3" json:"charged_amount,omitempty"`
	CanceledAmount   int64         `protobuf:"varint,3,opt,name=canceled_amount,json=canceledAmount,proto3" json:"canceled_amount,omitempty"`
	RefundedAmount   int64         `protobuf:"varint,4,opt,name=refunded_amount,json=refundedAmount,proto3" json:"refunded_amount,omitempty"`
	Currency         string        `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Status           *relay.Status `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetSummaryResponse) Reset() {
	*x = GetSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSummaryResponse) ProtoMessage() {}

func (x *GetSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSummaryResponse.ProtoReflect.Descriptor instead.
func (*GetSummaryResponse) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetSummaryResponse) GetAuthorizedAmount() int64 {
	if x != nil {
		return x.AuthorizedAmount
	}
	return 0
}

func (x *GetSummaryResponse) GetChargedAmount() int64 {
	if x != nil {
		return x.ChargedAmount
	}
	return 0
}

func (x *GetSummaryResponse) GetCanceledAmount() int64 {
	if x != nil {
		return x.CanceledAmount
	}
	return 0
}

func (x *GetSummaryResponse) GetRefundedAmount() int64 {
	if x != nil {
		return x.RefundedAmount
	}
	return 0
}

func (x *GetSummaryResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetSummaryResponse) GetStatus() *relay.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type GetOperationsSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentId string `protobuf:"bytes,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *GetOperationsSummaryRequest) Reset() {
	*x = GetOperationsSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOperationsSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationsSummaryRequest) ProtoMessage() {}

func (x *GetOperationsSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationsSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetOperationsSummaryRequest) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetOperationsSummaryRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type GetOperationsSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operations []*OperationSummary `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	Status     *relay.Status       `protobuf:"bytes,101,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetOperationsSummaryResponse) Reset() {
	*x = GetOperationsSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOperationsSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOperationsSummaryResponse) ProtoMessage() {}

func (x *GetOperationsSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_processing_v1_payment_processing_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOperationsSummaryResponse.ProtoReflect.Descriptor instead.
func (*GetOperationsSummaryResponse) Descriptor() ([]byte, []int) {
	return file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetOperationsSummaryResponse) GetOperations() []*OperationSummary {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *GetOperationsSummaryResponse) GetStatus() *relay.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_relay_processing_v1_payment_processing_service_proto protoreflect.FileDescriptor

var file_relay_processing_v1_payment_processing_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a,
	0x0d, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45,
	0x0a, 0x1f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x30, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9d, 0x01, 0x0a, 0x0e,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x0d,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7e, 0x0a, 0x0e, 0x52,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x32, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0xfd, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x3c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x8c, 0x01,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9d, 0x03, 0x0a,
	0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x12, 0x53, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x30, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x64, 0x76, 0x65,
	0x64, 0x37, 0x35, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x74, 0x73, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x67, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0xaa, 0x02, 0x13, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_processing_v1_payment_processing_service_proto_rawDescOnce sync.Once
	file_relay_processing_v1_payment_processing_service_proto_rawDescData = file_relay_processing_v1_payment_processing_service_proto_rawDesc
)

func file_relay_processing_v1_payment_processing_service_proto_rawDescGZIP() []byte {
	file_relay_processing_v1_payment_processing_service_proto_rawDescOnce.Do(func() {
		file_relay_processing_v1_payment_processing_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_processing_v1_payment_processing_service_proto_rawDescData)
	})
	return file_relay_processing_v1_payment_processing_service_proto_rawDescData
}

var file_relay_processing_v1_payment_processing_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_relay_processing_v1_payment_processing_service_proto_goTypes = []interface{}{
	(*ChargeRequest)(nil),                // 0: relay.processing.v1.ChargeRequest
	(*ChargeResponse)(nil),               // 1: relay.processing.v1.ChargeResponse
	(*RefundRequest)(nil),                // 2: relay.processing.v1.RefundRequest
	(*RefundResponse)(nil),               // 3: relay.processing.v1.RefundResponse
	(*GetSummaryRequest)(nil),            // 4: relay.processing.v1.GetSummaryRequest
	(*GetSummaryResponse)(nil),           // 5: relay.processing.v1.GetSummaryResponse
	(*GetOperationsSummaryRequest)(nil),  // 6: relay.processing.v1.GetOperationsSummaryRequest
	(*GetOperationsSummaryResponse)(nil), // 7: relay.processing.v1.GetOperationsSummaryResponse
	(relay.PaymentMethod)(0),             // 8: relay.PaymentMethod
	(*Order)(nil),                        // 9: relay.processing.v1.Order
	(*anypb.Any)(nil),                    // 10: google.protobuf.Any
	(*relay.Status)(nil),                 // 11: relay.Status
	(*OperationSummary)(nil),             // 12: relay.processing.v1.OperationSummary
}
var file_relay_processing_v1_payment_processing_service_proto_depIdxs = []int32{
	8,  // 0: relay.processing.v1.ChargeRequest.payment_method:type_name -> relay.PaymentMethod
	9,  // 1: relay.processing.v1.ChargeRequest.order:type_name -> relay.processing.v1.Order
	10, // 2: relay.processing.v1.ChargeRequest.data:type_name -> google.protobuf.Any
	10, // 3: relay.processing.v1.ChargeResponse.data:type_name -> google.protobuf.Any
	11, // 4: relay.processing.v1.ChargeResponse.status:type_name -> relay.Status
	10, // 5: relay.processing.v1.RefundRequest.data:type_name -> google.protobuf.Any
	10, // 6: relay.processing.v1.RefundResponse.data:type_name -> google.protobuf.Any
	11, // 7: relay.processing.v1.RefundResponse.status:type_name -> relay.Status
	11, // 8: relay.processing.v1.GetSummaryResponse.status:type_name -> relay.Status
	12, // 9: relay.processing.v1.GetOperationsSummaryResponse.operations:type_name -> relay.processing.v1.OperationSummary
	11, // 10: relay.processing.v1.GetOperationsSummaryResponse.status:type_name -> relay.Status
	0,  // 11: relay.processing.v1.PaymentProcessing.Charge:input_type -> relay.processing.v1.ChargeRequest
	2,  // 12: relay.processing.v1.PaymentProcessing.Refund:input_type -> relay.processing.v1.RefundRequest
	4,  // 13: relay.processing.v1.PaymentProcessing.GetSummary:input_type -> relay.processing.v1.GetSummaryRequest
	6,  // 14: relay.processing.v1.PaymentProcessing.GetOperationsSummary:input_type -> relay.processing.v1.GetOperationsSummaryRequest
	1,  // 15: relay.processing.v1.PaymentProcessing.Charge:output_type -> relay.processing.v1.ChargeResponse
	3,  // 16: relay.processing.v1.PaymentProcessing.Refund:output_type -> relay.processing.v1.RefundResponse
	5,  // 17: relay.processing.v1.PaymentProcessing.GetSummary:output_type -> relay.processing.v1.GetSummaryResponse
	7,  // 18: relay.processing.v1.PaymentProcessing.GetOperationsSummary:output_type -> relay.processing.v1.GetOperationsSummaryResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_relay_processing_v1_payment_processing_service_proto_init() }
func file_relay_processing_v1_payment_processing_service_proto_init() {
	if File_relay_processing_v1_payment_processing_service_proto != nil {
		return
	}
	file_relay_processing_v1_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeRequest); i {
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
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChargeResponse); i {
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
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundRequest); i {
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
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefundResponse); i {
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
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryRequest); i {
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
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSummaryResponse); i {
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
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOperationsSummaryRequest); i {
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
		file_relay_processing_v1_payment_processing_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOperationsSummaryResponse); i {
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
			RawDescriptor: file_relay_processing_v1_payment_processing_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relay_processing_v1_payment_processing_service_proto_goTypes,
		DependencyIndexes: file_relay_processing_v1_payment_processing_service_proto_depIdxs,
		MessageInfos:      file_relay_processing_v1_payment_processing_service_proto_msgTypes,
	}.Build()
	File_relay_processing_v1_payment_processing_service_proto = out.File
	file_relay_processing_v1_payment_processing_service_proto_rawDesc = nil
	file_relay_processing_v1_payment_processing_service_proto_goTypes = nil
	file_relay_processing_v1_payment_processing_service_proto_depIdxs = nil
}
