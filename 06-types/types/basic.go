package types

import (
  "fmt"
  "myapp/util"
)

var myInt int     // encouraged to USE ONLY THIS
var myInt16 int16 // encouraged NOT to use
var myInt32 int32 // encouraged NOT to use -- slightly more efficient if targeting ONLY 32-bit architecture
var myInt64 int64 // encouraged NOT to use -- slightly more efficient if targeting ONLY 64-bit architecture
var myUint uint   // unsigned

var myFloat float32   // large numbers, default use
var myFloat64 float64 // really large numbers, if you know you'll need it

func Numerics() {
	util.Section("Numerics:")

	myInt = 10
	myUint = 20
	myFloat = 10.1
	myFloat64 = 100.1

	fmt.Println(myInt, myUint, myFloat, myFloat64)
}

// STRINGS ARE IMMUTABLE IN GO
func Strings() {
	util.Section("Strings:")

	myString := "Trevor"
	fmt.Println(myString)

	myString = "Tom" // IMMUTABLE => this will create an entirely new string and store into your variable
	fmt.Println(myString)
}

func Booleans() {
	util.Section("Booleans:")

	var myBool = true
	fmt.Println(myBool)

	myBool = false
	fmt.Println(myBool)
}
