package main

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/whodatXLIV/MazeSolving/maze"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("File name of maze: ")
	filename, _ := reader.ReadString('\n')
	filename = filename[:len(filename)-1]

	imfile, err := os.Open("MazeImages/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer imfile.Close()
	img, err := png.Decode(imfile)
	if err != nil {
		log.Fatal(err)
	}

	m, graph, err := maze.PrepareMaze(img)
	if err != nil {
		log.Fatal(err)
	}

	ent, err := maze.FindEntrance(img)
	if err != nil {
		log.Fatal(err)
	}

	ext, err := maze.FindExit(img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Maze Succesfully Loaded!!!\n")

	fmt.Print("Please specify solver to be used\nSolvers include BFS, DFS. The default solver is BFS: ")
	solver, _ := reader.ReadString('\n')
	solver = solver[:len(solver)-1]

	var sol []int
	start := time.Now()
	switch solver {
	case "BFS":
		sol = maze.BreadthFirst(m, ent, ext)
	case "DFS":
		sol = maze.DepthFirst(m, ent, ext)
		//	case "AStar":
		//		sol = maze.AStar(m, ent, ext)
	default:
		sol = maze.BreadthFirst(m, ent, ext)
		solver = "BFS"
	}
	elapsed := time.Since(start)

	sImg := maze.SolvedColor(graph, ent, ext, sol)

	outfile, err := os.Create("MazeImages/" + filename[:len(filename)-4] + "_" + solver + ".png")
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	pl := maze.PathLength(sol)
	fmt.Println("The length of the path found is: ", pl)
	fmt.Println("The time taken to solve is: ", elapsed)
	png.Encode(outfile, sImg)

}
