package savetools_test

import (
	"encoding/binary"
	"github.com/mcred/savetools"
	"testing"
)

func TestSlot_GetValue(t *testing.T) {
	card1.SetActiveSlot(1)
	slot := card1.GetActiveSlot()
	v := slot.GetValue(savetools.Attribute{0x203, 8, binary.LittleEndian})
	if v != 0 {
		t.Errorf("GetValueForSlot(0x203, false): expected %d, actual %d", 70, v)
	}
}

func TestSlot_SetValue(t *testing.T) {
	card1.Reload()
	card1.SetActiveSlot(0)
	slot := card1.GetActiveSlot()
	a := savetools.Attribute{0x203, 8, binary.LittleEndian}
	slot.SetValue(a, 22)
	v := slot.GetValue(a)
	if v != 22 {
		t.Errorf("SetValueForSlot(0x203, false): expected %d, actual %d", 22, v)
	}
	if card1.GetValue(a) != v {
		t.Errorf("Value for Card does not match: expected %d, actual %d", v, card1.GetValue(a))
	}
}

func TestSlot_Erase(t *testing.T) {
	slot := card1.GetActiveSlot()
	a := savetools.Attribute{0x203, 8, binary.LittleEndian}
	beforeValue := slot.GetValue(a) //70
	beforeSize := len(slot.Data)
	slot.Erase()
	afterValue := slot.GetValue(a) //0
	afterSize := len(slot.Data)
	if afterValue == beforeValue {
		t.Errorf("Erase Slot Failed: expected %d, actual %d", afterValue, beforeValue)
	}
	if beforeSize != afterSize {
		t.Errorf("Erased Slot Wrong Size: expected %d, actual %d", beforeSize, afterSize)
	}
	card1.Reload()
}

func TestSlot_CopyTo(t *testing.T) {
	slot1 := card1.GetSlotById(0)
	slot2 := card1.GetSlotById(2)
	a := savetools.Attribute{0x203, 8, binary.LittleEndian}
	if slot1.GetValue(a) == slot2.GetValue(a) {
		t.Errorf("Slots already match")
	}
	if slot2.GetValue(a) != 0 {
		t.Errorf("Slot 2 has incorrect starting value: expected %d, actual %d", 0, slot2.GetValue(a))
	}
	slot1.CopyTo(slot2)
	if slot1.GetValue(a) != slot2.GetValue(a) {
		t.Errorf("Slots do not match")
	}
	if slot2.GetValue(a) != 70 {
		t.Errorf("Slot 2 has incorrect ending value: expected %d, actual %d", 70, slot2.GetValue(a))
	}
	card1.Reload()
}

func TestSlot_CopyFrom(t *testing.T) {
	slot1 := card1.GetSlotById(0)
	slot2 := card1.GetSlotById(2)
	a := savetools.Attribute{0x203, 8, binary.LittleEndian}
	if slot1.GetValue(a) == slot2.GetValue(a) {
		t.Errorf("Slots already match")
	}
	if slot1.GetValue(a) != 70 {
		t.Errorf("Slot 1 has incorrect starting value: expected %d, actual %d", 0, slot1.GetValue(a))
	}
	slot1.CopyFrom(slot2)
	if slot1.GetValue(a) != slot2.GetValue(a) {
		t.Errorf("Slots do not match")
	}
	if slot1.GetValue(a) != 00 {
		t.Errorf("Slot 1 has incorrect ending value: expected %d, actual %d", 70, slot1.GetValue(a))
	}
	card1.Reload()
}
