package src

import (
	"log"
	"testing"
)

/*----------------------------------------------------*/
type ColorType string

const (
	SOLID ColorType = "Solid"
	DEEP  ColorType = "Deep"
)

/*----------------------------------------------------*/
type Size uint8

const (
	small      Size = 0
	medium     Size = 1
	large      Size = 2
	extraLarge Size = 3
)

/*----------------------------------------------------*/

func TestDataTypeEnum(t *testing.T) {
	log.Println(SOLID)
	log.Println(DEEP)

	log.Println(small)
	log.Println(medium)
	log.Println(large)
	log.Println(extraLarge)

	//t.Errorf("error")
	t.Logf("DONE")
}
