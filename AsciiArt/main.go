package main

import (
	"fmt"
	"github.com/OLivitskyi/AsciiArt/const"
	"os"
	"strings"
)

func main() {
	text := os.Args[_const.ArgInputText]
	bannerFilename := _const.DefaultBannerName + ".txt"

	if len(os.Args) > _const.ArgBannerName {
		requestedBanner := os.Args[_const.ArgBannerName]
		availableBanners := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}
		if availableBanners[requestedBanner] {
			bannerFilename = requestedBanner + ".txt"
		} else {
			fmt.Println("The Banner", requestedBanner, "is not available. Using", _const.DefaultBannerName+".txt", "as a default ASCII Art.")
		}
	}

	asciiArt := convertToAsciiArt(text, _const.ResourcesPath+bannerFilename)
	fmt.Println(asciiArt)
}

func convertToAsciiArt(text, bannerFilename string) string {
	asciiMap := createMap(bannerFilename)
	lines := strings.Split(text, "\n")
	art := printArt(lines, asciiMap)
	return art
}

func createMap(filename string) map[rune][]string {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	bannerLines := strings.Split(string(fileContent), "\n")
	asciiMap := make(map[rune][]string, _const.AsciiSymbolsCount)

	totalLines := len(bannerLines)

	linesPerSymbol := _const.WithoutSpacingSeparator
	if (totalLines / _const.AsciiSymbolsCount) == _const.WithSpacingSeparator {
		linesPerSymbol = _const.WithSpacingSeparator
	}

	for char := _const.AsciiSpace; char <= _const.AsciiTilde; char++ {
		startLine := _const.AsciiArtStartOffset + (int(char)-int(_const.AsciiSpace))*linesPerSymbol
		endLine := startLine + _const.RowCountPerChar
		fullArt := bannerLines[startLine:endLine]
		asciiMap[char] = fullArt
	}

	return asciiMap
}

func printArt(lines []string, asciiMap map[rune][]string) string {
	var result string
	for _, word := range lines {
		if word == "" {
			result += "\n"
			continue
		}
		for row := 0; row < _const.RowCountPerChar; row++ {
			for _, ch := range word {
				result += asciiMap[ch][row]
			}
			result += "\n"
		}
	}
	return result
}
