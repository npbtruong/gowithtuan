//go:build ignore
package main

import "fmt"

func main() {
	//Bài 1 hiện số 1 -> 100
	//Bỏ các con số ngẫu nhiên 6, 48, 75, 89
	// mỗi số cách nhau = dấu , 
	// bỏ dấu phây cuối

	for i:=1;i<=100;i++{
		switch i {
			case 6,48,75,89:
				continue
			default :
				if(i==100){
					fmt.Printf("%d",i)
					break
				} 
				fmt.Printf("%d,",i)
		}

	}

	fmt.Printf("\n\n")
	//optimal 1
	var s string
	for i:=1;i<=100;i++{
		switch i {
			case 6,48,75,89:
				continue
			default :
				if(i==100){
				s += fmt.Sprintf("%d",i)
					break
				} 
				s += fmt.Sprintf("%d,",i)
		}

	}
	fmt.Println(s)


	fmt.Printf("\n\n")
	// Bài 2 hiện 1 -> 100
	// chỉ số lẻ
	// cứ 3 số xuống dòng ko có dấy phẩy cuối
	// mỗi số cách nhau bằng ,
	// bỏ dấu , ở vị trí cuối
    j:=3
	for i:=1;i<=100;i++{
		if(i%2!=0){
			if(j>0 && j<3){
				fmt.Print(",")
			}
			j--
			if(j==0){
				fmt.Printf("%d\n",i)
				j=3
			}else{
				fmt.Printf("%d",i)
			}
			
			
		}
	}


	fmt.Printf("\n\n")
	//Bài 3 bảng cửu chương
	// cho phép nhập số bắt đầu kết thúc
	// kiểm tra input
	// mỗi bảng cửu chương có title

	var from int
	var to int

	fmt.Println("Nhập số nguyên dương bắt đầu: ")
	fmt.Scanln(&from)
	fmt.Println("Nhập số nguyên dương kết thúc: ")
	fmt.Scanln(&to)

	if(from > 1 && to > 1 && from < to && from <= 9 && to <= 9){
		for i:=from;i<=to;i++{
			fmt.Printf("Bảng cửu chương %d\n",i)
			for j:=1;j<=9;j++{
				fmt.Printf("%d * %d = %d\n",i,j,i*j)
			}
		}
	}else{
		fmt.Println("Nhập sai, vui lòng nhập lại")
		return
	}
}