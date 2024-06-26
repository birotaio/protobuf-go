// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v4.25.3
// source: protofif.proto

package protofif

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int64  `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32  `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	Loc     string `protobuf:"bytes,3,opt,name=loc,proto3" json:"loc,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofif_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_protofif_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_protofif_proto_rawDescGZIP(), []int{0}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

func (x *Timestamp) GetLoc() string {
	if x != nil {
		return x.Loc
	}
	return ""
}

type Duration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nanoseconds int64 `protobuf:"varint,1,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
}

func (x *Duration) Reset() {
	*x = Duration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofif_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Duration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duration) ProtoMessage() {}

func (x *Duration) ProtoReflect() protoreflect.Message {
	mi := &file_protofif_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duration.ProtoReflect.Descriptor instead.
func (*Duration) Descriptor() ([]byte, []int) {
	return file_protofif_proto_rawDescGZIP(), []int{1}
}

func (x *Duration) GetNanoseconds() int64 {
	if x != nil {
		return x.Nanoseconds
	}
	return 0
}

var file_protofif_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         65001,
		Name:          "protofif.bsontags",
		Tag:           "varint,65001,opt,name=bsontags",
		Filename:      "protofif.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         65002,
		Name:          "protofif.non_editable",
		Tag:           "varint,65002,opt,name=non_editable",
		Filename:      "protofif.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         65003,
		Name:          "protofif.database_default",
		Tag:           "varint,65003,opt,name=database_default",
		Filename:      "protofif.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         65001,
		Name:          "protofif.json",
		Tag:           "bytes,65001,opt,name=json",
		Filename:      "protofif.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         65002,
		Name:          "protofif.moretags",
		Tag:           "bytes,65002,opt,name=moretags",
		Filename:      "protofif.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         65003,
		Name:          "protofif.embed",
		Tag:           "varint,65003,opt,name=embed",
		Filename:      "protofif.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Defaults tag bson:"<jsonname>,omitempty" on all go struct fields
	//
	// optional bool bsontags = 65001;
	E_Bsontags = &file_protofif_proto_extTypes[0]
	// Defaults tag editable:"false" on all go struct fields
	//
	// optional bool non_editable = 65002;
	E_NonEditable = &file_protofif_proto_extTypes[1]
	// Defaults tag database:"default" on all go struct fields
	//
	// optional bool database_default = 65003;
	E_DatabaseDefault = &file_protofif_proto_extTypes[2]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string json = 65001;
	E_Json = &file_protofif_proto_extTypes[3]
	// optional string moretags = 65002;
	E_Moretags = &file_protofif_proto_extTypes[4]
	// optional bool embed = 65003;
	E_Embed = &file_protofif_proto_extTypes[5]
)

var File_protofif_proto protoreflect.FileDescriptor

var file_protofif_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x66, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x09,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x63, 0x22, 0x2c, 0x0a, 0x08, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x3a, 0x40, 0x0a, 0x08, 0x62, 0x73, 0x6f,
	0x6e, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe9, 0xfb, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x62, 0x73, 0x6f, 0x6e, 0x74, 0x61, 0x67, 0x73, 0x88, 0x01, 0x01, 0x3a, 0x47, 0x0a, 0x0c, 0x6e,
	0x6f, 0x6e, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xea, 0xfb, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e, 0x6f, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x88, 0x01, 0x01, 0x3a, 0x4f, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xeb, 0xfb, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x36, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe9, 0xfb, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x3a, 0x3e, 0x0a,
	0x08, 0x6d, 0x6f, 0x72, 0x65, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xea, 0xfb, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x6f, 0x72, 0x65, 0x74, 0x61, 0x67, 0x73, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x0a,
	0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xeb, 0xfb, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofif_proto_rawDescOnce sync.Once
	file_protofif_proto_rawDescData = file_protofif_proto_rawDesc
)

func file_protofif_proto_rawDescGZIP() []byte {
	file_protofif_proto_rawDescOnce.Do(func() {
		file_protofif_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofif_proto_rawDescData)
	})
	return file_protofif_proto_rawDescData
}

var file_protofif_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protofif_proto_goTypes = []interface{}{
	(*Timestamp)(nil),                   // 0: protofif.Timestamp
	(*Duration)(nil),                    // 1: protofif.Duration
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
}
var file_protofif_proto_depIdxs = []int32{
	2, // 0: protofif.bsontags:extendee -> google.protobuf.MessageOptions
	2, // 1: protofif.non_editable:extendee -> google.protobuf.MessageOptions
	2, // 2: protofif.database_default:extendee -> google.protobuf.MessageOptions
	3, // 3: protofif.json:extendee -> google.protobuf.FieldOptions
	3, // 4: protofif.moretags:extendee -> google.protobuf.FieldOptions
	3, // 5: protofif.embed:extendee -> google.protobuf.FieldOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	0, // [0:6] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protofif_proto_init() }
func file_protofif_proto_init() {
	if File_protofif_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofif_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
		file_protofif_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Duration); i {
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
			RawDescriptor: file_protofif_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_protofif_proto_goTypes,
		DependencyIndexes: file_protofif_proto_depIdxs,
		MessageInfos:      file_protofif_proto_msgTypes,
		ExtensionInfos:    file_protofif_proto_extTypes,
	}.Build()
	File_protofif_proto = out.File
	file_protofif_proto_rawDesc = nil
	file_protofif_proto_goTypes = nil
	file_protofif_proto_depIdxs = nil
}
