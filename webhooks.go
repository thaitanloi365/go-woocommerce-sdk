package woocommerce

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func (c *Client) CreateWebHook() (*WebHook, error) {
	url := c.option.buildURL("webhooks")

	var data WebHook
	err := mapstructure.Decode(c.request.Data, &data)

	if err != nil {
		return nil, err
	}

	resp, body, errs := c.request.Clone().
		Post(url).
		Set("Content-Type", "application/json").
		SetBasicAuth(c.option.Key, c.option.Secret).
		Send(&data).
		EndBytes()
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

	var webHook WebHook
	err = json.Unmarshal(body, &webHook)
	if err != nil {
		return nil, err
	}

	return &webHook, nil
}

func (c *Client) DeleteWebHook(id int) (*WebHook, error) {
	url := c.option.buildURL(fmt.Sprintf("webhooks/%d", id))

	resp, body, errs := c.request.Clone().
		SetBasicAuth(c.option.Key, c.option.Secret).
		Delete(url).
		EndBytes()
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

	var webHook WebHook
	err := json.Unmarshal(body, &webHook)
	if err != nil {
		return nil, err
	}

	return &webHook, nil
}
