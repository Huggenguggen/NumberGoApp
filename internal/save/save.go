// Package save implemenets save file
// encodes to json such that game can continue
package save

import (
	"encoding/json"
	"os"
)

func SaveGame(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func LoadGame(path string, data any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(data)
}
