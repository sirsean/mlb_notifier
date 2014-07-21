package mlb

import (
    "testing"
)

func TestLeaderEmpty(t *testing.T) {
    game := Game{}
    leader := game.Leader()
    if leader != nil {
        t.Errorf("Leader should be nil")
    }
}

func TestLeaderNoScore(t *testing.T) {
    game := Game{
        AwayTeam: Team{Name: "Team 1"},
        HomeTeam: Team{Name: "Team 2"},
    }
    leader := game.Leader()
    if leader != nil {
        t.Errorf("Leader should be nil")
    }
}

func TestLeaderTied(t *testing.T) {
    game := Game{
        AwayTeam: Team{Name: "Team 1"},
        HomeTeam: Team{Name: "Team 2"},
        Score: LineScore{Runs:HomeAway{Away:0, Home:0}},
    }
    leader := game.Leader()
    if leader != nil {
        t.Errorf("Leader should be nil")
    }
}

func TestLeaderAway(t *testing.T) {
    game := Game{
        AwayTeam: Team{Name: "Team 1"},
        HomeTeam: Team{Name: "Team 2"},
        Score: LineScore{Runs:HomeAway{Away:1, Home:0}},
    }
    leader := game.Leader()
    if *leader != game.AwayTeam {
        t.Errorf("Leader should be AwayTeam")
    }
}

func TestLeaderHome(t *testing.T) {
    game := Game{
        AwayTeam: Team{Name: "Team 1"},
        HomeTeam: Team{Name: "Team 2"},
        Score: LineScore{Runs:HomeAway{Away:1, Home:2}},
    }
    leader := game.Leader()
    if *leader != game.HomeTeam {
        t.Errorf("Leader should be HomeTeam")
    }
}
