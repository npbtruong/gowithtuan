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
	input := ReadInput(promt)
	value, err := strconv.Atoi(input)
	if err == nil && value > 0 {
		return value
	}
	fmt.Println("Invalid Input!")
	return 0
}

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "window":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
	if err:= cmd.Run();err != nil{
		fmt.Println("Error clearing screen: ", err)
	}
}
