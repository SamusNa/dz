package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func main() {
	operation := inputOperation()
	str := inputNumbers()
	result := result(operation, str)
	fmt.Println(result)
}

func inputOperation() string {
	for {
		var operation string
		fmt.Println("Введите операцию AVG, SUM или MED: ")
		fmt.Scan(&operation)
		switch operation {
		case "AVG", "SUM", "MED":
			return operation
		default:
			fmt.Println("Ошибка: введите операцию заново")
		}
	}
}

func inputNumbers() string {
	var str string
	fmt.Println("Введите числа через запятую: ")
	fmt.Scan(&str)
	return str
}

func result(operation string, str string) string {
	numbers := strings.Split(str, ",")
	switch operation {
	case "AVG":
		a := 0
		sum := 0
		for _, value := range numbers {
			num, _ := strconv.Atoi(value)
			sum += num
			a++
		}
		avg := float64(sum)/float64(a)
		return fmt.Sprintf("Среднее равно: %v", avg)
	case "SUM":
		sum := 0
		for _, value := range numbers {
			num, _ := strconv.Atoi(value)
			sum += num
		}
		return fmt.Sprintf("Сумма равна: %v", sum)
	case "MED":
		var nums []float64
		for _, x := range strings.Split(str, ",") {
			x = strings.TrimSpace(x)
			if x == "" {
				continue
			}
			num, _ := strconv.ParseFloat(x, 64)
			nums = append(nums, num)
		}
		sort.Float64s(nums)
		b := len(nums)
		med := nums[b/2]
		if b%2 == 0 {
			med = (nums[b/2-1] + nums[b/2])/2
		}
		return fmt.Sprintf("Медиана равна: %v", med)
	}
	return ""
}