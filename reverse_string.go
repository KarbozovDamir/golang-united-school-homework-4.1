package reverse_string

// package main

// package main

func ReverseString(input string) (output string) {

	dataRune := []rune(input)
	strRune := make([]rune, 0, len(dataRune))
	for i := len(dataRune) - 1; i >= 0; i-- {
		strRune = append(strRune, dataRune[i])
	}
	return string(strRune)
}

// func main() {
// 	s := "Hello World!"
// 	s = ReverseString(s)
// 	fmt.Println(s)
// }
