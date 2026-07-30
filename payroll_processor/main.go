package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommisionEmployee struct {
	Name          string
	BaseSalary    float64
	CommisionRate float64
	SalasAmount   float64
}

func (ce CommisionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommisionRate * ce.SalasAmount)
}

func (ce CommisionEmployee) String() string {
	return fmt.Sprintf("Commision: %s (Base: $%.2f/hr, CommRate: %.2f%%). Sales: $%.2f", ce.Name, ce.BaseSalary, ce.CommisionRate*100, ce.SalasAmount)
}

func printEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("- Processing: %s\n", employee)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("\n--- Processing Payroll ---")
	totalPayroll := 0.0

	for _, emp := range employees {
		printEmployeeSummary(emp)
		pay := emp.CalculatePay()
		fmt.Printf("Monthly Pay: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nMonthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("------------------------------")
}
func main() {
	fmt.Println("Welcome to the Payroll processor!")

	salEmp := SalariedEmployee{
		Name:         "Md Jihanur Rahman",
		AnnualSalary: 720000.00,
	}
	hrEmp := HourlyEmployee{
		Name:        "Nishak",
		HourlyRate:  25.00,
		HoursWorked: 160.00,
	}
	comEmp := CommisionEmployee{
		Name:          "Nishak JR",
		BaseSalary:    2000.00,
		CommisionRate: 0.10,
		SalasAmount:   15000.00,
	}

	payRollList := []Payable{
		salEmp,
		hrEmp,
		comEmp,
		HourlyEmployee{Name: "JR Nishak", HourlyRate: 30.00, HoursWorked: 150.00},
	}
	ProcessPayroll(payRollList)
}
