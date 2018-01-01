package gdax

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"net/url"
	"strconv"
	"time"
	"github.com/google/uuid"
)

type Product struct {
	ID             string          `json:"id"`
	BaseCurrency   string          `json:"base_currency"`
	QuoteCurrency  string          `json:"quote_currency"`
	BaseMinSize    decimal.Decimal `json:"base_min_size,string"`
	BaseMaxSize    decimal.Decimal `json:"base_max_size,string"`
	QuoteIncrement decimal.Decimal `json:"quote_increment,string"`
}

type Ticker struct {
	TradeId int             `json:"trade_id,number"`
	Price   decimal.Decimal `json:"price,string"`
	Size    decimal.Decimal `json:"size,string"`
	Time    Time            `json:"time,string"`
	Bid     decimal.Decimal `json:"bid,string"`
	Ask     decimal.Decimal `json:"ask,string"`
	Volume  decimal.Decimal `json:"volume,string"`
}

type Trade struct {
	TradeId int             `json:"trade_id,number"`
	Price   decimal.Decimal `json:"price,string"`
	Size    decimal.Decimal `json:"size,string"`
	Time    Time            `json:"time,string"`
	Side    string          `json:"side"`
}

type HistoricRate struct {
	Time   time.Time
	Low    float64
	High   float64
	Open   float64
	Close  float64
	Volume float64
}

type Stats struct {
	Low          decimal.Decimal `json:"low,string"`
	High         decimal.Decimal `json:"high,string"`
	Open         decimal.Decimal `json:"open,string"`
	Volume       decimal.Decimal `json:"volume,string"`
	Last         decimal.Decimal `json:"last,string"`
	Volume_30Day decimal.Decimal `json:"volume_30day,string"`
}

type BookEntry struct {
	Price          decimal.Decimal
	Size           decimal.Decimal
	NumberOfOrders int
	OrderId        uuid.UUID
}

type Book struct {
	Sequence int         `json:"sequence"`
	Bids     []BookEntry `json:"bids"`
	Asks     []BookEntry `json:"asks"`
}

type ListTradesParams struct {
	Pagination PaginationParams
}

type GetHistoricRatesParams struct {
	Start       time.Time
	End         time.Time
	Granularity int
}

func (e *BookEntry) UnmarshalJSON(data []byte) error {
	var entry []interface{}

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	priceString, ok := entry[0].(string)
	if !ok {
		return errors.New("Expected string")
	}

	sizeString, ok := entry[1].(string)
	if !ok {
		return errors.New("Expected string")
	}

	price, err := decimal.NewFromString(priceString)
	if err != nil {
		return err
	}

	size, err := decimal.NewFromString(sizeString)
	if err != nil {
		return err
	}

	*e = BookEntry{
		Price: price,
		Size:  size,
	}

	var stringOrderId string
	numberOfOrdersFloat, ok := entry[2].(float64)
	if !ok {
		// Try to see if it's a string
		stringOrderId, ok = entry[2].(string)
		if !ok {
			return errors.New("Could not parse 3rd column, tried float and string")
		}
		e.OrderId, err = uuid.Parse(stringOrderId)
		if err != nil {
			return err
		}

	} else {
		e.NumberOfOrders = int(numberOfOrdersFloat)
	}

	return nil
}

func (e *HistoricRate) UnmarshalJSON(data []byte) error {
	var entry []interface{}

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	timeFloat, ok := entry[0].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	lowFloat, ok := entry[1].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	highFloat, ok := entry[2].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	openFloat, ok := entry[3].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	closeFloat, ok := entry[4].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	volumeFloat, ok := entry[5].(float64)
	if !ok {
		return errors.New("Expected float")
	}

	*e = HistoricRate{
		Time:   time.Unix(int64(timeFloat), 0),
		Low:    lowFloat,
		High:   highFloat,
		Open:   openFloat,
		Close:  closeFloat,
		Volume: volumeFloat,
	}

	return nil
}

func (c *Client) GetBook(ctx context.Context, product string, level int) (Book, error) {
	var book Book

	requestURL := fmt.Sprintf("/products/%s/book?level=%d", product, level)
	_, err := c.request(ctx, false, "GET", requestURL, nil, &book)
	return book, err
}

func (c *Client) GetTicker(ctx context.Context, product string) (Ticker, error) {
	var ticker Ticker

	requestURL := fmt.Sprintf("/products/%s/ticker", product)
	_, err := c.request(ctx, false, "GET", requestURL, nil, &ticker)
	return ticker, err
}

func (c *Client) ListTrades(ctx context.Context, product string, p ...ListTradesParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
	}

	return NewCursor(
		ctx,
		c,
		false,
		"GET",
		fmt.Sprintf("/products/%s/trades", product),
		&paginationParams,
	)
}

func (c *Client) GetProducts(ctx context.Context) ([]Product, error) {
	var products []Product

	requestURL := fmt.Sprintf("/products")
	_, err := c.request(ctx, false, "GET", requestURL, nil, &products)
	return products, err
}

func (c *Client) GetHistoricRates(ctx context.Context, product string, p ...GetHistoricRatesParams) ([]HistoricRate, error) {
	var historicRates []HistoricRate
	requestURL := fmt.Sprintf("/products/%s/candles", product)
	params := GetHistoricRatesParams{}
	if len(p) > 0 {
		params = p[0]
	}

	if !params.Start.IsZero() && !params.End.IsZero() && params.Granularity != 0 {
		values := url.Values{}
		layout := "2006-01-02T15:04:05Z"
		values.Add("start", params.Start.UTC().Format(layout))
		values.Add("end", params.End.UTC().Format(layout))
		values.Add("granularity", strconv.Itoa(params.Granularity))

		requestURL = fmt.Sprintf("%s?%s", requestURL, values.Encode())
	}

	_, err := c.request(ctx, false, "GET", requestURL, nil, &historicRates)
	return historicRates, err
}

func (c *Client) GetStats(ctx context.Context, product string) (Stats, error) {
	var stats Stats
	requestURL := fmt.Sprintf("/products/%s/stats", product)
	_, err := c.request(ctx, false, "GET", requestURL, nil, &stats)
	return stats, err
}
