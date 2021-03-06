package rpc

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *MsgpProcessor) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z MsgpProcessor) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 0
	err = en.Append(0x80)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MsgpProcessor) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 0
	o = append(o, 0x80)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgpProcessor) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MsgpProcessor) Msgsize() (s int) {
	s = 1
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MsgpRpcRequestData) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "ServiceMethod":
			z.ServiceMethod, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ServiceMethod")
				return
			}
		case "NoReply":
			z.NoReply, err = dc.ReadBool()
			if err != nil {
				err = msgp.WrapError(err, "NoReply")
				return
			}
		case "InParam":
			z.InParam, err = dc.ReadBytes(z.InParam)
			if err != nil {
				err = msgp.WrapError(err, "InParam")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MsgpRpcRequestData) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Seq"
	err = en.Append(0x84, 0xa3, 0x53, 0x65, 0x71)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Seq)
	if err != nil {
		err = msgp.WrapError(err, "Seq")
		return
	}
	// write "ServiceMethod"
	err = en.Append(0xad, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.ServiceMethod)
	if err != nil {
		err = msgp.WrapError(err, "ServiceMethod")
		return
	}
	// write "NoReply"
	err = en.Append(0xa7, 0x4e, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBool(z.NoReply)
	if err != nil {
		err = msgp.WrapError(err, "NoReply")
		return
	}
	// write "InParam"
	err = en.Append(0xa7, 0x49, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.InParam)
	if err != nil {
		err = msgp.WrapError(err, "InParam")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MsgpRpcRequestData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Seq"
	o = append(o, 0x84, 0xa3, 0x53, 0x65, 0x71)
	o = msgp.AppendUint64(o, z.Seq)
	// string "ServiceMethod"
	o = append(o, 0xad, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64)
	o = msgp.AppendString(o, z.ServiceMethod)
	// string "NoReply"
	o = append(o, 0xa7, 0x4e, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79)
	o = msgp.AppendBool(o, z.NoReply)
	// string "InParam"
	o = append(o, 0xa7, 0x49, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d)
	o = msgp.AppendBytes(o, z.InParam)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgpRpcRequestData) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "ServiceMethod":
			z.ServiceMethod, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ServiceMethod")
				return
			}
		case "NoReply":
			z.NoReply, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NoReply")
				return
			}
		case "InParam":
			z.InParam, bts, err = msgp.ReadBytesBytes(bts, z.InParam)
			if err != nil {
				err = msgp.WrapError(err, "InParam")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MsgpRpcRequestData) Msgsize() (s int) {
	s = 1 + 4 + msgp.Uint64Size + 14 + msgp.StringPrefixSize + len(z.ServiceMethod) + 8 + msgp.BoolSize + 8 + msgp.BytesPrefixSize + len(z.InParam)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MsgpRpcResponseData) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Err":
			z.Err, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Err")
				return
			}
		case "Reply":
			z.Reply, err = dc.ReadBytes(z.Reply)
			if err != nil {
				err = msgp.WrapError(err, "Reply")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MsgpRpcResponseData) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Seq"
	err = en.Append(0x83, 0xa3, 0x53, 0x65, 0x71)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Seq)
	if err != nil {
		err = msgp.WrapError(err, "Seq")
		return
	}
	// write "Err"
	err = en.Append(0xa3, 0x45, 0x72, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Err)
	if err != nil {
		err = msgp.WrapError(err, "Err")
		return
	}
	// write "Reply"
	err = en.Append(0xa5, 0x52, 0x65, 0x70, 0x6c, 0x79)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Reply)
	if err != nil {
		err = msgp.WrapError(err, "Reply")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MsgpRpcResponseData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Seq"
	o = append(o, 0x83, 0xa3, 0x53, 0x65, 0x71)
	o = msgp.AppendUint64(o, z.Seq)
	// string "Err"
	o = append(o, 0xa3, 0x45, 0x72, 0x72)
	o = msgp.AppendString(o, z.Err)
	// string "Reply"
	o = append(o, 0xa5, 0x52, 0x65, 0x70, 0x6c, 0x79)
	o = msgp.AppendBytes(o, z.Reply)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgpRpcResponseData) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Seq":
			z.Seq, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Seq")
				return
			}
		case "Err":
			z.Err, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Err")
				return
			}
		case "Reply":
			z.Reply, bts, err = msgp.ReadBytesBytes(bts, z.Reply)
			if err != nil {
				err = msgp.WrapError(err, "Reply")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MsgpRpcResponseData) Msgsize() (s int) {
	s = 1 + 4 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.Err) + 6 + msgp.BytesPrefixSize + len(z.Reply)
	return
}
