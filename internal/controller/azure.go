package controller

import (
	"encoding/json"
	"net/http"

	"github.com/feirm/azure/pkg/coins"
)

// Endpoint to return list of all coins
func ListAllCoins(w http.ResponseWriter, r *http.Request) {
	var coinsList []coins.CoinInformation

	// Iterate through the coins map and append to slice
	for _, v := range coins.Coins {
		coinsList = append(coinsList, v.CoinInformation)
	}

	// Return as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coinsList)
}
