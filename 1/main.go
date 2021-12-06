package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if(e != nil) {
		panic(e)
	}
}

func main() {
	file, err := os.Open("assets/data.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var depth []int32
	for scanner.Scan() {
		tmp, err := strconv.Atoi(scanner.Text())
		check(err)
		depth = append(depth, int32(tmp))
	}
	file.Close()
	sum := 0

	// Part 1
	for i := 1; i < len(depth); i++ {
		if(depth[i-1] < depth[i]) {
			sum++
		}
	}
	fmt.Println(sum)

	// Part 2
	sum = 0
	for i := 0; i < len(depth)-3; i++ {
		if(depth[i] < depth[i+3]) {
			sum++
		}
	}
	fmt.Println(sum)
}
