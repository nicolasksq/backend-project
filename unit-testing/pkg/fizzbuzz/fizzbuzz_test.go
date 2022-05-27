package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FizzBuzzTestSuite struct {
	suite.Suite
}

func TestFizSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzTestSuite))
}

type testsStruct struct {
	name     string
	total    int64
	fizzAt   int64
	buzzAt   int64
	expected []string
}

func TestFizzBuzzWithoutBuzz(t *testing.T) {
	fizzEverywhere := []string{"Fizz", "Fizz", "Fizz", "Fizz", "Fizz", "Fizz", "Fizz", "Fizz", "Fizz", "Fizz"}
	fizzInEvenNumbers := []string{"1", "Fizz", "3", "Fizz", "5", "Fizz", "7", "Fizz", "9", "Fizz"}
	fizzInDivisibleByThree := []string{"1", "2", "Fizz", "4", "5", "Fizz", "7", "8", "Fizz", "10"}

	testsToRun := []testsStruct{
		buildTest("fizz happy path", 10, 1, 99, fizzEverywhere),
		buildTest("fizz in even number", 10, 2, 99, fizzInEvenNumbers),
		buildTest("fizz in odd number", 10, 3, 99, fizzInDivisibleByThree),
	}

	for _, test := range testsToRun {
		t.Run(test.name, func(t *testing.T) {
			result := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFizzBuzzWithoutFizz(t *testing.T) {
	fizzEverywhere := []string{"Buzz", "Buzz", "Buzz", "Buzz", "Buzz", "Buzz", "Buzz", "Buzz", "Buzz", "Buzz"}
	fizzInEvenNumbers := []string{"1", "Buzz", "3", "Buzz", "5", "Buzz", "7", "Buzz", "9", "Buzz"}
	fizzInDivisibleByThree := []string{"1", "2", "Buzz", "4", "5", "Buzz", "7", "8", "Buzz", "10"}

	testsToRun := []testsStruct{
		buildTest("buzz happy path", 10, 99, 1, fizzEverywhere),
		buildTest("buzz in even number", 10, 99, 2, fizzInEvenNumbers),
		buildTest("buzz in odd number", 10, 99, 3, fizzInDivisibleByThree),
	}

	for _, test := range testsToRun {
		t.Run(test.name, func(t *testing.T) {
			result := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFizzBuzzWithFizzBuzz(t *testing.T) {
	fizzEverywhere := []string{"FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz"}
	fizzInEvenNumbers := []string{"1", "FizzBuzz", "3", "FizzBuzz", "5", "FizzBuzz", "7", "FizzBuzz", "9", "FizzBuzz"}
	fizzInDivisibleByThree := []string{"1", "2", "FizzBuzz", "4", "5", "FizzBuzz", "7", "8", "FizzBuzz", "10"}

	testsToRun := []testsStruct{
		buildTest("fizzbuzz happy path", 10, 1, 1, fizzEverywhere),
		buildTest("fizzbuzz in even number", 10, 2, 2, fizzInEvenNumbers),
		buildTest("fizzbuzz in odd number", 10, 3, 3, fizzInDivisibleByThree),
	}

	for _, test := range testsToRun {
		t.Run(test.name, func(t *testing.T) {
			result := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFizzBuzzWithoutFizzAndBuzz(t *testing.T) {
	fizzEverywhere := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98"}
	testsToRun := []testsStruct{
		buildTest("fizzbuzz happy path", 98, 99, 99, fizzEverywhere),
	}

	for _, test := range testsToRun {
		t.Run(test.name, func(t *testing.T) {
			result := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFizzBuzzTotalWithTotalOne(t *testing.T) {
	var empryString = []string{}
	var oneInTotal = []string{"1"}
	var fizzResult = []string{"Fizz"}
	var buzzResult = []string{"Buzz"}
	var fizzBuzzResult = []string{"FizzBuzz"}
	testsToRun := []testsStruct{
		buildTest("should get empty string", 0, 99, 99, empryString),
		buildTest("should get just one as a result", 1, 99, 99, oneInTotal),
		buildTest("should get fizz result", 1, 1, 99, fizzResult),
		buildTest("should get buzz result", 1, 99, 1, buzzResult),
		buildTest("should get fizzbuzz result", 1, 1, 1, fizzBuzzResult),
	}
	for _, test := range testsToRun {
		t.Run(test.name, func(t *testing.T) {
			result := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFizzBuzzTotalWithZeroValue(t *testing.T) {
	var expected []string
	// this should log an error but still
	testsToRun := []testsStruct{
		buildTest("should print an error", 10, 0, 0, expected),
	}
	for _, test := range testsToRun {
		t.Run(test.name, func(t *testing.T) {
			result := FizzBuzz(test.total, test.fizzAt, test.buzzAt)
			assert.Equal(t, test.expected, result)
		})
	}
}

func buildTest(name string, total, fizzAt, buzzAt int64, expected []string) testsStruct {
	return testsStruct{
		name:     name,
		total:    total,
		fizzAt:   fizzAt,
		buzzAt:   buzzAt,
		expected: expected,
	}
}
