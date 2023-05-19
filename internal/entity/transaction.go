package entity

type Transaction struct {
	Hash      string `json:"tx_hash" gorm:"primaryKey;unique"`
	From      string `json:"from"`
	To        string `json:"to"`
	Nonce     uint64 `json:"nonce"`
	Data      string `json:"data"`
	Value     int64  `json:"value"`
	BlockHash string `json:"_" gorm:"size:128;not null"`
	Logs      string `json:"logs"`
}

type Log struct {
	Index uint   `json:"index"`
	Data  string `json:"data"`
}
