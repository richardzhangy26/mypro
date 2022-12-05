package model

type student struct {
	Name  string
	score float32
}

func NewStudent(n string, s float32) *student {
	return &student{
		Name:  n,
		score: s,
	}
}
func (stu *student) GetScore() float32 {
	return stu.score
}
