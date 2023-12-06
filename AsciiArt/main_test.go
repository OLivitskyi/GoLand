package main

import "testing"

func TestConvertToAsciiArtStandart(t *testing.T) {
	got := convertToAsciiArt("A", "resources/standard.txt")
	want :=
		"           \n" +
			"    /\\     \n" +
			"   /  \\    \n" +
			"  / /\\ \\   \n" +
			" / ____ \\  \n" +
			"/_/    \\_\\ \n"

	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
func TestConvertToAsciiArtWithThinkertoy(t *testing.T) {
	got := convertToAsciiArt("A", "resources/thinkertoy.txt")
	want := "      \n" +
		"  O   \n" +
		" / \\  \n" +
		"o---o \n" +
		"|   | \n" +
		"o   o \n"

	if got != want {
		t.Errorf("\ngot  %q\nwant %q", got, want)
	}
}
