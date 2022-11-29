package main

import (
	"fmt"
)

type product struct {
	Title    string
	Price    money
	Released timestamp
}

func (p *product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Title, p.Price, p.Released)
}

func (p *product) discount(ratio float64) {
	p.Price *= money(1 - ratio)
}
