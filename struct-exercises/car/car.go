package car

import "fmt"

type car struct {
	brand string
	model string
	year  int
	speed int
}

func (c *car) Accelerate(amount int) {
	fmt.Println("Accelarated Speed:", c.speed+amount)
}
func (c *car) Brake(amount int) {
	c.speed -= amount

	fmt.Println("Speed:", c.speed)

}
func (c *car) DisplayInfo() {
	fmt.Println("Car Brand:", c.brand)
	fmt.Println("Car Model:", c.model)
	fmt.Println("Car Year:", c.year)
	fmt.Println("Car speed:", c.speed)

}

func RegisterCar() {
	car1 := car{
		brand: "Toyota",
		model: "f1",
		year:  2019,
		speed: 200,
	}

	car1.Accelerate(100)
	car1.Brake(50)
	car1.DisplayInfo()
}
