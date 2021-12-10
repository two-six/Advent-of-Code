package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	"sort"
)

type stack []rune

func(s stack) push(r rune) stack {
	return append(s, r)
}

func(s stack) pop() (stack, rune, error) {
	l := len(s)
	if l == 0 {
		return nil, 0, errors.New("Stack is empty!")
	}
	return s[:l-1], s[l-1], nil

}

func check(e error) {
	if(nil != e) {
		panic(e)
	}
}

func getLines(s string) []string {
	out := make([]string, 0, 10)
	file, err := os.Open(s)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}

func illegalRune(l string) (rune, stack) {
	var s stack
	for _, r := range []rune(string(l[:])) {
		switch r {
		case '(', '[', '{', '<':
			s = s.push(r)
		case ')', ']', '}', '>':
			var tmp rune
			var e error
			s, tmp, e = s.pop()
			if e != nil {
				return r, nil
			}
			switch tmp {
			case '(':
				if r != ')' {
					return r, nil
				}
			case '[':
				if r != ']' {
					return r, nil
				}
			case '{':
				if r != '}' {
					return r, nil
				}
			case '<':
				if r != '>' {
					return r, nil
				}
			}
		}
	}
	return 0, s
}

func countMissingParts(s stack) uint64 {
	var sum uint64
	for true {
		var r rune
		var e error
		s, r, e = s.pop()
		if e != nil {
			return sum
		}
		sum *= 5
		switch r {
		case '(':
			sum += 1
		case '[':
			sum += 2
		case '{':
			sum += 3
		case '<':
			sum += 4
		}
	}
	return 0
}

func parts(l *[]string) (uint, uint64) {
	type ps struct {
		a uint
		b uint64
	}
	done := make(chan ps, len(*l))
	for _, line := range *l {
		go func(line string) {
			tmp, s := illegalRune(line)
			switch tmp {
			case ')':
				done <- ps{3, 0}
			case ']':
				done <- ps{57, 0}
			case '}':
				done <- ps{1197, 0}
			case '>':
				done <- ps{25137, 0}
			case 0:
				done <- ps{0, countMissingParts(s)}
			}
		}(line)
	}
	var sum uint
	middle := make([]uint64, 0, len(*l))
	for i := 0; i < len(*l); i++ {
		tmp := <-done
		sum += tmp.a
		if tmp.b != 0 {
			middle = append(middle, tmp.b)
		}
	}
	sort.Slice(middle, func(i, j int) bool {
		return middle[i] < middle[j]
	})
	return sum, middle[len(middle)/2]
}

func main() {
	l := getLines("assets/data.txt")
	fmt.Println(parts(&l))
}
