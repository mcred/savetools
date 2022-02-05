package saveutils_test

import (
	"SaveUtils/pkg/saveutils"
	"encoding/binary"
	"testing"
)

func TestCanGet8BitValue(t *testing.T) {
	v := card1.GetValue(saveutils.Attribute{0x203, 8, binary.LittleEndian})
	if v != 70 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 70, v)
	}
}

func TestCanGet16BitValue(t *testing.T) {
	v := card2.GetValue(saveutils.Attribute{0x1FF0, 16, binary.LittleEndian})
	if v != 25819 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 25819, v)
	}
}

func TestCanSet8BitValue(t *testing.T) {
	card1.SetValue(saveutils.Attribute{0x203, 8, binary.LittleEndian}, 255)
	v := card1.GetValue(saveutils.Attribute{0x203, 8, binary.LittleEndian})
	if v != 255 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 255, v)
	}
}

func TestCanSet16BitValue(t *testing.T) {
	card2.SetValue(saveutils.Attribute{0x203, 16, binary.LittleEndian}, 2003)
	v := card2.GetValue(saveutils.Attribute{0x203, 16, binary.LittleEndian})
	if v != 2003 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 2003, v)
	}
}
