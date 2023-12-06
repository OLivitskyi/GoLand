package main

import "testing"

func TestConvertToAsciiArt(t *testing.T) {
	got := convertToAsciiArt("A", "resources/standard.txt")
	want := "           \n" +
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
