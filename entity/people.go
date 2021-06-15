package entity

import "fmt"

type People struct {
	Id      string
	Name    string
	Age     int
	Company string
	Address string
}

func (p People) SayHello() {
	fmt.Printf("Hello, I'm %s, %d years old\n", p.Name, p.Age)
}

func (p *People) UpdateAddress(address string) {
	p.Address = address
}

func (p *People) UpdateCompany(company string) {
	p.Company = company
}

func (p People) ToString() string {
	return fmt.Sprintf("Id: %s, Name: %s, Age: %d, Company: %s, Address: %s", p.Id, p.Name, p.Age, p.Company, p.Address)
}
