//go:build ignore

package main

import "fmt"

func main() {
	/*
		# make() dùng để cấp phát ô nhớ cho 3 kiểu tham chiếu GO: các kiểu này ko có ô chứa 8 byte sẵn
		- slice,array s:=make([]int) , make([]int, 5) (độ dài 5) hoặc make([]int, 0, 10) (độ dài 0, sức chứa 10)
		- map m := make(map[int]string)
		- channel c := make(Chan int)
	*/

	//Map tương tự array và slice nhưng được tùy biến key thay vì index 0->...
	// map[key type]value type
	// Khởi tạo gán trực tiếp
	menu := map[string]float32{
		"cháo": 10.1223232,
		"phở":  11.13232,
		"bún":  12.232323,
	}
	menu["cơm tấm"] = 13
	fmt.Println(menu)
	// khởi tạo với make phổ biến
	city := make(map[string]string)
	city["hcm"] = "Hồ Chí Minh City"
	city["hn"] = "Hà Nội"
	city["dn"] = "Đà Nẵng"
	fmt.Println(city)
	// khởi tạo map rỗng ko ít dùng
	var mssv map[int]string
	mssv = make(map[int]string)
	mssv[1] = "An"
	mssv[2] = "Bình"
	mssv[3] = "Giang"
	fmt.Println(mssv)

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	// value,exist
	value, exist := menu["bún"]
	if exist {
		fmt.Printf("tồn tại value= %.2f\n", value)
	} else {
		fmt.Println("Key không tồn tại trong map")
	}
	// range

	for key, value := range menu {
		fmt.Printf("Món %s giá %v \n", key, value)
	}

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	//Duyệt dữ liệu map+struct và map+slice
	// map struct
	list_employee_department := map[string]Employee{
		"RnD":      {Name: "Phúc", Age: 23, Role: "Php Developer"},
		"security": {Name: "tèo", Age: 20, Role: "Bảo vệ"},
		"HR":       {Name: "Thúy", Age: 30, Role: "Admin"},
	}
	
	val,exi := list_employee_department["HR"]
	if exi {
		fmt.Println(list_employee_department["HR"].Name)
		fmt.Println(val.Age)
	} else {
		fmt.Println("Key không tồn tại trong map")
	}

	for key, value := range list_employee_department {
		fmt.Printf("Phòng ban %s nhân viên %s tổi %d chức vụ %s \n", key, value.Name, value.Age, value.Role)
	}

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	// slice struct
	studentClass := map[string] []string {
		"tuan" : {"cntt","ctdl"},
		"tu" : {"lsd","htn","vm"},
		"an" : {"doan"},
	}

	for key,value := range studentClass {
		fmt.Printf("Học sinh: %s học các môn %v \n",key,value)

		for _,val := range value {
			fmt.Printf("Môn Học : %s \n",val)
		}
	}

	
}

type Employee struct {
	Name string
	Age  int
	Role string
}
