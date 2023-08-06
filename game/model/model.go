package model

type Human struct {
	Strength int
	Agility  int
}

type GameStage struct {
	Text         string
	CurrentIndex int
	NextSteps    []int
	Enemy        Human
}

type GameState struct {
	Player Human
	Stage  GameStage
}
