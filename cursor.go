package gdax

import (
	"context"
	"fmt"
)

type Cursor struct {
	HasMore bool

	ctx        context.Context
	client     *Client
	pagination *PaginationParams
	method     string
	params     interface{}
	url        string
	private    bool
}

func NewCursor(ctx context.Context, client *Client, private bool, method, url string, paginationParams *PaginationParams) *Cursor {
	return &Cursor{
		ctx:        ctx,
		client:     client,
		method:     method,
		url:        url,
		pagination: paginationParams,
		HasMore:    true,
		private:    private,
	}
}

func (c *Cursor) Page(i interface{}, direction string) error {
	url := c.url
	if c.pagination.Encode(direction) != "" {
		url = fmt.Sprintf("%s?%s", c.url, c.pagination.Encode(direction))
	}

	res, err := c.client.request(c.ctx, c.private, c.method, url, c.params, i)
	if err != nil {
		c.HasMore = false
		return err
	}

	c.pagination.Before = res.Header.Get("CB-BEFORE")
	c.pagination.After = res.Header.Get("CB-AFTER")

	if c.pagination.Done(direction) {
		c.HasMore = false
	}

	return nil
}

func (c *Cursor) NextPage(i interface{}) error {
	return c.Page(i, "next")
}

func (c *Cursor) PrevPage(i interface{}) error {
	return c.Page(i, "prev")
}
