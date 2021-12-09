package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	a = 1 << iota
	b
	c
	d
	e
	f
	g
)

const (
	zero = a+b+c+e+f+g
	one = c+f
	two = a+c+d+e+g
	three = a+c+d+f+g
	four = b+c+d+f
	five = a+b+d+f+g
	six = a+b+d+e+f+g
	seven = a+c+f
	eight = a+b+c+d+e+f+g
	nine = a+b+c+d+f+g
)

type line struct {
	signals []string
	output []string
}

func check(e error) {
	if(nil != e) {
		panic(e)
	}
}

func read_values(s string) []line {
	var lines []line
	file, err := os.Open(s)
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		out := strings.Split(strings.Split(scanner.Text(), " | ")[1], " ")
		sig := strings.Split(strings.Split(scanner.Text(), " | ")[0], " ")
		lines = append(lines, line{sig, out})
	}
	return lines
}

func part_one(lines *[]line) uint {
	var sum uint
	for _, l := range *lines {
		for _, el := range l.output {
			if (len(el) >= 2 && len(el) <= 4) || len(el) == 7 {
				sum++
			}
		}
	}
	return sum
}

func sort_str(s string) string {
	tmp := []rune(string(s[:]))
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	return string(tmp)
}

func sort_rune(a []rune) []rune {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}

func find_uniq(a []rune, b []rune) []rune {
	tmp := make(map[rune]uint)
	for _, la := range a {
		tmp[la]++
	}
	for _, lb := range b {
		tmp[lb]++
	}

	var out []rune
	for k, el := range tmp {
		if el == 1 {
			out = append(out, k)
		}
	}

	return out
}

func decipher(s []string) [7]struct{a rune; b uint} {
	out := [7]struct{a rune; b uint}{}
	for i, l := range s {
		s[i] =  sort_str(l)
	}
	var n [10][]rune

	var tmp5 [3][]rune
	var tmp6 [3][]rune

	for _, l := range s {
		switch len(l) {
		case 2:
			n[1] = []rune(string(l[:]))
		case 3:
			n[7] = []rune(string(l[:]))
		case 4:
			n[4] = []rune(string(l[:]))
		case 5:
			if len(tmp5[0]) == 0 {
				tmp5[0] = []rune(string(l[:]))
			} else if len(tmp5[1]) == 0 {
				tmp5[1] = []rune(string(l[:]))
			} else {
				tmp5[2] = []rune(string(l[:]))
			}
		case 6:
			if len(tmp6[0]) == 0 {
				tmp6[0] = []rune(string(l[:]))
			} else if len(tmp6[1]) == 0 {
				tmp6[1] = []rune(string(l[:]))
			} else {
				tmp6[2] = []rune(string(l[:]))
			}
		case 7:
			n[8] = []rune(string(l[:]))
		}
	}

	// find a
	out[0] = struct{a rune; b uint}{find_uniq(n[1], n[7])[0], a}

	// find e
	for _, el5 := range tmp5 {
		if len(find_uniq(n[4], el5)) == 3 {
			if len(find_uniq(el5, n[1])) == 5 {
				n[5] = el5
				for _, el6 := range tmp6 {
					if len(find_uniq(n[1], el6)) == 6 {
						n[6] = el6
						out[4] = struct{a rune; b uint}{find_uniq(n[5], n[6])[0], e}
						break
					}
				}
				break
			}
		}
	}
	// find c
	for _, el6 := range tmp6 {
		if len(find_uniq(n[5], el6)) == 1 {
			if find_uniq(n[5], el6)[0] != out[4].a {
				n[9] = el6
				out[2] = struct{a rune; b uint}{find_uniq(n[5], n[9])[0], c}
				break
			}
		}
	}
	// find f
	for _, el2 := range n[1] {
		if el2 != out[2].a {
			out[5] = struct{a rune; b uint}{el2, f}
			break
		}
	}

	// find g
	for _, el := range find_uniq(n[9], n[4]){
		if el != out[0].a {
			out[6] = struct{a rune; b uint}{el, g}
			break
		}
	}

	// find d
	for _, el5 := range tmp5 {
		if len(find_uniq(n[1], el5)) == 3 {
			n[3] = el5
			break
		}
	}
	for _, r := range n[3] {
		if (r != out[0].a && r != out[2].a) && (r != out[5].a && r != out[6].a) {
			out[3] = struct{a rune; b uint}{r, d}
			break
		}
	}

	// find b
	for _, r := range n[5] {
		if (r != out[0].a && r != out[3].a) && (r != out[5].a && r != out[6].a) {
			out[1] = struct{a rune; b uint}{r, b}
			break
		}
	}


	return out
}

func out_number(lines *line, sol *[7]struct{a rune; b uint}) uint {
	var sum string
	for _, n := range lines.output {
		number := uint(0)
		for _, r := range []rune(string(n[:])) {
			for _, v := range sol {
				if v.a == r {
					number += v.b
				}
			}
		}
		switch number {
		case zero:
			sum += "0"
		case one:
			sum += "1"
		case two:
			sum += "2"
		case three:
			sum += "3"
		case four:
			sum += "4"
		case five:
			sum += "5"
		case six:
			sum += "6"
		case seven:
			sum += "7"
		case eight:
			sum += "8"
		case nine:
			sum += "9"
		}
	}
	out, err := strconv.ParseUint(sum, 10, 32)
	check(err)

	return uint(out)
}

func part_two(lines *[]line) uint {
	done := make(chan uint, len(*lines))
	var sum uint
	for _, l := range *lines {
		go func(l line) {
			sol := decipher(l.signals)
			done <- out_number(&l, &sol)
		}(l)
	}
	for i := 0; i < len(*lines); i++ {
		sum += <-done
	}

	return sum
}

func main() {
	lines := read_values("assets/big-boy.txt")
	fmt.Println("Part 1: ", part_one(&lines))
	fmt.Println("Part 2: ", part_two(&lines))
}
