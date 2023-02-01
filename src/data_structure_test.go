package src

import (
	"log"
	"testing"
)

func TestDataStructure(t *testing.T) {
	arrayExampleWithOutDataInitialization()
	arrayExampleWithDataInitialization()
	multidimensionalArrayExample()

	sliceExample()
	emptySliceExample()
	multidimensionalSliceExample()

	mapExample()
}

func mapExample() {
	//map[key_type]value_type
	mapDataType := map[string]interface{}{
		"1": "test1",
		"2": "test2",
		"3": "test3",
	}

	mapDataType["4"] = "test4"
	mapDataType["5"] = "test5"

	for key, value := range mapDataType {
		log.Println("key: ", key, " value: ", value)
	}

	mapDataType["2"] = "Rakib"

	getValueByKey := mapDataType["1"]
	log.Println(getValueByKey)

	delete(mapDataType, "3")
	log.Println(mapDataType)

}

func multidimensionalSliceExample() {
	var multidimensionalSlice [][]interface{}

	multidimensionalSlice = append(multidimensionalSlice, []interface{}{"1", "test1", "hello"})
	multidimensionalSlice = append(multidimensionalSlice, []interface{}{"2", "test2"})
	multidimensionalSlice = append(multidimensionalSlice, []interface{}{"3", "test3"})

	for i, data := range multidimensionalSlice {
		log.Println(i, " : ", data)
	}
}

func emptySliceExample() {
	var emptySlice []string

	log.Println(emptySlice)

	emptySlice = append(emptySlice, "test1")
	emptySlice = append(emptySlice, "test2")
	emptySlice = append(emptySlice, "test3")
	emptySlice = append(emptySlice, "test4")

	for index, data := range emptySlice {
		log.Println("using index ", index, " data ", data)
		log.Println("without index ", emptySlice[index])
	}
}

func sliceExample() {
	var sliceWillBeWithOutLength = []string{"Rakib"}

	for index, data := range sliceWillBeWithOutLength {
		log.Println(index, " ", data)
	}

	sliceWillBeWithOutLength = append(sliceWillBeWithOutLength, "Rifat")
	sliceWillBeWithOutLength = append(sliceWillBeWithOutLength, "Rony")

	i2 := len(sliceWillBeWithOutLength)
	log.Println(i2)
	for i := 0; i < i2; i++ {
		log.Println(sliceWillBeWithOutLength[i])
	}
}

func multidimensionalArrayExample() {
	var multiDimensionalArray = [3][2]interface{}{{1, 2}, {1, 2}, {1, 2}}
	for index, data := range multiDimensionalArray {
		log.Println("Data of column: ", data)
		for _, value := range data {
			log.Println("data of row: ", index, " data ", value)
		}
	}
}

func arrayExampleWithDataInitialization() {
	//var arrayDemo [3]interface{}
	//arrayDemo := [3]interface{}{1, 2, 3}
	var arrayDemo = [3]interface{}{1, 2, 3}
	for index := range arrayDemo {
		log.Println("without casting interface ", arrayDemo[index])
		log.Println("after casting interface to int ", arrayDemo[index].(int))
	}
}

func arrayExampleWithOutDataInitialization() {
	//var arrayWithFixedLength [3]string{}
	//arrayWithFixedLength := [3]string{}
	//arrayWithFixedLength := [...]string{}
	var arrayWithFixedLength = [3]string{}

	arrayWithFixedLength[0] = "rakib0"
	arrayWithFixedLength[1] = "rakib1"
	arrayWithFixedLength[2] = "rakib2"

	for index, data := range arrayWithFixedLength {
		log.Println(index, " ", data)
	}
}
