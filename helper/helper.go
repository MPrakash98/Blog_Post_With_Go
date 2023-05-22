package helper

import "fmt"

// Function for  messages
func PrintMessage(message string) {
	fmt.Println(message)
}

// Function for  errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
