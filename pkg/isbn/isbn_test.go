package isbn

import "testing"

type (
	ConvertResult struct {
		input    string
		expected string
	}
)

func TestConvertISBN10ToISBN13(t *testing.T) {
	tests := []ConvertResult{
		{"0596000480", "9780596000484"},
		{"0321751043", "9780321751041"},
		{"0131103628", "9780131103627"},
	}

	for _, test := range tests {
		result, err := ConvertISBN10ToISBN13(test.input)
		if err != nil {
			t.Errorf("Error converting ISBN-10 to ISBN-13: %v", err)
		}

		if result != test.expected {
			t.Errorf("ISBN-10 to ISBN-13 conversion failed for input %s. Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}

func TestConvertISBN13ToISBN10(t *testing.T) {
	tests := []ConvertResult{
		{"9780596000484", "0596000480"},
		{"9780321751041", "0321751043"},
		{"9780131103627", "0131103628"},
	}

	for _, test := range tests {
		result, err := ConvertISBN13ToISBN10(test.input)
		if err != nil {
			t.Errorf("Error converting ISBN-13 to ISBN-10: %v", err)
		}

		if result != test.expected {
			t.Errorf("ISBN-13 to ISBN-10 conversion failed for input %s. Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
