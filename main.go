package main

import (
	"fmt"
	"os"
	"reloaded"
	"strconv"
)

func First_letter(letter rune, i int, file string) bool {
	if file[i+1] != ' ' && file[i-1] != ' ' {
		return false
	} else {
		return true
	}
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
			if string(v) == "(" { // for hex, bin, up, low and cap
				i++
				switch string(sample_data[i]) {
				case "u":
					end, start = reloaded.if_words(i, string(sample_data))
					sample_data = reloaded.to_upper(start, end, sample_data)

				case "l":
					end, start = reloaded.if_words(i, string(sample_data))
					sample_data = reloaded.to_lower(start, end, sample_data)
					fmt.Println(string(sample_data))

				case "c":
					end, start = reloaded.if_words(i, string(sample_data))
					sample_data = []byte(reloaded.Capitalize(start, end, string(sample_data)))
					fmt.Println(string(sample_data))

				case "h":
					end, start = reloaded.if_words(prev_word, i, string(sample_data))
					h_decimal := sample_data[start:end]
					decimal, err := strconv.ParseInt(string(h_decimal), 16, 64)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Decimal :", decimal)
					}

				case "b":
					end, start = reloaded.prev_word(i, string(sample_data))
					b_decimal := sample_data[start:end]
					decimal, err := strconv.ParseInt(string(b_decimal), 2, 64)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Decimal :", decimal)
					}
				}
			}
			// if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			// first := First_letter(rune(v), i, string(sample_data))
			// }
		}
		result_data := reloaded.transfer(string(sample_data))
		fmt.Println(result_data)
	}
}

// end, start := prev_word(i, string(data))
// fmt.Println(string(data[start:end]))

// decimal, err := strconv.ParseInt(hexStr, 16, 64)
//     if err != nil {
//         fmt.Println("Error:", err)
