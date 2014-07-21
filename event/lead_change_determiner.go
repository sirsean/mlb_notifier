package event

import (
	"fmt"
	"github.com/sirsean/mlb_notifier/mlb"
)

type LeadChangeDeterminer struct {
	Inning int
}

func (determiner LeadChangeDeterminer) Determine(from mlb.Game, to mlb.Game) []Event {
	events := make([]Event, 0)
	if to.Status.InningInt() < determiner.Inning {
		return events
	}

	fromLeader := from.Leader()
	toLeader := to.Leader()

	switch {
	case fromLeader == nil && toLeader != nil:
		// it was tied, but now it isn't
		winningScore, losingScore := to.CurrentScore()
		event := Event{
			GameId:   from.GameId,
			AwayTeam: from.AwayTeam,
			HomeTeam: from.HomeTeam,
			Text:     fmt.Sprintf("%v broke the tie in the %v, %v-%v", toLeader.Abbreviation, to.Status.FriendlyInning(), winningScore, losingScore),
		}
		events = append(events, event)
	case fromLeader != nil && toLeader == nil:
		// it wasn't tied, but now it is
		winningScore, losingScore := to.CurrentScore()
		tyingTeam := to.HomeTeam
		if fromLeader.SameAs(to.HomeTeam) {
			tyingTeam = to.AwayTeam
		}
		event := Event{
			GameId:   from.GameId,
			AwayTeam: from.AwayTeam,
			HomeTeam: to.HomeTeam,
			Text:     fmt.Sprintf("%v tied it up in the %v, %v-%v", tyingTeam.Abbreviation, to.Status.FriendlyInning(), winningScore, losingScore),
		}
		events = append(events, event)
	case fromLeader != nil && !fromLeader.SameAs(*toLeader):
		// it wasn't tied, but the leader has changed
		winningScore, losingScore := to.CurrentScore()
		event := Event{
			GameId:   from.GameId,
			AwayTeam: from.AwayTeam,
			HomeTeam: from.HomeTeam,
			Text:     fmt.Sprintf("%v took the lead in the %v, %v-%v", toLeader.Abbreviation, to.Status.FriendlyInning(), winningScore, losingScore),
		}
		events = append(events, event)
	}
	return events
}
