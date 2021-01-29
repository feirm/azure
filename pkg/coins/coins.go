package coins

import (
	"github.com/feirm/azure/pkg/models"
)

var Coins = map[string]*models.CoinData{
	"XFE":  &Feirm,
	"BTC":  &Bitcoin,
	"DOGE": &Doge,
}
