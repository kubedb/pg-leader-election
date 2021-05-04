/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kmodules.xyz/monitoring-agent-api/api/v1/generated.proto

package v1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	v1 "k8s.io/api/core/v1"
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

func (m *AgentSpec) Reset()      { *m = AgentSpec{} }
func (*AgentSpec) ProtoMessage() {}
func (*AgentSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb8a0907217b332, []int{0}
}
func (m *AgentSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AgentSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AgentSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentSpec.Merge(m, src)
}
func (m *AgentSpec) XXX_Size() int {
	return m.Size()
}
func (m *AgentSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentSpec.DiscardUnknown(m)
}

var xxx_messageInfo_AgentSpec proto.InternalMessageInfo

func (m *PrometheusExporterSpec) Reset()      { *m = PrometheusExporterSpec{} }
func (*PrometheusExporterSpec) ProtoMessage() {}
func (*PrometheusExporterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb8a0907217b332, []int{1}
}
func (m *PrometheusExporterSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrometheusExporterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PrometheusExporterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrometheusExporterSpec.Merge(m, src)
}
func (m *PrometheusExporterSpec) XXX_Size() int {
	return m.Size()
}
func (m *PrometheusExporterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PrometheusExporterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PrometheusExporterSpec proto.InternalMessageInfo

func (m *PrometheusSpec) Reset()      { *m = PrometheusSpec{} }
func (*PrometheusSpec) ProtoMessage() {}
func (*PrometheusSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb8a0907217b332, []int{2}
}
func (m *PrometheusSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PrometheusSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *PrometheusSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrometheusSpec.Merge(m, src)
}
func (m *PrometheusSpec) XXX_Size() int {
	return m.Size()
}
func (m *PrometheusSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PrometheusSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PrometheusSpec proto.InternalMessageInfo

func (m *ServiceMonitorSpec) Reset()      { *m = ServiceMonitorSpec{} }
func (*ServiceMonitorSpec) ProtoMessage() {}
func (*ServiceMonitorSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdb8a0907217b332, []int{3}
}
func (m *ServiceMonitorSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceMonitorSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ServiceMonitorSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMonitorSpec.Merge(m, src)
}
func (m *ServiceMonitorSpec) XXX_Size() int {
	return m.Size()
}
func (m *ServiceMonitorSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMonitorSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMonitorSpec proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AgentSpec)(nil), "kmodules.xyz.monitoring_agent_api.api.v1.AgentSpec")
	proto.RegisterType((*PrometheusExporterSpec)(nil), "kmodules.xyz.monitoring_agent_api.api.v1.PrometheusExporterSpec")
	proto.RegisterType((*PrometheusSpec)(nil), "kmodules.xyz.monitoring_agent_api.api.v1.PrometheusSpec")
	proto.RegisterType((*ServiceMonitorSpec)(nil), "kmodules.xyz.monitoring_agent_api.api.v1.ServiceMonitorSpec")
	proto.RegisterMapType((map[string]string)(nil), "kmodules.xyz.monitoring_agent_api.api.v1.ServiceMonitorSpec.LabelsEntry")
}

func init() {
	proto.RegisterFile("kmodules.xyz/monitoring-agent-api/api/v1/generated.proto", fileDescriptor_fdb8a0907217b332)
}

