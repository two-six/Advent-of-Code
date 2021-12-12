package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsLowerCase(s string) bool {
	return !(s == strings.ToUpper(string(s)))
}

func readPaths(s string) map[string][]string {
	out := make(map[string][]string)
	file, err := os.Open(s)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), "-")
		out[tmp[0]] = append(out[tmp[0]], tmp[1])
		out[tmp[1]] = append(out[tmp[1]], tmp[0])
	}
	return out
}

func findPath(s string, p *map[string][]string, vsc map[string]uint, partOne bool, done chan uint) {
	if s == "start" {
		var out uint
		d := make(chan uint, len(*p))
		for _, el := range (*p)[s] {
			go findPath(el, p, vsc, partOne, d)
			out += <-d
		}
		done <- out
	} else if s == "end" {
		done <- 1
	} else {
		copyVSC := make(map[string]uint)
		for k, el := range vsc {
			copyVSC[k] = el
		}
		if IsLowerCase(s) {
			copyVSC[s]++
			if copyVSC[s] == 2 {
				if partOne {
					done <- 0
					return
				}
				for k, el := range copyVSC {
					if k != s && el > 1 {
						done <- 0
						return
					}
				}
			} else if copyVSC[s] > 2 {
				done <- 0
				return
			}
		}
		var out uint
		d := make(chan uint, len(*p))
		for _, el := range (*p)[s] {
			if el != "start" {
				go findPath(el, p, copyVSC, partOne, d)
				out += <-d
			}
		}
		done <- out
		return
	}
}

func partOne(p *map[string][]string) uint {
	done := make(chan uint, 1)
	findPath("start", p, make(map[string]uint), true, done)
	return <-done
}

func partTwo(p *map[string][]string) uint {
	done := make(chan uint, 1)
	findPath("start", p, make(map[string]uint), false, done)
	return <-done
}

func main() {
	paths := readPaths("assets/data.txt")
	fmt.Println("Silver: ", partOne(&paths))
	fmt.Println("Gold: ", partTwo(&paths))
}
