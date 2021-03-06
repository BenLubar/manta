// Code generated by protoc-gen-go.
// source: toolevents.proto
// DO NOT EDIT!

package dota

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ChangeMapToolEvent struct {
	Mapname          *string `protobuf:"bytes,1,opt,name=mapname" json:"mapname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChangeMapToolEvent) Reset()         { *m = ChangeMapToolEvent{} }
func (m *ChangeMapToolEvent) String() string { return proto.CompactTextString(m) }
func (*ChangeMapToolEvent) ProtoMessage()    {}

func (m *ChangeMapToolEvent) GetMapname() string {
	if m != nil && m.Mapname != nil {
		return *m.Mapname
	}
	return ""
}

type TraceRayServerToolEvent struct {
	Start            *CMsgVector `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End              *CMsgVector `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TraceRayServerToolEvent) Reset()         { *m = TraceRayServerToolEvent{} }
func (m *TraceRayServerToolEvent) String() string { return proto.CompactTextString(m) }
func (*TraceRayServerToolEvent) ProtoMessage()    {}

func (m *TraceRayServerToolEvent) GetStart() *CMsgVector {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TraceRayServerToolEvent) GetEnd() *CMsgVector {
	if m != nil {
		return m.End
	}
	return nil
}

type ToolTraceRayResult struct {
	Hit              *bool       `protobuf:"varint,1,opt,name=hit" json:"hit,omitempty"`
	Impact           *CMsgVector `protobuf:"bytes,2,opt,name=impact" json:"impact,omitempty"`
	Normal           *CMsgVector `protobuf:"bytes,3,opt,name=normal" json:"normal,omitempty"`
	Distance         *float32    `protobuf:"fixed32,4,opt,name=distance" json:"distance,omitempty"`
	Fraction         *float32    `protobuf:"fixed32,5,opt,name=fraction" json:"fraction,omitempty"`
	Ehandle          *int32      `protobuf:"varint,6,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ToolTraceRayResult) Reset()         { *m = ToolTraceRayResult{} }
func (m *ToolTraceRayResult) String() string { return proto.CompactTextString(m) }
func (*ToolTraceRayResult) ProtoMessage()    {}

func (m *ToolTraceRayResult) GetHit() bool {
	if m != nil && m.Hit != nil {
		return *m.Hit
	}
	return false
}

func (m *ToolTraceRayResult) GetImpact() *CMsgVector {
	if m != nil {
		return m.Impact
	}
	return nil
}

func (m *ToolTraceRayResult) GetNormal() *CMsgVector {
	if m != nil {
		return m.Normal
	}
	return nil
}

func (m *ToolTraceRayResult) GetDistance() float32 {
	if m != nil && m.Distance != nil {
		return *m.Distance
	}
	return 0
}

func (m *ToolTraceRayResult) GetFraction() float32 {
	if m != nil && m.Fraction != nil {
		return *m.Fraction
	}
	return 0
}

func (m *ToolTraceRayResult) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type SpawnEntityToolEvent struct {
	EntityKeyvalues  []byte `protobuf:"bytes,1,opt,name=entity_keyvalues" json:"entity_keyvalues,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SpawnEntityToolEvent) Reset()         { *m = SpawnEntityToolEvent{} }
func (m *SpawnEntityToolEvent) String() string { return proto.CompactTextString(m) }
func (*SpawnEntityToolEvent) ProtoMessage()    {}

func (m *SpawnEntityToolEvent) GetEntityKeyvalues() []byte {
	if m != nil {
		return m.EntityKeyvalues
	}
	return nil
}

func (m *SpawnEntityToolEvent) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type SpawnEntityToolEventResult struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SpawnEntityToolEventResult) Reset()         { *m = SpawnEntityToolEventResult{} }
func (m *SpawnEntityToolEventResult) String() string { return proto.CompactTextString(m) }
func (*SpawnEntityToolEventResult) ProtoMessage()    {}

func (m *SpawnEntityToolEventResult) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type DestroyEntityToolEvent struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DestroyEntityToolEvent) Reset()         { *m = DestroyEntityToolEvent{} }
func (m *DestroyEntityToolEvent) String() string { return proto.CompactTextString(m) }
func (*DestroyEntityToolEvent) ProtoMessage()    {}

func (m *DestroyEntityToolEvent) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

type DestroyAllEntitiesToolEvent struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DestroyAllEntitiesToolEvent) Reset()         { *m = DestroyAllEntitiesToolEvent{} }
func (m *DestroyAllEntitiesToolEvent) String() string { return proto.CompactTextString(m) }
func (*DestroyAllEntitiesToolEvent) ProtoMessage()    {}

type RestartMapToolEvent struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RestartMapToolEvent) Reset()         { *m = RestartMapToolEvent{} }
func (m *RestartMapToolEvent) String() string { return proto.CompactTextString(m) }
func (*RestartMapToolEvent) ProtoMessage()    {}

type ToolEvent_GetEntityInfo struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_GetEntityInfo) Reset()         { *m = ToolEvent_GetEntityInfo{} }
func (m *ToolEvent_GetEntityInfo) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInfo) ProtoMessage()    {}

func (m *ToolEvent_GetEntityInfo) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_GetEntityInfo) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInfoResult struct {
	Cppclass         *string     `protobuf:"bytes,1,opt,name=cppclass,def=shithead" json:"cppclass,omitempty"`
	Classname        *string     `protobuf:"bytes,2,opt,name=classname" json:"classname,omitempty"`
	Name             *string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Origin           *CMsgVector `protobuf:"bytes,4,opt,name=origin" json:"origin,omitempty"`
	Mins             *CMsgVector `protobuf:"bytes,5,opt,name=mins" json:"mins,omitempty"`
	Maxs             *CMsgVector `protobuf:"bytes,6,opt,name=maxs" json:"maxs,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ToolEvent_GetEntityInfoResult) Reset()         { *m = ToolEvent_GetEntityInfoResult{} }
func (m *ToolEvent_GetEntityInfoResult) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInfoResult) ProtoMessage()    {}

const Default_ToolEvent_GetEntityInfoResult_Cppclass string = "shithead"

func (m *ToolEvent_GetEntityInfoResult) GetCppclass() string {
	if m != nil && m.Cppclass != nil {
		return *m.Cppclass
	}
	return Default_ToolEvent_GetEntityInfoResult_Cppclass
}

func (m *ToolEvent_GetEntityInfoResult) GetClassname() string {
	if m != nil && m.Classname != nil {
		return *m.Classname
	}
	return ""
}

func (m *ToolEvent_GetEntityInfoResult) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ToolEvent_GetEntityInfoResult) GetOrigin() *CMsgVector {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *ToolEvent_GetEntityInfoResult) GetMins() *CMsgVector {
	if m != nil {
		return m.Mins
	}
	return nil
}

