package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := os.Args[1:]

	// sampleFile, err := os.Open(input[1])
	// if err == nil {
	// 	l := len(sampleFile)
	// 	data := make([]byte,l)
	// 	sampleFile.Read(data)
	// 	fmt.Println(string(data))
	// } else {
	// 	fmt.Println(err)
	// }

	data, err := os.ReadFile(input[1])
	if err != nil {
		fmt.Println(err)
	} else {
		for i, v := range data {
			if string(v) == "(" {
				i++
				switch data[i] {
				case "h":
					decimal, error := strconv.ParseInt(hexadecimal, 16, 64)
				}
			}

		}
		fmt.Println(string(data[:5]))
	}
}
