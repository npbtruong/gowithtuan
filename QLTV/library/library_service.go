package library

import (
	"fmt"
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

// khỏi tạo vùng nhớ 1 lần duy nhất ta dùng con trỏ đở  tốn ram dọn
func AddBook(lib *Library) error {
	id := utils.GenerateUuid()
	title := utils.GetNonEmptyString("Nhập tên sách: ")
	author := utils.GetNonEmptyString("Nhập tên tác giả: ")

	if err := lib.AddBookStore(id, title, author); err != nil{
		return err
	}

	fmt.Printf("Thêm sách thành công ! Id: %s", id)
	fmt.Println(lib.book)

	return nil
}

func ListBook() error {
	return nil
}

func AddBorrower() error {
	return nil
}

func ListBorrower() error {
	return nil
}

func BorrowBook() error {
	return nil
}

func ListBorrowHistory() error {
	return nil
}

func ReturnBook() error {
	return nil
}

func SearchBook() error {
	return nil
}
