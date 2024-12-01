package decisions

const (
	ROCK = 0      // beats scissors (scissors + 1) % 3 == 0
	PAPER = 1     // beats rock (rock + 1) % 3 == 1
	SCISSORS = 2  // beats paper (paper + 1) % 3 == 2
)
