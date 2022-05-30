package cache

//TODO: Redis support in future for horizontal deployments.

var storedData string

func GetCurrentMap() string {
	return storedData
}

func SetCurrentMap(jsonString string) {
	storedData = jsonString
}
