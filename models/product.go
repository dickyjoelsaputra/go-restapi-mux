package models

type Product struct {
	Id    int64   `gorm:"primaryKey" json:"id"`
	Nama  string  `gorm:"varchar(300)" json:"nama"`
	Stok  int     `gorm:"int(5)" json:"stok"`
	Harga float64 `gorm:"decimal(14,2)" json:"harga"`
}