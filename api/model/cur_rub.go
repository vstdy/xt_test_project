package model

import (
	"encoding/json"
	"time"

	canonical "github.com/vstdy/xt_test_project/model"
)

type CurRubLatestResponse struct {
	Date time.Time `json:"date"`
	AUD  float32   `json:"AUD,omitempty"`
	AZN  float32   `json:"AZN,omitempty"`
	GBP  float32   `json:"GBP,omitempty"`
	AMD  float32   `json:"AMD,omitempty"`
	BYN  float32   `json:"BYN,omitempty"`
	BGN  float32   `json:"BGN,omitempty"`
	BRL  float32   `json:"BRL,omitempty"`
	HUF  float32   `json:"HUF,omitempty"`
	HKD  float32   `json:"HKD,omitempty"`
	DKK  float32   `json:"DKK,omitempty"`
	USD  float32   `json:"USD,omitempty"`
	EUR  float32   `json:"EUR,omitempty"`
	INR  float32   `json:"INR,omitempty"`
	KZT  float32   `json:"KZT,omitempty"`
	CAD  float32   `json:"CAD,omitempty"`
	KGS  float32   `json:"KGS,omitempty"`
	CNY  float32   `json:"CNY,omitempty"`
	MDL  float32   `json:"MDL,omitempty"`
	NOK  float32   `json:"NOK,omitempty"`
	PLN  float32   `json:"PLN,omitempty"`
	RON  float32   `json:"RON,omitempty"`
	XDR  float32   `json:"XDR,omitempty"`
	SGD  float32   `json:"SGD,omitempty"`
	TJS  float32   `json:"TJS,omitempty"`
	TRY  float32   `json:"TRY,omitempty"`
	TMT  float32   `json:"TMT,omitempty"`
	UZS  float32   `json:"UZS,omitempty"`
	UAH  float32   `json:"UAH,omitempty"`
	CZK  float32   `json:"CZK,omitempty"`
	SEK  float32   `json:"SEK,omitempty"`
	CHF  float32   `json:"CHF,omitempty"`
	ZAR  float32   `json:"ZAR,omitempty"`
	KRW  float32   `json:"KRW,omitempty"`
	JPY  float32   `json:"JPY,omitempty"`
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

type CurRubHistoryRequest struct {
	Pagination
	DateTimeFilter
	Currency string `json:"currency"`
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
