// package declaration
package main

import (
	"flag"
	"fmt"
)

func main() {

	parsename := flag.String("lump", "", "The demo file to be read through")

	flag.Parse()

	if *parsename == "" {

		fmt.Println("Please enter a valid file name")
		return

	}

	file := *parsename

	demobytes, err := readIntoByteSlice(file)
	if err != nil {
		return
	}
	vanillaDemoCheck(demobytes)

}
