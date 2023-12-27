package menu

import (
	"fmt"
	"github.com/Diaszano/goMysticNumber/internal/game"
	"github.com/Diaszano/goMysticNumber/internal/scoreboard"
	"time"
)

// Menu displays the main menu and handles user input for different options.
func Menu() {
	clearTerminal()
	fmt.Print("Menu\n1 - Play\n2 - Scoreboard\n3 - Exit\nChoose an option: ")

	var option int
	_, _ = fmt.Scanf("%d", &option)

	switch option {
	case 1:
		clearTerminal()
		points := game.Play()
		scores, _ := scoreboard.Load()
		var highestScore uint8 = 0

		if len(scores) > 0 {
			highestScore = scores[0].Points
		}

		if points >= highestScore && points > 0 {
			if points == highestScore {
				fmt.Println("Congratulations! You are among the best!")
			} else {
				fmt.Println("Wow, you set a new record!")
			}
			var name string
			fmt.Print("Enter your name to save your score: ")
			_, _ = fmt.Scanf("%s", &name)
			_, _ = scoreboard.Save(scoreboard.Score{Name: name, Points: points, Created: time.Now()})
		}
		pause()
	case 2:
		clearTerminal()
		scores, _ := scoreboard.Load()
		fmt.Println("Scoreboard in order of points:")
		for i := 0; i < len(scores); i++ {
			fmt.Printf("  %dº %v", i+1, scores[i])
		}
		pause()
	case 3:
		clearTerminal()
		return
	}
	Menu()
}

// clearTerminal clears the terminal screen.
func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

// pause prompts the user to press Enter to continue.
func pause() {
	fmt.Println("Press Enter to continue...")
	_, _ = fmt.Scanln()
}
