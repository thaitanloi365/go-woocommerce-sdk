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

/* func TestUpdateOrder(t *testing.T) {
	client := NewClient(&BasicOptions{
		URL:             os.Getenv("URL"),
		Key:             os.Getenv("KEY"),
		Secret:          os.Getenv("SECRET"),
		AdvancedOptions: AdvancedOptions{},
	})

	order, err := client.UpdateOrder(28)
	assert.NoError(t, err)
	fmt.Println(order)
}
*/

// URL=https://eziesell.com KEY=ck_66237fdf062b1f3bba71538c7ec7412fd87e74ab SECRET=cs_d0533a64b02a549cbddc2f634a84689c2c93a6c2 go test -timeout 2m -run ^TestGetOrder$ github.com/thaitanloi365/go-woocommerce-sdk -v
