// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package A6

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TextEntry struct {
	_tab flatbuffers.Table
}

func GetRootAsTextEntry(buf []byte, offset flatbuffers.UOffsetT) *TextEntry {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TextEntry{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsTextEntry(buf []byte, offset flatbuffers.UOffsetT) *TextEntry {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &TextEntry{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *TextEntry) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TextEntry) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TextEntry) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TextEntry) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TextEntryStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func TextEntryAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func TextEntryAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func TextEntryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}