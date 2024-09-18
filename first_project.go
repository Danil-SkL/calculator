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
	{"I", 1}, {"II", 2}, {"III", 3}, {"IV", 4}, {"V", 5}, {"VI", 6},
	{"VII", 7}, {"VIII", 8}, {"IX", 9}, {"X", 10}, {"L", 50},
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

	for _, pair := range romanList {
        if pair.symbol == sym {
            fmt.Println(pair.inputNum)
            return pair.inputNum
        }

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

func checkLegit(numOne string) bool {
	checkLegit := false

	for _, check := range romanList {
		if check.symbol == numOne{
			checkLegit = true
			
		}
	}

	if checkLegit == false{
		panic("Вы ввели число которого нет в римской системе счисления")
	}

	return checkLegit
}

func main() {
	var numOne string
	var numTwo string
	var sym string
	checkRomOne := false
	checkRomTwo := false
	answer := 0
	//i := false

	//Считывание ввода пользователем
	fmt.Println("ВВедите значение...")
	n, err := fmt.Scanln(&numOne, &sym, &numTwo)

	//Обработка ошибки, если больше трех значений
	if n != 3 && n > 1 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	//Обработка ошибки, если меньше 3 элементов
	if err != nil {
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	

	//Конвертация из строки в число
	convertOne, err := strconv.Atoi(numOne)
	if err != nil {
		checkLegit(numOne)
		convertOne = serchNum(numOne)
		checkRomOne = true
	}
	convertTwo, err := strconv.Atoi(numTwo)
	if err != nil {
		checkLegit(numTwo)
		convertTwo = serchNum(numTwo)
		checkRomTwo = true
	}

	if convertOne <= 0 || convertOne > 10 || convertTwo <= 0 || convertTwo > 10 {
     panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более.") 
     } else if checkRomOne != checkRomTwo { //Проверка, чтобы оба числа были одной системы счисления
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

		if checkRomOne && checkRomTwo && answer <= 0 { //Проверка чтобы не было минуса
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
