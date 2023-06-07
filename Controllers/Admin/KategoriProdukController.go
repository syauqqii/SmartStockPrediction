package Admin

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
}

func CreateKategoriProduk(w http.ResponseWriter, r *http.Request) {
	var kategoriProdukInput Models.KategoriProdukInput

	if err := Utils.DecodeJSONBody(w, r, &kategoriProdukInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var existingKategoriProduk Models.KategoriProduk

	if err := Database.DB.Where("nama_kategori_produk = ?", kategoriProdukInput.NamaKategoriProduk).First(&existingKategoriProduk).Error; err == nil {
		response := map[string]string{"message": "nama kategori sudah ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		return
	}

	kategoriProduk := Models.KategoriProduk{
		NamaKategoriProduk: kategoriProdukInput.NamaKategoriProduk,
	}

	if err := Database.DB.Create(&kategoriProduk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil menambahkan kategori produk"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
}

func GetKategoriProdukByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kategoriProdukID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var kategoriProduk Models.KategoriProduk
	
	if err := Database.DB.First(&kategoriProduk, kategoriProdukID).Error; err != nil {
		response := map[string]string{"message": "kategori produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := Models.KategoriProdukResponse{
		ID:                 kategoriProduk.ID,
		NamaKategoriProduk: kategoriProduk.NamaKategoriProduk,
	}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func UpdateKategoriProduk(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kategoriProdukID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var kategoriProduk Models.KategoriProduk
	if err := Database.DB.First(&kategoriProduk, kategoriProdukID).Error; err != nil {
		response := map[string]string{"message": "kategori produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var kategoriProdukInput Models.KategoriProdukInput
	if err := Utils.DecodeJSONBody(w, r, &kategoriProdukInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var existingKategoriProduk Models.KategoriProduk
	if err := Database.DB.Where("nama_kategori_produk = ? AND id_kategori_produk != ?", kategoriProdukInput.NamaKategoriProduk, kategoriProdukID).First(&existingKategoriProduk).Error; err == nil {
		response := map[string]string{"message": "nama kategori sudah ada"}
		Utils.ResponseJSON(w, http.StatusConflict, response)
		return
	}

	kategoriProduk.NamaKategoriProduk = kategoriProdukInput.NamaKategoriProduk

	if err := Database.DB.Save(&kategoriProduk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil update kategori produk"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func DeleteKategoriProduk(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kategoriProdukID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var kategoriProduk Models.KategoriProduk
	if err := Database.DB.First(&kategoriProduk, kategoriProdukID).Error; err != nil {
		response := map[string]string{"message": "kategori produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := Database.DB.Delete(&kategoriProduk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil hapus kategori produk"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}
