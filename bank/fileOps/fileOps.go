package fileOps

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value int64) {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		return
	}
}

func GetFloatFromFile(fileName string) int64 {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0
	}
	valueText := string(data)
	value, convError := strconv.ParseInt(valueText, 64, 10)
	if convError != nil {
		return 0
	}

	return value
}
