package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func main2() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	findOffset(data, 0, 2)
	findOffset(data, 0, 40)
	findOffset(data, 0, 11)
}

// pass n the hours and minutes of your save to find the offset, it is not exact and may return multiple values
func findOffset(data []byte, hours, minutes int) {
	min := float64(hours*60*60 + minutes*60)
	max := float64(hours*60*60 + minutes*60 + 60)
	fmt.Printf("Searching between %.0f and %.0f seconds\n", min, max)

	start := uint64(0)
	for i := 0; i < len(data)*8; i++ {
		bit := (data[i/8] >> (7 - (i % 8))) & 1
		start = (start << 1) | uint64(bit)
		start = start & 0xFFFFFFFF_FFFFFFFF

		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, start)
		double := math.Float64frombits(binary.BigEndian.Uint64(buf))

		if min < double && double < max {
			offset := i - 64
			fmt.Printf("offset: %d\n", i)
			fmt.Printf("%d %.6f %d.%d %x\n", offset, double, int(offset/8), offset%8, buf)
		}
	}
}
