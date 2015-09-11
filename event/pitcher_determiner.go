package event

import (
	"fmt"
	"github.com/sirsean/mlb_notifier/mlb"
	"strings"
)

type PitcherDeterminer struct {
	Perfect  int
	NoHitter int
	Shutout  int
	K        int
}

func (d PitcherDeterminer) Determine(from, to mlb.Game) []Event {
	events := make([]Event, 0)

	events = append(events, d.processPitchers(to, to.Boxscore.AwayPitchers)...)
	events = append(events, d.processPitchers(to, to.Boxscore.HomePitchers)...)

	return events
}

func (d PitcherDeterminer) processPitchers(game mlb.Game, pitchers []mlb.Pitcher) []Event {
	events := make([]Event, 0)

	if pitchers == nil || len(pitchers) == 0 {
		return events
	}

	// get the starter (which is the only pitcher we care about)
	pitcher := pitchers[0]

	feats := make([]string, 0)

	switch {
	case pitcher.Outs >= d.Perfect*3 && pitcher.Outs == pitcher.BF && pitcher.H == 0 && pitcher.BB == 0:
		// they have a perfect game
		// TODO: if someone reaches via error and is doubled off, we think it's a perfect game
		feats = append(feats, fmt.Sprintf("a perfect game through %v", pitcher.Innings()))
	case pitcher.Outs >= d.NoHitter*3 && pitcher.H == 0:
		// they have a no-hitter
		feats = append(feats, fmt.Sprintf("a no-hitter through %v", pitcher.Innings()))
	case pitcher.Outs >= d.Shutout*3 && pitcher.R == 0:
		// they have a shutout
		feats = append(feats, fmt.Sprintf("a shutout through %v", pitcher.Innings()))
	}

	// they have more than the minimum number of strikouts
	if d.K > 0 && pitcher.K >= d.K {
		feats = append(feats, fmt.Sprintf("%v K", pitcher.K))
	}

	if len(feats) > 0 {
		event := Event{
			GameId:   game.GameId,
			AwayTeam: game.AwayTeam,
			HomeTeam: game.HomeTeam,
			Text:     fmt.Sprintf("%v has %v", pitcher.Name, strings.Join(feats, ", ")),
		}
		events = append(events, event)
	}

	return events
}
