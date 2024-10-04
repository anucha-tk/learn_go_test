package main

import "fmt"

func main() {
	n := 105

	b, s := isPrime(n)
	fmt.Println(b, s)
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	if n < 0 {
		return false, "negative number is not prime, by definition!"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because is it divisble by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
