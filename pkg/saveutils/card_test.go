package saveutils_test

import (
	"SaveUtils/pkg/saveutils"
	"reflect"
	"testing"
)

var card1, err1 = saveutils.Load("../../tests/files/snes/save1.sav", 3, 0xA00)
var card2, err2 = saveutils.Load("../../tests/files/snes/save2.sav", 3, 2560)

func TestCanLoad(t *testing.T) {
	if reflect.TypeOf(card1) != reflect.TypeOf(saveutils.Card{}) {
		t.Error("saveutils.Card Assertion Failed")
	}
}

func TestCanNotLoad(t *testing.T) {
	_, err := saveutils.Load("not/a/valid/path", 0, 0x00)
	if err.Error() != "unable to load file" {
		t.Errorf("saveutils.Card Load Error Failed. expected %s, actual %s", "unable to load file", err.Error())
	}
}

func TestCanSave(t *testing.T) {
	err := card2.Save()
	if err != nil {
		t.Error("saveutils.Save Failed")
	}
}

func TestCanNotSave(t *testing.T) {
	failCard := saveutils.Card{"bad/file", []byte{}, []saveutils.Slot{}, 0}
	err := failCard.Save()
	if err.Error() != "unable to save file" {
		t.Errorf("saveutils.Save Error Failed. expected %s, actual %s", "unable to save file", err.Error())
	}
}
