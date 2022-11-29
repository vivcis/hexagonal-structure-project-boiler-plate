package repository

import (
	"errors"
)

var (
	errInvalidDatabaseInstance = errors.New("invalid db instance")
)

const (
	InstanceMysql int = iota
	InstanceMongo
)

// NewDatabaseFactory returns Db type based of the db instance provided
func NewDatabaseFactory(instance int) (*Db, error) {
	switch instance {
	case InstanceMongo:
		return ConnectMongoDB(newConfigMongoDB())
	default:
		return nil, errInvalidDatabaseInstance
	}
}
