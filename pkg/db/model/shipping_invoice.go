package model

type ShippingInvoice struct {
	BaseModel
	RouteID           string `json:"routeId"`
	Route             string `json:"-"`
	Side              string `json:"side"`
	Status            string `json:"status"`
	ExternalInvoiceID string `json:"externalInvoiceId"`
}

func (a ShippingInvoice) ResouceName() string {
	return "shippinginvoices"
}

func (a ShippingInvoice) TableName() string {
	return "shipping_invoices"
}
