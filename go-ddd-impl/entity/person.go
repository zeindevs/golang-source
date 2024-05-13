package entity

import "github.com/google/uuid"

// Person is an entity that represents a person in all Domain
type Person struct {
  // Id an the identifier of the entity
  ID uuid.UUID
  Name string
  Age int
}
