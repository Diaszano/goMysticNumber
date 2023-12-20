package scoreboard

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Score struct {
	Name    string    `json:"name"`
	Points  uint8     `json:"points"`
	Created time.Time `json:"created"`
}

type scoreboardCollection struct {
	Scores []Score `json:"scores"`
}

const fileName = "points.json"

func Save(score Score) (scores []Score, err error) {
	scores, err = Load()
	if err != nil {
		return scores, fmt.Errorf("")
	}
	scores = append(scores, score)
	ad := scoreboardCollection{Scores: scores}
	_ = write(ad)
	return scores, nil
}

func write(scoreboard scoreboardCollection) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create the scoreboardCollection %s: %v", fileName, err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(scoreboard, "", "\t")
	if err != nil {
		return fmt.Errorf(": %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write the scoreboardCollection %s: %v", fileName, err)
	}
	return nil
}

// Load function reads the scoreboardCollection data from a file and returns the list of scores.
//
// It internally uses the read function to read the data from the file specified by fileName.
// If the read operation is successful, the function returns the list of scores and a nil error.
// If any error occurs during the file reading or JSON parsing, an error message is returned,
// and the scores slice is empty.
//
// Example:
//
//	scores, err := Load()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Now 'scores' contains the list of scores from the scoreboardCollection.
//
// Returns:
//   - scores: A slice of Score struct representing the scores loaded from the scoreboardCollection.
//   - err: An error indicating the success or failure of the load operation.
func Load() (scores []Score, err error) {
	scoreboard, err := read()
	if err != nil {
		return scores, fmt.Errorf("failed to load the scoreboardCollection %s: %v", fileName, err)
	}

	scores = scoreboard.Scores
	return scores, nil
}

// read function reads a JSON file and parses its content into a scoreboardCollection.
// The file specified by the fileName parameter should contain valid JSON data
// representing the scoreboardCollection structure.
//
// If the file is successfully read and parsed, the function returns the populated
// scoreboardCollection and a nil error. If any error occurs during file reading or JSON parsing,
// an error message is returned along with an empty scoreboardCollection.
//
// Example:
//
//	table, err := read()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Now 'table' contains the parsed data.
//
// Returns:
//   - table: A scoreboardCollection struct containing the parsed data from the file.
//   - err: An error indicating the success or failure of the read operation.
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
		return scoreboard, fmt.Errorf("fails to parse JSON: %v", err)
	}

	return scoreboard, nil
}
