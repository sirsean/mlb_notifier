package event

import (
	"fmt"
	"github.com/sirsean/mlb_notifier/mlb"
	"strings"
)

type BatterDeterminer struct {
	H   int
	HR  int
	RBI int
	R   int
	SB  int
}

func (d BatterDeterminer) Determine(from mlb.Game, to mlb.Game) []Event {
	events := make([]Event, 0)

	events = append(events, d.processBatters(to, to.Boxscore.AwayBatters)...)
	events = append(events, d.processBatters(to, to.Boxscore.HomeBatters)...)

	return events
}

func (d BatterDeterminer) processBatters(game mlb.Game, batters []mlb.Batter) []Event {
	events := make([]Event, 0)
	for _, batter := range batters {
		feats := make([]string, 0)
		if batter.HR >= d.HR {
			feats = append(feats, fmt.Sprintf("%v HR", batter.HR))
		}
		if batter.H >= d.H {
			feats = append(feats, fmt.Sprintf("%v H", batter.H))
		}
		if batter.RBI >= d.RBI {
			feats = append(feats, fmt.Sprintf("%v RBI", batter.RBI))
		}
		if batter.R >= d.R {
			feats = append(feats, fmt.Sprintf("%v R", batter.R))
		}
		if batter.SB >= d.SB {
			feats = append(feats, fmt.Sprintf("%v SB", batter.SB))
		}
		if len(feats) > 0 {
			event := Event{
				GameId:   game.GameId,
				AwayTeam: game.AwayTeam,
				HomeTeam: game.HomeTeam,
				Text:     fmt.Sprintf("%v has %v", batter.Name, strings.Join(feats, ", ")),
			}
			events = append(events, event)
		}
	}
	return events
}
