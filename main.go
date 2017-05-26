package hello

import (
	"fmt"
	"net/http"
	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"strconv"
	"github.com/gorilla/mux"
)

const (
	GET_METHOD = "GET"
)

func init() {
	r := mux.NewRouter()
	r.Methods(GET_METHOD).Path("/customers").HandlerFunc(GetCustomers)
	r.HandleFunc("/", handler)
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		GetProduct(w)
	case http.MethodPost:
		AddProduct(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("HTTP %s not implemented!", r.Method)))
	}
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddCustomer(w, r)
	case http.MethodGet:
		GetCustomers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("HTTP %s not implemented!", r.Method)))
	}
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var customer Customer
	err := decoder.Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	datastoreKey, storeErr := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, CUSTOMER_DATASTORE_KIND, nil), &customer)
	if storeErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, marshalErr := json.Marshal(&customer)
	if marshalErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.FormatInt(datastoreKey.IntID(), 10)))
}

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := GetAllCustomers(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, marshalErr := json.Marshal(&customers)
	if marshalErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(b))
}

func GetProduct(w http.ResponseWriter) {
	p := Product{Id: "123", Description: "A product", Price: 3.14}
	productJson, err := json.Marshal(p)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, string(productJson))
	return
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	p := Product{Id: "123", Description: "A product", Price: 3.14}
	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "product", nil), &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, fmt.Sprintf("Added in datastore with key: %s", key))
}
