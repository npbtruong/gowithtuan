package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"uuid"
)

func GenerateUuid() string{
	return uuid.New().String()
}

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

func GetOptionalString(promt string, oldValue string) string {

	input := ReadInput(promt)

	if input != "" {
		return input
	}
	return oldValue

}

func OptionalPositiveFloat(promt string, oldValue float64) float64 {

	input := ReadInput(promt)
	value, err := strconv.ParseFloat(input, 64)
	if err == nil && value > 0 {
		return value
	}
	return oldValue

}

func OptionalPositiveNumber(promt string, oldValue int) int {

	input := ReadInput(promt)
	value, err := strconv.Atoi(input)
	if err == nil && value > 0 {
		return value
	}
	return oldValue

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