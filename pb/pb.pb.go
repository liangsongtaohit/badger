// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ManifestChange_Operation int32

const (
	ManifestChange_CREATE ManifestChange_Operation = 0
	ManifestChange_DELETE ManifestChange_Operation = 1
)

var ManifestChange_Operation_name = map[int32]string{
	0: "CREATE",
	1: "DELETE",
}
var ManifestChange_Operation_value = map[string]int32{
	"CREATE": 0,
	"DELETE": 1,
}

func (x ManifestChange_Operation) String() string {
	return proto.EnumName(ManifestChange_Operation_name, int32(x))
}
func (ManifestChange_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_pb_7d6edc481ce0a5cd, []int{3, 0}
}

type KV struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	UserMeta             []byte   `protobuf:"bytes,3,opt,name=user_meta,json=userMeta,proto3" json:"user_meta,omitempty"`
	Version              uint64   `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	ExpiresAt            uint64   `protobuf:"varint,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	Meta                 []byte   `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_7d6edc481ce0a5cd, []int{0}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KV.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(dst, src)
}
func (m *KV) XXX_Size() int {
	return m.Size()
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KV) GetUserMeta() []byte {
	if m != nil {
		return m.UserMeta
	}
	return nil
}

func (m *KV) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *KV) GetExpiresAt() uint64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *KV) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

