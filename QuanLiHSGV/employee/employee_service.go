package employee



import (
	"fmt"
	"go/quanli/utils"
)

const menu = `
===== QUAN LY NHAN VIEN =====
1. Them nhan vien
2. Xoa nhan vien
3. Sua nhan vien
4. Danh nhan giang vien
5. Tim kiem nhan vien
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

var list_employee []Employee

func add() {
	var id int
	for {
		id = utils.PositiveNumber("Nhập Id:")
		if utils.IsUnique(id, list_employee) {
			break
		}
		fmt.Println("Id đã tồn tại nhập Id khác:")
	}

	name := utils.GetNonEmptyString("Nhập Tên:")
	salary := utils.PositiveFloat("Nhập lương  :")
	phone := utils.PositiveNumber("Nhập sdt:")

	teacher := Employee{
		Id:         id,
		Name:       name,
		Salary:    salary,
		Phone: phone,
	
	}

	list_employee = append(list_employee, teacher)

	fmt.Println("Thêm thành công")
}
func remove() {}

func edit() {

	fmt.Println("-=-=-=-=-=-=-=-= Sữa Thông tin nhân viên =-=-=-=-=-=-=")

	id := utils.PositiveNumber("Nhập ID nhân viên cần sữa:")

	for index, e := range list_employee {
		if e.GetId() == id {
			fmt.Println("Nhap thong tin mới (Ấn Enter để giử giá trị củ)")

			name := utils.GetOptionalString(fmt.Sprintf("Nhập Tên:(%s)", e.Name), e.Name)
			salary := utils.OptionalPositiveFloat(fmt.Sprintf("Nhập lương :(%.2f)", e.Salary), e.Salary)
			phone := utils.OptionalPositiveNumber(fmt.Sprintf("Nhập sdt:(%d)", e.Phone), e.Phone)

			list_employee[index] = Employee{
				Id:         id,
				Name:       name,
				Salary: salary,
				Phone: phone,
			}

			fmt.Println("CAP NHAT THANH CONG")
			return
		}
	}
	fmt.Println("Chưa có Nhân viên Id này!")

}

func list() {
	for _, value := range list_employee {
		value.GetInfo()
	}
}


func find() {
	fmt.Println("-=-=-=-=-=-=-=-= Tìm kiếm Thông tin nhân viên =-=-=-=-=-=-=")

	id := utils.PositiveNumber("Nhập ID nhân viên cần tìm:")

	for _,t := range list_employee{
		if (t.GetId() == id) {
				fmt.Println("ĐÃ tìm thấy nhân viên")
				t.GetInfo()
				return
		}
	}
		fmt.Println("X Không tìm thấy Thông tin nhân viên ID ")
}
