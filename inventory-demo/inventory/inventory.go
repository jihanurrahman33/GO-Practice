package inventory

import "fmt"

type Inventory struct {
	Items []string
}

func (inv *Inventory) AddItem(items ...string) {
	if len(items) == 0 {
		fmt.Println("Please add a item")
		return
	}

	inv.Items = append(inv.Items, items...)

}
func (inv *Inventory) RemoveItem(item string) {
	if len(item) == 0 {
		fmt.Println("please give an item to remove")
		return
	}
	rmvItemIdx := 0
	for idx, val := range inv.Items {
		if item == val {
			rmvItemIdx = idx
		}
	}
	inv.Items = append(inv.Items[:rmvItemIdx], inv.Items[rmvItemIdx+1:]...)
}
func (inv *Inventory) DisplayItem() {
	for idx, val := range inv.Items {
		fmt.Printf("%d. %s\n", idx+1, val)
	}
}
