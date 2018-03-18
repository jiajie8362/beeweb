package main

import (
	"regexp"
	"fmt"
)

const text = "My email is ccmouse@gmail.com"
func main() {
	re := regexp.MustCompile(".+@.+\\..+")
	match := re.FindString(text)
	fmt.Println(match)
}
