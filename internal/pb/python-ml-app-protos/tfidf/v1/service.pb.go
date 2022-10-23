// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: internal/python-ml-app-protos/tfidf/v1/service.proto

package v1

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

type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Synopsis string `protobuf:"bytes,1,opt,name=synopsis,proto3" json:"synopsis,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetSynopsis() string {
	if x != nil {
		return x.Synopsis
	}
	return ""
}

type PredictReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predictions []*Prediction `protobuf:"bytes,1,rep,name=predictions,proto3" json:"predictions,omitempty"`
}

func (x *PredictReply) Reset() {
	*x = PredictReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictReply) ProtoMessage() {}

func (x *PredictReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictReply.ProtoReflect.Descriptor instead.
func (*PredictReply) Descriptor() ([]byte, []int) {
	return file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *PredictReply) GetPredictions() []*Prediction {
	if x != nil {
		return x.Predictions
	}
	return nil
}

type Prediction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genre string `protobuf:"bytes,1,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *Prediction) Reset() {
	*x = Prediction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prediction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prediction) ProtoMessage() {}

func (x *Prediction) ProtoReflect() protoreflect.Message {
	mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prediction.ProtoReflect.Descriptor instead.
func (*Prediction) Descriptor() ([]byte, []int) {
	return file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *Prediction) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type TrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TrainRequest) Reset() {
	*x = TrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainRequest) ProtoMessage() {}

func (x *TrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainRequest.ProtoReflect.Descriptor instead.
func (*TrainRequest) Descriptor() ([]byte, []int) {
	return file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescGZIP(), []int{3}
}

type TrainReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TrainReply) Reset() {
	*x = TrainReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainReply) ProtoMessage() {}

func (x *TrainReply) ProtoReflect() protoreflect.Message {
	mi := &file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainReply.ProtoReflect.Descriptor instead.
func (*TrainReply) Descriptor() ([]byte, []int) {
	return file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *TrainReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_python_ml_app_protos_tfidf_v1_service_proto protoreflect.FileDescriptor

var file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f,
	0x6e, 0x2d, 0x6d, 0x6c, 0x2d, 0x61, 0x70, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x74, 0x66, 0x69, 0x64, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x66, 0x69, 0x64, 0x66, 0x22, 0x2c, 0x0a,
	0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x70, 0x73, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x70, 0x73, 0x69, 0x73, 0x22, 0x43, 0x0a, 0x0c, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x0b, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x66, 0x69, 0x64, 0x66, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x22, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x65, 0x6e, 0x72, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x26, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x78, 0x0a, 0x0a,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x74, 0x66, 0x69, 0x64, 0x66, 0x2e, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74,
	0x66, 0x69, 0x64, 0x66, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x74,
	0x66, 0x69, 0x64, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x74, 0x66, 0x69, 0x64, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescOnce sync.Once
	file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescData = file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDesc
)

func file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescGZIP() []byte {
	file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescOnce.Do(func() {
		file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescData)
	})
	return file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDescData
}

var file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_python_ml_app_protos_tfidf_v1_service_proto_goTypes = []interface{}{
	(*PredictRequest)(nil), // 0: tfidf.PredictRequest
	(*PredictReply)(nil),   // 1: tfidf.PredictReply
	(*Prediction)(nil),     // 2: tfidf.Prediction
	(*TrainRequest)(nil),   // 3: tfidf.TrainRequest
	(*TrainReply)(nil),     // 4: tfidf.TrainReply
}
var file_internal_python_ml_app_protos_tfidf_v1_service_proto_depIdxs = []int32{
	2, // 0: tfidf.PredictReply.predictions:type_name -> tfidf.Prediction
	0, // 1: tfidf.MovieGenre.Predict:input_type -> tfidf.PredictRequest
	3, // 2: tfidf.MovieGenre.Train:input_type -> tfidf.TrainRequest
	1, // 3: tfidf.MovieGenre.Predict:output_type -> tfidf.PredictReply
	4, // 4: tfidf.MovieGenre.Train:output_type -> tfidf.TrainReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_python_ml_app_protos_tfidf_v1_service_proto_init() }
func file_internal_python_ml_app_protos_tfidf_v1_service_proto_init() {
	if File_internal_python_ml_app_protos_tfidf_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictReply); i {
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
		file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prediction); i {
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
		file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainRequest); i {
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
		file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainReply); i {
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
			RawDescriptor: file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_python_ml_app_protos_tfidf_v1_service_proto_goTypes,
		DependencyIndexes: file_internal_python_ml_app_protos_tfidf_v1_service_proto_depIdxs,
		MessageInfos:      file_internal_python_ml_app_protos_tfidf_v1_service_proto_msgTypes,
	}.Build()
	File_internal_python_ml_app_protos_tfidf_v1_service_proto = out.File
	file_internal_python_ml_app_protos_tfidf_v1_service_proto_rawDesc = nil
	file_internal_python_ml_app_protos_tfidf_v1_service_proto_goTypes = nil
	file_internal_python_ml_app_protos_tfidf_v1_service_proto_depIdxs = nil
}
