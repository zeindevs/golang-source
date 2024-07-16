package client

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type db interface {
	StoreName(name string) (string, error)
	GetName(uuid string) (string, error)
}

type logger interface {
	Debug(message string)
	Error(err error)
}

type client struct {
	logger      logger
	db          db
	numPlayers  int
	playersKeys []string
}

func NewClient(logger logger, db db) (*client, error) {
	if db == nil {
		return nil, errors.New("Tried to create new Client but db was nil")
	}
	return &client{
		logger:      logger,
		db:          db,
		playersKeys: []string{},
	}, nil
}

func (c *client) AddPlayers() error {
	if c.db == nil {
		return errors.New("Can't AddPlayers, db was nil")
	}
	key, _ := c.db.StoreName("Erik the Red")
	c.playersKeys = append(c.playersKeys, key)
	key, _ = c.db.StoreName("Attila the Hun")
	c.playersKeys = append(c.playersKeys, key)
	key, _ = c.db.StoreName("Vald the Impaler")
	c.playersKeys = append(c.playersKeys, key)
	key, _ = c.db.StoreName("Khal Drogho")
	c.playersKeys = append(c.playersKeys, key)
	key, _ = c.db.StoreName("Hoan of Arc")
	c.playersKeys = append(c.playersKeys, key)
	key, _ = c.db.StoreName("Caesar")
	c.playersKeys = append(c.playersKeys, key)
	key, _ = c.db.StoreName("Stephen Hawking")
	c.playersKeys = append(c.playersKeys, key)

	c.numPlayers = 7
	return nil
}

func (c *client) PlayRounds(numRounds int) error {
	if c.db == nil {
		return errors.New("Tried to play a round but db was nil")
	}
	i := 0
	for i < numRounds {
		c.logger.Debug(fmt.Sprintf("=== Playing Round %d ===", i+1))
		p1 := rand.Intn(c.numPlayers - 1)
		p2 := rand.Intn(c.numPlayers - 1)
		if p1 == p2 {
			c.logger.Debug("Round ended is a draw!")
			i++
			time.Sleep(time.Second * 2)
			continue
		}

		player1 := c.playersKeys[p1]
		p1Name, err := c.db.GetName(player1)
		if err != nil {
			c.logger.Error(err)
		}

		player2 := c.playersKeys[p2]
		p2Name, err := c.db.GetName(player2)
		if err != nil {
			c.logger.Error(err)
		}

		c.logger.Debug(fmt.Sprintf("Player %s attacked Player %s!", p1Name, p2Name))
		i++
		time.Sleep(time.Second * 2)
		winner := rand.Intn(1)
		if winner == 0 {
			c.logger.Debug(fmt.Sprintf("%s won!", p1Name))
		} else {
			c.logger.Debug(fmt.Sprintf("%s won!", p2Name))
		}
		c.logger.Debug("=== ROUND FINISHED ===\n")
	}
	return nil
}
