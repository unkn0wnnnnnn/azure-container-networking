// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: transport.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DatapathPodMetadata_APIVersion int32

const (
	DatapathPodMetadata_V1 DatapathPodMetadata_APIVersion = 0
)

// Enum value maps for DatapathPodMetadata_APIVersion.
var (
	DatapathPodMetadata_APIVersion_name = map[int32]string{
		0: "V1",
	}
	DatapathPodMetadata_APIVersion_value = map[string]int32{
		"V1": 0,
	}
)

func (x DatapathPodMetadata_APIVersion) Enum() *DatapathPodMetadata_APIVersion {
	p := new(DatapathPodMetadata_APIVersion)
	*p = x
	return p
}

func (x DatapathPodMetadata_APIVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatapathPodMetadata_APIVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_proto_enumTypes[0].Descriptor()
}

func (DatapathPodMetadata_APIVersion) Type() protoreflect.EnumType {
	return &file_transport_proto_enumTypes[0]
}

func (x DatapathPodMetadata_APIVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatapathPodMetadata_APIVersion.Descriptor instead.
func (DatapathPodMetadata_APIVersion) EnumDescriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0, 0}
}

// The following events are supported:
type Events_EventType int32

const (
	// Send a CREATE/UPDATE event to a datapath Pod
	Events_APPLY Events_EventType = 0
	// Send a DELETE event to a datapath Pod
	Events_REMOVE Events_EventType = 1
)

// Enum value maps for Events_EventType.
var (
	Events_EventType_name = map[int32]string{
		0: "APPLY",
		1: "REMOVE",
	}
	Events_EventType_value = map[string]int32{
		"APPLY":  0,
		"REMOVE": 1,
	}
)

func (x Events_EventType) Enum() *Events_EventType {
	p := new(Events_EventType)
	*p = x
	return p
}

func (x Events_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Events_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_proto_enumTypes[1].Descriptor()
}

func (Events_EventType) Type() protoreflect.EnumType {
	return &file_transport_proto_enumTypes[1]
}

func (x Events_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Events_EventType.Descriptor instead.
func (Events_EventType) EnumDescriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1, 0}
}

// Supported objects are IPSet and NetworkPolicy objects.
type Events_ObjectType int32

const (
	Events_IPSET  Events_ObjectType = 0
	Events_POLICY Events_ObjectType = 1
)

// Enum value maps for Events_ObjectType.
var (
	Events_ObjectType_name = map[int32]string{
		0: "IPSET",
		1: "POLICY",
	}
	Events_ObjectType_value = map[string]int32{
		"IPSET":  0,
		"POLICY": 1,
	}
)

func (x Events_ObjectType) Enum() *Events_ObjectType {
	p := new(Events_ObjectType)
	*p = x
	return p
}

func (x Events_ObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Events_ObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_proto_enumTypes[2].Descriptor()
}

func (Events_ObjectType) Type() protoreflect.EnumType {
	return &file_transport_proto_enumTypes[2]
}

func (x Events_ObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Events_ObjectType.Descriptor instead.
func (Events_ObjectType) EnumDescriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1, 1}
}

// DatapathPodMetadata is the metadata for a datapath pod
type DatapathPodMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName    string                         `protobuf:"bytes,1,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`                                    // Daemonset Pod ID
	NodeName   string                         `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`                                 // Node name
	ApiVersion DatapathPodMetadata_APIVersion `protobuf:"varint,3,opt,name=apiVersion,proto3,enum=protos.DatapathPodMetadata_APIVersion" json:"apiVersion,omitempty"` // Controlplane API version to support backwards compatibility
}

func (x *DatapathPodMetadata) Reset() {
	*x = DatapathPodMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatapathPodMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatapathPodMetadata) ProtoMessage() {}

func (x *DatapathPodMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatapathPodMetadata.ProtoReflect.Descriptor instead.
func (*DatapathPodMetadata) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0}
}

func (x *DatapathPodMetadata) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *DatapathPodMetadata) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *DatapathPodMetadata) GetApiVersion() DatapathPodMetadata_APIVersion {
	if x != nil {
		return x.ApiVersion
	}
	return DatapathPodMetadata_V1
}

// Events defines the operation (event type) and object type being
// streamed to the datapath client. A events message may carry one or
// more Event objects.
type Events struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Events_EventType  `protobuf:"varint,1,opt,name=type,proto3,enum=protos.Events_EventType" json:"type,omitempty"`
	Object Events_ObjectType `protobuf:"varint,2,opt,name=object,proto3,enum=protos.Events_ObjectType" json:"object,omitempty"`
	// Payload can contain one or more Event objects.
	Event []*Event `protobuf:"bytes,3,rep,name=event,proto3" json:"event,omitempty"`
}

func (x *Events) Reset() {
	*x = Events{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Events) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events) ProtoMessage() {}

func (x *Events) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events.ProtoReflect.Descriptor instead.
func (*Events) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1}
}

func (x *Events) GetType() Events_EventType {
	if x != nil {
		return x.Type
	}
	return Events_APPLY
}

func (x *Events) GetObject() Events_ObjectType {
	if x != nil {
		return x.Object
	}
	return Events_IPSET
}

func (x *Events) GetEvent() []*Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// Event is a generic object that can be Created,
// Updated, Deleted by the controlplane.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data can contain one or more instances of IPSet or NetworkPolicy
	// objects.
	Data []*structpb.Struct `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetData() []*structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x61, 0x74, 0x68, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x61, 0x74, 0x68, 0x50, 0x6f, 0x64,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x14, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a,
	0x02, 0x56, 0x31, 0x10, 0x00, 0x22, 0xd7, 0x01, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x31,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x22, 0x23, 0x0a, 0x0a, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x50, 0x53, 0x45,
	0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x01, 0x22,
	0x34, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4b, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x70, 0x61, 0x74, 0x68, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x30, 0x01, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x6e, 0x70, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_rawDescOnce sync.Once
	file_transport_proto_rawDescData = file_transport_proto_rawDesc
)

func file_transport_proto_rawDescGZIP() []byte {
	file_transport_proto_rawDescOnce.Do(func() {
		file_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_rawDescData)
	})
	return file_transport_proto_rawDescData
}

var file_transport_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transport_proto_goTypes = []interface{}{
	(DatapathPodMetadata_APIVersion)(0), // 0: protos.DatapathPodMetadata.APIVersion
	(Events_EventType)(0),               // 1: protos.Events.EventType
	(Events_ObjectType)(0),              // 2: protos.Events.ObjectType
	(*DatapathPodMetadata)(nil),         // 3: protos.DatapathPodMetadata
	(*Events)(nil),                      // 4: protos.Events
	(*Event)(nil),                       // 5: protos.Event
	(*structpb.Struct)(nil),             // 6: google.protobuf.Struct
}
var file_transport_proto_depIdxs = []int32{
	0, // 0: protos.DatapathPodMetadata.apiVersion:type_name -> protos.DatapathPodMetadata.APIVersion
	1, // 1: protos.Events.type:type_name -> protos.Events.EventType
	2, // 2: protos.Events.object:type_name -> protos.Events.ObjectType
	5, // 3: protos.Events.event:type_name -> protos.Event
	6, // 4: protos.Event.data:type_name -> google.protobuf.Struct
	3, // 5: protos.DataplaneEvents.Connect:input_type -> protos.DatapathPodMetadata
	4, // 6: protos.DataplaneEvents.Connect:output_type -> protos.Events
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatapathPodMetadata); i {
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
		file_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Events); i {
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
		file_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
		EnumInfos:         file_transport_proto_enumTypes,
		MessageInfos:      file_transport_proto_msgTypes,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}
