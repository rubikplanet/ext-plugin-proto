// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package ExtraInfo

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type RespBody struct {
	_tab flatbuffers.Table
}

func GetRootAsRespBody(buf []byte, offset flatbuffers.UOffsetT) *RespBody {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &RespBody{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsRespBody(buf []byte, offset flatbuffers.UOffsetT) *RespBody {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &RespBody{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *RespBody) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *RespBody) Table() flatbuffers.Table {
	return rcv._tab
}

func RespBodyStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func RespBodyEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
