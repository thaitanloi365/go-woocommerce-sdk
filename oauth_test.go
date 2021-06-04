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

	storeURL := "https://eziesell.wpcomstaging.com"
	appName := "WooCommerce API Client-Go"
	scope := "read_write"
	userID := "123"
	returnURL := "https://en1m4618qmvrd.x.pipedream.net"
	callbackURL := "https://en1m4618qmvrd.x.pipedream.net"

	authParams := &AuthParams{
		StoreURL:    storeURL,
		AppName:     appName,
		Scope:       scope,
		UserID:      userID,
		ReturnURL:   returnURL,
		CallbackURL: callbackURL,
	}

	authURL, err := client.GenerateAPIKey(authParams)
	assert.NoError(t, err)
	assert.NotZero(t, authURL)
	fmt.Println(authURL)
}
