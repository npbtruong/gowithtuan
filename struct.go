//go:build ignore

package main

import "fmt"

type GiangVien struct {
	name   string
	gender string
	age    int
}

// func ("Receiver") "Method"(){}
func (gv GiangVien)letterInfo(){
	fmt.Printf("Thư gửi %v %v %v  \n", gv.name, gv.age, gv.gender)
	gv.age++
}

func (gv *GiangVien)clearInfo(){
	gv.name = ""
	gv.gender = ""
	gv.age = 0
}

func main() {
	LHP := GiangVien{
		name:   "Le Hong Phong",
		gender: "male",
		age:    34,
	}

	ADC := GiangVien{
		name: "Andrade Cien Almas",
		gender: "LK",
		age: 37,
	}

	hienThongTin(&LHP)
	hienThongTin(&ADC)

	fmt.Println()
	
	LHP.letterInfo()

	LHP.clearInfo()

	fmt.Println(LHP)
}

// Tại sao Dùng con trỏ để đở copy vô ô nhớ mới đở tốn dung lượng
func hienThongTin(gv *GiangVien){
	//khi sài truct có thể lược bỏ (*gv) thành gv.NẾU là biến bắt buộc có * truct bỏ được
	fmt.Printf("Name:%v Age:%v Gender:%v \n", (*gv).name, gv.age, gv.gender)
}