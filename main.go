package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/tateexon/primeInGameTime/lib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename>\n", os.Args[0])
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

	t1 := lib.FindFromOffsets(data, 44617)
	lib.PrintTime(t1, 1)
	t2 := lib.FindFromOffsets(data, 52137)
	lib.PrintTime(t2, 2)
	t3 := lib.FindFromOffsets(data, 59657)
	lib.PrintTime(t3, 3)
}
