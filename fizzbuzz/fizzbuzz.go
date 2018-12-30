// Package fizzbuzz FizzBuzzの結果をコンソールに出力
package fizzbuzz

import fmt "fmt"

// PrintFizzBuzz return void
func PrintFizzBuzz(first int, end int) {
	maxCount := end
	for index := first; index <= maxCount; index++ {

		fmt.Printf("%3d %s\n", index, ResultFizzBuzz(index))

	}
}

// ResultFizzBuzz return string
func ResultFizzBuzz(value int) string {

	if value%15 == 0 {
		return "Fizz Buzz"
	} else if value%5 == 0 {
		return "Buzz"
	} else if value%3 == 0 {
		return "Fizz"
	} else {
		return ""
	}

}
