package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	AUD string = "aud"
	AZN string = "azn"
	GBP string = "gbp"
	AMD string = "amd"
	BYN string = "byn"
	BGN string = "bgn"
	BRL string = "brl"
	HUF string = "huf"
	HKD string = "hkd"
	DKK string = "dkk"
	USD string = "usd"
	EUR string = "eur"
	INR string = "inr"
	KZT string = "kzt"
	CAD string = "cad"
	KGS string = "kgs"
	CNY string = "cny"
	MDL string = "mdl"
	NOK string = "nok"
	PLN string = "pln"
	RON string = "ron"
	XDR string = "xdr"
	SGD string = "sgd"
	TJS string = "tjs"
	TRY string = "try"
	TMT string = "tmt"
	UZS string = "uzs"
	UAH string = "uah"
	CZK string = "czk"
	SEK string = "sek"
	CHF string = "chf"
	ZAR string = "zar"
	KRW string = "krw"
	JPY string = "jpy"
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

// ValidateCurRub performs currencies validation.
func ValidateCurRub(cur string) error {
	switch strings.ToLower(cur) {
	case AUD, AZN, GBP, AMD, BYN, BGN, BRL, HUF, HKD, DKK, USD, EUR,
		INR, KZT, CAD, KGS, CNY, MDL, NOK, PLN, RON, XDR, SGD, TJS,
		TRY, TMT, UZS, UAH, CZK, SEK, CHF, ZAR, KRW, JPY:
		return nil
	default:
		return fmt.Errorf("unknown currency: %s", cur)
	}
}
