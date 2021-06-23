// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/project_domain_attributes.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Defines a set of custom matching attributes which defines resource defaults for a project and domain.
// For more info on matchable attributes, see - :ref:`ref_flyteidl/admin/matchable_resource.proto`.
type ProjectDomainAttributes struct {
	// Unique project id for which this set of attributes will be applied.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id for which this set of attributes will be applied.
	Domain               string              `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	MatchingAttributes   *MatchingAttributes `protobuf:"bytes,3,opt,name=matching_attributes,json=matchingAttributes,proto3" json:"matching_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProjectDomainAttributes) Reset()         { *m = ProjectDomainAttributes{} }
func (m *ProjectDomainAttributes) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributes) ProtoMessage()    {}
func (*ProjectDomainAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{0}
}

func (m *ProjectDomainAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributes.Unmarshal(m, b)
}
func (m *ProjectDomainAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributes.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributes.Merge(m, src)
}
func (m *ProjectDomainAttributes) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributes.Size(m)
}
func (m *ProjectDomainAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributes proto.InternalMessageInfo

func (m *ProjectDomainAttributes) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectDomainAttributes) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ProjectDomainAttributes) GetMatchingAttributes() *MatchingAttributes {
	if m != nil {
		return m.MatchingAttributes
	}
	return nil
}

// Sets custom attributes for a project-domain combination.
type ProjectDomainAttributesUpdateRequest struct {
	// +required
	Attributes           *ProjectDomainAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ProjectDomainAttributesUpdateRequest) Reset()         { *m = ProjectDomainAttributesUpdateRequest{} }
func (m *ProjectDomainAttributesUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesUpdateRequest) ProtoMessage()    {}
func (*ProjectDomainAttributesUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{1}
}

func (m *ProjectDomainAttributesUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Merge(m, src)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Size(m)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesUpdateRequest proto.InternalMessageInfo

func (m *ProjectDomainAttributesUpdateRequest) GetAttributes() *ProjectDomainAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Purposefully empty, may be populated in the future.
type ProjectDomainAttributesUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectDomainAttributesUpdateResponse) Reset()         { *m = ProjectDomainAttributesUpdateResponse{} }
func (m *ProjectDomainAttributesUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesUpdateResponse) ProtoMessage()    {}
func (*ProjectDomainAttributesUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{2}
}

func (m *ProjectDomainAttributesUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Merge(m, src)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Size(m)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesUpdateResponse proto.InternalMessageInfo

// Request to get an individual project domain attribute override.
type ProjectDomainAttributesGetRequest struct {
	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Which type of matchable attributes to return.
	// +required
	ResourceType         MatchableResource `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProjectDomainAttributesGetRequest) Reset()         { *m = ProjectDomainAttributesGetRequest{} }
func (m *ProjectDomainAttributesGetRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesGetRequest) ProtoMessage()    {}
func (*ProjectDomainAttributesGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{3}
}

func (m *ProjectDomainAttributesGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesGetRequest.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesGetRequest.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesGetRequest.Merge(m, src)
}
func (m *ProjectDomainAttributesGetRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesGetRequest.Size(m)
}
func (m *ProjectDomainAttributesGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesGetRequest proto.InternalMessageInfo

func (m *ProjectDomainAttributesGetRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectDomainAttributesGetRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ProjectDomainAttributesGetRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Response to get an individual project domain attribute override.
type ProjectDomainAttributesGetResponse struct {
	Attributes           *ProjectDomainAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ProjectDomainAttributesGetResponse) Reset()         { *m = ProjectDomainAttributesGetResponse{} }
func (m *ProjectDomainAttributesGetResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesGetResponse) ProtoMessage()    {}
func (*ProjectDomainAttributesGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{4}
}

func (m *ProjectDomainAttributesGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesGetResponse.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesGetResponse.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesGetResponse.Merge(m, src)
}
func (m *ProjectDomainAttributesGetResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesGetResponse.Size(m)
}
func (m *ProjectDomainAttributesGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesGetResponse proto.InternalMessageInfo

func (m *ProjectDomainAttributesGetResponse) GetAttributes() *ProjectDomainAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Request to delete a set matchable project domain attribute override.
type ProjectDomainAttributesDeleteRequest struct {
	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Which type of matchable attributes to delete.
	// +required
	ResourceType         MatchableResource `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProjectDomainAttributesDeleteRequest) Reset()         { *m = ProjectDomainAttributesDeleteRequest{} }
func (m *ProjectDomainAttributesDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesDeleteRequest) ProtoMessage()    {}
func (*ProjectDomainAttributesDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{5}
}

func (m *ProjectDomainAttributesDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Merge(m, src)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Size(m)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesDeleteRequest proto.InternalMessageInfo

func (m *ProjectDomainAttributesDeleteRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectDomainAttributesDeleteRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ProjectDomainAttributesDeleteRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Purposefully empty, may be populated in the future.
type ProjectDomainAttributesDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectDomainAttributesDeleteResponse) Reset()         { *m = ProjectDomainAttributesDeleteResponse{} }
func (m *ProjectDomainAttributesDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesDeleteResponse) ProtoMessage()    {}
func (*ProjectDomainAttributesDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{6}
}

func (m *ProjectDomainAttributesDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Merge(m, src)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Size(m)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProjectDomainAttributes)(nil), "flyteidl.admin.ProjectDomainAttributes")
	proto.RegisterType((*ProjectDomainAttributesUpdateRequest)(nil), "flyteidl.admin.ProjectDomainAttributesUpdateRequest")
	proto.RegisterType((*ProjectDomainAttributesUpdateResponse)(nil), "flyteidl.admin.ProjectDomainAttributesUpdateResponse")
	proto.RegisterType((*ProjectDomainAttributesGetRequest)(nil), "flyteidl.admin.ProjectDomainAttributesGetRequest")
	proto.RegisterType((*ProjectDomainAttributesGetResponse)(nil), "flyteidl.admin.ProjectDomainAttributesGetResponse")
	proto.RegisterType((*ProjectDomainAttributesDeleteRequest)(nil), "flyteidl.admin.ProjectDomainAttributesDeleteRequest")
	proto.RegisterType((*ProjectDomainAttributesDeleteResponse)(nil), "flyteidl.admin.ProjectDomainAttributesDeleteResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/project_domain_attributes.proto", fileDescriptor_e8ab0b551a649f05)
}

