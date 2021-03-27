package main

import "fmt"

func main() {
	/*
	 * Arrrays examples
	 */

	// declaring an empty array of strings
	var days []string

	// declaring an array with elements
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	fmt.Println(days[0]) // prints 'monday'
	fmt.Println(days[5]) // prints 'saturday'

	/*
	 * Slices exmamples
	 */

	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekdays := days[0:5]
	fmt.Println(weekdays)
	// This returns: [Monday Tuesday Wednesday Thursday Friday]

	/*
	 * Maps examples
	 */

	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}

	fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350

	/*
	 * Struct examples
	 */

	type Person struct {
		name string
		age  int
	}

	// our Team struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty 'Team'
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forrest"}, Person{name: "Gordon"}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}
