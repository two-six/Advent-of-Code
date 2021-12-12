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

func findPath(s string, p map[string][]string, vsc map[string]uint) uint {
	if s == "start" {
		var out uint
		for _, el := range p[s] {
			out += findPath(el, p, vsc)
		}
		return out
	} else if s == "end" {
		return 1
	} else {
		copyVSC := make(map[string]uint)
		for k, el := range vsc {
			copyVSC[k] = el
		}
		if IsLowerCase(s) {
			copyVSC[s]++
			if copyVSC[s] == 2 {
				for k, el := range copyVSC {
					if k != s && el > 1 {
						return 0
					}
				}
			} else if copyVSC[s] > 2 {
				return 0
			}
		}
		var out uint
		for _, el := range p[s] {
			if el != "start" {
				out += findPath(el, p, copyVSC)
			}
		}
		return out
	}
}

func partOne(p *map[string][]string) uint {
	return findPath("start", *p, make(map[string]uint))
}

func main() {
	paths := readPaths("assets/data.txt")
	fmt.Println(partOne(&paths))
}
