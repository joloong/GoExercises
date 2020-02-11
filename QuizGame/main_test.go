package main

import (
	"testing"
)

func TestParseLines(t *testing.T) {
	lines := [][]string{
		[]string{"5+5", "10"},
		[]string{"1+1", "2"},
	}

	got := parseLines(lines)
	problems := []problem{
		problem{
			q: "5+5",
			a: "10",
		},
		problem{
			q: "1+1",
			a: "2",
		},
	}

	for i, p := range problems {
		if p.q != got[i].q || p.a != got[i].a {
			t.Errorf("ParseLines is incorrect, got: %v, want: %v.\n", got[i], p)
			break
		}
	}
}
