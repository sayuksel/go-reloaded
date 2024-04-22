package reloaded

func Valid(text string) bool { //checks which edit has to be made only if no special characters in given changes

	if text == "(hex)" || text == "(bin)" || text == "(up)" || text == "(low)" || text == "(cap)" {
		return true
	} else {
		return false
	}
}
