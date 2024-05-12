package writer

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

var (
	testFixtures = "test-fixtures"
	testFile     = filepath.Join(".", testFixtures, "test.csv")
	tempFile     = filepath.Join(".", testFixtures, "temp.csv")
	newFile      = filepath.Join(".", testFixtures, "new-test.csv")
	noFile       = filepath.Join(".", testFixtures, "none.csv")
	header       = []string{"timestamp", "Header1", "Header2"}
)

func TestWriteCsv(t *testing.T) {
	err := WriteCsv([]string{"data1", "data2"}, tempFile, header)

	if err != nil {
		t.Errorf("Test Failed: WriteCsv failed: %v", err)
	}

	_ = os.Remove(tempFile)
}

func TestDataCsvExists(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{testFile, true},
		{noFile, false},
	}

	for _, test := range tests {
		if output := dataCsvExists(test.input); output != test.expected {
			t.Errorf("Test Failed: Expected: %v, Got: %v", test.expected, output)
		}
	}
}

func TestInitCsvFile(t *testing.T) {
	err := initCsvFile(newFile, header)

	if err != nil {
		t.Errorf("Test Failed: %v", err)
	}

	_, err = os.Stat(newFile)

	if os.IsNotExist(err) {
		t.Errorf("Test Failed: Test File Not Found")
	}

	_ = os.Remove(newFile)
}

func TestUpdateCsvFile(t *testing.T) {
	err := initCsvFile(tempFile, header)

	if err != nil {
		t.Errorf("Test Failed: %v", err)
	}

	err = updateCsvFile([]string{"data1", "data2"}, tempFile)

	if err != nil {
		t.Errorf("Test Failed: CSV file update failed: %v", err)
	}

	file, err := os.Open(tempFile)

	if err != nil {
		t.Errorf("Test Failed: Error opening edited file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if lineCount != 2 {
		t.Errorf("Test Failed: Did not get expected number of lines.")
	}
	_ = os.Remove(tempFile)
}
