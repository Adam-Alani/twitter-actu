package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Substitution struct {
	original string
	replacement string
}

func unpackSub(str, sep string) (string, string) {
	el := strings.Split(str, sep)
	return el[0], el[1]
}

func parseSub(path string) []Substitution {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []Substitution

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		original, replacement := unpackSub(scanner.Text(), ", ")
		output = append(output, Substitution{original,replacement})
	}

	return output

}