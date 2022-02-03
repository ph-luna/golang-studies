package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	option, firstNumber, secondNumber := reader()

	n1, _ := strconv.Atoi(firstNumber)
	n2, _ := strconv.Atoi(secondNumber)

	switch option {
	case "1":
		result := soma(n1, n2)

		fmt.Println("Resultado: ", result)
		break

	case "2":
		result := sub(n1, n2)

		fmt.Println("Resultado: ", result)
		break

	case "3":
		result := mult(n1, n2)

		fmt.Println("Resultado: ", result)
		break

	case "4":
		result := division(n1, n2)

		fmt.Println("Resultado: ", result)
		break
	}
}

func reader() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	var option, firstNumber, secondNumber string

	fmt.Println("Informe a operação desejada:")
	fmt.Println("---------------------")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtracao")
	fmt.Println("3 - Multiplicacao")
	fmt.Println("4 - Divisao")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	option = strings.Replace(text, "\r\n", "", -1)

	fmt.Println("Informe o primeiro numero")

	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')

	firstNumber = strings.Replace(text, "\r\n", "", -1)

	fmt.Println("Informe o segundo numero")

	fmt.Print("-> ")
	text, _ = reader.ReadString('\n')

	secondNumber = strings.Replace(text, "\r\n", "", -1)

	return option, firstNumber, secondNumber
}

func soma(n1, n2 int) int {
	return n1 + n2
}

func sub(n1, n2 int) int {
	return n1 - n2
}

func mult(n1, n2 int) int {
	return n1 * n2
}

func division(n1, n2 int) int {
	if n2 == 0 {
		fmt.Println("Erro: Impossivel executar divisao por zero.")

		return 0
	}

	return n1 / n2
}
