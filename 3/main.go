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
	var gamma, epsilon string
	var lines [2000]string
	count := [12]struct{one uint; zero uint}{}
	tmp_n := 0
	for scanner.Scan() {
		chars := []rune(scanner.Text())
		for i, r := range chars {
			if r == '1' {
				count[i].one++;
			} else {
				count[i].zero++;
			}
		}
		lines[tmp_n] = scanner.Text()
		tmp_n++
	}
	file.Close()

	// part 1
	for _, x := range count {
		if(x.zero > x.one) {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	var e error
	part_one := struct{gamma uint64; epsilon uint64} {}
	part_one.gamma, e = strconv.ParseUint(gamma, 2, 32)
	check(e)
	part_one.epsilon, e = strconv.ParseUint(epsilon, 2, 32)
	check(e)
	fmt.Println("Part 1:", part_one.gamma * part_one.epsilon)

	// part 2
}
