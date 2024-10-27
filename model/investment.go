package model

type Investment struct {
	items map[Category]int
}

func NewInvestment(items map[Category]int) *Investment {
	return &Investment{items: items}
}

func (i *Investment) Items() map[Category]int {
	return i.items
}
