package reloaded

func To_cap(text string) string { //converts the given text to uppercase

	new_string := ""

	for i, v := range text {
		if v >= 'a' && v <= 'z' {
			if i == 0 {
				new_string += string(v - 32)
			} else {
				new_string += string(v)
			}
		} else if v >= 'A' && v <= 'Z' {
			new_string += string(v)
		} else {
			return text
		}

	}
	return new_string
}
