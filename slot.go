package savetools

type Slot struct {
	Start int `json:"start"`
	Size  int `json:"size"`
}

func (c *Card) SetActiveSlot(i int) {
	c.ActiveSlot = i
}

func (c *Card) GetValueForSlot(a Attribute) uint {
	a.Location = a.Location + (c.ActiveSlot * c.Slots[c.ActiveSlot].Size)
	return c.GetValue(a)
}

func (c *Card) SetValueForSlot(a Attribute, v uint) {
	a.Location = a.Location + (c.ActiveSlot * c.Slots[c.ActiveSlot].Size)
	c.SetValue(a, v)
}
