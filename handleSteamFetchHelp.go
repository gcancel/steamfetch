package main

import "fmt"

func handleSteamFetchHelp(s *state, cmd command, cmds commands) error {
	help := cmds.descriptions["steamfetch"]
	fmt.Println(help)
	fmt.Println("")
	for name, desc := range cmds.descriptions {
		if name == "steamfetch" {
			continue
		} else {
			fmt.Printf("%s:	 %s\n", name, desc)
		}

	}
	return nil
}
