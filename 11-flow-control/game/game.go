package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"myapp/util"
)


type Game struct {
	DisplayCh chan string
	RoundCh chan int
	Round Round
}

type Round struct {
	RoundNum int
	PlayerScore int
	ComputerScore int
}

const (
	ROCK = 0
	PAPER = 1
	SCISSORS = 2
)

var reader = bufio.NewReader(os.Stdin)


func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayCh <- "Computer wins!"
	<- g.DisplayCh
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayCh <- "Player wins!"
	<- g.DisplayCh
}


func (g *Game) ClearScreen() {
	util.ClearScreen()
}


func (g *Game) StartGame() {
	// use select to process input in channels
		// Print to screen
		// Keep track of round number
	// runs forever
	for {
		select {
		case round := <- g.RoundCh:
			g.Round.RoundNum = g.Round.RoundNum + round
			g.RoundCh <- 1  // send something back, as main program is waiting for some result...
		case msg := <- g.DisplayCh:
			fmt.Println(msg)
			g.DisplayCh <- ""
		}

	}
}


func (g *Game) PrintIntro() {
	gameName := "Rock, Paper & Scissors"
	g.DisplayCh <- fmt.Sprintf(`
%s
%s
Game is played for three rounds, and best of three wins the game!

`, gameName, strings.Repeat("=", len(gameName)))
	<- g.DisplayCh
}


func (g *Game) PlayRound() {
	rand.Seed(time.Now().UnixNano())

	// this is a bit artificial, probably just to force using channels due to select
	g.RoundCh <- 1
	<- g.RoundCh  // wait for a result from g.RoundCh

	roundTitle := fmt.Sprintf("Round %d", g.Round.RoundNum)
  playerValue := -1
	computerValue := rand.Intn(3)
	roundFailed := false

	g.DisplayCh <- fmt.Sprintf(`%s
%s`, roundTitle, strings.Repeat("-", len(roundTitle)))
	<- g.DisplayCh

 	fmt.Print("Please enter rock, paper, or scissors -> ")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	switch playerChoice {
	case "rock":
		playerValue = ROCK
	case "paper":
		playerValue = PAPER
	case "scissors":
		playerValue = SCISSORS
	}

	g.DisplayCh <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	<- g.DisplayCh


	var computerChoice string
	switch computerValue {
	case ROCK:
		computerChoice = "ROCK"
	case PAPER:
		computerChoice = "PAPER"
	case SCISSORS:
		computerChoice = "SCISSORS"
	}
	g.DisplayCh <- fmt.Sprintf("Computer chose %s", computerChoice)
	<- g.DisplayCh



	if playerValue == computerValue {
		roundFailed = true  // round failed for a draw
		g.DisplayCh <- "It's a draw"
		<- g.DisplayCh
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
		default:
			roundFailed = true  // round failed an invalid choice
			g.DisplayCh <- "Invalid choice!"
			<- g.DisplayCh
		}
	}

	if roundFailed {
		g.RoundCh <- -1  // reduce so that it can be increased at the beginning of the for loop
		<- g.RoundCh     // wait for a result from g.RoundCh
	}

	g.DisplayCh <- ""
	<- g.DisplayCh
}


func (g *Game) PrintGameSummary() {
	var winner string
	title := "Final score"
	scoreSummary := fmt.Sprintf("Player: %d/3, Computer: %d/3", g.Round.PlayerScore, g.Round.ComputerScore)
	if (g.Round.PlayerScore > g.Round.ComputerScore) {
		winner = "Player"
	} else {
		winner = "Computer"
	}


	g.DisplayCh <- fmt.Sprintf(`%s
%s
%s
%s wins the game!`, title, strings.Repeat("-", len(title)), scoreSummary, winner)
	<- g.DisplayCh

}
