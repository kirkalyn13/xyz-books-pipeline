package util

import (
	"fmt"
	"strings"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/model"
)

// FormatAuthors formats Book  authors to a single string presentation
func FormatAuthors(authors []*model.Author) string {
	var authorList []string

	for _, a := range authors {
		middleName := a.MiddleName

		if middleName == "" {
			middleName = " "
		} else {
			middleName = fmt.Sprintf(" %s ", middleName)
		}

		author := a.FirstName + middleName + a.LastName
		authorList = append(authorList, author)
	}

	return strings.Join(authorList, ", ")
}
