package strings

import (
	"fmt"
	"strings"
	"myapp/util"
)

func StringsManipulation() {

	// =====================================================================

	fmt.Println("=====================================")
	util.Title("Strings Manipulation using 'strings' Package in Go")

	//            0         1         2         3         4
	//            012345678901234567890123456789012345678901234
	newString := "Go is a great programming language, Go for it!"

	if strings.Contains(newString, "Go") {
		var replacedString string

		replacedString = strings.Replace(newString, "Go", "Golang", -1)
		fmt.Println(replacedString)

		replacedString = strings.Replace(newString, "Go", "Golang", 1)
		fmt.Println(replacedString)

		replacedString = strings.ReplaceAll(newString, "Go", "Golang")
		fmt.Println(replacedString)
	}

	// =====================================================================

	util.Title("String comparison:")
	a := "A"
	if a == "A" {
		fmt.Println("a is equal to A")
	}

	if "A" > "B" {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("A is NOT greater than B")
	}

	if "Alpha" > "Absolute" {
		fmt.Println("Alpha is greater than Absolute")
	} else {
		fmt.Println("Alpha is NOT greater than Absolute")
	}

	// =====================================================================

	util.Title("String trimming:")
	badEmail := " me@here.com  "
	goodEmail := strings.TrimSpace(badEmail)
	fmt.Printf("badEmail: '%s'\n", badEmail)
	fmt.Printf("goodEmail: '%s'\n", goodEmail)

	// =====================================================================

	util.Title("String replace at position:")
	str := "alpha alpha alpha alpha alpha"
  str = replaceNth(str, "alpha", "beta", 3)
  fmt.Println(str)

	// =====================================================================

  util.Title("String casing:")
  myString := "This is a clear EXAMPLE of why we search in one case only."

  if strings.Contains(myString, "this") {
  	fmt.Println("Searching for 'this'.... Found it!")
  } else {
  	fmt.Println("'this' not found")
  }

  if strings.Contains(strings.ToLower(myString), "this") {
  	fmt.Println("Searching for 'this'.... Found it!")
  } else {
  	fmt.Println("'this' not found")
  }

  fmt.Println(strings.ToLower(myString))
  fmt.Println(strings.ToUpper(myString))
  fmt.Println(strings.Title(myString))
  // the final line has all Title Case
  fmt.Println(strings.Title(strings.ToLower(myString)))

}



func replaceNth(s, old, new string, n int) string {
	replacementIdx := 0	// index
	lenOld := len(old)

	for j := 1; j <= n; j++ {
		x := strings.Index(s[replacementIdx:], old)
		if (x < 0) {
			break // we did not find the old in s
		}

		// we have found the old in s... continue logic

		replacementIdx = replacementIdx + x
		if j == n {			// start of chars we want to replace
			remainderIdx := replacementIdx + lenOld
			return s[:replacementIdx] + new + s[remainderIdx:]
		}

		replacementIdx += lenOld
	}

	return s
}
