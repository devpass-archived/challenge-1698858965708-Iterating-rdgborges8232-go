package main

import (
	"testing"
)

func TestPrintMap(t *testing.T) {
	employeeSalaries := map[string]int{
		"John Doe": 50000,
		"Jane Smith": 60000,
		"Mike Johnson": 55000,
	}

	PrintMap(employeeSalaries)
}