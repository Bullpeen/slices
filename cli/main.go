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

	fmt.Println(slices.GetScores())
}