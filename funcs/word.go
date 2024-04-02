package reloaded

func prev_word(end int, file string) (int, int) { // this function will give u the start and end valus of the word to be manipulated if there is only one word
	flag := false
	start := end - 3
	end -= 2
	for !flag {
		if string(file[start]) != " " {
			start--
		} else {
			flag = true
		}
	}
	start++
	return end, start
}
