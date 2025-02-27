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
}
