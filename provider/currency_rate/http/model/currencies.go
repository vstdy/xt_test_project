package model

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	canonical "github.com/vstdy/xt_test_project/model"
)

// BtcUsdtRate keeps BTC-USDT rate data.
type BtcUsdtRate struct {
	Code string `json:"code"`
	Data struct {
		Time             int64  `json:"time"`
		Symbol           string `json:"symbol"`
		Buy              string `json:"buy"`
		Sell             string `json:"sell"`
		ChangeRate       string `json:"changeRate"`
		ChangePrice      string `json:"changePrice"`
		High             string `json:"high"`
		Low              string `json:"low"`
		Vol              string `json:"vol"`
		VolValue         string `json:"volValue"`
		Last             string `json:"last"`
		AveragePrice     string `json:"averagePrice"`
		TakerFeeRate     string `json:"takerFeeRate"`
		MakerFeeRate     string `json:"makerFeeRate"`
		TakerCoefficient string `json:"takerCoefficient"`
		MakerCoefficient string `json:"makerCoefficient"`
	} `json:"data"`
}

// ToCanonical converts a currency_rate model to canonical model.
func (bur BtcUsdtRate) ToCanonical() (canonical.BtcUsdt, error) {
	buy, err := strconv.ParseFloat(bur.Data.Buy, 32)
	if err != nil {
		return canonical.BtcUsdt{}, err
	}
	sell, err := strconv.ParseFloat(bur.Data.Sell, 32)
	if err != nil {
		return canonical.BtcUsdt{}, err
	}

	obj := canonical.BtcUsdt{
		Timestamp: time.UnixMilli(bur.Data.Time),
		Buy:       float32(buy),
		Sell:      float32(sell),
	}

	return obj, nil
}

// CurRubRates keeps currencies to RUB rates.
type (
	CurRubRates struct {
		XMLName xml.Name     `xml:"ValCurs"`
		Date    string       `xml:"Date,attr"`
		Name    string       `xml:"name,attr"`
		Rates   []CurRubRate `xml:"Valute"`
	}

	CurRubRate struct {
		XMLName  xml.Name `xml:"Valute"`
		ID       string   `xml:"ID,attr"`
		NumCode  int      `xml:"NumCode"`
		CharCode string   `xml:"CharCode"`
		Nominal  int      `xml:"Nominal"`
		Name     string   `xml:"Name"`
		Value    string   `xml:"Value"`
	}
)

// ToCanonical converts a currency_rate model to canonical model.
func (crr CurRubRates) ToCanonical() (canonical.CurRub, error) {
	date, err := time.Parse("02.01.2006", crr.Date)
	if err != nil {
		return canonical.CurRub{}, err
	}

	m := make(map[string]float32)
	for _, cur := range crr.Rates {
		dotSeparatedValue := strings.ReplaceAll(cur.Value, ",", ".")
		value, err := strconv.ParseFloat(dotSeparatedValue, 32)
		if err != nil {
			return canonical.CurRub{}, err
		}
		m[cur.CharCode] = float32(value)
	}

	obj := canonical.CurRub{
		Date: date,
		AUD:  m["AUD"],
		AZN:  m["AZN"],
		GBP:  m["GBP"],
		AMD:  m["AMD"],
		BYN:  m["BYN"],
		BGN:  m["BGN"],
		BRL:  m["BRL"],
		HUF:  m["HUF"],
		HKD:  m["HKD"],
		DKK:  m["DKK"],
		USD:  m["USD"],
		EUR:  m["EUR"],
		INR:  m["INR"],
		KZT:  m["KZT"],
		CAD:  m["CAD"],
		KGS:  m["KGS"],
		CNY:  m["CNY"],
		MDL:  m["MDL"],
		NOK:  m["NOK"],
		PLN:  m["PLN"],
		RON:  m["RON"],
		XDR:  m["XDR"],
		SGD:  m["SGD"],
		TJS:  m["TJS"],
		TRY:  m["TRY"],
		TMT:  m["TMT"],
		UZS:  m["UZS"],
		UAH:  m["UAH"],
		CZK:  m["CZK"],
		SEK:  m["SEK"],
		CHF:  m["CHF"],
		ZAR:  m["ZAR"],
		KRW:  m["KRW"],
		JPY:  m["JPY"],
	}

	return obj, nil
}
