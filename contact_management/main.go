package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact

var contactIndexByName map[string]int

var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func AddContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %+v\n", name)
		return
	}
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %+v\n", name)
}

func findContact(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func ListContacts() {
	fmt.Println("---Listing Contacts---")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	for idx, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", idx+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("")
}

func main() {
	AddContact("Nishak", "nishak@example.com", "01111111111")
	AddContact("Nishak", "nishak@example.com", "01111111111")
	AddContact("Jihan", "nishak@example.com", "01111111111")
	AddContact("Nishak JR", "nishak@example.com", "01111111111")
	AddContact("JR Nishak", "nishak@example.com", "01111111111")
	ListContacts()

	jihan := findContact("Jihan")
	if jihan == nil {
		fmt.Println("no contact found.")
	} else {
		fmt.Println("contact found.")
		fmt.Println(*jihan)
	}
}
