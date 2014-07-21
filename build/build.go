package build

import (
	"encoding/json"
	"fmt"
	"github.com/sirsean/mlb_notifier/mlb"
	"io"
	"strconv"
)

func BuildGames(reader io.Reader) ([]mlb.Game, error) {
	dec := json.NewDecoder(reader)
	var v map[string]interface{}
	if err := dec.Decode(&v); err != nil {
		fmt.Printf("JSON decode error: %v\n", err)
		return nil, err
	}
	data := v["data"].(map[string]interface{})
	gamesObj := data["games"].(map[string]interface{})
	var gamesList []interface{}
	if _, ok := gamesObj["game"].([]interface{}); ok {
		gamesList = gamesObj["game"].([]interface{})
	} else if _, ok := gamesObj["game"].(map[string]interface{}); ok {
		gamesList = []interface{}{gamesObj["game"].(map[string]interface{})}
	}

	games := make([]mlb.Game, 0)
	for _, g := range gamesList {
		gameObj := g.(map[string]interface{})
		game := buildGame(gameObj)
		games = append(games, game)
	}

	return games, nil
}

func BuildBoxscore(reader io.Reader) (*mlb.Boxscore, error) {
	dec := json.NewDecoder(reader)
	var v map[string]interface{}
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	boxscore := new(mlb.Boxscore)
	data := v["data"].(map[string]interface{})
	boxscoreObj := data["boxscore"].(map[string]interface{})
	battingList := boxscoreObj["batting"].([]interface{})
	for _, o := range battingList {
		teamBattingObj := o.(map[string]interface{})
		teamFlag := teamBattingObj["team_flag"].(string)
		battersList := teamBattingObj["batter"].([]interface{})
		batters := make([]mlb.Batter, 0)
		for _, b := range battersList {
			batterObj := b.(map[string]interface{})
			batter := buildBatter(batterObj)
			batters = append(batters, batter)
		}
		switch teamFlag {
		case "home":
			boxscore.HomeBatters = batters
		case "away":
			boxscore.AwayBatters = batters
		}
	}
	pitchingList := boxscoreObj["pitching"].([]interface{})
	for _, o := range pitchingList {
		teamPitchingObj := o.(map[string]interface{})
		teamFlag := teamPitchingObj["team_flag"].(string)
		var pitchersList []interface{}
		if _, ok := teamPitchingObj["pitcher"].([]interface{}); ok {
			pitchersList = teamPitchingObj["pitcher"].([]interface{})
		} else if _, ok := teamPitchingObj["pitcher"].(map[string]interface{}); ok {
			pitchersList = []interface{}{teamPitchingObj["pitcher"]}
		}
		pitchers := make([]mlb.Pitcher, 0)
		for _, p := range pitchersList {
			pitcher := buildPitcher(p.(map[string]interface{}))
			pitchers = append(pitchers, pitcher)
		}
		switch {
		case teamFlag == "home":
			boxscore.HomePitchers = pitchers
		case teamFlag == "away":
			boxscore.AwayPitchers = pitchers
		}
	}

	return boxscore, nil
}

func buildGame(obj map[string]interface{}) mlb.Game {
	var status mlb.Status
	if _, ok := obj["status"]; ok {
		status = buildStatus(obj["status"].(map[string]interface{}))
	}
	var score mlb.LineScore
	if _, ok := obj["linescore"]; ok {
		score = buildLineScore(obj["linescore"].(map[string]interface{}))
	}

	return mlb.Game{
		GameId:   obj["id"].(string),
		Status:   status,
		AwayTeam: buildTeam("away", obj),
		HomeTeam: buildTeam("home", obj),
		Score:    score,
	}
}

func buildStatus(obj map[string]interface{}) mlb.Status {
	return mlb.Status{
		Status:      grabString(obj, "status"),
		Inning:      grabString(obj, "inning"),
		InningState: grabString(obj, "inning_state"),
	}
}

func buildTeam(homeAwayKey string, gameObj map[string]interface{}) mlb.Team {
	abbreviationKey := fmt.Sprintf("%v_name_abbrev", homeAwayKey)
	cityKey := fmt.Sprintf("%v_team_city", homeAwayKey)
	nameKey := fmt.Sprintf("%v_team_name", homeAwayKey)

	return mlb.Team{
		Abbreviation: grabString(gameObj, abbreviationKey),
		City:         grabString(gameObj, cityKey),
		Name:         grabString(gameObj, nameKey),
	}
}

func buildLineScore(obj map[string]interface{}) mlb.LineScore {
	var r, h, e mlb.HomeAway
	if _, ok := obj["r"]; ok {
		r = buildHomeAway(obj["r"].(map[string]interface{}))
	}
	if _, ok := obj["h"]; ok {
		h = buildHomeAway(obj["h"].(map[string]interface{}))
	}
	if _, ok := obj["e"]; ok {
		e = buildHomeAway(obj["e"].(map[string]interface{}))
	}

	return mlb.LineScore{
		Runs:   r,
		Hits:   h,
		Errors: e,
	}
}

func buildHomeAway(obj map[string]interface{}) mlb.HomeAway {
	return mlb.HomeAway{
		Home: grabInt(obj, "home"),
		Away: grabInt(obj, "away"),
	}
}

func buildBatter(obj map[string]interface{}) mlb.Batter {
	return mlb.Batter{
		Name: grabString(obj, "name_display_first_last"),
		AB:   grabInt(obj, "ab"),
		K:    grabInt(obj, "k"),
		BB:   grabInt(obj, "bb"),
		HBP:  grabInt(obj, "hbp"),
		H:    grabInt(obj, "h"),
		HR:   grabInt(obj, "hr"),
		R:    grabInt(obj, "r"),
		RBI:  grabInt(obj, "rbi"),
		SB:   grabInt(obj, "sb"),
		CS:   grabInt(obj, "cs"),
	}
}

func buildPitcher(obj map[string]interface{}) mlb.Pitcher {
	return mlb.Pitcher{
		Name:    grabString(obj, "name_display_first_last"),
		Outs:    grabInt(obj, "out"),
		BF:      grabInt(obj, "bf"),
		BB:      grabInt(obj, "bb"),
		K:       grabInt(obj, "so"),
		H:       grabInt(obj, "h"),
		ER:      grabInt(obj, "er"),
		R:       grabInt(obj, "r"),
		Pitches: grabInt(obj, "np"),
		Strikes: grabInt(obj, "s"),
	}
}

func grabString(obj map[string]interface{}, key string) string {
	if _, ok := obj[key]; ok {
		return obj[key].(string)
	}
	return ""
}

func grabInt(obj map[string]interface{}, key string) int {
	if _, ok := obj[key]; ok {
		i, _ := strconv.Atoi(obj[key].(string))
		return i
	}
	return 0
}
