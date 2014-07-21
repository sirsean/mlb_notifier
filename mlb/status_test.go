package mlb

import (
    "testing"
)

func TestInning1(t *testing.T) {
    status := Status{Inning: "1"}
    friendly := status.FriendlyInning()
    correct := "1st"
    if friendly != correct {
        t.Errorf("Wanted %v, got %v", correct, friendly)
    }
}

func TestInning4(t *testing.T) {
    status := Status{Inning: "4"}
    friendly := status.FriendlyInning()
    correct := "4th"
    if friendly != correct {
        t.Errorf("Wanted %v, got %v", correct, friendly)
    }
}
