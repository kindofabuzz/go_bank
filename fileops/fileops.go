package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("Failed to find file. Creating new file with balance of 0.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse stored value.")
	}
	return value, nil
}

func WriteFloatToFile(value float64, fileName string) error {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		return errors.New("Failed to write value to file.")
	}
	return nil
}
