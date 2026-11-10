package slice

import (
	"fmt"
	"reflect"
)

func Slice() {
	slice_ := []int{1, 2, 3, 4, 5}
	fmt.Println(slice_)
	fmt.Println(reflect.TypeOf(slice_).Kind())



	array := [5]int{10,20,30,40,50}
	// slice cắt từ mãng array
	slice := array[0:3] // bắt đầu từ 0 đến n-1 tức 0 đến 2 => 10 20 30
	fmt.Println(slice)

	// khỏi tạo slice mới make
	// make( []kiểu dữ liệu , chiều dài , capacity chiều dài tối đa được append thường x2 cap hiện có)
	// capacity dung lượng tối đa slice đạt được mà không phải cấp phát bộ nhớ mới, nhưng nó không phải giới hạn cứng.
	make_slice := make([]int,3,5) // tương đương slice []int{0,0,0}
	make_slice[0]=1
	fmt.Println(make_slice)
	make_slice = append(make_slice, 3,7,9)
	fmt.Println(make_slice,cap(make_slice))

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	fruit_slice := []string{"apple","strawberry","grape","watermelon"}
	//for i:=0 ; i < len(fruit_slice);i++{}
	for _,val := range fruit_slice{
		fmt.Println(val)
	}

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	// matrix slide
	matrix_slice := [][]int{
		{1,2,3},
		{0,0,0},
	}
	
	for _,value := range matrix_slice{
		for index,val := range value{
			fmt.Print(val)
			if(index == (len(value)-1)){fmt.Println()}
		}
	}

	// gộp 2 slide ... là lấy tất cả
	slice1 := []int{1,2}
	slice2 := []int{3,4}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	// tạo subslice từ slice cha
	subSlice := slice1[1:3]
	fmt.Println(subSlice)// trên  0 1 dưới 1 2 3 => [2 3]
	fmt.Println(slice1[:2])// tức trên 0 dưới 2 => [1 2]
	fmt.Println(slice1[:2])// tức trên 0 dưới 2 => [1 2]
}
