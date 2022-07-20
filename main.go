package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"tofee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps

	for k, v := range menu {

		fmt.Println(k, "-", v)

	}

	// ints as key type

	phonebook := map[int]string{

		26667: "mario",
		34589: "luigi",
		89992: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[26667])

	phonebook[34589] = "bowser"
	fmt.Println(phonebook)

	phonebook[89992] = "yoshi"
	fmt.Println(phonebook)

}
