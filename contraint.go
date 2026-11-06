//go:build ignore

package main

import (
	"fmt"
)

// contraint là biến ràng buộc
type Number interface {
	int | float64
}

// hoặc T int|float64
func sum[T Number](a, b T) T {
	return a + b
}
func PrintValue[T any](val T) {
	fmt.Println(val)
}
func main() {
	PrintValue(sum(5,5.5))
}
