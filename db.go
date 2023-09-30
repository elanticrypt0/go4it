package go4it

import "gorm.io/gorm"

type DB struct {
	Primary   *gorm.DB
	Secondary *gorm.DB
	Actives   []DBActive
}

func (b *DB) SetPrimaryDB(index uint8) {
	b.Primary = b.Actives[index].Conn
}

func (b *DB) SetSecondaryDB(index uint8) {
	b.Secondary = b.Actives[index].Conn
}
