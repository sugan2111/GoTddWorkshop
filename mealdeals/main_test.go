package mealdeals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyTillReturnsZero(t *testing.T) {
	myTill := NewTill()
	assert.Equal(t,0,myTill.getTotal())
}

func TestOneCrisps(t *testing.T) {
	myTill := NewTill()
	myTill.scanItem("crisps")
	assert.Equal(t,90,myTill.getTotal())
}

func TestOneSandwich(t *testing.T) {
	myTill := NewTill()
	myTill.scanItem("sandwich")
	assert.Equal(t,200,myTill.getTotal())
}

func TestOneDrink(t *testing.T) {
	myTill := NewTill()
	myTill.scanItem("drink")
	assert.Equal(t,135,myTill.getTotal())
}

func TestMealDeals(t *testing.T) {
	myTill := NewTill()
	myTill.scanItem("crisps")
	myTill.scanItem("drink")
	myTill.scanItem("sandwich")

	assert.Equal(t, 300, myTill.getTotal())
}

func TestTwoCrisps(t *testing.T) {
	myTill := NewTill()
	myTill.scanItem("crisps")
	myTill.scanItem("crisps")
	assert.Equal(t,180,myTill.getTotal())
}

func TestTwoSandwich(t *testing.T) {
	myTill := NewTill()
	myTill.scanItem("sandwich")
	myTill.scanItem("sandwich")
	assert.Equal(t,400,myTill.getTotal())
}

func TestTwoMealDeals(t *testing.T) {
	myTill := NewTill()
	myTill.scanItem("crisps")
	myTill.scanItem("drink")
	myTill.scanItem("sandwich")
	myTill.scanItem("crisps")
	myTill.scanItem("drink")
	myTill.scanItem("sandwich")
	assert.Equal(t,600,myTill.getTotal())
}