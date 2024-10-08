package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	intro()

	doneChan := make(chan bool)
	go readUserInput(os.Stdin, doneChan)
	<-doneChan
	close(doneChan)
	fmt.Println("Goodbye")
}

func intro() {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumber(scanner)
		if done {
			doneChan <- true
		}
		fmt.Println(res)
		prompt()
	}
}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "quit", true
	}

	numtoCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}
	_, msg := isPrime(numtoCheck)

	return msg, false
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
