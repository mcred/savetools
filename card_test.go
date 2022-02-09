package savetools_test

import (
	"encoding/binary"
	"github.com/mcred/savetools"
	"reflect"
	"testing"
)

var card1, err1 = savetools.Load("./tests/files/snes/save1.sav", 3, 0xA00)
var card2, err2 = savetools.Load("./tests/files/snes/save2.sav", 3, 2560)

func TestLoad(t *testing.T) {
	if reflect.TypeOf(card1) != reflect.TypeOf(savetools.Card{}) {
		t.Error("savetools.Card Assertion Failed")
	}
}

func TestLoadFail(t *testing.T) {
	_, err := savetools.Load("not/a/valid/path", 0, 0x00)
	if err.Error() != "unable to load file" {
		t.Errorf("savetools.Card Load Error Failed. expected %s, actual %s", "unable to load file", err.Error())
	}
}

func TestCard_Save(t *testing.T) {
	err := card2.Save()
	if err != nil {
		t.Error("savetools.Save Failed")
	}
}

func TestCard_SaveFail(t *testing.T) {
	failCard := savetools.Card{"bad/file", []byte{}, []savetools.Slot{}, 0}
	err := failCard.Save()
	if err.Error() != "unable to save file" {
		t.Errorf("savetools.Save Error Failed. expected %s, actual %s", "unable to save file", err.Error())
	}
}

func TestCard_GetValue8Bit(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x203, 8, binary.LittleEndian})
	if v != 70 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 70, v)
	}
}

func TestCard_SetValue8Bit(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 8, binary.LittleEndian}, 255)
	v := card1.GetValue(savetools.Attribute{0x203, 8, binary.LittleEndian})
	if v != 255 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 255, v)
	}
}

func TestCard_GetValue16Bit(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x1FF0, 16, binary.LittleEndian})
	if v != 46426 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 46426, v)
	}
}

func TestCard_SetValue16Bi(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 16, binary.LittleEndian}, 2003)
	v := card1.GetValue(savetools.Attribute{0x203, 16, binary.LittleEndian})
	if v != 2003 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 2003, v)
	}
}

func TestCard_GetValue32Bit(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x203, 32, binary.LittleEndian})
	if v != 4589523 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 4589523, v)
	}
}

func TestCard_SetValue32Bit(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 32, binary.LittleEndian}, 4294967295)
	v := card1.GetValue(savetools.Attribute{0x203, 32, binary.LittleEndian})
	if v != 4294967295 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 4294967295, v)
	}
}

func TestCard_GetValue64Bi(t *testing.T) {
	v := card1.GetValue(savetools.Attribute{0x203, 64, binary.LittleEndian})
	if v != 2251838468390911 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 2251838468390911, v)
	}
}

func TestCard_Set(t *testing.T) {
	card1.SetValue(savetools.Attribute{0x203, 64, binary.LittleEndian}, 9223372036854775807)
	v := card1.GetValue(savetools.Attribute{0x203, 64, binary.LittleEndian})
	if v != 9223372036854775807 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 9223372036854775807, v)
	}
}
