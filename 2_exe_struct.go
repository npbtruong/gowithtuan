//go:build ignore

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//Bài 1 áp dung struct tính chu vi,diện tích HCN ,kiểm tra input

type Shape struct {
	Width  int
	Height int
}

type Output struct {
	Cv int `json:"chuvi"`
	Dt int `json:"dientich"`
}

func main() {
	var rectangle Shape
	rectangle.nhapThongTin()

	output := Output{
		Cv: rectangle.chuVi(),
		Dt: rectangle.dienTich(),
	}

	data, err := json.Marshal(output)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println(string(data))
}

// Dùng pointer *Shape chủ yếu tiết kiệm ô nhớ khỏi copy chứ Shape vẫn được
func (s *Shape) chuVi() int {
	return (s.Width + s.Height) * 2
}

func (s Shape) dienTich() int {
	return s.Width * s.Height
}

func (s *Shape) nhapThongTin() {

	for {

		fmt.Println("Nhập Width :")
		if err := ReadData(&s.Width); err != nil || s.Width <= 0 {
			fmt.Println("error w")
			continue
		}

		for {
			fmt.Println("Nhập Height :")
			if err1 := ReadData(&s.Height); err1 != nil || s.Height <= 0 {
				fmt.Println("error h")
				continue
			} else {
				return
			}
		}
	}

}

var reader = bufio.NewReader(os.Stdin)

// ReadData đọc nguyên 1 dòng từ bàn phím rồi chuyển vào data.
// data phải là pointer (vd: &s.Width) thì mới ghi được vào biến gốc.
func ReadData(data any) error {
	line, err := reader.ReadString('\n') // đọc hết dòng → không còn ký tự thừa
	if err != nil {                      // hết input (Ctrl+Z) → thoát, tránh lặp vô hạn
		fmt.Println("Không còn input")
		os.Exit(1)
	}
	_, err = fmt.Sscanln(line, data) // giống Scanln nhưng đọc từ string
	return err
}
