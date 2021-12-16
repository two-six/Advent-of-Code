package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type intPair struct {
	y, x int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func remove(slice []intPair, s int) []intPair {
	return append(slice[:s], slice[s+1:]...)
}

func readData(s string) [][]int {
	var out [][]int
	file, err := os.Open(s)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var i uint
	for scanner.Scan() {
		out = append(out, make([]int, 0, 0))
		for _, el := range scanner.Text() {
			tmp, err := strconv.Atoi(string(el))
			check(err)
			out[i] = append(out[i], int(tmp))
		}
		i++
	}
	return out
}

func visitNeighbours(bp [][]int, dist map[intPair]int, pos intPair) []intPair {
	var out []intPair
	y, x := pos.y, pos.x
	if y-1 >= 0 {
		if dist[intPair{y-1, x}] == 0  || dist[intPair{y-1, x}] > dist[pos] + bp[y-1][x] {
			dist[intPair{y-1, x}] = dist[pos] + bp[y-1][x]
			out = append(out, intPair{y-1, x})
		}
	}
	if x+1 < len(bp) {
		if dist[intPair{y, x+1}] == 0  || dist[intPair{y, x+1}] > dist[pos] + bp[y][x+1] {
			dist[intPair{y, x+1}] = dist[pos] + bp[y][x+1]
			out = append(out, intPair{y, x+1})
		}
	}
	if y+1 < len(bp) {
		if dist[intPair{y+1, x}] == 0  || dist[intPair{y+1, x}] > dist[pos] + bp[y+1][x] {
			dist[intPair{y+1, x}] = dist[pos] + bp[y+1][x]
			out = append(out, intPair{y+1, x})
		}
	}
	if x-1 >= 0 {
		if dist[intPair{y, x-1}] == 0  || dist[intPair{y, x-1}] > dist[pos] + bp[y][x-1] {
			dist[intPair{y, x-1}] = dist[pos] + bp[y][x-1]
			out = append(out, intPair{y, x-1})
		}
	}
	return out
}

func algo(bp [][]int) map[intPair]int {
	dist := make(map[intPair]int)
	var vis [][]bool
	for i := 0; i < len(bp); i++ {
		vis = append(vis, make([]bool, 0, 0))
		for j := 0; j < len(bp); j++ {
			vis[i] = append(vis[i], false)
		}
	}
	var tmp []intPair
	var emptyTMP []intPair
	tmp = visitNeighbours(bp, dist, intPair{0, 0})
	for true {
		var tmp2 [][]intPair
		tmp1 := make(map[intPair]int)
		for _, pos := range tmp {
			tmp2 = append(tmp2, visitNeighbours(bp, dist, pos))
			for _, el2 := range tmp2 {
				for _, el := range el2 {
					if tmp1[el] == 0 {
						tmp1[el] = 1
					}
				}
			}
		}

		tmp = emptyTMP
		for k := range tmp1 {
			tmp = append(tmp, k)
		}
		if len(tmp1) == 0 {
			return dist
		}
	}
	return dist
}

func silver(bp [][]int) int {
	return algo(bp)[intPair{len(bp)-1, len(bp)-1}]
}

func enlargeSlice(bp [][]int, n int) [][]int {
	siz := len(bp)
	for k := 0; k < n-1; k++ {
		for i := 0; i < siz; i++ {
			for j := 0; j < siz; j++ {
				tmp := bp[i][j+(k*siz)]+1
				if tmp > 9 {
					tmp = 1
				}
				bp[i] = append(bp[i], tmp)
			}
		}
	}
	for k := 0; k < n-1; k++ {
		for i := 0; i < siz; i++ {
			var tmp []int
			for j, el := range bp[i+(k*siz)] {
				tmp = append(tmp, el+1)
				if tmp[j] > 9 {
					tmp[j] = 1
				}
			}
			bp = append(bp, tmp)
		}
	}
	return bp
}

func gold(bp [][]int) int {
	tmp := enlargeSlice(bp, 5)
	return algo(tmp)[intPair{len(tmp)-1, len(tmp)-1}]
}

func main() {
	bp := readData("assets/data.txt")
	fmt.Println("Silver: ", silver(bp))
	fmt.Println("Gold: ", gold(bp))
}
