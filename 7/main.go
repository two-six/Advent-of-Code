package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

func check(e error) {
	if(e != nil) {
		panic(e)
	}
}

func main() {
	crabs := make(map[uint64]uint64)
	file, err := os.Open("assets/data.txt")
	check(err)
	defer file.Close()
	crabarrino, err := csv.NewReader(file).Read()
	check(err)
	file.Close()
	var max uint64
	for _, elem := range crabarrino {
		tmp, err := strconv.ParseUint(elem, 10, 32)
		check(err)
		crabs[tmp]++
		if max < tmp {
			max = tmp
		}
	}
	var min1, min2 uint64
	for i := uint64(0); i < max; i++ {
		tmp, tmp2 := uint64(0), uint64(0)
		for k, el := range crabs {
			if i == k {
				continue
			}
			var val uint64
			if k > i {
				val = k-i
			} else {
				val = i-k
			}
			tmp += val*el;
			tmp2 += ((val*val+val)/2)*el
		}
		if i == 0 {
			min1 = tmp
			min2 = tmp2
		} else if tmp < min1 {
			min1 = tmp
		} else if tmp2 < min2 {
			min2 = tmp2
		}
	}
	fmt.Println(min1)
	fmt.Println(min2)
}
