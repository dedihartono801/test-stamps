package main

import (
	"fmt"
	"strconv"
	"strings"
)

func defineValue(n int) string {
	switch {
	case n%3 == 0 && n%5 == 0:
		return "FooBar"
	case n%5 == 0:
		return "Bar"
	case n%3 == 0:
		return "Foo"
	default:
		return strconv.Itoa(n)
	}
}

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func collectResult(n int) []string {
	results := []string{}
	for i := n; i >= 1; i-- {
		if !isPrime(i) {
			value := defineValue(i)
			results = append(results, value)
		}
	}
	return results
}

func main() {
	// Generate prime numbers up to 100
	limit := 100
	primes := collectResult(limit)
	fmt.Println(strings.Join(primes, ", "))
}
