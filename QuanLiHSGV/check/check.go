// Package check chứa các hàm kiểm tra dữ liệu người dùng nhập.
//
// Mỗi hàm nhận chuỗi người dùng gõ vào, trả về giá trị đã chuẩn hoá
// hoặc error mô tả vì sao không hợp lệ. Các hàm ở đây KHÔNG tự in ra
// màn hình, việc in lỗi và hỏi lại là của package input.
package check

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IntInRange: s phải là số nguyên nằm trong khoảng [lo, hi]. Dùng cho chọn menu.
func IntInRange(s string, lo, hi int) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n < lo || n > hi {
		return 0, fmt.Errorf("Gia tri khong hop le! Hay nhap so tu %d den %d.", lo, hi)
	}
	return n, nil
}

// PositiveInt: s phải là số nguyên > 0. Dùng cho ID.
// "abc", "1.5", "-3", "" đều bị từ chối.
func PositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n <= 0 {
		return 0, errors.New("Gia tri khong hop le! Hay nhap so nguyen duong.")
	}
	return n, nil
}

// Name: chỉ gồm chữ cái (có dấu cũng được) và khoảng trắng, dài 2-50 ký tự.
// Trả về tên đã chuẩn hoá: "  nguyen   VAN  a " -> "Nguyen Van A".
func Name(s string) (string, error) {
	words := strings.Fields(s) // tách theo khoảng trắng, tự bỏ khoảng trắng thừa
	if len(words) == 0 {
		return "", errors.New("Ten khong duoc de trong.")
	}
	for i, w := range words {
		for _, r := range w {
			// IsMark: cho phép dấu tiếng Việt khi bộ gõ dùng kiểu "Unicode to hop"
			if !unicode.IsLetter(r) && !unicode.IsMark(r) {
				return "", errors.New("Ten chi duoc chua chu cai va khoang trang.")
			}
		}
		words[i] = capitalize(w)
	}

	name := strings.Join(words, " ")
	if n := utf8.RuneCountInString(name); n < 2 || n > 50 {
		return "", errors.New("Ten phai dai tu 2 den 50 ky tu.")
	}
	return name, nil
}

// capitalize viết hoa chữ đầu, viết thường phần còn lại: "nGUYEN" -> "Nguyen".
func capitalize(w string) string {
	r, size := utf8.DecodeRuneInString(w)
	return string(unicode.ToUpper(r)) + strings.ToLower(w[size:])
}

// ClassCode: mã lớp chỉ gồm chữ cái không dấu, chữ số và dấu '-', dài 1-20 ký tự.
// Trả về dạng viết hoa: "k15-cntt" -> "K15-CNTT".
func ClassCode(s string) (string, error) {
	if s == "" {
		return "", errors.New("Lop khong duoc de trong.")
	}
	for _, r := range s {
		ok := ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9') || r == '-'
		if !ok {
			return "", errors.New("Lop chi duoc chua chu cai khong dau, chu so va dau '-'.")
		}
	}
	if len(s) > 20 {
		return "", errors.New("Lop toi da 20 ky tu.")
	}
	return strings.ToUpper(s), nil
}

// Score: điểm là số từ 0 đến 10, chấp nhận cả dấu phẩy "8,5".
func Score(s string) (float64, error) {
	s = strings.Replace(s, ",", ".", 1)

	// Chỉ cho phép chữ số và tối đa 1 dấu chấm.
	// Nếu không chặn, strconv.ParseFloat sẽ nhận cả "1e1", "NaN", "Inf"...
	onlyDigits := strings.Trim(s, "0123456789.") == ""
	if s == "" || !onlyDigits || strings.Count(s, ".") > 1 {
		return 0, errors.New("Diem khong hop le! Hay nhap so tu 0 den 10.")
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil || f < 0 || f > 10 {
		return 0, errors.New("Diem khong hop le! Hay nhap so tu 0 den 10.")
	}
	return f, nil
}
