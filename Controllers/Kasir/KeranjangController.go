package Kasir

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func CreateKeranjang(w http.ResponseWriter, r *http.Request) {
	var keranjangInput Models.KeranjangInput

	if err := Utils.DecodeJSONBody(w, r, &keranjangInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var pelanggan Models.Pelanggan

	if err := Database.DB.First(&pelanggan, keranjangInput.IDPelanggan).Error; err != nil {
		response := map[string]string{"message": "id pelanggan tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var produk Models.Produk

	if err := Database.DB.First(&produk, keranjangInput.IDProduk).Error; err != nil {
		response := map[string]string{"message": "id produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var keranjang Models.Keranjang

	if err := Database.DB.Where("id_pelanggan = ? AND id_produk = ?", keranjangInput.IDPelanggan, keranjangInput.IDProduk).First(&keranjang).Error; err == nil {
		keranjang.JumlahProduk += keranjangInput.JumlahProduk
		if err := Database.DB.Save(&keranjang).Error; err != nil {
			response := map[string]string{"message": err.Error()}
			Utils.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		response := map[string]string{"message": "jumlah produk pada keranjang berhasil diperbarui"}
		Utils.ResponseJSON(w, http.StatusOK, response)
		return
	}

	keranjang = Models.Keranjang{
		IDPelanggan:  keranjangInput.IDPelanggan,
		IDProduk:     keranjangInput.IDProduk,
		JumlahProduk: keranjangInput.JumlahProduk,
	}

	if err := Database.DB.Create(&keranjang).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil menambahkan keranjang"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
}


func GetKeranjangByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keranjangID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var keranjang Models.Keranjang
	if err := Database.DB.First(&keranjang, keranjangID).Error; err != nil {
		response := map[string]string{"message": "keranjang tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	keranjangResponse := Models.KeranjangResponse{
		ID:            keranjang.ID,
		IDPelanggan:   keranjang.IDPelanggan,
		IDProduk:      keranjang.IDProduk,
		JumlahProduk:  keranjang.JumlahProduk,
	}

	Utils.ResponseJSON(w, http.StatusOK, keranjangResponse)
}

func UpdateKeranjang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keranjangID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var keranjang Models.Keranjang
	if err := Database.DB.First(&keranjang, keranjangID).Error; err != nil {
		response := map[string]string{"message": "keranjang tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var keranjangInput Models.KeranjangInput
	if err := Utils.DecodeJSONBody(w, r, &keranjangInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var pelanggan Models.Pelanggan

	if err := Database.DB.First(&pelanggan, keranjangInput.IDPelanggan).Error; err != nil {
		response := map[string]string{"message": "id pelanggan tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var produk Models.Produk

	if err := Database.DB.First(&produk, keranjangInput.IDProduk).Error; err != nil {
		response := map[string]string{"message": "id produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	// Validasi jika ID pelanggan tidak berubah
	if keranjang.IDPelanggan != keranjangInput.IDPelanggan {
		response := map[string]string{"message": "Tidak dapat mengubah ID pelanggan"}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Validasi jika ID produk tidak berubah
	if keranjang.IDProduk != keranjangInput.IDProduk {
		response := map[string]string{"message": "Tidak dapat mengubah ID produk"}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	keranjang.JumlahProduk = keranjangInput.JumlahProduk

	if err := Database.DB.Save(&keranjang).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil mengupdate keranjang"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func DeleteKeranjang(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	keranjangID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var keranjang Models.Keranjang
	if err := Database.DB.First(&keranjang, keranjangID).Error; err != nil {
		response := map[string]string{"message": "Keranjang tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := Database.DB.Delete(&keranjang).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Berhasil menghapus keranjang"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func ListKeranjangs(w http.ResponseWriter, r *http.Request) {
	var keranjangs []Models.Keranjang
	if err := Database.DB.Find(&keranjangs).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var keranjangsResponse Models.KeranjangListResponse
	for _, keranjang := range keranjangs {
		keranjangResponse := Models.KeranjangResponse{
			ID:            keranjang.ID,
			IDPelanggan:   keranjang.IDPelanggan,
			IDProduk:      keranjang.IDProduk,
			JumlahProduk:  keranjang.JumlahProduk,
		}
		keranjangsResponse.Keranjangs = append(keranjangsResponse.Keranjangs, keranjangResponse)
	}

	Utils.ResponseJSON(w, http.StatusOK, keranjangsResponse)
}