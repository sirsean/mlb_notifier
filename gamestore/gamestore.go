package gamestore

import (
    "fmt"
    "sync"
    "github.com/sirsean/mlb_notifier/mlb"
    "github.com/sirsean/mlb_notifier/event"
    "github.com/sirsean/mlb_notifier/comm"
)

var store = struct {
    m map[string]mlb.Game
    sync.Mutex
}{m: make(map[string]mlb.Game)}

var determiners = []event.Determiner{
    event.LeadChangeDeterminer{Inning: 7},
    event.BatterDeterminer{H: 4, HR: 2, RBI: 4, R: 4, SB: 3},
    event.PitcherDeterminer{Perfect: 5, NoHitter: 5, Shutout: 8},
}

func AddGame(game mlb.Game) {
    store.Lock()
    existing, ok := store.m[game.GameId]
    store.Unlock()

    if ok {
        events := determine(existing, game)
        fmt.Printf("EVENTS for %v: ", game.GameId)
        fmt.Println(events)
        comm.Send(events)
    }

    store.Lock()
    store.m[game.GameId] = game
    store.Unlock()
}

func determine(from mlb.Game, to mlb.Game) []event.Event {
    events := make([]event.Event, 0)
    for _, determiner := range determiners {
        events = append(events, determiner.Determine(from, to)...)
    }
    return events
}
