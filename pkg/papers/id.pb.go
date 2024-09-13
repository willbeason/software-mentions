// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: papers/id.proto

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

type LicenseType int32

const (
	LicenseType_LICENSE_UNSPECIFIED                          LicenseType = 0
	LicenseType_LICENSE_CC_BY                                LicenseType = 1
	LicenseType_LICENSE_CC_BY_NC_ND                          LicenseType = 2
	LicenseType_LICENSE_CC_BY_NC                             LicenseType = 3
	LicenseType_LICENCE_ARXIV                                LicenseType = 4
	LicenseType_LICENSE_CC_BY_NC_SA                          LicenseType = 5
	LicenseType_LICENSE_CC_BY_SA                             LicenseType = 6
	LicenseType_LICENSE_CC0                                  LicenseType = 7
	LicenseType_LICENSE_ELSEVIER_SPECIFIC_OA_USER_LICENSE    LicenseType = 8
	LicenseType_LICENSE_IMPLIED_OA                           LicenseType = 9
	LicenseType_LICENSE_PUBLIC_DOMAIN                        LicenseType = 10
	LicenseType_LICENSE_NO_CC_CODE                           LicenseType = 11
	LicenseType_LICENSE_CC_BY_ND                             LicenseType = 12
	LicenseType_LICENSE_PUBLISHER_SPECIFIC_LICENSE           LicenseType = 13
	LicenseType_LICENSE_PUBLISHER_SPECIFIC_AUTHOR_MANUSCRIPT LicenseType = 14
	LicenseType_LICENSE_ACS_SPECIFIC_CHOICE_USAGE_AGREEMENT  LicenseType = 15
	LicenseType_LICENSE_OPEN_GOVERNMENT_LICENSE_CANADA       LicenseType = 16
)

// Enum value maps for LicenseType.
var (
	LicenseType_name = map[int32]string{
		0:  "LICENSE_UNSPECIFIED",
		1:  "LICENSE_CC_BY",
		2:  "LICENSE_CC_BY_NC_ND",
		3:  "LICENSE_CC_BY_NC",
		4:  "LICENCE_ARXIV",
		5:  "LICENSE_CC_BY_NC_SA",
		6:  "LICENSE_CC_BY_SA",
		7:  "LICENSE_CC0",
		8:  "LICENSE_ELSEVIER_SPECIFIC_OA_USER_LICENSE",
		9:  "LICENSE_IMPLIED_OA",
		10: "LICENSE_PUBLIC_DOMAIN",
		11: "LICENSE_NO_CC_CODE",
		12: "LICENSE_CC_BY_ND",
		13: "LICENSE_PUBLISHER_SPECIFIC_LICENSE",
		14: "LICENSE_PUBLISHER_SPECIFIC_AUTHOR_MANUSCRIPT",
		15: "LICENSE_ACS_SPECIFIC_CHOICE_USAGE_AGREEMENT",
		16: "LICENSE_OPEN_GOVERNMENT_LICENSE_CANADA",
	}
	LicenseType_value = map[string]int32{
		"LICENSE_UNSPECIFIED":                          0,
		"LICENSE_CC_BY":                                1,
		"LICENSE_CC_BY_NC_ND":                          2,
		"LICENSE_CC_BY_NC":                             3,
		"LICENCE_ARXIV":                                4,
		"LICENSE_CC_BY_NC_SA":                          5,
		"LICENSE_CC_BY_SA":                             6,
		"LICENSE_CC0":                                  7,
		"LICENSE_ELSEVIER_SPECIFIC_OA_USER_LICENSE":    8,
		"LICENSE_IMPLIED_OA":                           9,
		"LICENSE_PUBLIC_DOMAIN":                        10,
		"LICENSE_NO_CC_CODE":                           11,
		"LICENSE_CC_BY_ND":                             12,
		"LICENSE_PUBLISHER_SPECIFIC_LICENSE":           13,
		"LICENSE_PUBLISHER_SPECIFIC_AUTHOR_MANUSCRIPT": 14,
		"LICENSE_ACS_SPECIFIC_CHOICE_USAGE_AGREEMENT":  15,
		"LICENSE_OPEN_GOVERNMENT_LICENSE_CANADA":       16,
	}
)

func (x LicenseType) Enum() *LicenseType {
	p := new(LicenseType)
	*p = x
	return p
}

