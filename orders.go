package woocommerce

import (
	"encoding/json"
	"errors"
	"fmt"
)

type GetOrdersParams struct {
	Context string `json:"context"` // view or edit
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Search  string `json:"search"`
	After   string `json:"after"` // Format ISO8601
	Before  string `json:"before"`
}

func (c *Client) GetOrders(params GetOrdersParams) ([]*Order, error) {
	var url = c.option.buildURL("orders")

	resp, body, errs := c.request.Clone().Query(params).Get(url).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode >= 400 {
		var apiErr ApiError
		var err = json.Unmarshal(body, &apiErr)
		if err != nil {
			return nil, err
		}

		return nil, &apiErr
	}

	var orders []*Order
	var err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil

}

func (c *Client) GetOrder(id int) (*Order, error) {
	var url = c.option.buildURL(fmt.Sprintf("orders/%d", id))
	resp, body, errs := c.request.Clone().Get(url).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode >= 400 {
		var apiErr ApiError
		var err = json.Unmarshal(body, &apiErr)
		if err != nil {
			return nil, err
		}

		return nil, &apiErr
	}

	var order Order
	var err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil

}

func (c *Client) UpdateOrder(id int) (*Order, error) {
	order, err := c.GetOrder(id)
	if order == nil {
		return nil, err
	}

	status, ok := c.request.Data["status"].(string)
	if !ok {
		return nil, errors.New("Order's status must be a string")
	}

	order.Status = status
	return order, nil
}

func (c *Client) CreateOrder() {

}
