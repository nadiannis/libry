package utils

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

func GetInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func FormatDate(datetime time.Time) string {
	format := "02 Jan 2006"
	return datetime.Format(format)
}
