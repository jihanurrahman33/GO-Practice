package main

import (
	"fmt"
	"inventory_system/inventory"
)

func main() {
	inv := inventory.Inventory{}

	inv.AddItem("Mouse", "Keyboard", "Monitor")
	inv.DisplayItem()

	fmt.Println("------")

	inv.RemoveItem("Keyboard")
	inv.DisplayItem()
}
