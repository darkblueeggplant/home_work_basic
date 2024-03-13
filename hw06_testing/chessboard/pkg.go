package chessboard

import (
	"strings"
)

func PrintDesk(size int) string {
	black := "#"
	white := " "
	finalblack := ""
	finalwhite := ""
	var result strings.Builder

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			finalblack += black
			finalwhite += white
		} else {
			finalblack += white
			finalwhite += black
		}
	}

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			result.WriteString(finalblack)
			result.WriteString("\n")
		} else {
			result.WriteString(finalwhite)
			result.WriteString("\n")
		}
	}

	return result.String()
}
