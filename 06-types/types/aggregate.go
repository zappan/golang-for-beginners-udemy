package types

import (
	"fmt"
	"myapp/util"
)

// Array type is not used as frequently in Go as in other languages
// due to existence of slices (arrays with additional functionalities)
// would slices be lists? Maybe.. TBD

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func Arrays() {
	util.Section("Arrays:")

	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("Fist element in array is", myStrings[0])
}

func Structs() {
	util.Section("Structs:")

	// var myCar Car
	// myCar.NumberOfTires = 4
	// myCar.Luxury = false
	// myCar.Make = "Volkswagen"

	// shorthand for declaring and populating a struct...
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC90",
		Year:          2019,
	}

	fmt.Printf("My car is a %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)
}
