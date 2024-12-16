const maze = [
    "xxxxxxxxxx x",
    "x        x x",
    "x        x x",
    "x xxxxxxxx x",
    "x          x",
    "x xxxxxxxxxx",
];

type Point = {
    x: number,
    y: number,
}

const mazeResult = [
    { x: 10, y: 0 },
    { x: 10, y: 1 },
    { x: 10, y: 2 },
    { x: 10, y: 3 },
    { x: 10, y: 4 },
    { x: 9, y: 4 },
    { x: 8, y: 4 },
    { x: 7, y: 4 },
    { x: 6, y: 4 },
    { x: 5, y: 4 },
    { x: 4, y: 4 },
    { x: 3, y: 4 },
    { x: 2, y: 4 },
    { x: 1, y: 4 },
    { x: 1, y: 5 },
];

const startingPoint = { x: 10, y: 0 }
const finshingPoint = { x: 1, y: 5 }

const dir = [
    [-1, 0],
    [1, 0],
    [0, -1],
    [0, 1],
];

function walk(maze: string[], wall: string, curr: Point, end: Point, seen: boolean[][], path: Point[]): boolean {
    // Base case
    // 1. off the map
    if (curr.x < 0 || curr.x >= maze[0].length || curr.y < 0 || curr.y >= maze.length) {
        return false
    }

    // 2. on a wall
    if (maze[curr.y][curr.x] == wall) {
        return false
    }

    // 3. end point
    if (curr.x == end.x && curr.y == end.y) {
        path.push(end)
        return true
    }

    // 4. seen
    if (seen[curr.y][curr.x]) {
        return false
    }

    // RECURSE

    // pre
    seen[curr.y][curr.y] = true
    path.push(curr)

    // recurse
    for (let i = 0; i < dir.length; ++i) {
        const [x, y] = dir[i]

        if (walk(maze, wall, {
            x: curr.x + x,
            y: curr.y + y,
        }, end, seen, path)) {
            return true
        }
    }


    // post
    path.pop()

    return false
}

export default function MazeSolver(maze: string[], wall: string, start: Point, end: Point): Point[] {
    const seen: boolean[][] = []
    const path: Point[] = []

    for (let i = 0; i < maze.length; ++i) {
        seen.push(new Array(maze[0].length).fill(false))
    }

    walk(maze, wall, start, end, seen, path);

    return path
}

console.log(MazeSolver(maze, "x", startingPoint, finshingPoint))