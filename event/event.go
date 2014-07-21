package event

import (
    "fmt"
    "github.com/sirsean/mlb_notifier/mlb"
)

type Determiner interface {
    Determine(from mlb.Game, to mlb.Game) []Event
}

type Event struct {
    GameId string
    AwayTeam mlb.Team
    HomeTeam mlb.Team
    Text string
}

func (e Event) Summary() string {
    return fmt.Sprintf("[%v @ %v] %v", e.AwayTeam.Abbreviation, e.HomeTeam.Abbreviation, e.Text)
}

func (e Event) DateKey() string {
    return DateKey(mlb.GameIdParts(e.GameId))
}

func DateKey(year, month, day int) string {
    return fmt.Sprintf("%v/%v/%v", year, month, day)
}
