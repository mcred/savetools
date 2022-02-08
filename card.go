package savetools

import (
	"errors"
	"io/ioutil"
)

type Card struct {
	Path       string `json:"path"`
	Data       []byte `json:"data"`
	Slots      []Slot `json:"slots"`
	ActiveSlot int    `json:"activeSlot"`
}

func Load(path string, blocks int, size int) (Card, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return Card{}, errors.New("unable to load file")
	}
	var slots []Slot
	for i := 0; i < blocks; i++ {
		start := i * size
		slots = append(slots, Slot{start, size, f[start : start+size]})
	}
	return Card{Path: path, Data: f, Slots: slots, ActiveSlot: 0}, nil
}

func (c *Card) Save() error {
	err := ioutil.WriteFile(c.Path, c.Data, 0644)
	if err != nil {
		return errors.New("unable to save file")
	}
	return nil
}

func (c *Card) GetValue(a Attribute) int {
	return getValueFromBytes(c.Data, a)
}

func (c *Card) SetValue(a Attribute, v int) {
	setValueInBytes(c.Data, a, v)
}

func (c *Card) SetActiveSlot(i int) {
	c.ActiveSlot = i
}

func (c *Card) GetActiveSlot() Slot {
	return c.Slots[c.ActiveSlot]
}
