package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	canonical "github.com/vstdy/xt_test_project/model"
)

// CurBtc keeps currencies to BTC rates data.
type (
	CurBtc struct {
		bun.BaseModel `bun:"cur_btc,alias:cb"`
		ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
		Timestamp     time.Time `bun:"timestamp,notnull"`
		RUB           int       `bun:"rub,notnull"`
		AUD           int       `bun:"aud,notnull"`
		AZN           int       `bun:"azn,notnull"`
		GBP           int       `bun:"gbp,notnull"`
		AMD           int       `bun:"amd,notnull"`
		BYN           int       `bun:"byn,notnull"`
		BGN           int       `bun:"bgn,notnull"`
		BRL           int       `bun:"brl,notnull"`
		HUF           int       `bun:"huf,notnull"`
		HKD           int       `bun:"hkd,notnull"`
		DKK           int       `bun:"dkk,notnull"`
		USD           int       `bun:"usd,notnull"`
		EUR           int       `bun:"eur,notnull"`
		INR           int       `bun:"inr,notnull"`
		KZT           int       `bun:"kzt,notnull"`
		CAD           int       `bun:"cad,notnull"`
		KGS           int       `bun:"kgs,notnull"`
		CNY           int       `bun:"cny,notnull"`
		MDL           int       `bun:"mdl,notnull"`
		NOK           int       `bun:"nok,notnull"`
		PLN           int       `bun:"pln,notnull"`
		RON           int       `bun:"ron,notnull"`
		XDR           int       `bun:"xdr,notnull"`
		SGD           int       `bun:"sgd,notnull"`
		TJS           int       `bun:"tjs,notnull"`
		TRY           int       `bun:"try,notnull"`
		TMT           int       `bun:"tmt,notnull"`
		UZS           int       `bun:"uzs,notnull"`
		UAH           int       `bun:"uah,notnull"`
		CZK           int       `bun:"czk,notnull"`
		SEK           int       `bun:"sek,notnull"`
		CHF           int       `bun:"chf,notnull"`
		ZAR           int       `bun:"zar,notnull"`
		KRW           int       `bun:"krw,notnull"`
		JPY           int       `bun:"jpy,notnull"`
	}

	CurBtcs []CurBtc
)

// ToCanonical converts a CurBtc DB object to canonical model.
func (cb CurBtc) ToCanonical() canonical.CurBtc {
	return canonical.CurBtc{
		ID:        cb.ID,
		Timestamp: cb.Timestamp,
		AUD:       float32(cb.AUD) / 1000,
		AZN:       float32(cb.AZN) / 1000,
		GBP:       float32(cb.GBP) / 1000,
		AMD:       float32(cb.AMD) / 1000,
		BYN:       float32(cb.BYN) / 1000,
		BGN:       float32(cb.BGN) / 1000,
		BRL:       float32(cb.BRL) / 1000,
		HUF:       float32(cb.HUF) / 1000,
		HKD:       float32(cb.HKD) / 1000,
		DKK:       float32(cb.DKK) / 1000,
		USD:       float32(cb.USD) / 1000,
		EUR:       float32(cb.EUR) / 1000,
		INR:       float32(cb.INR) / 1000,
		KZT:       float32(cb.KZT) / 1000,
		CAD:       float32(cb.CAD) / 1000,
		KGS:       float32(cb.KGS) / 1000,
		CNY:       float32(cb.CNY) / 1000,
		MDL:       float32(cb.MDL) / 1000,
		NOK:       float32(cb.NOK) / 1000,
		PLN:       float32(cb.PLN) / 1000,
		RON:       float32(cb.RON) / 1000,
		XDR:       float32(cb.XDR) / 1000,
		SGD:       float32(cb.SGD) / 1000,
		TJS:       float32(cb.TJS) / 1000,
		TRY:       float32(cb.TRY) / 1000,
		TMT:       float32(cb.TMT) / 1000,
		UZS:       float32(cb.UZS) / 1000,
		UAH:       float32(cb.UAH) / 1000,
		CZK:       float32(cb.CZK) / 1000,
		SEK:       float32(cb.SEK) / 1000,
		CHF:       float32(cb.CHF) / 1000,
		ZAR:       float32(cb.ZAR) / 1000,
		KRW:       float32(cb.KRW) / 1000,
		JPY:       float32(cb.JPY) / 1000,
	}
}

// ToCanonical converts list of CurBtc DB objects to list of canonical models.
func (cbs CurBtcs) ToCanonical() []canonical.CurBtc {
	objs := make([]canonical.CurBtc, 0, len(cbs))
	for _, dbObj := range cbs {
		objs = append(objs, dbObj.ToCanonical())
	}

	return objs
}

// NewCurBtcFromCanonical creates a new CurBtc DB object from canonical model.
func NewCurBtcFromCanonical(obj canonical.CurBtc) CurBtc {
	return CurBtc{
		ID:        obj.ID,
		Timestamp: obj.Timestamp,
		RUB:       int(obj.RUB * 1000),
		AUD:       int(obj.AUD * 1000),
		AZN:       int(obj.AZN * 1000),
		GBP:       int(obj.GBP * 1000),
		AMD:       int(obj.AMD * 1000),
		BYN:       int(obj.BYN * 1000),
		BGN:       int(obj.BGN * 1000),
		BRL:       int(obj.BRL * 1000),
		HUF:       int(obj.HUF * 1000),
		HKD:       int(obj.HKD * 1000),
		DKK:       int(obj.DKK * 1000),
		USD:       int(obj.USD * 1000),
		EUR:       int(obj.EUR * 1000),
		INR:       int(obj.INR * 1000),
		KZT:       int(obj.KZT * 1000),
		CAD:       int(obj.CAD * 1000),
		KGS:       int(obj.KGS * 1000),
		CNY:       int(obj.CNY * 1000),
		MDL:       int(obj.MDL * 1000),
		NOK:       int(obj.NOK * 1000),
		PLN:       int(obj.PLN * 1000),
		RON:       int(obj.RON * 1000),
		XDR:       int(obj.XDR * 1000),
		SGD:       int(obj.SGD * 1000),
		TJS:       int(obj.TJS * 1000),
		TRY:       int(obj.TRY * 1000),
		TMT:       int(obj.TMT * 1000),
		UZS:       int(obj.UZS * 1000),
		UAH:       int(obj.UAH * 1000),
		CZK:       int(obj.CZK * 1000),
		SEK:       int(obj.SEK * 1000),
		CHF:       int(obj.CHF * 1000),
		ZAR:       int(obj.ZAR * 1000),
		KRW:       int(obj.KRW * 1000),
		JPY:       int(obj.JPY * 1000),
	}
}
