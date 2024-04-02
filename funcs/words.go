package reloaded

func prev_words(end, words int, file string) (int, int) { // this function will give u the start and end valus of the word to be manipulated if there is more than one word

	flag := 0
	start := end - 3
	end -= 2
	for flag < words {
		if (string(file[start]) != " ") && (start > 0) {
			start--
		} else {
			flag++
			start--
		}
	}
	start++
	return end, start
}
