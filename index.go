package main

import (
	"context"
	"fmt"
	"go-parrot/effect"
	"go-parrot/entity"
	"go-parrot/game"
	"go-parrot/game/model"
	"strconv"
	"strings"
)

func Handler(ctx context.Context, question entity.Question) (*entity.ResponseEntity, error) {
	fmt.Println(ctx)

	var text = question.Request.OriginalUtterance
	var voice = text
	if strings.Contains(strings.ToLower(question.Request.OriginalUtterance), "олю") {
		text = "Привет, Оля"
		voice = effect.AddRandom(text)
	} else if strings.ToLower(question.Request.OriginalUtterance) == "1" {
		player1 := model.Human{
			Strength: 5,
			Agility:  10,
		}
		player2 := model.Human{
			Strength: 10,
			Agility:  5,
		}
		text = game.Fight(&player1, &player2)
		fmt.Printf("Player1 %v\n", strconv.Itoa(player1.Strength))
		fmt.Printf("Player2 %v\n", strconv.Itoa(player2.Strength))
	}
	return &entity.ResponseEntity{
			Version: "1.0",
			Session: question.Session.SessionID,
			Response: entity.Response{
				EndSession: false,
				Text:       text,
				Tts:        voice,
			},
		},
		nil
}
