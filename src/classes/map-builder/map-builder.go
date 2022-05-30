package map_builder

import (
	"math/rand"
	"tuble/src/classes/block"
	"tuble/src/config"
	"tuble/src/extensions"
)

type Generator struct {
	mapSize       int8
	validPath     [][2]int8
	creatorEdgesX [2]int8
	creatorEdgesY [2]int8
	creatorEdges  [][2]int8

	gameMap [][]block.Block
}

func New() [][]block.Block {
	g := Generator{}
	g.initializeVariables()
	g.createValidPath()
	g.connectBlocksOnPath()
	return g.gameMap
}

//Called methods.

func (g *Generator) initializeVariables() {
	g.mapSize = int8(config.Map.Size)
	g.creatorEdgesX = [2]int8{0, g.mapSize - 1}
	g.creatorEdgesY = [2]int8{0, g.mapSize - 1}
	g.creatorEdges = [][2]int8{}
	g.gameMap = make([][]block.Block, g.mapSize)
	for valueX := int8(0); valueX < g.mapSize; valueX++ {
		g.gameMap[valueX] = make([]block.Block, g.mapSize)
		for valueY := int8(0); valueY < g.mapSize; valueY++ {
			if extensions.ArraySearch(valueX, g.creatorEdgesX) >= 0 || extensions.ArraySearch(valueY, g.creatorEdgesY) >= 0 {
				g.creatorEdges = append(g.creatorEdges, [2]int8{valueX, valueY})
			}
			//Initialize blocks and assign initial randomness.
			b := block.Create(valueX, valueY)
			b.SetRandomConnections() //Random connections == rotation.
			b.SetRandomSpecialType()
			g.gameMap[valueX][valueY] = b
		}
	}
}

func (g *Generator) createValidPath() {
	minConnected := int(config.Map.MinConnected)
	mustFinishOnEdge := config.Map.ForceEdgeFinish
	coords := g.pickRandomEdgeCoords()

	for { //While loop.
		g.addCoordsToPathHistory(coords)
		if (minConnected <= len(g.validPath)) && (!mustFinishOnEdge || g.isEdgeCoords(coords)) {
			//Halt if current coords are finished.
			return
		}

		coordCandidates := g.getValidSurroundingCoords(coords)
		if coordCandidates == nil {
			//No valid coordinates around -> rewind validPath until.
			coords = g.rewindPathUntilValid()
		} else {
			//Valid path available -> continue.
			coords = pickRandomCoords(coordCandidates)
		}
	}
}

func (g *Generator) connectBlocksOnPath() {
	lastKey := len(g.validPath) - 1
	prevKey := 0
	for key := 1; key <= lastKey; key++ {
		coords := g.validPath[key]
		activeBlock := &g.gameMap[coords[0]][coords[1]]
		prevCoords := g.validPath[prevKey]
		prevBlock := &g.gameMap[prevCoords[0]][prevCoords[1]]

		//Connect previous and current blocks.
		if prevKey == 0 {
			prevBlock.SetStartingBlock(coords)
		} else {
			prevBlock.SetConnectionToCoords(coords, true)
		}
		activeBlock.SetConnectionToCoords(prevCoords, false)
		if key == lastKey {
			activeBlock.SetType(block.TypeEndpoint)
			activeBlock.SetSecondConnection(block.NoConnection)
		}

		//Rotate connections in a controlled manner if not start or finish.
		if activeBlock.IsMoveable() {
			//activeBlock.RandomRotate() TODO: Re-enable this.
		}
		prevKey = key
	}
}

//Internal methods.

func (g *Generator) rewindPathUntilValid() [2]int8 {
	//Get previous element and remove from path history.
	prevValidCoords := g.getLastValidPathHistory()
	coordCandidates := g.getValidSurroundingCoords(prevValidCoords) //Pick a different candidate.
	g.rewindLastValidPathHistory()
	if coordCandidates == nil {
		return g.rewindPathUntilValid()
	} else {
		return pickRandomCoords(coordCandidates)
	}
}

//Helper methods.

func (g *Generator) pickRandomEdgeCoords() [2]int8 {
	return pickRandomCoords(g.creatorEdges)
}

func pickRandomCoords(coords [][2]int8) [2]int8 {
	rand.Seed(extensions.NewSeed())
	return coords[rand.Intn(len(coords))]
}

func (g *Generator) addCoordsToPathHistory(coords [2]int8) {
	g.validPath = append(g.validPath, coords)
}

func (g *Generator) getValidSurroundingCoords(coords [2]int8) [][2]int8 {
	initialSet := [4][2]int8{
		0: {coords[0], coords[1] - 1}, //Towards top.
		1: {coords[0] + 1, coords[1]}, //Towards right.
		2: {coords[0], coords[1] + 1}, //Towards bottom.
		3: {coords[0] - 1, coords[1]}, //Towards left.
	}
	//Must be within limits.
	var limitedSet [][2]int8
	for _, coords := range initialSet {
		if block.CoordsAreWithinLimits(coords, g.creatorEdgesX, g.creatorEdgesY) {
			limitedSet = append(limitedSet, coords)
		}
	}
	//Must not be previously stepped on.
	var validatedSet [][2]int8
	for _, coords := range limitedSet {
		if extensions.ArraySearch(coords, g.validPath) < 0 {
			validatedSet = append(validatedSet, coords)
		}
	}
	return validatedSet
}

func (g *Generator) getLastValidPathHistory() [2]int8 {
	pos := len(g.validPath) - 2
	if pos < 0 {
		panic("PANIC: Map generator failure!")
	}
	return g.validPath[pos] // -1 is current, -2 is previous.
}

func (g *Generator) rewindLastValidPathHistory() {
	g.validPath = g.validPath[:len(g.validPath)-1]
}

func (g *Generator) isEdgeCoords(coords [2]int8) bool {
	return extensions.ArraySearch(coords[0], g.creatorEdgesX) >= 0 || extensions.ArraySearch(coords[1], g.creatorEdgesY) >= 0
}
