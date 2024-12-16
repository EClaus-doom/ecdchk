// package declaration
package ecgodemocheck

//imports, should be just io, os, and fmt for this little guy
import ()

func main() {

	demobytes, err := readIntoByteSlice("demo.lmp")
	if err != nil {
		return
	}
	vanillaDemoCheck(demobytes)

}
