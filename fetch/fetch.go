package fetch

import (
	"github.com/sirsean/mlb_notifier/build"
	"github.com/sirsean/mlb_notifier/mlb"
	"log"
	"net/http"
)

func FetchScoreboard(url string) ([]mlb.Game, error) {
	log.Printf("Fetching scoreboard from: %v\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("There was an error: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	return build.BuildGames(res.Body)
}

func FetchBoxscore(url string) (*mlb.Boxscore, error) {
	log.Printf("Fetching boxscore from: %v\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("There was an error: %v\n", err)
		return nil, err
	}
	defer res.Body.Close()
	return build.BuildBoxscore(res.Body)
}
