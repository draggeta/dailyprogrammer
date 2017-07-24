package main

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"strconv"
	"errors"
)

/*
See the section 'More accurate rough estimation' on the following page:
http://www.hermes-press.com/square_root.htm

Provided that n » 1

1. Take the integer part of the number n. Z = int(n)
2. Count the number of digits in Z. Let D be the number of digits.
3. Calculate the value of 3.16 ^ D where the symbol " ^ " is to-the-n-power operator. Note: 3.16 ≈ √10
4. The estimate is the value obtain in step 3 multiple by an adjustment factor based on the first (left most) digit of the number n. E = ADJ * 3.16 ^ D

Table of ADJ

First Digit	ADJ
1 			0.32
2			0.45
3			0.55
4			0.63
5			0.71
6			0.78
7			0.84
8			0.90
9			0.95
*/

var adjCorrection = map[float64]float64 {
	1: 0.32,
	2: 0.45,
	3: 0.55,
	4: 0.63,
	5: 0.71,
	6: 0.78,
	7: 0.84,
	8: 0.90,
	9: 0.95,
}

func sqrEstimate (f float64) float64 {
	digitCount := math.Floor(math.Log10(f))
	div10Pow := f / (math.Pow(10, digitCount))
	firstDigit := math.Floor(math.Abs(div10Pow))
	return math.Pow(3.16, digitCount) * adjCorrection[firstDigit]
}

func sqrNewton (f float64, precision float64) (string, error) {
	var x float64
	if f > 100 {
		x = sqrEstimate(f)
	} else if f < 0 {
		return "0", errors.New("sqrNewton: square root of negative number")
	} else {
		x = 1
	}
	d := (x*x - f) / (2 * x)
	for math.Abs(d) > math.Pow(10, -precision) {
		x = x - d
		d = (x*x - f) / (2 * x)
	}
	return big.NewFloat(x).Text('f', int(precision)), nil
}

func main() {
	challenge := []string{
		"0 12345",
		"8 123456",
		"1 12345678901234567890123456789",
	}
	for _,s := range challenge {
		v := strings.Split(s, " ")
		precision, _ := strconv.ParseFloat(v[0],64)
		f, _ := strconv.ParseFloat(v[1], 64)
		if sqrt, err := sqrNewton(f, precision); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(sqrt)
		}
	}
}