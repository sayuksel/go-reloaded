package reloaded

func Vowels(words []string) {

	vowels := "aeiouAEIOU"

	for i, v := range words {
		if v == "a" || v == "A" {
			for _, w := range vowels {
				if w == rune(words[i+1][0]) {
					words[i] = words[i] + "n"

				}
			}

		}

	}

}
