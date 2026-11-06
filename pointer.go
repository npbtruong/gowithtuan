//go:build ignore

package main

import "fmt"

/*
pointer giúp tránh copy dữ liệu ko cần thiết và update dữ liệu trực tiếp
Mặc định dùng T. Chuyển sang *T khi cần sửa dữ liệu, struct lớn, hoặc cần dùng chung. 
Một type đã có method nào dùng *T thì cho tất cả method của type đó dùng *T luôn.
 */
func lythuyenPointer() {
	name := "UrGod"

	fmt.Println("------- ÌNOMATION ------")
	fmt.Printf("Data Type %T \n",name)
	fmt.Printf("Data Value %v \n",name)
	fmt.Printf("Data Address %v \n",&name)

	// lưu địa chỉ con trỏ của biến name vào 1 biến khác
	ptrName := &name
	fmt.Printf("name Address %v \n",ptrName)
	//bản thân cái biến lưu cái địa chỉ cũng có địa chỉ của nó
	fmt.Printf("ptrName Address %v \n",&ptrName)

	// dựa vào địa chỉ con trỏ có thể check value địa chỉ đó lưu gì
	fmt.Printf("Nhìn vô * 1 địa chỉ thấy gì , nhìn vô *ptrName: %v \n", *ptrName)


	// tạo biến chứa biến chứa địa chỉ của name
	ptrName2 := &ptrName

	fmt.Printf("dòm vô ptrName thấy địa chỉ rồi lại dòm vào địa chỉ thấy name value : %v \n", **ptrName2)
}


// khi re vào con trỏ sẽ có kiểu dữ liệu của nó bạn nên set đúng kiểu dữ liệu vd *string cho name
func updateName(name *string){
	*name = "Your God"// biến mới vùng nhớ mới
}

func main(){
	name := "Ur God"
	
	name = "Urgot?" // cập nhật giá trị

	updateName(&name)
	fmt.Printf("Data Type %T \n",name)
	fmt.Printf("Data Value %v \n",name)
	fmt.Printf("Data Address %v \n",&name)



}