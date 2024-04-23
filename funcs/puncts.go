package reloaded

import (
	"fmt"
	"os"
)

func Puncts(s []string) []string {
	pun := []string{".", ",", "!", "?", ":", ";"}
	flag := true
	for i, v := range s {
		for _, w := range pun {
			if len(v) > 0 {
				if string(v[0]) == w { // if there is a punc in the first word
					for _, x := range pun {
						if string(v[len(v)-1]) == x { // if there is a punc in the last word
							if i == 0 {
								fmt.Println("Error: There seems to be an error with the input text.")
								os.Exit(0)
							}
							s[i-1] = s[i-1] + v
							if i == len(s)-1 {
								s = s[:i]
								return s
							} else {
								s = append(s[:i], s[i+1:]...)
								flag = false
							}
						}
					}
					if flag {
						if i < len(s) {
							s[i-1] = s[i-1] + w
							s[i] = v[1:]

						}
					}
					flag = true
				}
			}
		}
	}
	return s
}
