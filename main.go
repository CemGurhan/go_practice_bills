package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {

	capital := strings.ToUpper(n)
	names := strings.Split(capital, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func main() {

	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cem gurhan")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("barret")
	fmt.Println(fn3, sn3)

}
