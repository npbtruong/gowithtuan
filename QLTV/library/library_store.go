package library

import (
	"fmt"
	"go/qltv/models"
	"time"
)

type Library struct {
	book        map[string]models.Book
	borrower    map[string]models.Borrower
	transaction map[string]models.Transaction
}

// construct
func NewLibrary() *Library {
	return &Library{
		book:        make(map[string]models.Book),
		borrower:    make(map[string]models.Borrower),
		transaction: make(map[string]models.Transaction),
	}
}

// khỏi tạo vùng nhớ 1 lần duy nhất ta dùng con trỏ đở  tốn ram dọn
func (lib *Library) AddBookStore(id, title, author string) error {

	if _, exist := lib.book[id]; exist {
		return fmt.Errorf("Sách với ID: %s đã tồn tại", id)
	}

	lib.book[id] = models.Book{
		Id:     id,
		Title:  title,
		Author: author,
	}

	return nil
}

// map
// func (lib *Library) ListBooksStore() map[string]models.Book {
// 	return lib.book
// }

// slide để có hổ trợ sort/search
func (lib *Library) ListBooksStore() []models.Book {

	books := make([]models.Book, 0, len(lib.book))
	for _, book := range lib.book {
		books = append(books, book)
	}

	return books
}

func (lib *Library) AddBorrowerStore(id, name, email string) error {

	if _, exist := lib.borrower[id]; exist {
		return fmt.Errorf("Người mượn với ID: %s đã tồn tại", id)
	}

	lib.borrower[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrower))
	for _, borrower := range lib.borrower {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func (lib *Library) BorrowBookStore(transactionId, bookId, borrowerId string) error {
	book, bookExist := lib.book[bookId]
	if !bookExist {
		return fmt.Errorf("sach với ID %s không tồn tại", bookId)
	}

	if book.IsBorrowed {
		return fmt.Errorf("sach với ID %s đã bị mượn", bookId)
	}

	if _, borrowerExist := lib.borrower[borrowerId]; !borrowerExist {
		return fmt.Errorf("Người mượn sách Id %s không tồn tại", borrowerId)
	}

	if _, transactionExist := lib.transaction[transactionId]; transactionExist {
		return fmt.Errorf("Transaction Id %s đã tồn tại", transactionId)
	}

	book.IsBorrowed = true
	lib.book[bookId] = book

	lib.transaction[transactionId] = models.Transaction{
		TransactionId: transactionId,
		BookId:        bookId,
		BorrowerId:    borrowerId,
		BorrowDate:    time.Now(),
	}

	return nil
}

func (lib *Library) ListTransactionStore() []models.Transaction {
	transactions := make([]models.Transaction, 0, len(lib.transaction))
	for _, transaction := range lib.transaction {
		transactions = append(transactions, transaction)
	}
	return transactions
}
