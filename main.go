package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/singh-harveer/g-sign-assign/ds/heap"
	"github.com/singh-harveer/g-sign-assign/ds/trie"
)

func readLines(filename string) (string, error) {
	var line string
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return line, err
	}
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if len(line) == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}
				return line, err
			}
		}

	}
	return line, nil
}

func isValidFile(fileName string) bool {

	validExtentionSlice := []string{"txt"}
	s := strings.Split(fileName, ".")

	if len(s) < 2 {
		return false
	}
	for _, v := range validExtentionSlice {

		if s[1] == v {
			return true
		}
	}
	return false

}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	fileName := os.Args[1]

	if !isValidFile(fileName) {
		fmt.Println("invalid file!")
		return
	}

	t := trie.NewTrie()
	h := heap.NewMinheap()

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//read file line by line
		line := scanner.Text()
		words := strings.Split(line, " ")
		for i := 0; i < len(words); i++ {
			currentWord := words[i] //strings.ToLower(words[i])
			// fmt.Println(currentWord)
			result := t.Insert(currentWord)

			if result == nil {
				fmt.Println("Word is: ", words[i])
				fmt.Println("ignoring current word which contains special char/rune")
				continue
			}

			h.Insert(result, words[i])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	h.Display()

}
