package boot

import (
	"tuble/src/storage/cache"
	"tuble/src/storage/files"
	"tuble/src/utils"
)

func RestoreOrCreateMap() {
	existMapJson := files.Read(files.MapFile)
	if len(existMapJson) > 0 {
		cache.SetCurrentMap(string(existMapJson))
	} else {
		//Generate and store new.
		utils.MapToJson()
	}
}
