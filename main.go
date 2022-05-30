package main

import (
	"tuble/src/boot"
	"tuble/src/config"
)

func main() {
	config.Reload()
	boot.RestoreOrCreateMap()
	boot.StartScheduler()
	boot.StartHTTP()
}
