package mazesolver

import "fmt"

var maze = []string{
	"xxxxxxxxxx x",
	"x        x x",
	"x        x x",
	"x xxxxxxxx x",
	"x          x",
	"x xxxxxxxxxx",
}

type Point struct {
	x int
	y int
}

var mazeResult = []Point{
	{x: 10, y: 0},
	{x: 10, y: 1},
	{x: 10, y: 2},
	{x: 10, y: 3},
	{x: 10, y: 4},
	{x: 9, y: 4},
	{x: 8, y: 4},
	{x: 7, y: 4},
	{x: 6, y: 4},
	{x: 5, y: 4},
	{x: 4, y: 4},
	{x: 3, y: 4},
	{x: 2, y: 4},
	{x: 1, y: 4},
	{x: 1, y: 5},
}

var startingPoint = Point{x: 10, y: 0}
var finshingPoint = Point{x: 1, y: 5}

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall string, curr, end Point, seen [][]bool, path *[]Point) bool {
	fmt.Println(*path)
	// Base case
	// 1. off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y < len((maze)) {
		return false
	}

	// 2. on the wall
	if string(maze[curr.y][curr.x]) == wall {
		return false
	}

	// 3. end point
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}

	// 4. seen
	if seen[curr.y][curr.x] {
		return false
	}

	// recurse
	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// recurse
	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]

		if walk(maze, wall, Point{
			x: curr.x + x,
			y: curr.y + y,
		}, end, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]

	return false
}

func Solve(maze []string, wall string, start, end Point) []Point {
	var seen = [][]bool{}
	var path = []Point{}

	for i := 0; i < len(maze); i++ {
		row := make([]bool, len(maze[0]))
		seen = append(seen, row)
	}

	walk(maze, wall, start, end, seen, &path)

	return path
}

func RunMazeSolver() {
	result := Solve(maze, "x", startingPoint, finshingPoint)
	fmt.Println(result)
}
