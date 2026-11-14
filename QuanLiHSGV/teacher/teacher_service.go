package teacher

import (
	"fmt"
	"go/quanli/utils"
)

const menu = `
===== QUAN LY GIANG VIEN =====
1. Them giang vien
2. Xoa giang vien
3. Sua giang vien
4. Danh sach giang vien
5. Tim kiem giang vien
6. Quay lai`

// Menu hiển thị menu quản lý sinh viên, chạy tới khi người dùng chọn "Quay lai".
func Menu() {
	for {
		utils.ClearScreen()
		fmt.Println(menu)
		switch utils.PositiveNumber("Chọn chức năng : ") {
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
		default:
			fmt.Println("Lựa chọn ko hợp lệ!!")
		}
		utils.ReadInput("\nNhan Enter De Tiep Tuc")
	}
}

var list_teacher []Teacher

func add() {
	var id int
	for {
		id = utils.PositiveNumber("Nhập Id:")
		if utils.IsUnique(id, list_teacher) {
			break
		}
		fmt.Println("Id đã tồn tại nhập Id khác:")
	}

	name := utils.GetNonEmptyString("Nhập Tên:")
	subject := utils.GetNonEmptyString("Nhập môn dạy:")
	baseSalary := utils.PositiveFloat("Nhập lương cơ bản:")
	bonus := utils.PositiveFloat("Nhập thưởng:")

	teacher := Teacher{
		Id:         id,
		Name:       name,
		Subject:    subject,
		BaseSalary: baseSalary,
		Bonus:      bonus,
	}

	list_teacher = append(list_teacher, teacher)

	fmt.Println("Thêm thành công")
}
func remove() {}
func edit()   {}
func list() {
	for _, value := range list_teacher {
		value.GetInfo()
	}
}
func find() {}
