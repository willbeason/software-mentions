// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: papers/software-mentions.proto

package papers

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

type SoftwareMentions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application      string       `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	Version          string       `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Date             string       `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Md5              string       `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5,omitempty"`
	Pages            []*Page      `protobuf:"bytes,5,rep,name=pages,proto3" json:"pages,omitempty"`
	Mentions         []*Mention   `protobuf:"bytes,6,rep,name=mentions,proto3" json:"mentions,omitempty"`
	References       []*Reference `protobuf:"bytes,7,rep,name=references,proto3" json:"references,omitempty"`
	Runtime          int32        `protobuf:"varint,8,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Id               string       `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
	Metadata         *Mention     `protobuf:"bytes,10,opt,name=metadata,proto3" json:"metadata,omitempty"`
	OriginalFilePath string       `protobuf:"bytes,11,opt,name=original_file_path,json=originalFilePath,proto3" json:"original_file_path,omitempty"`
	FileName         string       `protobuf:"bytes,12,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *SoftwareMentions) Reset() {
	*x = SoftwareMentions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoftwareMentions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoftwareMentions) ProtoMessage() {}

func (x *SoftwareMentions) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoftwareMentions.ProtoReflect.Descriptor instead.
func (*SoftwareMentions) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{0}
}

func (x *SoftwareMentions) GetApplication() string {
	if x != nil {
		return x.Application
	}
	return ""
}

func (x *SoftwareMentions) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SoftwareMentions) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *SoftwareMentions) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *SoftwareMentions) GetPages() []*Page {
	if x != nil {
		return x.Pages
	}
	return nil
}

func (x *SoftwareMentions) GetMentions() []*Mention {
	if x != nil {
		return x.Mentions
	}
	return nil
}

func (x *SoftwareMentions) GetReferences() []*Reference {
	if x != nil {
		return x.References
	}
	return nil
}

func (x *SoftwareMentions) GetRuntime() int32 {
	if x != nil {
		return x.Runtime
	}
	return 0
}

func (x *SoftwareMentions) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SoftwareMentions) GetMetadata() *Mention {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SoftwareMentions) GetOriginalFilePath() string {
	if x != nil {
		return x.OriginalFilePath
	}
	return ""
}

func (x *SoftwareMentions) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageHeight float32 `protobuf:"fixed32,1,opt,name=page_height,json=pageHeight,proto3" json:"page_height,omitempty"`
	PageWidth  float32 `protobuf:"fixed32,2,opt,name=page_width,json=pageWidth,proto3" json:"page_width,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetPageHeight() float32 {
	if x != nil {
		return x.PageHeight
	}
	return 0
}

func (x *Page) GetPageWidth() float32 {
	if x != nil {
		return x.PageWidth
	}
	return 0
}

type Mention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoftwareName              *SoftwareName      `protobuf:"bytes,1,opt,name=software_name,json=softwareName,proto3" json:"software_name,omitempty"`
	Type                      string             `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Version                   *Version           `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Language                  *SoftwareName      `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Url                       *SoftwareName      `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Context                   string             `protobuf:"bytes,6,opt,name=context,proto3" json:"context,omitempty"`
	MentionContextAttributes  *ContextAttributes `protobuf:"bytes,7,opt,name=mentionContextAttributes,proto3" json:"mentionContextAttributes,omitempty"`
	DocumentContextAttributes *ContextAttributes `protobuf:"bytes,8,opt,name=documentContextAttributes,proto3" json:"documentContextAttributes,omitempty"`
	References                []*Reference       `protobuf:"bytes,9,rep,name=references,proto3" json:"references,omitempty"`
}

func (x *Mention) Reset() {
	*x = Mention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mention) ProtoMessage() {}

func (x *Mention) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mention.ProtoReflect.Descriptor instead.
func (*Mention) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{2}
}

func (x *Mention) GetSoftwareName() *SoftwareName {
	if x != nil {
		return x.SoftwareName
	}
	return nil
}

func (x *Mention) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Mention) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Mention) GetLanguage() *SoftwareName {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *Mention) GetUrl() *SoftwareName {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *Mention) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *Mention) GetMentionContextAttributes() *ContextAttributes {
	if x != nil {
		return x.MentionContextAttributes
	}
	return nil
}

func (x *Mention) GetDocumentContextAttributes() *ContextAttributes {
	if x != nil {
		return x.DocumentContextAttributes
	}
	return nil
}

func (x *Mention) GetReferences() []*Reference {
	if x != nil {
		return x.References
	}
	return nil
}

