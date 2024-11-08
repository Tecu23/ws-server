// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: proto/chess_engine.proto

package enginepb

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

type MoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fen  string `protobuf:"bytes,1,opt,name=fen,proto3" json:"fen,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	mi := &file_proto_chess_engine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_engine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_chess_engine_proto_rawDescGZIP(), []int{0}
}

func (x *MoveRequest) GetFen() string {
	if x != nil {
		return x.Fen
	}
	return ""
}

func (x *MoveRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type MoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BestMove string `protobuf:"bytes,1,opt,name=best_move,json=bestMove,proto3" json:"best_move,omitempty"`
}

func (x *MoveResponse) Reset() {
	*x = MoveResponse{}
	mi := &file_proto_chess_engine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResponse) ProtoMessage() {}

func (x *MoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chess_engine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResponse.ProtoReflect.Descriptor instead.
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return file_proto_chess_engine_proto_rawDescGZIP(), []int{1}
}

func (x *MoveResponse) GetBestMove() string {
	if x != nil {
		return x.BestMove
	}
	return ""
}

var File_proto_chess_engine_proto protoreflect.FileDescriptor

var file_proto_chess_engine_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0b, 0x4d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x2b, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x32, 0x3f, 0x0a, 0x0b,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x11, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x65, 0x73, 0x74, 0x4d, 0x6f, 0x76, 0x65,
	0x12, 0x0c, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a,
	0x09, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_chess_engine_proto_rawDescOnce sync.Once
	file_proto_chess_engine_proto_rawDescData = file_proto_chess_engine_proto_rawDesc
)

func file_proto_chess_engine_proto_rawDescGZIP() []byte {
	file_proto_chess_engine_proto_rawDescOnce.Do(func() {
		file_proto_chess_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chess_engine_proto_rawDescData)
	})
	return file_proto_chess_engine_proto_rawDescData
}

var file_proto_chess_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_chess_engine_proto_goTypes = []any{
	(*MoveRequest)(nil),  // 0: MoveRequest
	(*MoveResponse)(nil), // 1: MoveResponse
}
var file_proto_chess_engine_proto_depIdxs = []int32{
	0, // 0: ChessEngine.CalculateBestMove:input_type -> MoveRequest
	1, // 1: ChessEngine.CalculateBestMove:output_type -> MoveResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_chess_engine_proto_init() }
func file_proto_chess_engine_proto_init() {
	if File_proto_chess_engine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_chess_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chess_engine_proto_goTypes,
		DependencyIndexes: file_proto_chess_engine_proto_depIdxs,
		MessageInfos:      file_proto_chess_engine_proto_msgTypes,
	}.Build()
	File_proto_chess_engine_proto = out.File
	file_proto_chess_engine_proto_rawDesc = nil
	file_proto_chess_engine_proto_goTypes = nil
	file_proto_chess_engine_proto_depIdxs = nil
}