package controller

import (
	"net/http"

	"github.com/feirm/azure/pkg/coins"
	"github.com/feirm/azure/pkg/models"
	"github.com/feirm/tatsuya/pkg/utils"
)

// Endpoint to return list of all coins
func ListAllCoins(w http.ResponseWriter, r *http.Request) {
	var coinsList []models.Coin

	// Iterate through the coins map and append to slice
	for _, v := range coins.Coins {
		coinsList = append(coinsList, v.Coin)
	}

	// Return as JSON
	utils.RespondJSON(w, http.StatusOK, coinsList)
}
