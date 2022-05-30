package files

import (
	"fmt"
	"os"
)

const MapFile = ".mapcache.json"

func Read(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func Write(name string, contents []byte) {
	err := os.WriteFile(name, contents, 0777)
	if err != nil {
		fmt.Println(err)
	}
}
