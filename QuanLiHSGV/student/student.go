// Package student quản lý danh sách sinh viên: thêm, xoá, sửa, xem, tìm kiếm.
//
// Bên ngoài chỉ cần gọi student.Menu(). Các hàm và biến viết thường
// (add, edit, students...) là private, chỉ dùng được trong package này.
package student

import (
	"fmt"
	"slices"

	"go/quanli/input"
)

type Student struct {
	ID    int
	Name  string
	Class string
	Point float64
}

// String quy định cách in 1 sinh viên. fmt.Println(s) sẽ tự gọi hàm này,
// nên định dạng in chỉ cần viết 1 chỗ.
func (s Student) String() string {
	return fmt.Sprintf("- ID: %d | Ten: %s | Lop: %s | Diem: %.2f", s.ID, s.Name, s.Class, s.Point)
}

// students lưu danh sách sinh viên.
var students []Student

const menu = `
===== QUAN LY SINH VIEN =====
1. Them sinh vien
2. Xoa sinh vien
3. Sua sinh vien
4. Danh sach sinh vien
5. Tim kiem sinh vien
6. Quay lai`

// Menu hiển thị menu quản lý sinh viên, chạy tới khi người dùng chọn "Quay lai".
func Menu() {
	for {
		fmt.Println(menu)
		switch input.Option("👉 Chon chuc nang: ", 1, 6) {
		case 1:
			add()
		case 2:
			remove()
		case 3:
			edit()
		case 4:
			list()
		case 5:
			find()
		case 6:
			return
		}
		input.Pause()
	}
}

// indexOf trả về vị trí sinh viên có ID = id, không có thì trả về -1.
func indexOf(id int) int {
	return slices.IndexFunc(students, func(s Student) bool {
		return s.ID == id
	})
}

// askExisting hỏi ID rồi tìm sinh viên đó. Dùng chung cho xoá, sửa, tìm kiếm.
// ok = false nếu không tìm thấy.
func askExisting() (index int, ok bool) {
	id := input.PositiveInt("- Nhap ID sinh vien: ")
	index = indexOf(id)
	if index == -1 {
		fmt.Println("❌ Khong tim thay sinh vien co ID:", id)
		return -1, false
	}
	return index, true
}

// askInfo nhập tên, lớp, điểm vào s. Dùng chung cho thêm và sửa.
// Nhận con trỏ *Student để sửa thẳng vào biến của nơi gọi.
func askInfo(s *Student) {
	s.Name = input.Name("- Nhap ten: ")
	s.Class = input.ClassCode("- Nhap lop: ")
	s.Point = input.Score("- Nhap diem: ")
}

func add() {
	fmt.Println("===== Them Sinh Vien =====")
	var s Student
	for {
		s.ID = input.PositiveInt("- Nhap ID: ")
		if indexOf(s.ID) == -1 {
			break
		}
		fmt.Println("❌ ID da ton tai! Hay nhap ID khac.")
	}
	askInfo(&s)

	students = append(students, s)
	fmt.Println("✅ Them sinh vien thanh cong!")
}

func remove() {
	fmt.Println("===== Xoa Sinh Vien =====")
	i, ok := askExisting()
	if !ok {
		return
	}
	fmt.Println(students[i]) // in trước khi xoá, sau khi xoá vị trí i là người khác
	students = slices.Delete(students, i, i+1)
	fmt.Println("✅ Xoa sinh vien thanh cong!")
}

func edit() {
	fmt.Println("===== Sua Sinh Vien =====")
	i, ok := askExisting()
	if !ok {
		return
	}
	fmt.Println("Thong tin hien tai:")
	fmt.Println(students[i])

	// &students[i] trỏ thẳng vào phần tử trong slice, nên askInfo sửa
	// trực tiếp sinh viên đó. ID không bị đụng tới nên giữ nguyên.
	askInfo(&students[i])

	fmt.Println("✅ Sua sinh vien thanh cong!")
	fmt.Println(students[i])
}

func list() {
	fmt.Println("======= Danh Sach Sinh Vien =======")
	if len(students) == 0 {
		fmt.Println("(Danh sach trong)")
		return
	}
	for _, s := range students {
		fmt.Println(s)
	}
}

func find() {
	fmt.Println("===== Tim Kiem Sinh Vien =====")
	i, ok := askExisting()
	if !ok {
		return
	}
	fmt.Println("Tim thay sinh vien:")
	fmt.Println(students[i])
}
