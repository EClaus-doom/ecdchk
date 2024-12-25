package main

import (
	"fmt"
)

// mbf demo function for displaying info about MBF demos
func mbfDemoCheck(slice []byte) {
	//put slice into header
	header := slice
	//assigned variables for various flags from the demo header bytes
	version := header[0]
	skill := header[8]
	episode := header[9]
	mapnum := header[10]
	multimode := header[11]
	respawn := header[19]
	fast := header[20]
	nomo := header[21]

	//show header version
	fmt.Println("")
	fmt.Println("MBF DEMO HEADER")
	fmt.Println("")

	switch version {
	case 203:
		fmt.Println("Version: MBF 2.03")
	default:
		fmt.Println("Version: Other MBF Version", version)
	}

	switch skill {
	case 0:
		fmt.Println("Skill: 1 - I'm Too Young to Die!")
	case 1:
		fmt.Println("Skill: 2 - Hey, Not Too Rough!")
	case 2:
		fmt.Println("Skill: 3 - Hurt Me Plenty!")
	case 3:
		fmt.Println("Skill: 4 - Ultraviolence!")
	case 4:
		fmt.Println("Skill: 5 - Nightmare!")

	}

	fmt.Println("Episode: ", episode)

	fmt.Println("Map: ", mapnum)

	switch multimode {
	case 0:
		fmt.Println("Multiplayer Mode: Single Player or Co-op")
	case 1:
		fmt.Println("Multiplayer Mode: Deathmatch")
	case 2:
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

func mbfDemoUmapInfo(slice []byte) {
	//put slice into header
	header := slice
	//
	version := header[26]
	skill := header[34]
	episode := header[35]
	mapnum := header[36]
	multimode := header[37]
	respawn := header[45]
	fast := header[46]
	nomo := header[47]

	//show header version
	fmt.Println("")
	fmt.Println("MBF DEMO HEADER")
	fmt.Println("")

	switch version {
	case 203:
		fmt.Println("Version: MBF 2.03")
	default:
		fmt.Println("Version: Other MBF Version", version)
	}

	switch skill {
	case 0:
		fmt.Println("Skill: 1 - I'm Too Young to Die!")
	case 1:
		fmt.Println("Skill: 2 - Hey, Not Too Rough!")
	case 2:
		fmt.Println("Skill: 3 - Hurt Me Plenty!")
	case 3:
		fmt.Println("Skill: 4 - Ultraviolence!")
	case 4:
		fmt.Println("Skill: 5 - Nightmare!")

	}

	fmt.Println("Episode: ", episode)

	fmt.Println("Map: ", mapnum)

	switch multimode {
	case 0:
		fmt.Println("Multiplayer Mode: Single Player or Co-op")
	case 1:
		fmt.Println("Multiplayer Mode: Deathmatch")
	case 2:
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
