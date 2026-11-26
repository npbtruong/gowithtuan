package library

import (
	"fmt"
	"go/qltv/models"
	"time"
)

type Library struct {
	books        map[string]models.Book
	borrowers    map[string]models.Borrower
	transactions map[string]models.Transaction
}

// construct
func NewLibrary() *Library {
	return &Library{
		books:        make(map[string]models.Book),
		borrowers:    make(map[string]models.Borrower),
		transactions: make(map[string]models.Transaction),
	}
}

// khỏi tạo vùng nhớ 1 lần duy nhất ta dùng con trỏ đở  tốn ram dọn
func (lib *Library) AddBookStore(id, title, author string) error {

	if _, exist := lib.books[id]; exist {
		return fmt.Errorf("Sách với ID: %s đã tồn tại", id)
	}

	lib.books[id] = models.Book{
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

	books := make([]models.Book, 0, len(lib.books))
	for _, book := range lib.books {
		books = append(books, book)
	}

	return books
}

func (lib *Library) AddBorrowerStore(id, name, email string) error {

	if _, exist := lib.borrowers[id]; exist {
		return fmt.Errorf("Người mượn với ID: %s đã tồn tại", id)
	}

	lib.borrowers[id] = models.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}

	return nil
}

func (lib *Library) ListBorrowersStore() []models.Borrower {
	borrowers := make([]models.Borrower, 0, len(lib.borrowers))
	for _, borrower := range lib.borrowers {
		borrowers = append(borrowers, borrower)
	}
	return borrowers
}

func (lib *Library) BorrowBookStore(transactionId, bookId, borrowerId string) error {
	book, bookExist := lib.books[bookId]
	if !bookExist {
		return fmt.Errorf("sach với ID %s không tồn tại", bookId)
	}

	if book.IsBorrowed {
		return fmt.Errorf("sach với ID %s đã bị mượn", bookId)
	}

	if _, borrowerExist := lib.borrowers[borrowerId]; !borrowerExist {
		return fmt.Errorf("Người mượn sách Id %s không tồn tại", borrowerId)
	}

	if _, transactionExist := lib.transactions[transactionId]; transactionExist {
		return fmt.Errorf("Transaction Id %s đã tồn tại", transactionId)
	}

	book.IsBorrowed = true
	lib.books[bookId] = book

	lib.transactions[transactionId] = models.Transaction{
		TransactionId: transactionId,
		BookId:        bookId,
		BorrowerId:    borrowerId,
		BorrowDate:    time.Now(),
	}

	return nil
}

func (lib *Library) ListTransactionByBorrowerStore(borrowerId string) []models.Transaction {

	if _, borrowerExist := lib.borrowers[borrowerId]; !borrowerExist {
		fmt.Printf("Người mượn sách Id %s không tồn tại", borrowerId)
		return nil
	}

	transactions := make([]models.Transaction, 0, len(lib.transactions))

	for _, transaction := range lib.transactions {
		if(transaction.BorrowerId == borrowerId){
			transactions = append(transactions, transaction)
		}
	}

	return transactions
}
