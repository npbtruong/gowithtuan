package main

import (
	"fmt"

	"go/quanli/input"
	"go/quanli/student"
)

const menu = `
📚 CHUONG TRINH QUAN LY
1. Quan ly sinh vien
2. Quan ly giang vien
3. Thoat`

func main() {
	for {
		fmt.Println(menu)
		switch input.Option("👉 Chon chuc nang: ", 1, 3) {
		case 1:
			student.Menu()
		case 2:
			fmt.Println("Chuc nang quan ly giang vien dang phat trien.")
		case 3:
			fmt.Println("Tam biet!")
			return
		}
	}
}
