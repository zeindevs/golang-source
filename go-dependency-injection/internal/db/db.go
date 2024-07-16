package db

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type logger interface {
	Debug(message string)
	Error(err error)
}

type dbService struct {
	logger logger
	dbMap  map[string]string
}

func NewDbService(logger logger) (*dbService, error) {
	if logger == nil {
		return nil, errors.New("Invalid logger passed to DbService init")
	}
	return &dbService{
		logger: logger,
		dbMap:  map[string]string{},
	}, nil
}

func (d *dbService) StoreName(name string) (string, error) {
	uuid := uuid.NewString()
	d.logger.Debug(fmt.Sprintf("Created new UUID for insertion: %s", uuid))
	d.dbMap[uuid] = name
	return uuid, nil
}

func (d *dbService) GetName(uuid string) (string, error) {
	name, ok := d.dbMap[uuid]
	if !ok {
		return "", errors.New("UUID entry not found in DB")
	}
	return name, nil
}

func (d *dbService) GetAllKeys() ([]string, error) {
	if d.dbMap == nil {
		return nil, errors.New("Db not initialized")
	}
	keys := []string{}
	for k := range d.dbMap {
		keys = append(keys, k)
	}
	return keys, nil
}
