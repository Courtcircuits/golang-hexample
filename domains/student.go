package domains

type Student struct {
	ID    string
	Name  string
	Grade int
}

type StudentRepository interface {
	Save(student *Student) error
	Get(id string) (*Student, error)
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
