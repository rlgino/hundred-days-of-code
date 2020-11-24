package database

import "fmt"

var (
	instance Connection
)

// Connection of database
type Connection interface {
	GetConnection()
}

type connectionsInstance struct {
	HOST string
	PORT int
	DB   string
}

// GetInstance get instance of connection
func GetInstance() Connection {
	if instance == nil {
		instance = &connectionsInstance{
			HOST: "localhost",
			PORT: 3000,
			DB:   "pruebas",
		}
	}

	return instance
}

func (instance *connectionsInstance) GetConnection() {
	fmt.Println(instance.HOST)
}