func (x LicenseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LicenseType) Descriptor() protoreflect.EnumDescriptor {
	return file_papers_id_proto_enumTypes[0].Descriptor()
}

func (LicenseType) Type() protoreflect.EnumType {
	return &file_papers_id_proto_enumTypes[0]
}

func (x LicenseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LicenseType.Descriptor instead.
func (LicenseType) EnumDescriptor() ([]byte, []int) {
	return file_papers_id_proto_rawDescGZIP(), []int{0}
}

type ResourceType int32

const (
	ResourceType_RESOURCE_UNSPECIFIED ResourceType = 0
	ResourceType_RESOURCE_JSON        ResourceType = 1
	ResourceType_RESOURCE_PDF         ResourceType = 2
	ResourceType_RESOURCE_LATEX       ResourceType = 3
	ResourceType_RESOURCE_XML         ResourceType = 4
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "RESOURCE_UNSPECIFIED",
		1: "RESOURCE_JSON",
		2: "RESOURCE_PDF",
		3: "RESOURCE_LATEX",
		4: "RESOURCE_XML",
	}
	ResourceType_value = map[string]int32{
		"RESOURCE_UNSPECIFIED": 0,
		"RESOURCE_JSON":        1,
		"RESOURCE_PDF":         2,
		"RESOURCE_LATEX":       3,
		"RESOURCE_XML":         4,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_papers_id_proto_enumTypes[1].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_papers_id_proto_enumTypes[1]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_papers_id_proto_rawDescGZIP(), []int{1}
}

type PaperId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // json = "id"
	Doi       string         `protobuf:"bytes,2,opt,name=doi,proto3" json:"doi,omitempty"`
	Arxiv     string         `protobuf:"bytes,3,opt,name=arxiv,proto3" json:"arxiv,omitempty"`
	PmId      string         `protobuf:"bytes,4,opt,name=pm_id,json=pmid,proto3" json:"pm_id,omitempty"`
	PmcId     string         `protobuf:"bytes,5,opt,name=pmc_id,json=pmcid,proto3" json:"pmc_id,omitempty"`
	IstexId   string         `protobuf:"bytes,6,opt,name=istex_id,json=istexId,proto3" json:"istex_id,omitempty"`
	Resources []ResourceType `protobuf:"varint,7,rep,packed,name=resources,proto3,enum=ResourceType" json:"resources,omitempty"` // json = "resources"
	License   LicenseType    `protobuf:"varint,8,opt,name=license,proto3,enum=LicenseType" json:"license,omitempty"`             // json = "license"
	OaLink    string         `protobuf:"bytes,9,opt,name=oa_link,json=oaLink,proto3" json:"oa_link,omitempty"`
}

func (x *PaperId) Reset() {
	*x = PaperId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_papers_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaperId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaperId) ProtoMessage() {}

