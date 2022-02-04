// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/redis/v1/maintenance.proto

package redis

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type WeeklyMaintenanceWindow_WeekDay int32

const (
	WeeklyMaintenanceWindow_WEEK_DAY_UNSPECIFIED WeeklyMaintenanceWindow_WeekDay = 0
	WeeklyMaintenanceWindow_MON                  WeeklyMaintenanceWindow_WeekDay = 1
	WeeklyMaintenanceWindow_TUE                  WeeklyMaintenanceWindow_WeekDay = 2
	WeeklyMaintenanceWindow_WED                  WeeklyMaintenanceWindow_WeekDay = 3
	WeeklyMaintenanceWindow_THU                  WeeklyMaintenanceWindow_WeekDay = 4
	WeeklyMaintenanceWindow_FRI                  WeeklyMaintenanceWindow_WeekDay = 5
	WeeklyMaintenanceWindow_SAT                  WeeklyMaintenanceWindow_WeekDay = 6
	WeeklyMaintenanceWindow_SUN                  WeeklyMaintenanceWindow_WeekDay = 7
)

var WeeklyMaintenanceWindow_WeekDay_name = map[int32]string{
	0: "WEEK_DAY_UNSPECIFIED",
	1: "MON",
	2: "TUE",
	3: "WED",
	4: "THU",
	5: "FRI",
	6: "SAT",
	7: "SUN",
}

var WeeklyMaintenanceWindow_WeekDay_value = map[string]int32{
	"WEEK_DAY_UNSPECIFIED": 0,
	"MON":                  1,
	"TUE":                  2,
	"WED":                  3,
	"THU":                  4,
	"FRI":                  5,
	"SAT":                  6,
	"SUN":                  7,
}

func (x WeeklyMaintenanceWindow_WeekDay) String() string {
	return proto.EnumName(WeeklyMaintenanceWindow_WeekDay_name, int32(x))
}

func (WeeklyMaintenanceWindow_WeekDay) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d9e0e3353f6a6e54, []int{2, 0}
}