type KVList struct {
	Kv                   []*KV    `protobuf:"bytes,1,rep,name=kv" json:"kv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVList) Reset()         { *m = KVList{} }
func (m *KVList) String() string { return proto.CompactTextString(m) }
func (*KVList) ProtoMessage()    {}
func (*KVList) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_7d6edc481ce0a5cd, []int{1}
}
func (m *KVList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KVList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KVList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *KVList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVList.Merge(dst, src)
}
func (m *KVList) XXX_Size() int {
	return m.Size()
}
func (m *KVList) XXX_DiscardUnknown() {
	xxx_messageInfo_KVList.DiscardUnknown(m)
}

var xxx_messageInfo_KVList proto.InternalMessageInfo

func (m *KVList) GetKv() []*KV {
	if m != nil {
		return m.Kv
	}
	return nil
}

type ManifestChangeSet struct {
	// A set of changes that are applied atomically.
	Changes              []*ManifestChange `protobuf:"bytes,1,rep,name=changes" json:"changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ManifestChangeSet) Reset()         { *m = ManifestChangeSet{} }
func (m *ManifestChangeSet) String() string { return proto.CompactTextString(m) }
func (*ManifestChangeSet) ProtoMessage()    {}
func (*ManifestChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_7d6edc481ce0a5cd, []int{2}
}
func (m *ManifestChangeSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestChangeSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ManifestChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestChangeSet.Merge(dst, src)
}
func (m *ManifestChangeSet) XXX_Size() int {
	return m.Size()
}
func (m *ManifestChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestChangeSet proto.InternalMessageInfo

func (m *ManifestChangeSet) GetChanges() []*ManifestChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

type ManifestChange struct {
	Id                   uint64                   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Op                   ManifestChange_Operation `protobuf:"varint,2,opt,name=Op,proto3,enum=pb.ManifestChange_Operation" json:"Op,omitempty"`
	Level                uint32                   `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ManifestChange) Reset()         { *m = ManifestChange{} }
func (m *ManifestChange) String() string { return proto.CompactTextString(m) }
func (*ManifestChange) ProtoMessage()    {}
func (*ManifestChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_7d6edc481ce0a5cd, []int{3}
}
func (m *ManifestChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManifestChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManifestChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ManifestChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestChange.Merge(dst, src)
}
func (m *ManifestChange) XXX_Size() int {
	return m.Size()
}
func (m *ManifestChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestChange.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestChange proto.InternalMessageInfo

func (m *ManifestChange) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ManifestChange) GetOp() ManifestChange_Operation {
	if m != nil {
		return m.Op
	}
	return ManifestChange_CREATE
}

func (m *ManifestChange) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*KV)(nil), "pb.KV")
	proto.RegisterType((*KVList)(nil), "pb.KVList")
	proto.RegisterType((*ManifestChangeSet)(nil), "pb.ManifestChangeSet")
	proto.RegisterType((*ManifestChange)(nil), "pb.ManifestChange")
	proto.RegisterEnum("pb.ManifestChange_Operation", ManifestChange_Operation_name, ManifestChange_Operation_value)
}
func (m *KV) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KV) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if len(m.UserMeta) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.UserMeta)))
		i += copy(dAtA[i:], m.UserMeta)
	}
	if m.Version != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPb(dAtA, i, uint64(m.Version))
	}
	if m.ExpiresAt != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintPb(dAtA, i, uint64(m.ExpiresAt))
	}
	if len(m.Meta) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPb(dAtA, i, uint64(len(m.Meta)))
		i += copy(dAtA[i:], m.Meta)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *KVList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KVList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Kv) > 0 {
		for _, msg := range m.Kv {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ManifestChangeSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChangeSet) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, msg := range m.Changes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintPb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ManifestChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManifestChange) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPb(dAtA, i, uint64(m.Id))
	}
	if m.Op != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintPb(dAtA, i, uint64(m.Op))
	}
	if m.Level != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintPb(dAtA, i, uint64(m.Level))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *KV) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	l = len(m.UserMeta)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovPb(uint64(m.Version))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovPb(uint64(m.ExpiresAt))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovPb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *KVList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Kv) > 0 {
		for _, e := range m.Kv {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManifestChangeSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovPb(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ManifestChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPb(uint64(m.Id))
	}
	if m.Op != 0 {
		n += 1 + sovPb(uint64(m.Op))
	}
	if m.Level != 0 {
		n += 1 + sovPb(uint64(m.Level))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPb(x uint64) (n int) {
	return sovPb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KV) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KV: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KV: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserMeta", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserMeta = append(m.UserMeta[:0], dAtA[iNdEx:postIndex]...)
			if m.UserMeta == nil {
				m.UserMeta = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = append(m.Meta[:0], dAtA[iNdEx:postIndex]...)
			if m.Meta == nil {
				m.Meta = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *KVList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: KVList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KVList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kv = append(m.Kv, &KV{})
			if err := m.Kv[len(m.Kv)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestChangeSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestChangeSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChangeSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, &ManifestChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ManifestChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManifestChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManifestChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= (ManifestChange_Operation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPb(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb.proto", fileDescriptor_pb_7d6edc481ce0a5cd) }

var fileDescriptor_pb_7d6edc481ce0a5cd = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4f, 0x4e, 0xfa, 0x40,
	0x14, 0x66, 0x86, 0x52, 0xe0, 0xfd, 0x7e, 0x92, 0xfa, 0x62, 0x4c, 0x13, 0xb5, 0x69, 0xea, 0x86,
	0x05, 0xe9, 0x02, 0x4f, 0x80, 0xd8, 0x05, 0x01, 0x42, 0x32, 0x12, 0xb6, 0xa4, 0x95, 0xa7, 0x36,
	0x60, 0x3b, 0x69, 0x87, 0x46, 0x8f, 0xe0, 0x05, 0x8c, 0x47, 0x72, 0xe9, 0x11, 0x0c, 0x5e, 0xc4,
	0x74, 0x00, 0x13, 0xe2, 0xee, 0xfb, 0xf7, 0xbe, 0xc5, 0xf7, 0xa0, 0x21, 0x23, 0x5f, 0x66, 0xa9,
	0x4a, 0x91, 0xcb, 0xc8, 0x7b, 0x63, 0xc0, 0x87, 0x33, 0xb4, 0xa0, 0xba, 0xa4, 0x17, 0x9b, 0xb9,
	0xac, 0xfd, 0x5f, 0x94, 0x10, 0x4f, 0xa0, 0x56, 0x84, 0xab, 0x35, 0xd9, 0x5c, 0x6b, 0x5b, 0x82,
	0x67, 0xd0, 0x5c, 0xe7, 0x94, 0xcd, 0x9f, 0x48, 0x85, 0x76, 0x55, 0x3b, 0x8d, 0x52, 0x18, 0x93,
	0x0a, 0xd1, 0x86, 0x7a, 0x41, 0x59, 0x1e, 0xa7, 0x89, 0x6d, 0xb8, 0xac, 0x6d, 0x88, 0x3d, 0xc5,
	0x0b, 0x00, 0x7a, 0x96, 0x71, 0x46, 0xf9, 0x3c, 0x54, 0x76, 0x4d, 0x9b, 0xcd, 0x9d, 0xd2, 0x53,
	0x88, 0x60, 0xe8, 0x42, 0x53, 0x17, 0x6a, 0xec, 0xb9, 0x60, 0x0e, 0x67, 0xa3, 0x38, 0x57, 0x78,
	0x0a, 0x7c, 0x59, 0xd8, 0xcc, 0xad, 0xb6, 0xff, 0x75, 0x4d, 0x5f, 0x46, 0xfe, 0x70, 0x26, 0xf8,
	0xb2, 0xf0, 0x7a, 0x70, 0x3c, 0x0e, 0x93, 0xf8, 0x9e, 0x72, 0xd5, 0x7f, 0x0c, 0x93, 0x07, 0xba,
	0x25, 0x85, 0x1d, 0xa8, 0xdf, 0x69, 0x92, 0xef, 0x2e, 0xb0, 0xbc, 0x38, 0xcc, 0x89, 0x7d, 0xc4,
	0x7b, 0x65, 0xd0, 0x3a, 0xf4, 0xb0, 0x05, 0x7c, 0xb0, 0xd0, 0x43, 0x18, 0x82, 0x0f, 0x16, 0xd8,
	0x01, 0x3e, 0x91, 0x7a, 0x84, 0x56, 0xf7, 0xfc, 0x6f, 0x97, 0x3f, 0x91, 0x94, 0x85, 0x2a, 0x4e,
	0x13, 0xc1, 0x27, 0xb2, 0x5c, 0x6d, 0x44, 0x05, 0xad, 0xf4, 0x36, 0x47, 0x62, 0x4b, 0xbc, 0x4b,
	0x68, 0xfe, 0xc6, 0x10, 0xc0, 0xec, 0x8b, 0xa0, 0x37, 0x0d, 0xac, 0x4a, 0x89, 0x6f, 0x82, 0x51,
	0x30, 0x0d, 0x2c, 0x76, 0x6d, 0x7d, 0x6c, 0x1c, 0xf6, 0xb9, 0x71, 0xd8, 0xd7, 0xc6, 0x61, 0xef,
	0xdf, 0x4e, 0x25, 0x32, 0xf5, 0x9b, 0xae, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x0e, 0xe6,
	0x7c, 0xb2, 0x01, 0x00, 0x00,
}
