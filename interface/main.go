package main

import (
	"fmt"
	"go/go/cat"
	"go/go/dog"
	"go/go/mouse"
	"go/go/service"
)

//interface là tập hợp các method 1 kiểu dữ liệu phải triển khai
/*
	Nếu interface chỉ 1 method thì nên đặt tên giống method thêm er
	type Sender interface {
		send() float32
	}

	Mặc định dùng T. Chuyển sang *T khi cần sửa dữ liệu, struct lớn, hoặc cần dùng chung.
	Một type đã có method nào dùng *T thì cho tất cả method của type đó dùng *T luôn.
*/

func MakeSound(a service.Animal) {
	fmt.Printf("Tiếng của %v \n", a.GetName())
	fmt.Println(a.Sound())
}

func MakeSoundExtendsion(a service.AnimalExtension) {
	fmt.Printf("Tiếng của %v \n", a.GetName())
	fmt.Println(a.Sound())
	fmt.Println(a.EatSound())
}

func PrintValue(val interface{}) {
	// (kiểu) assertion
	// s, ok := val.(string) 
	// if ok {
	// 	fmt.Println(s)
	// } else {
	// 	fmt.Println("Not a string")
	// }

	// type switch  .(type) chỉ áp dụng cho switch case
	switch val.(type) {
		case int:
			fmt.Println(val)
		case string:
			fmt.Println(val)
		default:
			fmt.Println("-=-=Not a string nor int=-=-")
	}
}

func main() {
	mydog := dog.Dog{
		Name: "GouDog",
	}

	// Đi làm viết sẵn hàm tạo mới tên New()
	mycat, err := cat.New(" PussyCat  ", "")
	if err != nil {
		panic(err)
	}

	shecat, err := cat.New(" HerCat  ", "")
	if err != nil {
		panic(err)
	}

	PrintValue("-=-=-=-=-=-=-=-=-=-")
	MakeSound(mydog)
	PrintValue("-=-=-=-=-=-=-=-=-=-")
	MakeSound(mycat)
	PrintValue("-=-=-=-=-=-=-=-=-=-")
	MakeSoundExtendsion(shecat)

	mymouse, err := mouse.New("CHÍT")
	if err != nil {
		panic(err)
	}
	PrintValue(12.3)
	fmt.Println(mymouse.Bite())

}
