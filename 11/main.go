package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readGrid(s string) [10][10]int {
	var out [10][10]int
	file, err := os.Open(s)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var i int
	for scanner.Scan() {
		for j, r := range []rune(string(scanner.Text()[:])) {
			out[i][j], err = strconv.Atoi(string(r))
			check(err)
		}
		i++
	}
	return out
}

func nextRound(g *[10][10]int) [10][10]int {
	var out [10][10]int
	for i, el1 := range *g {
		for j, el2 := range el1 {
			out[i][j] = el2+1
		}
	}
	return out
}

func flash(g [10][10]int, a, b int) ([10][10]int, uint) {
	if a < 0 || a > 9 || b < 0 || b > 9 {
		return g, 0
	}
	if g[a][b] == -1 {
		return g, 0
	}
	var out uint
	g[a][b]++
	if g[a][b] > 10 {
		out++
		var tmp uint
		g[a][b] = -1
		g, tmp = flash(g, a-1, b-1)
		out += tmp
		g, tmp = flash(g, a-1, b)
		out += tmp
		g, tmp = flash(g, a-1, b+1)
		out += tmp
		g, tmp = flash(g, a, b-1)
		out += tmp
		g, tmp = flash(g, a, b+1)
		out += tmp
		g, tmp = flash(g, a+1, b-1)
		out += tmp
		g, tmp = flash(g, a+1, b)
		out += tmp
		g, tmp = flash(g, a+1, b+1)
		out += tmp
	}
	return g, out
}

func flashing(g *[10][10]int) (uint, [10][10]int) {
	grid := *g
	var out uint
Loop:
	for i, el1 := range grid {
		for j, el2 := range el1 {
			if el2 > 9 {
				var tmp uint
				grid, tmp = flash(grid, i, j)
				out, grid = flashing(&grid)
				out += tmp
				break Loop
			}
		}
	}
	return out, grid
}

func repairBoard(g [10][10]int) ([10][10]int, uint) {
	var sum uint
	for i, el1 := range g {
		for j, el2 := range el1 {
			if el2 == -1 {
				sum++
				g[i][j] = 0
			}
		}
	}
	return g, sum
}

func partOne(g *[10][10]int) uint {
	var out uint
	grid := *g
	var tmp uint
	for i := 0; ; i++ {
		grid = nextRound(&grid)
		tmp, grid = flashing(&grid)
		out += tmp
		grid, tmp = repairBoard(grid)
		if i == 99 {
			fmt.Println(out)
		}
		if tmp == 100 {
			return uint(i+1)
		}
	}
	return 0
}

func main() {
	octopuses := readGrid("assets/data.txt")
	fmt.Println(partOne(&octopuses))
}
