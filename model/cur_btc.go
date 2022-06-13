package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const RUB string = "rub"

// CurBtc keeps currencies to BTC rates data.
type CurBtc struct {
	ID        uuid.UUID
	Timestamp time.Time
	RUB       float32
	AUD       float32
	AZN       float32
	GBP       float32
	AMD       float32
	BYN       float32
	BGN       float32
	BRL       float32
	HUF       float32
	HKD       float32
	DKK       float32
	USD       float32
	EUR       float32
	INR       float32
	KZT       float32
	CAD       float32
	KGS       float32
	CNY       float32
	MDL       float32
	NOK       float32
	PLN       float32
	RON       float32
	XDR       float32
	SGD       float32
	TJS       float32
	TRY       float32
	TMT       float32
	UZS       float32
	UAH       float32
	CZK       float32
	SEK       float32
	CHF       float32
	ZAR       float32
	KRW       float32
	JPY       float32
}

//NewCurBtc creates a new CurBtc model from CurRub and BtcUsdt models.
func NewCurBtc(curRub CurRub, btcUsdt BtcUsdt) CurBtc {
	btcRub := btcUsdt.Sell * curRub.USD

	return CurBtc{
		Timestamp: btcUsdt.Timestamp,
		RUB:       btcRub,
		AUD:       btcRub / curRub.AUD,
		AZN:       btcRub / curRub.AZN,
		GBP:       btcRub / curRub.GBP,
		AMD:       btcRub / curRub.AMD,
		BYN:       btcRub / curRub.BYN,
		BGN:       btcRub / curRub.BGN,
		BRL:       btcRub / curRub.BRL,
		HUF:       btcRub / curRub.HUF,
		HKD:       btcRub / curRub.HKD,
		DKK:       btcRub / curRub.DKK,
		USD:       btcUsdt.Sell,
		EUR:       btcRub / curRub.EUR,
		INR:       btcRub / curRub.INR,
		KZT:       btcRub / curRub.KZT,
		CAD:       btcRub / curRub.CAD,
		KGS:       btcRub / curRub.KGS,
		CNY:       btcRub / curRub.CNY,
		MDL:       btcRub / curRub.MDL,
		NOK:       btcRub / curRub.NOK,
		PLN:       btcRub / curRub.PLN,
		RON:       btcRub / curRub.RON,
		XDR:       btcRub / curRub.XDR,
		SGD:       btcRub / curRub.SGD,
		TJS:       btcRub / curRub.TJS,
		TRY:       btcRub / curRub.TRY,
		TMT:       btcRub / curRub.TMT,
		UZS:       btcRub / curRub.UZS,
		UAH:       btcRub / curRub.UAH,
		CZK:       btcRub / curRub.CZK,
		SEK:       btcRub / curRub.SEK,
		CHF:       btcRub / curRub.CHF,
		ZAR:       btcRub / curRub.ZAR,
		KRW:       btcRub / curRub.KRW,
		JPY:       btcRub / curRub.JPY,
	}
}

// ValidateCurBtc performs currencies validation.
func ValidateCurBtc(cur string) error {
	switch strings.ToLower(cur) {
	case RUB, AUD, AZN, GBP, AMD, BYN, BGN, BRL, HUF, HKD, DKK, USD, EUR,
		INR, KZT, CAD, KGS, CNY, MDL, NOK, PLN, RON, XDR, SGD, TJS,
		TRY, TMT, UZS, UAH, CZK, SEK, CHF, ZAR, KRW, JPY:
		return nil
	default:
		return fmt.Errorf("unknown currency: %s", cur)
	}
}
