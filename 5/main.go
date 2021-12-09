package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func check(e error) {
	if(e != nil) {
		panic(e)
	}
}

func main() {
	points := make(map[struct{a, b uint64}]uint)

	file, err := os.Open("assets/data.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tmp := scanner.Text()
		tmp = strings.ReplaceAll(tmp, " -> ", ",")
		coordinates_s := strings.Split(tmp, ",")
		var coordinates [4]uint64
		for i, elem := range coordinates_s {
			coordinates[i], err = strconv.ParseUint(elem, 10, 32)
			check(err)
		}
		if coordinates[0] == coordinates[2] {
			var y1, y2 uint64
			if(coordinates[1] > coordinates[3]) {
				y1 = coordinates[1]
				y2 = coordinates[3]
			} else {
				y1 = coordinates[3]
				y2 = coordinates[1]
			}
			for i := uint64(0); i <= y1-y2; i++ {
				points[struct{a, b uint64}{coordinates[0], y2+i}]++
			}
		} else {
			var x1, x2 uint64
			if(coordinates[0] > coordinates[2]) {
				x1 = coordinates[0]
				x2 = coordinates[2]
			} else {
				x1 = coordinates[2]
				x2 = coordinates[0]
			}
			for i := uint64(0); i <= x1-x2; i++ {
				points[struct{a, b uint64}{x2+i, coordinates[1]}]++
			}
		}
	}
	file.Close()
	var sum uint
	for k, elem := range points {
		fmt.Println(k, elem)
		if elem >= 2 {
			sum++
		}
	}
	fmt.Println(sum)
}
