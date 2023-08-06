package game

import (
	"fmt"
	"go-parrot/game/cube"
	"go-parrot/game/model"
	"strconv"
)

func Fight(player1 *model.Human, player2 *model.Human) string {
	for player1.Strength > 0 && player2.Strength > 0 {
		power := cube.Dice()*2 + player1.Strength
		power2 := cube.Dice()*2 + player2.Strength
		if power > power2 {
			player1.Strength = player1.Strength - 2
		} else {
			player2.Strength = player2.Strength - 2
		}
		fmt.Printf("Player1 %v\n", strconv.Itoa(player1.Strength))
		fmt.Printf("Player2 %v\n", strconv.Itoa(player2.Strength))
	}
	if player1.Strength < 0 {
		return "Player1 is looser"
	} else {
		return "Player2 is looser"
	}

}
