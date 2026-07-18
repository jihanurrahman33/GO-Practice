package main

import (
	"fmt"
)

func main() {

	fmt.Println("Day 1 of practicing maps in golang")

	Nishak := map[string]any{
		"Name":   "Nishak",
		"Age":    23,
		"Gender": "Male",
	}
	Nishak["Age"] = "24"
	fmt.Println(Nishak["Name"])
	fmt.Println(Nishak["Age"])
	fmt.Println(Nishak["Gender"])
	Nishak["Age"] = "23"
	for key, value := range Nishak {
		fmt.Printf("%s:%s", key, value)
		fmt.Println()
	}
	changeMapValue(&Nishak, "Age", 26)

	fmt.Println(Nishak)

	newemtMap := createEmptyMap()
	fmt.Println(*newemtMap)
	(*newemtMap)["Name"] = "Newly Created Empty Map"
	fmt.Println(*newemtMap)
	assignedMap := *newemtMap
	fmt.Println(assignedMap)

}

func createEmptyMap() *map[string]any {
	return &map[string]any{}
}
func changeMapValue(mp *map[string]any, givenKey string, newValue any) {
	(*mp)[givenKey] = newValue
	fmt.Println("Updated")
}
