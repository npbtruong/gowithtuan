package library

import (
	"fmt"
	"go/qltv/models"
)

type Library struct {
	book map[string]models.Book
}

// construct
func NewLibrary() *Library {
	return &Library{
		book: make(map[string]models.Book),
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
