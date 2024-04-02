package reloaded

func if_words(end int, file string) (int, int) { // this function will make sure if there is more than one word to be manipulated
	flag := false
	words := 1

	for _, v := range file[end:] {
		if !flag {
			if v == ')' {
				flag = true
				break
			} else {
				if v >= '1' && v <= '9' {
					words = int(v) - 48
				}
			}
		}
	}
	if words > 1 {
		r_end, start := prev_words(end, words, file)
		return r_end, start
	}

	r_end, start := prev_word(end, file)
	return r_end, start
}
