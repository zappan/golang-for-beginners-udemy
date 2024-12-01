package types

import (
	"fmt"
	"sort"
	"myapp/util"
	"github.com/eiannone/keyboard"
)


// ====================================================================


// num is more numAddr, actually
func changeValueOfAPointer(num *int, value int) {
	*num = value
}

func Pointers() {
	util.Section("Pointers:")
	// // direct store of the information
	// var myInt int
	// myInt = 10
	// fmt.Println(myInt)

	// also possible to declare a variable that doesn't hold the information
	// but point to the location in memory that holds the infomation

	// with integers
	x := 10
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer address is", myFirstPointer) 	// prints out the memory address
	fmt.Println("myFirstPointer is", *myFirstPointer)         // prints out the value at the given memory address
	fmt.Println("x addresss is", &x)
	util.Spacer()

	*myFirstPointer = 15

	fmt.Println("x is now", x)
	fmt.Println("myFirstPointer address is now", myFirstPointer)  // prints out the value at the given memory address
	fmt.Println("x addresss is now", &x)
	fmt.Println("myFirstPointer is now", *myFirstPointer)         // prints out the value at the given memory address
	util.Spacer()

	changeValueOfAPointer(myFirstPointer, 919)
	fmt.Println("After function call, x is now", x)
	changeValueOfAPointer(&x, 323)
	fmt.Println("After function call, x is now", x)

	fmt.Println("----------------")

	// with strings -- note the new memory location after change NOT happening
	// although string is immutable and created entirely new...
	// interesting it takes (can take?) the same memory location

	pet := "I own a dog"
	fmt.Println("pet is:", pet)
	fmt.Println("pet address is:", &pet)
	util.Spacer()

	pet = "I own a cat and a fish and a squirrel and a bat and a turtle"
	fmt.Println("pet is now:", pet)
	fmt.Println("pet address is now:", &pet)
}


// ====================================================================


func deleteFromSlice(slice []string, i int) []string {
	slice[i] = slice[len(slice) - 1]  // 1) copy the last element into the position 'i'
	slice[len(slice) - 1] = ""        // 2) set that last element to nothing (why not nil?)
	slice = slice[:len(slice) - 1]	  // 3) trucnate the slice by removing that last el. (by taking the portion of the slice up to the last el.)
	return slice
}

// ? slices keep the order of elements ? (do they ? if so, they're lists)
func Slices() {
	util.Section("Slices:")

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	// for index, el := range animals {..}
	// for _, x := range animals {..}
	for idx, x := range animals {
		// fmt.Println(idx, x)
		fmt.Println(x, "@ idx", idx)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First to elements are", animals[0:2])
	fmt.Println("The 'animals' slice is", len(animals), "elements long.")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)
	util.Spacer()

	pos := 1
	fmt.Println("Deleting from slice", animals, "at position", pos)
	animals = deleteFromSlice(animals, pos)
	fmt.Println(animals)
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)
}


// ====================================================================


// this function should somehow support generics or 'any' tyoe for elVal...
func isElFoundOutput(elVal int, elExists bool) {
	if (elExists) {
		fmt.Println("element", elVal, "found in map")
	} else {
		fmt.Println("element", elVal, "NOT found in map")
	}
}

// MAP === KEY-VALUE PAIRS === DICTIONARY === JSON
// a reference type, so it's passed by reference by default
// rarely if ever a need to use pointers to it
func Maps() {
	util.Section("Maps:")


	// var x = map[int]int   				// creates a nil-map, can't do much with it
	intMap := make(map[string]int)  // need to use 'make()'' keyword instead

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	// map never sorted, run this multiple times, each time they're in different order
	for key, value := range intMap {
		fmt.Println(key, value)
	}
	util.Spacer()

	delete(intMap, "four")
	for key, value := range intMap {
		fmt.Println(key, value)
	}
	util.Spacer()

	el, elExists := intMap["four"]
	isElFoundOutput(el, elExists)

	el, elExists = intMap["two"]
	isElFoundOutput(el, elExists)

	intMap["two"] = 4
	el, elExists = intMap["two"]
	isElFoundOutput(el, elExists)
}


// ====================================================================


func addTwoNumbers(x, y int) int {
	return x + y
}

// "Naked Return" - almost never used in practice
// -- only use for very short functions, as it makes readability worse
func addTwoNumbersNaked(x, y int) (sum int) {
	sum = x + y
	return
}

// "a variatic function" (variatic param always comes last)
func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total += x
	}
	return total
}

type Animal struct {
	Name string
	Sound string
	NumLegs int
}

// "a receiver function"
func (a *Animal) says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) legs() {
	fmt.Printf("A %s has %d legs\n", a.Name, a.NumLegs)
}

func Functions() {
	util.Section("Functions:")

	z := addTwoNumbers(2, 4)
	fmt.Println("value of z is", z)
	// util.Spacer()

	zNaked := addTwoNumbersNaked(3, 8)
	fmt.Println("value of zNaked is", zNaked)
	// util.Spacer()

	myTotal := sumMany(1,2,4,6)
	fmt.Println("value of myTotal as a sum of many is", myTotal)
	util.Spacer()

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumLegs = 4
	dog.says()		// NOTICE a function got attached to a dog using a "reciever function"

	cat := Animal{
		Name: "cat",
		Sound: "meow",
		NumLegs: 4,
	}
	cat.says()

	dog.legs()
	cat.legs()
	// util.Spacer()
}

// ====================================================================

func doSomething(s string) {
	for i := 1; i < 5; i++ {
		fmt.Println("s is", s)
	}
}


var keyPressChannel chan rune   		// 'chan' is a channel type

func listenForKeyPress() {
	for {
		key := <- keyPressChannel
		fmt.Println("Key pressed:", string(key))
	}
}

// UNIQUE TO GO: channels are means allowing your program to send information from one place to another
// aha - that's that thing used with GoRoutines to synchronize data between them
// That is de-facto messaging built-in, to run without a Redis/etc for a monolith
// (similar to that in-process queue in Laravel LEVELS used, IIRC?)
func Channels() {
	util.Section("Channels:")

	go doSomething("Hello world")
	fmt.Println("This is a message from the main program")
	// for {
	// 	// do nothing -- so it doesn't end the program, as we're interested in the goroutine running
	// }

	//

	keyPressChannel = make(chan rune) 	// you need to create it using the 'make' keyword, same as for maps

	go listenForKeyPress()

	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	util.Spacer()
	fmt.Println("Listening to key-presses in a Goroutine")
	fmt.Println("Press any key, or q to quit")
	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChannel <- char  // send the information via the channel to the go-routine
	}
}
