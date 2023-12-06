package main

import "testing"

func TestConvertToAsciiArt(t *testing.T) {
	got := ConvertToAsciiArt("A", "/Users/oleg_livitskyi/SandBox/Go/GoLand/AsciArt/resources/standard.txt")
	want := `_|_
/ \
X
\_/\
_|_
(/)` + "\n"
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
