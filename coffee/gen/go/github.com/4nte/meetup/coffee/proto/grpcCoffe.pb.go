// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: grpcCoffe.proto

package proto

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

type LongBlack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 2 shots recommended
	ShotsOfEspresso int32 `protobuf:"varint,1,opt,name=shots_of_espresso,json=shotsOfEspresso,proto3" json:"shots_of_espresso,omitempty"`
	// 90 ml recommended
	HowWaterMl int32 `protobuf:"varint,2,opt,name=how_water_ml,json=howWaterMl,proto3" json:"how_water_ml,omitempty"`
}

func (x *LongBlack) Reset() {
	*x = LongBlack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcCoffe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongBlack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongBlack) ProtoMessage() {}

func (x *LongBlack) ProtoReflect() protoreflect.Message {
	mi := &file_grpcCoffe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongBlack.ProtoReflect.Descriptor instead.
func (*LongBlack) Descriptor() ([]byte, []int) {
	return file_grpcCoffe_proto_rawDescGZIP(), []int{0}
}

func (x *LongBlack) GetShotsOfEspresso() int32 {
	if x != nil {
		return x.ShotsOfEspresso
	}
	return 0
}

func (x *LongBlack) GetHowWaterMl() int32 {
	if x != nil {
		return x.HowWaterMl
	}
	return 0
}

// Caffè macchiato, sometimes called espresso macchiato, is an espresso coffee drink with a small amount of milk, usually foamed.
type Macchiato struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1 shot recommended
	ShotsOfEspresso int32 `protobuf:"varint,1,opt,name=shots_of_espresso,json=shotsOfEspresso,proto3" json:"shots_of_espresso,omitempty"`
	// 1-2 ml recommended
	SteamedMilkMl int32 `protobuf:"varint,2,opt,name=steamed_milk_ml,json=steamedMilkMl,proto3" json:"steamed_milk_ml,omitempty"`
}

func (x *Macchiato) Reset() {
	*x = Macchiato{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcCoffe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Macchiato) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Macchiato) ProtoMessage() {}

func (x *Macchiato) ProtoReflect() protoreflect.Message {
	mi := &file_grpcCoffe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Macchiato.ProtoReflect.Descriptor instead.
func (*Macchiato) Descriptor() ([]byte, []int) {
	return file_grpcCoffe_proto_rawDescGZIP(), []int{1}
}

func (x *Macchiato) GetShotsOfEspresso() int32 {
	if x != nil {
		return x.ShotsOfEspresso
	}
	return 0
}

func (x *Macchiato) GetSteamedMilkMl() int32 {
	if x != nil {
		return x.SteamedMilkMl
	}
	return 0
}

// Traditionally a cappuccino is viewed as a drink that should only be consumed in the morning
type Cappuccino struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 1-2 shots of espresso
	ShotsOfEspresso int32 `protobuf:"varint,1,opt,name=shots_of_espresso,json=shotsOfEspresso,proto3" json:"shots_of_espresso,omitempty"`
	// 60 ml recommended
	SteamedMilkMl int32 `protobuf:"varint,2,opt,name=steamed_milk_ml,json=steamedMilkMl,proto3" json:"steamed_milk_ml,omitempty"`
	// 60 ml recommended
	FoamedMilkMl int32 `protobuf:"varint,3,opt,name=foamed_milk_ml,json=foamedMilkMl,proto3" json:"foamed_milk_ml,omitempty"`
	// optional sprinkling of chocolate powder
	SprinklinkgOfChocolatePoweder bool `protobuf:"varint,4,opt,name=sprinklinkg_of_chocolate_poweder,json=sprinklinkgOfChocolatePoweder,proto3" json:"sprinklinkg_of_chocolate_poweder,omitempty"`
}

func (x *Cappuccino) Reset() {
	*x = Cappuccino{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcCoffe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cappuccino) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cappuccino) ProtoMessage() {}

