package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Student struct {
	roll_number int
	name        string
	address     string
	subjects    []string
}

type StudentList struct {
	list []*Student
}

func new_student(roll_number int, name string, address string, subjects []string) *Student {
	s := new(Student)
	s.roll_number = roll_number
	s.name = name
	s.address = address
	s.subjects = subjects
	return s
}

func (ls *StudentList) create_student(roll_number int, name string, address string, subjects []string) *Student {
	st := new_student(roll_number, name, address, subjects)
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
		fmt.Println("Student Subjects: ", student_list[i].subjects)
	}
}

func CalculateHash(string_to_hash string) string {
	fmt.Printf("String Received: %s\n", string_to_hash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(string_to_hash)))
}

func main() {
	students := new(StudentList)
	students.create_student(21, "Tony Stark", "10-8-80 Malibu Point 90625", []string{"Physics", "QuantumMechanics"})
	students.create_student(22, "Son Goku-Kun", "Planet Earth", []string{"basic maths: failed", "elementary physics: failed"})
	students.create_student(21, "Rycon Sama", "Planet Xivor", []string{"ehh", "dont care"})
	students.create_student(22, "Gogeta", "The Name Says it all :)", []string{":^)", "v_v"})
	fmt.Println("Now getting Hash of a particular string `ThroughOutHeavenAndEarthIAloneAmHonoredOne`: ", CalculateHash("ThroughOutHeavenAndEarthIAloneAmHonoredOne"))

	print_list(students)
}
