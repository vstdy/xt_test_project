package model

import (
	"time"

	"github.com/google/uuid"
)

// CurRub keeps currencies to RUB rates data.
type CurRub struct {
	ID   uuid.UUID
	Date time.Time
	AUD  float32
	AZN  float32
	GBP  float32
	AMD  float32
	BYN  float32
	BGN  float32
	BRL  float32
	HUF  float32
	HKD  float32
	DKK  float32
	USD  float32
	EUR  float32
	INR  float32
	KZT  float32
	CAD  float32
	KGS  float32
	CNY  float32
	MDL  float32
	NOK  float32
	PLN  float32
	RON  float32
	XDR  float32
	SGD  float32
	TJS  float32
	TRY  float32
	TMT  float32
	UZS  float32
	UAH  float32
	CZK  float32
	SEK  float32
	CHF  float32
	ZAR  float32
	KRW  float32
	JPY  float32
}
