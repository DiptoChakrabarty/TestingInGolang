package main

import "testing"

func Test_checkPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		num      int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime no"},
		{"not prime", 8, false, "8 is not prime divisble by 2"},
		{"zero", 0, false, "0 is not prime"},
		{"one", 1, false, "1 is not prime"},
		{"two", 2, true, "2 is a prime no"},
	}

	for _, p := range primeTests {
		result, msg := checkPrime(p.num)
		if p.expected != result {
			t.Errorf("%s expected %t but got %t", p.name, p.expected, result)
		}

		if p.msg != msg {
			t.Errorf("%s expected msg %s but got %s", p.name, p.msg, msg)
		}

	}
}
