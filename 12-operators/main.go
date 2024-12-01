package main

import (
	"errors"
	"fmt"
	"math"
)

func operatorsPrecedence() {
	answer := 7 + 3 * 4  // follows the normal order of operations
	fmt.Println("Answer:", answer)

	answer = (7 + 3) * 4
	fmt.Println("Answer is now:", answer)
}


func primaryOperators() {
	// multiplication
	// area = pi*r^2
	var radius = 12.0   // to make it a float
	area := math.Pi * radius * radius
	fmt.Println("Area is:", area)

	// integer division
	half := 1 / 2
	fmt.Println("half with integer division:", half)

	// float division (32 or 64 -- must not combine different floats, compiler error)
	halfFloat := 1.0 / 2.0
	fmt.Println("half with float division:", halfFloat)

	// squaring (raising to the power)
	badThreeSquared := 3^2    // caret in Go is bitwise XOR operator (only)
	fmt.Println("badThreeSquared:", badThreeSquared)

	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("goodThreeSquared: ", goodThreeSquared)

	goodThreeSquared = math.Pow(3.0, 3.0)
	fmt.Println("goodThreeSquared is now: ", goodThreeSquared)

	// modulus
	remainder := 50 % 3
	fmt.Println("remainder:", remainder)   // 48 + 2

	// unary operators
	x := 3
	x++
	fmt.Println("x is now", x)
	x--
	x--
	fmt.Println("x is now", x)
	// --x DOES NOT EXIST IN GO
}


func precedence() {
	// Go Language Reference
	// -- The Go Programming Language Specification :: Arithmetic operator
	// GoLangPrograms.com -- simpler variation of Arithmetic operations
	// tutorials.com cheatsheet precedence

	// multiplication
	a := 12.0 * 3.0 / 4.0
	b := (12.0 * 3.0) / 4.0
	c := 12.0 * (3.0 / 4.0)
	fmt.Println("a:", a, "b:", b, "c:", c)  // all results are '9' -- parentheses are redundant

	// integer division
	unclear := 12 * (3 / 4)
	fmt.Println("unclear:", unclear)   // evaluated as int

	// parenthesis
	f := 12.0 / 3.0 / 4.0
	fmt.Println("f", f)
	f = 12.0 / (3.0 / 4.0)
	fmt.Println("f is now", f)  // paranthesis have precedence here
	fmt.Println()

	// addition/subtraction
	x := 12 + 3 - 4
	y := (12 + 3) - 4
	z := 12 + (3 - 4)
	fmt.Println("x:", x, "y:", y, "z:", z)  // all results are '11' -- parentheses are redundant
	// addition/subtraction with multiplication
	x = 12 + 3 * 4
	y = (12 + 3) * 4
	z = 12 + (3 * 4)
	fmt.Println("x:", x, "y:", y, "z:", z)  // all results are different -- parentheses do matter here
}


func modulus() {
	// does one number divide exactly into another
	x := 12
	y := 3
	if x % y == 0 {
		fmt.Println(y, "divides exactly into", x)
	} else {
		fmt.Println(y, "does not divides exactly into", x)
	}

	// what next month is -- wrong algo
	thisMonth := 11
	nextMonth := (thisMonth + 1) % 12
	fmt.Println("The month after", thisMonth, "is", nextMonth)

	// what next month is -- correct algo
	fmt.Println("-------------------------------------")
	for m := 1; m <= 12; m++ {
		nm := m % 12 + 1
		fmt.Println("The month after", m, "is", nm)
	}
	fmt.Println("-------------------------------------")

	// ...and you can use modulus in ROCK-PAPER-SCISSORS game to simplify the logic
	// see 11-flow-control/switch.go for an example of that simplifiction
}


func relationalAndConditionalOperators() {
	second := 31
	minute := 1

	// PRECEDENCE (evaluation order):
	// 1) addition operator
	// 2) less-than & greater-than
	// 3) AND logical operator
	if minute < 59 && second + 1 > 59 {
		minute++
	}
	fmt.Println(minute)
}


// short-circuit or McCarthy(?) evaluation
func shortCircuitEvaluation() {
	a := 12
	b := 6

	divideTwoNumbersUnsafe := func (x, y int) int {
		return x / y
	}

	// bad code, as it allows divide by zero if 'b' comes from some operation yielding 0
	c := divideTwoNumbersUnsafe(a, b)
	if (c == 2) {
		fmt.Println("We found a two")
	}

	// ### set an error condition
	b = 0

	// short-circuit evaluation -- will exit early
	if b != 0 && divideTwoNumbersUnsafe(a, b) == 2 {
		fmt.Println("We found a two again")
	}

	// build in the protection inside the divide function
	divideTwoNumbers := func (x, y int) (int, error) {
		if (y == 0) {
			return 0, errors.New("cannot divide by 0")
		}
		return x / y, nil
	}

	c, err := divideTwoNumbers(a, b)
	if (err != nil) {
		fmt.Println(err)
	} else {
		if (c == 2) {
			fmt.Println("We found a two once more")
		}
	}
}


func assignmentOperators() {
	x := 12
	x++
	fmt.Println(x)
	x --
	fmt.Println(x)

	y := 10
	y *= 2
	fmt.Println(y)
	y /= 4
	fmt.Println(y)

	// ### because of (unintended) side-effects, this won't work in Go (unlike in Java, where it would)
	// z := y -= 8
}


func main() {
	operatorsPrecedence()
	primaryOperators()
	precedence()
	modulus()
	relationalAndConditionalOperators()
	shortCircuitEvaluation()
	assignmentOperators()
}
