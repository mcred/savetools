package savetools

type Slot struct {
	Start int    `json:"start"`
	Size  int    `json:"size"`
	Data  []byte `json:"data"`
}

func getValueFromBytes(b []byte, a Attribute) int {
	l := a.Location
	switch a.Bits {
	case 64:
		return int(a.Endianness.Uint64(b[l : l+8]))
	case 32:
		return int(a.Endianness.Uint32(b[l : l+4]))
	case 16:
		return int(a.Endianness.Uint16(b[l : l+2]))
	default:
		return int(b[a.Location])
	}
}

func setValueInBytes(b []byte, a Attribute, v int) {
	l := a.Location
	switch a.Bits {
	case 64:
		r := make([]byte, 8)
		a.Endianness.PutUint64(r, uint64(v))
		for i := 0; i < 8; i++ {
			b[l+i] = r[i]
		}
	case 32:
		r := make([]byte, 4)
		a.Endianness.PutUint32(r, uint32(v))
		for i := 0; i < 4; i++ {
			b[l+i] = r[i]
		}
	case 16:
		r := make([]byte, 2)
		a.Endianness.PutUint16(r, uint16(v))
		for i := 0; i < 2; i++ {
			b[l+i] = r[i]
		}
	default:
		b[a.Location] = byte(v)
	}
}

func (s *Slot) GetValue(a Attribute) int {
	return getValueFromBytes(s.Data, a)
}

func (s *Slot) SetValue(a Attribute, v int) {
	setValueInBytes(s.Data, a, v)
}

func (s *Slot) Erase() {
	s.Data = make([]byte, s.Size)
}

func (s *Slot) CopyTo(destination *Slot) {
	destination.Data = s.Data
}

func (s *Slot) CopyFrom(source *Slot) {
	s.Data = source.Data
}
