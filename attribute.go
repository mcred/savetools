package savetools

import "encoding/binary"

type Attribute struct {
	Location   int              `json:"location"`
	Bits       int              `json:"bits"`
	Endianness binary.ByteOrder `json:"endianness"`
}
