package main

import (
	"fmt"

	"github.com/Courtcircuits/students/domains"
)


func main() {
  student := domains.Student{
    ID: "001",
    Name: "Mihai",
    Grade: 10,
  }

  sr := domains.NewInMemoryStudentRepository()

  sr.Save(&student)
  s, _ := sr.Get(student.ID)

  fmt.Println(s.Name)

	//with the service

	service := domains.StudentServiceImpl{
		StudentRepository: sr,
	}

	student, err := service.EnrollStudent("Mihai", 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(student.Name)

	student, err = service.GetStudent(student.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(student.Name)
}
