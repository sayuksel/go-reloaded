package main

import (
	"fmt"
	"os"
	// "strconv"
)

func capitalize(start, end int, file []byte) []byte { // this function receives the start and end point of the words and capitalizes them
	i := start
	for _, v := range file[start:end] {
		if v >= 'a' && v <= 'z' {
			file[i] = v - 32
		}
		i++
	}
	return file

}

func to_lower(start, end int, file []byte) []byte {
	i := start
	for _, v := range file[start:end] {
		if v >= 'A' && v <= 'Z' {
			file[i] = v + 32
		}
		i++
	}
	return file
}

func transfer(sample_file string, result_file *os.File) *os.File {

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

func prev_word(end int, file string) (int, int) { // this function will give u the start and end valus of the word to be manipulated
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
					sample_data = capitalize(start, end, sample_data)

				case "l":
					end, start = if_words(i, string(sample_data))
					sample_data = to_lower(start, end, sample_data)
					fmt.Println(string(sample_data))

				case "c":

				}

			}

		}
		// result_data := transfer(string(sample_data), result_data)
		// fmt.Println(result_data)
	}
}

// end, start := prev_word(i, string(data))
// fmt.Println(string(data[start:end]))
