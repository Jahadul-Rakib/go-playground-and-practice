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

	log.Println(" \n ")
	//another approach use of interface
	castToCatFromInterface := animalInterfaceForCat.(*Cat)
	log.Println(castToCatFromInterface)
}
