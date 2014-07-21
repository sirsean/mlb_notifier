package event

import (
    "fmt"
    "github.com/sirsean/mlb_notifier/mlb"
)

type PitcherDeterminer struct {
    Perfect int
    NoHitter int
    Shutout int
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

    switch {
        case pitcher.Outs >= d.Perfect * 3 && pitcher.Outs == pitcher.BF && pitcher.H == 0 && pitcher.BB == 0:
            // check if they have a perfect game
            // TODO: if someone reaches via error and is doubled off, we think it's a perfect game
            event := Event{
                GameId: game.GameId,
                AwayTeam: game.AwayTeam,
                HomeTeam: game.HomeTeam,
                Text: fmt.Sprintf("%s is perfect through %v", pitcher.Name, pitcher.Innings()),
            }
            events = append(events, event)
        case pitcher.Outs >= d.NoHitter * 3 && pitcher.H == 0:
            // they have a no-hitter
            event := Event{
                GameId: game.GameId,
                AwayTeam: game.AwayTeam,
                HomeTeam: game.HomeTeam,
                Text: fmt.Sprintf("%s has a no-hitter through %v", pitcher.Name, pitcher.Innings()),
            }
            events = append(events, event)
        case pitcher.Outs >= d.Shutout * 3 && pitcher.R == 0:
            // they have a shutout
            event := Event{
                GameId: game.GameId,
                AwayTeam: game.AwayTeam,
                HomeTeam: game.HomeTeam,
                Text: fmt.Sprintf("%s has a shutout through %v", pitcher.Name, pitcher.Innings()),
            }
            events = append(events, event)
    }

    return events
}
