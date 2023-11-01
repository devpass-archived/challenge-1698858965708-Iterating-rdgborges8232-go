package main

import (
	"fmt"
)

func PrintMap(employeeSalaries map[string]int) {
	for name, salary := range employeeSalaries {
		fmt.Printf("%s: %d\n", name, salary)
	}
}

func main() {
	employeeSalaries := map[string]int{
		"John Doe":    50000,
		"Jane Smith":  60000,
		"Mike Johnson": 55000,
	}

	PrintMap(employeeSalaries)
}
