package main

import (
	"github.com/ToySin/finance/metric"
	"github.com/ToySin/finance/portfolio"
)

func main() {
	p := portfolio.GetTestPortfolio()
	w := &metric.FileWriter{}
	w.Portfolio("output.md", *p)
}
