package main

import "fmt"

func main() {

	myBill := newBill("marios bill")

	myBill.updateTip(10)

	fmt.Println(myBill.format())

}
