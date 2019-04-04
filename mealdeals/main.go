package mealdeals

import "math"

type Till struct {
	Till []interface{}
	crisps int
	drink int
	sandwich int
	offers int
}

func NewTill() *Till {
	return &Till{}
}

func (t *Till) getTotal() int {

	var total = 0
	if t.sandwich >= 1 && t.drink >= 1 && t.crisps >= 1 {
		t.offers = int(math.Min(float64(t.sandwich),math.Min(float64(t.crisps),float64(t.drink))))
		total += t.offers * 300
	} else {
		total += 200 * t.sandwich + 90 * t.crisps + 135 * t.drink
	}
	return total

}

func (t *Till) scanItem(item string) {
	if item == "sandwich" {
		t.sandwich++
	} else if item == "crisps" {
		t.crisps++
	} else if item == "drink" {
		t.drink++
	}
}