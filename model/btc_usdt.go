package model

import (
	"time"

	"github.com/google/uuid"
)

// BtcUsdt keeps BTC-USDT rate data.
type BtcUsdt struct {
	ID        uuid.UUID
	Timestamp time.Time
	Buy       float32
	Sell      float32
}
