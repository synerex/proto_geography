// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geography.proto

package proto_geography // import "github.com/synerex/proto_geography"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// for BarGraph type
type BarType int32

const (
	BarType_BT_BOX_FIXCOLOR BarType = 0
	BarType_BT_BOX_VARCOLOR BarType = 1
	BarType_BT_HEX_FIXCOLOR BarType = 2
	BarType_BT_HEX_VARCOLOR BarType = 3
)

var BarType_name = map[int32]string{
	0: "BT_BOX_FIXCOLOR",
	1: "BT_BOX_VARCOLOR",
	2: "BT_HEX_FIXCOLOR",
	3: "BT_HEX_VARCOLOR",
}
var BarType_value = map[string]int32{
	"BT_BOX_FIXCOLOR": 0,
	"BT_BOX_VARCOLOR": 1,
	"BT_HEX_FIXCOLOR": 2,
	"BT_HEX_VARCOLOR": 3,
}

func (x BarType) String() string {
	return proto.EnumName(BarType_name, int32(x))
}
func (BarType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{0}
}

// geographic data  (mainly for supply)
type Geo struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Options              string   `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Geo) Reset()         { *m = Geo{} }
func (m *Geo) String() string { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()    {}
func (*Geo) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{0}
}
func (m *Geo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Geo.Unmarshal(m, b)
}
func (m *Geo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Geo.Marshal(b, m, deterministic)
}
func (dst *Geo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Geo.Merge(dst, src)
}
func (m *Geo) XXX_Size() int {
	return xxx_messageInfo_Geo.Size(m)
}
func (m *Geo) XXX_DiscardUnknown() {
	xxx_messageInfo_Geo.DiscardUnknown(m)
}

var xxx_messageInfo_Geo proto.InternalMessageInfo

func (m *Geo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Geo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Geo) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Geo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Geo) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

type Line struct {
	From                 []float64 `protobuf:"fixed64,1,rep,packed,name=from,proto3" json:"from,omitempty"`
	To                   []float64 `protobuf:"fixed64,2,rep,packed,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Line) Reset()         { *m = Line{} }
func (m *Line) String() string { return proto.CompactTextString(m) }
func (*Line) ProtoMessage()    {}
func (*Line) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{1}
}
func (m *Line) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Line.Unmarshal(m, b)
}
func (m *Line) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Line.Marshal(b, m, deterministic)
}
func (dst *Line) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Line.Merge(dst, src)
}
func (m *Line) XXX_Size() int {
	return xxx_messageInfo_Line.Size(m)
}
func (m *Line) XXX_DiscardUnknown() {
	xxx_messageInfo_Line.DiscardUnknown(m)
}

var xxx_messageInfo_Line proto.InternalMessageInfo

func (m *Line) GetFrom() []float64 {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Line) GetTo() []float64 {
	if m != nil {
		return m.To
	}
	return nil
}

type Lines struct {
	Lines                []*Line  `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Color                []int32  `protobuf:"varint,3,rep,packed,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lines) Reset()         { *m = Lines{} }
func (m *Lines) String() string { return proto.CompactTextString(m) }
func (*Lines) ProtoMessage()    {}
func (*Lines) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{2}
}
func (m *Lines) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lines.Unmarshal(m, b)
}
func (m *Lines) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lines.Marshal(b, m, deterministic)
}
func (dst *Lines) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lines.Merge(dst, src)
}
func (m *Lines) XXX_Size() int {
	return xxx_messageInfo_Lines.Size(m)
}
func (m *Lines) XXX_DiscardUnknown() {
	xxx_messageInfo_Lines.DiscardUnknown(m)
}

var xxx_messageInfo_Lines proto.InternalMessageInfo

func (m *Lines) GetLines() []*Line {
	if m != nil {
		return m.Lines
	}
	return nil
}

func (m *Lines) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Lines) GetColor() []int32 {
	if m != nil {
		return m.Color
	}
	return nil
}

