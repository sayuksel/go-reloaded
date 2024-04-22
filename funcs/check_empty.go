package reloaded

import "strings"

func Empty(text string) (bool, int) {

	if len(text) == 0 {
		return true, 1
	} else if strings.TrimSpace(text) == "" {
		return true, 2
	}
	return false, 0
}