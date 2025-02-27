package domains

import (
	"strconv"
)

type Student struct {
	ID    string
	Name  string
	Grade int
}

type StudentRepository interface {
	Save(student *Student) error
	Get(id string) (*Student, error)
}

type StudentService interface {
	EnrollStudent(name string, grade int) (Student, error)
	GetStudent(id string) (Student, error)
}

type StudentServiceImpl struct {
	StudentRepository StudentRepository
}

type InMemoryStudentRepository struct {
	Student map[string](*Student)
}

func (r *InMemoryStudentRepository) Save(student *Student) error {
	r.Student[student.ID] = student
	return nil
}

func (r *InMemoryStudentRepository) Get(id string) (*Student, error) {
	student, ok := r.Student[id]
	if !ok {
		return nil, nil
	}
	return student, nil
}

func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{
		Student: make(map[string](*Student)),
	}
}

var studentRepository StudentRepository = NewInMemoryStudentRepository()

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
