package database

import "fmt"

type Database struct {
	StorageU2SU map[string]string
	StorageSU2U map[string]string
	Connected   bool
}

func NewDB() *Database {
	db := new(Database)
	db.StorageU2SU = make(map[string]string)
	db.StorageSU2U = make(map[string]string)
	db.Connected = false
	return db
}

var DB *Database = NewDB()

func (db *Database) OpenConnection() {
	if !db.Connected {
		db.Connected = true
	} else {
		fmt.Println("DB is CONNECTED!")
	}
}

func (db *Database) CloseConnection() {
	if db.Connected {
		db.Connected = false
	} else {
		fmt.Println("DB is DISCONNECTED!")
	}
}
