package main

import (
	"fmt"
	"go/qltv/library"
	"go/qltv/utils"
)

const menu = `
📚 CHUONG QUAN LY THU VIEN
1. Them sach
2. Xem danh sach sach
3. Them nguoi muon
4. Xem danh sach nguoi muon
5. Muon sach
6. Xem lich su muon
7. Tra sach
8. Tim kiem sach
9. Thoat`
func main() {
	for {
		utils.ClearScreen()
		fmt.Println(menu)
		 
		data:= utils.PositiveNumber("👉 Chọn chức năng : ")

		switch data{
		case 1:
			fmt.Println("===========THÊM SÁCH=============")
			if err := library.AddBook();err != nil{
				fmt.Printf(" Lỗi thêm sách: %v\n",err)
			}
			
		case 2:
			fmt.Println("==========DANH SÁCH SÁCH==============")
			if err := library.ListBook();err != nil{
				fmt.Printf(" Lỗi hiện danh sách sách: %v\n",err)
			}
			
		case 3:
			fmt.Println("============THEM NGƯỜI MƯỢN============")
			if err := library.AddBorrower();err != nil{
				fmt.Printf(" Lỗi thêm người mượn: %v\n",err)
			}
			
		case 4:
			fmt.Println("=============DANH SACH NGƯỜI MƯỢN===========")
			if err := library.ListBorrower();err != nil{
				fmt.Printf(" Lỗi hiện danh sách người mượn: %v\n",err)
			}
			
		case 5:
			fmt.Println("==========MƯỢN SÁCH==============")
			if err := library.BorrowBook();err != nil{
				fmt.Printf(" Lỗi mượn sách: %v\n",err)
			}
			
		case 6:
			fmt.Println("===========LỊCH SỬ MƯỢN SÁCH=============")
			if err := library.ListBorrowHistory();err != nil{
				fmt.Printf(" Lỗi hiện lịch sử mượn sách: %v\n",err)
			}
			
		case 7:
			fmt.Println("===========TRẢ SÁCH=============")
			if err := library.ReturnBook();err != nil{
				fmt.Printf(" Lỗi trả sach: %v\n",err)
			}
			
		case 8:
			fmt.Println("===========TÌM KIẾM SÁCH=============")
			if err := library.SearchBook();err != nil{
				fmt.Printf(" Lỗi tìm kiếm sách: %v\n",err)
			}
			
		case 9:
			fmt.Println("Tam biet!")
			return
		default:
			fmt.Println("Chọn tính năng không hợp lệ!!")
		}
		utils.ReadInput("\nNhan Enter De Tiep Tuc")
	}

}