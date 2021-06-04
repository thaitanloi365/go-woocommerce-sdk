package woocommerce

import (
	"fmt"
	"net/url"
)

func (c *Client) GenerateAPIKey(authParams *AuthParams) (string, error) {
	authURL, err := url.Parse(fmt.Sprintf("%s/wc-auth/v1/authorize?", authParams.StoreURL))
	if err != nil {
		return "", err
	}

	params := url.Values{}
	params.Add("app_name", authParams.AppName)
	params.Add("scope", authParams.Scope)
	params.Add("user_id", authParams.UserID)
	params.Add("return_url", authParams.ReturnURL)
	params.Add("callback_url", authParams.CallbackURL)

	authURL.RawQuery = params.Encode()

	return authURL.String(), nil
}
