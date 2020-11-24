package database

import (
	"fmt"
	"time"

	"github.com/rlgino/hundred-days-of-code/patterns/singleton/model"
)

var (
	instance Connection
)

// Connection of database
type Connection interface {
	GetID() int64
	Select(ID int) model.DBTable
}

type connectionsInstance struct {
	HOST string
	PORT int
	DB   string
	ID   int64
}

// GetInstance get instance of connection
func GetInstance() Connection {
	if instance == nil {
		id := time.Now().Unix()
		instance = &connectionsInstance{
			HOST: "localhost",
			PORT: 3000,
			DB:   "pruebas",
			ID:   id,
		}
	}

	return instance
}

func (instance *connectionsInstance) Select(ID int) model.DBTable {
	fmt.Println("Finding ID: ", ID, "into", fmt.Sprintf("%s:%d/%s", instance.HOST, instance.PORT, instance.DB))
	return model.DBTable{}
}

func (instance *connectionsInstance) GetID() int64 {
	return instance.ID
}
