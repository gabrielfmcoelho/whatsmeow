// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: waAdv/WAAdv.proto

package waAdv

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

type ADVEncryptionType int32

const (
	ADVEncryptionType_E2EE   ADVEncryptionType = 0
	ADVEncryptionType_HOSTED ADVEncryptionType = 1
)

// Enum value maps for ADVEncryptionType.
var (
	ADVEncryptionType_name = map[int32]string{
		0: "E2EE",
		1: "HOSTED",
	}
	ADVEncryptionType_value = map[string]int32{
		"E2EE":   0,
		"HOSTED": 1,
	}
)

func (x ADVEncryptionType) Enum() *ADVEncryptionType {
	p := new(ADVEncryptionType)
	*p = x
	return p
}

func (x ADVEncryptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ADVEncryptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_waAdv_WAAdv_proto_enumTypes[0].Descriptor()
}

func (ADVEncryptionType) Type() protoreflect.EnumType {
	return &file_waAdv_WAAdv_proto_enumTypes[0]
}

func (x ADVEncryptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ADVEncryptionType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ADVEncryptionType(num)
	return nil
}

// Deprecated: Use ADVEncryptionType.Descriptor instead.
func (ADVEncryptionType) EnumDescriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{0}
}

type ADVKeyIndexList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawID        *uint32            `protobuf:"varint,1,opt,name=rawID" json:"rawID,omitempty"`
	Timestamp    *uint64            `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	CurrentIndex *uint32            `protobuf:"varint,3,opt,name=currentIndex" json:"currentIndex,omitempty"`
	ValidIndexes []uint32           `protobuf:"varint,4,rep,packed,name=validIndexes" json:"validIndexes,omitempty"`
	AccountType  *ADVEncryptionType `protobuf:"varint,5,opt,name=accountType,enum=WAAdv.ADVEncryptionType" json:"accountType,omitempty"`
}

func (x *ADVKeyIndexList) Reset() {
	*x = ADVKeyIndexList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVKeyIndexList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVKeyIndexList) ProtoMessage() {}

func (x *ADVKeyIndexList) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVKeyIndexList.ProtoReflect.Descriptor instead.
func (*ADVKeyIndexList) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{0}
}

func (x *ADVKeyIndexList) GetRawID() uint32 {
	if x != nil && x.RawID != nil {
		return *x.RawID
	}
	return 0
}

func (x *ADVKeyIndexList) GetTimestamp() uint64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *ADVKeyIndexList) GetCurrentIndex() uint32 {
	if x != nil && x.CurrentIndex != nil {
		return *x.CurrentIndex
	}
	return 0
}

func (x *ADVKeyIndexList) GetValidIndexes() []uint32 {
	if x != nil {
		return x.ValidIndexes
	}
	return nil
}

func (x *ADVKeyIndexList) GetAccountType() ADVEncryptionType {
	if x != nil && x.AccountType != nil {
		return *x.AccountType
	}
	return ADVEncryptionType_E2EE
}

type ADVSignedKeyIndexList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details             []byte `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	AccountSignature    []byte `protobuf:"bytes,2,opt,name=accountSignature" json:"accountSignature,omitempty"`
	AccountSignatureKey []byte `protobuf:"bytes,3,opt,name=accountSignatureKey" json:"accountSignatureKey,omitempty"`
}

func (x *ADVSignedKeyIndexList) Reset() {
	*x = ADVSignedKeyIndexList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVSignedKeyIndexList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVSignedKeyIndexList) ProtoMessage() {}

func (x *ADVSignedKeyIndexList) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVSignedKeyIndexList.ProtoReflect.Descriptor instead.
func (*ADVSignedKeyIndexList) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{1}
}

func (x *ADVSignedKeyIndexList) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ADVSignedKeyIndexList) GetAccountSignature() []byte {
	if x != nil {
		return x.AccountSignature
	}
	return nil
}

func (x *ADVSignedKeyIndexList) GetAccountSignatureKey() []byte {
	if x != nil {
		return x.AccountSignatureKey
	}
	return nil
}

type ADVDeviceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawID       *uint32            `protobuf:"varint,1,opt,name=rawID" json:"rawID,omitempty"`
	Timestamp   *uint64            `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	KeyIndex    *uint32            `protobuf:"varint,3,opt,name=keyIndex" json:"keyIndex,omitempty"`
	AccountType *ADVEncryptionType `protobuf:"varint,4,opt,name=accountType,enum=WAAdv.ADVEncryptionType" json:"accountType,omitempty"`
	DeviceType  *ADVEncryptionType `protobuf:"varint,5,opt,name=deviceType,enum=WAAdv.ADVEncryptionType" json:"deviceType,omitempty"`
}

func (x *ADVDeviceIdentity) Reset() {
	*x = ADVDeviceIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVDeviceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVDeviceIdentity) ProtoMessage() {}

func (x *ADVDeviceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVDeviceIdentity.ProtoReflect.Descriptor instead.
func (*ADVDeviceIdentity) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{2}
}

func (x *ADVDeviceIdentity) GetRawID() uint32 {
	if x != nil && x.RawID != nil {
		return *x.RawID
	}
	return 0
}

func (x *ADVDeviceIdentity) GetTimestamp() uint64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *ADVDeviceIdentity) GetKeyIndex() uint32 {
	if x != nil && x.KeyIndex != nil {
		return *x.KeyIndex
	}
	return 0
}

func (x *ADVDeviceIdentity) GetAccountType() ADVEncryptionType {
	if x != nil && x.AccountType != nil {
		return *x.AccountType
	}
	return ADVEncryptionType_E2EE
}

func (x *ADVDeviceIdentity) GetDeviceType() ADVEncryptionType {
	if x != nil && x.DeviceType != nil {
		return *x.DeviceType
	}
	return ADVEncryptionType_E2EE
}

type ADVSignedDeviceIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details             []byte `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	AccountSignatureKey []byte `protobuf:"bytes,2,opt,name=accountSignatureKey" json:"accountSignatureKey,omitempty"`
	AccountSignature    []byte `protobuf:"bytes,3,opt,name=accountSignature" json:"accountSignature,omitempty"`
	DeviceSignature     []byte `protobuf:"bytes,4,opt,name=deviceSignature" json:"deviceSignature,omitempty"`
}

func (x *ADVSignedDeviceIdentity) Reset() {
	*x = ADVSignedDeviceIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVSignedDeviceIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVSignedDeviceIdentity) ProtoMessage() {}

func (x *ADVSignedDeviceIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVSignedDeviceIdentity.ProtoReflect.Descriptor instead.
func (*ADVSignedDeviceIdentity) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{3}
}

func (x *ADVSignedDeviceIdentity) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ADVSignedDeviceIdentity) GetAccountSignatureKey() []byte {
	if x != nil {
		return x.AccountSignatureKey
	}
	return nil
}

func (x *ADVSignedDeviceIdentity) GetAccountSignature() []byte {
	if x != nil {
		return x.AccountSignature
	}
	return nil
}

func (x *ADVSignedDeviceIdentity) GetDeviceSignature() []byte {
	if x != nil {
		return x.DeviceSignature
	}
	return nil
}