type SoftwareName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawForm        string         `protobuf:"bytes,1,opt,name=rawForm,proto3" json:"rawForm,omitempty"`
	NormalizedForm string         `protobuf:"bytes,2,opt,name=normalizedForm,proto3" json:"normalizedForm,omitempty"`
	OffsetStart    int32          `protobuf:"varint,3,opt,name=offsetStart,proto3" json:"offsetStart,omitempty"`
	OffsetEnd      int32          `protobuf:"varint,4,opt,name=offsetEnd,proto3" json:"offsetEnd,omitempty"`
	BoundingBoxes  []*BoundingBox `protobuf:"bytes,5,rep,name=boundingBoxes,proto3" json:"boundingBoxes,omitempty"`
}

func (x *SoftwareName) Reset() {
	*x = SoftwareName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoftwareName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoftwareName) ProtoMessage() {}

func (x *SoftwareName) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoftwareName.ProtoReflect.Descriptor instead.
func (*SoftwareName) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{3}
}

func (x *SoftwareName) GetRawForm() string {
	if x != nil {
		return x.RawForm
	}
	return ""
}

func (x *SoftwareName) GetNormalizedForm() string {
	if x != nil {
		return x.NormalizedForm
	}
	return ""
}

func (x *SoftwareName) GetOffsetStart() int32 {
	if x != nil {
		return x.OffsetStart
	}
	return 0
}

func (x *SoftwareName) GetOffsetEnd() int32 {
	if x != nil {
		return x.OffsetEnd
	}
	return 0
}

func (x *SoftwareName) GetBoundingBoxes() []*BoundingBox {
	if x != nil {
		return x.BoundingBoxes
	}
	return nil
}

type BoundingBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P int32   `protobuf:"varint,1,opt,name=p,proto3" json:"p,omitempty"`
	X float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	W float32 `protobuf:"fixed32,4,opt,name=w,proto3" json:"w,omitempty"`
	H float32 `protobuf:"fixed32,5,opt,name=h,proto3" json:"h,omitempty"`
}

func (x *BoundingBox) Reset() {
	*x = BoundingBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoundingBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundingBox) ProtoMessage() {}

func (x *BoundingBox) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundingBox.ProtoReflect.Descriptor instead.
func (*BoundingBox) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{4}
}

func (x *BoundingBox) GetP() int32 {
	if x != nil {
		return x.P
	}
	return 0
}

func (x *BoundingBox) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *BoundingBox) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *BoundingBox) GetW() float32 {
	if x != nil {
		return x.W
	}
	return 0
}

func (x *BoundingBox) GetH() float32 {
	if x != nil {
		return x.H
	}
	return 0
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawForm        string `protobuf:"bytes,1,opt,name=rawForm,proto3" json:"rawForm,omitempty"`
	NormalizedForm string `protobuf:"bytes,2,opt,name=normalizedForm,proto3" json:"normalizedForm,omitempty"`
	OffsetStart    int32  `protobuf:"varint,3,opt,name=offsetStart,proto3" json:"offsetStart,omitempty"`
	OffsetEnd      int32  `protobuf:"varint,4,opt,name=offsetEnd,proto3" json:"offsetEnd,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{5}
}

func (x *Version) GetRawForm() string {
	if x != nil {
		return x.RawForm
	}
	return ""
}

func (x *Version) GetNormalizedForm() string {
	if x != nil {
		return x.NormalizedForm
	}
	return ""
}

func (x *Version) GetOffsetStart() int32 {
	if x != nil {
		return x.OffsetStart
	}
	return 0
}

func (x *Version) GetOffsetEnd() int32 {
	if x != nil {
		return x.OffsetEnd
	}
	return 0
}

type MentionReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MentionReference) Reset() {
	*x = MentionReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MentionReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MentionReference) ProtoMessage() {}

func (x *MentionReference) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MentionReference.ProtoReflect.Descriptor instead.
func (*MentionReference) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{6}
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefKey int32  `protobuf:"varint,1,opt,name=refKey,proto3" json:"refKey,omitempty"`
	Tei    string `protobuf:"bytes,2,opt,name=tei,proto3" json:"tei,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{7}
}

func (x *Reference) GetRefKey() int32 {
	if x != nil {
		return x.RefKey
	}
	return 0
}

func (x *Reference) GetTei() string {
	if x != nil {
		return x.Tei
	}
	return ""
}

type ContextAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Used    *Attribute `protobuf:"bytes,1,opt,name=used,proto3" json:"used,omitempty"`
	Created *Attribute `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Shared  *Attribute `protobuf:"bytes,3,opt,name=shared,proto3" json:"shared,omitempty"`
}

func (x *ContextAttributes) Reset() {
	*x = ContextAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextAttributes) ProtoMessage() {}

func (x *ContextAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextAttributes.ProtoReflect.Descriptor instead.
func (*ContextAttributes) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{8}
}

func (x *ContextAttributes) GetUsed() *Attribute {
	if x != nil {
		return x.Used
	}
	return nil
}

func (x *ContextAttributes) GetCreated() *Attribute {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ContextAttributes) GetShared() *Attribute {
	if x != nil {
		return x.Shared
	}
	return nil
}

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{9}
}

func (x *Attribute) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

func (x *Attribute) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_software_mentions_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_papers_software_mentions_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_papers_software_mentions_proto_rawDescGZIP(), []int{10}
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_papers_software_mentions_proto protoreflect.FileDescriptor

var file_papers_software_mentions_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
	0x65, 0x2d, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfe, 0x02, 0x0a, 0x10, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x4d, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x24, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x46, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x57, 0x69, 0x64, 0x74, 0x68, 0x22, 0xa9, 0x03, 0x0a, 0x07, 0x4d, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0d, 0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53,
	0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0c, 0x73, 0x6f, 0x66,
	0x74, 0x77, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x6f, 0x66, 0x74,
	0x77, 0x61, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x4e, 0x0a, 0x18, 0x6d, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x18, 0x6d,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x19, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x19,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x0c, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61,
	0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x46, 0x6f, 0x72,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x46, 0x6f, 0x72, 0x6d,
	0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x46, 0x6f,
	0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x0d, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x52, 0x0d, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x0b,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x77, 0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x68, 0x22, 0x8b, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x61, 0x77, 0x46, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x61, 0x77, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x22,
	0x12, 0x0a, 0x10, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x35, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x66, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x69, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x69, 0x22, 0x7d, 0x0a, 0x11, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x22, 0x37, 0x0a, 0x09, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x0c,
	0x5a, 0x0a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_papers_software_mentions_proto_rawDescOnce sync.Once
	file_papers_software_mentions_proto_rawDescData = file_papers_software_mentions_proto_rawDesc
)

func file_papers_software_mentions_proto_rawDescGZIP() []byte {
	file_papers_software_mentions_proto_rawDescOnce.Do(func() {
		file_papers_software_mentions_proto_rawDescData = protoimpl.X.CompressGZIP(file_papers_software_mentions_proto_rawDescData)
	})
	return file_papers_software_mentions_proto_rawDescData
}

var file_papers_software_mentions_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_papers_software_mentions_proto_goTypes = []interface{}{
	(*SoftwareMentions)(nil),  // 0: SoftwareMentions
	(*Page)(nil),              // 1: Page
	(*Mention)(nil),           // 2: Mention
	(*SoftwareName)(nil),      // 3: SoftwareName
	(*BoundingBox)(nil),       // 4: BoundingBox
	(*Version)(nil),           // 5: Version
	(*MentionReference)(nil),  // 6: MentionReference
	(*Reference)(nil),         // 7: Reference
	(*ContextAttributes)(nil), // 8: ContextAttributes
	(*Attribute)(nil),         // 9: Attribute
	(*Metadata)(nil),          // 10: Metadata
}
var file_papers_software_mentions_proto_depIdxs = []int32{
	1,  // 0: SoftwareMentions.pages:type_name -> Page
	2,  // 1: SoftwareMentions.mentions:type_name -> Mention
	7,  // 2: SoftwareMentions.references:type_name -> Reference
	2,  // 3: SoftwareMentions.metadata:type_name -> Mention
	3,  // 4: Mention.software_name:type_name -> SoftwareName
	5,  // 5: Mention.version:type_name -> Version
	3,  // 6: Mention.language:type_name -> SoftwareName
	3,  // 7: Mention.url:type_name -> SoftwareName
	8,  // 8: Mention.mentionContextAttributes:type_name -> ContextAttributes
	8,  // 9: Mention.documentContextAttributes:type_name -> ContextAttributes
	7,  // 10: Mention.references:type_name -> Reference
	4,  // 11: SoftwareName.boundingBoxes:type_name -> BoundingBox
	9,  // 12: ContextAttributes.used:type_name -> Attribute
	9,  // 13: ContextAttributes.created:type_name -> Attribute
	9,  // 14: ContextAttributes.shared:type_name -> Attribute
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_papers_software_mentions_proto_init() }
func file_papers_software_mentions_proto_init() {
	if File_papers_software_mentions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_papers_software_mentions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoftwareMentions); i {
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
		file_papers_software_mentions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_papers_software_mentions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mention); i {
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
		file_papers_software_mentions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoftwareName); i {
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
		file_papers_software_mentions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoundingBox); i {
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
		file_papers_software_mentions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_papers_software_mentions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MentionReference); i {
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
		file_papers_software_mentions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_papers_software_mentions_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextAttributes); i {
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
		file_papers_software_mentions_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_papers_software_mentions_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_papers_software_mentions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_papers_software_mentions_proto_goTypes,
		DependencyIndexes: file_papers_software_mentions_proto_depIdxs,
		MessageInfos:      file_papers_software_mentions_proto_msgTypes,
	}.Build()
	File_papers_software_mentions_proto = out.File
	file_papers_software_mentions_proto_rawDesc = nil
	file_papers_software_mentions_proto_goTypes = nil
	file_papers_software_mentions_proto_depIdxs = nil
}