var fileDescriptor_e8ab0b551a649f05 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x65, 0x7f, 0x3f, 0xa8, 0x38, 0x6a, 0x0f, 0x11, 0x34, 0x78, 0x6a, 0x17, 0xa5, 0xbd, 0x98,
	0x40, 0x45, 0x3c, 0x2b, 0xc5, 0x9e, 0x04, 0x89, 0x7a, 0xf1, 0x12, 0xf2, 0x67, 0x4c, 0x23, 0x49,
	0x76, 0xdd, 0x4c, 0x0e, 0xfd, 0x30, 0x42, 0x3f, 0xaa, 0xb8, 0xc9, 0xf6, 0x9f, 0x46, 0x11, 0x04,
	0x8f, 0x93, 0xbc, 0x79, 0xef, 0xed, 0x7b, 0x0c, 0x38, 0x4f, 0xd9, 0x8c, 0x30, 0x8d, 0x33, 0x37,
	0x88, 0xf3, 0xb4, 0x70, 0xa5, 0x12, 0xcf, 0x18, 0x91, 0x1f, 0x8b, 0x3c, 0x48, 0x0b, 0x3f, 0x20,
	0x52, 0x69, 0x58, 0x11, 0x96, 0x8e, 0x54, 0x82, 0x84, 0xd5, 0x35, 0x78, 0x47, 0xe3, 0x8f, 0x06,
	0x1b, 0xfb, 0x79, 0x40, 0xd1, 0x34, 0x08, 0x33, 0xf4, 0x15, 0x96, 0xa2, 0x52, 0x11, 0xd6, 0x8b,
	0x7c, 0xce, 0xe0, 0xf0, 0xb6, 0x26, 0x1f, 0x6b, 0xee, 0xcb, 0x05, 0xb5, 0x65, 0xc3, 0x56, 0xa3,
	0x6b, 0xb3, 0x1e, 0x1b, 0x6e, 0x7b, 0x66, 0xb4, 0x0e, 0xa0, 0x53, 0x3b, 0xb1, 0xff, 0xe9, 0x1f,
	0xcd, 0x64, 0xdd, 0xc1, 0xbe, 0x56, 0x4a, 0x8b, 0x64, 0xc5, 0xa3, 0xfd, 0xbf, 0xc7, 0x86, 0x3b,
	0x23, 0xee, 0xac, 0x9b, 0x74, 0x6e, 0x1a, 0xe8, 0x52, 0xd2, 0xb3, 0xf2, 0x0f, 0xdf, 0xb8, 0x80,
	0xe3, 0x16, 0x87, 0x0f, 0x32, 0x0e, 0x08, 0x3d, 0x7c, 0xa9, 0xb0, 0x24, 0x6b, 0x02, 0xb0, 0xa2,
	0xc9, 0xb4, 0xe6, 0x60, 0x53, 0xb3, 0x85, 0xc9, 0x5b, 0x59, 0xe5, 0x03, 0x38, 0xf9, 0x46, 0xb0,
	0x94, 0xa2, 0x28, 0x91, 0xbf, 0x32, 0xe8, 0xb7, 0x20, 0x27, 0x48, 0xc6, 0xd7, 0xcf, 0x63, 0xbc,
	0x86, 0x3d, 0x53, 0x93, 0x4f, 0x33, 0x89, 0x3a, 0xc0, 0xee, 0xa8, 0xff, 0x69, 0x80, 0xef, 0xad,
	0x7a, 0x0d, 0xda, 0xdb, 0x35, 0x7b, 0xf7, 0x33, 0x89, 0x3c, 0x07, 0xfe, 0x95, 0xbd, 0xfa, 0x15,
	0xbf, 0x97, 0xdb, 0x9c, 0xb5, 0x36, 0x35, 0xc6, 0x0c, 0x97, 0x4d, 0xfd, 0x5d, 0x22, 0xed, 0xd5,
	0x1a, 0x87, 0x75, 0x28, 0x57, 0x17, 0x8f, 0xe7, 0x49, 0x4a, 0xd3, 0x2a, 0x74, 0x22, 0x91, 0xbb,
	0x5a, 0x45, 0xa8, 0xc4, 0x5d, 0x9c, 0x55, 0x82, 0x85, 0x2b, 0xc3, 0xd3, 0x44, 0xb8, 0xeb, 0x97,
	0x16, 0x76, 0xf4, 0x5d, 0x9d, 0xbd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x77, 0x7d, 0xa6, 0xc2,
	0x03, 0x00, 0x00,
}
