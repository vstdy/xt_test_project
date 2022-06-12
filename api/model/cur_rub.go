package model

import (
	"encoding/json"
	"time"

	canonical "github.com/vstdy/xt_test_project/model"
)

type CurRubLatestResponse struct {
	Date time.Time `json:"date"`
	AUD  float32   `json:"AUD"`
	AZN  float32   `json:"AZN"`
	GBP  float32   `json:"GBP"`
	AMD  float32   `json:"AMD"`
	BYN  float32   `json:"BYN"`
	BGN  float32   `json:"BGN"`
	BRL  float32   `json:"BRL"`
	HUF  float32   `json:"HUF"`
	HKD  float32   `json:"HKD"`
	DKK  float32   `json:"DKK"`
	USD  float32   `json:"USD"`
	EUR  float32   `json:"EUR"`
	INR  float32   `json:"INR"`
	KZT  float32   `json:"KZT"`
	CAD  float32   `json:"CAD"`
	KGS  float32   `json:"KGS"`
	CNY  float32   `json:"CNY"`
	MDL  float32   `json:"MDL"`
	NOK  float32   `json:"NOK"`
	PLN  float32   `json:"PLN"`
	RON  float32   `json:"RON"`
	XDR  float32   `json:"XDR"`
	SGD  float32   `json:"SGD"`
	TJS  float32   `json:"TJS"`
	TRY  float32   `json:"TRY"`
	TMT  float32   `json:"TMT"`
	UZS  float32   `json:"UZS"`
	UAH  float32   `json:"UAH"`
	CZK  float32   `json:"CZK"`
	SEK  float32   `json:"SEK"`
	CHF  float32   `json:"CHF"`
	ZAR  float32   `json:"ZAR"`
	KRW  float32   `json:"KRW"`
	JPY  float32   `json:"JPY"`
}

// MarshalJSON implements interface json.Marshaler.
func (cr CurRubLatestResponse) MarshalJSON() ([]byte, error) {
	type CurRubLatestResponseAlias CurRubLatestResponse

	btcUsdtLatestResponse := struct {
		Date string `json:"date"`
		CurRubLatestResponseAlias
	}{
		Date:                      cr.Date.Format("2006-01-02"),
		CurRubLatestResponseAlias: CurRubLatestResponseAlias(cr),
	}

	return json.Marshal(btcUsdtLatestResponse)
}

// NewCurRubLatestResponseFromCanonical creates a new CurRubLatestResponse object from canonical model.
func NewCurRubLatestResponseFromCanonical(obj canonical.CurRub) CurRubLatestResponse {
	return CurRubLatestResponse{
		Date: obj.Date,
		AUD:  obj.AUD,
		AZN:  obj.AZN,
		GBP:  obj.GBP,
		AMD:  obj.AMD,
		BYN:  obj.BYN,
		BGN:  obj.BGN,
		BRL:  obj.BRL,
		HUF:  obj.HUF,
		HKD:  obj.HKD,
		DKK:  obj.DKK,
		USD:  obj.USD,
		EUR:  obj.EUR,
		INR:  obj.INR,
		KZT:  obj.KZT,
		CAD:  obj.CAD,
		KGS:  obj.KGS,
		CNY:  obj.CNY,
		MDL:  obj.MDL,
		NOK:  obj.NOK,
		PLN:  obj.PLN,
		RON:  obj.RON,
		XDR:  obj.XDR,
		SGD:  obj.SGD,
		TJS:  obj.TJS,
		TRY:  obj.TRY,
		TMT:  obj.TMT,
		UZS:  obj.UZS,
		UAH:  obj.UAH,
		CZK:  obj.CZK,
		SEK:  obj.SEK,
		CHF:  obj.CHF,
		ZAR:  obj.ZAR,
		KRW:  obj.KRW,
		JPY:  obj.JPY,
	}
}

type CurRubHistoryResponse struct {
	Total   int                    `json:"total"`
	History []CurRubLatestResponse `json:"history"`
}

// NewCurRubHistoryResponseFromCanonical creates a new CurRubHistoryResponse object from canonical model.
func NewCurRubHistoryResponseFromCanonical(total int, objs []canonical.CurRub) CurRubHistoryResponse {
	var history []CurRubLatestResponse
	for _, obj := range objs {
		history = append(history, NewCurRubLatestResponseFromCanonical(obj))
	}

	return CurRubHistoryResponse{
		Total:   total,
		History: history,
	}
}
