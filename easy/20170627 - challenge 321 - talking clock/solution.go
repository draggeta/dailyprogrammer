package	main

import (
	"fmt"
	//"math"
	"strconv"
	"strings"
	//"time"
)

func strToTime (s string) (string) {
	splitTime := strings.Split(s, ":")
	var err error
	var h int
	var hour string
	var minute []string
	var ampm string
	if h,err = strconv.Atoi(splitTime[0]); err == nil {
		switch h % 12 {
			case 0: hour = "twelve"
			case 1: hour = "one"
			case 2: hour = "two"
			case 3: hour = "three"
			case 4: hour = "four"
			case 5: hour = "five"
			case 6: hour = "six"
			case 7: hour = "seven"
			case 8: hour = "eight"
			case 9: hour = "nine"
			case 10: hour = "ten"
			case 11: hour = "eleven"
		}
	}
	if m,err := strconv.Atoi(splitTime[1]); err == nil {
		switch m {
			case 0:
			case 10: minute = append(minute, "ten")
			case 11: minute = append(minute, "eleven")
			case 12: minute = append(minute, "twelve")
			case 13: minute = append(minute, "thirteen")
			case 15: minute = append(minute, "fifteen")
			case 16: minute = append(minute, "sixteen")
			case 17: minute = append(minute, "seventeen")
			case 18: minute = append(minute, "eighteen")
			case 19: minute = append(minute, "nineteen")
		}
		if m < 10 || m >= 20 {
			switch {
				case m > 0 && m < 10: minute = append(minute, "oh")
				case m >= 20 && m < 30: minute = append(minute, "twenty")
				case m >= 30 && m < 40: minute = append(minute, "thirty")
				case m >= 40 && m < 50: minute = append(minute, "forty")
				case m >= 50 && m < 60: minute = append(minute, "fifty")
			}
			switch m % 10 {
				case 1: minute = append(minute, "one")
				case 2: minute = append(minute, "two")
				case 3: minute = append(minute, "three")
				case 4: minute = append(minute, "four")
				case 5: minute = append(minute, "five")
				case 6: minute = append(minute, "six")
				case 7: minute = append(minute, "seven")
				case 8: minute = append(minute, "eight")
				case 9: minute = append(minute, "nine")
			}
		}
	}
	if h >= 12 {
		ampm = "PM"
	} else {
		ampm = "AM"
	}
	strMinutes := strings.Join(minute," ")
	var strSpace string
	if cap(minute) > 0 {
		strSpace = " "
	}
	return fmt.Sprintf("It's %s %s%s%s\n", hour, strMinutes, strSpace, ampm)
}

func main() {
	test :=[6]string{"00:00","01:30","12:05","14:01","20:29","21:00"}
	for index := range(test) {
		fmt.Printf(strToTime(test[index]))
	}
}