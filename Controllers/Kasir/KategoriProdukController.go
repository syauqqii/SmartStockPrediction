package Kasir

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func ListKategoriProduk(w http.ResponseWriter, r *http.Request) {
	var kategoriProduks []Models.KategoriProduk

	if err := Database.DB.Find(&kategoriProduks).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Kasir/kategoriProdukController.go -> ListKategoriProduk() - 1")
		return
	}

	kategoriProdukResponses := make([]Models.KategoriProdukResponse, len(kategoriProduks))

	for i, kategoriProduk := range kategoriProduks {
		kategoriProdukResponses[i] = Models.KategoriProdukResponse{
			ID:                 kategoriProduk.ID,
			NamaKategoriProduk: kategoriProduk.NamaKategoriProduk,
		}
	}

	response := Models.KategoriProdukListResponse{KategoriProduks: kategoriProdukResponses}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Kasir/kategoriProdukController.go -> ListKategoriProduk()")
}

func GetKategoriProdukByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kategoriProdukID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Kasir/kategoriProdukController.go -> GetKategoriProdukByID() - 1")
		return
	}

	var kategoriProduk Models.KategoriProduk
	
	if err := Database.DB.First(&kategoriProduk, kategoriProdukID).Error; err != nil {
		response := map[string]string{"message": "kategori produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Kasir/kategoriProdukController.go -> GetKategoriProdukByID() - 2")
		return
	}

	response := Models.KategoriProdukResponse{
		ID:                 kategoriProduk.ID,
		NamaKategoriProduk: kategoriProduk.NamaKategoriProduk,
	}
	Utils.ResponseJSON(w, http.StatusOK, response)
	Utils.Logger(3, "Kasir/kategoriProdukController.go -> GetKategoriProdukByID()")
}