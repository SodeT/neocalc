package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var (
	pl = fmt.Println
)

// TODO: Implement these functions again so that they actually work...
func SaveVariable(identifier string, value float64) error {
	file, err := os.OpenFile("variables.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	csvFile := csv.NewWriter(file)
	defer csvFile.Flush()

	strValue := strconv.FormatFloat(value, 'f', -1, 64)
	err = csvFile.Write([]string{identifier, strValue})
	if err != nil {
		return err
	}
	return nil
}

func SaveFunction(identifier string, body string) error {
	file, err := os.OpenFile("functions.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	csvFile := csv.NewWriter(file)
	defer csvFile.Flush()

	err = csvFile.Write([]string{identifier, body})
	if err != nil {
		return err
	}
	return nil
}

func LoadFunctions() ([]string, error) {
	file, err := os.Open("functions.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvFile := csv.NewReader(file)

	content, err := csvFile.ReadAll()
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, row := range content {
		ret = append(ret, row[1])
	}
	return ret, nil
}

func LoadVariables() (map[string]float64, error) {
	file, err := os.Open("variables.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	csvFile := csv.NewReader(file)

	content, err := csvFile.ReadAll()
	if err != nil {
		return nil, err
	}

	ret := make(map[string]float64)
	for _, row := range content {
		ret[row[0]], err = strconv.ParseFloat(row[1], 64)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}


