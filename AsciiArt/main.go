package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	resourcesPath           string = "resources/"
	defaultBannerName              = "standard"
	argInputText                   = 1
	argBannerName                  = 2
	asciiSpace              rune   = 32
	asciiTilde              rune   = 126
	asciiSymbolsCount       int    = 95
	withoutSpacingSeparator        = 9
	withSpacingSeparator           = 10
	rowCountPerChar                = 6
	asciiArtStartOffset            = 1
)

func main() {
	text := os.Args[argInputText]
	bannerFilename := defaultBannerName + ".txt"

	if len(os.Args) > argBannerName {
		requestedBanner := os.Args[argBannerName]
		availableBanners := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}
		if availableBanners[requestedBanner] {
			bannerFilename = requestedBanner + ".txt"
		} else {
			fmt.Println("The Banner", requestedBanner, "is not available. Using", defaultBannerName+".txt", "as a default ASCII Art.")
		}
	}

	asciiArt := convertToAsciiArt(text, resourcesPath+bannerFilename)
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
	asciiMap := make(map[rune][]string, asciiSymbolsCount)

	totalLines := len(bannerLines)

	linesPerSymbol := withoutSpacingSeparator
	if (totalLines / asciiSymbolsCount) == withSpacingSeparator {
		linesPerSymbol = withSpacingSeparator
	}

	for char := asciiSpace; char <= asciiTilde; char++ {
		startLine := asciiArtStartOffset + (int(char)-int(asciiSpace))*linesPerSymbol
		endLine := startLine + rowCountPerChar
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
		for row := 0; row < rowCountPerChar; row++ {
			for _, ch := range word {
				result += asciiMap[ch][row]
			}
			result += "\n"
		}
	}
	return result
}
