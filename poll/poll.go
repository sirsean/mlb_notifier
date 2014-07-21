package poll

import (
	"fmt"
	"github.com/sirsean/mlb_notifier/comm"
	"github.com/sirsean/mlb_notifier/fetch"
	"github.com/sirsean/mlb_notifier/gamestore"
	"github.com/sirsean/mlb_notifier/mlb"
	"time"
)

func Start() {
	gameChannel := make(chan mlb.Game)
	pollTick := time.Tick(1 * time.Minute)
	cleanCommTick := time.Tick(10 * time.Minute)
	tick(gameChannel)
	for {
		select {
		case <-pollTick:
			tick(gameChannel)
		case <-cleanCommTick:
			comm.Clean()
		case game := <-gameChannel:
			fmt.Println("Received: ", game)
			gamestore.AddGame(game)
		}
	}
}

func tick(gameChannel chan mlb.Game) {
	getGamesFor(gameChannel, time.Now())
	getGamesFor(gameChannel, time.Now().Add(-24*time.Hour))
}

func getGamesFor(gameChannel chan mlb.Game, day time.Time) {
	start := time.Now()
	games, _ := fetch.FetchScoreboard(mlb.ScoreboardUrlFor(day))
	fmt.Println(len(games))
	fmt.Println(time.Since(start))
	for _, g := range games {
		go func(game mlb.Game) {
			if game.IsInProgress() {
				boxscore, err := fetch.FetchBoxscore(mlb.BoxscoreUrlFor(game))
				if err != nil {
					fmt.Println(err)
				}
				if boxscore != nil {
					game.Boxscore = *boxscore
				}
				gameChannel <- game
			}
		}(g)
	}
}
