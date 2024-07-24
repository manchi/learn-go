package main

import "fmt"

func main() {
	var students []string
	fmt.Println("students::", students)

	students = append(students, "Ram")
	students = append(students, "Keshav")
	students = append(students, "Arjun")
	students = append(students, "Krishna")
	students = append(students, "Shiva")
	fmt.Println("students before ::", students)

	for _, student := range students {
		fmt.Printf("%v, ", student)
	}
	fmt.Println()

	students = removeStudent(students, "Arjun")
	fmt.Println("students after::", students)
}

func removeStudent(slice []string, name string) []string {
	for i, v := range slice {
		if v == name {
			fmt.Printf("removing %v\n", v)
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
