//go:build ignore
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Nhập số chọn hàm ex1:1 ex2:2")
	_, err := fmt.Scanln(&num)
	if err != nil || (num != 1 && num != 2) {
		fmt.Println("chọn số không có trong danh sách")
		return
	}

	switch num {
		case 1:ex1()
		case 2:ex2()
		default:return
	}
	
}

// Bài 01 xây hàm đệ quy tính tổng N
// người dùng nhập vào số N, tính tổng từ 1 -> N
func ex1 (){
	var num int
	fmt.Print("nhập số N để tính tổng 1->N: ")
	fmt.Scanln(&num)
	if(num > 0){
		fmt.Println("Tổng từ 1 ->", num, "là:", sum(num))
	} else {
		fmt.Println("Vui lòng nhập một số dương.")
	}
}

func sum(num int)int{
	if num >= 1 {
		num = num + sum(num-1)
	}
	return num
}

// Bài 02 xây dưng finonacci : bắt đầu bằng 0 1 số tiếp theo bằng tổng 2 số trước 0 1 2 3 5 7 12
// nhập số lượng hiện và hiện dãy fibonacci
// tính tổng dãy

type Fibo struct {
	numList string // chuỗi số fibo
	lenght int32 //độ dài fibo
	total int // tổng chuỗi
}

func ex2(){
	f:=Fibo{
		numList : "0,1",
		lenght : 2,
		total : 1,
	}

	fmt.Print("Nhập độ dài chuổi FIBO: ")
	fmt.Scanln(&f.lenght)

	if f.lenght < 2 {
		fmt.Println("Dãy Fibonacci tối thiểu có 2 số") 
		return
	}else {
		f.lenght = f.lenght - 2
		if(f.lenght > 0){
			f.Fibonacii(0,1)
		}
	}

	
	fmt.Printf("list Fibonacci: %s \n",f.numList)
	fmt.Printf("tổng các số Fibonacci là %d \n",f.total)

}

func (f *Fibo) Fibonacii(num1 int,num2 int)(int){
	num3 :=num1+num2
	f.lenght--
	f.total += num3
	f.numList += fmt.Sprintf(",%d",num3)
	if(f.lenght > 0){
		return f.Fibonacii(num2,num3)
	}
	return num3
}


//Bài 3 Xây tập hợp chức năng
//chuyển bài tập thành Function
//xây dựng menu chọn hàm