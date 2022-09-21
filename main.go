package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	country string
	city    string
	street  string
}

type Person struct {
	f_name   string
	l_name   string
	phone    string
	age      int
	is_alive bool
	address  Address
}

func init_function() Person {
	person := Person{
		f_name: "Haroon",
		l_name: "Ayaz",
		phone:  "+92-315-5992431",
		age:    22, is_alive: true,
		address: Address{country: "Pakistan", city: "Rawalpindi", street: "2"},
	}
	return person
}

func reciever_func(person Person) {
	fmt.Printf("Recieved something of type %s \n", reflect.TypeOf(person).String())
}

func main() {
	var person = init_function()
	fmt.Printf("Recieved a struct of type %s \n", reflect.TypeOf(person).String())
	fmt.Printf("person.f_name: %s \n", person.f_name)
	fmt.Printf("person.l_name: %s \n", person.l_name)
	fmt.Printf("Now passing this struct into another func\n")
	reciever_func(person)
}
