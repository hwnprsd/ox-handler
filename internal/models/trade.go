package models

type TradeV1 struct {
	Id          string `json:"id"`
	UserAddress string `json:"user_address"`
	AssetA      string `json:"asset_a"`
	AssetB      string `json:"asset_b"`
	AmountA     string `json:"amount_a"`
	AmountB     string `json:"amoubt_b"`
	TxHash      string `json:"tx_hash"`
	Fee         string `json:"fee"`
}