type MaintenanceWindow struct {
	// Types that are valid to be assigned to Policy:
	//	*MaintenanceWindow_Anytime
	//	*MaintenanceWindow_WeeklyMaintenanceWindow
	Policy               isMaintenanceWindow_Policy `protobuf_oneof:"policy"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MaintenanceWindow) Reset()         { *m = MaintenanceWindow{} }
func (m *MaintenanceWindow) String() string { return proto.CompactTextString(m) }
func (*MaintenanceWindow) ProtoMessage()    {}
func (*MaintenanceWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e0e3353f6a6e54, []int{0}
}

func (m *MaintenanceWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintenanceWindow.Unmarshal(m, b)
}
func (m *MaintenanceWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintenanceWindow.Marshal(b, m, deterministic)
}
func (m *MaintenanceWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintenanceWindow.Merge(m, src)
}
func (m *MaintenanceWindow) XXX_Size() int {
	return xxx_messageInfo_MaintenanceWindow.Size(m)
}
func (m *MaintenanceWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintenanceWindow.DiscardUnknown(m)
}

var xxx_messageInfo_MaintenanceWindow proto.InternalMessageInfo

type isMaintenanceWindow_Policy interface {
	isMaintenanceWindow_Policy()
}

type MaintenanceWindow_Anytime struct {
	Anytime *AnytimeMaintenanceWindow `protobuf:"bytes,1,opt,name=anytime,proto3,oneof"`
}

type MaintenanceWindow_WeeklyMaintenanceWindow struct {
	WeeklyMaintenanceWindow *WeeklyMaintenanceWindow `protobuf:"bytes,2,opt,name=weekly_maintenance_window,json=weeklyMaintenanceWindow,proto3,oneof"`
}

func (*MaintenanceWindow_Anytime) isMaintenanceWindow_Policy() {}

func (*MaintenanceWindow_WeeklyMaintenanceWindow) isMaintenanceWindow_Policy() {}

func (m *MaintenanceWindow) GetPolicy() isMaintenanceWindow_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *MaintenanceWindow) GetAnytime() *AnytimeMaintenanceWindow {
	if x, ok := m.GetPolicy().(*MaintenanceWindow_Anytime); ok {
		return x.Anytime
	}
	return nil
}

func (m *MaintenanceWindow) GetWeeklyMaintenanceWindow() *WeeklyMaintenanceWindow {
	if x, ok := m.GetPolicy().(*MaintenanceWindow_WeeklyMaintenanceWindow); ok {
		return x.WeeklyMaintenanceWindow
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MaintenanceWindow) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MaintenanceWindow_Anytime)(nil),
		(*MaintenanceWindow_WeeklyMaintenanceWindow)(nil),
	}
}

type AnytimeMaintenanceWindow struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnytimeMaintenanceWindow) Reset()         { *m = AnytimeMaintenanceWindow{} }
func (m *AnytimeMaintenanceWindow) String() string { return proto.CompactTextString(m) }
func (*AnytimeMaintenanceWindow) ProtoMessage()    {}
func (*AnytimeMaintenanceWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e0e3353f6a6e54, []int{1}
}

func (m *AnytimeMaintenanceWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnytimeMaintenanceWindow.Unmarshal(m, b)
}
func (m *AnytimeMaintenanceWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnytimeMaintenanceWindow.Marshal(b, m, deterministic)
}
func (m *AnytimeMaintenanceWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnytimeMaintenanceWindow.Merge(m, src)
}
func (m *AnytimeMaintenanceWindow) XXX_Size() int {
	return xxx_messageInfo_AnytimeMaintenanceWindow.Size(m)
}
func (m *AnytimeMaintenanceWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_AnytimeMaintenanceWindow.DiscardUnknown(m)
}

var xxx_messageInfo_AnytimeMaintenanceWindow proto.InternalMessageInfo

type WeeklyMaintenanceWindow struct {
	Day WeeklyMaintenanceWindow_WeekDay `protobuf:"varint,1,opt,name=day,proto3,enum=yandex.cloud.mdb.redis.v1.WeeklyMaintenanceWindow_WeekDay" json:"day,omitempty"`
	// Hour of the day in UTC.
	Hour                 int64    `protobuf:"varint,2,opt,name=hour,proto3" json:"hour,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeeklyMaintenanceWindow) Reset()         { *m = WeeklyMaintenanceWindow{} }
func (m *WeeklyMaintenanceWindow) String() string { return proto.CompactTextString(m) }
func (*WeeklyMaintenanceWindow) ProtoMessage()    {}
func (*WeeklyMaintenanceWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e0e3353f6a6e54, []int{2}
}

func (m *WeeklyMaintenanceWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeeklyMaintenanceWindow.Unmarshal(m, b)
}
func (m *WeeklyMaintenanceWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeeklyMaintenanceWindow.Marshal(b, m, deterministic)
}
func (m *WeeklyMaintenanceWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeeklyMaintenanceWindow.Merge(m, src)
}
func (m *WeeklyMaintenanceWindow) XXX_Size() int {
	return xxx_messageInfo_WeeklyMaintenanceWindow.Size(m)
}
func (m *WeeklyMaintenanceWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_WeeklyMaintenanceWindow.DiscardUnknown(m)
}

var xxx_messageInfo_WeeklyMaintenanceWindow proto.InternalMessageInfo

func (m *WeeklyMaintenanceWindow) GetDay() WeeklyMaintenanceWindow_WeekDay {
	if m != nil {
		return m.Day
	}
	return WeeklyMaintenanceWindow_WEEK_DAY_UNSPECIFIED
}

func (m *WeeklyMaintenanceWindow) GetHour() int64 {
	if m != nil {
		return m.Hour
	}
	return 0
}

type MaintenanceOperation struct {
	Info                 string               `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	DelayedUntil         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=delayed_until,json=delayedUntil,proto3" json:"delayed_until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MaintenanceOperation) Reset()         { *m = MaintenanceOperation{} }
func (m *MaintenanceOperation) String() string { return proto.CompactTextString(m) }
func (*MaintenanceOperation) ProtoMessage()    {}
func (*MaintenanceOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9e0e3353f6a6e54, []int{3}
}

func (m *MaintenanceOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaintenanceOperation.Unmarshal(m, b)
}
func (m *MaintenanceOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaintenanceOperation.Marshal(b, m, deterministic)
}
func (m *MaintenanceOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaintenanceOperation.Merge(m, src)
}
func (m *MaintenanceOperation) XXX_Size() int {
	return xxx_messageInfo_MaintenanceOperation.Size(m)
}
func (m *MaintenanceOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_MaintenanceOperation.DiscardUnknown(m)
}

var xxx_messageInfo_MaintenanceOperation proto.InternalMessageInfo

func (m *MaintenanceOperation) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *MaintenanceOperation) GetDelayedUntil() *timestamp.Timestamp {
	if m != nil {
		return m.DelayedUntil
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1.WeeklyMaintenanceWindow_WeekDay", WeeklyMaintenanceWindow_WeekDay_name, WeeklyMaintenanceWindow_WeekDay_value)
	proto.RegisterType((*MaintenanceWindow)(nil), "yandex.cloud.mdb.redis.v1.MaintenanceWindow")
	proto.RegisterType((*AnytimeMaintenanceWindow)(nil), "yandex.cloud.mdb.redis.v1.AnytimeMaintenanceWindow")
	proto.RegisterType((*WeeklyMaintenanceWindow)(nil), "yandex.cloud.mdb.redis.v1.WeeklyMaintenanceWindow")
	proto.RegisterType((*MaintenanceOperation)(nil), "yandex.cloud.mdb.redis.v1.MaintenanceOperation")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/redis/v1/maintenance.proto", fileDescriptor_d9e0e3353f6a6e54)
}

var fileDescriptor_d9e0e3353f6a6e54 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0xc6, 0x4d, 0xe8, 0x02, 0x95, 0x59, 0x55, 0x6a, 0x1a, 0x11, 0x81, 0x7c, 0x42,
	0x42, 0x59, 0xcb, 0x2e, 0x70, 0xe0, 0x8f, 0x50, 0xd2, 0xb8, 0x6a, 0x04, 0x4d, 0x90, 0x1b, 0x2b,
	0x82, 0x8b, 0xb5, 0xce, 0x6e, 0xdd, 0x15, 0xf6, 0xae, 0x95, 0xda, 0x09, 0x7e, 0x05, 0x9e, 0x0a,
	0x4e, 0xf0, 0x24, 0x48, 0xbc, 0x02, 0x27, 0xe4, 0x71, 0x22, 0x1a, 0xd1, 0x20, 0x71, 0xfb, 0x34,
	0xfb, 0x9b, 0xf9, 0x66, 0xc6, 0x63, 0xf4, 0xb8, 0xa0, 0x92, 0xf1, 0x4f, 0xd6, 0x34, 0x56, 0x39,
	0xb3, 0x12, 0x16, 0x5a, 0x33, 0xce, 0xc4, 0x95, 0x35, 0xb7, 0xad, 0x84, 0x0a, 0x99, 0x71, 0x49,
	0xe5, 0x94, 0x93, 0x74, 0xa6, 0x32, 0x85, 0x0f, 0x2b, 0x98, 0x00, 0x4c, 0x12, 0x16, 0x12, 0x80,
	0xc9, 0xdc, 0x6e, 0x3d, 0x88, 0x94, 0x8a, 0x62, 0x6e, 0x01, 0x18, 0xe6, 0x17, 0x56, 0x26, 0x12,
	0x7e, 0x95, 0xd1, 0x24, 0xad, 0x72, 0x5b, 0xed, 0x35, 0xa3, 0x39, 0x8d, 0x05, 0xa3, 0x99, 0x50,
	0xb2, 0x7a, 0x36, 0x7f, 0x68, 0xe8, 0xde, 0xd9, 0x1f, 0xc3, 0x89, 0x90, 0x4c, 0x2d, 0xf0, 0x08,
	0x35, 0xa8, 0x2c, 0xca, 0x52, 0x4d, 0xed, 0xa1, 0xf6, 0xe8, 0xb6, 0x73, 0x44, 0x36, 0xb6, 0x40,
	0xba, 0x15, 0xf9, 0x57, 0x95, 0xd3, 0x2d, 0x6f, 0x55, 0x05, 0xa7, 0xe8, 0x70, 0xc1, 0xf9, 0xc7,
	0xb8, 0x08, 0xae, 0x4d, 0x17, 0x2c, 0x80, 0x6b, 0x6e, 0x83, 0x85, 0xf3, 0x0f, 0x8b, 0x09, 0xe4,
	0xde, 0xe4, 0x70, 0xb0, 0xb8, 0xf9, 0xa9, 0xb7, 0x87, 0xea, 0xa9, 0x8a, 0xc5, 0xb4, 0xc0, 0xfa,
	0x97, 0xaf, 0xb6, 0x66, 0xb6, 0x50, 0x73, 0x53, 0xa3, 0xe6, 0x4f, 0x0d, 0x1d, 0x6c, 0xb0, 0xc0,
	0x6f, 0x51, 0x8d, 0xd1, 0x02, 0xd6, 0xb0, 0xe7, 0x3c, 0xff, 0xff, 0x1e, 0x21, 0xde, 0xa7, 0x85,
	0x57, 0x96, 0xc1, 0xf7, 0x91, 0x7e, 0xa9, 0xf2, 0x19, 0x8c, 0x5c, 0xeb, 0xdd, 0xfa, 0xf5, 0xcd,
	0xd6, 0xed, 0x8e, 0xf3, 0xc4, 0x83, 0xa8, 0x19, 0xa2, 0xc6, 0x92, 0xc6, 0x4d, 0xb4, 0x3f, 0x71,
	0xdd, 0x37, 0x41, 0xbf, 0xfb, 0x3e, 0xf0, 0x87, 0xe7, 0xef, 0xdc, 0xe3, 0xc1, 0xc9, 0xc0, 0xed,
	0x1b, 0x5b, 0xb8, 0x81, 0x6a, 0x67, 0xa3, 0xa1, 0xa1, 0x95, 0x62, 0xec, 0xbb, 0xc6, 0x76, 0x29,
	0x26, 0x6e, 0xdf, 0xa8, 0x41, 0xe4, 0xd4, 0x37, 0xf4, 0x52, 0x9c, 0x78, 0x03, 0x63, 0xa7, 0x14,
	0xe7, 0xdd, 0xb1, 0x51, 0x07, 0xe1, 0x0f, 0x8d, 0x86, 0x39, 0x47, 0xfb, 0xd7, 0x7a, 0x1c, 0xa5,
	0x7c, 0x06, 0xe7, 0x80, 0xdb, 0x48, 0x17, 0xf2, 0x42, 0xc1, 0xa0, 0xbb, 0xbd, 0xdd, 0xcf, 0xdf,
	0xed, 0x9d, 0x97, 0xaf, 0x9c, 0xa7, 0xcf, 0x3c, 0x08, 0xe3, 0xd7, 0xe8, 0x2e, 0xe3, 0x31, 0x2d,
	0x38, 0x0b, 0x72, 0x99, 0x89, 0x78, 0xf9, 0xd1, 0x5a, 0xa4, 0xba, 0x3f, 0xb2, 0xba, 0x3f, 0x32,
	0x5e, 0xdd, 0x9f, 0x77, 0x67, 0x99, 0xe0, 0x97, 0x7c, 0x8f, 0xa1, 0xf6, 0xda, 0xee, 0x68, 0x2a,
	0xd6, 0xf6, 0xf7, 0xe1, 0x38, 0x12, 0xd9, 0x65, 0x1e, 0x92, 0xa9, 0x4a, 0xac, 0x8a, 0xec, 0x54,
	0x37, 0x1b, 0xa9, 0x4e, 0xc4, 0x25, 0x18, 0x58, 0x1b, 0xff, 0x9a, 0x17, 0x20, 0xc2, 0x3a, 0x60,
	0x47, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x07, 0xa7, 0xef, 0x5f, 0x03, 0x00, 0x00,
}
