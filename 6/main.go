package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(s string) [9]uint {
	var out [9]uint
	file, err := os.Open(s)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	f := strings.Split(scanner.Text(), ",")
	for _, el := range f {
		tmp, err := strconv.Atoi(el)
		check(err)
		out[tmp]++
	}
	return out
}

func solve(fishes [9]uint, rep int) uint {
	for i := 0; i < rep; i++ {
		tmp := fishes[0]
		for j := uint(0); j < 8; j++ {
			fishes[j] = fishes[j+1]
		}
		fishes[6] += tmp
		fishes[8] = tmp
	}
	var sum uint
	for _, val := range fishes {
		sum += val
	}
	return sum
}

func main() {
	fishes := readData("assets/data.txt")
	fmt.Println("Silver: ", solve(fishes, 80))
	fmt.Println("Gold: ", solve(fishes, 256))
}
