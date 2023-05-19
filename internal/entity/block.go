package entity

type Block struct {
	Number       uint64        `gorm:"primaryKey;unique"`
	Hash         string        `gorm:"size:128;uniqueIndex;not null"`
	Timestamp    uint64        `gorm:"not null"`
	ParentHash   string        `gorm:"size:128;not null"`
	Transactions []Transaction `gorm:"foreignKey:BlockHash;references:Hash"`
}
