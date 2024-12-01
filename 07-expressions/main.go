package main

import (
	"fmt"
)

func main() {
  // variables having assigned a <type> literal
  age := 10							// this is an integer LITERAL
  name := "Jack"				// this is a string LITERAL
  rightHanded := true   // this is a boolean LITERAL

  // although assigned as literals, those variables get evalated as expressions when printed out
  // especially for the int & boolean - they each  become a string
  // -----------------------------------------------------------------------
  // AN EXPRESSION IS A VALUE (OR A SOME BIT OF CODE) THAT CAN BE EVALUATED TO A SINGLE VALUE
  // -----------------------------------------------------------------------
  fmt.Printf("%s is %d years old. Right haned: %t\n", name, age, rightHanded)  // all three vars are EXPRESSIONS here

  ageInTenYears := age + 10   // is also an EXPRESSION, as it can be evaluated to a single value
  fmt.Printf("In ten years, %s will be %d years old.\n", name, ageInTenYears)  // also EXPRESSIONS

  isATeenager := age >= 13    // also an EXPRESSION, evaluates to a single value
  fmt.Println(name, "is a teenager:", isATeenager)  // also EXPRESSIONS

  // age + 10 := 20   // not permitted, expressions on assignment must appear on the right hand side of the assignment operator
}