type Point struct {
	Lat                  float64  `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float64  `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{3}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Point) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Path struct {
	Points               []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Color                []int32  `protobuf:"varint,3,rep,packed,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{4}
}
func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (dst *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(dst, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Path) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Path) GetColor() []int32 {
	if m != nil {
		return m.Color
	}
	return nil
}

type Paths struct {
	Paths                []*Path  `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paths) Reset()         { *m = Paths{} }
func (m *Paths) String() string { return proto.CompactTextString(m) }
func (*Paths) ProtoMessage()    {}
func (*Paths) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{5}
}
func (m *Paths) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paths.Unmarshal(m, b)
}
func (m *Paths) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paths.Marshal(b, m, deterministic)
}
func (dst *Paths) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paths.Merge(dst, src)
}
func (m *Paths) XXX_Size() int {
	return xxx_messageInfo_Paths.Size(m)
}
func (m *Paths) XXX_DiscardUnknown() {
	xxx_messageInfo_Paths.DiscardUnknown(m)
}

var xxx_messageInfo_Paths proto.InternalMessageInfo

func (m *Paths) GetPaths() []*Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

type Scatter struct {
	Points               []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Radiuses             []int32  `protobuf:"varint,2,rep,packed,name=radiuses,proto3" json:"radiuses,omitempty"`
	Color                []int32  `protobuf:"varint,3,rep,packed,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scatter) Reset()         { *m = Scatter{} }
func (m *Scatter) String() string { return proto.CompactTextString(m) }
func (*Scatter) ProtoMessage()    {}
func (*Scatter) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{6}
}
func (m *Scatter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scatter.Unmarshal(m, b)
}
func (m *Scatter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scatter.Marshal(b, m, deterministic)
}
func (dst *Scatter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scatter.Merge(dst, src)
}
func (m *Scatter) XXX_Size() int {
	return xxx_messageInfo_Scatter.Size(m)
}
func (m *Scatter) XXX_DiscardUnknown() {
	xxx_messageInfo_Scatter.DiscardUnknown(m)
}

var xxx_messageInfo_Scatter proto.InternalMessageInfo

func (m *Scatter) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Scatter) GetRadiuses() []int32 {
	if m != nil {
		return m.Radiuses
	}
	return nil
}

func (m *Scatter) GetColor() []int32 {
	if m != nil {
		return m.Color
	}
	return nil
}

type ViewState struct {
	Lat                  float64  `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float64  `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	Zoom                 int32    `protobuf:"varint,3,opt,name=zoom,proto3" json:"zoom,omitempty"`
	Pitch                float64  `protobuf:"fixed64,4,opt,name=pitch,proto3" json:"pitch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewState) Reset()         { *m = ViewState{} }
func (m *ViewState) String() string { return proto.CompactTextString(m) }
func (*ViewState) ProtoMessage()    {}
func (*ViewState) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{7}
}
func (m *ViewState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewState.Unmarshal(m, b)
}
func (m *ViewState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewState.Marshal(b, m, deterministic)
}
func (dst *ViewState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewState.Merge(dst, src)
}
func (m *ViewState) XXX_Size() int {
	return xxx_messageInfo_ViewState.Size(m)
}
func (m *ViewState) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewState.DiscardUnknown(m)
}

var xxx_messageInfo_ViewState proto.InternalMessageInfo

func (m *ViewState) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *ViewState) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *ViewState) GetZoom() int32 {
	if m != nil {
		return m.Zoom
	}
	return 0
}

func (m *ViewState) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

