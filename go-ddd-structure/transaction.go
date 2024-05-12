package tavern

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a valueobject because has no identifier and is unmutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
