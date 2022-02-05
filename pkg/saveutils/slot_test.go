package saveutils_test

import (
	"SaveUtils/pkg/saveutils"
	"encoding/binary"
	"testing"
)

func TestCanGetValueAtSlot(t *testing.T) {
	card1.SetActiveSlot(1)
	v := card1.GetValueForSlot(saveutils.Attribute{0x203, 8, binary.LittleEndian})
	if v != 0 {
		t.Errorf("GetValueForSlot(0x203, false): expected %d, actual %d", 70, v)
	}
}

func TestCanSetValueAtSlot(t *testing.T) {
	card1.SetActiveSlot(1)
	a := saveutils.Attribute{0x203, 8, binary.LittleEndian}
	card1.SetValueForSlot(a, 22)
	v := card1.GetValueForSlot(a)
	if v != 22 {
		t.Errorf("SetValueForSlot(0x203, false): expected %d, actual %d", 22, v)
	}
}
