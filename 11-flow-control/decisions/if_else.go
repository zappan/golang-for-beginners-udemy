package decisions

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"strings"
	"time"
)


func rockPaperScissors_viaIfElse() {
	rand.Seed(time.Now().UnixNano())

	playerChoice := ""
	playerValue := -1
	computerValue := rand.Intn(3)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter rock, paper, or scissors -> ")

	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println("Player chose", playerChoice, "and value is", playerValue)
	fmt.Println("Computer value is", computerValue)
	fmt.Println()
}


func IfElse() {
	rockPaperScissors_viaIfElse()
}
