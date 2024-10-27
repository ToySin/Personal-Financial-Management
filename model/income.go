package model

type Income struct {
	items map[Category]int
}

func NewIncome(items map[Category]int) *Income {
	return &Income{items: items}
}
