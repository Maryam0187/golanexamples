package golangexamples

import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {

	var temp string
	for x, _ := range sliceToConcat {
		temp = temp + string(sliceToConcat[x]) + "-"
	}

	return temp
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for x, _ := range sliceToEncrypt {
		sliceToEncrypt[x] = byte(int(sliceToEncrypt[x]) + ceaserCount)
	}

}

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
