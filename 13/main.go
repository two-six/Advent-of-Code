package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type point struct {
	x int
	y int
}

type board struct {
	points []point
	sizeX int
	sizeY int
}

type fold struct {
	x 	bool
	val int
}

func foldBoard(f fold, b board) board {
	if f.x {
		var first, second board
		for _, el := range b.points {
			if el.x < f.val {
				first.points = append(first.points, el)
			} else {
				second.points = append(second.points, el)
			}
		}
		for _, el := range second.points {
			el.x -= f.val*2
			if el.x < 0 {
				el.x = -el.x
			}
			first.points = append(first.points, el)
		}
		first.sizeX = f.val-1
		first.sizeY = b.sizeY
		return removeDuplicates(first)
	} else {
		var first, second board
		for _, el := range b.points {
			if el.y < f.val {
				first.points = append(first.points, el)
			} else {
				second.points = append(second.points, el)
			}
		}
		for _, el := range second.points {
			el.y -= f.val*2
			if el.y < 0 {
				el.y = -el.y
			}
			first.points = append(first.points, el)
		}
		first.sizeX = b.sizeX
		first.sizeY = f.val-1
		return removeDuplicates(first)
	}
}

func removeDuplicates(b board) board {
	keys := make(map[point]bool)
	var out board
	for _, el := range b.points {
		if _, val := keys[el]; !val {
			keys[el] = true
			out.points = append(out.points, el)
		}
	}
	out.sizeX = b.sizeX
	out.sizeY = b.sizeY
	return out
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPointsFolds(s string) (board, []fold) {
	var outPoint []point
	var outFold []fold
	file, err := os.Open(s)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var sx, sy int
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), ",")
		if len(tmp) == 2 {
			tmp2, e := strconv.ParseInt(tmp[0], 10, 32)
			check(e)
			tmp3, e2 := strconv.ParseInt(tmp[1], 10, 32)
			check(e2)
			if int(tmp2) > sx {
				sx = int(tmp2)
			}
			if int(tmp3) > sy {
				sy = int(tmp3)
			}
			outPoint = append(outPoint, point{int(tmp2), int(tmp3)})
		} else if len(scanner.Text()) == 0 {
			continue
		} else {
			tmp := strings.Split(strings.Split(scanner.Text(), " ")[2], "=")
			tmp2, e := strconv.ParseUint(tmp[1], 10, 32)
			check(e)
			outFold = append(outFold, fold{tmp[0] == "x", int(tmp2)})
		}
	}
	return board{outPoint, sx, sy}, outFold
}

func partOne(b board, f *[]fold) uint {
	b = foldBoard((*f)[0], b)
	return uint(len(b.points))
}

func partTwo(b board, f *[]fold) {
	for _, el := range *f {
		b = foldBoard(el, b)
	}
	for i := 0; i <= b.sizeY; i++ {
		fmt.Printf("\n")
	XCon:
		for j := 0; j <= b.sizeX; j++ {
			for _, el := range b.points {
				if el.x == j && el.y == i {
					fmt.Printf("#")
					continue XCon
				}
			}
			fmt.Printf(" ")
		}
	}
	fmt.Println()
}

func main() {
	board, folds := getPointsFolds("assets/data.txt")
	fmt.Println("Silver: ", partOne(board, &folds))
	fmt.Println("Gold: ")
	partTwo(board, &folds)
}
