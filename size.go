package main

import (
	"strconv"
)

type Size int64

var suffix = [...]string{"Bytes", "KBytes", "MBytes", "GBytes", "TBytes"}

func (s Size) String() string {
	i := 0
	for int64(s) > 50000 && i < len(suffix)-1 {
		s = s / 1000
		i++
	}
	return strconv.FormatInt(int64(s), 10) + " " + suffix[i]
}
