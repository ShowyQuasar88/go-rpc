// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: person/person.proto

package person

import (
	home "github.com/ShowyQuasar88/rpc_demo/go_grpc/pb/home"
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

type Person_GENDER int32

const (
	Person_MAN   Person_GENDER = 0
	Person_WOMAN Person_GENDER = 1
	Person_GIRL  Person_GENDER = 1
	Person_OTHER Person_GENDER = 2
)

// Enum value maps for Person_GENDER.
var (
	Person_GENDER_name = map[int32]string{
		0: "MAN",
		1: "WOMAN",
		// Duplicate value: 1: "GIRL",
		2: "OTHER",
	}
	Person_GENDER_value = map[string]int32{
		"MAN":   0,
		"WOMAN": 1,
		"GIRL":  1,
		"OTHER": 2,
	}
)

func (x Person_GENDER) Enum() *Person_GENDER {
	p := new(Person_GENDER)
	*p = x
	return p
}

func (x Person_GENDER) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_GENDER) Descriptor() protoreflect.EnumDescriptor {
	return file_person_person_proto_enumTypes[0].Descriptor()
}

func (Person_GENDER) Type() protoreflect.EnumType {
	return &file_person_person_proto_enumTypes[0]
}

func (x Person_GENDER) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_GENDER.Descriptor instead.
func (Person_GENDER) EnumDescriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1, 0}
}

type House struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	V       *House_V  `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"` // 声明了需要使用才会产生，否则是不会生成的
}

func (x *House) Reset() {
	*x = House{}
	mi := &file_person_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *House) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*House) ProtoMessage() {}

func (x *House) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use House.ProtoReflect.Descriptor instead.
func (*House) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0}
}

func (x *House) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

func (x *House) GetV() *House_V {
	if x != nil {
		return x.V
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32             `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Sex     bool              `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	Test    []string          `protobuf:"bytes,4,rep,name=test,proto3" json:"test,omitempty"`                                                                                                              // repeated 声明切片
	TestMap map[string]string `protobuf:"bytes,5,rep,name=test_map,json=testMap,proto3" json:"test_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 声明 map，第一个只能是 string，第二个可以是之前定义过的类型，比如 Person
	// Types that are assignable to TestOneOf:
	//
	//	*Person_One
	//	*Person_Two
	//	*Person_Three
	TestOneOf isPerson_TestOneOf `protobuf_oneof:"TestOneOf"`
	Gender    Person_GENDER      `protobuf:"varint,6,opt,name=gender,proto3,enum=person.Person_GENDER" json:"gender,omitempty"`
	IHome     *home.Home         `protobuf:"bytes,10,opt,name=i_home,json=iHome,proto3" json:"i_home,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_person_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *Person) GetTest() []string {
	if x != nil {
		return x.Test
	}
	return nil
}

func (x *Person) GetTestMap() map[string]string {
	if x != nil {
		return x.TestMap
	}
	return nil
}

func (m *Person) GetTestOneOf() isPerson_TestOneOf {
	if m != nil {
		return m.TestOneOf
	}
	return nil
}

func (x *Person) GetOne() string {
	if x, ok := x.GetTestOneOf().(*Person_One); ok {
		return x.One
	}
	return ""
}

func (x *Person) GetTwo() string {
	if x, ok := x.GetTestOneOf().(*Person_Two); ok {
		return x.Two
	}
	return ""
}

func (x *Person) GetThree() string {
	if x, ok := x.GetTestOneOf().(*Person_Three); ok {
		return x.Three
	}
	return ""
}

func (x *Person) GetGender() Person_GENDER {
	if x != nil {
		return x.Gender
	}
	return Person_MAN
}

func (x *Person) GetIHome() *home.Home {
	if x != nil {
		return x.IHome
	}
	return nil
}

type isPerson_TestOneOf interface {
	isPerson_TestOneOf()
}

type Person_One struct {
	One string `protobuf:"bytes,7,opt,name=one,proto3,oneof"`
}

type Person_Two struct {
	Two string `protobuf:"bytes,8,opt,name=two,proto3,oneof"`
}

type Person_Three struct {
	Three string `protobuf:"bytes,9,opt,name=three,proto3,oneof"`
}

func (*Person_One) isPerson_TestOneOf() {}

func (*Person_Two) isPerson_TestOneOf() {}

func (*Person_Three) isPerson_TestOneOf() {}

type House_V struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *House_V) Reset() {
	*x = House_V{}
	mi := &file_person_person_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *House_V) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*House_V) ProtoMessage() {}

func (x *House_V) ProtoReflect() protoreflect.Message {
	mi := &file_person_person_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use House_V.ProtoReflect.Descriptor instead.
func (*House_V) Descriptor() ([]byte, []int) {
	return file_person_person_proto_rawDescGZIP(), []int{0, 0}
}

func (x *House_V) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_person_person_proto protoreflect.FileDescriptor

var file_person_person_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x0f, 0x68,
	0x6f, 0x6d, 0x65, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69,
	0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x12, 0x1d, 0x0a, 0x01, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x56, 0x52, 0x01, 0x76,
	0x1a, 0x17, 0x0a, 0x01, 0x56, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9e, 0x03, 0x0a, 0x06, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x03,
	0x74, 0x77, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x74, 0x77, 0x6f,
	0x12, 0x16, 0x0a, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x06, 0x69, 0x5f, 0x68, 0x6f, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x48,
	0x6f, 0x6d, 0x65, 0x52, 0x05, 0x69, 0x48, 0x6f, 0x6d, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x35, 0x0a, 0x06, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x4d,
	0x41, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x49, 0x52, 0x4c, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x6f, 0x77, 0x79, 0x51, 0x75,
	0x61, 0x73, 0x61, 0x72, 0x38, 0x38, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_person_proto_rawDescOnce sync.Once
	file_person_person_proto_rawDescData = file_person_person_proto_rawDesc
)

func file_person_person_proto_rawDescGZIP() []byte {
	file_person_person_proto_rawDescOnce.Do(func() {
		file_person_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_person_proto_rawDescData)
	})
	return file_person_person_proto_rawDescData
}

var file_person_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_person_person_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_person_person_proto_goTypes = []any{
	(Person_GENDER)(0), // 0: person.Person.GENDER
	(*House)(nil),      // 1: person.House
	(*Person)(nil),     // 2: person.Person
	(*House_V)(nil),    // 3: person.House.V
	nil,                // 4: person.Person.TestMapEntry
	(*home.Home)(nil),  // 5: home.Home
}
var file_person_person_proto_depIdxs = []int32{
	2, // 0: person.House.persons:type_name -> person.Person
	3, // 1: person.House.v:type_name -> person.House.V
	4, // 2: person.Person.test_map:type_name -> person.Person.TestMapEntry
	0, // 3: person.Person.gender:type_name -> person.Person.GENDER
	5, // 4: person.Person.i_home:type_name -> home.Home
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_person_person_proto_init() }
func file_person_person_proto_init() {
	if File_person_person_proto != nil {
		return
	}
	file_person_person_proto_msgTypes[1].OneofWrappers = []any{
		(*Person_One)(nil),
		(*Person_Two)(nil),
		(*Person_Three)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_person_proto_goTypes,
		DependencyIndexes: file_person_person_proto_depIdxs,
		EnumInfos:         file_person_person_proto_enumTypes,
		MessageInfos:      file_person_person_proto_msgTypes,
	}.Build()
	File_person_person_proto = out.File
	file_person_person_proto_rawDesc = nil
	file_person_person_proto_goTypes = nil
	file_person_person_proto_depIdxs = nil
}
