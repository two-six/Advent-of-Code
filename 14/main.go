package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type runePair struct {
	a, b rune
}

func extractPairs(s string) (map[runePair]uint64, map[rune]uint64) {
	out := make(map[runePair]uint64)
	outC := make(map[rune]uint64)
	for i := 0; i < len(s)-1; i++ {
		tmp := runePair{rune(s[i]), rune(s[i+1])}
		out[tmp]++
		outC[rune(s[i])]++
	}
	outC[rune(s[len(s)-1])]++
	return out, outC
}

func readData(s string) (map[runePair]uint64, [][3]rune, map[rune]uint64) {
	outStart := make(map[runePair]uint64)
	var outList [][3]rune
	outC := make(map[rune]uint64)
	file, err := os.Open(s)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		tmp := strings.Split(scanner.Text(), " -> ")
		if len(tmp) == 1 {
			outStart, outC = extractPairs(scanner.Text())
		} else {
			var tmp2 [3]rune
			tmp2[0] = rune(tmp[0][0])
			tmp2[1] = rune(tmp[0][1])
			tmp2[2] = rune(tmp[1][0])
			outList = append(outList, tmp2)
		}
	}
	return outStart, outList, outC
}

func solve(p map[runePair]uint64, l [][3]rune, c map[rune]uint64, rep int) uint64 {
	counting := make(map[rune]uint64)
	for k, el := range c {
		counting[k] = el
	}
	for i := 0; i < rep; i++ {
		tmpP := make(map[runePair]uint64)
		for k, val := range p {
			tmpP[k] = val
		}
		for k, val := range p {
			for _, el := range l {
				if k.a == el[0] && k.b == el[1] && val != 0 {
					tmp := runePair{el[0], el[1]}
					tmpP[tmp] -= val
					tmp = runePair{el[0], el[2]}
					tmpP[tmp] += val
					tmp = runePair{el[2], el[1]}
					tmpP[tmp] += val
					counting[el[2]] += val
					break
				}
			}
		}
		p = tmpP
	}
	var max, min uint64
	for _, val := range counting {
		if min == 0 {
			min = val
		}
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}
	return max-min
}

func main() {
	pairs, list, c := readData("assets/data.txt")
	fmt.Println("Silver: ", solve(pairs, list, c, 10))
	fmt.Println("Gold: ", solve(pairs, list, c, 40))
}
