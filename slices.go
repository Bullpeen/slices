package slices

import (
	"context"

	log "github.com/Sirupsen/logrus"
	"github.com/jirwin/quadlek/quadlek"
	"sort"
	"strings"
	"fmt"
)

var slicers []string


func GetScores() (string) {
	results := make(chan *User, len(slicers))

	for _, uid := range slicers {
		go func(uid string) {
			uu, err := GetUser(uid)

			results <- uu

			if err != nil {
				log.Info("Error fetching user: %s", err)
			}
		}(uid)
	}

	var output []*User

	for a := 0; a < len(slicers); a++ {
		if u := <-results; u != nil {
			output = append(output, u)
		}
	}

	close(results)
	sort.Sort(ByTotalPP(output))

	var outStr []string

	for _, user := range output {
		outStr = append(outStr, fmt.Sprintf("*%s*: %f", user.UserName, user.TotalPP))
	}

	return strings.Join(outStr, "\n")
}

func slices(ctx context.Context, cmdChannel <-chan *quadlek.CommandMsg) {
	for {
		select {
		case cmdMsg := <-cmdChannel:

			cmdMsg.Command.Reply() <- &quadlek.CommandResp{
				Text:         "Fetching Scores..",
				ResponseType: "ephemeral",
			}

			cmdMsg.Bot.RespondToSlashCommand(cmdMsg.Command.ResponseUrl, &quadlek.CommandResp{
				Text:      GetScores(),
				InChannel: true,
			})

		case <-ctx.Done():
			log.Info("slices: stopping plugin")
		}
	}
}

func Register(who []string) quadlek.Plugin {
	slicers = who
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
