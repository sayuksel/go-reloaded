package reloaded

import (
	"os"
	"strconv"
)

func To_bin(text string) string {

	decimal, err := strconv.ParseInt(text, 2, 64) //converts the binary value to decimal
	if err != nil {

		println("Error: Unable to convert the binary value to decimal.")
		println(err.Error())
		os.Exit(0)
	}
	return strconv.Itoa(int(decimal))
}
