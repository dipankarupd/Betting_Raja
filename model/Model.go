package model

type Game struct {
	Id     int    `json:id`
	Team1  string `json:team1`
	Team2  string `json:team2`
	Winner string `json:winner`
}

var games []Game

func GetGames() []Game {
	games = []Game{
		{Id: 1, Team1: "Team A", Team2: "Team B", Winner: "Team B"},
		{Id: 2, Team1: "Team C", Team2: "Team D", Winner: "Team C"},
		{Id: 3, Team1: "Team E", Team2: "Team F", Winner: "Team F"},
		{Id: 4, Team1: "Team G", Team2: "Team H", Winner: "Team G"},
		{Id: 5, Team1: "Team I", Team2: "Team J", Winner: "Team J"},
	}

	return games
}

func GetGame(id int) Game {

	for _, game := range games {

		if game.Id == id {
			return game
		}
	}
	return Game{}
}
