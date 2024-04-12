package calculations

func factorial(number int64, setlogging bool) int64 {
	if number <= 1 {
		return 1
	} else {
		return number * factorial(number-1, setlogging)
	}
}