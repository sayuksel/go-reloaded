package reloaded

func Apos(words []string) []string {

	flag := true
	for i, v := range words {
		if (v[0] == 34 || v[0] == 39) && len(v) > 1 { //if the apostrophe is stuck to the first word
			flag = false
		}

		if v == string(rune(34)) || v == string(rune(39)) {
			if flag { //if the apostrophe isnt stuck to the first word
				words[i+1] = v + words[i+1]
				words = append(words[:i], words[i+1:]...)
				flag = false
			} else {
				words[i-1] = words[i-1] + v
				if i == len(words)-1 {
					return words[:i]
				}
				words = append(words[:i], words[i+1:]...)
				flag = true
			}
		}
	}
	return words
}
