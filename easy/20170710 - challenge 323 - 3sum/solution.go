package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func threeSum (slice []int) (outputSlice [][3]int) {
	sort.Ints(slice)
	outputMap := make(map[[3]int]struct{})

	for i := range slice[:len(slice)-3] {
		a := slice[i]
		start := i + 1
		end := len(slice) - 1

		for start < end {
			b := slice[start]
			c := slice[end]

			if a + b + c == 0 {
				test := [3]int{a, b, c}
				_,ok := outputMap[test]

				if ok == false {
					outputMap[test] = struct{}{}
				}
				end = end - 1
			} else if a + b + c >= 0 {
				end = end -1
			} else {
				start = start + 1
			}
		}
	}

	outputSlice = [][3]int{}
	for k := range outputMap {
		outputSlice = append(outputSlice, k)
	}

	return
}

func strToIntSlice (s string) (o []int) {
	stringSlice := strings.Fields(s)

	for _,elem := range stringSlice {
		conv,_ := strconv.Atoi(elem)
		o = append(o, conv)
	}

	return
}

func main() {
	test:= []string{
		"4 5 -1 -2 -7 2 -5 -3 -7 -3 1",
		"-1 -6 -3 -7 5 -8 2 -8 1",
		"-5 -1 -4 2 9 -9 -6 -1 -7",
	}

	for _,elemStr := range test {
		intSlice := strToIntSlice(elemStr)

		for _,elem := range threeSum(intSlice) {
			fmt.Println(strings.Trim(fmt.Sprintf("%v", elem), "[]"))
		}

		fmt.Printf("\n")
	}
}
