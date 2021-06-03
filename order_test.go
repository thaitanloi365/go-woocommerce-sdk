package woocommerce

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrder(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	order, err := client.GetOrder(28)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestGetOrders(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	orders, err := client.GetOrders(nil)
	assert.NoError(t, err)
	assert.NotNil(t, orders)
}

func TestUpdateOrder(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	// fields to be updated
	requestData := make(map[string]interface{})
	requestData["status"] = "on-hold"
	client.request.Data = requestData

	order, err := client.UpdateOrder(28)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestCreateOrder(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	// data of the order that will be created
	reuestData := make(map[string]interface{})
	reuestData["payment_method"] = "bacs"
	reuestData["payment_method_title"] = "Direct Bank Transfer"
	reuestData["set_paid"] = true
	reuestData["currency"] = "USD"
	reuestData["status"] = "pending"
	reuestData["billing"] = map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"address_1":  "969 Market",
		"address_2":  "",
		"city":       "San Francisco",
		"state":      "CA",
		"postcode":   "94103",
		"country":    "US",
		"email":      "john.doe@example.com",
		"phone":      "(555) 555-5555",
	}
	reuestData["shipping"] = map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"address_1":  "969 Market",
		"address_2":  "",
		"city":       "San Francisco",
		"state":      "CA",
		"postcode":   "94103",
		"country":    "US",
	}
	reuestData["line_items"] = []interface{}{
		map[string]interface{}{
			"product_id": 93,
			"quantity":   2,
		},
		map[string]interface{}{
			"product_id":   22,
			"variation_id": 23,
			"quantity":     1,
		},
	}
	reuestData["shipping_lines"] = []interface{}{
		map[string]interface{}{
			"method_id":    "flat_rate",
			"method_title": "Flat Rate",
			"total":        "10.00",
		},
	}
	client.request.Data = reuestData

	order, err := client.CreateOrder()
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

// URL=https://eziesell.com KEY=ck_66237fdf062b1f3bba71538c7ec7412fd87e74ab SECRET=cs_d0533a64b02a549cbddc2f634a84689c2c93a6c2 go test -timeout 2m -run ^TestGetOrder$ github.com/thaitanloi365/go-woocommerce-sdk -v
