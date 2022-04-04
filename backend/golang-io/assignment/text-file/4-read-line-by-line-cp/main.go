package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Print("Hello World")
}

func ScanToArray(arr *[]string, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		*arr = append(*arr, string(line))
	}
	return nil // TODO: replace this
}

func ScanToMap(dataMap map[string]string, fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	// var deee map[string]string
	var c []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c = append(c, scanner.Text())
	}
	for i := 0; i < len(dataMap); i++ {
		dataMap[c[i]] = c[i+1]
	}

	return nil // TODO: replace this
}
