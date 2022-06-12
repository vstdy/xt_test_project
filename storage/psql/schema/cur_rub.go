package schema

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	canonical "github.com/vstdy/xt_test_project/model"
)

// CurRub keeps currencies to RUB rates data.
type (
	CurRub struct {
		bun.BaseModel `bun:"cur_rub,alias:cr"`
		ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()"`
		Date          time.Time `bun:"date,notnull"`
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

	CurRubs []CurRub
)

// NewCurRubFromCanonical creates a new CurRub DB object from canonical model.
func NewCurRubFromCanonical(obj canonical.CurRub) CurRub {
	return CurRub{
		ID:   obj.ID,
		Date: obj.Date,
		AUD:  int(obj.AUD * 10000),
		AZN:  int(obj.AZN * 10000),
		GBP:  int(obj.GBP * 10000),
		AMD:  int(obj.AMD * 10000),
		BYN:  int(obj.BYN * 10000),
		BGN:  int(obj.BGN * 10000),
		BRL:  int(obj.BRL * 10000),
		HUF:  int(obj.HUF * 10000),
		HKD:  int(obj.HKD * 10000),
		DKK:  int(obj.DKK * 10000),
		USD:  int(obj.USD * 10000),
		EUR:  int(obj.EUR * 10000),
		INR:  int(obj.INR * 10000),
		KZT:  int(obj.KZT * 10000),
		CAD:  int(obj.CAD * 10000),
		KGS:  int(obj.KGS * 10000),
		CNY:  int(obj.CNY * 10000),
		MDL:  int(obj.MDL * 10000),
		NOK:  int(obj.NOK * 10000),
		PLN:  int(obj.PLN * 10000),
		RON:  int(obj.RON * 10000),
		XDR:  int(obj.XDR * 10000),
		SGD:  int(obj.SGD * 10000),
		TJS:  int(obj.TJS * 10000),
		TRY:  int(obj.TRY * 10000),
		TMT:  int(obj.TMT * 10000),
		UZS:  int(obj.UZS * 10000),
		UAH:  int(obj.UAH * 10000),
		CZK:  int(obj.CZK * 10000),
		SEK:  int(obj.SEK * 10000),
		CHF:  int(obj.CHF * 10000),
		ZAR:  int(obj.ZAR * 10000),
		KRW:  int(obj.KRW * 10000),
		JPY:  int(obj.JPY * 10000),
	}
}

// ToCanonical converts a CurRub DB object to canonical model.
func (cr CurRub) ToCanonical() canonical.CurRub {
	return canonical.CurRub{
		ID:   cr.ID,
		Date: cr.Date,
		AUD:  float32(cr.AUD) / 10000,
		AZN:  float32(cr.AZN) / 10000,
		GBP:  float32(cr.GBP) / 10000,
		AMD:  float32(cr.AMD) / 10000,
		BYN:  float32(cr.BYN) / 10000,
		BGN:  float32(cr.BGN) / 10000,
		BRL:  float32(cr.BRL) / 10000,
		HUF:  float32(cr.HUF) / 10000,
		HKD:  float32(cr.HKD) / 10000,
		DKK:  float32(cr.DKK) / 10000,
		USD:  float32(cr.USD) / 10000,
		EUR:  float32(cr.EUR) / 10000,
		INR:  float32(cr.INR) / 10000,
		KZT:  float32(cr.KZT) / 10000,
		CAD:  float32(cr.CAD) / 10000,
		KGS:  float32(cr.KGS) / 10000,
		CNY:  float32(cr.CNY) / 10000,
		MDL:  float32(cr.MDL) / 10000,
		NOK:  float32(cr.NOK) / 10000,
		PLN:  float32(cr.PLN) / 10000,
		RON:  float32(cr.RON) / 10000,
		XDR:  float32(cr.XDR) / 10000,
		SGD:  float32(cr.SGD) / 10000,
		TJS:  float32(cr.TJS) / 10000,
		TRY:  float32(cr.TRY) / 10000,
		TMT:  float32(cr.TMT) / 10000,
		UZS:  float32(cr.UZS) / 10000,
		UAH:  float32(cr.UAH) / 10000,
		CZK:  float32(cr.CZK) / 10000,
		SEK:  float32(cr.SEK) / 10000,
		CHF:  float32(cr.CHF) / 10000,
		ZAR:  float32(cr.ZAR) / 10000,
		KRW:  float32(cr.KRW) / 10000,
		JPY:  float32(cr.JPY) / 10000,
	}
}

// ToCanonical converts list of CurRub DB objects to list of canonical models.
func (crs CurRubs) ToCanonical() []canonical.CurRub {
	objs := make([]canonical.CurRub, 0, len(crs))
	for _, dbObj := range crs {
		objs = append(objs, dbObj.ToCanonical())
	}

	return objs
}
