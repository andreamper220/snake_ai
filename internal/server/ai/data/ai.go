package data

import (
	"snake_ai/internal/logger"
	"strings"
)

func GenerateAiFunctions(ai string) []func(snake *Snake) {
	aiStrings := strings.Split(ai, `,`)
	aiFunctions := make([]func(snake *Snake), len(aiStrings)-1)
	for i, aiString := range aiStrings {
		if aiString != "" {
			switch aiString {
			case "right":
				aiFunctions[i] = func(snake *Snake) { snake.Right() }
			case "left":
				aiFunctions[i] = func(snake *Snake) { snake.Left() }
			case "move":
				aiFunctions[i] = func(snake *Snake) { snake.Move() }
			}
		}
	}
	logger.Log.Infof("AI functions: %v", aiFunctions)

	return aiFunctions
}
