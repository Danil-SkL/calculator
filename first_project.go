package main

import (
	"fmt"
	"strconv"
)

// Карта на римский вариант
var romanList = []struct {
	symbol   string
	inputNum int
}{
	{"I", 1}, {"V", 5}, {"X", 10}, {"L", 50},
}

// Карта с римского варианта
var romanMap = []struct {
	decVal int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

// Перевод с римского
func serchNum(sym string) int {
	var zeroInt = 0

	for j := 0; j < len(sym); j++ {
		for _, i := range romanList {
			if string(sym[j]) == i.symbol {
				zeroInt += i.inputNum
			}
		}
	}

	return zeroInt
}

// Перевод в римские цифры
func decimalToRomanRecursive(num int) string {
	for _, pair := range romanMap {
		if num >= pair.decVal {
			return pair.symbol + decimalToRomanRecursive(num-pair.decVal)
		}
	}
	return ""
}

func main() {
	var numOne string
	var numTwo string
	var sym string
	checkRomOne := false
	checkRomTwo := false
	answer := 0

	//Считывание ввода пользователем
	fmt.Println("ВВедите значение...")
	n, err := fmt.Scanln(&numOne, &sym, &numTwo)

	//Обработка ошибки, если меньше 3 элементов
	if err != nil {
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	//Обработка ошибки, если больше трех значений
	if n != 3 && n > 2 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	//Конвертация из строки в число
	convertOne, err := strconv.Atoi(numOne)
	if err != nil {
		convertOne = serchNum(numOne)
		checkRomOne = true
	}
	convertTwo, err := strconv.Atoi(numTwo)
	if err != nil {
		convertTwo = serchNum(numTwo)
		checkRomTwo = true
	}

	if checkRomOne != checkRomTwo { //Проверка, чтобы оба числа были одной системы счисления
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	} else if checkRomOne == checkRomTwo { //Подсчет ответа
		switch sym {
		case "+":
			answer = convertOne + convertTwo
		case "-":
			answer = convertOne - convertTwo
		case "*":
			answer = convertOne * convertTwo
		case "/":
			answer = convertOne / convertTwo
		}

		if checkRomOne && checkRomTwo && answer < 0 { //Проверка чтобы не было минуса
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		} else if checkRomOne && checkRomTwo { //Вывод для римского варианта
			fmt.Println("Ваш ответ...")
			fmt.Println(decimalToRomanRecursive(answer))
		} else { //Обычный вывод
			fmt.Println("Ваш ответ...")
			fmt.Println(answer)
		}
	}
}
