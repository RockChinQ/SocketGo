package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile(`^[>>*]$`)
	b := reg.MatchString("3.io>>")
	fmt.Println(b)
}
