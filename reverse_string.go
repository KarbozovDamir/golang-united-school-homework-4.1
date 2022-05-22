package reverse_string

// package main

func ReverseString(input string) (output string) {

	strBytes := make([]byte, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		strBytes = append(strBytes, input[i])
	}
	return string(strBytes)
}

// func main() {
// 	s := "Hello World!"
// 	s = ReverseString(s)
// 	fmt.Println(s)
// }
