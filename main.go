package main

import "fmt"

func main() {

	data:="continue"

	var num int

	process(data,num)
}

func process(data string,num int) {
	// golang luôn phải truyền đủ các tham số vào func, không có giá trị mặc định
	// golang không cho khai giá trị mặc định trên func muốn mặc định if bên dưới
	if(num == 0){
		num = 1
	}


	switch data {
	case "break","something":
		for i:=0;i<10;i+=2{
			if (i==4) {
				break
			}
			fmt.Println("Hello")
		}
	case "continue":
		for i:=0;i<10;i+=2{
			if (i==4) {
				continue
			}
			fmt.Println("Hello")
		}
	default:
		fmt.Println("Unknown greeting")
	}
}