package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := os.Args[1]
	resourcesPath := "resources/"
	bannerFilename := "standard.txt"

	if len(os.Args) > 2 {
		requestedBanner := os.Args[2]
		availableBanners := map[string]bool{
			"standard":   true,
			"shadow":     true,
			"thinkertoy": true,
		}
		if availableBanners[requestedBanner] {
			bannerFilename = requestedBanner + ".txt"
		} else {
			fmt.Println("The Banner", requestedBanner, "is not available. Using standard.txt as a default ASCII Art.")
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
	asciiMap := make(map[rune][]string, 95)

	// Get total number of lines
	totalLines := len(bannerLines)

	// Calculate number of lines per symbol based on total lines
	// For thinkertoy without space line, number of lines would be 9
	// For standard.txt and shadow.txt that use space line, number of lines would be 10
	linesPerSymbol := totalLines / 95

	for char := 32; char <= 126; char++ {
		startLine := 1 + (int(char)-32)*linesPerSymbol
		endLine := startLine + 8
		fullArt := bannerLines[startLine:endLine]
		asciiMap[rune(char)] = fullArt
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
		for row := 0; row < 6; row++ {
			for _, ch := range word {
				result += asciiMap[ch][row]
			}
			result += "\n"
		}
	}
	return result
}