func (x *Cappuccino) ProtoReflect() protoreflect.Message {
	mi := &file_grpcCoffe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cappuccino.ProtoReflect.Descriptor instead.
func (*Cappuccino) Descriptor() ([]byte, []int) {
	return file_grpcCoffe_proto_rawDescGZIP(), []int{2}
}

func (x *Cappuccino) GetShotsOfEspresso() int32 {
	if x != nil {
		return x.ShotsOfEspresso
	}
	return 0
}

func (x *Cappuccino) GetSteamedMilkMl() int32 {
	if x != nil {
		return x.SteamedMilkMl
	}
	return 0
}

func (x *Cappuccino) GetFoamedMilkMl() int32 {
	if x != nil {
		return x.FoamedMilkMl
	}
	return 0
}

func (x *Cappuccino) GetSprinklinkgOfChocolatePoweder() bool {
	if x != nil {
		return x.SprinklinkgOfChocolatePoweder
	}
	return false
}

// instructs the coffe machine to make a coffe
type MakeCoffeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to CoffeRecipe:
	//
	//	*MakeCoffeRequest_Cappuccino
	//	*MakeCoffeRequest_Macchiato
	//	*MakeCoffeRequest_LongBlack
	CoffeRecipe isMakeCoffeRequest_CoffeRecipe `protobuf_oneof:"coffe_recipe"`
	// amount of sugar in grams
	GramsOfSugar int32 `protobuf:"varint,4,opt,name=grams_of_sugar,json=gramsOfSugar,proto3" json:"grams_of_sugar,omitempty"`
	// delays the start of the coffe brewing for specified amount in seconds
	Delay int32 `protobuf:"varint,5,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *MakeCoffeRequest) Reset() {
	*x = MakeCoffeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcCoffe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeCoffeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeCoffeRequest) ProtoMessage() {}

func (x *MakeCoffeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcCoffe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeCoffeRequest.ProtoReflect.Descriptor instead.
func (*MakeCoffeRequest) Descriptor() ([]byte, []int) {
	return file_grpcCoffe_proto_rawDescGZIP(), []int{3}
}

func (m *MakeCoffeRequest) GetCoffeRecipe() isMakeCoffeRequest_CoffeRecipe {
	if m != nil {
		return m.CoffeRecipe
	}
	return nil
}

func (x *MakeCoffeRequest) GetCappuccino() *Cappuccino {
	if x, ok := x.GetCoffeRecipe().(*MakeCoffeRequest_Cappuccino); ok {
		return x.Cappuccino
	}
	return nil
}

func (x *MakeCoffeRequest) GetMacchiato() *Macchiato {
	if x, ok := x.GetCoffeRecipe().(*MakeCoffeRequest_Macchiato); ok {
		return x.Macchiato
	}
	return nil
}

func (x *MakeCoffeRequest) GetLongBlack() *LongBlack {
	if x, ok := x.GetCoffeRecipe().(*MakeCoffeRequest_LongBlack); ok {
		return x.LongBlack
	}
	return nil
}

func (x *MakeCoffeRequest) GetGramsOfSugar() int32 {
	if x != nil {
		return x.GramsOfSugar
	}
	return 0
}

func (x *MakeCoffeRequest) GetDelay() int32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

type isMakeCoffeRequest_CoffeRecipe interface {
	isMakeCoffeRequest_CoffeRecipe()
}

type MakeCoffeRequest_Cappuccino struct {
	Cappuccino *Cappuccino `protobuf:"bytes,1,opt,name=cappuccino,proto3,oneof"`
}

type MakeCoffeRequest_Macchiato struct {
	Macchiato *Macchiato `protobuf:"bytes,2,opt,name=macchiato,proto3,oneof"`
}

type MakeCoffeRequest_LongBlack struct {
	LongBlack *LongBlack `protobuf:"bytes,3,opt,name=long_black,json=longBlack,proto3,oneof"`
}

func (*MakeCoffeRequest_Cappuccino) isMakeCoffeRequest_CoffeRecipe() {}

func (*MakeCoffeRequest_Macchiato) isMakeCoffeRequest_CoffeRecipe() {}

func (*MakeCoffeRequest_LongBlack) isMakeCoffeRequest_CoffeRecipe() {}

// When a coffe is successfully brewed, this message is sent to the client
type MakeCoffeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MakeCoffeResponse) Reset() {
	*x = MakeCoffeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcCoffe_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeCoffeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeCoffeResponse) ProtoMessage() {}

func (x *MakeCoffeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcCoffe_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeCoffeResponse.ProtoReflect.Descriptor instead.
func (*MakeCoffeResponse) Descriptor() ([]byte, []int) {
	return file_grpcCoffe_proto_rawDescGZIP(), []int{4}
}

var File_grpcCoffe_proto protoreflect.FileDescriptor

var file_grpcCoffe_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x22, 0x59, 0x0a, 0x09, 0x4c, 0x6f, 0x6e,
	0x67, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x5f,
	0x6f, 0x66, 0x5f, 0x65, 0x73, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x4f, 0x66, 0x45, 0x73, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x68, 0x6f, 0x77, 0x5f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f,
	0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x6f, 0x77, 0x57, 0x61, 0x74,
	0x65, 0x72, 0x4d, 0x6c, 0x22, 0x5f, 0x0a, 0x09, 0x4d, 0x61, 0x63, 0x63, 0x68, 0x69, 0x61, 0x74,
	0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x65, 0x73,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x68,
	0x6f, 0x74, 0x73, 0x4f, 0x66, 0x45, 0x73, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x12, 0x26, 0x0a,
	0x0f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x6d, 0x69, 0x6c, 0x6b, 0x5f, 0x6d, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x4d,
	0x69, 0x6c, 0x6b, 0x4d, 0x6c, 0x22, 0xcf, 0x01, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x70, 0x75, 0x63,
	0x63, 0x69, 0x6e, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x5f, 0x6f, 0x66,
	0x5f, 0x65, 0x73, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x4f, 0x66, 0x45, 0x73, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x12, 0x26, 0x0a, 0x0f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x6d, 0x69, 0x6c, 0x6b,
	0x5f, 0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x65, 0x64, 0x4d, 0x69, 0x6c, 0x6b, 0x4d, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x6f, 0x61, 0x6d,
	0x65, 0x64, 0x5f, 0x6d, 0x69, 0x6c, 0x6b, 0x5f, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x66, 0x6f, 0x61, 0x6d, 0x65, 0x64, 0x4d, 0x69, 0x6c, 0x6b, 0x4d, 0x6c, 0x12, 0x47,
	0x0a, 0x20, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x67, 0x5f, 0x6f, 0x66,
	0x5f, 0x63, 0x68, 0x6f, 0x63, 0x6f, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x64,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x73, 0x70, 0x72, 0x69, 0x6e, 0x6b,
	0x6c, 0x69, 0x6e, 0x6b, 0x67, 0x4f, 0x66, 0x43, 0x68, 0x6f, 0x63, 0x6f, 0x6c, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x77, 0x65, 0x64, 0x65, 0x72, 0x22, 0xfb, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x6b, 0x65,
	0x43, 0x6f, 0x66, 0x66, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0a,
	0x63, 0x61, 0x70, 0x70, 0x75, 0x63, 0x63, 0x69, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x2e, 0x43, 0x61, 0x70, 0x70, 0x75, 0x63,
	0x63, 0x69, 0x6e, 0x6f, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x70, 0x75, 0x63, 0x63, 0x69,
	0x6e, 0x6f, 0x12, 0x31, 0x0a, 0x09, 0x6d, 0x61, 0x63, 0x63, 0x68, 0x69, 0x61, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x2e, 0x4d,
	0x61, 0x63, 0x63, 0x68, 0x69, 0x61, 0x74, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x63,
	0x68, 0x69, 0x61, 0x74, 0x6f, 0x12, 0x32, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x62, 0x6c,
	0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x72, 0x61,
	0x6d, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x75, 0x67, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x75, 0x67, 0x61, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x5f, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x6f, 0x66,
	0x66, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x52, 0x0a, 0x0c, 0x43, 0x6f,
	0x66, 0x66, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x4d, 0x61,
	0x6b, 0x65, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x2e, 0x4d, 0x61, 0x6b, 0x65, 0x43,
	0x6f, 0x66, 0x66, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7c,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x42, 0x0e, 0x47, 0x72,
	0x70, 0x63, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34, 0x6e, 0x74, 0x65, 0x2f,
	0x6d, 0x65, 0x65, 0x74, 0x75, 0x70, 0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x00, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0xca, 0x02, 0x06, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0xe2, 0x02,
	0x12, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcCoffe_proto_rawDescOnce sync.Once
	file_grpcCoffe_proto_rawDescData = file_grpcCoffe_proto_rawDesc
)

func file_grpcCoffe_proto_rawDescGZIP() []byte {
	file_grpcCoffe_proto_rawDescOnce.Do(func() {
		file_grpcCoffe_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcCoffe_proto_rawDescData)
	})
	return file_grpcCoffe_proto_rawDescData
}

var file_grpcCoffe_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpcCoffe_proto_goTypes = []interface{}{
	(*LongBlack)(nil),         // 0: coffee.LongBlack
	(*Macchiato)(nil),         // 1: coffee.Macchiato
	(*Cappuccino)(nil),        // 2: coffee.Cappuccino
	(*MakeCoffeRequest)(nil),  // 3: coffee.MakeCoffeRequest
	(*MakeCoffeResponse)(nil), // 4: coffee.MakeCoffeResponse
}
var file_grpcCoffe_proto_depIdxs = []int32{
	2, // 0: coffee.MakeCoffeRequest.cappuccino:type_name -> coffee.Cappuccino
	1, // 1: coffee.MakeCoffeRequest.macchiato:type_name -> coffee.Macchiato
	0, // 2: coffee.MakeCoffeRequest.long_black:type_name -> coffee.LongBlack
	3, // 3: coffee.CoffeService.MakeCoffe:input_type -> coffee.MakeCoffeRequest
	4, // 4: coffee.CoffeService.MakeCoffe:output_type -> coffee.MakeCoffeResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpcCoffe_proto_init() }
func file_grpcCoffe_proto_init() {
	if File_grpcCoffe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcCoffe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongBlack); i {
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
		file_grpcCoffe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Macchiato); i {
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
		file_grpcCoffe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cappuccino); i {
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
		file_grpcCoffe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeCoffeRequest); i {
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
		file_grpcCoffe_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeCoffeResponse); i {
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
	file_grpcCoffe_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*MakeCoffeRequest_Cappuccino)(nil),
		(*MakeCoffeRequest_Macchiato)(nil),
		(*MakeCoffeRequest_LongBlack)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcCoffe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcCoffe_proto_goTypes,
		DependencyIndexes: file_grpcCoffe_proto_depIdxs,
		MessageInfos:      file_grpcCoffe_proto_msgTypes,
	}.Build()
	File_grpcCoffe_proto = out.File
	file_grpcCoffe_proto_rawDesc = nil
	file_grpcCoffe_proto_goTypes = nil
	file_grpcCoffe_proto_depIdxs = nil
}