var fileDescriptor_fdb8a0907217b332 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x3d, 0x6f, 0xd3, 0x4e,
	0x18, 0xcf, 0xc5, 0x49, 0x15, 0x5f, 0xfe, 0x4a, 0xfb, 0x3f, 0x10, 0xb2, 0x22, 0x64, 0x47, 0x61,
	0xf1, 0x40, 0x1d, 0xda, 0x29, 0x45, 0x0c, 0xd4, 0x28, 0x12, 0x48, 0x20, 0x55, 0x57, 0x84, 0x04,
	0x4b, 0xe5, 0xb8, 0x0f, 0xae, 0xd5, 0xe4, 0xce, 0x9c, 0xcf, 0x56, 0xc2, 0xc4, 0x47, 0xe0, 0x5b,
	0xf0, 0x55, 0x3a, 0x56, 0x4c, 0x9d, 0x2c, 0x6a, 0x3e, 0x00, 0x7b, 0x27, 0xe4, 0xb3, 0xdb, 0x24,
	0x25, 0x43, 0x05, 0x83, 0xa5, 0x7b, 0x5e, 0x7e, 0x2f, 0xcf, 0x73, 0x67, 0x3c, 0x3c, 0x9d, 0xf2,
	0xe3, 0x64, 0x02, 0xb1, 0x33, 0x9b, 0x7f, 0x1e, 0x4c, 0x39, 0x0b, 0x25, 0x17, 0x21, 0x0b, 0xb6,
	0xbd, 0x00, 0x98, 0xdc, 0xf6, 0xa2, 0x70, 0x50, 0x7c, 0xe9, 0xce, 0x20, 0x00, 0x06, 0xc2, 0x93,
	0x70, 0xec, 0x44, 0x82, 0x4b, 0x4e, 0xec, 0x65, 0xa4, 0xb3, 0x40, 0x1e, 0x29, 0xe4, 0x91, 0x17,
	0x85, 0x4e, 0xf1, 0xa5, 0x3b, 0xdd, 0xed, 0x20, 0x94, 0x27, 0xc9, 0xd8, 0xf1, 0xf9, 0x74, 0x10,
	0xf0, 0x80, 0x0f, 0x14, 0xc1, 0x38, 0xf9, 0xa8, 0x22, 0x15, 0xa8, 0x53, 0x49, 0xdc, 0xed, 0x9f,
	0x0e, 0x63, 0x27, 0xe4, 0x4a, 0xd7, 0xe7, 0x02, 0xd6, 0x88, 0xf7, 0xbf, 0x21, 0xac, 0xef, 0x17,
	0x3a, 0x87, 0x11, 0xf8, 0xe4, 0x09, 0x6e, 0x2a, 0x51, 0x03, 0xf5, 0x90, 0xad, 0xbb, 0xdd, 0xb3,
	0xcc, 0xaa, 0xe5, 0x99, 0xd5, 0x54, 0x1d, 0x57, 0x99, 0x55, 0xb6, 0xbe, 0x9d, 0x47, 0x40, 0xcb,
	0x46, 0x72, 0x82, 0x71, 0x24, 0xf8, 0x14, 0xe4, 0x09, 0x24, 0xb1, 0x51, 0xef, 0x21, 0xbb, 0xbd,
	0x3b, 0x74, 0xee, 0x3a, 0x91, 0x73, 0x70, 0x83, 0x2d, 0xf4, 0xdd, 0x4e, 0x9e, 0x59, 0x78, 0x91,
	0xa3, 0x4b, 0xdc, 0xfd, 0xef, 0x75, 0xfc, 0x60, 0x51, 0x1a, 0xcd, 0x22, 0x2e, 0x24, 0x08, 0x65,
	0xbb, 0x87, 0x1b, 0x45, 0xa4, 0x5c, 0x37, 0xdd, 0xff, 0x2a, 0xd7, 0x8d, 0x03, 0x2e, 0x24, 0x55,
	0x15, 0xf2, 0x10, 0x37, 0x3c, 0x11, 0x14, 0x06, 0x35, 0x5b, 0x77, 0x5b, 0x45, 0x75, 0x5f, 0x04,
	0x31, 0x55, 0x59, 0xb2, 0x87, 0x35, 0x60, 0xa9, 0xa1, 0xf5, 0x34, 0xbb, 0xbd, 0xdb, 0x75, 0xca,
	0xb5, 0x29, 0x8b, 0xc5, 0xda, 0x0a, 0x9f, 0x23, 0x96, 0xbe, 0xf3, 0x84, 0xdb, 0xae, 0xa8, 0xb5,
	0x11, 0x4b, 0x69, 0x81, 0x21, 0xef, 0xb1, 0x2e, 0x20, 0xe6, 0x89, 0xf0, 0x21, 0x36, 0x1a, 0x6a,
	0x7c, 0x7b, 0x1d, 0x01, 0xad, 0x9a, 0x28, 0x7c, 0x4a, 0x42, 0x01, 0x53, 0x60, 0x32, 0x76, 0xff,
	0xaf, 0xe8, 0xf4, 0xeb, 0x6a, 0x4c, 0x17, 0x6c, 0x64, 0x8c, 0x37, 0x63, 0xf0, 0x13, 0x11, 0xca,
	0xf9, 0x0b, 0xce, 0x24, 0xcc, 0xa4, 0xd1, 0x54, 0x02, 0x8f, 0xd6, 0x09, 0x1c, 0xae, 0xb6, 0xba,
	0xf7, 0xf2, 0xcc, 0xda, 0xbc, 0x95, 0xa4, 0xb7, 0x09, 0xfb, 0x57, 0x08, 0x77, 0x56, 0xef, 0x80,
	0x30, 0xdc, 0x82, 0x6a, 0xb9, 0x6a, 0xa1, 0xed, 0xdd, 0xe7, 0x7f, 0x73, 0x9f, 0xcb, 0x17, 0xe4,
	0x6e, 0x55, 0x83, 0xb6, 0xae, 0xb3, 0xf4, 0x46, 0x83, 0xcc, 0x70, 0x27, 0x06, 0x91, 0x86, 0x3e,
	0xbc, 0x29, 0x89, 0xab, 0x57, 0xf4, 0xec, 0xee, 0xaa, 0x87, 0x2b, 0x78, 0xa5, 0x48, 0xf2, 0xcc,
	0xea, 0xac, 0xe6, 0xe9, 0x2d, 0x9d, 0xfe, 0x2f, 0x84, 0xc9, 0x9f, 0x50, 0x12, 0xe1, 0x8d, 0x89,
	0x37, 0x86, 0x49, 0x6c, 0x20, 0xf5, 0x20, 0x5e, 0xfe, 0x8b, 0x11, 0xe7, 0xb5, 0xa2, 0x1a, 0x31,
	0x29, 0xe6, 0x6e, 0xa7, 0x5a, 0xc3, 0x46, 0x99, 0xa4, 0x95, 0x0e, 0x79, 0x8c, 0x5b, 0x21, 0x93,
	0x20, 0x52, 0x6f, 0xa2, 0x86, 0xd7, 0x17, 0x0b, 0x7b, 0x55, 0xe5, 0xe9, 0x4d, 0x47, 0x77, 0x0f,
	0xb7, 0x97, 0x48, 0xc9, 0x16, 0xd6, 0x4e, 0x61, 0x5e, 0xfe, 0xb1, 0xb4, 0x38, 0x92, 0xfb, 0xb8,
	0x99, 0x7a, 0x93, 0x04, 0x4a, 0x2e, 0x5a, 0x06, 0x4f, 0xeb, 0x43, 0xe4, 0xda, 0x67, 0x97, 0x66,
	0xed, 0xfc, 0xd2, 0xac, 0x5d, 0x5c, 0x9a, 0xb5, 0x2f, 0xb9, 0x89, 0xce, 0x72, 0x13, 0x9d, 0xe7,
	0x26, 0xba, 0xc8, 0x4d, 0xf4, 0x23, 0x37, 0xd1, 0xd7, 0x9f, 0x66, 0xed, 0x43, 0x3d, 0xdd, 0xf9,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x19, 0x94, 0x8b, 0x76, 0xcf, 0x04, 0x00, 0x00,
}

