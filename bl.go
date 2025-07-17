package bl

import (
	"fmt"
	"log"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

func Println(v ...interface{}) {
	message := fmt.Sprintln(v...)
	blockMessage := createBlockMessage(message)
	log.Print(blockMessage)
}

func Printf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	blockMessage := createBlockMessage(message)
	log.Print(blockMessage)
}

func Print(v ...interface{}) {
	message := fmt.Sprint(v...)
	blockMessage := createBlockMessage(message)
	log.Print(blockMessage)
}

func createBlockMessage(message string) string {
	if len(message) > 0 && message[len(message)-1] == '\n' {
		message = message[:len(message)-1]
	}

	width := len(message) + 4
	if width < 20 {
		width = 20
	}

	topLine := fmt.Sprintf("%s%s%s", red, repeatChar('-', width), reset)
	contentLine := fmt.Sprintf("%s| %s%s |%s", red, message, repeatChar(' ', width-len(message)-4), reset)
	bottomLine := fmt.Sprintf("%s%s%s", red, repeatChar('-', width), reset)

	return fmt.Sprintf("\n%s\n%s\n%s\n", topLine, contentLine, bottomLine)
}

func repeatChar(char rune, count int) string {
	result := make([]rune, count)
	for i := range result {
		result[i] = char
	}
	return string(result)
}
