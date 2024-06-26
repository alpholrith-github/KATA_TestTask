package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInteger(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func romanToInt(romanNumber string) (int, bool) {
	table := []struct {
		number int
		symbol string
	}{
		{1, "I"}, {2, "II"}, {3, "III"},
		{4, "IV"}, {5, "V"}, {6, "VI"},
		{7, "VII"}, {8, "VIII"}, {9, "IX"},
		{10, "X"},
	}
	i := 0
	for i < len(table) {
		if strings.EqualFold(romanNumber, table[i].symbol) == true {
			return table[i].number, true
		}
		i++
	}
	return -1, false
}

func intToRoman(x int) string {
	table := []struct {
		number int
		symbol string
	}{
		{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"},
		{10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"},
		{100, "C"},
	}
	result := ""
	i := len(table) - 1

	for x > 0 {
		for table[i].number > x {
			i--
		}
		result += table[i].symbol
		x -= table[i].number
	}

	return result
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")
		if len(args) != 3 {
			fmt.Println("Неверное количество аргументов")
			return
		}

		aStr := args[0]
		operator := args[1]
		bStr := args[2]

		validOperators := "+-*/"
		if !strings.Contains(validOperators, operator) || len(operator) != 1 {
			fmt.Println("Ошибка ввода оператора")
			return
		}

		if isInteger(aStr) && isInteger(bStr) {
			a, _ := strconv.Atoi(aStr)
			b, _ := strconv.Atoi(bStr)
			if a < 1 || b < 1 {
				fmt.Println("Один или оба аргумента меньше 1")
				return
			}
			if a > 10 || b > 10 {
				fmt.Println("Один или оба аргумента больше 10")
				return
			}
			switch operator {
			case "+":
				fmt.Println(a + b)
			case "-":
				fmt.Println(a - b)
			case "*":
				fmt.Println(a * b)
			case "/":
				fmt.Println(a / b)
			}
		} else {
			a, aStrConfirm := romanToInt(aStr)
			b, bStrConfirm := romanToInt(bStr)
			if (aStrConfirm && bStrConfirm) == true {

				switch operator {
				case "+":
					fmt.Println(intToRoman(a + b))
				case "*":
					fmt.Println(intToRoman(a * b))
				case "/":
					if a > b {
						fmt.Println(intToRoman(a / b))
					} else {
						fmt.Println("PANIC: result < 1")
						continue
					}
				case "-":
					if a > b {
						fmt.Println(intToRoman(a - b))
					} else {
						fmt.Println("PANIC: Oтрицательных римских чисел не существует или результат < 1")
						continue
					}
				}
			} else {
				fmt.Println("Вводимые аргументы не соответствуют требованиям")
				return
			}
		}
	}
}