type ADVSignedDeviceIdentityHMAC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details     []byte             `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	HMAC        []byte             `protobuf:"bytes,2,opt,name=HMAC" json:"HMAC,omitempty"`
	AccountType *ADVEncryptionType `protobuf:"varint,3,opt,name=accountType,enum=WAAdv.ADVEncryptionType" json:"accountType,omitempty"`
}

func (x *ADVSignedDeviceIdentityHMAC) Reset() {
	*x = ADVSignedDeviceIdentityHMAC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waAdv_WAAdv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ADVSignedDeviceIdentityHMAC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ADVSignedDeviceIdentityHMAC) ProtoMessage() {}

func (x *ADVSignedDeviceIdentityHMAC) ProtoReflect() protoreflect.Message {
	mi := &file_waAdv_WAAdv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ADVSignedDeviceIdentityHMAC.ProtoReflect.Descriptor instead.
func (*ADVSignedDeviceIdentityHMAC) Descriptor() ([]byte, []int) {
	return file_waAdv_WAAdv_proto_rawDescGZIP(), []int{4}
}

func (x *ADVSignedDeviceIdentityHMAC) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ADVSignedDeviceIdentityHMAC) GetHMAC() []byte {
	if x != nil {
		return x.HMAC
	}
	return nil
}

func (x *ADVSignedDeviceIdentityHMAC) GetAccountType() ADVEncryptionType {
	if x != nil && x.AccountType != nil {
		return *x.AccountType
	}
	return ADVEncryptionType_E2EE
}

var File_waAdv_WAAdv_proto protoreflect.FileDescriptor

var file_waAdv_WAAdv_proto_rawDesc = []byte{
	0x0a, 0x11, 0x77, 0x61, 0x41, 0x64, 0x76, 0x2f, 0x57, 0x41, 0x41, 0x64, 0x76, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x57, 0x41, 0x41, 0x64, 0x76, 0x22, 0xcd, 0x01, 0x0a, 0x0f, 0x41,
	0x44, 0x56, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x61, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72,
	0x61, 0x77, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01,
	0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x57, 0x41, 0x41, 0x64, 0x76, 0x2e, 0x41, 0x44, 0x56, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x41,
	0x44, 0x56, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2a,
	0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x22, 0xd9, 0x01, 0x0a,
	0x11, 0x41, 0x44, 0x56, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x77, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x72, 0x61, 0x77, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x57, 0x41, 0x41, 0x64, 0x76, 0x2e,
	0x41, 0x44, 0x56, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38,
	0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x57, 0x41, 0x41, 0x64, 0x76, 0x2e, 0x41, 0x44, 0x56, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x17, 0x41, 0x44, 0x56,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x30,
	0x0a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x2a, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x1b, 0x41, 0x44, 0x56, 0x53, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x48, 0x4d, 0x41, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x48, 0x4d, 0x41, 0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x48, 0x4d, 0x41, 0x43, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x57, 0x41, 0x41, 0x64,
	0x76, 0x2e, 0x41, 0x44, 0x56, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x2a, 0x29, 0x0a, 0x11, 0x41, 0x44, 0x56, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x32, 0x45, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x48, 0x4f, 0x53, 0x54, 0x45, 0x44, 0x10, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x62, 0x72, 0x69, 0x65,
	0x6c, 0x66, 0x6d, 0x63, 0x6f, 0x65, 0x6c, 0x68, 0x6f, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x73, 0x6d,
	0x65, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x41, 0x64, 0x76,
}

var (
	file_waAdv_WAAdv_proto_rawDescOnce sync.Once
	file_waAdv_WAAdv_proto_rawDescData = file_waAdv_WAAdv_proto_rawDesc
)

func file_waAdv_WAAdv_proto_rawDescGZIP() []byte {
	file_waAdv_WAAdv_proto_rawDescOnce.Do(func() {
		file_waAdv_WAAdv_proto_rawDescData = protoimpl.X.CompressGZIP(file_waAdv_WAAdv_proto_rawDescData)
	})
	return file_waAdv_WAAdv_proto_rawDescData
}

var file_waAdv_WAAdv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_waAdv_WAAdv_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_waAdv_WAAdv_proto_goTypes = []any{
	(ADVEncryptionType)(0),              // 0: WAAdv.ADVEncryptionType
	(*ADVKeyIndexList)(nil),             // 1: WAAdv.ADVKeyIndexList
	(*ADVSignedKeyIndexList)(nil),       // 2: WAAdv.ADVSignedKeyIndexList
	(*ADVDeviceIdentity)(nil),           // 3: WAAdv.ADVDeviceIdentity
	(*ADVSignedDeviceIdentity)(nil),     // 4: WAAdv.ADVSignedDeviceIdentity
	(*ADVSignedDeviceIdentityHMAC)(nil), // 5: WAAdv.ADVSignedDeviceIdentityHMAC
}
var file_waAdv_WAAdv_proto_depIdxs = []int32{
	0, // 0: WAAdv.ADVKeyIndexList.accountType:type_name -> WAAdv.ADVEncryptionType
	0, // 1: WAAdv.ADVDeviceIdentity.accountType:type_name -> WAAdv.ADVEncryptionType
	0, // 2: WAAdv.ADVDeviceIdentity.deviceType:type_name -> WAAdv.ADVEncryptionType
	0, // 3: WAAdv.ADVSignedDeviceIdentityHMAC.accountType:type_name -> WAAdv.ADVEncryptionType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_waAdv_WAAdv_proto_init() }
func file_waAdv_WAAdv_proto_init() {
	if File_waAdv_WAAdv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waAdv_WAAdv_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ADVKeyIndexList); i {
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
		file_waAdv_WAAdv_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ADVSignedKeyIndexList); i {
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
		file_waAdv_WAAdv_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ADVDeviceIdentity); i {
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
		file_waAdv_WAAdv_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ADVSignedDeviceIdentity); i {
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
		file_waAdv_WAAdv_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ADVSignedDeviceIdentityHMAC); i {
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
			RawDescriptor: file_waAdv_WAAdv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waAdv_WAAdv_proto_goTypes,
		DependencyIndexes: file_waAdv_WAAdv_proto_depIdxs,
		EnumInfos:         file_waAdv_WAAdv_proto_enumTypes,
		MessageInfos:      file_waAdv_WAAdv_proto_msgTypes,
	}.Build()
	File_waAdv_WAAdv_proto = out.File
	file_waAdv_WAAdv_proto_rawDesc = nil
	file_waAdv_WAAdv_proto_goTypes = nil
	file_waAdv_WAAdv_proto_depIdxs = nil
}
