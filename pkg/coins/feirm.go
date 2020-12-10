package coins

var (
	Feirm = Coin{
		CoinInformation{
			Name:      "Feirm",
			Ticker:    "XFE",
			Icon:      "",
			Blockbook: "https://blockbook.feirm.com",
			TxVersion: 1,
			BIP44:     193,
			Networks: map[string]CoinNetwork{
				"p2pkh": {
					BIP32: BIP32Information{
						Private: 0x0488ade4,
						Public:  0x0488b21e,
					},
					PubKeyHash: []int{0x12},
					ScriptHash: []int{0x57},
					WIF:        []int{0x55},
				},
			},
		},
	}
)
