package main

type Interface interface {
	String() string
}

type Student struct {
	name string
	t Teacher
}

type Teacher struct {
	name string
	sex int
}

func (s *Student) GetName() string{
	return s.name
}

func (s *Student) String() string {
	return "[ Alice ]"
}

func main() {

}

