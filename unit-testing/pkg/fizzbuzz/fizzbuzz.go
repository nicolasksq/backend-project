package fizzbuzz

import (
	"log"
	"strconv"
)

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if is is not divisible by either the `fizzAt` or
//   the `buzzAt` values.

// FizzBuzz TODO (done) added a recover function in case of panic -> divide by zero
// another options:
//	- check if the number is 0 and then return a <nil, error>
//	- check if the number is 0 and then return <[]string>

func FizzBuzz(total, fizzAt, buzzAt int64) []string {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Printf("hey, we have a problem here: %s", e.Error())
		}
	}()
	result := make([]string, total)
	for i := int64(1); i <= total; i++ {
		if !(i%fizzAt == 0) && !(i%buzzAt == 0) {
			result[i-1] = strconv.FormatInt(i, 10)
			continue
		}
		if i%fizzAt == 0 {
			result[i-1] = "Fizz"
		}
		if i%buzzAt == 0 {
			result[i-1] += "Buzz"
		}
	}
	return result
}
