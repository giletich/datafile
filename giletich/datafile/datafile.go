package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(GetName string) ([]float64, error) {
	num := make([]float64, 3)
	file, err := os.Open(GetName)
	if err != nil {
		return nil, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	number := 0.0
	for scanner.Scan() {
		number, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		num = append(num, number)
		i++
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return num, nil
}

func GetStrings(GetName string) ([]string, error) {
	num := make([]string, 3)
	file, err := os.Open(GetName)
	if err != nil {
		return nil, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    number := scanner.Text()
		num = append(num, number)
		i++
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return num, nil
}
