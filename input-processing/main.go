package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var lineWithErr = make(chan string)
var endChan = make(chan struct{})

const (
	binarySearch = "binary"
	linearSearch = "linear"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()
	r := bufio.NewReader(os.Stdin)
	searchErrorsAndCreateFile(r, binarySearch)
}

func searchErrorsAndCreateFile(r *bufio.Reader, searchType string) {
	wordToSearch := "ERROR"
	go func() {
		for {
			l, err := r.ReadString('\n')
			line := strings.TrimSuffix(l, "\n")

			// TODO tech-debt. we should have a better way to solve this before and add more things.
			if err != nil {
				endChan <- struct{}{}
				break
			}

			switch searchType {
			case linearSearch:
				if strings.Contains(line, wordToSearch) {
					lineWithErr <- line
				}
			case binarySearch:
				// create an array of words
				words := strings.Split(line, " ")
				// binary search -> O(log n) comparisons, where n is the size of the slice.
				// sort array of words
				sort.Strings(words)
				if words[sort.SearchStrings(words, wordToSearch)] == wordToSearch {
					lineWithErr <- line
				}
			default:
				endChan <- struct{}{}
				break
			}
		}
	}()

	for {
		select {
		case result := <-lineWithErr:
			f, err := os.OpenFile("./errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("error opening file")
				return
			}

			writer := bufio.NewWriter(f)
			_, err = writer.WriteString(result + "\n")
			if err != nil {
				fmt.Println("Error writing to buffer")
				return
			}
			writer.Flush()
			f.Close()
		case <-endChan:
			// we finish our process
			fmt.Println("exiting")
			return
		}
	}
}
