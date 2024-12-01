package strings

import (
	"fmt"
	"strings"
	"myapp/util"
)

const TAB = "\t"

func BasicStrings() {

	// ========================================================================
	util.Title("String:")
	name := "Hello world"
	fmt.Println(name)


	// ========================================================================
	util.Title("Bytes:")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()

	// ========================================================================
	util.Title("Index\t Hex\t Rune\t String")
	for x, y := range name {
		fmt.Printf("%3d\t %3x\t %4d\t %3s\n", x, y, y, string(y))
	}

	cyrillicName := "Здраво свете"
	util.Title("Index\t Hex\t Rune\t String")
	for x, y := range cyrillicName {
		fmt.Printf("%3d\t %3x\t %4d\t %3s\n", x, y, y, string(y))
	}


	// ========================================================================
	util.Title("Three ways to concatenate strings:")
	h := "Hello, "
	w := "world."

	// using plus-sign -- not the most efficient way, as strings are immutable
	// (lecture comment, not sure what it has to do with a plus-sign;
	// I guess in case of multiple plus-signs on the single line it yields intermediary results per each pair
	myString := h + w
	fmt.Println(myString)

	// using fmt -- [far] more efficient way
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// using stringbuilder - very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())


	// ========================================================================
	util.Title("Getting a substring:")
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(name[0:13])
	fmt.Println(name[13:])
	fmt.Println(name[6:21])


	// ========================================================================
	util.Title("String indexing:")
	//             0         1         2         3
	//             01234567890123456789012345678901234
	courseName := "Learn Go for Beginners Crash Course"

	fmt.Println(courseName[0])  // will not print 'L' but its int value (Rune)
	fmt.Println(string(courseName[0]))  // will now print 'L'
	fmt.Println(string(courseName[6]))  // will now print 'G'

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	for i := 13; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()


	// ========================================================================
	util.Title("String length:")
	fmt.Println("Lenght of courseName is:", len(courseName))


	// ========================================================================
	util.Title("Slices length:")
	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")
	fmt.Println("mySlice has", len(mySlice), "elements")
	// fmt.Println("the last element in mySlice is", mySlice[-1]) -- does not work
	fmt.Println("the last element in mySlice is", mySlice[len(mySlice)-1])
}
