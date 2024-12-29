// package declaration, keeping as just main as this setup is currently very simple, things might change down the line
package main

// bring in fmt for the console output and flag for CLI flags, other files work in same directory as the same package so no need to import at this time
import ()

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

	cliFlow()
}
