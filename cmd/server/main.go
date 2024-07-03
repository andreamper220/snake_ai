package main

import "github.com/andreamper220/snakeai.git/internal/application"

func main() {
	application.ParseFlags()
	if err := application.Run(); err != nil {
		panic(err)
	}
}
