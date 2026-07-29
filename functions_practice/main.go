package main

import "fmt"

func main() {
	fmt.Println("=== Different Types of Functions in Go ===")
	fmt.Println()

	// 1. Basic function (no params, no return)
	fmt.Println("--- 1. Basic function ---")
	sayHello()

	// 2. Function with parameters
	fmt.Println("\n--- 2. Function with parameters ---")
	greet("Nishak")

	// 3. Function with return value
	fmt.Println("\n--- 3. Function with return value ---")
	sum := add(10, 20)
	fmt.Println("10 + 20 =", sum)

	// 4. Multiple return values
	fmt.Println("\n--- 4. Multiple return values ---")
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", quotient, remainder)

	// 5. Named return values
	fmt.Println("\n--- 5. Named return values ---")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle 5x3 -> area: %d, perimeter: %d\n", area, perimeter)

	// 6. Variadic function (accepts any number of args)
	fmt.Println("\n--- 6. Variadic function ---")
	fmt.Println("Sum of 1,2,3,4,5 =", sumAll(1, 2, 3, 4, 5))
	fmt.Println("Sum of 10,20 =", sumAll(10, 20))

	// 7. Anonymous function
	fmt.Println("\n--- 7. Anonymous function ---")
	func(name string) {
		fmt.Println("Hello from anonymous function,", name)
	}("Nishak")

	// 8. Function as a variable (first-class function)
	fmt.Println("\n--- 8. Function as a variable ---")
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("6 * 7 =", multiply(6, 7))

	// 9. Function as a parameter (higher-order function)
	fmt.Println("\n--- 9. Function as a parameter ---")
	applyOperation(8, 4, add)
	applyOperation(8, 4, subtract)

	// 10. Function that returns a function (closure)
	fmt.Println("\n--- 10. Closure (function returns a function) ---")
	counter := makeCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3

	// 11. Recursive function
	fmt.Println("\n--- 11. Recursive function ---")
	fmt.Println("Factorial of 5 =", factorial(5))

	// 12. Deferred function (runs when surrounding function returns)
	fmt.Println("\n--- 12. Deferred function ---")
	deferExample()

	// 13. Methods (functions with receivers)
	fmt.Println("\n--- 13. Methods (value & pointer receivers) ---")
	p := Person{Name: "Nishak", Age: 23}
	p.Introduce()       // value receiver
	p.HaveBirthday()    // pointer receiver
	fmt.Println("Age after birthday:", p.Age)

	// 14. Blank identifier — ignore a return value
	fmt.Println("\n--- 14. Ignoring a return value ---")
	q, _ := divide(20, 3)
	fmt.Println("20 / 3 quotient only =", q)
}

// ---------- 1. Basic function ----------
func sayHello() {
	fmt.Println("Hello, Go functions!")
}

// ---------- 2. Parameters ----------
func greet(name string) {
	fmt.Println("Hello,", name)
}

// ---------- 3. Single return value ----------
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// ---------- 4. Multiple return values ----------
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// ---------- 5. Named return values ----------
// Named returns are declared in the signature; bare return uses them.
func rectangle(length, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // bare return
}

// ---------- 6. Variadic function ----------
// ...int means "zero or more int arguments"
func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// ---------- 9. Function as parameter ----------
// op is a function type: takes two ints, returns one int
func applyOperation(a, b int, op func(int, int) int) {
	result := op(a, b)
	fmt.Printf("Result of operation on %d and %d = %d\n", a, b, result)
}

// ---------- 10. Closure ----------
// makeCounter returns a function that remembers "count"
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// ---------- 11. Recursion ----------
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// ---------- 12. defer ----------
func deferExample() {
	fmt.Println("  Start of deferExample")
	defer fmt.Println("  Deferred: runs last (cleanup)")
	fmt.Println("  Middle of deferExample")
	fmt.Println("  End of deferExample")
	// deferred line prints after this function returns
}

// ---------- 13. Methods ----------
type Person struct {
	Name string
	Age  int
}

// Value receiver — works on a copy (fine for reading)
func (p Person) Introduce() {
	fmt.Printf("Hi, I am %s and I am %d years old.\n", p.Name, p.Age)
}

// Pointer receiver — can modify the original
func (p *Person) HaveBirthday() {
	p.Age++
}
