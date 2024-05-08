package storage

import (
	"encoding/csv"
	"strconv"
	"errors"
	"os"
)

func StoreVariable(name string, value float64) error {
	file, err := os.Open("variables.csv")
	if err != nil {
		file, err = os.Create("variables.csv")
		if err != nil {
			return err
		}
	}
	defer file.Close()
	csv := csv.NewWriter(file)
	defer csv.Flush()

	strValue := strconv.FormatFloat(value, 'f', -1, 64)
	csv.Write([]string{name, strValue})
	return nil
}

func ReadVariable(name string) (float64, error) {
	file, err := os.Open("variables.csv")
	if err != nil {
		return 0, err
	}
	defer file.Close()
	csv := csv.NewReader(file)

	content, err := csv.ReadAll()
	if err != nil {
		return 0, err
	}

	for _, row := range content {
		if row[0] == name {
			return strconv.ParseFloat(row[1], 64)
		}
	}
	return 0, errors.New("Variable not found...")
}
