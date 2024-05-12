// Package tavern hold all the entitites that are shared across subdomains
package tavern

import "github.com/google/uuid"

// Person is an tavern that represents a perosn in all Domains
type Person struct {
	// Id an the identifier of the tavern
	ID   uuid.UUID
	Name string
	Age  int
}
