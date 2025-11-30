package input

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadInputString(path string) string {
	content, err := os.ReadFile(fmt.Sprintf("%s", path))

	if err != nil {
		panic(fmt.Sprintf("Error when reading input file at %s: %s\n", path, err))
	}

	text := string(content)
	text = strings.TrimSuffix(text, "\n")
	return text
}

func ReadInputLines(path string) []string {
	inputString := ReadInputString(path)
	return strings.Split(inputString, "\n")
}

func ReadInputLetters(path string) [][]string {
	splitString := [][]string{}
	inputString := ReadInputLines(path)
	for _, row := range inputString {
		splitString = append(splitString, strings.Split(row, ""))
	}
	return splitString
}

func ReadInputRegex(path string, regex string) []string {
	inputString := ReadInputString(path)
	pattern, err := regexp.Compile(regex)

	if err != nil {
		panic(fmt.Sprintf("Error when compiling regex to ReadInputRegex: %s\n", err))
	}
	return pattern.FindAllString(inputString, -1)
}
