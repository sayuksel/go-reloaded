package reloaded

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
