// Storing here the functions that will handle the errors
package mathskills

import "fmt"

func ErrorHandler(s string) {
	fmt.Println(s)
}

func CheckError(err error) bool {
	return err != nil
}
