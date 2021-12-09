package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var capacityX, capacityY int

func check(e error) {
	if(nil != e) {
		panic(e)
	}
}

func getGrid(s string) [][]uint {
	file, err := os.Open(s)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var i uint
	out := make([][]uint, 0)
	for scanner.Scan() {
		out = append(out, make([]uint, 0))
		tmp := scanner.Text()
		capacityY++
		for _, r := range []rune(string(tmp[:])) {
			tmpUint, err := strconv.ParseUint(string(r), 10, 32)
			check(err)
			out[i] = append(out[i], uint(tmpUint))
		}
		i++
	}
	capacityX = len(out[0])
	return out
}

func sumOfTheRiskPoints(grid [][]uint) uint {
	const (
		up = iota
		left
		down
		right
	)

	done := make(chan uint, len(grid))
	for i, l := range grid {
		go func(i int, l []uint) {
			var sum uint
			for j := 0; j < len(l); j++ {
				var dir [4]uint
				if i == 0 {
					dir[up] = 10
				} else {
					dir[up] = grid[i-1][j]
				}
				if i == len(grid)-1 {
					dir[down] = 10
				} else {
					dir[down] = grid[i+1][j]
				}
				if j == 0 {
					dir[left] = 10
				} else {
					dir[left] = grid[i][j-1]
				}
				if j == len(grid[i])-1 {
					dir[right] = 10
				} else {
					dir[right] = grid[i][j+1]
				}
				min := uint(10)
				for _, n := range dir {
					if n < min {
						min = n
					}
				}
				if grid[i][j] < min {
					sum += grid[i][j]+1
				}
			}
			done <- sum
		}(i, l)
	}
	var out uint
	for i := 0; i < len(grid); i++ {
		out += <-done
	}
	return out
}

type coor struct {
	x int
	y int
}

func findPoints(grid *[][]uint) []coor {
	var out []coor
	for indX, i := range *grid {
		for indY, j := range i {
			if j != 9 {
				out = append(out, coor{indX, indY})
			}
		}
	}
	return out
}

func findBasins(mp map[coor]uint, c coor) uint {
	sum := uint(1)
	mp[c] = 1
	if a, b := mp[coor{c.x-1, c.y}]; b {
		if a == 0 {
			sum += findBasins(mp, coor{c.x-1, c.y})
		}
	}
	if a, b := mp[coor{c.x+1, c.y}]; b {
		if a == 0 {
			sum += findBasins(mp, coor{c.x+1, c.y})
		}
	}
	if a, b := mp[coor{c.x, c.y-1}]; b {
		if a == 0 {
			sum += findBasins(mp, coor{c.x, c.y-1})
		}
	}
	if a, b := mp[coor{c.x, c.y+1}]; b {
		if a == 0 {
			sum += findBasins(mp, coor{c.x, c.y+1})
		}
	}
	return sum
}

func multOfTheBasins(grid *[][]uint) uint {
	p := findPoints(grid)
	mp := make(map[coor]uint)
	var basSize []int
	for _, el := range p {
		mp[el] = 0
	}
	for k, el := range mp {
		if el == 0 {
			basSize = append(basSize, int(findBasins(mp, k)))
		}
	}

	sort.Slice(basSize, func(i, j int) bool {
		return basSize[i] < basSize[j]
	})

	s := len(basSize)-1
	out := basSize[s]*basSize[s-1]*basSize[s-2]

	return uint(out)
}

func partOne(grid *[][]uint) uint {
	return sumOfTheRiskPoints(*grid)
}

func partTwo(grid *[][]uint) uint {
	return multOfTheBasins(grid)
}

func main() {
	grid := getGrid("assets/data.txt")
	fmt.Println(partOne(&grid))
	fmt.Println(partTwo(&grid))
}
