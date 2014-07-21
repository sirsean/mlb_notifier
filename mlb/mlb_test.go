package mlb

import (
	"testing"
	"time"
)

func TestScoreboardUrlFor(t *testing.T) {
	date := time.Date(2013, time.September, 2, 12, 0, 0, 0, time.UTC)
	url := ScoreboardUrlFor(date)
	correct := "http://gdx.mlb.com/components/game/mlb/year_2013/month_09/day_02/master_scoreboard.json"
	if url != correct {
		t.Errorf("Wanted %v, got %v", correct, url)
	}
}

func TestScoreboardUrl(t *testing.T) {
	date := time.Now()
	correct := ScoreboardUrlFor(date)
	url := ScoreboardUrl()
	if url != correct {
		t.Errorf("Wanted %v, got %v", correct, url)
	}
}

func TestBoxscoreUrl(t *testing.T) {
	game := Game{GameId: "2013/09/06/milmlb-chnmlb-1"}
	correct := "http://gdx.mlb.com/components/game/mlb/year_2013/month_09/day_06/gid_2013_09_06_milmlb_chnmlb_1/boxscore.json"
	url := BoxscoreUrlFor(game)
	if url != correct {
		t.Errorf("Wanted %v, got %v", correct, url)
	}
}
