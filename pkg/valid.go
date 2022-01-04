package pkg

import (
	"fmt"
	"regexp"
)

func Printer(i string) string {
	fmt.Printf(i)
	return i

}
func ValidEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if emailRegex.MatchString(email) != true {
		return false
	}
	return true
}
func NameVald(name string, min, max int) bool {
	if len(name) < min || len(name) > max {
		return false
	}
	return true
}
func PasswordValid(password string, min int) bool {
	if len(password) < min {
		return false
	}
	return true
}
