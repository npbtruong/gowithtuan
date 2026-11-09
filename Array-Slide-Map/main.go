package main

import (
	//"arrayslidemap/array"
	"fmt"
	"reflect"
)

func main() {

	//array.Array()
	slice := []int{1,2,3,4,5}
	num := 12
	fmt.Println(slice)
	//kiểm tra type bằng reflect
	fmt.Println(reflect.TypeOf(slice).Kind())
	fmt.Println(reflect.TypeOf(num).Kind())
}