type BarGraph struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ts                   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=ts,proto3" json:"ts,omitempty"`
	Type                 BarType              `protobuf:"varint,3,opt,name=type,proto3,enum=pagent.BarType" json:"type,omitempty"`
	Color                int32                `protobuf:"varint,4,opt,name=color,proto3" json:"color,omitempty"`
	Lat                  float64              `protobuf:"fixed64,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float64              `protobuf:"fixed64,6,opt,name=lon,proto3" json:"lon,omitempty"`
	Width                float64              `protobuf:"fixed64,7,opt,name=width,proto3" json:"width,omitempty"`
	Radius               float64              `protobuf:"fixed64,8,opt,name=radius,proto3" json:"radius,omitempty"`
	Value                float64              `protobuf:"fixed64,9,opt,name=value,proto3" json:"value,omitempty"`
	Min                  float64              `protobuf:"fixed64,10,opt,name=min,proto3" json:"min,omitempty"`
	Max                  float64              `protobuf:"fixed64,11,opt,name=max,proto3" json:"max,omitempty"`
	Text                 string               `protobuf:"bytes,12,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BarGraph) Reset()         { *m = BarGraph{} }
func (m *BarGraph) String() string { return proto.CompactTextString(m) }
func (*BarGraph) ProtoMessage()    {}
func (*BarGraph) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{8}
}
func (m *BarGraph) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarGraph.Unmarshal(m, b)
}
func (m *BarGraph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarGraph.Marshal(b, m, deterministic)
}
func (dst *BarGraph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarGraph.Merge(dst, src)
}
func (m *BarGraph) XXX_Size() int {
	return xxx_messageInfo_BarGraph.Size(m)
}
func (m *BarGraph) XXX_DiscardUnknown() {
	xxx_messageInfo_BarGraph.DiscardUnknown(m)
}

var xxx_messageInfo_BarGraph proto.InternalMessageInfo

func (m *BarGraph) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BarGraph) GetTs() *timestamp.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *BarGraph) GetType() BarType {
	if m != nil {
		return m.Type
	}
	return BarType_BT_BOX_FIXCOLOR
}

