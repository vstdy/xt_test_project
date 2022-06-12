package model

import (
	"encoding/json"
	"time"

	canonical "github.com/vstdy/xt_test_project/model"
)

type BtcUsdtLatestResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Sell      float32   `json:"value"`
	Buy       float32   `json:"sell"`
}

// MarshalJSON implements interface json.Marshaler.
func (bu BtcUsdtLatestResponse) MarshalJSON() ([]byte, error) {
	type BtcUsdtLatestResponseAlias BtcUsdtLatestResponse

	btcUsdtLatestResponse := struct {
		Timestamp string `json:"timestamp"`
		BtcUsdtLatestResponseAlias
	}{
		Timestamp:                  bu.Timestamp.Format("2006-01-02 15:04:05"),
		BtcUsdtLatestResponseAlias: BtcUsdtLatestResponseAlias(bu),
	}

	return json.Marshal(btcUsdtLatestResponse)
}

// NewBtcUsdtLatestResponseFromCanonical creates a new BtcUsdtLatestResponse object from canonical model.
func NewBtcUsdtLatestResponseFromCanonical(obj canonical.BtcUsdt) BtcUsdtLatestResponse {
	return BtcUsdtLatestResponse{
		Timestamp: obj.Timestamp,
		Buy:       obj.Buy,
		Sell:      obj.Sell,
	}
}

type BtcUsdtHistoryResponse struct {
	Total   int                     `json:"total"`
	History []BtcUsdtLatestResponse `json:"history"`
}

// NewBtcUsdtHistoryResponseFromCanonical creates a new BtcUsdtHistoryResponse object from canonical model.
func NewBtcUsdtHistoryResponseFromCanonical(total int, objs []canonical.BtcUsdt) BtcUsdtHistoryResponse {
	var history []BtcUsdtLatestResponse
	for _, obj := range objs {
		history = append(history, NewBtcUsdtLatestResponseFromCanonical(obj))
	}

	return BtcUsdtHistoryResponse{
		Total:   total,
		History: history,
	}
}
