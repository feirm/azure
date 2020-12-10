package coins

// Representation of a full "coin"
type Coin struct {
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
	BIP32         BIP32Information `json:"bip32"`
	PubKeyHash    []int            `json:"pubKeyHash"`
	ScriptHash    []int            `json:"scriptHash"`
	WIF           []int            `json:"wif"`
}

type BIP32Information struct {
	Private int `json:"private"`
	Public  int `json:"public"`
}
