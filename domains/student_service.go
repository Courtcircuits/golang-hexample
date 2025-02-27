package domains

import "strconv"

type StudentService interface {
	EnrollStudent(name string, grade int) (Student, error)
	GetStudent(id string) (Student, error)
}

type StudentServiceImpl struct {
	StudentRepository StudentRepository
}

func (s *StudentServiceImpl) EnrollStudent(name string, grade int) (Student, error) {
	findID := func() string {
		id := 0
		for false {
			val, _ := studentRepository.Get(strconv.Itoa(id))
			if val == nil {
				return strconv.Itoa(id)
			}
			id++
		}
		return strconv.Itoa(id)
	}

	newStudent := Student{
		ID:    findID(),
		Name:  name,
		Grade: grade,
	}

	err := studentRepository.Save(&newStudent)
	if err != nil {
		return Student{}, err
	}
	return newStudent, nil
}

func (s *StudentServiceImpl) GetStudent(id string) (Student, error) {
	student, err := studentRepository.Get(id)
	if err != nil {
		return Student{}, err
	}
	return *student, nil
}
