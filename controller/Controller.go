package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dipankarupd/betting_raja/model"
)

type PredictResponse struct {
	ID     int    `json:"id"`
	Winner string `json:"winner"`
}

type GameRequest struct {
	ID int `json:"id"`
}

type GameResponse struct {
	ID    int    `json:"id"`
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	games := model.GetGames()

	gameResponses := make([]GameResponse, len(games))

	for i, game := range games {
		gameResponses[i] = GameResponse{
			ID:    game.Id,
			Team1: game.Team1,
			Team2: game.Team2,
		}
	}

	res, err := json.Marshal(gameResponses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func PredictWinner(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var gameRequest GameRequest

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(body, &gameRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wantedGame := model.GetGame(gameRequest.ID)

	predictResponse := PredictResponse{
		ID:     wantedGame.Id,
		Winner: wantedGame.Winner,
	}

	res, err := json.Marshal(predictResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
