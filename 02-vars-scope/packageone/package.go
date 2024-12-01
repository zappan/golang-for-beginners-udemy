package packageone

import "fmt"

// camepCase - private variable
var privateVar = "I am private"

// PascalCase - public variable
var PublicVar = "I am public (or exported)"


func notExported() {
	// private function - expoted outside of this package
	fmt.Println("From packageone -- I'm a private function")
}

func Exported() {
	// public function - expoted outside of this package
	fmt.Println("From packageone -- I'm a public function")
}
