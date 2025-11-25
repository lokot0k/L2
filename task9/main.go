package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func unpack(s string) (string, error) {
	var result []rune
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			return "", errors.New("невалидная строка")
		}

		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			j := i + 1
			for j < len(runes) && unicode.IsDigit(runes[j]) {
				j++
			}

			count, err := strconv.Atoi(string(runes[i+1 : j]))
			if err != nil {
				return "", err
			}

			for k := 0; k < count; k++ {
				result = append(result, runes[i])
			}
			i = j - 1
		} else {
			result = append(result, runes[i])
		}
	}

	return string(result), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")
	if scanner.Scan() {
		input := scanner.Text()
		result, err := unpack(input)
		if err != nil {
			panic("panic: " + err.Error())
		}
		fmt.Println(result)
	}
}
