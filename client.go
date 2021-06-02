package woocommerce

import (
	"crypto/tls"

	"github.com/parnurzeal/gorequest"
)

type Client struct {
	option  *BasicOptions
	request *gorequest.SuperAgent
}

func NewClient(option *BasicOptions) *Client {
	return &Client{
		request: gorequest.New().Timeout(defaultTimeout).TLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		}).SetDebug(true),
		option: option,
	}

}
