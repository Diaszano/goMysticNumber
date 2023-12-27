package game

import (
	"fmt"
	"github.com/Diaszano/goMysticNumber/internal/random"
)

const (
	gameLimit         = 5
	pointsStart uint8 = 100
)

// Play initiates the game, generating a random number and allowing the player
// to make up to gameLimit attempts to guess it. The function returns the final
// score after the game.
func Play() uint8 {
	number := random.Range(1, 100)
	points := pointsStart

	for i := 0; i < gameLimit; i++ {
		fmt.Print("Enter the value you think is between 1 and 100: ")
		var attempt int
		_, _ = fmt.Scanf("%d", &attempt)

		if attempt == number {
			printMessage("You guessed the correct number! It was %d", number, 32)
			return points
		}

		if attempt > number {
			printMessage("The number is less than %d", attempt, 36)
		} else {
			printMessage("The number is greater than %d", attempt, 34)
		}

		points -= pointsStart / gameLimit
	}

	printMessage("You ran out of attempts! The number was %d", number, 31)
	return 0
}

// printMessage formats and prints a colored message to the console.
func printMessage(message string, value int, colorCode int) {
	fmt.Printf("\033[1;%dm%s\033[0;0m\n", colorCode, fmt.Sprintf(message, value))
}
