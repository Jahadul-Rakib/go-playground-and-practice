package src

import (
	"log"
	"testing"
)

// interface and method-------------------------------------------------------
type animal interface {
	breathe()
}

type Cat struct {
	Name string `json:"name"`
}

type Tiger struct {
	Name string `json:"name"`
}

func (c *Cat) breathe() {
	c.Name = "my name is " + c.Name
}

func (t *Tiger) breathe() {
	t.Name = "my name is " + t.Name
}

// interface and function-------------------------------------------------------
type Car interface {
	engineType(t string) string
}
type T struct{}

func (T) engineType(t string) string {
	log.Print("inside--- " + t + " ---inside")
	return t
}

// another approch of interface
type Engine interface {
	model(modelName string) string
	createCar(fuelCapacity int, model string) interface{}
}
type NewCar struct {
	model        string
	fuelCapacity int
}
type CarName string

func (c CarName) createCar(fuelCapacity int, model string) interface{} {
	return NewCar{model: model, fuelCapacity: fuelCapacity}
}

func (c CarName) model(modelName string) string {
	return modelName
}

func Test_INTERFACE(t *testing.T) {
	//use of interface
	animalInterfaceForCat := animal(&Cat{Name: "C Rakib"})
	animalInterfaceForCat.breathe()
	log.Print(animalInterfaceForCat)

	//use of interface
	tiger := Tiger{Name: "T Rakib"}
	animal(&tiger).breathe()
	log.Println(tiger)

	x := Car(T{}).engineType("from method")
	log.Print(x)

	//another approach use of interface
	castToCatFromInterface := animalInterfaceForCat.(*Cat)
	log.Println(castToCatFromInterface)

	data := Engine(new(CarName))
	model := data.model("A300")
	log.Println(model)
	fuelCapacity := data.createCar(120, model)
	log.Println(fuelCapacity)

	//convert interface to struct
	newCar := fuelCapacity.(NewCar)
	log.Println("car model: ", newCar.model)
	log.Println("car fuel-capacity: ", newCar.fuelCapacity)
}
