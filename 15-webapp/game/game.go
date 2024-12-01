package game

import (
	"math/rand"
	"time"
)


const (
	ROCK = 0      // beats scissors (scissors + 1) % 3 == 0
	PAPER = 1     // beats rock (rock + 1) % 3 == 1
	SCISSORS = 2  // beats paper (paper + 1) % 3 == 2
	DRAW = 0
	PLAYER_WINS = 1
	COMPUTER_WINS = 2
)

type Round struct {
	Winner int 						`json:"winner"`
	PlayerChoice string   `json:"playerChoice"`
	ComputerChoice string `json:"computerChoice"`
	RoundResult string 		`json:"roundResult"`
	WisdomNugget string		`json:"wisdomNugget"`
}

func getRandomMessage(messages []string) string {
	return messages[rand.Intn(len(messages))]
}

func getRandomWinMessage() string {
	return getRandomMessage([]string{
		"You should buy a lottery ticket!",
		"Rock 'n roll!",
		"Go for it, champ!",
	})
}

func getRandomLoseMessage() string {
	return getRandomMessage([]string{
		"Too bad!",
		"Give it another chance...",
		"Come on, the statistics will work for you in the long run!",
	})
}

func getRandomDrawMessage() string {
	return getRandomMessage([]string{
		"Great minds think alike!",
		"Uh, oh. Try again.",
		"Nobody wins, but you can try again.",
	})
}


func PlayRound(playerValue int) (Round) {
	computerValue := rand.Intn(3)
	playerChoice := ""
	computerChoice := ""
	roundResult := ""
	wisdomNugget := ""
	winner := -1

	rand.Seed(time.Now().UnixNano())

	switch playerValue {
	case ROCK:
		playerChoice = "Player chose ROCK"
	case PAPER:
		playerChoice = "Player chose PAPER"
	case SCISSORS:
		playerChoice = "Player chose SCISSORS"
	}

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	}

  if playerValue == computerValue {
    winner = DRAW
    roundResult = "It's a draw"
    wisdomNugget = getRandomDrawMessage()
  } else if playerValue == (computerValue + 1) % 3 {
    winner = PLAYER_WINS
    roundResult = "Player wins!"
    wisdomNugget = getRandomWinMessage()
  } else {
    winner = COMPUTER_WINS
    roundResult = "Computer wins!"
    wisdomNugget = getRandomLoseMessage()
  }

	result := Round {
		Winner: winner,
		PlayerChoice: playerChoice,
		ComputerChoice: computerChoice,
		RoundResult: roundResult,
		WisdomNugget: wisdomNugget,
	}
	return result
}
