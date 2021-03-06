package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func IsPresent(numbers []int, value int) bool {
	for _, number := range numbers {
		if number == value {
			return true
		}
	}

	return false
}

func CanHavePandigitalProduct(multiplier, multiplicand int) bool {
	both := strconv.Itoa(multiplier) + strconv.Itoa(multiplicand)
	if len(both) != 5 {
		return false
	}

	for _, number := range both {
		if number == '0' || strings.Count(both, string(number)) > 1 {
			return false
		}
	}

	return true
}

func HasPandigitalProduct(multiplier, multiplicand int) (int, bool) {
	sequence := []int{ '1', '2', '3', '4', '5', '6', '7', '8', '9', }

	product := multiplier * multiplicand
	all := strconv.Itoa(multiplier) + strconv.Itoa(multiplicand) + strconv.Itoa(product)
	allBytes := []int(all)

	if len(allBytes) != 9 {
		return 0, false
	}

	sort.Ints(allBytes)
	for i, number := range allBytes {
		if number != sequence[i] {
			return 0, false
		}
	}

	return product, true
}

func main() {
	products := make([]int, 0)

	for multiplier := 1; multiplier < 10000; multiplier++ {
		for multiplicand := 1; multiplicand < 100; multiplicand++ {
			if CanHavePandigitalProduct(multiplier, multiplicand) {
				if product, ok := HasPandigitalProduct(multiplier, multiplicand); ok && !IsPresent(products, product) {
					products = append(products, product)
				}
			}
		}
	}

	sum := 0
	for _, product := range products {
		sum += product
	}
	fmt.Println(sum)
}
