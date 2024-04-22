package main

import (
	"os"
	"strings"

	reloaded "reloaded/funcs"
)

func main() {
	arguments := os.Args // takes the arguments from the command line

	if len(arguments) != 3 {
		println("Error: Invalid number of arguments provided.")
		println("Usage: go run main.go <arg1> <arg2>")
		return
	}

	input := arguments[1]
	output := arguments[2]

	text, err := os.ReadFile(input) // reads the input file
	if err != nil {
		println("Error: Unable to read the input file.")
		println(err.Error())
		return
	}

	empty_ret, empty_val := reloaded.Empty(string(text)) // checks if the input is empty or not

	if empty_ret { // gives results based on the return value of check_empty
		if empty_val == 1 {
			println("Error: No input provided.")
		} else if empty_val == 2 {
			println("Error: Only whitespace provided.")
		}
		os.Exit(0)
	}

	words := strings.Split(string(text), " ") // splits the input into words
	reloaded.Space(words)

	for i, v := range words {
		valid := reloaded.Valid(string(v))
		if valid {
			if i == 0 {
				println("Error: There seems to be no text to be edited (beginning of text file).")
				os.Exit(0)
			} else {
				words[i-1] = reloaded.Edit(words[i-1], v)

				if i < len(words)-1 {
					words = append(words[:i], words[i+1:]...)
				} else {
					words = words[:i]
					break
				}
			}
		}
	}

	for i, v := range words {
		valid := reloaded.Valid_num(string(v))
		if valid {
			count := reloaded.Count(words[i+1]) // counts the number of words to be edited
			if i == 0 {
				println("Error: There seems to be no text to be edited (beginning of text file).")
				os.Exit(0)

			} else if count > i {
				println("Error: There are not enough words to edit.")
				os.Exit(0)
			} else if count == 0 {
				println("Error: Negative input not allowed.")
				os.Exit(0)
			}
			for j := count; j > 0; j-- {
				v = strings.ReplaceAll(v, ",", ")")
				words[i-j] = reloaded.Edit(words[i-j], v)
			}
			if i < len(words)-2 {
				words = append(words[:i], words[i+2:]...)
			} else {
				words = words[:i]
				break

			}
		}
	}

	reloaded.Vowels(words) // corrects the vowels in the words

	words = reloaded.Puncts(words) // corrects the punctuation in the words

	words = reloaded.Apos(words)                     // corrects the apostrophes in the words
	output_text := strings.Join(words, " ")          // joins the words back into a single string
	os.WriteFile(output, []byte(output_text), 0o644) // writes the output to the output file
}
