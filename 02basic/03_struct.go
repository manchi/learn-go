package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) HasBirthday() {
	p.Age++
}

func main() {
	person1 := Person{"Rama", "Bhagavath", 30}
	fmt.Println(person1)

	person2 := Person{FirstName: "Krishna", LastName: "Bhat", Age: 25}
	fmt.Println(person2)

	person3 := new(Person)
	person3.FirstName = "Bharath"
	person3.LastName = "Acharya"
	person3.Age = 33
	fmt.Println(person3)
	fmt.Println(&person3)
	fmt.Println(*person3)

	fmt.Println(person1.FullName())

	person2.HasBirthday()
	fmt.Println(person2.Age)

}
