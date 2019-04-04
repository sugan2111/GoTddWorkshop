package offers

import (
	"errors"
	"math"
)

type Till struct {
	Till []interface{}
	bakedBeans int
	spaghettiHoops int
}

func NewTill() *Till {
	return &Till{}
}

func (t *Till) getTotal() int {
	var offers, total = 0,0
	total += 75 * t.bakedBeans + 80 * t.spaghettiHoops

	if t.bakedBeans >= 2 {
		offers = int(math.Floor(float64(t.bakedBeans)/float64(2)))
		total -= offers * 75
	}

	if t.spaghettiHoops >= 3 {
		offers = int(math.Floor(float64(t.spaghettiHoops)/float64(3)))
		total -= offers * 40
	}
	return total
}

func (t *Till) scanItem(item string) error {
	if item == "spaghetti hoops" {
		t.spaghettiHoops++
		return nil
	} else if item == "baked beans" {
		t.bakedBeans++
		return nil
	} else {
		return errors.New("Invalid item to be scanned")
	}
}