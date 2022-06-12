package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	canonical "github.com/vstdy/xt_test_project/model"
)

// BtcUsdt keeps BTC-USDT rate data.
type (
	BtcUsdt struct {
		bun.BaseModel `bun:"btc_usdt,alias:bu"`
		ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
		Buy           int       `bun:"buy,notnull"`
		Sell          int       `bun:"sell,notnull"`
		Timestamp     time.Time `bun:"timestamp,notnull"`
	}

	BtcUsdts []BtcUsdt
)

// NewBtcUsdtFromCanonical creates a new BtcUsdt DB object from canonical model.
func NewBtcUsdtFromCanonical(obj canonical.BtcUsdt) BtcUsdt {
	return BtcUsdt{
		ID:        obj.ID,
		Buy:       int(obj.Buy * 100),
		Sell:      int(obj.Sell * 100),
		Timestamp: obj.Timestamp,
	}
}

// ToCanonical converts a BtcUsdt DB object to canonical model.
func (bu BtcUsdt) ToCanonical() canonical.BtcUsdt {
	return canonical.BtcUsdt{
		ID:        bu.ID,
		Buy:       float32(bu.Buy) / 100,
		Sell:      float32(bu.Sell) / 100,
		Timestamp: bu.Timestamp,
	}
}

// ToCanonical converts list of BtcUsdt DB objects to list of canonical models.
func (bus BtcUsdts) ToCanonical() []canonical.BtcUsdt {
	objs := make([]canonical.BtcUsdt, 0, len(bus))
	for _, dbObj := range bus {
		objs = append(objs, dbObj.ToCanonical())
	}

	return objs
}
