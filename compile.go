package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".js") {
			fmt.Println(name)
			compile(name)
		}
	}
}

func compile(name string) {
	input, err := readLines(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := ""

	for _, line := range input {
		clear := strings.TrimSpace(line)
		clear = strings.ReplaceAll(clear, "\"", "%27")

		if !strings.HasPrefix(clear, "//") {
			result = result + " " + clear
		}
	}

	slice := []string{result}

	writeLines(slice, name+".bm")

}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
