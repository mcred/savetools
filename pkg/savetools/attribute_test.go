package savetools_test

import (
	"SaveUtils/pkg/savetools"
	"encoding/binary"
	"testing"
)

func TestCanGet8BitValue(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x203, 8, binary.LittleEndian})
	if v != 70 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 70, v)
	}
}

func TestCanSet8BitValue(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 8, binary.LittleEndian}, 255)
	v := card1.GetValue(savetools.Attribute{0x203, 8, binary.LittleEndian})
	if v != 255 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 255, v)
	}
}

func TestCanGet16BitValue(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x1FF0, 16, binary.LittleEndian})
	if v != 46426 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 46426, v)
	}
}

func TestCanSet16BitValue(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 16, binary.LittleEndian}, 2003)
	v := card1.GetValue(savetools.Attribute{0x203, 16, binary.LittleEndian})
	if v != 2003 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 2003, v)
	}
}

func TestCanGet32BitValue(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x203, 32, binary.LittleEndian})
	if v != 4589523 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 4589523, v)
	}
}

func TestCanSet32BitValue(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 32, binary.LittleEndian}, 4294967295)
	v := card1.GetValue(savetools.Attribute{0x203, 32, binary.LittleEndian})
	if v != 4294967295 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 4294967295, v)
	}
}

func TestCanGet64BitValue(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x203, 64, binary.LittleEndian})
	if v != 2251838468390911 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 2251838468390911, v)
	}
}

func TestCanSet64BitValue(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 64, binary.LittleEndian}, 9223372036854775807)
	v := card1.GetValue(savetools.Attribute{0x203, 64, binary.LittleEndian})
	if v != 9223372036854775807 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 9223372036854775807, v)
	}
}
