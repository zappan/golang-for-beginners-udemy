package decisions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func rockPaperScissors_viaSwitch() {
	rand.Seed(time.Now().UnixNano())

	playerChoice := ""
	playerValue := -1
	computerValue := rand.Intn(3)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter rock, paper, or scissors -> ")

	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	switch playerChoice {
	case "rock":
		playerValue = ROCK
	case "paper":
		playerValue = PAPER
	case "scissors":
		playerValue = SCISSORS
	default:
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose ROCK")
	case PAPER:
		fmt.Println("Computer chose PAPER")
	case SCISSORS:
		fmt.Println("Computer chose SCISSORS")
	}

  fmt.Println("Player chose", strings.ToUpper(playerChoice))

  if playerValue == computerValue {
    fmt.Println("It's a draw")
  } else if playerValue == -1 {
    fmt.Println("Invalid choice!")
  } else if playerValue == (computerValue + 1) % 3 {
    fmt.Println("Player wins")
  } else {
    fmt.Println("Computer wins")
  }

  // ### SIMPLIFIED VERSION ABOVE BY USING THE MODULUS OPERATOR
	// if playerValue == computerValue {
	// 	fmt.Println("It's a draw")
	// } else {
	// 	switch playerValue {
	// 	case ROCK:
	// 		if computerValue == PAPER {
	// 			fmt.Println("Computer wins")
	// 		} else {
	// 			fmt.Println("Player wins")
	// 		}
	// 	case PAPER:
	// 		if computerValue == SCISSORS {
	// 			fmt.Println("Computer wins")
	// 		} else {
	// 			fmt.Println("Player wins")
	// 		}
	// 	case SCISSORS:
	// 		if computerValue == ROCK {
	// 			fmt.Println("Computer wins")
	// 		} else {
	// 			fmt.Println("Player wins")
	// 		}
	// 	default:
	// 		fmt.Println("Invalid choice!")
	// 	}
	// }


	fmt.Println()
}

func Switch() {
	rockPaperScissors_viaSwitch()
}
