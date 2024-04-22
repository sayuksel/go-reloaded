package reloaded

func To_low(text string) string { //converts the given text to uppercase

	new_string := ""

	for _, v := range text {
		if v >= 'A' && v <= 'Z' {
			new_string += string(v + 32)
		} else if v >= 'a' && v <= 'z' {
			new_string += string(v)
		} else {
			return text
		}

	}
	return new_string
}
