package reloaded

import (
	"os"
	"strconv"
	// "strings"
	// "os"
)

func To_hex(text string) string {

	decimal, err := strconv.ParseInt(text, 16, 64) //converts the hexadecimal value to decimal
	if err != nil {

		println("Error: Unable to convert the hexadecimal value to decimal.")
		println(err.Error())
		os.Exit(0)
	} 

	return strconv.Itoa(int(decimal))
}
