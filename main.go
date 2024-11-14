package main

import (
	"fmt"
	"tournaments/server/models"
)

func main() {
	playerCount := 16
	players := make([]models.Player, playerCount)
	// matches := make([]models.Match, playerCount/2)

	for i := 0; i < playerCount; i++ {
		players[i] = models.NewPlayerData(i + 1)
	}

	// for i := 1; i < playerCount; i += 2 {
	// 	matches[i/2] = models.NewMatchData(players[i], players[i-1])
	// }

	tournament := models.NewTournamentData(players)

	winner := tournament.Winner()

	fmt.Printf("winner: %s\n", winner.Id())
}
