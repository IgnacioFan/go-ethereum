package entity

type Block struct {
	Number     uint64 `json:"block_num"`
	Hash       string `json:"block_hash"`
	Timestamp  uint64 `json:"block_time"`
	ParentHash string `json:"parent_hash"`
}