func (m *ToolEvent_GetEntityInfoResult) GetMaxs() *CMsgVector {
	if m != nil {
		return m.Maxs
	}
	return nil
}

type ToolEvent_GetEntityInputs struct {
	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_GetEntityInputs) Reset()         { *m = ToolEvent_GetEntityInputs{} }
func (m *ToolEvent_GetEntityInputs) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInputs) ProtoMessage()    {}

func (m *ToolEvent_GetEntityInputs) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_GetEntityInputs) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInputsResult struct {
	InputList        []string `protobuf:"bytes,1,rep,name=input_list" json:"input_list,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ToolEvent_GetEntityInputsResult) Reset()         { *m = ToolEvent_GetEntityInputsResult{} }
func (m *ToolEvent_GetEntityInputsResult) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_GetEntityInputsResult) ProtoMessage()    {}

func (m *ToolEvent_GetEntityInputsResult) GetInputList() []string {
	if m != nil {
		return m.InputList
	}
	return nil
}

type ToolEvent_FireEntityInput struct {
	Ehandle          *int32  `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool   `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	InputName        *string `protobuf:"bytes,3,opt,name=input_name" json:"input_name,omitempty"`
	InputParam       *string `protobuf:"bytes,4,opt,name=input_param" json:"input_param,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ToolEvent_FireEntityInput) Reset()         { *m = ToolEvent_FireEntityInput{} }
func (m *ToolEvent_FireEntityInput) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_FireEntityInput) ProtoMessage()    {}

func (m *ToolEvent_FireEntityInput) GetEhandle() int32 {
	if m != nil && m.Ehandle != nil {
		return *m.Ehandle
	}
	return 0
}

func (m *ToolEvent_FireEntityInput) GetClientsideentity() bool {
	if m != nil && m.Clientsideentity != nil {
		return *m.Clientsideentity
	}
	return false
}

func (m *ToolEvent_FireEntityInput) GetInputName() string {
	if m != nil && m.InputName != nil {
		return *m.InputName
	}
	return ""
}

func (m *ToolEvent_FireEntityInput) GetInputParam() string {
	if m != nil && m.InputParam != nil {
		return *m.InputParam
	}
	return ""
}

type ToolEvent_SFMRecordingStateChanged struct {
	Isrecording      *bool  `protobuf:"varint,1,opt,name=isrecording" json:"isrecording,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_SFMRecordingStateChanged) Reset()         { *m = ToolEvent_SFMRecordingStateChanged{} }
func (m *ToolEvent_SFMRecordingStateChanged) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_SFMRecordingStateChanged) ProtoMessage()    {}

func (m *ToolEvent_SFMRecordingStateChanged) GetIsrecording() bool {
	if m != nil && m.Isrecording != nil {
		return *m.Isrecording
	}
	return false
}

type ToolEvent_SFMToolActiveStateChanged struct {
	Isactive         *bool  `protobuf:"varint,1,opt,name=isactive" json:"isactive,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ToolEvent_SFMToolActiveStateChanged) Reset()         { *m = ToolEvent_SFMToolActiveStateChanged{} }
func (m *ToolEvent_SFMToolActiveStateChanged) String() string { return proto.CompactTextString(m) }
func (*ToolEvent_SFMToolActiveStateChanged) ProtoMessage()    {}

func (m *ToolEvent_SFMToolActiveStateChanged) GetIsactive() bool {
	if m != nil && m.Isactive != nil {
		return *m.Isactive
	}
	return false
}

func init() {
}
