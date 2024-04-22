package reloaded

import (
	"os"
)

func Space(words []string) {

	for _, v := range words {
		if v == "" {
			println("Error: Too many spaces provided.")
			os.Exit(0)
		}
	}
}
