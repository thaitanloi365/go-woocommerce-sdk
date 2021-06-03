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
	client.request.Data = requestData

	client.CreateWebHook()
	webHook, err := client.UpdateOrder(28)
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

	webHook, err := client.DeleteWebHook(1)
	assert.NoError(t, err)
	assert.NotNil(t, webHook)
}
