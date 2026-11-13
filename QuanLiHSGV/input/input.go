// Package input đọc dữ liệu từ bàn phím bằng bufio.
//
// Mỗi lần đọc lấy nguyên 1 dòng (kể cả '\n'), nên không còn bị sót
// ký tự xuống dòng như fmt.Scan, và đọc được tên có dấu cách.
// Các hàm ở đây tự hỏi lại cho tới khi người dùng nhập đúng.
package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go/quanli/check"
)

// Dùng chung 1 reader cho cả chương trình.
// Nếu mỗi hàm tự tạo bufio.NewReader riêng, phần dữ liệu reader cũ đã
// nạp vào bộ đệm mà chưa dùng tới sẽ bị mất.
var reader = bufio.NewReader(os.Stdin)

func Line(prompt string) string {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		fmt.Println("\nKet thuc chuong trinh.")
		os.Exit(0)
	}
	return strings.TrimSpace(line)
}

// ask hỏi đi hỏi lại cho tới khi hàm parse chấp nhận giá trị.
// T là kiểu kết quả (int, string, float64...). Nhờ generic nên vòng lặp
// "nhập -> kiểm tra -> sai thì nhập lại" chỉ cần viết 1 lần ở đây.
func ask[T any](prompt string, parse func(string) (T, error)) T {
	for {
		value, err := parse(Line(prompt))
		if err == nil {
			return value
		}
		fmt.Println("❌", err)
	}
}

func Option(prompt string, lo, hi int) int {
	return ask(prompt, func(s string) (int, error) {
		return check.IntInRange(s, lo, hi)
	})
}

// PositiveInt đọc số nguyên dương (ví dụ ID).
func PositiveInt(prompt string) int {
	return ask(prompt, check.PositiveInt)
}

// Name đọc họ tên, chỉ gồm chữ cái và khoảng trắng.
func Name(prompt string) string {
	return ask(prompt, check.Name)
}

// ClassCode đọc mã lớp, ví dụ "CNTT", "10A1", "K15-QTKD".
func ClassCode(prompt string) string {
	return ask(prompt, check.ClassCode)
}

// Score đọc điểm từ 0 đến 10.
func Score(prompt string) float64 {
	return ask(prompt, check.Score)
}

// Pause dừng màn hình cho tới khi người dùng bấm Enter.
func Pause() {
	Line("\nNhan phim Enter de tiep tuc...")
}
