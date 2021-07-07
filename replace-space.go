package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.Replace(s, " ", `%20`, -1)
}

func main() {
	fmt.Println(replaceSpace("we are happy"))
}