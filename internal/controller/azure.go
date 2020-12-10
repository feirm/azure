package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/feirm/azure/pkg/coins"
	"github.com/feirm/tatsuya/pkg/utils"
)

// Endpoint to return list of all coins
func ListAllCoins(w http.ResponseWriter, r *http.Request) {
	var coinsList []coins.CoinInformation

	// Iterate through the coins map and append to slice
	for _, v := range coins.Coins {
		coinsList = append(coinsList, v.CoinInformation)
	}

	// Return as JSON
	utils.RespondJSON(w, http.StatusOK, coinsList)
}

// Endpoint to return an individual coin
func FetchCoin(w http.ResponseWriter, r *http.Request) {
	var payload coins.RequestPayload

	// Decode the request body in the the payload model
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Check if coin ticker exists
	coin, ok := coins.Coins[strings.ToUpper(payload.Ticker)]
	if !ok {
		utils.RespondError(w, http.StatusBadRequest, "Coin not found!")
		return
	}

	utils.RespondJSON(w, http.StatusOK, coin)
}
