package main

import (
	"fmt"
	"strconv"
)

func BuildPrimesList(limit float64) []int {
	primes := make([]int, 0)
	primesCh := GetPrimes(limit)

	for prime := range primesCh {
		if prime == 0 {
			break
		}

		primes = append(primes, prime)
	}

	return primes
}

func IsPresent(slice []int, number int) bool {
	for _, element := range slice {
		if element == number {
			return true
		}
	}

	return false
}

func Truncate(number int, from string) int {
	digits := []int(strconv.Itoa(number))
	length := len(digits)

	if length == 1 {
		return 0
	}

	if from == "right" {
		digits = digits[:length - 1]
	} else {
		digits = digits[1:]
	}

	newNumber, _ := strconv.Atoi(string(digits))
	return newNumber
}

func IsTruncatable(number int, primes []int) bool {
	if number < 10 {
		return false
	}

	operations := []string{"right", "left"}
	for _, operation := range operations {
		copyNumber := number
		for copyNumber = Truncate(copyNumber, operation); copyNumber != 0; copyNumber = Truncate(copyNumber, operation) {
			if !IsPresent(primes, copyNumber) {
				return false
			}
		}
	}

	return true
}

func main() {
	truncatables := make([]int, 11)
	truncatablesCount := 0
	primes := BuildPrimesList(1000000)

	for _, prime := range primes {
		if IsTruncatable(prime, primes) {
			truncatables[truncatablesCount] = prime
			truncatablesCount++
		}

		if truncatablesCount > 10 {
			break
		}
	}

	sum := 0
	for _, truncatable := range truncatables {
		sum += truncatable
	}
	fmt.Println(sum)
}
