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

	// >> CRUD PRODUK ROUTE
	admin.HandleFunc("/produk", Admin.GetAllProduk).Methods("GET")
	admin.HandleFunc("/produk", Admin.CreateProduk).Methods("POST")
	admin.HandleFunc("/produk/{id}", Admin.GetProdukByID).Methods("GET")
	admin.HandleFunc("/produk/{id}", Admin.UpdateProduk).Methods("PUT")
	admin.HandleFunc("/produk/{id}", Admin.DeleteProduk).Methods("DELETE")

	// >> CRUD KERANJANG ROUTE
	admin.HandleFunc("/keranjang", Admin.ListKeranjangs).Methods("GET")
	admin.HandleFunc("/keranjang", Admin.CreateKeranjang).Methods("POST")
	admin.HandleFunc("/keranjang/{id}", Admin.GetKeranjangByID).Methods("GET")
	admin.HandleFunc("/keranjang/{id}", Admin.UpdateKeranjang).Methods("PUT")
	admin.HandleFunc("/keranjang/{id}", Admin.DeleteKeranjang).Methods("DELETE")

	// >> CRUD TRANSAKSI ROUTE
	admin.HandleFunc("/transaksi", Admin.GetAllTransaksi).Methods("GET")
	admin.HandleFunc("/transaksi", Admin.CreateTransaksi).Methods("POST")
	admin.HandleFunc("/transaksi/{id}", Admin.GetTransaksiByID).Methods("GET")
	admin.HandleFunc("/transaksi/{id}", Admin.DeleteTransaksi).Methods("DELETE")

	// ------------------------[ AUTH KASIR ]----------------------------------
	kasir := r.PathPrefix("/kasir").Subrouter()
	kasir.Use(Middleware.JWTKasirMiddleware)

	// >> CRUD PELANGGAN ROUTE
	kasir.HandleFunc("/pelanggan", Kasir.ListPelanggan).Methods("GET")
	kasir.HandleFunc("/pelanggan", Kasir.CreatePelanggan).Methods("POST")
	kasir.HandleFunc("/pelanggan/{id}", Kasir.GetPelangganByID).Methods("GET")
	kasir.HandleFunc("/pelanggan/{id}", Kasir.UpdatePelanggan).Methods("PUT")

	// >> CRUD KATEGORI PRODUK ROUTE
	kasir.HandleFunc("/kategori-produk", Admin.ListKategoriProduk).Methods("GET")
	kasir.HandleFunc("/kategori-produk/{id}", Kasir.GetKategoriProdukByID).Methods("GET")

	// >> CRUD PRODUK ROUTE
	kasir.HandleFunc("/produk", Admin.GetAllProduk).Methods("GET")
	kasir.HandleFunc("/produk/{id}", Kasir.GetProdukByID).Methods("GET")

	// >> CRUD KERANJANG ROUTE
	kasir.HandleFunc("/keranjang", Kasir.ListKeranjangs).Methods("GET")
	kasir.HandleFunc("/keranjang", Kasir.CreateKeranjang).Methods("POST")
	kasir.HandleFunc("/keranjang/{id}", Kasir.GetKeranjangByID).Methods("GET")
	kasir.HandleFunc("/keranjang/{id}", Kasir.UpdateKeranjang).Methods("PUT")
	kasir.HandleFunc("/keranjang/{id}", Kasir.DeleteKeranjang).Methods("DELETE")

	// >> CRUD TRANSAKSI ROUTE
	kasir.HandleFunc("/transaksi", Kasir.GetAllTransaksi).Methods("GET")
	kasir.HandleFunc("/transaksi", Kasir.CreateTransaksi).Methods("POST")
	kasir.HandleFunc("/transaksi/{id}", Kasir.GetTransaksiByID).Methods("GET")
	kasir.HandleFunc("/transaksi/{id}", Kasir.DeleteTransaksi).Methods("DELETE")

	// >> CRUD DETAIL TRANSAKSI ROUTE
	kasir.HandleFunc("/detail-transaksi", Kasir.ListDetailTransaksi).Methods("GET")
	kasir.HandleFunc("/detail-transaksi", Kasir.CreateDetailTransaksi).Methods("POST")
	kasir.HandleFunc("/detail-transaksi/{id}", Kasir.GetDetailTransaksiByID).Methods("GET")
	kasir.HandleFunc("/detail-transaksi/{id}", Kasir.UpdateDetailTransaksi).Methods("PUT")
	kasir.HandleFunc("/detail-transaksi/{id}", Kasir.DeleteDetailTransaksi).Methods("DELETE")

	// ------------------------[  START APP  ]----------------------------------
	if Utils.IS_DISPLAY == "1" {
		fmt.Printf("\n > Server running on: %s", Utils.Serv.Sprintf("http://%s:%s\n", Utils.APP_HOST, Utils.APP_PORT))
		Utils.Display()
	} else {
		fmt.Printf("\n > Server running on: %s", Utils.Serv.Sprintf("http://%s:%s", Utils.APP_HOST, Utils.APP_PORT))
	}

	fmt.Println("\n\n ------------------------------------------- [ LOG History ] -------------------------------------------\n")

	log.Fatal(http.ListenAndServe(Utils.APP_CONF, r))
}