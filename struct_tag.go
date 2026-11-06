//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// tag `type desciptionName` `key:giá trị`
// bỏ struct vô nhiều hàm khác nhau sẽ dựa vào cái tag để biết cách thực thi

// tag sinh ra để đổi tên Field thành name theo ý bạn in rra
type Employee struct {
	Name   string `json:"Họ Tên"`
	Age    int    `json:"Tuổi"`
	gender string
	//  gender G ko viết hoa sẽ ko export được
}

func main() {
	phuc := Employee{
		Name:   "Ưng Hoàng Phúc",
		Age:    36,
		gender: "male",
	}

	// chuyển về json sài packAge json.Marshal()
	outputByte, err := json.Marshal(phuc) // muốn export phải INHOA chử đầy tiên
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // thoát systerm
	}

	fmt.Println(string(outputByte))

}
