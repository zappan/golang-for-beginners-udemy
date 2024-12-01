package decisions

import (
	"fmt"
	"time"
	"myapp/game"
)

var ch1 = make(chan string)
var ch2 = make(chan string)

func taskOne() {
	time.Sleep(1 * time.Second)
	ch1 <- "one"
}

func taskTwo() {
	time.Sleep(2 * time.Second)
	ch2 <- "two"
}

// -------------------------------------------
// select statement is ONLY USED WITH CHANNELS
// => it's like a switch, but for channels
// -------------------------------------------
func simpleSelectExample() {
	go taskOne()
	go taskTwo()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- ch1:
			fmt.Println("received", msg1)
		case msg2 := <- ch2:
			fmt.Println("received", msg2)
		}
	}
	fmt.Println()
	time.Sleep(1 * time.Second)
}

// -------------------------------------------------------------

func rockPaperScissors_viaSelect() {
	displayCh := make(chan string)
	roundCh := make(chan int)

	game := game.Game {
		DisplayCh: displayCh,
		RoundCh: roundCh,
		Round: game.Round{
			RoundNum: 0,
			PlayerScore: 0,
			ComputerScore: 0,
		},
	}

	go game.StartGame()  // starts the game...

	game.ClearScreen()
	game.PrintIntro()

	for game.Round.RoundNum < 3 {
		game.PlayRound()
	}

	game.PrintGameSummary()
	fmt.Println()
}

func Select() {
	simpleSelectExample()
	rockPaperScissors_viaSelect()
}
