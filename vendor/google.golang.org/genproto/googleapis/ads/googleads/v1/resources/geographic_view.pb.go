// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/geographic_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A geographic view.
//
// Geographic View includes all metrics aggregated at the country level,
// one row per country. It reports metrics at either actual physical location of
// the user or an area of interest. If other segment fields are used, you may
// get more than one row per country.
type GeographicView struct {
	// The resource name of the geographic view.
	// Geographic view resource names have the form:
	//
	//
	// `customers/{customer_id}/geographicViews/{country_criterion_id}~{location_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// CriterionId for the geo target for a country.
	CountryGeoTargetConstant *wrappers.StringValue `protobuf:"bytes,2,opt,name=country_geo_target_constant,json=countryGeoTargetConstant,proto3" json:"country_geo_target_constant,omitempty"`
	// Type of the geo targeting of the campaign.
	LocationType         enums.GeoTargetingTypeEnum_GeoTargetingType `protobuf:"varint,3,opt,name=location_type,json=locationType,proto3,enum=google.ads.googleads.v1.enums.GeoTargetingTypeEnum_GeoTargetingType" json:"location_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *GeographicView) Reset()         { *m = GeographicView{} }
func (m *GeographicView) String() string { return proto.CompactTextString(m) }
func (*GeographicView) ProtoMessage()    {}
func (*GeographicView) Descriptor() ([]byte, []int) {
	return fileDescriptor_45a03e49a5e672c5, []int{0}
}

func (m *GeographicView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeographicView.Unmarshal(m, b)
}
func (m *GeographicView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeographicView.Marshal(b, m, deterministic)
}
func (m *GeographicView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeographicView.Merge(m, src)
}
func (m *GeographicView) XXX_Size() int {
	return xxx_messageInfo_GeographicView.Size(m)
}
func (m *GeographicView) XXX_DiscardUnknown() {
	xxx_messageInfo_GeographicView.DiscardUnknown(m)
}

var xxx_messageInfo_GeographicView proto.InternalMessageInfo

func (m *GeographicView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *GeographicView) GetCountryGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.CountryGeoTargetConstant
	}
	return nil
}

func (m *GeographicView) GetLocationType() enums.GeoTargetingTypeEnum_GeoTargetingType {
	if m != nil {
		return m.LocationType
	}
	return enums.GeoTargetingTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*GeographicView)(nil), "google.ads.googleads.v1.resources.GeographicView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/geographic_view.proto", fileDescriptor_45a03e49a5e672c5)
}

var fileDescriptor_45a03e49a5e672c5 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x8b, 0xd4, 0x30,
	0x18, 0xa5, 0x5d, 0x10, 0xac, 0xbb, 0x7b, 0xa8, 0x97, 0xb2, 0x2e, 0x32, 0xab, 0x2c, 0xcc, 0x29,
	0xa5, 0x2b, 0x28, 0xc4, 0x53, 0x57, 0xa5, 0xe0, 0x41, 0x96, 0x3a, 0xf4, 0xa0, 0x85, 0x92, 0x69,
	0x3f, 0x63, 0xa0, 0x4d, 0x42, 0x92, 0xce, 0x30, 0x37, 0x7f, 0x8b, 0x47, 0x7f, 0x8a, 0x3f, 0xc5,
	0xff, 0x20, 0x48, 0xd3, 0x26, 0x22, 0x32, 0x7a, 0x7b, 0xc9, 0xf7, 0xde, 0xe3, 0xbd, 0x2f, 0x89,
	0x5e, 0x50, 0x21, 0x68, 0x0f, 0x29, 0xe9, 0x74, 0x3a, 0xc3, 0x09, 0xed, 0xb2, 0x54, 0x81, 0x16,
	0xa3, 0x6a, 0x41, 0xa7, 0x14, 0x04, 0x55, 0x44, 0x7e, 0x66, 0x6d, 0xb3, 0x63, 0xb0, 0x47, 0x52,
	0x09, 0x23, 0xe2, 0xab, 0x99, 0x8d, 0x48, 0xa7, 0x91, 0x17, 0xa2, 0x5d, 0x86, 0xbc, 0xf0, 0xe2,
	0xf9, 0x31, 0x6f, 0xe0, 0xe3, 0x60, 0x7d, 0x1b, 0x43, 0x14, 0x05, 0xc3, 0x38, 0x6d, 0xcc, 0x41,
	0xc2, 0x6c, 0x7d, 0xf1, 0x78, 0xd1, 0xd9, 0xd3, 0x76, 0xfc, 0x94, 0xee, 0x15, 0x91, 0x12, 0x94,
	0x5e, 0xe6, 0x97, 0xce, 0x57, 0xb2, 0x94, 0x70, 0x2e, 0x0c, 0x31, 0x4c, 0xf0, 0x65, 0xfa, 0xe4,
	0x67, 0x10, 0x9d, 0x17, 0x3e, 0x72, 0xc5, 0x60, 0x1f, 0x3f, 0x8d, 0xce, 0x5c, 0xaa, 0x86, 0x93,
	0x01, 0x92, 0x60, 0x15, 0xac, 0xef, 0x97, 0xa7, 0xee, 0xf2, 0x1d, 0x19, 0x20, 0xfe, 0x18, 0x3d,
	0x6a, 0xc5, 0xc8, 0x8d, 0x3a, 0x34, 0xbf, 0x93, 0x35, 0xad, 0xe0, 0xda, 0x10, 0x6e, 0x92, 0x70,
	0x15, 0xac, 0x1f, 0xdc, 0x5c, 0x2e, 0x5d, 0x91, 0xcb, 0x86, 0xde, 0x1b, 0xc5, 0x38, 0xad, 0x48,
	0x3f, 0x42, 0x99, 0x2c, 0x06, 0x05, 0x88, 0x8d, 0x95, 0xbf, 0x5a, 0xd4, 0x31, 0x8b, 0xce, 0x7a,
	0xd1, 0xda, 0x9c, 0xb6, 0x69, 0x72, 0xb2, 0x0a, 0xd6, 0xe7, 0x37, 0xaf, 0xd1, 0xb1, 0x2d, 0xda,
	0x15, 0x21, 0x6f, 0xc4, 0x38, 0xdd, 0x1c, 0x24, 0xbc, 0xe1, 0xe3, 0xf0, 0xd7, 0x65, 0x79, 0xea,
	0xac, 0xa7, 0xd3, 0xed, 0x97, 0x30, 0xba, 0x6e, 0xc5, 0x80, 0xfe, 0xfb, 0x3e, 0xb7, 0x0f, 0xff,
	0x5c, 0xd3, 0xdd, 0x54, 0xe9, 0x2e, 0xf8, 0xf0, 0x76, 0x51, 0x52, 0xd1, 0x13, 0x4e, 0x91, 0x50,
	0x34, 0xa5, 0xc0, 0x6d, 0x61, 0xf7, 0x8c, 0x92, 0xe9, 0x7f, 0xfc, 0x98, 0x97, 0x1e, 0x7d, 0x0d,
	0x4f, 0x8a, 0x3c, 0xff, 0x16, 0x5e, 0x15, 0xb3, 0x65, 0xde, 0x69, 0x34, 0xc3, 0x09, 0x55, 0x19,
	0x2a, 0x1d, 0xf3, 0xbb, 0xe3, 0xd4, 0x79, 0xa7, 0x6b, 0xcf, 0xa9, 0xab, 0xac, 0xf6, 0x9c, 0x1f,
	0xe1, 0xf5, 0x3c, 0xc0, 0x38, 0xef, 0x34, 0xc6, 0x9e, 0x85, 0x71, 0x95, 0x61, 0xec, 0x79, 0xdb,
	0x7b, 0x36, 0xec, 0xb3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x75, 0x6e, 0x6b, 0xdd, 0x02,
	0x00, 0x00,
}