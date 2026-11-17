package employee

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float64
	Phone  int
}

func (e Employee) GetId() int {
	return e.Id
}

func (e Employee) GetInfo() {
	fmt.Printf("- ID: %d | Ten: %s | Luong: %.2f | sdt: %d \n", e.Id, e.Name, e.Salary, e.Phone)
}
