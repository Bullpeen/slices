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

	var output []*User

	for _, uid := range slicers {
		uu, err := GetUser(uid)

		if err == nil {
			output = append(output, uu)
		} else {
			log.Info("Error fetching user: %s", err)
		}
	}

	sort.Sort(ByTotalPP(output))

	var outStr []string

	for _, user := range output {
		outStr = append(outStr, fmt.Sprintf("*%s*, %f", user.UserName, user.TotalPP))
	}

	return strings.Join(outStr, "\n")
}

func slices(ctx context.Context, cmdChannel <-chan *quadlek.CommandMsg) {
	for {
		select {
		 case cmdMsg := <-cmdChannel:
			 cmdMsg.Command.Reply() <- nil
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