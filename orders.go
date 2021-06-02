package woocommerce

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type GetOrdersParams struct {
	Context       string        `json:"context"` // view or edit
	Page          int           `json:"page"`
	PerPage       int           `json:"per_page"`
	Search        string        `json:"search"`
	After         string        `json:"after"` // Format ISO8601
	Before        string        `json:"before"`
	Exclude       []interface{} `json:"exclude"`
	Include       []interface{} `json:"include"`
	Offset        int           `json:"offset"`
	Order         string        `json:"order"`
	OrderBy       string        `json:"orderby"`
	Parent        []interface{} `json:"parent"`
	ParentExclude []interface{} `json:"parent_exclude"`
	Status        []interface{} `json:"status"`
	Customer      int           `json:"customer"`
	Product       int           `json:"product"`
	DP            int           `json:"dp"`
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
	url := c.option.buildURL(fmt.Sprintf("orders/%d", id))

	data := c.request.Clone().Data
	resp, body, errs := c.request.Clone().Put(url).Send(&data).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode >= 400 {
		var apiErr ApiError
		err := json.Unmarshal(body, &apiErr)
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

// CreateOrder create an order
func (c *Client) CreateOrder() (*Order, error) {
	url := c.option.buildURL("orders")

	// get order data from POST request's body
	var data Order
	err := mapstructure.Decode(c.request.Data, &data)
	if err != nil {
		return nil, err
	}

	resp, body, errs := c.request.Clone().Post(url).Send(&data).EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode >= 400 {
		var apiErr ApiError
		err := json.Unmarshal(body, &apiErr)
		if err != nil {
			return nil, err
		}
		return nil, &apiErr
	}

	var order Order
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil

}
