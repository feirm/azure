package models

type CoinData struct {
	Coin Coin `json:"coin"`
}

// Representation of a full "coin"
type Coin struct {
	Name      string             `json:"name"`
	Ticker    string             `json:"ticker"`
	Icon      string             `json:"icon"`
	Blockbook string             `json:"blockbook"`
	TxVersion int                `json:"txVersion"`
	BIP44     int                `json:"bip44"`
	Networks  map[string]Network `json:"networks"`
}

type Network struct {
	MessagePrefix string           `json:"messagePrefix"`
	Bech32        string           `json:"bech32"`
	BIP32         BIP32Information `json:"bip32"`
	PubKeyHash    int              `json:"pubKeyHash"`
	ScriptHash    int              `json:"scriptHash"`
	WIF           int              `json:"wif"`
}

type BIP32Information struct {
	Public  int `json:"public"`
	Private int `json:"private"`
}
