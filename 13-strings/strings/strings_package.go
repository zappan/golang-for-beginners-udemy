package strings

import (
	"fmt"
	"strings"
	"myapp/util"
)

func StringsPackage() {
	fmt.Println("=====================================")
	util.Title("Strings Package in Go")

	courses := []string {
		"Learn Go for Beginners Crash Course",
		"Learn Java for Beginners Crash Course",
		"Learn Python for Beginners Crash Course",
		"Learn C for Beginners Crash Course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in", x, "and index is", strings.Index(x, "Go"))
		}
	}

	//            0         1         2         3         4
	//            012345678901234567890123456789012345678901234
	newString := "Go is a great programming language, Go for it!"

	fmt.Println("Variable 'newString' has prefix 'Go':", strings.HasPrefix(newString, "Go"))
	fmt.Println("Variable 'newString' has prefix 'great':", strings.HasPrefix(newString, "great"))
	fmt.Println("Variable 'newString' has prefix 'Python':", strings.HasPrefix(newString, "Python"))

	fmt.Println("Variable 'newString' has suffix '!'", strings.HasSuffix(newString, "!"))
	fmt.Println("Variable 'newString' has suffix 'it!'", strings.HasSuffix(newString, "it!"))

	fmt.Println("'Go' is found", strings.Count(newString, "Go"), "times in the variable 'newString'")

	fmt.Println("'Go' is first found at the index of", strings.Index(newString, "Go"), "in the variable 'newString'")
	fmt.Println("'Python' is first found at the index of", strings.Index(newString, "Python"), "in the variable 'newString'")

	fmt.Println("'Go' is last found at the index of", strings.LastIndex(newString, "Go"), "in the variable 'newString'")
}
