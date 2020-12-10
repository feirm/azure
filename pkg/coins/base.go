package coins

import "encoding/json"

type Coin struct {
	CoinInformation CoinInformation `json:"coinInformation"`
}

// Representation of a full "coin"
type CoinInformation struct {
	Name      string                 `json:"name"`
	Ticker    string                 `json:"ticker"`
	Icon      string                 `json:"icon"`
	Blockbook string                 `json:"blockbook"`
	TxVersion int                    `json:"txVersion"`
	BIP44     int                    `json:"bip44"`
	Networks  map[string]CoinNetwork `json:"networks"`
}

type CoinNetwork struct {
	MessagePrefix string           `json:"messagePrefix"`
	Bech32        string           `json:"bech32"`
	BIP32         BIP32Information `json:"bip32"`
	PubKeyHash    []int            `json:"pubKeyHash"`
	ScriptHash    []int            `json:"scriptHash"`
	WIF           []int            `json:"wif"`
}

type BIP32Information struct {
	Private int `json:"private"`
	Public  int `json:"public"`
}

// Feirm services use BitcoinJS for client-side cryptographic operations,
// so all coins should be compatible with that in mind.

// Convert coin to BitcoinJS Compatible network
func ToBitcoinJS(coin Coin, network string) ([]byte, error) {
	// Coins can have multiple networks (mainnet, testnet, regest etc), so we want support for those too if needed.
	coinNetwork := coin.CoinInformation.Networks[network]

	// Marshal to JSON
	networkBytes, err := json.Marshal(coinNetwork)
	if err != nil {
		return nil, err
	}

	return networkBytes, nil
}
