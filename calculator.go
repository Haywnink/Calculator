package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var RomanToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var ArabicToRoman = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 1 + 1 или X - V):")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := calculate(input)
	fmt.Println("Результат:", result)
}

func calculate(input string) string {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Неправильный формат ввода.")
	}

	num1, isRoman1 := parseNumber(parts[0])
	num2, isRoman2 := parseNumber(parts[2])

	if isRoman1 != isRoman2 {
		panic("Нельзя смешивать римские и арабские числа.")
	}

	result := applyOperation(num1, num2, parts[1])

	if isRoman1 {
		if result < 1 {
			panic("В римской системе нет отрицательных чисел.")
		}
		return toRoman(result)
	}
	return strconv.Itoa(result)
}

func parseNumber(s string) (int, bool) {
	if value, isRoman := RomanToArabic[s]; isRoman {
		return value, true
	}
	num, err := strconv.Atoi(s)
	if err != nil || num < 1 || num > 10 {
		panic("Число должно быть от 1 до 10.")
	}
	return num, false
}

func applyOperation(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль")
		}
		return num1 / num2
	default:
		panic("Неправильный оператор")
	}
}

func toRoman(num int) string {
	var roman strings.Builder

	romanValues := []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, rv := range romanValues {
		for num >= rv.Value {
			roman.WriteString(rv.Symbol)
			num -= rv.Value
		}
	}
	return roman.String()
}
