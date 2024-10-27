package model

type FixedExpense struct {
	items map[Category]int
}

func NewFixedExpense(items map[Category]int) *FixedExpense {
	return &FixedExpense{items: items}
}

func (f *FixedExpense) Items() map[Category]int {
	return f.items
}
