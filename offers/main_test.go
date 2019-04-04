package offers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyTillReturnsZero(t *testing.T) {
	myTill := NewTill()
	assert.Equal(t,0,myTill.getTotal())
}

func TestOneBakedBeans(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("baked beans")
	assert.Equal(t, 75, myTill.getTotal())
	assert.NoError(t,err)
}

func TestThreeBakedBeans(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("baked beans")
	err1 := myTill.scanItem("baked beans")
	err2 := myTill.scanItem("baked beans")
	assert.Equal(t,150,myTill.getTotal())
	assert.NoError(t,err)
	assert.NoError(t,err1)
	assert.NoError(t,err2)
}

func TestFourBakedBeans(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("baked beans")
	err1 := myTill.scanItem("baked beans")
	err2 := myTill.scanItem("baked beans")
	err3 := myTill.scanItem("baked beans")
	assert.Equal(t,150,myTill.getTotal())
	assert.NoError(t,err)
	assert.NoError(t,err1)
	assert.NoError(t,err2)
	assert.NoError(t,err3)
}

func TestBOGOF(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("baked beans")
	err1 := myTill.scanItem("baked beans")
	assert.Equal(t, 75, myTill.getTotal())
	assert.NoError(t,err)
	assert.NoError(t,err1)
}

func TestOneSpaghettiHoop(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("spaghetti hoops")
	assert.Equal(t,80,myTill.getTotal())
	assert.NoError(t, err)
}

func TestBTGTHP(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("spaghetti hoops")
	err1 := myTill.scanItem("spaghetti hoops")
	err2 := myTill.scanItem("spaghetti hoops")
	assert.Equal(t,200,myTill.getTotal())
	assert.NoError(t,err)
	assert.NoError(t,err1)
	assert.NoError(t,err2)
}

func TestTwoBTGTHP(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("spaghetti hoops")
	err1 := myTill.scanItem("spaghetti hoops")
	err2 := myTill.scanItem("spaghetti hoops")
	err3 := myTill.scanItem("spaghetti hoops")
	err4 := myTill.scanItem("spaghetti hoops")
	err5 := myTill.scanItem("spaghetti hoops")
	assert.Equal(t,400,myTill.getTotal())
	assert.NoError(t,err)
	assert.NoError(t,err1)
	assert.NoError(t,err2)
	assert.NoError(t,err3)
	assert.NoError(t,err4)
	assert.NoError(t,err5)
}

func TestTwoBeansThreeHoops(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("baked beans")
	err1 := myTill.scanItem("baked beans")
	err2 := myTill.scanItem("spaghetti hoops")
	err3 := myTill.scanItem("spaghetti hoops")
	err4 := myTill.scanItem("spaghetti hoops")
	assert.Equal(t,275,myTill.getTotal())
	assert.NoError(t,err)
	assert.NoError(t,err1)
	assert.NoError(t,err2)
	assert.NoError(t,err3)
	assert.NoError(t,err4)
}

func TestReturnsErrorWhenScannedItemIsInvalid(t *testing.T) {
	myTill := NewTill()
	err := myTill.scanItem("white bread")
	assert.Error(t, err, "Invalid item to be scanned")
}