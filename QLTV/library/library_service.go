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

	borrowerId := utils.GetNonEmptyString("Nhập Id người mượn: ")

	transactions := lib.ListTransactionByBorrowerStore(borrowerId)

	if len(transactions) == 0 {
		fmt.Println("Không có lịch sử mượn sách với borrowerId này!")
		return nil
	}

	returnDate := "chưa trả"
	for _, transaction := range transactions {
		if !transaction.ReturnDate.IsZero() {
			returnDate = transaction.ReturnDate.Format("01-01-2001")
		}
		fmt.Printf("Giao Dịch: %s ;Sách: %s ;Ngày mượn: %s ;Ngày trả: %s",
			transaction.TransactionId,
			transaction.BookId,
			transaction.BorrowDate.Format("01-01-2001"),
			returnDate)
	}

	return nil
}

func ReturnBook(lib *Library) error {
	transactionId := utils.GetNonEmptyString("Nhập Id Giao Dịch: ")

	if err := lib.ReturnBookStore(transactionId); err != nil {
		return err
	}

	fmt.Printf("Trả sách thành công thành công ! Id Giao dịch: %s \n", transactionId)

	return nil
}

func SearchBook(lib *Library) error {
	query := utils.GetNonEmptyString("Nhập từ khóa tìm kiếm sách : ")

	books := lib.SearchBooksStore(query)
	if len(books) == 0 {
		fmt.Println("Không tìm ra sách !!")
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
