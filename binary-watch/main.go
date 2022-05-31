package main

import (
	"fmt"
	"log"
	"math/bits"
)

func main() {
	for i := 0; i < 10; i++ {
		log.Printf("turnedOn: %d %v \n\n", i, readBinaryWatch(i))
	}
}

func readBinaryWatch(turnedOn int) []string {
	r := []string{}

	if turnedOn > 8 {
		return r
	}

	minutes := uint64(1<<6) - 1
	hours := (uint64(1<<4) - 1) << 6
	max := uint64(1<<10) - 1
	var i uint64
	for i = 0; i < max; i++ {
		if bits.OnesCount64(i) == turnedOn {
			m := i & minutes
			h := i & hours >> 6
			if m < 60 && h < 12 {
				r = append(r, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	// log.Printf("Minutes: %b, Hours: %b, Max: %b", minutes, hours, max)
	// log.Println(minutes, hours, max)

	return r
}
