package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// Variable for the filename, which should be open and read.
var fileName string

// Total count of words in the file.
var totalCount int

// askFileName requests a .txt file name from standard input arguments,
// or scans string from standard input, or opens an available file.
func askFileName() {
	if len(os.Args) == 2 {
		fileName = os.Args[1]
	} else if len(os.Args) > 2 {
		log.Fatal("More than 1 filename")
	} else {
		fmt.Print("Please enter the filename (also with path to file) or press Enter: ")
		// For example: /home/username/Рабочий стол/t.txt
		// !! if you handle the error, a new variable 'fileName' is created !! To complete
		fileName, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		fileName = strings.TrimSpace(fileName)
		// If program ask filename and user press Enter, programm
		// will be open a file with suffix ".txt" in the same directory.
		if fileName == "" {
			files, _ := ioutil.ReadDir(".")
			for _, file := range files {
				if strings.HasSuffix(file.Name(), ".txt") {
					fileName = file.Name()
					break
				}
			}
		}
	}

	textFromFile(fileName)
}

// This function open the file and read data.
func textFromFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("err")
	}

	editWords(strings.Fields(string(data)))
}

// I plan to write this finction in the future. It will be
// delete dots and another extra symbols from words.
func editWords(words []string) {
	//for word, _ := range m {
	//	if strings.HasSuffix(word, ".") {
	//		//
	//		word = word[:strings.IndexRune(word, '.')]
	//	} else if strings.HasSuffix(word, ",") {
	//		word = word[:strings.IndexRune(word, ',')]
	//	}
	//}
	wordsCounter(words)
}

// Function counts the number of unique words in the file.
// It creates map of words and counts of them.
func wordsCounter(words []string) {
	m := make(map[string]int)

	totalCount = len(words)

	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	sortWords(m)
}

// This function sorts words by their counts and
// writes to standard output in decreasing order.
func sortWords(m map[string]int) {
	values := make([]int, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}
	sort.Ints(values)

	fmt.Printf("\nTotal count of words: %d.\n", totalCount)
	fmt.Printf("Find %d unique words from file %q:\n", len(m), fileName)

	for i := len(values) - 1; i >= 0; i-- {
		for word, count := range m {
			if values[i] == count {
				fmt.Printf("%d. %s: %d\n", i+1, word, count)
				delete(m, word)
				break
			}
		}
	}
}

// Main functiron just calls function which ask filename.
func main() {
	askFileName()
}
