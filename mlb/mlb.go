package mlb

import(
    "fmt"
    "strconv"
    "strings"
    "time"
)

type Game struct {
    GameId string
    Status Status
    AwayTeam Team
    HomeTeam Team
    Score LineScore
    Boxscore Boxscore
}

func (g *Game) Leader() *Team {
    if g.Score.Runs.Home > g.Score.Runs.Away {
        return &g.HomeTeam
    } else if g.Score.Runs.Home < g.Score.Runs.Away {
        return &g.AwayTeam
    }
    return nil
}

func (g *Game) CurrentScore() (int, int) {
    leader := g.Leader()
    winningScore, losingScore := 0, 0
    if leader != nil && *leader == g.AwayTeam {
        winningScore = g.Score.Runs.Away
        losingScore = g.Score.Runs.Home
    } else {
        winningScore = g.Score.Runs.Home
        losingScore = g.Score.Runs.Away
    }
    return winningScore, losingScore
}

func (g Game) IsInProgress() bool {
    return g.Status.IsInProgress()
}

type Status struct {
    Status string
    Inning string
    InningState string
}

// A game is "in progress" as long as it's started and hasn't been postponed or finished.
// We'll use this to make sure we don't download boxscores for games that don't have any
// chance of having interesting/relevant boxscore data.
func (s Status) IsInProgress() bool {
    return s.Status == "In Progress" || s.Status == "Delayed" || s.Status == "Replay" || s.Status == "Game Over"
}

func (s *Status) InningInt() int {
    inning, _ := strconv.Atoi(s.Inning)
    return inning
}

func (s *Status) FriendlyInning() string {
    inning := s.InningInt()
    switch inning {
        case 1: return "1st"
        case 2: return "2nd"
        case 3: return "3rd"
        case 21: return "21st"
        case 22: return "22nd"
        case 23: return "23rd"
        case 31: return "31st"
        case 32: return "32nd"
        case 33: return "33rd"
    }
    return fmt.Sprintf("%dth", inning)
}

type Team struct {
    Abbreviation string
    City string
    Name string
}

func (t Team) SameAs(other Team) bool {
    return t.Abbreviation == other.Abbreviation
}

type LineScore struct {
    Runs HomeAway
    Hits HomeAway
    Errors HomeAway
}

type HomeAway struct {
    Away int
    Home int
}

type Batter struct {
    Name string
    AB int
    K int
    BB int
    HBP int
    H int
    HR int
    R int
    RBI int
    SB int
    CS int
}

type Pitcher struct {
    Name string
    Outs int
    BF int
    BB int
    K int
    H int
    ER int
    R int
    Pitches int
    Strikes int
}

func (p Pitcher) Innings() int {
    return p.Outs / 3
}

type Boxscore struct {
    HomeBatters []Batter
    AwayBatters []Batter
    HomePitchers []Pitcher
    AwayPitchers []Pitcher
}

func ScoreboardUrl() string {
    return ScoreboardUrlFor(time.Now())
}

func ScoreboardUrlFor(today time.Time) string {
    return fmt.Sprintf("http://gdx.mlb.com/components/game/mlb/year_%04d/month_%02d/day_%02d/master_scoreboard.json", today.Year(), today.Month(), today.Day())
}

func BoxscoreUrlFor(game Game) string {
    year, month, day := GameIdParts(game.GameId)
    gameId := strings.Replace(game.GameId, "/", "_", -1)
    gameId = strings.Replace(gameId, "-", "_", -1)
    return fmt.Sprintf("http://gdx.mlb.com/components/game/mlb/year_%04d/month_%02d/day_%02d/gid_%v/boxscore.json", year, month, day, gameId)
}

func GameIdParts(gameId string) (year, month, day int) {
    parts := strings.Split(gameId, "/")
    year, _ = strconv.Atoi(parts[0])
    month, _ = strconv.Atoi(parts[1])
    day, _ = strconv.Atoi(parts[2])
    return
}
