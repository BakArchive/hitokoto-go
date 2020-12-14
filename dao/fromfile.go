package dao

import (
	"bufio"
	"io"
	"os"
)

var result []string

func GetStringFromFile(source string) ([]string, error) {
	file, err := os.Open(source)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	if result == nil {
		result = make([]string, 0)
		reader := bufio.NewReader(file)
		for {
			str, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			result = append(result, string(str))
		}
	}
	return result, nil
}
