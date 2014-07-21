package fetch

import (
    "fmt"
    "net/http"
    "github.com/sirsean/mlb_notifier/mlb"
    "github.com/sirsean/mlb_notifier/build"
)

func FetchScoreboard(url string) ([]mlb.Game, error) {
    fmt.Printf("Fetching scoreboard from: %v\n", url)
    res, err := http.Get(url)
    if err != nil {
        fmt.Printf("There was an error: %v\n", err)
        return nil, err
    }
    defer res.Body.Close()
    return build.BuildGames(res.Body)
}

func FetchBoxscore(url string) (*mlb.Boxscore, error) {
    fmt.Printf("Fetching boxscore from: %v\n", url)
    res, err := http.Get(url)
    if err != nil {
        fmt.Printf("There was an error: %v\n", err)
        return nil, err
    }
    defer res.Body.Close()
    return build.BuildBoxscore(res.Body)
}
