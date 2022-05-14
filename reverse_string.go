package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	// var output string
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	return output
}

// func main() {
// 	s := "Hello World!"
// 	s = ReverseString(s)
// 	fmt.Println(s)
// }
