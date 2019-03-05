package slices

import (
	"context"

	log "github.com/Sirupsen/logrus"
	"github.com/jirwin/quadlek/quadlek"
	"sort"
	"strings"
	"fmt"
	"sync"
)

var slicers []string

func GetScores() (string) {

	var output = make([]*User, len(slicers))
	var wg sync.WaitGroup

	wg.Add(len(slicers))

	for idx, uid := range slicers {
		go func(idx int, uid string) {
			defer wg.Done()
			uu, err := GetUser(uid)

			if err == nil {
				output[idx] = uu
			} else {
				log.Info("Error fetching user: %s", err)
			}
		}(idx, uid)
	}

	wg.Wait()

	sort.Sort(ByTotalPP(output))

	var outStr []string

	for _, user := range output {
		if user != nil {
			outStr = append(outStr, fmt.Sprintf("*%s*, %f", user.UserName, user.TotalPP))
		}
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
