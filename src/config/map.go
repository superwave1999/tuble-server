package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var Map *MapConfig

type MapConfig struct {
	ForceEdgeFinish    bool
	Size               uint8
	MinConnected       uint8
	ScoreMoveBenefit   uint32
	ScoreMovePenalty   uint32
	ScoreTimeBenefitMs uint32
	ScoreTimePenaltyMs uint32
}

func Reload() {
	//Hardcoded defaults
	c := MapConfig{
		ForceEdgeFinish:    false,
		Size:               6,
		MinConnected:       14,
		ScoreMoveBenefit:   3,
		ScoreMovePenalty:   2,
		ScoreTimeBenefitMs: 5000,
		ScoreTimePenaltyMs: 10000,
	}

	//Read from .env
	err := godotenv.Load()
	if err == nil {
		vbool, err := strconv.ParseBool(os.Getenv("MAP_EDGE_FINISH"))
		if err == nil {
			c.ForceEdgeFinish = vbool
		}
		vint8, err := strconv.ParseUint(os.Getenv("MAP_SIZE"), 10, 8)
		if err == nil {
			c.Size = uint8(vint8)
		}
		vint8, err = strconv.ParseUint(os.Getenv("MAP_MIN_CONNECTED_BLOCKS"), 10, 8)
		if err == nil {
			c.MinConnected = uint8(vint8)
		}
		vint32, err := strconv.ParseUint(os.Getenv("SCORE_MOVE_BENEFIT"), 10, 32)
		if err == nil {
			c.ScoreMoveBenefit = uint32(vint32)
		}
		vint32, err = strconv.ParseUint(os.Getenv("SCORE_MOVE_PENALTY"), 10, 32)
		if err == nil {
			c.ScoreMovePenalty = uint32(vint32)
		}
		vint32, err = strconv.ParseUint(os.Getenv("SCORE_TIME_BENEFIT_MS"), 10, 32)
		if err == nil {
			c.ScoreTimeBenefitMs = uint32(vint32)
		}
		vint32, err = strconv.ParseUint(os.Getenv("SCORE_TIME_PENALTY_MS"), 10, 32)
		if err == nil {
			c.ScoreTimePenaltyMs = uint32(vint32)
		}
	}
	checkConfig(c)
	Map = &c
}

func checkConfig(c MapConfig) {
	if c.Size < 2 {
		panic("PANIC: Configured map size is too small!")
	}
	if c.MinConnected > (c.Size * c.Size) {
		panic("PANIC: Min connected blocks cannot be greater than size^2 !")
	}
}
