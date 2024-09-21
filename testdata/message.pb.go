// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.23.3
// source: testdata/message.proto

package testdata

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

type Dog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Dog) Reset() {
	*x = Dog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dog) ProtoMessage() {}

func (x *Dog) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dog.ProtoReflect.Descriptor instead.
func (*Dog) Descriptor() ([]byte, []int) {
	return file_testdata_message_proto_rawDescGZIP(), []int{0}
}

func (x *Dog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dog) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Cat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind  string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Birth int32  `protobuf:"varint,3,opt,name=birth,proto3" json:"birth,omitempty"`
}

func (x *Cat) Reset() {
	*x = Cat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cat) ProtoMessage() {}

func (x *Cat) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cat.ProtoReflect.Descriptor instead.
func (*Cat) Descriptor() ([]byte, []int) {
	return file_testdata_message_proto_rawDescGZIP(), []int{1}
}

func (x *Cat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cat) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Cat) GetBirth() int32 {
	if x != nil {
		return x.Birth
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"column:id;autoIncrement" bson:"_id"`
	Sequence   int64       `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty" gorm:"column:sequence;type:bigint"`
	Addresser  int64       `protobuf:"varint,3,opt,name=addresser,proto3" json:"addresser,omitempty" gorm:"column:addresser;type:bigint" bson:"Addresser"`
	Addressee  int64       `protobuf:"varint,4,opt,name=addressee,proto3" json:"addressee,omitempty" gorm:"column:addressee;type:bigint"`
	Originator int64       `protobuf:"varint,5,opt,name=originator,proto3" json:"originator,omitempty"`
	Timestamp  int64       `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Flags      int32       `protobuf:"varint,7,opt,name=flags,proto3" json:"flags,omitempty"`
	Platform   int32       `protobuf:"varint,8,opt,name=platform,proto3" json:"platform,omitempty"`
	Category   int32       `protobuf:"varint,9,opt,name=category,proto3" json:"category,omitempty"`
	Type       int32       `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`
	S1         string      `protobuf:"bytes,11,opt,name=s1,proto3" json:"s1,omitempty"`
	S2         string      `protobuf:"bytes,12,opt,name=s2,proto3" json:"s2,omitempty"`
	S3         string      `protobuf:"bytes,13,opt,name=s3,proto3" json:"s3,omitempty"`
	S4         string      `protobuf:"bytes,14,opt,name=s4,proto3" json:"s4,omitempty"`
	N1         int64       `protobuf:"varint,15,opt,name=n1,proto3" json:"n1,omitempty"`
	N2         int64       `protobuf:"varint,16,opt,name=n2,proto3" json:"n2,omitempty"`
	N3         int64       `protobuf:"varint,17,opt,name=n3,proto3" json:"n3,omitempty"`
	N4         int64       `protobuf:"varint,18,opt,name=n4,proto3" json:"n4,omitempty"`
	Quote      []byte      `protobuf:"bytes,19,opt,name=quote,proto3" json:"quote,omitempty"`
	Extra      []byte      `protobuf:"bytes,20,opt,name=extra,proto3" json:"extra,omitempty"`
	From       *MessageMan `protobuf:"bytes,21,opt,name=from,proto3" json:"from,omitempty"`
	// @go.name=MyPet
	// @bson.tag=my_pet
	//
	// Types that are assignable to Pet:
	//
	//	*Message_Dog
	//	*Message_Cat
	Pet isMessage_Pet `protobuf_oneof:"pet"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_testdata_message_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetSequence() int64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Message) GetAddresser() int64 {
	if x != nil {
		return x.Addresser
	}
	return 0
}

func (x *Message) GetAddressee() int64 {
	if x != nil {
		return x.Addressee
	}
	return 0
}

func (x *Message) GetOriginator() int64 {
	if x != nil {
		return x.Originator
	}
	return 0
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *Message) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *Message) GetCategory() int32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *Message) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Message) GetS1() string {
	if x != nil {
		return x.S1
	}
	return ""
}

func (x *Message) GetS2() string {
	if x != nil {
		return x.S2
	}
	return ""
}

func (x *Message) GetS3() string {
	if x != nil {
		return x.S3
	}
	return ""
}

func (x *Message) GetS4() string {
	if x != nil {
		return x.S4
	}
	return ""
}

func (x *Message) GetN1() int64 {
	if x != nil {
		return x.N1
	}
	return 0
}

func (x *Message) GetN2() int64 {
	if x != nil {
		return x.N2
	}
	return 0
}

