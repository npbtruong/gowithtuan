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

	if err := lib.AddBookStore(id, title, author); err != nil {
		return err
	}

	fmt.Printf("Thêm sách thành công ! Id: %s \n", id)

	return nil
}

func ListBook(lib *Library) error {
	books := lib.ListBooksStore()
	if len(books) == 0 {
		fmt.Println("Thư viện chưa có sách , vui lòng thêm sách!!")
		return nil
	}

	for _, book := range books {
		status := "còn"
		if book.IsBorrowed {
			status = "Bị mượn"
		}
		fmt.Printf("Id:%s , Tác giả:%s , Tiêu đề:%s , Trạng Thái:%s \n", book.Id, book.Author, book.Title, status)
	}
	return nil
}

func AddBorrower(lib *Library) error {

	id := utils.GenerateUuid()
	name := utils.GetNonEmptyString("Nhập tên người mượn: ")
	email := utils.GetNonEmptyString("Nhập email người mượn: ")

	if err := lib.AddBorrowerStore(id, name, email); err != nil {
		return err
	}

	fmt.Printf("Thêm người mượn thành công ! Id: %s \n", id)

	return nil
}

func ListBorrower(lib *Library) error {
	borrowers := lib.ListBorrowersStore()
	if len(borrowers) == 0 {
		fmt.Println("Chưa có người mượn sách , vui lòng thêm người mượn!!")
		return nil
	}

	for _, borrower := range borrowers {
		fmt.Printf("Id:%s , Tên:%s , Email:%s \n", borrower.Id, borrower.Name, borrower.Email)
	}
	return nil
}

func BorrowBook(lib *Library) error {

	transactionId := utils.GenerateUuid()
	bookId := utils.GetNonEmptyString("Nhập Id sách mượn: ")
	borrowerId := utils.GetNonEmptyString("Nhập Id người mượn: ")

	if err := lib.BorrowBookStore(transactionId, bookId, borrowerId); err != nil {
		return err
	}

	fmt.Printf("Sách được mượn thành công ! Transaction Id: %s \n", transactionId)

	return nil

}

func ListBorrowHistory(lib *Library) error {

	transactions := lib.ListTransactionStore()
	for _, transaction := range transactions {
		fmt.Printf("Id Giao Dịch: %s ; Id sách %s ; Id người mượn %s ; ngày mượn %v ; ngày trả %v",
			transaction.TransactionId,
			transaction.BookId,
			transaction.BorrowerId,
			transaction.BorrowDate,
			transaction.ReturnDate)
	}

	return nil
}

func ReturnBook() error {
	return nil
}

func SearchBook() error {
	return nil
}
