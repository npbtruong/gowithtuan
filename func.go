//go:build ignore

package main

import "fmt"

func main() {
	// sum, sub, div := multiVal(10, 4)
	// println("Sum:", sum)
	// println("Sub:", sub)
	// println("Div:", div)

	var num int
	fmt.Print("nhập số: ")
	fmt.Scanln(&num)
	if(num > 0){
		countDown(num)
	}

}

func multiVal(num1, num2 int) (int, int, float64) {
	sum := num1 + num2
	sub := num1 - num2
	div := float64(num1) / float64(num2)
	return sum, sub, div
}

// ĐỆ QUY nhập số gọi đến khi số nhỏ hơn -20 thì dừng
func countDown(num int) {
	fmt.Println(num)
	if num > -20 {
		countDown(num - 1)
	}
}
