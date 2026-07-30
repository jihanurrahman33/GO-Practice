package embeddingex

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

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name string
	Address
	ContactInfo
	BussinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)
	fmt.Printf("Location: %s\n", c.FullAdress())
	fmt.Printf("Street (promoted): %s\n", c.Street)
	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("Bussiness Type: %s\n", c.BussinessType)
}

type CompanyWithOwnEmail struct {
	Name string
	Address
	ContactInfo
	Email string
}

func PrintEmbeddingInfo() {

	fmt.Println("----- Struct Embedding -----")

	comp := Company{
		Name: "Innovate Solutions Inc.",
		Address: Address{
			Street:  "123 Main St",
			City:    "Springfield",
			State:   "IL",
			ZipCode: "62701",
		},
		ContactInfo: ContactInfo{
			Email: "conact@example.com",
			Phone: "810-312-0000",
		},
		BussinessType: "Technology",
	}

	comp.GetProfile()

	fmt.Printf("\nDirect Access to comp.City%s\n", comp.City)
	fmt.Printf("\nDirect Access to comp.Phone%s\n", comp.Phone)

	fmt.Printf("\nEmbedded Address struct%+v\n", comp.Address)
	fmt.Printf("\nEmbedded ContactInfo struct%+v\n", comp.ContactInfo)
}
