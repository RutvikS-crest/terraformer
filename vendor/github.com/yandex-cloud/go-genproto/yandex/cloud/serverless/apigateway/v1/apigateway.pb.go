// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/serverless/apigateway/v1/apigateway.proto

package apigateway

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ApiGateway_Status int32

const (
	ApiGateway_STATUS_UNSPECIFIED ApiGateway_Status = 0
	// API gateway is being created.
	ApiGateway_CREATING ApiGateway_Status = 1
	// API gateway is ready for use.
	ApiGateway_ACTIVE ApiGateway_Status = 2
	// API gateway is being deleted.
	ApiGateway_DELETING ApiGateway_Status = 3
	// API gateway failed. The only allowed action is delete.
	ApiGateway_ERROR ApiGateway_Status = 4
	// API gateway is being updated.
	ApiGateway_UPDATING ApiGateway_Status = 5
)

// Enum value maps for ApiGateway_Status.
var (
	ApiGateway_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
		4: "ERROR",
		5: "UPDATING",
	}
	ApiGateway_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"DELETING":           3,
		"ERROR":              4,
		"UPDATING":           5,
	}
)

func (x ApiGateway_Status) Enum() *ApiGateway_Status {
	p := new(ApiGateway_Status)
	*p = x
	return p
}

func (x ApiGateway_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiGateway_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_enumTypes[0].Descriptor()
}

func (ApiGateway_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_enumTypes[0]
}

func (x ApiGateway_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiGateway_Status.Descriptor instead.
func (ApiGateway_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescGZIP(), []int{0, 0}
}

type ApiGateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the API gateway. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the API gateway belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp for the API-gateway.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the API gateway. The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the API gateway.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// API gateway labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status of the API gateway.
	Status ApiGateway_Status `protobuf:"varint,8,opt,name=status,proto3,enum=yandex.cloud.serverless.apigateway.v1.ApiGateway_Status" json:"status,omitempty"`
	// Default domain for the API gateway. Generated at creation time.
	Domain string `protobuf:"bytes,9,opt,name=domain,proto3" json:"domain,omitempty"`
	// ID of the log group for the API gateway.
	LogGroupId string `protobuf:"bytes,10,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	// List of domains attached to API gateway.
	AttachedDomains []*AttachedDomain `protobuf:"bytes,11,rep,name=attached_domains,json=attachedDomains,proto3" json:"attached_domains,omitempty"`
	// Network access. If specified the gateway will be attached to specified network/subnet(s).
	Connectivity *Connectivity `protobuf:"bytes,12,opt,name=connectivity,proto3" json:"connectivity,omitempty"`
}

func (x *ApiGateway) Reset() {
	*x = ApiGateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiGateway) ProtoMessage() {}

func (x *ApiGateway) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiGateway.ProtoReflect.Descriptor instead.
func (*ApiGateway) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescGZIP(), []int{0}
}

func (x *ApiGateway) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApiGateway) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ApiGateway) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApiGateway) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiGateway) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApiGateway) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ApiGateway) GetStatus() ApiGateway_Status {
	if x != nil {
		return x.Status
	}
	return ApiGateway_STATUS_UNSPECIFIED
}

func (x *ApiGateway) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ApiGateway) GetLogGroupId() string {
	if x != nil {
		return x.LogGroupId
	}
	return ""
}

func (x *ApiGateway) GetAttachedDomains() []*AttachedDomain {
	if x != nil {
		return x.AttachedDomains
	}
	return nil
}

func (x *ApiGateway) GetConnectivity() *Connectivity {
	if x != nil {
		return x.Connectivity
	}
	return nil
}

type AttachedDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the domain.
	DomainId string `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// ID of the domain certificate.
	CertificateId string `protobuf:"bytes,2,opt,name=certificate_id,json=certificateId,proto3" json:"certificate_id,omitempty"`
	// Enabling flag.
	Enabled bool `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Name of the domain.
	Domain string `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *AttachedDomain) Reset() {
	*x = AttachedDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachedDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachedDomain) ProtoMessage() {}

func (x *AttachedDomain) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachedDomain.ProtoReflect.Descriptor instead.
func (*AttachedDomain) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescGZIP(), []int{1}
}

func (x *AttachedDomain) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *AttachedDomain) GetCertificateId() string {
	if x != nil {
		return x.CertificateId
	}
	return ""
}

func (x *AttachedDomain) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AttachedDomain) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// Gateway connectivity specification.
type Connectivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Network the gateway will have access to.
	// It's essential to specify network with subnets in all availability zones.
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Complete list of subnets (from the same network) the gateway can be attached to.
	// It's essential to specify at least one subnet for each availability zones.
	SubnetId []string `protobuf:"bytes,2,rep,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
}

func (x *Connectivity) Reset() {
	*x = Connectivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connectivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connectivity) ProtoMessage() {}

func (x *Connectivity) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connectivity.ProtoReflect.Descriptor instead.
func (*Connectivity) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescGZIP(), []int{2}
}

func (x *Connectivity) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Connectivity) GetSubnetId() []string {
	if x != nil {
		return x.SubnetId
	}
	return nil
}

var File_yandex_cloud_serverless_apigateway_v1_apigateway_proto protoreflect.FileDescriptor

var file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDesc = []byte{
	0x0a, 0x36, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe6, 0x05, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x50, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x69,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x20,
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x60, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x61, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x22, 0x86, 0x01, 0x0a, 0x0e, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x4a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x42, 0x81,
	0x01, 0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x54, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescOnce sync.Once
	file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescData = file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDesc
)

func file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescGZIP() []byte {
	file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescData)
	})
	return file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDescData
}

var file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_goTypes = []interface{}{
	(ApiGateway_Status)(0),        // 0: yandex.cloud.serverless.apigateway.v1.ApiGateway.Status
	(*ApiGateway)(nil),            // 1: yandex.cloud.serverless.apigateway.v1.ApiGateway
	(*AttachedDomain)(nil),        // 2: yandex.cloud.serverless.apigateway.v1.AttachedDomain
	(*Connectivity)(nil),          // 3: yandex.cloud.serverless.apigateway.v1.Connectivity
	nil,                           // 4: yandex.cloud.serverless.apigateway.v1.ApiGateway.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.serverless.apigateway.v1.ApiGateway.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: yandex.cloud.serverless.apigateway.v1.ApiGateway.labels:type_name -> yandex.cloud.serverless.apigateway.v1.ApiGateway.LabelsEntry
	0, // 2: yandex.cloud.serverless.apigateway.v1.ApiGateway.status:type_name -> yandex.cloud.serverless.apigateway.v1.ApiGateway.Status
	2, // 3: yandex.cloud.serverless.apigateway.v1.ApiGateway.attached_domains:type_name -> yandex.cloud.serverless.apigateway.v1.AttachedDomain
	3, // 4: yandex.cloud.serverless.apigateway.v1.ApiGateway.connectivity:type_name -> yandex.cloud.serverless.apigateway.v1.Connectivity
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_init() }
func file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_init() {
	if File_yandex_cloud_serverless_apigateway_v1_apigateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiGateway); i {
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
		file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachedDomain); i {
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
		file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connectivity); i {
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
			RawDescriptor: file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_msgTypes,
	}.Build()
	File_yandex_cloud_serverless_apigateway_v1_apigateway_proto = out.File
	file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_rawDesc = nil
	file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_goTypes = nil
	file_yandex_cloud_serverless_apigateway_v1_apigateway_proto_depIdxs = nil
}
