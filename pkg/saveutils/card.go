package saveutils

import "io/ioutil"

type GenerateChecksum func(*Card) uint

type Card struct {
	Path       string
	Data       []byte
	Slots      []Slot
	ActiveSlot int
}

func Load(path string, blocks int, size int) Card {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var slots []Slot
	for i := 0; i < blocks; i++ {
		slots = append(slots, Slot{i * 8, size})
	}
	return Card{Path: path, Data: f, Slots: slots, ActiveSlot: 0}
}

func (c *Card) GetChecksum(checksum GenerateChecksum) uint {
	return checksum(c)
}

func (c *Card) SetActiveSlot(i int) {
	c.ActiveSlot = i
}
