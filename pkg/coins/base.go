package coins

import (
	"github.com/btcsuite/btcd/chaincfg"
)

// Representation of a full "coin"
type Coin struct {
	Information CoinInformation `json:"coinInformation"`
	NetParams   *chaincfg.Params
}

type CoinInformation struct {
	Name      string                 `json:"name"`
	Ticker    string                 `json:"ticker"`
	Icon      string                 `json:"icon"` // Base-64 encoded PNG
	Blockbook string                 `json:"blockbook"`
	TxVersion int                    `json:"txVersion"`
	TxBuilder string                 `json:"txBuilder"`
	BIP44     int                    `json:"bip44"`
	Networks  map[string]CoinNetwork `json:"networks"`
}

type CoinNetwork struct {
	MessagePrefix string           `json:"messagePrefix"`
	BIP32         BIP32Information `json:"bip32"`
	PubKeyHash    []int            `json:"pubKeyHash"`
	ScriptHash    []int            `json:"scriptHash"`
	WIF           []int            `json:"wif"`
}

type BIP32Information struct {
	Private string `json:"private"`
	Public  string `json:"public"`
}
