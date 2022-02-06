package savetools

import "encoding/binary"

type Attribute struct {
	Location   int              `json:"location"`
	Bits       int              `json:"bits"`
	Endianness binary.ByteOrder `json:"endianness"`
}

func (c *Card) GetValue(a Attribute) int {
	l := a.Location
	switch a.Bits {
	case 64:
		return int(a.Endianness.Uint64(c.Data[l : l+8]))
	case 32:
		return int(a.Endianness.Uint32(c.Data[l : l+4]))
	case 16:
		return int(a.Endianness.Uint16(c.Data[l : l+2]))
	default:
		return int(c.Data[a.Location])
	}
}

func (c *Card) SetValue(a Attribute, v int) {
	l := a.Location
	switch a.Bits {
	case 64:
		b := make([]byte, 8)
		a.Endianness.PutUint64(b, uint64(v))
		for i := 0; i < 8; i++ {
			c.Data[l+i] = b[i]
		}
	case 32:
		b := make([]byte, 4)
		a.Endianness.PutUint32(b, uint32(v))
		for i := 0; i < 4; i++ {
			c.Data[l+i] = b[i]
		}
	case 16:
		b := make([]byte, 2)
		a.Endianness.PutUint16(b, uint16(v))
		for i := 0; i < 2; i++ {
			c.Data[l+i] = b[i]
		}
	default:
		c.Data[a.Location] = byte(v)
	}
}
