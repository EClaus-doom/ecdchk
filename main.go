// package declaration, keeping as just main as this setup is currently very simple, things might change down the line
package main

// bring in fmt for the console output and flag for CLI flags, other files work in same directory as the same package so no need to import at this time
import (
	"flag"
	"fmt"
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
	if version < 109 {

		vanillaDemoCheck(demobytes)
	}

	//switch cases to feed the slice into the right function so the demo info comes out right
	switch version {
	//vanilla, this is the standard version in use in modern times
	case 109:
		{

			vanillaDemoCheck(demobytes)
		}
	//boom, including older 201 and 202 versions for legacy demos, more to show proper versioning
	case 200:
		{
			boomDemoCheck(demobytes)
		}

	case 201:
		{
			boomDemoCheck(demobytes)
		}

	case 202:
		{
			boomDemoCheck(demobytes)
		}
	//mbf, including both 203 and even the Lee Killough blessed 204 which does exist, which shouldn't change the header
	case 203:
		{
			mbfDemoCheck(demobytes)
		}

	case 204:
		{
			mbfDemoCheck(demobytes)
		}
	//mbf21 is pretty straightforward since its a new standard
	case 221:
		{
			mbf21DemoCheck(demobytes)
		}
	//umapinfo with header before the usual info, I don't know how many of these exist as it seems
	//umapinfo generally doesn't add this on in the beginning of the header? I think it was just for that inital prboom um release?
	//definitely won't contain legacy demos though so those are omitted, and don't need mbf 2.04
	case 255:
		{
			version := demobytes[26]
			switch version {
			case 109:
				{
					vanillaDemoCheck(demobytes)
				}

			case 202:
				{

					boomDemoCheck(demobytes)
				}

			case 203:
				{
					mbfDemoCheck(demobytes)
				}

			case 221:
				{
					mbf21DemoCheck(demobytes)
				}
			}

		}
	}

}
