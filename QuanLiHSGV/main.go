package main

import (
	"fmt"
	"go/quanli/employee"
	"go/quanli/input"
	"go/quanli/student"
	"go/quanli/teacher"
)

const menu = `
📚 CHUONG TRINH QUAN LY
1. Quan ly sinh vien
2. Quan ly giang vien
3. Quan ly nhan vien
4. Thoat`

func main() {
	for {
		fmt.Println(menu)
		switch input.Option("👉 Chon chuc nang: ", 1, 4) {
		case 1:
			student.Menu()
		case 2:
			teacher.Menu()
		case 3:
			employee.Menu()
		case 4:
			fmt.Println("Tam biet!")
			return
		}
	}
}
