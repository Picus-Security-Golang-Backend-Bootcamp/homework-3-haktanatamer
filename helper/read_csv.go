package helper

import (
	"encoding/csv"
	"os"
	"strconv"
)

func ReadBookCsv(filename string) ([]BookCsv, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []BookCsv

	for _, line := range lines[1:] {
		aId, _ := strconv.Atoi(line[1])

		data := BookCsv{
			Name:     line[0],
			AuthorId: aId,
		}

		result = append(result, data)
	}

	return result, nil
}

func ReadAuthorCsv(filename string) ([]AuthorCsv, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []AuthorCsv

	for _, line := range lines[1:] {
		aId, _ := strconv.Atoi(line[0])

		data := AuthorCsv{
			Name: line[1],
			ID:   aId,
		}

		result = append(result, data)
	}

	return result, nil
}
