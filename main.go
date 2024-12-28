// package declaration, keeping as just main as this setup is currently very simple, things might change down the line
package main

// bring in fmt for the console output and flag for CLI flags, other files work in same directory as the same package so no need to import at this time
import (
	"flag"
	"fmt"
)

// enums for use in the various switch cases for readability, may be expanded later for future categories or added in specific demo sources
const (
	//stock vanilla version
	vanillaVer = 109
	//current Boom version
	boomOldVerOne = 200
	boomOldVerTwo = 201
	boomVer       = 202
	//MBF versions
	mbfVer      = 203
	mbfVerVogon = 204
	//MBF21 version
	mbfTwentyOne = 221
	//UMAPINFO
	umap = 255
	//skill in shorthand, I'm too young to die, Hey not to rough etc...
	itytd     = 0
	hntr      = 1
	hmp       = 2
	ultra     = 3
	nightmare = 4
	//multiplayer mode flags ex
	spCoop     = 0
	deathMatch = 1
	altDm      = 2
	//
)

func main() {

	//so far only lump exists for options, mainly set up if I feel like adding further functionality
	parsename := flag.String("lump", "", "The demo file to be read through")

	flag.Parse()

	//check for empty file and tell user if nothing is there
	if *parsename == "" {

		fmt.Println("Please enter a valid file name")
		return

	}

	//put the string pointer from the flag to a string that isn't a pointer for the file name
	file := *parsename

	//call the function that actually reads the file from the name found in its own source file, puts everything into a slice
	demobytes, err := readIntoByteSlice(file)
	if err != nil {
		return
	}

	version := demobytes[0]

	//initial vanilla check for old vanilla versions and feeding it in before hitting the switch cases
	if version < vanillaVer {

		vanillaDemoCheck(demobytes)
	}

	//switch cases to feed the slice into the right function so the demo info comes out right
	switch version {
	//vanilla, this is the standard version in use in modern times
	case vanillaVer:
		{

			vanillaDemoCheck(demobytes)
		}
	//boom, including older 201 and 202 versions for legacy demos, more to show proper versioning
	case boomOldVerOne:
		{
			boomDemoCheck(demobytes)
		}

	case boomOldVerTwo:
		{
			boomDemoCheck(demobytes)
		}

	case boomVer:
		{
			boomDemoCheck(demobytes)
		}
	//mbf, including both 203 and even the Lee Killough blessed 204 which does exist from the vogons forum, which shouldn't change the header
	case mbfVer:
		{
			mbfDemoCheck(demobytes)
		}

	case mbfVerVogon:
		{
			mbfDemoCheck(demobytes)
		}
	//mbf21 is pretty straightforward since its a new standard
	case mbfTwentyOne:
		{
			mbf21DemoCheck(demobytes)
		}
	//umapinfo with header before the usual info, I don't know how many of these exist as it seems
	//umapinfo generally doesn't add this on in the beginning of the header? I think it was just for that inital prboom um release?
	//definitely won't contain legacy demos though so those are omitted, and don't need mbf 2.04
	case umap:
		{
			version := demobytes[26]
			switch version {
			case vanillaVer:
				{
					vanillaDemoCheck(demobytes)
				}

			case boomVer:
				{

					boomDemoCheck(demobytes)
				}

			case mbfVer:
				{
					mbfDemoCheck(demobytes)
				}

			case mbfTwentyOne:
				{
					mbf21DemoCheck(demobytes)
				}
			}

		}
	}

}
