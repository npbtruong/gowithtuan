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
func remove() {
	fmt.Println("-=-=-=-=-=-=-=-= Xóa Thông tin giảng viên =-=-=-=-=-=-=")

	id := utils.PositiveNumber("Nhập ID giảng viên cần Xóa:")
	
	for index,t := range list_teacher{
		if (t.GetId() == id) {
			//[ 1 2 3 4] [trên : dưới]
			list_front := list_teacher [:index]
			list_back := list_teacher[index+1:]
			list_teacher = append(list_front,list_back...)
			fmt.Println("XÓA Thành Công")
			return
		}
	}
	fmt.Printf("Nhân viên id %d không tồn tại",id)
}

func edit() {

	fmt.Println("-=-=-=-=-=-=-=-= Sữa Thông tin giảng viên =-=-=-=-=-=-=")

	id := utils.PositiveNumber("Nhập ID giảng viên cần sữa:")

	for index, teacher := range list_teacher {
		if teacher.GetId() == id {
			fmt.Println("Nhap thong tin mới (Ấn Enter để giử giá trị củ)")

			name := utils.GetOptionalString(fmt.Sprintf("Nhập Tên:(%s)", teacher.Name), teacher.Name)
			subject := utils.GetOptionalString(fmt.Sprintf("Nhập môn dạy:(%s)", teacher.Subject), teacher.Subject)
			baseSalary := utils.OptionalPositiveFloat(fmt.Sprintf("Nhập lương cơ bản:(%.2f)", teacher.BaseSalary), teacher.BaseSalary)
			bonus := utils.OptionalPositiveFloat(fmt.Sprintf("Nhập thưởng:(%.2f)", teacher.Bonus), teacher.Bonus)

			list_teacher[index] = Teacher{
				Id:         id,
				Name:       name,
				Subject:    subject,
				BaseSalary: baseSalary,
				Bonus:      bonus,
			}

			fmt.Println("CAP NHAT THANH CONG")
			return
		}
	}
	fmt.Println("Chưa có Giảng viên Id này!")

}

func list() {
	for _, value := range list_teacher {
		value.GetInfo()
	}
}

func find() {
	fmt.Println("-=-=-=-=-=-=-=-= Tìm kiếm Thông tin giảng viên =-=-=-=-=-=-=")

	id := utils.PositiveNumber("Nhập ID giảng viên cần tìm:")

	for _,t := range list_teacher{
		if (t.GetId() == id) {
				fmt.Println("ĐÃ tìm thấy giảng viên")
				t.GetInfo()
				return
		}
	}
		fmt.Println("X Không tìm thấy Thông tin giảng viên ID ")
}
