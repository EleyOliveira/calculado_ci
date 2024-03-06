package main

import "fmt"

func main() {
	fmt.Println(Multiplicacao(20, 10))
}

func Multiplicacao(num1 int, num2 int) int {
	return num1*num2 + 3
}
