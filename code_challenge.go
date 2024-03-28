package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isValidCreditCard(number string) bool {
	pattern := `^(?:4|5|6)(?:\d{15}|\d{3}(-\d{4}){3})$`
	matched, _ := regexp.MatchString(pattern, number)
	if !matched {
		return false
	}

	// Check for four or more consecutive repeated digits
	normalizedNumber := strings.ReplaceAll(number, "-", "")
	for i := 0; i < len(normalizedNumber)-3; i++ {
		if normalizedNumber[i] == normalizedNumber[i+1] &&
			normalizedNumber[i+1] == normalizedNumber[i+2] &&
			normalizedNumber[i+2] == normalizedNumber[i+3] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i < n; i++ {
		scanner.Scan()
		creditCardNumber := scanner.Text()
		// fmt.Println(creditCardNumber)
		if isValidCreditCard(creditCardNumber) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