func (m *BarGraph) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *BarGraph) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *BarGraph) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func (m *BarGraph) GetWidth() float64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *BarGraph) GetRadius() float64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *BarGraph) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *BarGraph) GetMin() float64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *BarGraph) GetMax() float64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *BarGraph) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type BarGraphs struct {
	Bars                 []*BarGraphs `protobuf:"bytes,1,rep,name=bars,proto3" json:"bars,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BarGraphs) Reset()         { *m = BarGraphs{} }
func (m *BarGraphs) String() string { return proto.CompactTextString(m) }
func (*BarGraphs) ProtoMessage()    {}
func (*BarGraphs) Descriptor() ([]byte, []int) {
	return fileDescriptor_geography_706fd81f378a201d, []int{9}
}
func (m *BarGraphs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarGraphs.Unmarshal(m, b)
}
func (m *BarGraphs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarGraphs.Marshal(b, m, deterministic)
}
func (dst *BarGraphs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarGraphs.Merge(dst, src)
}
func (m *BarGraphs) XXX_Size() int {
	return xxx_messageInfo_BarGraphs.Size(m)
}
func (m *BarGraphs) XXX_DiscardUnknown() {
	xxx_messageInfo_BarGraphs.DiscardUnknown(m)
}

var xxx_messageInfo_BarGraphs proto.InternalMessageInfo

func (m *BarGraphs) GetBars() []*BarGraphs {
	if m != nil {
		return m.Bars
	}
	return nil
}

func init() {
	proto.RegisterType((*Geo)(nil), "pagent.Geo")
	proto.RegisterType((*Line)(nil), "pagent.Line")
	proto.RegisterType((*Lines)(nil), "pagent.Lines")
	proto.RegisterType((*Point)(nil), "pagent.Point")
	proto.RegisterType((*Path)(nil), "pagent.Path")
	proto.RegisterType((*Paths)(nil), "pagent.Paths")
	proto.RegisterType((*Scatter)(nil), "pagent.Scatter")
	proto.RegisterType((*ViewState)(nil), "pagent.ViewState")
	proto.RegisterType((*BarGraph)(nil), "pagent.BarGraph")
	proto.RegisterType((*BarGraphs)(nil), "pagent.BarGraphs")
	proto.RegisterEnum("pagent.BarType", BarType_name, BarType_value)
}

func init() { proto.RegisterFile("geography.proto", fileDescriptor_geography_706fd81f378a201d) }

var fileDescriptor_geography_706fd81f378a201d = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x71, 0x6b, 0xdb, 0x3e,
	0x10, 0xfd, 0xd9, 0x8e, 0x93, 0xe6, 0xda, 0x5f, 0xdb, 0x69, 0x63, 0x88, 0xfe, 0xb3, 0xe0, 0xad,
	0x10, 0x3a, 0x70, 0x21, 0xfb, 0x04, 0xcb, 0xd8, 0xba, 0x41, 0xa1, 0x45, 0x2d, 0x5d, 0x3b, 0x18,
	0x41, 0x49, 0x54, 0x5b, 0x60, 0x5b, 0xc6, 0xbe, 0xac, 0xc9, 0xbe, 0xcd, 0xbe, 0xe9, 0xd0, 0xc9,
	0xf6, 0xc2, 0xc6, 0xa0, 0xfb, 0xef, 0xee, 0xdd, 0xd3, 0xf9, 0x9e, 0xee, 0x59, 0x70, 0x90, 0x28,
	0x93, 0x54, 0xb2, 0x4c, 0x37, 0x71, 0x59, 0x19, 0x34, 0xac, 0x5f, 0xca, 0x44, 0x15, 0x78, 0xf4,
	0x22, 0x31, 0x26, 0xc9, 0xd4, 0x29, 0xa1, 0xf3, 0xd5, 0xfd, 0x29, 0xea, 0x5c, 0xd5, 0x28, 0xf3,
	0xd2, 0x11, 0xa3, 0x1c, 0x82, 0x33, 0x65, 0x18, 0x83, 0x1e, 0x6e, 0x4a, 0xc5, 0xbd, 0x91, 0x37,
	0x1e, 0x0a, 0x8a, 0xd9, 0x3e, 0xf8, 0x7a, 0xc9, 0xfd, 0x91, 0x37, 0x0e, 0x85, 0xaf, 0x97, 0xec,
	0x19, 0x84, 0x99, 0x9c, 0xab, 0x8c, 0x07, 0x44, 0x72, 0x89, 0x3d, 0xb9, 0x94, 0x28, 0x79, 0x6f,
	0xe4, 0x8d, 0xf7, 0x04, 0xc5, 0x8c, 0xc3, 0xc0, 0x94, 0xa8, 0x4d, 0x51, 0xf3, 0x90, 0xb8, 0x6d,
	0x1a, 0x9d, 0x40, 0xef, 0x5c, 0x17, 0xca, 0x9e, 0xba, 0xaf, 0x4c, 0xce, 0xbd, 0x51, 0x30, 0xf6,
	0x04, 0xc5, 0xf6, 0x7b, 0x68, 0xb8, 0x4f, 0x88, 0x8f, 0x26, 0xfa, 0x0c, 0xa1, 0xe5, 0xd6, 0x2c,
	0x82, 0x30, 0xb3, 0x01, 0xb1, 0x77, 0x27, 0x7b, 0xb1, 0x13, 0x17, 0xdb, 0xaa, 0x70, 0x25, 0x3b,
	0xdc, 0x83, 0x5e, 0x62, 0xda, 0xcc, 0xeb, 0x12, 0x8b, 0x2e, 0x4c, 0x66, 0x2a, 0x1e, 0x8c, 0x02,
	0x8b, 0x52, 0x12, 0xbd, 0x86, 0xf0, 0xd2, 0xe8, 0x02, 0xd9, 0x21, 0x04, 0x99, 0x44, 0x12, 0xed,
	0x09, 0x1b, 0x12, 0x62, 0x0a, 0x6a, 0x62, 0x11, 0x53, 0x44, 0x77, 0xd0, 0xbb, 0x94, 0x98, 0xb2,
	0x63, 0xe8, 0x97, 0xf6, 0x50, 0x3b, 0xc5, 0xff, 0xed, 0x14, 0xd4, 0x4a, 0x34, 0xc5, 0x7f, 0x9e,
	0x43, 0x62, 0x4a, 0x02, 0x4b, 0x1b, 0xfc, 0x2e, 0xd0, 0x56, 0x85, 0x2b, 0x45, 0x73, 0x18, 0x5c,
	0x2d, 0x24, 0xa2, 0xaa, 0x1e, 0x3b, 0xca, 0x11, 0xec, 0x54, 0x72, 0xa9, 0x57, 0xb5, 0xaa, 0xe9,
	0x56, 0x43, 0xd1, 0xe5, 0x7f, 0x19, 0xe8, 0x0e, 0x86, 0x37, 0x5a, 0x3d, 0x5c, 0xa1, 0x44, 0xf5,
	0x98, 0xcb, 0xb1, 0x6b, 0xfc, 0x6e, 0x4c, 0x4e, 0x8e, 0x08, 0x05, 0xc5, 0xb6, 0x75, 0xa9, 0x71,
	0x91, 0x92, 0x23, 0x3c, 0xe1, 0x92, 0xe8, 0x87, 0x0f, 0x3b, 0x53, 0x59, 0x9d, 0x59, 0x93, 0x36,
	0xce, 0xf2, 0x3a, 0x67, 0x9d, 0x80, 0x8f, 0x35, 0xf5, 0xdd, 0x9d, 0x1c, 0xc5, 0xce, 0xb2, 0x71,
	0x6b, 0xd9, 0xf8, 0xba, 0xb5, 0xac, 0xf0, 0xb1, 0x66, 0x2f, 0x1b, 0xa7, 0xda, 0x4f, 0xee, 0x4f,
	0x0e, 0x5a, 0xe9, 0x53, 0x59, 0x5d, 0x6f, 0x4a, 0xd5, 0x58, 0xb7, 0x93, 0xd7, 0x73, 0x5b, 0xa0,
	0xa4, 0x55, 0x14, 0xfe, 0xa1, 0xa8, 0xff, 0x4b, 0x51, 0xb7, 0xbf, 0x81, 0x9b, 0xde, 0xed, 0xef,
	0x39, 0xf4, 0xdd, 0xd5, 0xf1, 0x1d, 0x82, 0x9b, 0xcc, 0xb2, 0xbf, 0xc9, 0x6c, 0xa5, 0xf8, 0xd0,
	0xb1, 0x29, 0xb1, 0x5d, 0x73, 0x5d, 0x70, 0x70, 0x5d, 0x73, 0x5d, 0x10, 0x22, 0xd7, 0x7c, 0xb7,
	0x41, 0xe4, 0x9a, 0x7e, 0x38, 0xb5, 0x46, 0xbe, 0xd7, 0xfc, 0x70, 0x6a, 0x8d, 0xd1, 0x04, 0x86,
	0xed, 0x15, 0xd5, 0xec, 0x18, 0x7a, 0x73, 0x59, 0xb5, 0x2b, 0x7e, 0xb2, 0xa5, 0xd3, 0x11, 0x04,
	0x95, 0x4f, 0xbe, 0xc2, 0xa0, 0x91, 0xce, 0x9e, 0xc2, 0xc1, 0xf4, 0x7a, 0x36, 0xbd, 0xb8, 0x9d,
	0x7d, 0xf8, 0x74, 0xfb, 0xee, 0xe2, 0xfc, 0x42, 0x1c, 0xfe, 0xb7, 0x05, 0xde, 0xbc, 0x15, 0x0e,
	0xf4, 0x1a, 0xf0, 0xe3, 0xfb, 0x2d, 0xa6, 0xbf, 0x05, 0x76, 0xcc, 0x60, 0xfa, 0xea, 0x4b, 0x94,
	0x68, 0x4c, 0x57, 0xf3, 0x78, 0x61, 0xf2, 0xd3, 0x7a, 0x53, 0xa8, 0x4a, 0xad, 0xdd, 0x6b, 0x32,
	0xeb, 0xde, 0x9c, 0x79, 0x9f, 0x80, 0x37, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x8c, 0x4a,
	0x47, 0x87, 0x04, 0x00, 0x00,
}
