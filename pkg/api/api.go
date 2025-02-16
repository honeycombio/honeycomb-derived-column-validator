package api

import (
	_ "github.com/gogo/protobuf/gogoproto"
)

type DeriveOp int32

const (
	DeriveOp_D_NONE        DeriveOp = 0
	DeriveOp_D_IF          DeriveOp = 1
	DeriveOp_D_COALESCE    DeriveOp = 2
	DeriveOp_D_NOT         DeriveOp = 3
	DeriveOp_D_LT          DeriveOp = 4
	DeriveOp_D_LTE         DeriveOp = 5
	DeriveOp_D_GT          DeriveOp = 6
	DeriveOp_D_GTE         DeriveOp = 7
	DeriveOp_D_IN          DeriveOp = 8
	DeriveOp_D_AND         DeriveOp = 9
	DeriveOp_D_OR          DeriveOp = 10
	DeriveOp_D_MIN         DeriveOp = 11
	DeriveOp_D_MAX         DeriveOp = 12
	DeriveOp_D_CONCAT      DeriveOp = 13
	DeriveOp_D_HAS_VALUE   DeriveOp = 14
	DeriveOp_D_PREFIX      DeriveOp = 15
	DeriveOp_D_SUFFIX      DeriveOp = 56
	DeriveOp_D_BUCKET      DeriveOp = 41
	DeriveOp_D_FORMAT_TIME DeriveOp = 44
	DeriveOp_D_LENGTH      DeriveOp = 46
	DeriveOp_D_SWITCH      DeriveOp = 48
	// All arithmetic ops coerce to float and return floats
	DeriveOp_D_SUM   DeriveOp = 16
	DeriveOp_D_SUB   DeriveOp = 17
	DeriveOp_D_MUL   DeriveOp = 18
	DeriveOp_D_DIV   DeriveOp = 19
	DeriveOp_D_LOG10 DeriveOp = 40
	DeriveOp_D_MOD   DeriveOp = 45
	// Explicit cast operations
	DeriveOp_D_INT    DeriveOp = 25
	DeriveOp_D_FLOAT  DeriveOp = 27
	DeriveOp_D_BOOL   DeriveOp = 38
	DeriveOp_D_STRING DeriveOp = 39
	// string-only operations
	DeriveOp_D_CONTAINS       DeriveOp = 20
	DeriveOp_D_REG_MATCH      DeriveOp = 21
	DeriveOp_D_REG_VALUE      DeriveOp = 22
	DeriveOp_D_UNIX_TIMESTAMP DeriveOp = 23
	DeriveOp_D_REG_COUNT      DeriveOp = 24
	DeriveOp_D_TO_LOWER       DeriveOp = 55
	DeriveOp_D_SEARCH         DeriveOp = 57
	// Event metadata
	DeriveOp_D_EVENT_TIMESTAMP  DeriveOp = 42
	DeriveOp_D_INGEST_TIMESTAMP DeriveOp = 43
)

