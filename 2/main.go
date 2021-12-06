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
	file, err := os.Open("assets/data.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	part_one := struct{x int; depth int}{}
	part_two := struct{x int; depth int; aim int}{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		val, err :=  strconv.Atoi(parts[1])
		check(err)
		switch parts[0] {
		case "forward":
			part_one.x += val
			part_two.x += val
			part_two.depth += part_two.aim*val
		case "down":
			part_one.depth += val
			part_two.aim += val
		case "up":
			part_one.depth -= val
			part_two.aim -= val
		}
	}
	file.Close()

	fmt.Println(part_one.x * part_one.depth)
	fmt.Println(part_two.x * part_two.depth)
}
