package savetools_test

import (
	"encoding/binary"
	"github.com/mcred/savetools"
	"testing"
)

func TestCanGetValueBySlot(t *testing.T) {
	card1.SetActiveSlot(1)
	slot := card1.GetActiveSlot()
	v := slot.GetValue(savetools.Attribute{0x203, 8, binary.LittleEndian})
	if v != 0 {
		t.Errorf("GetValueForSlot(0x203, false): expected %d, actual %d", 70, v)
	}
}

func TestCanSetValueBySlot(t *testing.T) {
	card1.SetActiveSlot(1)
	slot := card1.GetActiveSlot()
	a := savetools.Attribute{0x203, 8, binary.LittleEndian}
	slot.SetValue(a, 22)
	v := slot.GetValue(a)
	if v != 22 {
		t.Errorf("SetValueForSlot(0x203, false): expected %d, actual %d", 22, v)
	}
}
