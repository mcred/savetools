package savetools_test

import (
	"SaveUtils/pkg/savetools"
	"reflect"
	"testing"
)

var card1, err1 = savetools.Load("../../tests/files/snes/save1.sav", 3, 0xA00)
var card2, err2 = savetools.Load("../../tests/files/snes/save2.sav", 3, 2560)

func TestCanLoad(t *testing.T) {
	if reflect.TypeOf(card1) != reflect.TypeOf(savetools.Card{}) {
		t.Error("savetools.Card Assertion Failed")
	}
}

func TestCanNotLoad(t *testing.T) {
	_, err := savetools.Load("not/a/valid/path", 0, 0x00)
	if err.Error() != "unable to load file" {
		t.Errorf("savetools.Card Load Error Failed. expected %s, actual %s", "unable to load file", err.Error())
	}
}

func TestCanSave(t *testing.T) {
	err := card2.Save()
	if err != nil {
		t.Error("savetools.Save Failed")
	}
}

func TestCanNotSave(t *testing.T) {
	failCard := savetools.Card{"bad/file", []byte{}, []savetools.Slot{}, 0}
	err := failCard.Save()
	if err.Error() != "unable to save file" {
		t.Errorf("savetools.Save Error Failed. expected %s, actual %s", "unable to save file", err.Error())
	}
}
