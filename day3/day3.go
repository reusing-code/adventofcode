package day3

type coord struct {
	x int
	y int
}

func getSpiralMemorySteps(memoryCell int) int {
	res := 0
	coords := getCoordinate(memoryCell)
	res = manhattenDistance(coords)

	return res
}

func getRank(memoryCell int) int {
	for i := 1; ; i += 2 {
		if memoryCell <= i*i {
			return i
		}
	}
}

func getSizeOfRank(rank int) int {
	if rank == 1 {
		return 1
	}
	return square(rank) - square(rank-2)
}

func getCoordinate(memoryCell int) coord {
	if memoryCell <= 1 {
		return coord{0, 0}
	}
	rank := getRank(memoryCell)
	rankSize := getSizeOfRank(rank)
	rankQuadrantSize := rankSize / 4
	cellIndexInRank := memoryCell - 1 - square(rank-2)
	cellIndexInQuadrant := cellIndexInRank % rankQuadrantSize

	x := (rank - 1) / 2
	y := cellIndexInQuadrant - (x - 1)

	// shift to correct quadrant
	var quad int = cellIndexInRank / rankQuadrantSize
	switch quad {
	case 1:
		tmpx := -y
		y = x
		x = tmpx
	case 2:
		tmpx := -x
		y = -y
		x = tmpx
	case 3:
		tmpx := y
		y = -x
		x = tmpx
	}
	return coord{x, y}
}

func square(x int) int {
	return x * x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattenDistance(c coord) int {
	return abs(c.x) + abs(c.y)
}

var stressTestMap = make(map[coord]int)

func getStressTestValue(input int) int {
	for i := 0; ; i++ {
		v := getStressTestValueOfCell(i)
		if v > input {
			return v
		}
	}
}

func getStressTestValueOfCell(input int) int {
	stressTestMap[coord{0, 0}] = 1
	c := getCoordinate(input)
	res, exists := stressTestMap[c]
	if exists == false {
		res = calculateStressTestValueOfCell(input)
	}
	return res
}

func calculateStressTestValueOfCell(cell int) int {
	res := int(0)
	// recursively fills grid, if not already done
	if cell > 1 {
		getStressTestValueOfCell(cell - 1)
	}

	c := getCoordinate(cell)
	for xd := -1; xd <= 1; xd++ {
		for yd := -1; yd <= 1; yd++ {
			cd := coord{c.x + xd, c.y + yd}
			res += stressTestMap[cd]
		}
	}

	stressTestMap[c] = res
	return res
}
