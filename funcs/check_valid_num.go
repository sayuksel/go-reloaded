package reloaded

func Valid_num(text string) bool { //checks which edit has to be made only if no special characters in given changes

	if text == "(up," || text == "(low," || text == "(cap," {
		return true
	} else {
		return false
	}
}