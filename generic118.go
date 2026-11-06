//go:build ignore

// generic chỉ có từ go 1.18 đỗ lên thay cho interface empty chữ T trong []
package main

import (
	"cmp"
	"fmt"
)

func PrintValue[T any](val T) T {
	return val
}

// comparable return mà so sánh == or != mà sài any là báo đỏ phải comparable
func IsTheSame[T comparable](a, b T) bool {
	return a == b
}

// sài cmp.Ordered vì comparable ko so sánh > < được hoặc chọn cha nó các kiểu dữ liệu
func MaxBetween[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func MinBetween[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generic với struct
type Box[T any] struct {
	Content T
	Desc int  // ← luôn là int, không liên quan T
}

func main() {
	fmt.Println(PrintValue(1))
	fmt.Println(IsTheSame(1, 1.1))
	fmt.Println(MaxBetween(1, 1.1))
	fmt.Println(MinBetween(1, 1.1))

	Boxes := Box[int]{
		Content: 999,
	}

	WhoBox := Box[string]{
		Content: "Lan Box",
		Desc: 1111,  // ← luôn là int, không liên quan T
	}

	fmt.Println(PrintValue(Boxes.Content))
	fmt.Println(PrintValue(WhoBox.Content))
	fmt.Println(PrintValue(WhoBox.Desc))

}
