package main

import (
	"log/slog"
)

func main() {
	slog.Info("hello, world")

	//sample_slog()
}

func sample_slog() {
	// read a number
	var num1 int = 3030
	slog.Info("your number is: %v ", num1)

	// read a string
	var name string = "programming"
	slog.Info("your string input is:", name)
	slog.Info("the type of the variable name is: %T\n", name)

	slog.Info("this is a float: %.4f\n", 123.123123)

	slog.Info("a for loop using int x:")
	for i := 0; i < len(name); i++ {
		slog.Any("%x,", name[i])
	}

	slog.Info("\na for loop using byte b:")
	for i := 0; i < len(name); i++ {
		slog.Info("%b,", name[i])
	}

	slog.Info("\na for loop using char c:")
	for i := 0; i < len(name); i++ {
		slog.Info("%c,", name[i])
	}
}
