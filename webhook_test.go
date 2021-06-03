package woocommerce

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWebHook(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	requestData := make(map[string]interface{})
	requestData["name"] = "Order updated"
	requestData["topic"] = "order.updated"
	requestData["delivery_url"] = "http://requestb.in/1g0sxmo1"
	requestData["status"] = WebHookStatusPaused
	client.request.Data = requestData

	webHook, err := client.CreateWebHook()
	assert.NoError(t, err)
	assert.NotNil(t, webHook)
}

func TestDeleteWebHook(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	webHook, err := client.DeleteWebHook(1, "true")
	assert.NoError(t, err)
	assert.NotNil(t, webHook)
}
