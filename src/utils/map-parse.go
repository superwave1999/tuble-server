package utils

import (
	"encoding/json"
	"tuble/src/classes/block"
)

func JsonBytesToMap(bytes []byte) (MapInput, error) {
	var mapInput MapInput
	err := json.Unmarshal(bytes, &mapInput)
	if err != nil {
		return mapInput, err
	}
	return mapInput, nil
}

type MapInput struct {
	Map    [][]block.Block
	TimeMs int
	Moves  int
}
