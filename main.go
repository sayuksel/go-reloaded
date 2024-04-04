package main

import (
	"fmt"
	"os"
	"strconv"
)

func Capitalize(start, end int, file string) string { // this function takes a string and capitalizes the first letter and makes other lowercase if not the first letter
	var result string
	wordStart := true
	for _, char := range file {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if wordStart {
				if char >= 'a' && char <= 'z' {
					result += string(char - 32)
				} else {
					result += string(char)
				}
				wordStart = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char + 32)
				} else {
					result += string(char)
				}
			}
		} else {
			if char >= 'A' && char <= 'Z' {
				result += string(char + 32)
			} else if char >= '0' && char <= '9' {
				result += string(char)
			} else {
				wordStart = true
				result += string(char)
			}
		}
	}
	return result
}

func to_upper(start, end int, file []byte) []byte { // this function receives the start and end point of the words and capitalizes them
	i := start
	for _, v := range file[start:end] {
		if v >= 'a' && v <= 'z' {
			file[i] = v - 32
		}
		i++
	}
	return file
}

func to_lower(start, end int, file []byte) []byte { // this function receives the start and end point of the words and makes them lower case
	i := start
	for _, v := range file[start:end] {
		if v >= 'A' && v <= 'Z' {
			file[i] = v + 32
		}
		i++
	}
	return file
}

func transfer(sample_file string, result_file *os.File) *os.File { // this function will transfer the data from sample.txt to result.txt which will be outputted
	flag := false
	for _, v := range sample_file {
		if v == '(' {
			flag = true
		}
		if !flag {
			_, err := result_file.WriteString(string(v))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			if v == ')' {
				flag = false
			}
		}
	}
	return result_file
}

func if_words(end int, file string) (int, int) { // this function will make sure if there is more than one word to be manipulated
	flag := false
	words := 1
	for _, v := range file[end:] {
		if !flag {
			if v == ')' {
				flag = true
				break
			} else {
				if v >= '1' && v <= '9' {
					words = int(v) - 48
				}
			}
		}
	}
	if words > 1 {
		r_end, start := prev_words(end, words, file)
		return r_end, start
	}
	r_end, start := prev_word(end, file)
	return r_end, start
}

func prev_words(end, words int, file string) (int, int) { // this function will give u the start and end valus of the word to be manipulated if there is more than one word
	flag := 0
	start := end - 3
	end -= 2
	for flag < words {
		if (string(file[start]) != " ") && (start > 0) {
			start--
		} else {
			flag++
			start--
		}
	}
	start++
	return end, start
}

func prev_word(end int, file string) (int, int) { // this function will give u the start and end valus of the word to be manipulated if there is only one word
	flag := false
	start := end - 3
	end -= 2
	for !flag {
		if string(file[start]) != " " {
			start--
		} else {
			flag = true
		}
	}
	start++
	return end, start
}

func if_vowel(i int, file string) bool {
	if file[i] == 'a' || file[i] == 'e' || file[i] == 'i' || file[i] == 'o' || file[i] == 'u' || file[i] == 'h' {
		return true
	}
	return false
}

func main() {
	input := os.Args[1:]
	start := 0
	end := 0
	// sampleFile, err := os.Open(input[1])
	// if err == nil {
	// 	l := len(sampleFile)
	// 	data := make([]byte,l)
	// 	sampleFile.Read(data)
	// 	fmt.Println(string(data))
	// } else {
	// 	fmt.Println(err)
	// }
	// result_data, err := os.WriteFile(input[0])
	// result_data, err := os.Open(input[1])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	sample_data, err := os.ReadFile(input[0])
	if err != nil {
		fmt.Println(err)
	} else {
		for i, v := range sample_data {
			if string(v) == "(" {
				i++
				switch string(sample_data[i]) {
				case "u":
					end, start = if_words(i, string(sample_data))
					sample_data = to_upper(start, end, sample_data)
				case "l":
					end, start = if_words(i, string(sample_data))
					sample_data = to_lower(start, end, sample_data)
					fmt.Println(string(sample_data))
				case "c":
					end, start = if_words(i, string(sample_data))
					sample_data = []byte(Capitalize(start, end, string(sample_data)))
					fmt.Println(string(sample_data))
				case "h":
					end, start = prev_word(i, string(sample_data))
					h_decimal := sample_data[start:end]
					decimal, err := strconv.ParseInt(string(h_decimal), 16, 64)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Decimal :", decimal)
					}
				case "b":
					end, start = prev_word(i, string(sample_data))
					b_decimal := sample_data[start:end]
					decimal, err := strconv.ParseInt(string(b_decimal), 2, 64)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Decimal :", decimal)
					}
				}
			}

			if v == 'a' || v == 'A' { // check for vowels
				if i > 0 { // checks if not first letter
					if sample_data[i-1] == ' ' && sample_data[i+1] == ' ' { // it makes sure that the 'a' is alone with no letter before or after it
						vowel := if_vowel(i+2, string(sample_data))
						if vowel {
							// make a to an
						}
					}
				} else {
					if sample_data[i+1] == ' ' {
						vowel := if_vowel(i+2, string(sample_data))
						if vowel {
							// make a to an
						}
					}
				}
			}
		}
		// result_data := transfer(string(sample_data), result_data)
		fmt.Println(string(sample_data))
	}
}

// end, start := prev_word(i, string(data))
// fmt.Println(string(data[start:end]))
// decimal, err := strconv.ParseInt(hexStr, 16, 64)
//     if err != nil {
//         fmt.Println("Error:", err)

/* package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sourceFileName := "source.txt"
	destinationFileName := "destination.txt"

	// Open the source file
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	// Create the destination file
	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()
		// Define your condition here. For example, transfer lines containing "important"
		if containsImportantText(line) {
			_, err := destinationFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to destination file:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading source file:", err)
	}
}

// containsImportantText checks if the given line contains the keyword "important".
// Adjust this function to match your criteria.
func containsImportantText(line string) bool {
	return strings.Contains(line, "important")
}

*/
