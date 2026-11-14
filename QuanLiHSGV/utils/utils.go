package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func PositiveNumber(promt string) int {
	for {
		input := ReadInput(promt)
		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Invalid Input!")
	}
}

func GetNonEmptyString(promt string) string {
	for {
		input := ReadInput(promt)

		if input != "" {
			return input
		}
		fmt.Println("Invalid Input!")
	}
}

func PositiveFloat(promt string) float64 {
	for {
		input := ReadInput(promt)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value > 0 {
			return value
		}
		fmt.Println("Invalid Input!")
	}
}

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing screen: ", err)
	}
}



// HasId: interface làm constraint(ràng buộc) , kiểu nào có GetId() int thì tự động thỏa (không cần implements).
type HasId interface{
	GetId() int
}

// IsUnique: hàm generic, dùng chung cho mọi []T có GetId(), trả về true nếu id chưa tồn tại.
func IsUnique[T HasId](id int, list []T) bool {
	for _, value := range list {
		if value.GetId() == id {
			return false
		}
	}
	return true
}