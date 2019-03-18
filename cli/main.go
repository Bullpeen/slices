package main

import (
	"fmt"
	"github.com/Bullpeen/slices"
)

func main() {
	//user, err := slices.GetUser("76561197974723967")
	//if err == nil {
	//	fmt.Printf("%s: %f\n", user.UserName, user.TotalPP)
	//} else {
	//	fmt.Println(err)
	//}

	slices.Register([]string { "76561197976367183", "76561198057633471", "76561198002272597", "76561197974723967", "76561197969022064","76561197980107683"})

	fmt.Println(slices.GetScores())
}
