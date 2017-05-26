package hello

const CUSTOMER_DATASTORE_KIND = "customer"

type Product struct {
	Id          string
	Description string
	Price       float32
}

type Order struct {
	Id string
	OrderNumber string
	Description string
}

type Customer struct {
	CustomerNumber string
	Name string
	Orders []Order `json:",omitempty"`
}
