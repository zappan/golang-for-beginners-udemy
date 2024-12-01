package main

import (
	"fmt"
	"myapp/packageone"
)

// package-level variable
var one = "One"

func main() {
	// variable shadowing -- block-level variable
	// var one = "this is a block level variable one"
	fmt.Println(one)

	myFunc()

	// when PasalCased - it's a publi variable, available outside of that package
	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	// ### when camelCased - it's a private variable, not available outside of that package
	// secondString := packageone.privateVar
	// fmt.Println("From packageone:", secondString)

	// public function called
	packageone.Exported()
}


func myFunc() {
	// var one = "the nmber one"
	fmt.Println(one)
}
