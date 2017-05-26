package hello

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"fmt"
)

func GetAllCustomers(r *http.Request) ([]*Customer, error) {
	ctx := appengine.NewContext(r)
	customers := make([]*Customer, 0)
	query := datastore.NewQuery(CUSTOMER_DATASTORE_KIND)

	_, err := query.GetAll(ctx, &customers)
	if err != nil {
		return nil, fmt.Errorf("datastoredb: could not get customers: %v", err)
	}
	return customers, nil
}
