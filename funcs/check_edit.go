package reloaded

func Edit(text, change string) string { //checks which edit has to be made only if no special characters in given changes

	switch change {
	case "(hex)":
		return To_hex(text)

	case "(bin)":
		return To_bin(text)

	case "(up)":
		return To_up(text)

	case "(low)":
		return To_low(text)

	case "(cap)":
		return To_cap(text)

	default:
		return text

	}
	// words := strings.Fields(text) //splits the input into words

	// for i := 0; i < len(words); i++ {
	// 	switch words[i] {
	// 	case "(hex)":

	// 		new_string, valid := To_hex(words[i-1])

	// 		if !valid {
	// 			return new_string, false
	// 		} else {
	// 			words[i] = " "
	// 			words[i-1] += new_string

	// 		}

	// 	}
	// }
	// return text, true
}
