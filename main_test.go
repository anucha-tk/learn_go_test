package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_IsPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		msg      string
		testNum  int
		expected bool
	}{
		{"not prime", "1 is not prime, by definition!", 1, false},
		{"prime", "7 is a prime number!", 7, true},
		{"negative", "negative number is not prime, by definition!", -10, false},
		{"not prime", "8 is not a prime number because is it divisble by 2", 8, false},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if msg != e.msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_promt(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	prompt()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Error("incorrect prompt: expected -> ")
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Error("incorrect intro msg")
	}
}

func Test_checkNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"quit", "q", "quit"},
		{"not number", "w", "Please enter a whole number"},
		{"prime", "7", "7 is a prime number!"},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumber(reader)
		if !strings.EqualFold(res, test.expected) {
			t.Errorf("%s incorrect expected: 7 is a prime number!", test.name)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)

	<-doneChan

	close(doneChan)
}
