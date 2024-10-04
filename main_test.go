package main

import "testing"

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
