package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	if filepath.Ext(filename) != ".gci" {
		fmt.Println("Not a GCI file")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	t1 := findFromOffsets(data, 1, 44617)
	t2 := findFromOffsets(data, 2, 52137)
	t3 := findFromOffsets(data, 3, 59657)
}

func toHour(seconds float64) int {
	return int(seconds/60/60)
}

func toMinutes(seconds float64) int {
	return int(seconds/60)%60
}

func toSeconds(seconds float64) int {
	return int(seconds)%60
}

func fractionOfSecond(seconds float64) float64 {
	return (seconds - float64(int(seconds)))*1000
}

func printTime(time float64, fileNumber int) {
	fmt.Printf("Save %d: %.2f seconds: %02d:%02d:%02d.%3.0f\n", fileNumber, time, toHour(time), toMinutes(time), toSeconds(time), fractionOfSecond(time))
}

func findFromOffsets(data []byte, offset int, fileNumber int) {
	start := uint64(0)
	for i := 0; i <= offset; i++ {
		bit := (data[i/8] >> (7 - (i % 8))) & 1
		start = (start << 1) | uint64(bit)
		start = start & 0xFFFFFFFF_FFFFFFFF

		if i == offset {
			buf := make([]byte, 8)
			binary.BigEndian.PutUint64(buf, start)
			time := math.Float64frombits(binary.BigEndian.Uint64(buf))
			printTime(time, fileNumber)
			break
		}
	}
}