func (x *Message) GetN3() int64 {
	if x != nil {
		return x.N3
	}
	return 0
}

func (x *Message) GetN4() int64 {
	if x != nil {
		return x.N4
	}
	return 0
}

func (x *Message) GetQuote() []byte {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *Message) GetExtra() []byte {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *Message) GetFrom() *MessageMan {
	if x != nil {
		return x.From
	}
	return nil
}

func (m *Message) GetPet() isMessage_Pet {
	if m != nil {
		return m.Pet
	}
	return nil
}

func (x *Message) GetDog() *Dog {
	if x, ok := x.GetPet().(*Message_Dog); ok {
		return x.Dog
	}
	return nil
}

func (x *Message) GetCat() *Cat {
	if x, ok := x.GetPet().(*Message_Cat); ok {
		return x.Cat
	}
	return nil
}

type isMessage_Pet interface {
	isMessage_Pet()
}

type Message_Dog struct {
	Dog *Dog `protobuf:"bytes,22,opt,name=dog,proto3,oneof"`
}

type Message_Cat struct {
	Cat *Cat `protobuf:"bytes,23,opt,name=cat,proto3,oneof"`
}

func (*Message_Dog) isMessage_Pet() {}

func (*Message_Cat) isMessage_Pet() {}

type MessageMan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UName string `protobuf:"bytes,1,opt,name=name,proto3" json:"user_name,omitempty"`
	UA    int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *MessageMan) Reset() {
	*x = MessageMan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageMan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageMan) ProtoMessage() {}

func (x *MessageMan) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageMan.ProtoReflect.Descriptor instead.
func (*MessageMan) Descriptor() ([]byte, []int) {
	return file_testdata_message_proto_rawDescGZIP(), []int{2, 0}
}

func (x *MessageMan) GetUName() string {
	if x != nil {
		return x.UName
	}
	return ""
}

func (x *MessageMan) GetUA() int32 {
	if x != nil {
		return x.UA
	}
	return 0
}

var File_testdata_message_proto protoreflect.FileDescriptor

var file_testdata_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x2b, 0x0a, 0x03, 0x64, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22,
	0x43, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x22, 0xe2, 0x04, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x31, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x32, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x34, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x31, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6e, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x32, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6e, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x33, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6e, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x34, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6e, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x12, 0x29, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x21, 0x0a, 0x03, 0x64, 0x6f, 0x67, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x03, 0x64,
	0x6f, 0x67, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x61, 0x74, 0x48, 0x00,
	0x52, 0x03, 0x63, 0x61, 0x74, 0x1a, 0x2b, 0x0a, 0x03, 0x6d, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x70, 0x65, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testdata_message_proto_rawDescOnce sync.Once
	file_testdata_message_proto_rawDescData = file_testdata_message_proto_rawDesc
)

func file_testdata_message_proto_rawDescGZIP() []byte {
	file_testdata_message_proto_rawDescOnce.Do(func() {
		file_testdata_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_message_proto_rawDescData)
	})
	return file_testdata_message_proto_rawDescData
}

var file_testdata_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_testdata_message_proto_goTypes = []any{
	(*Dog)(nil),        // 0: testdata.dog
	(*Cat)(nil),        // 1: testdata.cat
	(*Message)(nil),    // 2: testdata.message
	(*MessageMan)(nil), // 3: testdata.message.man
}
var file_testdata_message_proto_depIdxs = []int32{
	3, // 0: testdata.message.from:type_name -> testdata.message.man
	0, // 1: testdata.message.dog:type_name -> testdata.dog
	1, // 2: testdata.message.cat:type_name -> testdata.cat
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_testdata_message_proto_init() }
func file_testdata_message_proto_init() {
	if File_testdata_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Dog); i {
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
		file_testdata_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Cat); i {
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
		file_testdata_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
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
		file_testdata_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MessageMan); i {
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
	file_testdata_message_proto_msgTypes[2].OneofWrappers = []any{
		(*Message_Dog)(nil),
		(*Message_Cat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testdata_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_message_proto_goTypes,
		DependencyIndexes: file_testdata_message_proto_depIdxs,
		MessageInfos:      file_testdata_message_proto_msgTypes,
	}.Build()
	File_testdata_message_proto = out.File
	file_testdata_message_proto_rawDesc = nil
	file_testdata_message_proto_goTypes = nil
	file_testdata_message_proto_depIdxs = nil
}
