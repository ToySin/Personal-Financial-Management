package model

type Model interface {
	Items() map[Category]int
}
