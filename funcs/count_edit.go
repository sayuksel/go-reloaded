package reloaded

func Count(text string) int { //counts the number of words in the input

	arr := []rune{}
	for _, v := range text {
		arr = append(arr, v)
	}

	x := 1
	num := 0

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] >= '0' && arr[i] <= '9' {
			num += int(arr[i]-48) * x
			x *= 10
		} else if arr[i] == '-' {
			return 0
		}
	}
	return num
}
