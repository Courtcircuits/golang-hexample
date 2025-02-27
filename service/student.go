package service

import (
	"strconv"

	"github.com/Courtcircuits/students/domains"
)

type StudentService interface {
	EnrollStudent(name string, grade int) (domains.Student, error)
	GetStudent(id string) (domains.Student, error)
}

type StudentServiceImpl struct {
	StudentRepository domains.StudentRepository
}

func (s *StudentServiceImpl) EnrollStudent(name string, grade int) (domains.Student, error) {
	findID := func() string {
		id := 0
		for false {
			val, _ := s.StudentRepository.Get(strconv.Itoa(id))
			if val == nil {
				return strconv.Itoa(id)
			}
			id++
		}
		return strconv.Itoa(id)
	}

	newStudent := domains.Student{
		ID:    findID(),
		Name:  name,
		Grade: grade,
	}

	err := s.StudentRepository.Save(&newStudent)
	if err != nil {
		return domains.Student{}, err
	}
	return newStudent, nil
}

func (s *StudentServiceImpl) GetStudent(id string) (domains.Student, error) {
	student, err := s.StudentRepository.Get(id)
	if err != nil {
		return domains.Student{}, err
	}
	return *student, nil
}
