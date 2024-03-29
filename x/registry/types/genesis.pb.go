// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/registry/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the registry module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// pool_list ...
	PoolList []Pool `protobuf:"bytes,2,rep,name=pool_list,json=poolList,proto3" json:"pool_list"`
	// pool_count ...
	PoolCount uint64 `protobuf:"varint,3,opt,name=pool_count,json=poolCount,proto3" json:"pool_count,omitempty"`
	// funder_list ...
	FunderList []Funder `protobuf:"bytes,4,rep,name=funder_list,json=funderList,proto3" json:"funder_list"`
	// staker_list ...
	StakerList []Staker `protobuf:"bytes,5,rep,name=staker_list,json=stakerList,proto3" json:"staker_list"`
	// delegator_list ...
	DelegatorList []Delegator `protobuf:"bytes,6,rep,name=delegator_list,json=delegatorList,proto3" json:"delegator_list"`
	// delegation_pool_data_list ...
	DelegationPoolDataList []DelegationPoolData `protobuf:"bytes,7,rep,name=delegation_pool_data_list,json=delegationPoolDataList,proto3" json:"delegation_pool_data_list"`
	// delegation_entries_list ...
	DelegationEntriesList []DelegationEntries `protobuf:"bytes,8,rep,name=delegation_entries_list,json=delegationEntriesList,proto3" json:"delegation_entries_list"`
	// proposal_list ...
	ProposalList []Proposal `protobuf:"bytes,9,rep,name=proposal_list,json=proposalList,proto3" json:"proposal_list"`
	// unbonding_state ...
	UnbondingState UnbondingState `protobuf:"bytes,10,opt,name=unbonding_state,json=unbondingState,proto3" json:"unbonding_state"`
	// unbonding_entries ...
	UnbondingEntries []UnbondingEntries `protobuf:"bytes,11,rep,name=unbonding_entries,json=unbondingEntries,proto3" json:"unbonding_entries"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_99000362002b89f1, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPoolList() []Pool {
	if m != nil {
		return m.PoolList
	}
	return nil
}

func (m *GenesisState) GetPoolCount() uint64 {
	if m != nil {
		return m.PoolCount
	}
	return 0
}

func (m *GenesisState) GetFunderList() []Funder {
	if m != nil {
		return m.FunderList
	}
	return nil
}

func (m *GenesisState) GetStakerList() []Staker {
	if m != nil {
		return m.StakerList
	}
	return nil
}

func (m *GenesisState) GetDelegatorList() []Delegator {
	if m != nil {
		return m.DelegatorList
	}
	return nil
}

func (m *GenesisState) GetDelegationPoolDataList() []DelegationPoolData {
	if m != nil {
		return m.DelegationPoolDataList
	}
	return nil
}

func (m *GenesisState) GetDelegationEntriesList() []DelegationEntries {
	if m != nil {
		return m.DelegationEntriesList
	}
	return nil
}

func (m *GenesisState) GetProposalList() []Proposal {
	if m != nil {
		return m.ProposalList
	}
	return nil
}

func (m *GenesisState) GetUnbondingState() UnbondingState {
	if m != nil {
		return m.UnbondingState
	}
	return UnbondingState{}
}

func (m *GenesisState) GetUnbondingEntries() []UnbondingEntries {
	if m != nil {
		return m.UnbondingEntries
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kyve.registry.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("kyve/registry/v1beta1/genesis.proto", fileDescriptor_99000362002b89f1)
}

var fileDescriptor_99000362002b89f1 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0x1a, 0x42, 0xbb, 0x6e, 0x0b, 0x58, 0x14, 0x4c, 0x50, 0xdd, 0xa8, 0x80, 0x08,
	0x07, 0x6c, 0xb5, 0x1c, 0x91, 0x38, 0x40, 0x4a, 0x25, 0xfe, 0x09, 0xb5, 0x80, 0x44, 0x2f, 0xd1,
	0x26, 0xde, 0xba, 0x4b, 0xd2, 0x1d, 0xcb, 0x3b, 0x2e, 0xe4, 0x2d, 0x78, 0xac, 0x1e, 0x7b, 0x83,
	0x13, 0x42, 0xc9, 0x8b, 0x20, 0xcf, 0xae, 0x5b, 0x5a, 0xe2, 0xe6, 0x66, 0xad, 0xbf, 0xef, 0xf7,
	0x5b, 0x8d, 0x66, 0xd9, 0xfd, 0xc1, 0xe8, 0x48, 0x44, 0x99, 0x48, 0xa4, 0xc6, 0x6c, 0x14, 0x1d,
	0x6d, 0xf4, 0x04, 0xf2, 0x8d, 0x28, 0x11, 0x4a, 0x68, 0xa9, 0xc3, 0x34, 0x03, 0x04, 0x6f, 0xa5,
	0x08, 0x85, 0x65, 0x28, 0xb4, 0xa1, 0xe6, 0xad, 0x04, 0x12, 0xa0, 0x44, 0x54, 0x7c, 0x99, 0x70,
	0x73, 0x7d, 0x3a, 0x31, 0xe5, 0x19, 0x3f, 0xb4, 0xc0, 0xe6, 0x83, 0xe9, 0x99, 0x53, 0x03, 0xa5,
	0xd6, 0x7f, 0x36, 0xd8, 0xe2, 0xb6, 0xb9, 0xc8, 0x2e, 0x72, 0x14, 0xde, 0x33, 0xd6, 0x30, 0x18,
	0xdf, 0x69, 0x39, 0x6d, 0x77, 0x73, 0x35, 0x9c, 0x7a, 0xb1, 0xf0, 0x03, 0x85, 0x5e, 0xd4, 0x8f,
	0x7f, 0xaf, 0xd5, 0x76, 0x6c, 0xc5, 0x7b, 0xce, 0x16, 0x52, 0x80, 0x61, 0x77, 0x28, 0x35, 0xfa,
	0x57, 0x5a, 0x73, 0x6d, 0x77, 0xf3, 0x5e, 0x55, 0x1f, 0x60, 0x68, 0xdb, 0xf3, 0x45, 0xe7, 0xad,
	0xd4, 0xe8, 0xad, 0x32, 0x46, 0xfd, 0x3e, 0xe4, 0x0a, 0xfd, 0xb9, 0x96, 0xd3, 0xae, 0xef, 0x10,
	0xf1, 0x65, 0x71, 0xe0, 0x75, 0x98, 0xbb, 0x9f, 0xab, 0x58, 0x64, 0x46, 0x50, 0x27, 0x41, 0xd5,
	0x05, 0x5f, 0x51, 0xd2, 0x2a, 0x98, 0xe9, 0x91, 0xa4, 0xc3, 0x5c, 0x8d, 0x7c, 0x50, 0x52, 0xae,
	0x5e, 0x4a, 0xd9, 0xa5, 0x64, 0x49, 0x31, 0x3d, 0xa2, 0xbc, 0x63, 0xcb, 0xb1, 0x18, 0x8a, 0x84,
	0x23, 0x58, 0x50, 0x83, 0x40, 0xad, 0x0a, 0x50, 0xa7, 0x0c, 0x5b, 0xd6, 0xd2, 0x69, 0x9b, 0x70,
	0x5f, 0xd9, 0x5d, 0x7b, 0x20, 0x41, 0x75, 0x69, 0x08, 0x31, 0x47, 0x6e, 0xc8, 0xd7, 0x88, 0xfc,
	0xf8, 0x72, 0xb2, 0x04, 0x55, 0xcc, 0xb4, 0xc3, 0x91, 0x5b, 0xc5, 0xed, 0xf8, 0xbf, 0x3f, 0xe4,
	0xda, 0x67, 0x77, 0xfe, 0x71, 0x09, 0x85, 0x99, 0x14, 0xda, 0x98, 0xe6, 0xc9, 0xd4, 0x9e, 0x69,
	0xda, 0x32, 0x25, 0x2b, 0x5a, 0x89, 0x2f, 0xfe, 0x20, 0xcf, 0x6b, 0xb6, 0x94, 0x66, 0x90, 0x82,
	0xe6, 0x76, 0x23, 0x16, 0x88, 0xbe, 0x56, 0xb5, 0x11, 0x36, 0x6b, 0xa1, 0x8b, 0x65, 0x97, 0x58,
	0x1f, 0xd9, 0xf5, 0x5c, 0xf5, 0x40, 0xc5, 0x52, 0x25, 0x5d, 0x5d, 0x6c, 0xaa, 0xcf, 0x68, 0x3f,
	0x1f, 0x56, 0xd0, 0x3e, 0x95, 0x69, 0x5a, 0x6b, 0xcb, 0x5c, 0xce, 0xcf, 0x9d, 0x7a, 0x7b, 0xec,
	0xe6, 0x19, 0xd5, 0x0e, 0xc2, 0x77, 0xe9, 0x96, 0x8f, 0x66, 0x71, 0xcf, 0x8f, 0xe0, 0x46, 0x7e,
	0xf1, 0x7c, 0xfb, 0x78, 0x1c, 0x38, 0x27, 0xe3, 0xc0, 0xf9, 0x33, 0x0e, 0x9c, 0x1f, 0x93, 0xa0,
	0x76, 0x32, 0x09, 0x6a, 0xbf, 0x26, 0x41, 0x6d, 0xef, 0x49, 0x22, 0xf1, 0x20, 0xef, 0x85, 0x7d,
	0x38, 0x8c, 0xde, 0x7c, 0xf9, 0xbc, 0xf5, 0x5e, 0xe0, 0x37, 0xc8, 0x06, 0x51, 0xff, 0x80, 0x4b,
	0x15, 0x7d, 0x3f, 0x7b, 0xb3, 0x38, 0x4a, 0x85, 0xee, 0x35, 0xe8, 0xa5, 0x3e, 0xfd, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x26, 0x07, 0x62, 0x25, 0x47, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UnbondingEntries) > 0 {
		for iNdEx := len(m.UnbondingEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbondingEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	{
		size, err := m.UnbondingState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if len(m.ProposalList) > 0 {
		for iNdEx := len(m.ProposalList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposalList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.DelegationEntriesList) > 0 {
		for iNdEx := len(m.DelegationEntriesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationEntriesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.DelegationPoolDataList) > 0 {
		for iNdEx := len(m.DelegationPoolDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationPoolDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.DelegatorList) > 0 {
		for iNdEx := len(m.DelegatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.StakerList) > 0 {
		for iNdEx := len(m.StakerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.FunderList) > 0 {
		for iNdEx := len(m.FunderList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FunderList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.PoolCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PoolCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PoolList) > 0 {
		for iNdEx := len(m.PoolList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.PoolList) > 0 {
		for _, e := range m.PoolList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PoolCount != 0 {
		n += 1 + sovGenesis(uint64(m.PoolCount))
	}
	if len(m.FunderList) > 0 {
		for _, e := range m.FunderList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StakerList) > 0 {
		for _, e := range m.StakerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegatorList) > 0 {
		for _, e := range m.DelegatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegationPoolDataList) > 0 {
		for _, e := range m.DelegationPoolDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegationEntriesList) > 0 {
		for _, e := range m.DelegationEntriesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProposalList) > 0 {
		for _, e := range m.ProposalList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.UnbondingState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.UnbondingEntries) > 0 {
		for _, e := range m.UnbondingEntries {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolList = append(m.PoolList, Pool{})
			if err := m.PoolList[len(m.PoolList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCount", wireType)
			}
			m.PoolCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunderList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunderList = append(m.FunderList, Funder{})
			if err := m.FunderList[len(m.FunderList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerList = append(m.StakerList, Staker{})
			if err := m.StakerList[len(m.StakerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorList = append(m.DelegatorList, Delegator{})
			if err := m.DelegatorList[len(m.DelegatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationPoolDataList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationPoolDataList = append(m.DelegationPoolDataList, DelegationPoolData{})
			if err := m.DelegationPoolDataList[len(m.DelegationPoolDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationEntriesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationEntriesList = append(m.DelegationEntriesList, DelegationEntries{})
			if err := m.DelegationEntriesList[len(m.DelegationEntriesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalList = append(m.ProposalList, Proposal{})
			if err := m.ProposalList[len(m.ProposalList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnbondingState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbondingEntries = append(m.UnbondingEntries, UnbondingEntries{})
			if err := m.UnbondingEntries[len(m.UnbondingEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
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
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
