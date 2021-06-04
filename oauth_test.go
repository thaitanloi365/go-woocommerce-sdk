package woocommerce

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAPIKey(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	appName := "WooCommerce API Client-Go"
	scope := "read_write"
	userID := "123"
	returnURL := "https://eziesell.wpcomstaging.com"
	callbackURL := "https://ensi7ppn341rg.x.pipedream.net"

	authParams := &AuthParams{
		AppName:     appName,
		Scope:       scope,
		UserID:      userID,
		ReturnURL:   returnURL,
		CallbackURL: callbackURL,
	}

	apiKey, err := client.GenerateAPIKey(authParams)
	assert.NoError(t, err)
	assert.NotZero(t, apiKey)
	fmt.Println(apiKey)
}
