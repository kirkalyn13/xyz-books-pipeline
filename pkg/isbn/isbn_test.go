package isbn

import (
	"testing"
)

func TestToISBN13(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0132350882", "9780132350884"},
		{"1593275846", "9781593275846"},
		{"1449325947", "9781449325947"},
		{"054792822X", "9780547928227"},
	}

	for _, test := range tests {
		isbn13, err := ToISBN13(test.input)

		if err != nil {
			t.Errorf("Test Failed: Error Occurred: %v", err)
		}
		if isbn13 != test.expected {
			t.Errorf("Test Failed: ISBN-13 conversion incorrect. Expected: %s, Got: %s", test.expected, isbn13)
		}
	}
}

func TestToISBN10(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"9780132350884", "0132350882"},
		{"9781593275846", "1593275846"},
		{"9781449325947", "1449325947"},
		{"9780547928227", "054792822X"},
	}

	for _, test := range tests {
		isbn10, err := ToISBN10(test.input)

		if err != nil {
			t.Errorf("Test Failed: Error Occurred: %v", err)
		}
		if isbn10 != test.expected {
			t.Errorf("Test Failed: ISBN-10 conversion incorrect. Expected: %s, Got: %s", test.expected, isbn10)
		}
	}
}

func TestCheckDigitISBN13(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"978013235088", 4},
		{"978159327584", 6},
		{"978144932594", 7},
		{"978054792822", 7},
	}

	for _, test := range tests {
		checkDigit, err := checkDigitISBN13(test.input)

		if err != nil {
			t.Errorf("Test Failed: Error Occurred: %v", err)
		}
		if checkDigit != test.expected {
			t.Errorf("Test Failed: Incorrect check digit. Expected: %v, Got: %v", test.expected, checkDigit)
		}
	}
}

func TestCheckDigitISBN10(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"013235088", 2},
		{"159327584", 6},
		{"144932594", 7},
		{"054792822", 10},
	}

	for _, test := range tests {
		checkDigit, err := checkDigitISBN10(test.input)

		if err != nil {
			t.Errorf("Test Failed: Error Occurred: %v", err)
		}
		if checkDigit != test.expected {
			t.Errorf("Test Failed: Incorrect check digit. Expected: %v, Got: %v", test.expected, checkDigit)
		}
	}
}
