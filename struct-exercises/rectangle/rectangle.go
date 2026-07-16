package rectangle

type rectangle struct {
	width  float32
	height float32
}

func (r rectangle) calculateArea() float32 {
	return r.width * r.height

}
func (r rectangle) calculatePerimeter() float32 {
	return 2 * (r.width + r.height)
}
func registerRectangle() (rectangle, rectangle) {
	rct1 := rectangle{
		width:  20,
		height: 20,
	}
	rct2 := rectangle{
		width:  40,
		height: 40,
	}

	return rct1, rct2
}
