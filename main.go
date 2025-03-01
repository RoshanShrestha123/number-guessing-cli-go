package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

type DifficultyState int

const (
	EASY DifficultyState = iota + 1
	MEDIUM
	HARD
)

type GuessNumber struct {
	chances         int
	userGuessNumber int
	systemNumber    int
}

func main() {

	var chancesFromDifficultyLevel = map[DifficultyState]int{
		EASY:   10,
		MEDIUM: 5,
		HARD:   3,
	}

	var levels = map[DifficultyState]string{
		EASY:   "Easy",
		MEDIUM: "Medium",
		HARD:   "Hard",
	}

	var difficulty string
	gameState := &GuessNumber{
		chances:      chancesFromDifficultyLevel[MEDIUM],
		systemNumber: rand.Intn(100),
	}

	fmt.Print(`
###  Welcome to the Number Guessing Game!  ###
I'm thinking of a number between 1 and 100.
`)

	fmt.Printf("You have %d chances to guess the correct number.", gameState.chances)

	fmt.Println(`
	
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)`)

	fmt.Print("\nEnter your choice: ")
	fmt.Scan(&difficulty)

	diffLevel, err := strconv.ParseInt(difficulty, 10, 64)
	if err != nil {
		log.Fatal("Cannot be parsed to int from string", err)

	}

	if diffLevel > 3 || diffLevel <= 0 {
		log.Fatal("Difficult level should be between 1-3")
	}

	gameState.chances = chancesFromDifficultyLevel[DifficultyState(diffLevel)]

	fmt.Printf("\nGreat! You have selected the %s difficulty level.\n", levels[DifficultyState(diffLevel)])
	fmt.Println("Lets start the game!")

	for gameState.chances > 0 {

		fmt.Print("\nEnter your guess: ")
		userGuess := ""
		fmt.Scan(&userGuess)
		value, err := strconv.ParseInt(userGuess, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if value > 100 || value <= 0 {
			log.Fatal("number should be between 1-100")

		}

		if value > int64(gameState.systemNumber) {
			fmt.Printf("Incorrect! The number is less than %d.", value)
		} else if value < int64(gameState.systemNumber) {
			fmt.Printf("Incorrect! The number is greater than %d.", value)
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", gameState.chances)
			return
		}

		gameState.chances--
		fmt.Printf("\n---%d chances left---\n", gameState.chances)

	}

	fmt.Println("No more chances for you!")

}
