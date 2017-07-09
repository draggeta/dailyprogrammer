package main

import (
	"fmt"
)

func testJolly (a []int) ([]int, string) {
	var l = len(a) - 1
	var c = make(map[int]bool)
	if l == 0 {
		return a, "JOLLY"
	} else if l >= 0 {
		for i := 0; i < l; i++  {
			c[abs(a[i] - a[i + 1])] = true
		}
		if len(c) == l {
			for j := 1; j <= l; j++ {
				if !c[j] {
					return a, "NOT JOLLY"
				}
			}
		} else {
			return a, "NOT JOLLY"
		}
	}
	return a, "JOLLY"
}


func abs (i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}

func main () {
	t1 := []int{1,4,2,3}
	fmt.Println(testJolly(t1))
	t2 := []int{1,4,2,-1,6}
	fmt.Println(testJolly(t2))
	t3 := []int{19,22,24,21}
	fmt.Println(testJolly(t3))
	t4 := []int{19,22,24,25}
	fmt.Println(testJolly(t4))
	t5 := []int{2,-1,0,2}
	fmt.Println(testJolly(t5))
}
