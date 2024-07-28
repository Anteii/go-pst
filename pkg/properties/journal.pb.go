// go-pst is a library for reading Personal Storage Table (.pst) files (written in Go/Golang).
//
// Copyright 2023 Marten Mooij
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate msgp -tests=false

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: journal.proto

package properties

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

type Journal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether the document was sent by email or posted to a server folder during journaling.
	LogDocumentPosted *bool `protobuf:"varint,1,opt,name=log_document_posted,json=logDocumentPosted,proto3,oneof" json:"log_document_posted,omitempty" msg:"26934511,omitempty"`  
	// Indicates whether the document was printed during journaling.
	LogDocumentPrinted *bool `protobuf:"varint,2,opt,name=log_document_printed,json=logDocumentPrinted,proto3,oneof" json:"log_document_printed,omitempty" msg:"26932611,omitempty"`  
	// Indicates whether the document was sent to a routing recipient during journaling.
	LogDocumentRouted *bool `protobuf:"varint,3,opt,name=log_document_routed,json=logDocumentRouted,proto3,oneof" json:"log_document_routed,omitempty" msg:"26934411,omitempty"`  
	// Indicates whether the document was saved during journaling.
	LogDocumentSaved *bool `protobuf:"varint,4,opt,name=log_document_saved,json=logDocumentSaved,proto3,oneof" json:"log_document_saved,omitempty" msg:"26932711,omitempty"`  
	// Contains the duration, in minutes, of the activity.
	LogDuration *int32 `protobuf:"varint,5,opt,name=log_duration,json=logDuration,proto3,oneof" json:"log_duration,omitempty" msg:"2693193,omitempty"`  
	// Contains the time, in UTC, at which the activity ended.
	LogEnd *int64 `protobuf:"varint,6,opt,name=log_end,json=logEnd,proto3,oneof" json:"log_end,omitempty" msg:"26932064,omitempty"`  
	// Contains metadata about the Journal object.
	LogFlags *int32 `protobuf:"varint,7,opt,name=log_flags,json=logFlags,proto3,oneof" json:"log_flags,omitempty" msg:"2693243,omitempty"`  
	// Contains the time, in UTC, at which the activity began.
	LogStart *int64 `protobuf:"varint,8,opt,name=log_start,json=logStart,proto3,oneof" json:"log_start,omitempty" msg:"26931864,omitempty"`  
	// Briefly describes the journal activity that is being recorded.
	LogType *string `protobuf:"bytes,9,opt,name=log_type,json=logType,proto3,oneof" json:"log_type,omitempty" msg:"26931231,omitempty"`  
	// Contains an expanded description of the journal activity that is being recorded.
	LogTypeDesc *string `protobuf:"bytes,10,opt,name=log_type_desc,json=logTypeDesc,proto3,oneof" json:"log_type_desc,omitempty" msg:"26934631,omitempty"`  
}

func (x *Journal) Reset() {
	*x = Journal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_journal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Journal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Journal) ProtoMessage() {}

func (x *Journal) ProtoReflect() protoreflect.Message {
	mi := &file_journal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Journal.ProtoReflect.Descriptor instead.
func (*Journal) Descriptor() ([]byte, []int) {
	return file_journal_proto_rawDescGZIP(), []int{0}
}

func (x *Journal) GetLogDocumentPosted() bool {
	if x != nil && x.LogDocumentPosted != nil {
		return *x.LogDocumentPosted
	}
	return false
}

func (x *Journal) GetLogDocumentPrinted() bool {
	if x != nil && x.LogDocumentPrinted != nil {
		return *x.LogDocumentPrinted
	}
	return false
}

func (x *Journal) GetLogDocumentRouted() bool {
	if x != nil && x.LogDocumentRouted != nil {
		return *x.LogDocumentRouted
	}
	return false
}

func (x *Journal) GetLogDocumentSaved() bool {
	if x != nil && x.LogDocumentSaved != nil {
		return *x.LogDocumentSaved
	}
	return false
}

func (x *Journal) GetLogDuration() int32 {
	if x != nil && x.LogDuration != nil {
		return *x.LogDuration
	}
	return 0
}

func (x *Journal) GetLogEnd() int64 {
	if x != nil && x.LogEnd != nil {
		return *x.LogEnd
	}
	return 0
}

func (x *Journal) GetLogFlags() int32 {
	if x != nil && x.LogFlags != nil {
		return *x.LogFlags
	}
	return 0
}

func (x *Journal) GetLogStart() int64 {
	if x != nil && x.LogStart != nil {
		return *x.LogStart
	}
	return 0
}

func (x *Journal) GetLogType() string {
	if x != nil && x.LogType != nil {
		return *x.LogType
	}
	return ""
}

func (x *Journal) GetLogTypeDesc() string {
	if x != nil && x.LogTypeDesc != nil {
		return *x.LogTypeDesc
	}
	return ""
}

var File_journal_proto protoreflect.FileDescriptor

var file_journal_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe8, 0x04, 0x0a, 0x07, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x13, 0x6c,
	0x6f, 0x67, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x11, 0x6c, 0x6f, 0x67, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x35, 0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01,
	0x52, 0x12, 0x6c, 0x6f, 0x67, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69,
	0x6e, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x6c, 0x6f, 0x67, 0x5f, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x11, 0x6c, 0x6f, 0x67, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12,
	0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x61, 0x76,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03, 0x52, 0x10, 0x6c, 0x6f, 0x67, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x26, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x5f, 0x65,
	0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x45,
	0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x07, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x6c, 0x6f, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x88,
	0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6c,
	0x6f, 0x67, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x65, 0x64, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x5f,
	0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x61, 0x76,
	0x65, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x6f, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6f, 0x69, 0x6a, 0x74, 0x65,
	0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x73, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_journal_proto_rawDescOnce sync.Once
	file_journal_proto_rawDescData = file_journal_proto_rawDesc
)

func file_journal_proto_rawDescGZIP() []byte {
	file_journal_proto_rawDescOnce.Do(func() {
		file_journal_proto_rawDescData = protoimpl.X.CompressGZIP(file_journal_proto_rawDescData)
	})
	return file_journal_proto_rawDescData
}

var file_journal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_journal_proto_goTypes = []any{
	(*Journal)(nil), // 0: Journal
}
var file_journal_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_journal_proto_init() }
func file_journal_proto_init() {
	if File_journal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_journal_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Journal); i {
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
	file_journal_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_journal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_journal_proto_goTypes,
		DependencyIndexes: file_journal_proto_depIdxs,
		MessageInfos:      file_journal_proto_msgTypes,
	}.Build()
	File_journal_proto = out.File
	file_journal_proto_rawDesc = nil
	file_journal_proto_goTypes = nil
	file_journal_proto_depIdxs = nil
}
