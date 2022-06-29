package action

import "fmt"

// Function for handling messages
func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
