package main

import (
    "fmt"
    "strings"
)

var lookupmap = map[string]string{
    "A": "4",
    "B": "6",
    "E": "3",
    "I": "1",
    "L": "|",
    "M": "(V)",
    "N": "(\\)",
    "O": "0",
    "S": "5",
    "T": "7",
    "V": "\\/",
    "W": "`//",
}

func toleet(s string) {
    original := s
    output   := strings.ToUpper(s)
    for key, value := range lookupmap {
        output = strings.Replace(output, key, value, -1)
    }
    fmt.Printf("%s -> %s\n", original, output)
}

func fromleet(s string) (original, output string) {
    original = s
    output   = strings.ToUpper(s)
    for key, value := range lookupmap {
        output = strings.Replace(output, value, key, -1)
    }
    return
}

func main() {
    toleet("Eye need help!")
    o,e := fromleet(`3Y3 (\)33d j00 t0 g37 d4 d0c70r.`)
	fmt.Printf("%s -> %s\n", o, e)
}