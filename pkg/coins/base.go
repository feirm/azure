package coins

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
	PubKeyHash    int              `json:"pubKeyHash"`
	ScriptHash    int              `json:"scriptHash"`
	WIF           int              `json:"wif"`
}

type BIP32Information struct {
	Private int `json:"private"`
	Public  int `json:"public"`
}

// Coin request payload
type RequestPayload struct {
	Ticker string `json:"ticker"`
}
