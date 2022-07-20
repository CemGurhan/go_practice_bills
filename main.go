package main

import "fmt"

func main() {

	age := 23
	name := "Cem"

	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	//println

	fmt.Println("115")
	fmt.Println("511")
	fmt.Println("My name is", name, "and my age is", age)

	//printf

	fmt.Printf("My name is %v and my age is %v \n", name, age)
	fmt.Printf("My name is %q and my age is %q \n", name, age)
	fmt.Printf("age is of type %T", age)
	fmt.Printf("you scored %f points \n", 225.33)
	fmt.Printf("you scored %0.2f points \n", 225.33)

	// sprintf
	str := fmt.Sprintf("My name is %v and my age is %v \n", name, age)
	fmt.Println("the saved string is:", str)

}
