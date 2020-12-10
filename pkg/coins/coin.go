package coins

import (
	"errors"
	"strings"
)

var Coins = map[string]*Coin{
	"XFE": &Feirm,
}

// Check if a coin exists and return its data
func FetchCoin(ticker string) (*Coin, error) {
	coin, ok := Coins[strings.ToUpper(ticker)]
	if !ok {
		return nil, errors.New("Coin does not exist.")
	}

	return coin, nil
}
