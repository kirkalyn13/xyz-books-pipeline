package service

import (
	"testing"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/model"
)

func TestFormatAuthors(t *testing.T) {
	authors := []*model.Author{
		{FirstName: "John", LastName: "Doe"},
		{FirstName: "Jane", LastName: "Smith", MiddleName: "A."},
		{FirstName: "Alice", LastName: "Johnson", MiddleName: "B."},
	}

	expected := "John Doe, Jane A. Smith, Alice B. Johnson"
	result := formatAuthors(authors)

	if result != expected {
		t.Errorf("Test Failed: Expected: %s, Got: %s", expected, result)
	}
}
