package game_test

import (
	"github.com/google/uuid"
	"time"

	gamedata "snake_ai/internal/domain/game/data"
	matchdata "snake_ai/internal/domain/match/data"
)

const (
	gameWidth  = 5
	gameHeight = 5
	initX      = 3
	initY      = 3
	initXTo    = 1
	initYTo    = 0
)

func (s *GameTestSuite) TestSnakeMove() {
	pa := matchdata.NewParty()
	g := gamedata.NewGame(gameWidth, gameHeight, &pa)
	g.Food = &gamedata.Food{
		Position: gamedata.Point{
			X: 1,
			Y: 1,
		},
	}
	s.games.AddGame(g)

	tests := []struct {
		name     string
		commands []func(snake *gamedata.Snake)
		x        int
		y        int
	}{
		{
			name: "move",
			commands: []func(snake *gamedata.Snake){
				func(snake *gamedata.Snake) { snake.Move() },
			},
			x: 4,
			y: 3,
		},
		{
			name: "right",
			commands: []func(snake *gamedata.Snake){
				func(snake *gamedata.Snake) { snake.Right() },
				func(snake *gamedata.Snake) { snake.Move() },
			},
			x: 3,
			y: 4,
		},
		{
			name: "left",
			commands: []func(snake *gamedata.Snake){
				func(snake *gamedata.Snake) { snake.Left() },
				func(snake *gamedata.Snake) { snake.Move() },
			},
			x: 3,
			y: 2,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			sn := gamedata.NewSnake(initX, initY, initXTo, initYTo, tt.commands)
			userId := uuid.New()
			g.AddSnake(sn, userId)
			for range tt.commands {
				g.Update()
				time.Sleep(100 * time.Millisecond)
			}
			s.Assert().Equal(tt.x, sn.Body[0].X)
			s.Assert().Equal(tt.y, sn.Body[0].Y)
			g.RemoveSnake(userId)
		})
	}

	s.games.RemoveGame(g)
}

func (s *GameTestSuite) TestSnakeRotation() {
	pa := matchdata.NewParty()
	g := gamedata.NewGame(gameWidth, gameHeight, &pa)
	g.Food = &gamedata.Food{
		Position: gamedata.Point{
			X: 1,
			Y: 1,
		},
	}
	s.games.AddGame(g)

	tests := []struct {
		name       string
		command    func(snake *gamedata.Snake)
		directions []gamedata.Point
	}{
		{
			name:    "right",
			command: func(snake *gamedata.Snake) { snake.Right() },
			directions: []gamedata.Point{
				{0, 1},
				{-1, 0},
				{0, -1},
				{1, 0},
			},
		},
		{
			name:    "left",
			command: func(snake *gamedata.Snake) { snake.Left() },
			directions: []gamedata.Point{
				{0, -1},
				{-1, 0},
				{0, 1},
				{1, 0},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			sn := gamedata.NewSnake(initX, initY, initXTo, initYTo, []func(snake *gamedata.Snake){tt.command})
			userId := uuid.New()
			g.AddSnake(sn, userId)
			for _, direction := range tt.directions {
				g.Update()
				time.Sleep(100 * time.Millisecond)
				s.Assert().Equal(direction.X, sn.Direction.X)
				s.Assert().Equal(direction.Y, sn.Direction.Y)
			}
			g.RemoveSnake(userId)
		})
	}

	s.games.RemoveGame(g)
}

func (s *GameTestSuite) TestSnakeEdgeCollision() {
	pa := matchdata.NewParty()
	g := gamedata.NewGame(gameWidth, gameHeight, &pa)
	g.Food = &gamedata.Food{
		Position: gamedata.Point{
			X: 3,
			Y: 3,
		},
	}
	s.games.AddGame(g)

	tests := []struct {
		name  string
		initX int
		initY int
		xTo   int
		yTo   int
	}{
		{
			name:  "up",
			initX: 1,
			initY: 5,
			xTo:   0,
			yTo:   1,
		},
		{
			name:  "right",
			initX: 5,
			initY: 1,
			xTo:   1,
			yTo:   0,
		},
		{
			name:  "down",
			initX: 1,
			initY: 1,
			xTo:   0,
			yTo:   -1,
		},
		{
			name:  "left",
			initX: 1,
			initY: 1,
			xTo:   -1,
			yTo:   0,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			sn := gamedata.NewSnake(tt.initX, tt.initY, tt.xTo, tt.yTo, []func(snake *gamedata.Snake){
				func(snake *gamedata.Snake) { snake.Move() },
			})
			userId := uuid.New()
			g.AddSnake(sn, userId)
			time.Sleep(100 * time.Millisecond)
			g.Update()
			time.Sleep(100 * time.Millisecond)
			s.Assert().Equal(0, len(g.Snakes.Data))
		})
	}

	s.games.RemoveGame(g)
}

func (s *GameTestSuite) TestSnakeFoodEating() {
	user := s.AddNewUser()

	pa := matchdata.NewParty()
	g := gamedata.NewGame(gameWidth, gameHeight, &pa)
	g.Food = &gamedata.Food{
		Position: gamedata.Point{
			X: 4,
			Y: 3,
		},
	}
	s.games.AddGame(g)

	sn := gamedata.NewSnake(initX, initY, initXTo, initYTo, []func(snake *gamedata.Snake){
		func(snake *gamedata.Snake) { snake.Move() },
	})
	g.AddSnake(sn, user.Id)
	g.Update()
	time.Sleep(100 * time.Millisecond)
	g.Update()
	time.Sleep(100 * time.Millisecond)
	s.Assert().Equal(2, len(sn.Body))

	s.games.RemoveGame(g)
}
