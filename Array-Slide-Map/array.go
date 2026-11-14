//go:build ignore

package main

import "fmt"

func main() {
	var numbers [5](float64)
	numbers[3] = 5.5

	var str = [6]string{"fafa", "fifi"}

	var list = [...]int{1, 2, 3, 4, 55}

	//Mattrix mãng đa chiều
	var matrix = [3][4]int{
		{999, 2, 3, 4},
		{2, 3, 1, 1},
		{1, 1, 1, 1},
	}
	fmt.Println(str, matrix[0][0])

	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-")
	for index, v := range list {
		fmt.Println(index, v)
	}

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-")
	for index, v := range matrix {
		fmt.Println(index, v)
		for _, val := range v {
			fmt.Println(val)
		}
	}
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-")
	DanhSachSv()

}

type Sinhvien struct {
	Mssv int
	Ten  string
	Tuoi int
}

func DanhSachSv() {
	list_sv := [...]Sinhvien{
		{Mssv: 123, Ten: "Tuấn", Tuoi: 12},
		{Mssv: 124, Ten: "Tú", Tuoi: 12},
		{Mssv: 125, Ten: "Nhi", Tuoi: 13},
		{Mssv: 126, Ten: "Tiên", Tuoi: 12},
	}

	for _, sv := range list_sv {
		fmt.Printf("%d Sinh viên %v Tuổi %d \n", sv.Mssv, sv.Ten, sv.Tuoi)
	}
}
