// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: listwatcher.proto

package listwatcher

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define a change event for an item
type Event struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`               // Unique identifier for the event
	Data          *anypb.Any             `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`           // The event data, stored as a generic protobuf 'Any' type
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // Timestamp of when the event occurred
	Type          string                 `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`           // Type of event (e.g., "added", "updated", "deleted")
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_listwatcher_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_listwatcher_proto_msgTypes[0]
	if x != nil {
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
	return file_listwatcher_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Event) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_listwatcher_proto protoreflect.FileDescriptor

var file_listwatcher_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32, 0x4b,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x12, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x3b, 0x6c, 0x69, 0x73,
	0x74, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_listwatcher_proto_rawDescOnce sync.Once
	file_listwatcher_proto_rawDescData = file_listwatcher_proto_rawDesc
)

func file_listwatcher_proto_rawDescGZIP() []byte {
	file_listwatcher_proto_rawDescOnce.Do(func() {
		file_listwatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_listwatcher_proto_rawDescData)
	})
	return file_listwatcher_proto_rawDescData
}

var file_listwatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_listwatcher_proto_goTypes = []any{
	(*Event)(nil),                 // 0: listwatcher.Event
	(*anypb.Any)(nil),             // 1: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_listwatcher_proto_depIdxs = []int32{
	1, // 0: listwatcher.Event.data:type_name -> google.protobuf.Any
	2, // 1: listwatcher.Event.timestamp:type_name -> google.protobuf.Timestamp
	0, // 2: listwatcher.ListWatchService.ListWatch:input_type -> listwatcher.Event
	0, // 3: listwatcher.ListWatchService.ListWatch:output_type -> listwatcher.Event
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_listwatcher_proto_init() }
func file_listwatcher_proto_init() {
	if File_listwatcher_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_listwatcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listwatcher_proto_goTypes,
		DependencyIndexes: file_listwatcher_proto_depIdxs,
		MessageInfos:      file_listwatcher_proto_msgTypes,
	}.Build()
	File_listwatcher_proto = out.File
	file_listwatcher_proto_rawDesc = nil
	file_listwatcher_proto_goTypes = nil
	file_listwatcher_proto_depIdxs = nil
}
