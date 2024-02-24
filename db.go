package go4it

import "gorm.io/gorm"

type DB struct {
	Primary   *gorm.DB
	Secondary *gorm.DB
	Security  *gorm.DB
	Auth      *gorm.DB
	Actives   []DBActive
}

func (b *DB) SetPrimaryDB(index uint8) {
	b.Primary = b.Actives[index].Conn
}

func (b *DB) SetSecondaryDB(index uint8) {
	b.Secondary = b.Actives[index].Conn
}

// Set the db to interact with security
func (b *DB) SetSecurityDB(index uint8) {
	b.Security = b.Actives[index].Conn
}

// Set the db to interact with authentication
func (b *DB) SetAuthDB(index uint8) {
	b.Auth = b.Actives[index].Conn
}
