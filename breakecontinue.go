//go:build ignore

package main

import "fmt"

func main() {
	
	loop:for {
		fmt.Println("================================")
		fmt.Println("Nhập số chọn mode, 0 là kết thúc")
		var mod int
		
		_,err:=fmt.Scan(&mod)
		if(err != nil){
			fmt.Println("Error Unknow 500")
			continue
		}

		switch mod{
			case 1:
				var number int
				for{
					//var err error
					fmt.Println("vui lòng nhập số dương")
					_,err:=fmt.Scan(&number)

					if(err != nil || number<=0){
						fmt.Println("Nhập sai")
						continue // méo cần continue cũng được nhưng đẻ cho tường minh
					}else{
						break
					}
				}

				fmt.Printf("case 1: %d \n",number)
				//continue go méo cần continue
			case 2:
				fmt.Println("case 2")
				//continue go méo cần continue
			case 0:
				fmt.Println("Kết con nhà thúc")
				break loop
			default:
				fmt.Println("default")
		}
	}
	fmt.Println("Tác giả : UrGOD")
}
