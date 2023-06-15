package Route

import (
	"os"
	"fmt"
	"syscall"
	"context"
	"net/http"
	"os/signal"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Middleware"
	"SmartStockPrediction/Controllers"
	"SmartStockPrediction/Controllers/Admin"
	"SmartStockPrediction/Controllers/Kasir"
)

func RunRoute() {
	r := mux.NewRouter()

	// ------------------------[  NO AUTH ROUTE ]----------------------------------
	r.HandleFunc("/login", Controllers.Login).Methods("POST")
	r.HandleFunc("/register", Controllers.Register).Methods("POST")
	r.HandleFunc("/logout", Controllers.Logout).Methods("GET")
	

	// ------------------------[ AUTH ADMIN ROUTE ]----------------------------------
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(Middleware.JWTAdminMiddleware)	// Set Middleware -> admin

	// >> CRUD USER ROUTE
	admin.HandleFunc("/users", Admin.CreateUser).Methods("POST")		// Create
	admin.HandleFunc("/users", Admin.ListUser).Methods("GET")			// Read
	admin.HandleFunc("/users/{id}", Admin.GetUserByID).Methods("GET")	// Read By ID
	admin.HandleFunc("/users/{id}", Admin.UpdateUser).Methods("PUT")	// Update
	admin.HandleFunc("/users/{id}", Admin.DeleteUser).Methods("DELETE")	// Delete

	// >> CRUD PELANGGAN ROUTE
	admin.HandleFunc("/pelanggan", Admin.CreatePelanggan).Methods("POST")			// Create
	admin.HandleFunc("/pelanggan", Admin.ListPelanggan).Methods("GET")				// Read
	admin.HandleFunc("/pelanggan/{id}", Admin.GetPelangganByID).Methods("GET")		// Read By ID
	admin.HandleFunc("/pelanggan/{id}", Admin.UpdatePelanggan).Methods("PUT")		// Update
	admin.HandleFunc("/pelanggan/{id}", Admin.DeletePelanggan).Methods("DELETE")	// Delete

	// >> CRUD KATEGORI PRODUK ROUTE
	admin.HandleFunc("/kategori-produk", Admin.CreateKategoriProduk).Methods("POST")		// Create
	admin.HandleFunc("/kategori-produk", Admin.ListKategoriProduk).Methods("GET")			// Read
	admin.HandleFunc("/kategori-produk/{id}", Admin.GetKategoriProdukByID).Methods("GET")	// Read By ID
	admin.HandleFunc("/kategori-produk/{id}", Admin.UpdateKategoriProduk).Methods("PUT")	// Update
	admin.HandleFunc("/kategori-produk/{id}", Admin.DeleteKategoriProduk).Methods("DELETE")	// Delete

	// >> CRUD PRODUK ROUTE
	admin.HandleFunc("/produk", Admin.CreateProduk).Methods("POST")			// Create
	admin.HandleFunc("/produk", Admin.GetAllProduk).Methods("GET")			// Read
	admin.HandleFunc("/produk/{id}", Admin.GetProdukByID).Methods("GET")	// Read By ID
	admin.HandleFunc("/produk/{id}", Admin.UpdateProduk).Methods("PUT")		// Update
	admin.HandleFunc("/produk/{id}", Admin.DeleteProduk).Methods("DELETE")	// Delete

	// >> MACHINE LEARNING ROUTE
	admin.HandleFunc("/rekomendasi-dari-stok", Admin.GetRecommendation).Methods("GET")						// Read
	admin.HandleFunc("/rekomendasi-dari-penjualan", Admin.GetRecommendationByTransaction).Methods("GET")	// Read

	// ------------------------[ AUTH KASIR ROUTE ]----------------------------------
	kasir := r.PathPrefix("/kasir").Subrouter()
	kasir.Use(Middleware.JWTKasirMiddleware) // Set Middleware -> kasir

	// >> CRU PELANGGAN ROUTE
	kasir.HandleFunc("/pelanggan", Kasir.CreatePelanggan).Methods("POST")		// Create
	kasir.HandleFunc("/pelanggan", Kasir.ListPelanggan).Methods("GET")			// Read
	kasir.HandleFunc("/pelanggan/{id}", Kasir.GetPelangganByID).Methods("GET")	// Read By ID
	kasir.HandleFunc("/pelanggan/{id}", Kasir.UpdatePelanggan).Methods("PUT")	// Update

	// >> R KATEGORI PRODUK ROUTE
	kasir.HandleFunc("/kategori-produk", Admin.ListKategoriProduk).Methods("GET")			// Read
	kasir.HandleFunc("/kategori-produk/{id}", Kasir.GetKategoriProdukByID).Methods("GET")	// Read By ID

	// >> R PRODUK ROUTE
	kasir.HandleFunc("/produk", Admin.GetAllProduk).Methods("GET")			// Read
	kasir.HandleFunc("/produk/{id}", Kasir.GetProdukByID).Methods("GET")	// Read By ID

	// >> CRUD KERANJANG ROUTE
	kasir.HandleFunc("/keranjang", Kasir.CreateKeranjang).Methods("POST")			// Create
	kasir.HandleFunc("/keranjang", Kasir.ListKeranjangs).Methods("GET")				// Read
	kasir.HandleFunc("/keranjang/{id}", Kasir.GetKeranjangByID).Methods("GET")		// Read By ID
	kasir.HandleFunc("/keranjang/{id}", Kasir.UpdateKeranjang).Methods("PUT")		// Update
	kasir.HandleFunc("/keranjang/{id}", Kasir.DeleteKeranjang).Methods("DELETE")	// Delete

	// >> CR TRANSAKSI ROUTE
	kasir.HandleFunc("/transaksi", Kasir.CreateTransaksi).Methods("POST")		// Create
	kasir.HandleFunc("/transaksi", Kasir.GetAllTransaksi).Methods("GET")		// Read
	kasir.HandleFunc("/transaksi/{id}", Kasir.GetTransaksiByID).Methods("GET")	// Read By ID

	// >> CRUD DETAIL TRANSAKSI ROUTE
	kasir.HandleFunc("/detail-transaksi", Kasir.CreateDetailTransaksi).Methods("POST")		// Create 
	kasir.HandleFunc("/detail-transaksi", Kasir.ListDetailTransaksi).Methods("GET")			// Read
	kasir.HandleFunc("/detail-transaksi/{id}", Kasir.GetDetailTransaksiByID).Methods("GET")	// Read By ID

	kasir.HandleFunc("/detail-transaksi-id/{id}", Kasir.GetDetailTransaksiByTransaksiID).Methods("GET")	// Read By ID Transaksi

	// ------------------------[  START APP  ]----------------------------------
	if Utils.IS_DISPLAY == "1" {
		fmt.Printf("\n > Server running on: %s", Utils.Serv.Sprintf("http://%s:%s\n", Utils.APP_HOST, Utils.APP_PORT))
		Utils.Display()
	} else {
		fmt.Printf("\n > Server running on: %s", Utils.Serv.Sprintf("http://%s:%s", Utils.APP_HOST, Utils.APP_PORT))
	}

	fmt.Println("\n\n ------------------------------------------- [ LOG History ] -------------------------------------------\n")

	// Start server with goroutine
	server := &http.Server{
		Addr:    Utils.APP_CONF,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Utils.Logger(4, err.Error())
		}
	}()

	// Wait for SIGINT (Ctrl+C)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Block until the signal is received
	<-stopChan

	// Shutdown the server gracefully
	if err := server.Shutdown(context.Background()); err != nil {
		Utils.Logger(4, err.Error())
	}

	Utils.Logger(3, "Berhasil keluar")
}