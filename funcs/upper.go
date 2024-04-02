package reloaded

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
