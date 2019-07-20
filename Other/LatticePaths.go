package Other

import "fmt"

func LatticePaths() {
	var a []cood
	//Grid size can be changed although higher numbers contain many paths and this records all
	//20 x 20 = 137846528820(number of paths) * 40(int per path) * 1 byte(8 bits per int) = 5.5 terabytes
	//Ergo, this program cannot be used to solve high values and serves only to find out the paths of lower sized grids
	c := calcPath(a, cood{0, 0}, 10, 10)
	for i, path := range c {
		fmt.Printf("Path %d is:%v\n", i+1, path)
	}
}

func calcPath(currPath []cood, currCood cood, maxX int8, maxY int8) [][]cood {
	currPath = append(currPath, currCood)
	var x, y [][]cood
	if currCood.x < maxX {
		nextCood := cood{currCood.x + 1, currCood.y}
		x = calcPath(currPath, nextCood, maxX, maxY)
	}
	if currCood.y < maxY {
		nextCood := cood{currCood.x, currCood.y + 1}
		y = calcPath(currPath, nextCood, maxX, maxY)
	}
	if currCood == (cood{maxX, maxY}) {
		x = append(x, currPath)
	}
	for _, b := range y {
		x = append(x, b)
	}
	return x
}

type cood struct {
	x int8
	y int8
}
