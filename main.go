package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Height int
	Weight int
	Grade  string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

// Another struct type used in Person
type Address struct {
	Street  string
	City    string
	ZipCode int
}

// func
func sayHello(name string, round int) {
	for i := 0; i < round; i++ {
		fmt.Printf("Hello %s\n", name)
	}

}

func add(a int, b int) int {
	return a + b
}

// oop in go
type Employee struct {
	Firstname string
	Lastname  string
}

func (e Employee) Fullname() string {
	return e.Firstname + " " + e.Lastname
}

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

	//struct
	var student [3]Student
	student[0].Name = "Failman"
	student[0].Weight = 60
	student[0].Height = 180
	student[0].Grade = "F"

	student[1].Name = "Anxious"
	student[1].Weight = 60
	student[1].Height = 180
	student[1].Grade = "A"

	for i := 0; i < len(student); i++ {
		fmt.Println("Student", (i + 1))
		fmt.Println(student[i].Name)
		fmt.Println(student[i].Weight)
		fmt.Println(student[i].Height)
		fmt.Println(student[i].Grade)
	}

	students := make(map[string]Student)

	// Add Student structs to the map
	students["st01"] = Student{Name: "Mikelopster", Weight: 60, Height: 180, Grade: "F"}
	students["st02"] = Student{Name: "Alice", Weight: 55, Height: 165, Grade: "A"}
	students["st03"] = Student{Name: "Bob", Weight: 68, Height: 175, Grade: "B"}

	// Print the map
	fmt.Println("Students Map:", students)

	// Access and print a specific student by key
	fmt.Println("Student st01:", students["st01"])

	//struct in struct
	// Create an instance of the Person struct
	var person Person
	person.Name = "Alice"
	person.Age = 30
	person.Address = Address{
		Street:  "123 Main St",
		City:    "Gotham",
		ZipCode: 12345,
	}

	// Alternative way to initialize a struct
	bob := Person{
		Name: "Bob",
		Age:  25,
		Address: Address{
			Street:  "456 Elm St",
			City:    "Metropolis",
			ZipCode: 67890,
		},
	}

	// Print struct values
	fmt.Println(person)
	fmt.Println(bob)

	//function

	sayHello("fail", 1)
	sayHello("Pathetic", 2)
	sayHello("looser", 3)

	number1 := 3
	number2 := 5
	sumNumber := add(number1, number2)
	fmt.Println(sumNumber)

	employee := Employee{
		Firstname: "Fail",
		Lastname:  "Pathetic",
	}

	fullName := employee.Fullname()
	fmt.Println("Fullname of this employee is: ", fullName)

}
