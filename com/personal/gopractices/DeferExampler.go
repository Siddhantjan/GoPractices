package main

import "fmt"

type Worker struct {
	Name string
}

func (e *Worker) UpdateName(name string) {
	e.Name = name
}

func doStuff(e *Worker) {
	name := "Elliot Forbes"
	defer e.UpdateName(name)
	fmt.Println("doing other exciting stuff")

	name = "Elliot M Forbes"
}

func main() {
	fmt.Println("defer in go")
	elliot := &Worker{
		Name: "Elliot",
	}

	fmt.Printf("%+v\n", elliot)
	doStuff(elliot)
	fmt.Printf("%+v\n", elliot)
}
