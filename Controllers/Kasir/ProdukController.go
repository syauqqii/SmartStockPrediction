package Kasir

import (
	"strconv"
	"net/http"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	produkID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var produk Models.Produk
	if err := Database.DB.Preload("KategoriProduk").First(&produk, produkID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response := map[string]string{"message": "produk tidak ditemukan"}
			Utils.ResponseJSON(w, http.StatusNotFound, response)
			return
		}
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	produkResponse := Models.ProdukResponse{
		ID:                produk.ID,
		NamaProduk:        produk.NamaProduk,
		HargaProduk:       produk.HargaProduk,
		StokProduk:        produk.StokProduk,
		IDKategoriProduk:  produk.IDKategoriProduk,
	}

	response := map[string]interface{}{"produk": produkResponse}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func GetAllProduk(w http.ResponseWriter, r *http.Request) {
	var produks []Models.Produk
	if err := Database.DB.Preload("KategoriProduk").Find(&produks).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var produkResponses []Models.ProdukResponse
	for _, produk := range produks {
		produkResponse := Models.ProdukResponse{
			ID:                produk.ID,
			NamaProduk:        produk.NamaProduk,
			HargaProduk:       produk.HargaProduk,
			StokProduk:        produk.StokProduk,
			IDKategoriProduk:  produk.IDKategoriProduk,
		}
		produkResponses = append(produkResponses, produkResponse)
	}

	response := Models.ProdukListResponse{Produks: produkResponses}
	Utils.ResponseJSON(w, http.StatusOK, response)
}