package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a message
	fmt.Println("is it a prime?")
	fmt.Println("\n enter a whole number and press q to exit")

	// make a channel when  user wants to quit
	doneChan := make(chan bool)

	go readUserInput(doneChan)

	<-doneChan

	close(doneChan)

	fmt.Println("\nover")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		result, done := checkNumber(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(result)
	}
}

func checkNumber(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter whole number", false
	}

	_, msg := checkPrime(numToCheck)
	return msg, false
}

func checkPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime", n)
	}

	if n < 0 {
		return false, fmt.Sprintf("%d is negative not prime", n)
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not prime divisble by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime no", n)
}
