package saveutils

type Slot struct {
	Start int
	Size  int
}

func (c *Card) GetSlot(i int) Slot {
	return c.Slots[i]
}
