package main

func main() {
	printFizzBuzz(1, 30)
}

func printFizzBuzz(first int, end int) {
	maxCount := end
	for index := first; index < maxCount; index++ {

		if index%15 == 0 {
			println(index, "Fizz Buzz")
		} else if index%5 == 0 {
			println(index, "Buzz")
		} else if index%3 == 0 {
			println(index, "Fizz")
		} else {
			println(index)
		}

	}
}
