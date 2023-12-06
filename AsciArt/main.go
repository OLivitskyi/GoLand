package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter text to convert to ASCII Art:")
	inputText, _ := reader.ReadString('\n')
	inputText = strings.TrimSuffix(inputText, "\n") // Remove the newline character

	bannerFilename := "/Users/oleg_livitskyi/SandBox/Go/GoLand/AsciArt/resources/standard.txt"
	fmt.Println(ConvertToAsciiArt(inputText, bannerFilename))
}

func ConvertToAsciiArt(text, bannerFilename string) string {
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
	asciiMap := make(map[rune][]string)

	for char := 32; char <= 126; char++ {
		startLine := 1 + (int(char)-32)*9
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
		for row := 0; row < 8; row++ {
			for _, ch := range word {
				result += asciiMap[ch][row]
			}
			result += "\n"
		}
	}
	return result
}
