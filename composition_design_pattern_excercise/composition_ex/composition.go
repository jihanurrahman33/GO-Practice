package compositionex

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAdress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}
	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerID      int
	Name            string
	Email           string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer ID: %d\n", c.CustomerID)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("Billing Address: %s\n", c.BillingAddress.FullAdress())
	fmt.Printf("Shipping Address: %s\n", c.ShippingAddress.FullAdress())
}

func PrintCompositionEx() {
	fmt.Println("----- Composition -----")

	cust := Customer{
		CustomerID: 101,
		Name:       "Jane Doe",
		Email:      "jane.doe@example.com",
		BillingAddress: Address{
			Street:  "123 Main St",
			City:    "Springfield",
			State:   "IL",
			ZipCode: "62701",
		},
		ShippingAddress: Address{
			Street:  "456 Market St",
			City:    "Springfield",
			State:   "IL",
			ZipCode: "62702",
		},
	}

	cust.PrintDetails()

	fmt.Println("----- customer with same billing and shipping adress -----")

	mainAdress := Address{
		Street:  "123 Main St",
		City:    "Springfield",
		State:   "IL",
		ZipCode: "62701",
	}

	cust2 := Customer{
		CustomerID:      101,
		Name:            "Jane Doe",
		Email:           "jane.doe@example.com",
		BillingAddress:  mainAdress,
		ShippingAddress: mainAdress,
	}

	cust2.PrintDetails()
}
