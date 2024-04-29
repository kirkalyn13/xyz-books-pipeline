package isbn

import (
	"fmt"
	"math"
	"strconv"
)

// ConvertISBN10ToISBN13 converts an ISBN-10 string to an ISBN-13 string.
func ConvertISBN10ToISBN13(isbn10 string) (string, error) {
	if len(isbn10) != 10 {
		return "", fmt.Errorf("invalid ISBN-10 length")
	}

	isbn10Prefix := isbn10[:9]
	isbn13Prefix := "978" + isbn10Prefix

	checkDigit, err := getCheckDigit(isbn13Prefix, true)
	if err != nil {
		return "", err
	}

	return isbn13Prefix + checkDigit, nil
}

// ConvertISBN13ToISBN10 converts an ISBN-13 string to an ISBN-10 string.
func ConvertISBN13ToISBN10(isbn13 string) (string, error) {
	if len(isbn13) != 13 {
		return "", fmt.Errorf("invalid ISBN-13 length")
	}

	isbn10Prefix := isbn13[3:12]

	checkDigit, err := getCheckDigit(isbn10Prefix, false)
	if err != nil {
		return "", err
	}

	return isbn10Prefix + checkDigit, nil
}

// getCheckDigit calculates the check digit for an ISBN-13 or ISBN-10 prefix.
func getCheckDigit(prefix string, isISBN13 bool) (string, error) {
	var sum int
	var modVal int
	var constantVal int

	for i, digit := range prefix {
		digitValue, err := strconv.Atoi(string(digit))

		if err != nil {
			return "", fmt.Errorf("invalid digit in ISBN: %v", err)
		}

		sum += digitValue * (i%2*2 + 1)
	}

	if isISBN13 {
		constantVal = 10
		modVal = 10
	} else {
		constantVal = 0
		modVal = 11
	}

	checkDigit := math.Abs(float64((constantVal - (sum % modVal)) % modVal))
	return strconv.Itoa(int(checkDigit)), nil
}
