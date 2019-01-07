package unitegrate

import "strconv"

func isFizz(number int) bool {
	return number%3 == 0
}

func isBuzz(number int) bool {
	return number%5 == 0
}

func isFizzBuzz(number int) bool {
	return isFizz(number) && isBuzz(number)
}

func FizzBuzz(text string) string {
	number, _ := strconv.Atoi(text)
	if isFizzBuzz(number) {
		return "FizzBuzz"
	}

	if isFizz(number) {
		return "Fizz"
	}

	if isBuzz(number) {
		return "Buzz"
	}
	return text
}
