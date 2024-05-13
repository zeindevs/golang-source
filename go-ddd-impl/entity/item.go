package entity

import "github.com/google/uuid"

// Item is an entity that represents a item in all Domain
type Item struct {
  // Id an the identifier of the entity
  ID uuid.UUID
  Name string
  Description string
}
