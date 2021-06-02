package woocommerce

import (
	"fmt"
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
	fmt.Println(order)
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
