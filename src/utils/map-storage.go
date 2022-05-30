package utils

import (
	"encoding/json"
	"strconv"
	"time"
	"tuble/src/classes/block"
	MapBuilder "tuble/src/classes/map-builder"
	"tuble/src/storage/cache"
	"tuble/src/storage/files"
)

func mapStorageStruct(builtMap [][]block.Block) string {
	utc := time.Now().UTC()
	jsonStruct := MapFile{
		Date:      utc.Format("2006-01-02"),
		Timestamp: strconv.FormatInt(utc.UnixMilli(), 10),
		Map:       builtMap,
		//TODO: Implement map safety. Secure token for frontend. This data contains unmovable block data.
	}
	marshal, _ := json.Marshal(jsonStruct)
	return string(marshal)
}

func MapToJson() {
	newMap := mapStorageStruct(MapBuilder.New())
	cache.SetCurrentMap(newMap)
	files.Write(files.MapFile, []byte(newMap))
}

type MapFile struct {
	Date      string
	Timestamp string
	Map       [][]block.Block
}
