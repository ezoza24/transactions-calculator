package main

import "fmt"

func main() {
	var sender string = "Sherali"
	var receiver string = "Alisher"
	var amount int = 500_000
	var amountAfterComission int
	commission := amount / 100
	fmt.Println("========= Чек =========")
	fmt.Printf("Отправитель: %s\n", sender)
	fmt.Printf("Получатель: %s\n", receiver)
	fmt.Printf("Сумма: %d сум\n ", amount)
	fmt.Printf("Комиссия: %d сум\n", commission)
	amountAfterComission = amount + (amount / 100)
	fmt.Printf("Итого: %d cум\n", amountAfterComission)
	fmt.Println("=====================")
}