func (x *PaperId) ProtoReflect() protoreflect.Message {
	mi := &file_papers_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaperId.ProtoReflect.Descriptor instead.
func (*PaperId) Descriptor() ([]byte, []int) {
	return file_papers_id_proto_rawDescGZIP(), []int{0}
}

func (x *PaperId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaperId) GetDoi() string {
	if x != nil {
		return x.Doi
	}
	return ""
}

func (x *PaperId) GetArxiv() string {
	if x != nil {
		return x.Arxiv
	}
	return ""
}

func (x *PaperId) GetPmId() string {
	if x != nil {
		return x.PmId
	}
	return ""
}

func (x *PaperId) GetPmcId() string {
	if x != nil {
		return x.PmcId
	}
	return ""
}

func (x *PaperId) GetIstexId() string {
	if x != nil {
		return x.IstexId
	}
	return ""
}

func (x *PaperId) GetResources() []ResourceType {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *PaperId) GetLicense() LicenseType {
	if x != nil {
		return x.License
	}
	return LicenseType_LICENSE_UNSPECIFIED
}

func (x *PaperId) GetOaLink() string {
	if x != nil {
		return x.OaLink
	}
	return ""
}

var File_papers_id_proto protoreflect.FileDescriptor

var file_papers_id_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x6f, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x69, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x72, 0x78, 0x69, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x72, 0x78, 0x69, 0x76, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6d, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x6d,
	0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6d, 0x63, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x74, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x74, 0x65, 0x78, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x6c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4c, 0x69, 0x63,
	0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x61, 0x4c, 0x69, 0x6e, 0x6b, 0x2a, 0x82, 0x04, 0x0a, 0x0b, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49,
	0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x43,
	0x43, 0x5f, 0x42, 0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x5f, 0x43, 0x43, 0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x43, 0x5f, 0x4e, 0x44, 0x10, 0x02, 0x12,
	0x14, 0x0a, 0x10, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x43, 0x5f, 0x42, 0x59,
	0x5f, 0x4e, 0x43, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x43, 0x45,
	0x5f, 0x41, 0x52, 0x58, 0x49, 0x56, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49, 0x43, 0x45,
	0x4e, 0x53, 0x45, 0x5f, 0x43, 0x43, 0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x43, 0x5f, 0x53, 0x41, 0x10,
	0x05, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x43, 0x5f,
	0x42, 0x59, 0x5f, 0x53, 0x41, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x49, 0x43, 0x45, 0x4e,
	0x53, 0x45, 0x5f, 0x43, 0x43, 0x30, 0x10, 0x07, 0x12, 0x2d, 0x0a, 0x29, 0x4c, 0x49, 0x43, 0x45,
	0x4e, 0x53, 0x45, 0x5f, 0x45, 0x4c, 0x53, 0x45, 0x56, 0x49, 0x45, 0x52, 0x5f, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x43, 0x5f, 0x4f, 0x41, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49,
	0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x43, 0x45, 0x4e,
	0x53, 0x45, 0x5f, 0x49, 0x4d, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x5f, 0x4f, 0x41, 0x10, 0x09, 0x12,
	0x19, 0x0a, 0x15, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49,
	0x43, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49,
	0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x43, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x43,
	0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x44, 0x10, 0x0c, 0x12, 0x26, 0x0a, 0x22, 0x4c, 0x49, 0x43, 0x45,
	0x4e, 0x53, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x52, 0x5f, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x0d,
	0x12, 0x30, 0x0a, 0x2c, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x53, 0x48, 0x45, 0x52, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x5f, 0x41,
	0x55, 0x54, 0x48, 0x4f, 0x52, 0x5f, 0x4d, 0x41, 0x4e, 0x55, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54,
	0x10, 0x0e, 0x12, 0x2f, 0x0a, 0x2b, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x41, 0x43,
	0x53, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x5f, 0x43, 0x48, 0x4f, 0x49, 0x43,
	0x45, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x47, 0x52, 0x45, 0x45, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x0f, 0x12, 0x2a, 0x0a, 0x26, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x4f,
	0x50, 0x45, 0x4e, 0x5f, 0x47, 0x4f, 0x56, 0x45, 0x52, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4c,
	0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x41, 0x44, 0x41, 0x10, 0x10, 0x2a,
	0x73, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x50, 0x44, 0x46, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x58,
	0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x58,
	0x4d, 0x4c, 0x10, 0x04, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_papers_id_proto_rawDescOnce sync.Once
	file_papers_id_proto_rawDescData = file_papers_id_proto_rawDesc
)

func file_papers_id_proto_rawDescGZIP() []byte {
	file_papers_id_proto_rawDescOnce.Do(func() {
		file_papers_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_papers_id_proto_rawDescData)
	})
	return file_papers_id_proto_rawDescData
}

var file_papers_id_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_papers_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_papers_id_proto_goTypes = []any{
	(LicenseType)(0),  // 0: LicenseType
	(ResourceType)(0), // 1: ResourceType
	(*PaperId)(nil),   // 2: PaperId
}
var file_papers_id_proto_depIdxs = []int32{
	1, // 0: PaperId.resources:type_name -> ResourceType
	0, // 1: PaperId.license:type_name -> LicenseType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_papers_id_proto_init() }
func file_papers_id_proto_init() {
	if File_papers_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_papers_id_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PaperId); i {
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
			RawDescriptor: file_papers_id_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_papers_id_proto_goTypes,
		DependencyIndexes: file_papers_id_proto_depIdxs,
		EnumInfos:         file_papers_id_proto_enumTypes,
		MessageInfos:      file_papers_id_proto_msgTypes,
	}.Build()
	File_papers_id_proto = out.File
	file_papers_id_proto_rawDesc = nil
	file_papers_id_proto_goTypes = nil
	file_papers_id_proto_depIdxs = nil
}
