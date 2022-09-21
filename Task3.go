package main

import (
	"fmt"
	"strings"
)

type Student struct {
	roll_number int
	name        string
	address     string
}

func new_student(roll_number int, name string, address string) *Student {
	s := new(Student)
	s.roll_number = roll_number
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) create_student(roll_number int, name string, address string) *Student {
	st := new_student(roll_number, name, address)
	ls.list = append(ls.list, st)
	return st
}

func print_list(st_ls *StudentList) {
	student_list := st_ls.list
	for i := 0; i < len(st_ls.list); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 10), i+1, strings.Repeat("=", 10))
		fmt.Println("Student Roll Number: ", student_list[i].roll_number)
		fmt.Println("Student Name: ", student_list[i].name)
		fmt.Println("Student Address: ", student_list[i].address)
	}
}

func main() {
	students := new(StudentList)
	students.create_student(21, "Tony Stark", "10-8-80 Malibu Point 90625")
	students.create_student(22, "Son Goku-Kun", "Planet Earth")
	students.create_student(21, "Rycon Sama", "Planet Xivor")
	students.create_student(22, "Gogeta", "The Name Says it all :)")
	print_list(students)
}
