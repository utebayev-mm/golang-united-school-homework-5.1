package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	var end Point
	end.x += s.start.x
	end.y += s.start.y
	return end
}

func (s *Square) Area() uint {
	square := s.a * s.a
	return square
}

func (s *Square) Perimeter() uint {
	perimeter := s.a * 4
	return perimeter
}
