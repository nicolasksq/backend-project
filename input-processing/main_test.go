package main

// quick test over main - o.O

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"
)

func TestInputProcessingFixedFile(t *testing.T) {
	file, err := os.Open("test_fixed.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	searchErrorsAndCreateFile(bufio.NewReader(file), binarySearch)
}

func TestInputProcessingRandomFile(t *testing.T) {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file")
		return
	}

	writer := bufio.NewWriter(f)
	for i := 0; i < 1000; i++ {
		_, err = writer.WriteString(largestRandomString(100, false))
		if err != nil {
			fmt.Println("Error writing to buffer")
			return
		}
		if i%7 == 0 {
			_, err = writer.WriteString(largestRandomString(100, true))
			if err != nil {
				fmt.Println("Error writing to buffer")
				return
			}
		}
	}

	writer.Flush()
	f.Close()

	file, _ := os.Open("test.txt")
	searchErrorsAndCreateFile(bufio.NewReader(file), binarySearch)
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func largestRandomString(times int, withError bool) string {
	var result string
	for i := 0; i <= times; i++ {
		result += " " + randomString(math.MaxInt8) + " \n"
	}
	if withError {
		result += " ERROR"
	}
	return result
}