func (m *AgentSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AgentSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AgentSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Prometheus != nil {
		{
			size, err := m.Prometheus.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	i -= len(m.Agent)
	copy(dAtA[i:], m.Agent)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Agent)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PrometheusExporterSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrometheusExporterSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrometheusExporterSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SecurityContext != nil {
		{
			size, err := m.SecurityContext.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	{
		size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Env) > 0 {
		for iNdEx := len(m.Env) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Env[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Args) > 0 {
		for iNdEx := len(m.Args) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Args[iNdEx])
			copy(dAtA[i:], m.Args[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Args[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	i = encodeVarintGenerated(dAtA, i, uint64(m.Port))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}

func (m *PrometheusSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PrometheusSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PrometheusSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ServiceMonitor != nil {
		{
			size, err := m.ServiceMonitor.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Exporter.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ServiceMonitorSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceMonitorSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceMonitorSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Interval)
	copy(dAtA[i:], m.Interval)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Interval)))
	i--
	dAtA[i] = 0x12
	if len(m.Labels) > 0 {
		keysForLabels := make([]string, 0, len(m.Labels))
		for k := range m.Labels {
			keysForLabels = append(keysForLabels, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
		for iNdEx := len(keysForLabels) - 1; iNdEx >= 0; iNdEx-- {
			v := m.Labels[string(keysForLabels[iNdEx])]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(keysForLabels[iNdEx])
			copy(dAtA[i:], keysForLabels[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(keysForLabels[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AgentSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Agent)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Prometheus != nil {
		l = m.Prometheus.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PrometheusExporterSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Port))
	if len(m.Args) > 0 {
		for _, s := range m.Args {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Env) > 0 {
		for _, e := range m.Env {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = m.Resources.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.SecurityContext != nil {
		l = m.SecurityContext.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *PrometheusSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Exporter.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.ServiceMonitor != nil {
		l = m.ServiceMonitor.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *ServiceMonitorSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	l = len(m.Interval)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AgentSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AgentSpec{`,
		`Agent:` + fmt.Sprintf("%v", this.Agent) + `,`,
		`Prometheus:` + strings.Replace(this.Prometheus.String(), "PrometheusSpec", "PrometheusSpec", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PrometheusExporterSpec) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForEnv := "[]EnvVar{"
	for _, f := range this.Env {
		repeatedStringForEnv += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForEnv += "}"
	s := strings.Join([]string{`&PrometheusExporterSpec{`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`Args:` + fmt.Sprintf("%v", this.Args) + `,`,
		`Env:` + repeatedStringForEnv + `,`,
		`Resources:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Resources), "ResourceRequirements", "v1.ResourceRequirements", 1), `&`, ``, 1) + `,`,
		`SecurityContext:` + strings.Replace(fmt.Sprintf("%v", this.SecurityContext), "SecurityContext", "v1.SecurityContext", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PrometheusSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PrometheusSpec{`,
		`Exporter:` + strings.Replace(strings.Replace(this.Exporter.String(), "PrometheusExporterSpec", "PrometheusExporterSpec", 1), `&`, ``, 1) + `,`,
		`ServiceMonitor:` + strings.Replace(this.ServiceMonitor.String(), "ServiceMonitorSpec", "ServiceMonitorSpec", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServiceMonitorSpec) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&ServiceMonitorSpec{`,
		`Labels:` + mapStringForLabels + `,`,
		`Interval:` + fmt.Sprintf("%v", this.Interval) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *AgentSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: AgentSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AgentSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Agent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Agent = AgentType(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prometheus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Prometheus == nil {
				m.Prometheus = &PrometheusSpec{}
			}
			if err := m.Prometheus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PrometheusExporterSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PrometheusExporterSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrometheusExporterSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Env", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Env = append(m.Env, v1.EnvVar{})
			if err := m.Env[len(m.Env)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityContext", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SecurityContext == nil {
				m.SecurityContext = &v1.SecurityContext{}
			}
			if err := m.SecurityContext.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PrometheusSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: PrometheusSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PrometheusSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exporter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Exporter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceMonitor", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ServiceMonitor == nil {
				m.ServiceMonitor = &ServiceMonitorSpec{}
			}
			if err := m.ServiceMonitor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ServiceMonitorSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: ServiceMonitorSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceMonitorSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interval = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