type isDerivedParam_Value interface {
	isDerivedParam_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type isData_Value_Value interface {
	isData_Value_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Data_Value struct {
	// Types that are valid to be assigned to Value:
	//
	//	*Data_Value_Str
	//	*Data_Value_Int
	//	*Data_Value_Flt
	//	*Data_Value_Bool
	//	*Data_Value_Histogram
	//	*Data_Value_Distinct
	//	*Data_Value_Tdigest
	//	*Data_Value_Quantile
	//	*Data_Value_Time
	//	*Data_Value_Concurrency
	//	*Data_Value_Rate
	Value isData_Value_Value `protobuf_oneof:"value"`
}

type DerivedParam struct {
	// Types that are valid to be assigned to Value:
	//
	//	*DerivedParam_Static
	//	*DerivedParam_Column
	//	*DerivedParam_Derived
	Value isDerivedParam_Value `protobuf_oneof:"value"`
}

type DerivedParam_Static struct {
	Static *Data_Value `protobuf:"bytes,1,opt,name=static,proto3,oneof" json:"static,omitempty"`
}
type DerivedParam_Column struct {
	Column int64 `protobuf:"varint,2,opt,name=column,proto3" json:"column"`
}
type DerivedParam_Derived struct {
	Derived *DerivedValue `protobuf:"bytes,3,opt,name=derived,proto3,oneof" json:"derived,omitempty"`
}

func (*DerivedParam_Derived) isDerivedParam_Value() {}
func (*DerivedParam_Static) isDerivedParam_Value()  {}

// For the type derived from proto: DerivedParam_Derived
func (m *DerivedParam_Derived) MarshalTo(data []byte) (n int, err error) {
	// Since marshalling isn't needed for the current use-case, return a dummy value.
	return 0, nil
}

func (m *DerivedParam_Static) MarshalTo(data []byte) (n int, err error) {
	// Since marshalling isn't needed for the current use-case, return a dummy value.
	return 0, nil
}

func (m *DerivedParam_Derived) Size() int {
	return 0
}

func (m *DerivedParam_Static) Size() int {
	return 0
}

// func (m *DerivedParam) Reset()         { *m = DerivedParam{} }
// func (m *DerivedParam) String() string { return proto.CompactTextString(m) }
// func (*DerivedParam) ProtoMessage()    {}
// func (*DerivedParam) Descriptor() ([]byte, []int) {
// 	return fileDescriptor_4dc296cbfe5ffcd5, []int{10}
// }
// func (m *DerivedParam) XXX_Unmarshal(b []byte) error {
// 	return m.Unmarshal(b)
// }
// func (m *DerivedParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
// 	if deterministic {
// 		return xxx_messageInfo_DerivedParam.Marshal(b, m, deterministic)
// 	} else {
// 		b = b[:cap(b)]
// 		n, err := m.MarshalToSizedBuffer(b)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return b[:n], nil
// 	}
// }
// func (m *DerivedParam) XXX_Merge(src proto.Message) {
// 	xxx_messageInfo_DerivedParam.Merge(m, src)
// }
// func (m *DerivedParam) XXX_Size() int {
// 	return m.Size()
// }
// func (m *DerivedParam) XXX_DiscardUnknown() {
// 	xxx_messageInfo_DerivedParam.DiscardUnknown(m)
// }

type DerivedValue struct {
	Op     DeriveOp        `protobuf:"varint,1,opt,name=op,proto3,enum=api.DeriveOp" json:"op,omitempty"`
	Params []*DerivedParam `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty"`
}

// func (m *DerivedValue) Reset()         { *m = DerivedValue{} }
// func (m *DerivedValue) String() string { return proto.CompactTextString(m) }
// func (*DerivedValue) ProtoMessage()    {}
// func (*DerivedValue) Descriptor() ([]byte, []int) {
// 	return fileDescriptor_4dc296cbfe5ffcd5, []int{11}
// }
// func (m *DerivedValue) XXX_Unmarshal(b []byte) error {
// 	return m.Unmarshal(b)
// }
// func (m *DerivedValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
// 	if deterministic {
// 		return xxx_messageInfo_DerivedValue.Marshal(b, m, deterministic)
// 	} else {
// 		b = b[:cap(b)]
// 		n, err := m.MarshalToSizedBuffer(b)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return b[:n], nil
// 	}
// }
// func (m *DerivedValue) XXX_Merge(src proto.Message) {
// 	xxx_messageInfo_DerivedValue.Merge(m, src)
// }
// func (m *DerivedValue) XXX_Size() int {
// 	return m.Size()
// }
// func (m *DerivedValue) XXX_DiscardUnknown() {
// 	xxx_messageInfo_DerivedValue.DiscardUnknown(m)
// }

// var xxx_messageInfo_DerivedValue proto.InternalMessageInfo

type FilterOp int32

const (
	FilterOp_UNKNOWN                FilterOp = 0
	FilterOp_EQUAL                  FilterOp = 1
	FilterOp_NOT_EQUAL              FilterOp = 2
	FilterOp_GT                     FilterOp = 3
	FilterOp_GTE                    FilterOp = 4
	FilterOp_LT                     FilterOp = 5
	FilterOp_LTE                    FilterOp = 6
	FilterOp_PREFIX                 FilterOp = 7
	FilterOp_NOT_PREFIX             FilterOp = 8
	FilterOp_HAS_VALUE              FilterOp = 9
	FilterOp_NOT_HAS_VALUE          FilterOp = 10
	FilterOp_CONTAINS               FilterOp = 11
	FilterOp_NOT_CONTAINS           FilterOp = 12
	FilterOp_IN                     FilterOp = 15
	FilterOp_NOT_IN                 FilterOp = 16
	FilterOp_SEARCH                 FilterOp = 17
	FilterOp_SUFFIX                 FilterOp = 18
	FilterOp_NOT_SUFFIX             FilterOp = 19
)

func FilterOpIsUnary(op FilterOp) bool {
	return op == FilterOp_HAS_VALUE || op == FilterOp_NOT_HAS_VALUE
}
