package utility

type Quadrato struct {
	X int
	Y int
}

type FormaGeometrica interface {
	Area() int
}

func (q Quadrato) Area() int {
	return q.X * q.Y
}