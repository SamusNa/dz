package main

import "fmt"

var baseCurrency string
const usdToEUR = 0.8553
const usdToRUB = 80.43
const eurToRUB = usdToRUB/usdToEUR

func main() {
	baseCurrency = getUserInputBase()
	sum := getUserInputSum()
	toCurrency := getUserInputToCurrency()
	result, currency := converter(baseCurrency, sum, toCurrency)
	fmt.Printf("%.2f %v\n", result, currency)
}

func converter(baseCurrency string, sum float64, toCurrency string) (float64, string) {
	switch baseCurrency {
	case "USD":
		if toCurrency == "EUR" {
			return sum * usdToEUR, "EUR"
		} else {
			return sum * usdToRUB, "RUB"
		}
	case "EUR":
		if toCurrency == "USD" {
			return sum / usdToEUR, "USD"
		} else {
			return sum * eurToRUB, "RUB"
		}
	default:
		if toCurrency == "USD" {
			return sum / usdToRUB, "USD"
		} else {
			return sum / eurToRUB, "EUR"
		}
	}
}

func getUserInputBase() string {
	for {
		fmt.Print("Введите исходную валюту (EUR, USD или RUB): ")
		fmt.Scan(&baseCurrency)
		switch baseCurrency {
		case "EUR", "USD", "RUB":
			return baseCurrency
		default:
			fmt.Println("Ошибка: введите исходную валюту заново (EUR, USD или RUB): ")
		}
	}
}

func getUserInputSum() float64 {
	var sum float64
	var input string
	for {
		fmt.Print("Введите сумму для конвертации: ")
		fmt.Scan(&input)
		n, err := fmt.Sscanf(input, "%f", &sum)
		if err == nil && n == 1 {
			return sum
		} else {
			fmt.Println("Ошибка: введите сумму заново")
		}
	}	
}

func getUserInputToCurrency () string {
	var toCurrency string
	for {
	switch baseCurrency {
	case "EUR":
		fmt.Println("Введите целевую валюту (USD или RUB): ")
	case "USD":
		fmt.Println("Введите целевую валюту (EUR или RUB): ")
	case "RUB":
		fmt.Println("Введите целевую валюту (EUR или USD): ")
	}
	fmt.Scan(&toCurrency)
	switch toCurrency {
		case "EUR", "USD", "RUB":
			return toCurrency
		default:
			switch baseCurrency {
			case "EUR":
				fmt.Println("Ошибка: введите целевую валюту заново (USD или RUB): ")
			case "USD":
				fmt.Println("Ошибка: введите целевую валюту заново (EUR или RUB): ")
			case "RUB":
				fmt.Println("Ошибка: введите целевую валюту заново (EUR или USD): ")
			}
		}
	}
}