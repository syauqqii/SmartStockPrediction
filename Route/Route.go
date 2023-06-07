package Route

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Middleware"
	"SmartStockPrediction/Controllers"
	"SmartStockPrediction/Controllers/Admin"
	"SmartStockPrediction/Controllers/Kasir"
)

func RunRoute() {
	r := mux.NewRouter()

	// ------------------------[  NO AUTH  ]----------------------------------
	r.HandleFunc("/login", Controllers.Login).Methods("POST")
	r.HandleFunc("/logout", Controllers.Logout).Methods("GET")

	// ------------------------[ AUTH ADMIN ]----------------------------------
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(Middleware.JWTAdminMiddleware)

	// >> CRUD USER ROUTE
	admin.HandleFunc("/users", Admin.ListUser).Methods("GET")
	admin.HandleFunc("/users", Admin.CreateUser).Methods("POST")
	admin.HandleFunc("/users/{id}", Admin.GetUserByID).Methods("GET")
	admin.HandleFunc("/users/{id}", Admin.UpdateUser).Methods("PUT")
	admin.HandleFunc("/users/{id}", Admin.DeleteUser).Methods("DELETE")

	// >> CRUD PELANGGAN ROUTE
	admin.HandleFunc("/pelanggan", Admin.ListPelanggan).Methods("GET")
	admin.HandleFunc("/pelanggan", Admin.CreatePelanggan).Methods("POST")
	admin.HandleFunc("/pelanggan/{id}", Admin.GetPelangganByID).Methods("GET")
	admin.HandleFunc("/pelanggan/{id}", Admin.UpdatePelanggan).Methods("PUT")
	admin.HandleFunc("/pelanggan/{id}", Admin.DeletePelanggan).Methods("DELETE")

	// >> CRUD KATEGORI PRODUK ROUTE
	admin.HandleFunc("/kategori-produk", Admin.ListKategoriProduk).Methods("GET")
	admin.HandleFunc("/kategori-produk", Admin.CreateKategoriProduk).Methods("POST")
	admin.HandleFunc("/kategori-produk/{id}", Admin.GetKategoriProdukByID).Methods("GET")
	admin.HandleFunc("/kategori-produk/{id}", Admin.UpdateKategoriProduk).Methods("PUT")
	admin.HandleFunc("/kategori-produk/{id}", Admin.DeleteKategoriProduk).Methods("DELETE")

	// ------------------------[ AUTH KASIR ]----------------------------------
	kasir := r.PathPrefix("/kasir").Subrouter()
	kasir.Use(Middleware.JWTKasirMiddleware)

	// >> CRUD PELANGGAN ROUTE
	kasir.HandleFunc("/pelanggan", Kasir.ListPelanggan).Methods("GET")
	kasir.HandleFunc("/pelanggan", Kasir.CreatePelanggan).Methods("POST")
	kasir.HandleFunc("/pelanggan/{id}", Kasir.GetPelangganByID).Methods("GET")
	kasir.HandleFunc("/pelanggan/{id}", Kasir.UpdatePelanggan).Methods("PUT")
	kasir.HandleFunc("/pelanggan/{id}", Kasir.DeletePelanggan).Methods("DELETE")

	// ------------------------[  START APP  ]----------------------------------
	fmt.Printf("\n > Server running on: %s", Utils.Serv.Sprintf("http://%s:%s\n", Utils.APP_HOST, Utils.APP_PORT))

	fmt.Printf("\n   > Endpoint (Admin) :")
	fmt.Printf("\n     - /users      -> GET        |  - /pelanggan      -> GET        |  - /kategori-produk      -> GET")
	fmt.Printf("\n     - /users      -> POST       |  - /pelanggan      -> POST       |  - /kategori-produk      -> POST")
	fmt.Printf("\n     - /users/{id} -> GET        |  - /pelanggan/{id} -> GET        |  - /kategori-produk/{id} -> GET")
	fmt.Printf("\n     - /users/{id} -> PUT        |  - /pelanggan/{id} -> PUT        |  - /kategori-produk/{id} -> PUT")
	fmt.Printf("\n     - /users/{id} -> DELETE     |  - /pelanggan/{id} -> DELETE     |  - /kategori-produk/{id} -> DELETE")

	fmt.Printf("\n\n   > Endpoint (Kasir) :")
	fmt.Printf("\n     - /pelanggan      -> GET    |")
	fmt.Printf("\n     - /pelanggan      -> POST   |")
	fmt.Printf("\n     - /pelanggan/{id} -> GET    |")
	fmt.Printf("\n     - /pelanggan/{id} -> PUT    |")
	fmt.Printf("\n     - /pelanggan/{id} -> DELETE |")

	fmt.Println("\n\n ------------------------------------------- [ LOG History ] -------------------------------------------\n")

	log.Fatal(http.ListenAndServe(Utils.APP_CONF, r))
}