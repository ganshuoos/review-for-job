package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string
}

func (a Animal) A() {
	fmt.Println("A")
}

func (a Animal) B() {
	fmt.Println("B")
}

func main() {
	a := Animal{
		Name: "ganshuoos",
	}
	t := reflect.TypeOf(a)
	fmt.Println(t.Kind(), t.Name(), t.NumMethod())

	v := reflect.ValueOf(&a)
	v = v.Elem()
	v.Field(0).SetString("new Eggo")
	fmt.Println(a)
}
