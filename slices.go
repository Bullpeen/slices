package slices

import (
	"context"

	log "github.com/Sirupsen/logrus"
	"github.com/jirwin/quadlek/quadlek"
	"sort"
	"fmt"
)

var users = []string { "1841289615922336", "76561197976367183", "76561198057633471", "76561198002272597", "76561197974723967", "76561197969022064"}

func GetScores() (string) {

	var output []User

	for _, uid := range users {
		uu, err := GetUser(uid)

		if err == nil {
			output = append(output, *uu)
		} else {
			log.Info("Error fetching user: %s", err)
		}
	}

	sort.Sort(ByTotalPP(output))

	var outStr = ""

	for _, user := range output {
		if len(outStr) != 0 {
			outStr += "\n"
		}

		outStr += fmt.Sprintf("*%s*, %f", user.UserName, user.TotalPP)
	}

	return outStr
}

func slices(ctx context.Context, cmdChannel <-chan *quadlek.CommandMsg) {
	for {
		select {
		 case cmdMsg := <-cmdChannel:
			 cmdMsg.Command.Reply() <- &quadlek.CommandResp{
				 Text:      GetScores(),
				 InChannel: true,
			 }

			 case <-ctx.Done():
			 	log.Info("slices: stopping plugin")
		}
	}
}

func Register() quadlek.Plugin {
	return quadlek.MakePlugin(
		"Slices",
		[]quadlek.Command{
			quadlek.MakeCommand("slices", slices),
		},
		nil,
		nil,
		nil,
		nil,
	)
}