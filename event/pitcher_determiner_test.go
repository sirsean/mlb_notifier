package event

import (
    "testing"
    "github.com/sirsean/mlb_notifier/mlb"
)

func TestNoPitchers(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 0) {
        t.Errorf("Should be zero events")
    }
}

func TestEmptyPitchers(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: make([]mlb.Pitcher, 0),
            AwayPitchers: make([]mlb.Pitcher, 0),
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 0) {
        t.Errorf("Should be zero events")
    }
}

func TestPerfectThru3(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 10,
                    BF: 10,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 0) {
        t.Errorf("Should be zero events")
    }
}

func TestPerfectThru5(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 15,
                    BF: 15,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 1) {
        t.Errorf("Should be one event")
    }
    e := events[0]
    if e.GameId != "111" {
        t.Errorf("Wrong GameId")
    }
    if e.AwayTeam != g.AwayTeam {
        t.Errorf("Wrong AwayTeam")
    }
    if e.HomeTeam != g.HomeTeam {
        t.Errorf("Wrong HomeTeam")
    }
    if e.Text != "Some Guy is perfect through 5" {
        t.Errorf("Wrong Text")
    }
}

func Test1WalkOutsEqualBattersFacedThru8(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 24,
                    BF: 24,
                    BB: 1,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 1) {
        t.Errorf("Should be one event")
    }
    e := events[0]
    if e.GameId != "111" {
        t.Errorf("Wrong GameId")
    }
    if e.AwayTeam != g.AwayTeam {
        t.Errorf("Wrong AwayTeam")
    }
    if e.HomeTeam != g.HomeTeam {
        t.Errorf("Wrong HomeTeam")
    }
    if e.Text != "Some Guy has a no-hitter through 8" {
        t.Errorf("Wrong Text")
    }
}

func TestNoHitterThru3(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 9,
                    BF: 10,
                    BB: 1,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 0) {
        t.Errorf("Should be zero events")
    }
}

func TestNoHitterThru7(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 21,
                    BF: 22,
                    BB: 1,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 1) {
        t.Errorf("Should be one event")
    }
    e := events[0]
    if e.GameId != "111" {
        t.Errorf("Wrong GameId")
    }
    if e.AwayTeam != g.AwayTeam {
        t.Errorf("Wrong AwayTeam")
    }
    if e.HomeTeam != g.HomeTeam {
        t.Errorf("Wrong HomeTeam")
    }
    if e.Text != "Some Guy has a no-hitter through 7" {
        t.Errorf("Wrong Text")
    }
}

func TestShutoutThru5(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 15,
                    BF: 17,
                    BB: 1,
                    H: 2,
                    R: 0,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 0) {
        t.Errorf("Should be zero events")
    }
}

func TestShutoutThru8(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            AwayPitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Away Starter",
                    Outs: 24,
                    BF: 30,
                    H: 4,
                    BB: 2,
                    R: 0,
                },
            },
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 26,
                    BF: 30,
                    H: 1,
                    BB: 1,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 2) {
        t.Errorf("Should be two events")
    }
    e := events[0]
    if e.GameId != "111" {
        t.Errorf("Wrong GameId")
    }
    if e.AwayTeam != g.AwayTeam {
        t.Errorf("Wrong AwayTeam")
    }
    if e.HomeTeam != g.HomeTeam {
        t.Errorf("Wrong HomeTeam")
    }
    if e.Text != "Away Starter has a shutout through 8" {
        t.Errorf("Wrong away starter text")
    }
    if events[1].Text != "Some Guy has a shutout through 8" {
        t.Errorf("Wrong home starter Text")
    }
}

func TestNoHitterNotShutoutThru9(t *testing.T) {
    d := determiner()
    g := mlb.Game{
        GameId: "111",
        AwayTeam: mlb.Team{Name: "Team 1"},
        HomeTeam: mlb.Team{Name: "Team 2"},
        Boxscore: mlb.Boxscore{
            HomePitchers: []mlb.Pitcher{
                mlb.Pitcher{
                    Name: "Some Guy",
                    Outs: 27,
                    BF: 28,
                    H: 0,
                    BB: 1,
                    R: 1,
                    ER: 0,
                },
            },
        },
    }
    events := d.Determine(g, g)
    if (len(events) != 1) {
        t.Errorf("Should be one event")
    }
    e := events[0]
    if e.GameId != "111" {
        t.Errorf("Wrong GameId")
    }
    if e.AwayTeam != g.AwayTeam {
        t.Errorf("Wrong AwayTeam")
    }
    if e.HomeTeam != g.HomeTeam {
        t.Errorf("Wrong HomeTeam")
    }
    if e.Text != "Some Guy has a no-hitter through 9" {
        t.Errorf("Wrong Text")
    }
}

func determiner() PitcherDeterminer {
    return PitcherDeterminer{
        Perfect: 5,
        NoHitter: 5,
        Shutout: 8,
    }
}
