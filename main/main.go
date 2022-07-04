package main

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
