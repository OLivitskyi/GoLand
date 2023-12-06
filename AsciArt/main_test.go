package main

import "testing"

func TestConvertToAsciiArt(t *testing.T) {
	got := ConvertToAsciiArt("A", "/Users/oleg_livitskyi/SandBox/Go/GoLand/AsciArt/resources/standard.txt")
	want :=
		"           \n" +
			"    /\\     \n" +
			"   /  \\    \n" +
			"  / /\\ \\   \n" +
			" / ____ \\  \n" +
			"/_/    \\_\\ \n" +
			"           \n" +
			"           \n"
	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
