package woocommerce

type Address struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Address1  string `json:"address_1"`
	Address2  string `json:"address_2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Postcode  string `json:"postcode"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Metadata struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Href struct {
	Href string `json:"href"`
}

type Links struct {
	Self       []*Href `json:"self"`
	Collection []*Href `json:"collection"`
	Customer   []*Href `json:"customer"`
}
type ShippingLine struct {
	ID          int         `json:"id"`
	MethodTitle string      `json:"method_title"`
	MethodID    string      `json:"method_id"`
	Total       string      `json:"total"`
	TotalTax    string      `json:"total_tax"`
	Taxes       []*Tax      `json:"taxes"`
	MetaData    []*Metadata `json:"meta_data"`
}

type Tax struct {
	ID       int    `json:"id"`
	Total    string `json:"total"`
	Subtotal string `json:"subtotal"`
}
type LineItem struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	ProductID   int         `json:"product_id"`
	VariationID int         `json:"variation_id"`
	Quantity    int         `json:"quantity"`
	TaxClass    string      `json:"tax_class"`
	Subtotal    string      `json:"subtotal"`
	SubtotalTax string      `json:"subtotal_tax"`
	Total       string      `json:"total"`
	TotalTax    string      `json:"total_tax"`
	Taxes       []*Tax      `json:"taxes"`
	MetaData    []*Metadata `json:"meta_data"`
	Sku         string      `json:"sku"`
	Price       int         `json:"price"`
}
type TaxLine struct {
	ID               int         `json:"id"`
	RateCode         string      `json:"rate_code"`
	RateID           int         `json:"rate_id"`
	Label            string      `json:"label"`
	Compound         bool        `json:"compound"`
	TaxTotal         string      `json:"tax_total"`
	ShippingTaxTotal string      `json:"shipping_tax_total"`
	MetaData         []*Metadata `json:"meta_data"`
}
type Order struct {
	ID                 int             `json:"id"`
	ParentID           int             `json:"parent_id"`
	Number             string          `json:"number"`
	OrderKey           string          `json:"order_key"`
	CreatedVia         string          `json:"created_via"`
	Version            string          `json:"version"`
	Status             string          `json:"status"`
	Currency           string          `json:"currency"`
	DateCreated        string          `json:"date_created"`
	DateCreatedGmt     string          `json:"date_created_gmt"`
	DateModified       string          `json:"date_modified"`
	DateModifiedGmt    string          `json:"date_modified_gmt"`
	DiscountTotal      string          `json:"discount_total"`
	DiscountTax        string          `json:"discount_tax"`
	ShippingTotal      string          `json:"shipping_total"`
	ShippingTax        string          `json:"shipping_tax"`
	CartTax            string          `json:"cart_tax"`
	Total              string          `json:"total"`
	TotalTax           string          `json:"total_tax"`
	PricesIncludeTax   bool            `json:"prices_include_tax"`
	CustomerID         int             `json:"customer_id"`
	CustomerIPAddress  string          `json:"customer_ip_address"`
	CustomerUserAgent  string          `json:"customer_user_agent"`
	CustomerNote       string          `json:"customer_note"`
	Billing            *Address        `json:"billing"`
	Shipping           *Address        `json:"shipping"`
	PaymentMethod      string          `json:"payment_method"`
	PaymentMethodTitle string          `json:"payment_method_title"`
	TransactionID      string          `json:"transaction_id"`
	DatePaid           string          `json:"date_paid"`
	DatePaidGmt        string          `json:"date_paid_gmt"`
	DateCompleted      interface{}     `json:"date_completed"`
	DateCompletedGmt   interface{}     `json:"date_completed_gmt"`
	CartHash           string          `json:"cart_hash"`
	MetaData           []*Metadata     `json:"meta_data"`
	LineItems          []*LineItem     `json:"line_items"`
	TaxLines           []*TaxLine      `json:"tax_lines"`
	ShippingLines      []*ShippingLine `json:"shipping_lines"`
	FeeLines           []interface{}   `json:"fee_lines"`
	CouponLines        []interface{}   `json:"coupon_lines"`
	Refunds            []interface{}   `json:"refunds"`
	Links              Links           `json:"_links,omitempty"`
}

type WebHook struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Status          string        `json:"status"`
	Topic           string        `json:"topic"`
	Resource        string        `json:"resource"`
	Event           string        `json:"event"`
	Hooks           []interface{} `json:"hooks"`
	DeliveryURL     string        `json:"delivery_url"`
	Secret          string        `json:"secret"`
	DateCreated     string        `json:"date_created"`
	DateCreatedGmt  string        `json:"date_created_gmt"`
	DateModified    string        `json:"date_modified"`
	DateModifiedGmt string        `json:"date_modified_gmt"`
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusOnHold     OrderStatus = "on-hold"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
	OrderStatusFailed     OrderStatus = "failed"
	OrderStatusTrash      OrderStatus = "trash"
)

type WebHookStatus string

const (
	WebHookStatusActive   WebHookStatus = "active"
	WebHookStatusPaused   WebHookStatus = "paused"
	WebHookStatusDisabled WebHookStatus = "disabled"
)
