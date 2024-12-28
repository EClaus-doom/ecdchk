package main

import "fmt"

func mbf21DemoCheck(slice []byte) {
	//put into another in function variable called header
	header := slice
	//assigned variables for various flags from the demo header bytes
	version := header[0]
	skill := header[7]
	episode := header[8]
	mapnum := header[9]
	multimode := header[10]
	respawn := header[15]
	fast := header[16]
	nomo := header[17]

	//show header version
	fmt.Println("")
	fmt.Println("MBF21 DEMO HEADER")
	fmt.Println("")

	switch version {
	case mbfTwentyOne:
		fmt.Println("Version: MBF21")
	default:
		fmt.Println("Not MBF 21", version)
	}

	switch skill {
	case itytd:
		fmt.Println("Skill: 1 - I'm Too Young to Die!")
	case hntr:
		fmt.Println("Skill: 2 - Hey, Not Too Rough!")
	case hmp:
		fmt.Println("Skill: 3 - Hurt Me Plenty!")
	case ultra:
		fmt.Println("Skill: 4 - Ultraviolence!")
	case nightmare:
		fmt.Println("Skill: 5 - Nightmare!")

	}

	fmt.Println("Episode: ", episode)

	fmt.Println("Map: ", mapnum)

	switch multimode {
	case spCoop:
		fmt.Println("Multiplayer Mode: Single Player or Co-op")
	case deathMatch:
		fmt.Println("Multiplayer Mode: Deathmatch")
	case altDm:
		fmt.Println("Multiplayer Mode: Alt-Death")
	}

	if respawn > 0 {
		fmt.Println("Respawn: On")
	} else {
		fmt.Println("Respawn: Off")
	}

	if fast > 0 {
		fmt.Println("Fast: On")
	} else {
		fmt.Println("Fast: Off")
	}

	if nomo > 0 {
		fmt.Println("No Monsters: On")
	} else {
		fmt.Println("No Monsters: Off")
	}

	fmt.Println("")
	fmt.Println("END HEADER")
	fmt.Println("")
}
