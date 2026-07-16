package rectangle

import "fmt"

func CallProblem2() {
	rct1, rct2 := registerRectangle()

	areaOfRCT1 := rct1.calculateArea()
	areaOfRCT2 := rct2.calculateArea()
	perimeterOfRCT1 := rct1.calculatePerimeter()
	perimeterOfRCT2 := rct2.calculatePerimeter()

	fmt.Println("Area of Rectangle 1 is:", areaOfRCT1)
	fmt.Println("Area of Rectangle 2 is:", areaOfRCT2)
	fmt.Println("Perimeter of Rectangle 1 is:", perimeterOfRCT1)
	fmt.Println("Perimeter of Rectangle 2 is:", perimeterOfRCT2)
}
