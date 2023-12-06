package main

import "testing"

func TestConvertToAsciiArt(t *testing.T) {
	got := convertToAsciiArt("A", "resources/standard.txt")
	want :=
		"           \n" +
			"    /\\     \n" +
			"   /  \\    \n" +
			"  / /\\ \\   \n" +
			" / ____ \\  \n" +
			"/_/    \\_\\ \n"
	if got != want {
		t.Errorf("\ngot  %q\n, want %q", got, want)
	}
}

func TestConvertToAsciiArtWithThinkertoy(t *testing.T) {
	got := convertToAsciiArt("A", "resources/thinkertoy.txt")
	want :=
		"      \n" +
			"  O   \n" +
			" / \\  \n" +
			"o---o \n" +
			"|   | \n" +
			"o   o \n"
	if got != want {
		t.Errorf("\ngot  %q\n, want %q", got, want)
	}
}

func TestConvertToAsciiArtWithShadow(t *testing.T) {
	got := convertToAsciiArt("A", "resources/shadow.txt")
	want := "         \n" +
		"  _|_|   \n" +
		"_|    _| \n" +
		"_|_|_|_| \n" +
		"_|    _| \n" +
		"_|    _| \n"
	if got != want {
		t.Errorf("\ngot  %q\n, want %q", got, want)
	}
}
