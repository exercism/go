package gradeschool

type School any

func New() *School {
	panic("Please implement the New function")
}

func (s *School) Add(student string, grade int) bool {
	panic("Please implement the Add function")
}

func (s *School) Grade(level int) []string {
	panic("Please implement the Grade function")
}

func (s *School) Enrollment() []string {
	panic("Please implement the Enrollment function")
}
