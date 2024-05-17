package isbn

import (
	"errors"
	"strconv"
	"strings"
)

// ToISBN13 converts ISBN10 to ISBN13
func ToISBN13(isbn10 string) (string, error) {
	isbn10 = strings.ReplaceAll(isbn10, "-", "")
	isbn10 = strings.ReplaceAll(isbn10, " ", "")

	if len(isbn10) != 10 {
		return "", errors.New("Invalid ISBN-10")
	}

	isbnSubstr := "978" + isbn10[:9]

	checkDigit, err := checkDigitISBN13(isbnSubstr)

	if err != nil {
		return "", err
	}

	isbn13 := isbnSubstr + strconv.Itoa(checkDigit)

	return isbn13, nil
}

// ToISBN10 converts ISBN13 to ISBN10
func ToISBN10(isbn13 string) (string, error) {
	isbn13 = strings.ReplaceAll(isbn13, "-", "")
	isbn13 = strings.ReplaceAll(isbn13, " ", "")

	if len(isbn13) != 13 {
		return "", errors.New("Invalid ISBN-13")
	}

	if !strings.HasPrefix(isbn13, "978") && !strings.HasPrefix(isbn13, "979") {
		return "", errors.New("Invalid ISBN-13 prefix")
	}

	isbnSubstr := isbn13[3:12]
	checkDigit, err := checkDigitISBN10(isbnSubstr)

	if err != nil {
		return "", err
	}

	isbn10 := isbnSubstr + strconv.Itoa(checkDigit)

	if checkDigit == 10 {
		isbn10 = isbnSubstr + "X"
	}

	return isbn10, nil
}

// checkDigitISBN13 returns the corresponding ISBN13 check digit
func checkDigitISBN13(isbnSubstr string) (int, error) {
	checkDigit := 0

	for i, c := range isbnSubstr {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, err
		}
		if i%2 == 0 {
			checkDigit += digit
		} else {
			checkDigit += digit * 3
		}
	}

	checkDigit = 10 - (checkDigit % 10)

	return checkDigit, nil
}

// checkDigitISBN10 returns the corresponding ISBN10 check digit
func checkDigitISBN10(isbnSubstr string) (int, error) {
	checkDigit := 0

	for i, c := range isbnSubstr {
		digit, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, err
		}
		checkDigit += (i + 1) * digit
	}
	checkDigit = checkDigit % 11

	return checkDigit, nil
}
