package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

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
		{"neg", -5, false, "-5 is negative not prime"},
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

func Test_checkNumber(t *testing.T) {
	// We need to simulate user input in this test
	// bufioNewScanner expects io reader so we provide that

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter whole number"},
		{name: "quit", input: "q", expected: ""},
		{name: "prime", input: "5", expected: "5 is a prime no"},
		{name: "not prime", input: "15", expected: "15 is not prime divisble by 3"},
		{name: "negative", input: "-6", expected: "-6 is negative not prime"},
		{name: "string", input: "six", expected: "Please enter whole number"},
		{name: "decimal", input: "1.9", expected: "Please enter whole number"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)

		result, _ := checkNumber(reader)

		if !strings.EqualFold(result, e.expected) {
			t.Errorf("incorrect value returned for %s", e.input)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// we need a channel and an instance of io.Reader
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
