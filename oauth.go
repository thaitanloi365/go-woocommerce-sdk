package woocommerce

import (
	"fmt"
	"net/url"
)

func (c *Client) GenerateAPIKey(authParams *AuthParams) (string, error) {
	authenURL := c.option.URL + "/wc-auth/v1/authorize?"

	params := url.Values{}
	params.Add("app_name", authParams.AppName)
	params.Add("scope", authParams.Scope)
	params.Add("user_id", authParams.UserID)
	params.Add("return_url", authParams.ReturnURL)
	params.Add("callback_url", authParams.CallbackURL)

	encodedURL := fmt.Sprintf("%s%s", authenURL, params.Encode())

	_, body, errs := c.request.Clone().
		Get(encodedURL).
		End()

	if len(errs) > 0 {
		return "", errs[0]
	}

	return body, nil
}
