package model

type Saving struct {
	items map[Category]int
}

func NewSaving(items map[Category]int) *Saving {
	return &Saving{items: items}
}

func (s *Saving) Items() map[Category]int {
	return s.items
}
