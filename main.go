package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	subSlice := mySlice[1:3]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice))
	fmt.Println(cap(subSlice))
	// list
	//append
	var appSlice []int
	appSlice = append(appSlice, 10)
	appSlice = append(appSlice, 20, 30)

	fmt.Println(appSlice)

	//map
	myMap := make(map[string]int)

	myMap["apple"] = 5
	myMap["banana"] = 10
	myMap["orange"] = 8

	// Access and print a value for a key
	fmt.Println("Apples:", myMap["apple"])

	delete(myMap, "banana")

	//for loop key val map
	for key, value := range myMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	myMap["pear"] = 45

	// Checking if key exist
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value: ", val)
	} else {
		fmt.Println("Pear not found in map")
	}

}
