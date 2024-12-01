package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

// own custom type
type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User
	var dogOwnership string

	user.UserName = readString("What is your name?")
	// fmt.Println("Your name is", user.UserName)

	user.Age = readInt("How old are you?")
	// fmt.Println("Your age is", user.Age)

	user.FavouriteNumber = readFloat("What's your favourite number?")
	user.OwnsADog = readBool("Do you own a dog (y/n)?")

	// ...adds spaces between the arguments, which isn't good for the full-stop after the username
	fmt.Println("Your name is", user.UserName, "and you are", user.Age, "years old.")
	fmt.Println("Your name is", user.UserName, ". You are", user.Age, "years old.")

	// easy, but inefficient way: concatenation operator joining strings
	// --> SLOW AND MEMORY USE IS HIGHER
	// --> ALSO, DOESN'T WORK ON AGE AS IT'S A NUMBER RATHER THAN A STRING
	// fmt.Println("Your name is" + user.UserName + ". You are" + user.Age + "years old.")
	fmt.Println("Your name is "+user.UserName+". You are", user.Age, "years old.")

	// the efficient way - (Sprintf = String-print-format)
	// -- C-like
	fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", user.UserName, user.Age))
	// and an even more efficient way with only one function and newline terminator
	// -- even more C-like with that \n at the end
	fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.2f\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber)

	fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.2f. Owns a dog: %t.\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwnsADog)

	if user.OwnsADog {
		dogOwnership = "are"
	} else {
		dogOwnership = "aren't"
	}
	fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.2f, and you %s a dog owner.\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		dogOwnership)
}

func prompt() {
	fmt.Print("-> ")
}

func askAQuestion(q string) {
	fmt.Println(q)
	prompt()
}

func readInput(q string) string {
	askAQuestion(q)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func readString(q string) string {
	for {
		// fmt.Println(q)
		// prompt()

		// userInput, err := reader.ReadString('\n')
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// userInput = strings.Replace(userInput, "\r\n", "", -1)
		// userInput = strings.Replace(userInput, "\n", "", -1)
		userInput := readInput(q)
		if userInput == "" {
			fmt.Println("Please enter value")
		} else {
			return userInput
		}
	}
}

func readInt(q string) int {
	for {
		// fmt.Println(q)
		// prompt()

		// userInput, err := reader.ReadString('\n')
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// userInput = strings.Replace(userInput, "\r\n", "", -1)
		// userInput = strings.Replace(userInput, "\n", "", -1)
		userInput := readInput(q)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(q string) float64 {
	for {
		userInput := readInput(q)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBool(q string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Printf("%s ", q)

		userInput, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Please enter a value")
		} else {
			// answer := string(userInput)
			// fmt.Println(answer)

			// // switch-case would be approproate here, or 'includes'
			// if answer == "y" || answer == "Y" {
			// 	return true
			// } else if answer == "n" || answer == "N" {
			// 	return false
			// } else {
			// 	// fmt.Println("Please enter 'y' or 'n")
			// }


			answer := strings.ToLower(string(userInput))
			fmt.Println(answer)

			if answer != "y" && answer != "n" {
				fmt.Println("Please enter 'y' or 'n")
			} else {
				return answer	== "y"
			}
		}
	}
}
