package reloaded

import (
	"fmt"
	"os"
)

func transfer(sample_file string) *os.File { // this function will transfer the data from sample.txt to result.txt which will be outputted
	var result_file *os.File

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
