// package declaration
package main

//imports, should be just io, os, and fmt for this little guy

func main() {

	demobytes, err := readIntoByteSlice("demo.lmp")
	if err != nil {
		return
	}
	vanillaDemoCheck(demobytes)

}
