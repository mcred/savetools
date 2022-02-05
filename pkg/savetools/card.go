package savetools

import (
	"errors"
	"io/ioutil"
)

type Card struct {
	Path       string
	Data       []byte
	Slots      []Slot
	ActiveSlot int
}

func Load(path string, blocks int, size int) (Card, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return Card{}, errors.New("unable to load file")
	}
	var slots []Slot
	for i := 0; i < blocks; i++ {
		slots = append(slots, Slot{i * 8, size})
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
