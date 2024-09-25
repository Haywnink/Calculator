package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	re := regexp.MustCompile(`^\"(.+?)\"\s*([\+\-\*\/])\s*(\".*\"|\d+)$`)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 4 {
		panic("Неверный формат ввода")
	}

	str1 := matches[1]
	op := matches[2]
	str2 := strings.Trim(matches[3], `"`)

	var num int
	var err error

	if op == "*" || op == "/" {
		num, err = strconv.Atoi(str2)
		if err != nil || num < 1 || num > 10 {
			panic("Неверное число")
		}
	}

	var result string
	switch op {
	case "+":
		result = str1 + str2
	case "-":
		if strings.Contains(str1, str2) {
			result = strings.Replace(str1, str2, "", 1)
		} else {
			result = str1
		}
	case "*":
		result = strings.Repeat(str1, num)
	case "/":
		result = str1[:len(str1)/num]
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Printf("\"%s\"\n", result)
}
