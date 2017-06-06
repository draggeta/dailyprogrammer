package	main

import (
	"fmt"
)

var tag = map[string]string {
	"a": "bc",
	"b": "a",
	"c": "aaa",
}



func cts(s string) {
	for len(s) > 1 {
		s = s[2:] + tag[s[:1]]
		fmt.Println(s)
	}
}

func main() {
	x := "aaaaaaa"
	cts(x)
}