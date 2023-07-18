// Copyright 2023 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: proto/experiment/command.proto

package experiment

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

type CreateGoalCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateGoalCommand) Reset() {
	*x = CreateGoalCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoalCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoalCommand) ProtoMessage() {}

func (x *CreateGoalCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoalCommand.ProtoReflect.Descriptor instead.
func (*CreateGoalCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGoalCommand) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateGoalCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGoalCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RenameGoalCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RenameGoalCommand) Reset() {
	*x = RenameGoalCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameGoalCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameGoalCommand) ProtoMessage() {}

func (x *RenameGoalCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameGoalCommand.ProtoReflect.Descriptor instead.
func (*RenameGoalCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{1}
}

func (x *RenameGoalCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChangeDescriptionGoalCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ChangeDescriptionGoalCommand) Reset() {
	*x = ChangeDescriptionGoalCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeDescriptionGoalCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeDescriptionGoalCommand) ProtoMessage() {}

func (x *ChangeDescriptionGoalCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeDescriptionGoalCommand.ProtoReflect.Descriptor instead.
func (*ChangeDescriptionGoalCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeDescriptionGoalCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ArchiveGoalCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArchiveGoalCommand) Reset() {
	*x = ArchiveGoalCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveGoalCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveGoalCommand) ProtoMessage() {}

func (x *ArchiveGoalCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveGoalCommand.ProtoReflect.Descriptor instead.
func (*ArchiveGoalCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{3}
}

type DeleteGoalCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGoalCommand) Reset() {
	*x = DeleteGoalCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoalCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoalCommand) ProtoMessage() {}

func (x *DeleteGoalCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoalCommand.ProtoReflect.Descriptor instead.
func (*DeleteGoalCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{4}
}

type CreateExperimentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureId       string   `protobuf:"bytes,1,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	StartAt         int64    `protobuf:"varint,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	StopAt          int64    `protobuf:"varint,4,opt,name=stop_at,json=stopAt,proto3" json:"stop_at,omitempty"`
	GoalIds         []string `protobuf:"bytes,5,rep,name=goal_ids,json=goalIds,proto3" json:"goal_ids,omitempty"`
	Name            string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description     string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	BaseVariationId string   `protobuf:"bytes,8,opt,name=base_variation_id,json=baseVariationId,proto3" json:"base_variation_id,omitempty"`
}

func (x *CreateExperimentCommand) Reset() {
	*x = CreateExperimentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExperimentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExperimentCommand) ProtoMessage() {}

func (x *CreateExperimentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExperimentCommand.ProtoReflect.Descriptor instead.
func (*CreateExperimentCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{5}
}

func (x *CreateExperimentCommand) GetFeatureId() string {
	if x != nil {
		return x.FeatureId
	}
	return ""
}

func (x *CreateExperimentCommand) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *CreateExperimentCommand) GetStopAt() int64 {
	if x != nil {
		return x.StopAt
	}
	return 0
}

func (x *CreateExperimentCommand) GetGoalIds() []string {
	if x != nil {
		return x.GoalIds
	}
	return nil
}

func (x *CreateExperimentCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateExperimentCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateExperimentCommand) GetBaseVariationId() string {
	if x != nil {
		return x.BaseVariationId
	}
	return ""
}

type ChangeExperimentPeriodCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartAt int64 `protobuf:"varint,1,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	StopAt  int64 `protobuf:"varint,2,opt,name=stop_at,json=stopAt,proto3" json:"stop_at,omitempty"`
}

func (x *ChangeExperimentPeriodCommand) Reset() {
	*x = ChangeExperimentPeriodCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeExperimentPeriodCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeExperimentPeriodCommand) ProtoMessage() {}

func (x *ChangeExperimentPeriodCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeExperimentPeriodCommand.ProtoReflect.Descriptor instead.
func (*ChangeExperimentPeriodCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{6}
}

func (x *ChangeExperimentPeriodCommand) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *ChangeExperimentPeriodCommand) GetStopAt() int64 {
	if x != nil {
		return x.StopAt
	}
	return 0
}

type ChangeExperimentNameCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ChangeExperimentNameCommand) Reset() {
	*x = ChangeExperimentNameCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeExperimentNameCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeExperimentNameCommand) ProtoMessage() {}

func (x *ChangeExperimentNameCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeExperimentNameCommand.ProtoReflect.Descriptor instead.
func (*ChangeExperimentNameCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{7}
}

func (x *ChangeExperimentNameCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChangeExperimentDescriptionCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ChangeExperimentDescriptionCommand) Reset() {
	*x = ChangeExperimentDescriptionCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeExperimentDescriptionCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeExperimentDescriptionCommand) ProtoMessage() {}

func (x *ChangeExperimentDescriptionCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeExperimentDescriptionCommand.ProtoReflect.Descriptor instead.
func (*ChangeExperimentDescriptionCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{8}
}

func (x *ChangeExperimentDescriptionCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type StopExperimentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopExperimentCommand) Reset() {
	*x = StopExperimentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopExperimentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopExperimentCommand) ProtoMessage() {}

func (x *StopExperimentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopExperimentCommand.ProtoReflect.Descriptor instead.
func (*StopExperimentCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{9}
}

type ArchiveExperimentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ArchiveExperimentCommand) Reset() {
	*x = ArchiveExperimentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveExperimentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveExperimentCommand) ProtoMessage() {}

func (x *ArchiveExperimentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveExperimentCommand.ProtoReflect.Descriptor instead.
func (*ArchiveExperimentCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{10}
}

type DeleteExperimentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteExperimentCommand) Reset() {
	*x = DeleteExperimentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExperimentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExperimentCommand) ProtoMessage() {}

func (x *DeleteExperimentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExperimentCommand.ProtoReflect.Descriptor instead.
func (*DeleteExperimentCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{11}
}

type StartExperimentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartExperimentCommand) Reset() {
	*x = StartExperimentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartExperimentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartExperimentCommand) ProtoMessage() {}

func (x *StartExperimentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartExperimentCommand.ProtoReflect.Descriptor instead.
func (*StartExperimentCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{12}
}

type FinishExperimentCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinishExperimentCommand) Reset() {
	*x = FinishExperimentCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_experiment_command_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishExperimentCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishExperimentCommand) ProtoMessage() {}

func (x *FinishExperimentCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_experiment_command_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishExperimentCommand.ProtoReflect.Descriptor instead.
func (*FinishExperimentCommand) Descriptor() ([]byte, []int) {
	return file_proto_experiment_command_proto_rawDescGZIP(), []int{13}
}

var File_proto_experiment_command_proto protoreflect.FileDescriptor

var file_proto_experiment_command_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x6f, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x27, 0x0a, 0x11, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x1c, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x6f, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a, 0x12,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xef, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x6f, 0x70, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x61, 0x6c, 0x49, 0x64,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x53, 0x0a, 0x1d, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x70, 0x41, 0x74, 0x22, 0x31,
	0x0a, 0x1b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x46, 0x0a, 0x22, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x74, 0x6f,
	0x70, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x19,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_experiment_command_proto_rawDescOnce sync.Once
	file_proto_experiment_command_proto_rawDescData = file_proto_experiment_command_proto_rawDesc
)

func file_proto_experiment_command_proto_rawDescGZIP() []byte {
	file_proto_experiment_command_proto_rawDescOnce.Do(func() {
		file_proto_experiment_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_experiment_command_proto_rawDescData)
	})
	return file_proto_experiment_command_proto_rawDescData
}

var file_proto_experiment_command_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_experiment_command_proto_goTypes = []interface{}{
	(*CreateGoalCommand)(nil),                  // 0: bucketeer.experiment.CreateGoalCommand
	(*RenameGoalCommand)(nil),                  // 1: bucketeer.experiment.RenameGoalCommand
	(*ChangeDescriptionGoalCommand)(nil),       // 2: bucketeer.experiment.ChangeDescriptionGoalCommand
	(*ArchiveGoalCommand)(nil),                 // 3: bucketeer.experiment.ArchiveGoalCommand
	(*DeleteGoalCommand)(nil),                  // 4: bucketeer.experiment.DeleteGoalCommand
	(*CreateExperimentCommand)(nil),            // 5: bucketeer.experiment.CreateExperimentCommand
	(*ChangeExperimentPeriodCommand)(nil),      // 6: bucketeer.experiment.ChangeExperimentPeriodCommand
	(*ChangeExperimentNameCommand)(nil),        // 7: bucketeer.experiment.ChangeExperimentNameCommand
	(*ChangeExperimentDescriptionCommand)(nil), // 8: bucketeer.experiment.ChangeExperimentDescriptionCommand
	(*StopExperimentCommand)(nil),              // 9: bucketeer.experiment.StopExperimentCommand
	(*ArchiveExperimentCommand)(nil),           // 10: bucketeer.experiment.ArchiveExperimentCommand
	(*DeleteExperimentCommand)(nil),            // 11: bucketeer.experiment.DeleteExperimentCommand
	(*StartExperimentCommand)(nil),             // 12: bucketeer.experiment.StartExperimentCommand
	(*FinishExperimentCommand)(nil),            // 13: bucketeer.experiment.FinishExperimentCommand
}
var file_proto_experiment_command_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_experiment_command_proto_init() }
func file_proto_experiment_command_proto_init() {
	if File_proto_experiment_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_experiment_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoalCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameGoalCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeDescriptionGoalCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveGoalCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoalCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExperimentCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeExperimentPeriodCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeExperimentNameCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeExperimentDescriptionCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopExperimentCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveExperimentCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteExperimentCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartExperimentCommand); i {
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
		file_proto_experiment_command_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishExperimentCommand); i {
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
			RawDescriptor: file_proto_experiment_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_experiment_command_proto_goTypes,
		DependencyIndexes: file_proto_experiment_command_proto_depIdxs,
		MessageInfos:      file_proto_experiment_command_proto_msgTypes,
	}.Build()
	File_proto_experiment_command_proto = out.File
	file_proto_experiment_command_proto_rawDesc = nil
	file_proto_experiment_command_proto_goTypes = nil
	file_proto_experiment_command_proto_depIdxs = nil
}
