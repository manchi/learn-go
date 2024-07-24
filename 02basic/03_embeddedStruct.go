package main

import "fmt"

type Employee struct {
	FirstName  string
	LastName   string
	Age        int
	Address    Address
	Department Department
}

type Address struct {
	Street   string
	City     string
	District string
	PinCode  int
}

type Department struct {
	DepartmentName string
	Designation    string
	Manager        string
}

func (p Employee) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Employee) HasBirthday() {
	p.Age++
}
func (p *Employee) PrintInfo() {
	fmt.Println(*p)
}

func main() {
	person1 := Employee{"Ram", "Prasad", 30,
		Address{
			Street:   "Kalyan Nagar",
			City:     "Kadri",
			District: "Mangalore",
			PinCode:  574111},
		Department{
			DepartmentName: "IT",
			Designation:    "Senior Engineer",
			Manager:        "Mahesh",
		},
	}
	fmt.Println(person1)

	person2 := Employee{
		FirstName: "Krishna",
		LastName:  "Bhat",
		Age:       25,
		Address: Address{
			Street:   "Ballal Nagar",
			City:     "Kundapur",
			District: "Udupi",
			PinCode:  576121},
		Department: Department{
			DepartmentName: "HR",
			Designation:    "Senior Assistant",
			Manager:        "Rajesh",
		},
	}
	fmt.Println(person2)

	fmt.Println(person2.FullName())
	person2.HasBirthday()
	fmt.Println(person2.Age)

	person2.PrintInfo()
}
