package main

import (
	"bufio"
	"os"
)

func readTextFile(path string) ([]string, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	fileReader := bufio.NewScanner(file)
	fileReader.Split(bufio.ScanLines)
	var fileLines []string

	for fileReader.Scan() {
		fileLines = append(fileLines, fileReader.Text())
	}

	file.Close()

	return fileLines, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
