// Package scoreboard provides functionality to manage and persist a collection of scores.
//
// The package includes methods to save and load scores from a JSON file, as well as a Score struct
// representing an individual score entry.
package scoreboard

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Score represents an individual score entry with a name, points, and creation timestamp.
type Score struct {
	Name    string    `json:"name"`
	Points  uint8     `json:"points"`
	Created time.Time `json:"created"`
}

func (sc Score) String() string {
	return fmt.Sprintf("Player %s has %d points.\n", sc.Name, sc.Points)
}

// scoreboardCollection represents a collection of scores.
type scoreboardCollection struct {
	Scores []Score `json:"scores"`
}

const fileName = "points.json"

// Save appends a new score to the existing scoreboard and saves it to the file.
//
// It loads the existing scores, appends the new score, and writes the updated scoreboard
// to the file. If any error occurs during the process, it returns the existing scores and the error.
func Save(score Score) (scores []Score, err error) {
	scores, err = Load()
	if err != nil {
		return scores, fmt.Errorf("failed to load scores: %v", err)
	}

	scores = append(scores, score)

	scoreboard := scoreboardCollection{Scores: scores}

	err = write(scoreboard)
	if err != nil {
		return scores, fmt.Errorf("failed to save scores: %v", err)
	}

	return scores, nil
}

// write writes the scoreboardCollection to the specified file in JSON format.
//
// It creates or truncates the file, marshals the scoreboardCollection to JSON, and writes
// the data to the file. If any error occurs during this process, it returns an error.
func write(scoreboard scoreboardCollection) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", fileName, err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(scoreboard, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", fileName, err)
	}
	return nil
}

// Load reads the scoreboardCollection data from a file and returns the list of scores.
//
// It internally uses the read function to read the data from the file specified by fileName.
// If any error occurs during the file reading or JSON parsing, an error message is returned,
// and the scores slice is empty.
func Load() (scores []Score, err error) {
	scoreboard, err := read()
	if err != nil {
		return scores, fmt.Errorf("failed to load scores: %v", err)
	}

	scores = scoreboard.Scores
	return sortScoreByPoints(scores), nil
}

// read reads a JSON file and parses its content into a scoreboardCollection.
//
// If any error occurs during file reading or JSON parsing, an error message is returned along
// with an empty scoreboardCollection.
func read() (scoreboard scoreboardCollection, err error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return scoreboard, nil
		}
		return scoreboard, fmt.Errorf("failed to read file %s: %v", fileName, err)
	}

	if !json.Valid(file) {
		err = os.Remove(fileName)
		if err != nil {
			return scoreboard, fmt.Errorf("failed to remove old file %s: %v", fileName, err)
		}
		return scoreboard, nil
	}

	err = json.Unmarshal(file, &scoreboard)
	if err != nil {
		return scoreboard, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return scoreboard, nil
}

// sortScoreByPoints sorts the scores in descending order based on points.
func sortScoreByPoints(scores []Score) []Score {
	for i := 0; i < len(scores)-1; i++ {
		for j := i + 1; j < len(scores); j++ {
			if scores[i].Points < scores[j].Points {
				scores[i], scores[j] = scores[j], scores[i]
			}
		}
	}
	return scores
}
