package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//func main() {
//	var myString string
//	myString = "Green"
//
//	log.Println("myString is set to", myString)
//	changeUsingPointers(&myString)
//	log.Println("after myString is set to", myString)
//}
//
//func changeUsingPointers(s *string) {
//	newValue := "Red"
//	*s = newValue
//}

//type User struct {
//	FirstName   string
//	LastName    string
//	PhoneNumber string
//	Age         int
//	BirthDate   time.Time
//}
//
//func main() {
//	user := User{
//		FirstName:   "Bobo",
//		LastName:    "Jenkins",
//		PhoneNumber: "505-555-5555",
//		Age:         55,
//	}
//
//	log.Println(user.BirthDate)
//}

//type myStruct struct {
//	FirstName string
//}
//
//// Tying function to a struct
//func (m *myStruct) printFirstName() string {
//	return m.FirstName
//}
//
//func main() {
//	var myVar myStruct
//	myVar.FirstName = "John"
//
//	myVar2 := myStruct{
//		FirstName: "Mary",
//	}
//
//	log.Println(myVar2.printFirstName())
//	log.Println(myVar.printFirstName())
//}

//func main() {
//	// Dunno what you want? Use interface as a placeholder
//	myMap := make(map[string]string)
//
//	myMap["dog"] = "Skippy"
//
//	log.Println(myMap["dog"])
//}

//func main() {
//	//animals := []string{"dog", "fish", "horse", "cow", "cat"}
//
//	cakes := map[string]string{"drows": "chocolate", "binky": "carrot", "gangk": "strawberry"}
//
//	//Looping over a slice requires Pythonesque dictionary loops, but if you don't want to deal with the iterator, use an underscore
//	for person, cake := range cakes {
//		log.Println(person, cake)
//	}
//}
//
//type Animal interface {
//	Says() string
//	NumberOfLegs() int
//}
//
//type Dog struct {
//	Name  string
//	Breed string
//}
//
//type Gorilla struct {
//	Name          string
//	Color         string
//	NumberOfTeeth int
//}
//
//func main() {
//	dog := Dog{
//		Name:  "Sweetums",
//		Breed: "German Shepherd",
//	}
//	PrintInfo(&dog)
//
//	gorilla := Gorilla{
//		Name:          "Jock",
//		Color:         "grey",
//		NumberOfTeeth: 38,
//	}
//
//	PrintInfo(&gorilla)
//}
//
//func PrintInfo(a Animal) {
//	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
//}
//
//func (d *Dog) Says() string {
//	return "Woof"
//}
//
//func (d *Dog) NumberOfLegs() int {
//	return 4
//}
//
//func (d *Gorilla) Says() string {
//	return "Grunt"
//}
//
//func (d *Gorilla) NumberOfLegs() int {
//	return 2
//}

//func main() {
//	log.Println("Hello")
//
//	var myVar helpers.SomeType
//	myVar.TypeName = "Some Name"
//	log.Println(myVar.TypeName)
//}
//
//const numPool = 10
//
//func CalculateValue(intChan chan int) {
//	randomNumber := helpers.RandomNumber(numPool)
//	intChan <- randomNumber
//}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent", 
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne", 
		"hair_color": "black",
		"has_dog": true
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []Person

	//write json from a struct
	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Kakky"
	m2.LastName = "Kiffit"
	m2.HairColor = "blue"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "       ")

	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
