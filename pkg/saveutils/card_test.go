package saveutils_test

import (
	"SaveUtils/pkg/saveutils"
	"encoding/binary"
	"reflect"
	"testing"
)

var card1 = saveutils.Load("../../tests/files/snes/save1.sav", 3, 0xA00)
var card2 = saveutils.Load("../../tests/files/snes/save2.sav", 3, 2560)

func TestCanLoad(t *testing.T) {
	if reflect.TypeOf(card1) != reflect.TypeOf(saveutils.Card{}) {
		t.Error("saveutils.Card Assertion Failed")
	}
}

/*
Thanks to https://github.com/mikearnos/snessum for help with this
*/
func generateChecksum(c *saveutils.Card) uint {
	var checksum uint = 0
	for slot := 1; slot <= 3; slot++ {
		max := (slot * 0xA00) - 2
		min := (slot - 1) * 0xA00
		for i := max; i >= min; i -= 2 {
			if checksum > 0xFFFF {
				checksum -= 0xFFFF
			}
			checksum += uint(c.GetValue(saveutils.Attribute{i, 16, binary.LittleEndian}))
		}
		checksum &= 0xFFFF
	}
	return checksum
}

func TestCanGenerateCheckSum(t *testing.T) {
	card := saveutils.Load("../../tests/files/snes/save1.sav", 3, 0xA00)
	c := card.GetChecksum(generateChecksum)
	if c != 46426 {
		t.Errorf("Checksum invalid: expected %d, actual %d", 46426, c)
	}
}
