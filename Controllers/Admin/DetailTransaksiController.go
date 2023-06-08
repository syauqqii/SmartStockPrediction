package Admin

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func CreateDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	var detailTransaksiInput Models.DetailTransaksiInput
	if err := Utils.DecodeJSONBody(w, r, &detailTransaksiInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var transaksi Models.Transaksi
	if err := Database.DB.First(&transaksi, detailTransaksiInput.IDTransaksi).Error; err != nil {
		response := map[string]string{"message": "id transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	idPelanggan := transaksi.IDPelanggan

	var keranjangs []Models.Keranjang
	
	if err := Database.DB.Where("id_pelanggan = ?", idPelanggan).Find(&keranjangs).Error; err != nil {
		response := map[string]string{"message": "gagal mendapatkan keranjang pelanggan"}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var detailTransaksis []Models.DetailTransaksi

	for _, keranjang := range keranjangs {
		produk := Models.Produk{}
		if err := Database.DB.First(&produk, keranjang.IDProduk).Error; err != nil {
			response := map[string]string{"message": "gagal mendapatkan detail produk dari keranjang"}
			Utils.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		produk.StokProduk -= keranjang.JumlahProduk
		if err := Database.DB.Save(&produk).Error; err != nil {
			response := map[string]string{"message": "gagal mengurangi stok produk"}
			Utils.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

		detailTransaksi := Models.DetailTransaksi{
			IDTransaksi:     detailTransaksiInput.IDTransaksi,
			IDProduk:        keranjang.IDProduk,
			JumlahProduk:    keranjang.JumlahProduk,
			HargaProduk:     produk.HargaProduk,
		}

		detailTransaksis = append(detailTransaksis, detailTransaksi)
	}

	// 4. Masukkan ke tabel detail transaksi sebelum menghapus keranjang
	if err := Database.DB.Create(&detailTransaksis).Error; err != nil {
		response := map[string]string{"message": "gagal membuat detail transaksi"}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Hapus keranjang berdasarkan ID pelanggan
	if err := Database.DB.Where("id_pelanggan = ?", idPelanggan).Delete(&Models.Keranjang{}).Error; err != nil {
		response := map[string]string{"message": "gagal menghapus keranjang"}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil membuat detail transaksi dan menghapus keranjang"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func GetDetailTransaksiByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	detailTransaksiID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var detailTransaksi Models.DetailTransaksi
	if err := Database.DB.First(&detailTransaksi, detailTransaksiID).Error; err != nil {
		response := map[string]string{"message": "detail transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	response := Models.DetailTransaksiResponse{
		ID:              detailTransaksi.ID,
		IDTransaksi:     detailTransaksi.IDTransaksi,
		IDProduk:        detailTransaksi.IDProduk,
		JumlahProduk:    detailTransaksi.JumlahProduk,
		HargaProduk:     detailTransaksi.HargaProduk,
	}

	Utils.ResponseJSON(w, http.StatusOK, response)
}


func GetAllDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	var detailTransaksis []Models.DetailTransaksi
	if err := Database.DB.Find(&detailTransaksis).Error; err != nil {
		response := map[string]string{"message": "Gagal mendapatkan detail transaksi"}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	var detailTransaksiResponses []Models.DetailTransaksiResponse
	for _, detailTransaksi := range detailTransaksis {
		detailTransaksiResponse := Models.DetailTransaksiResponse{
			ID:              detailTransaksi.ID,
			IDTransaksi:     detailTransaksi.IDTransaksi,
			IDProduk:        detailTransaksi.IDProduk,
			JumlahProduk:    detailTransaksi.JumlahProduk,
			HargaProduk:     detailTransaksi.HargaProduk,
		}
		detailTransaksiResponses = append(detailTransaksiResponses, detailTransaksiResponse)
	}

	response := Models.DetailTransaksiListResponse{
		DetailTransaksis: detailTransaksiResponses,
	}

	Utils.ResponseJSON(w, http.StatusOK, response)
}

func UpdateDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	detailTransaksiID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var detailTransaksi Models.DetailTransaksi
	if err := Database.DB.First(&detailTransaksi, detailTransaksiID).Error; err != nil {
		response := map[string]string{"message": "Detail transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var detailTransaksiInput Models.DetailTransaksiInput
	if err := Utils.DecodeJSONBody(w, r, &detailTransaksiInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	detailTransaksi.IDTransaksi = detailTransaksiInput.IDTransaksi
	detailTransaksi.IDProduk = detailTransaksiInput.IDProduk
	detailTransaksi.JumlahProduk = detailTransaksiInput.JumlahProduk
	detailTransaksi.HargaProduk = detailTransaksiInput.HargaProduk

	if err := Database.DB.Save(&detailTransaksi).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Berhasil mengupdate detail transaksi"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func DeleteDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	detailTransaksiID, err := strconv.Atoi(params["id"])
	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var detailTransaksi Models.DetailTransaksi
	if err := Database.DB.First(&detailTransaksi, detailTransaksiID).Error; err != nil {
		response := map[string]string{"message": "detail transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := Database.DB.Delete(&detailTransaksi).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil menghapus detail transaksi"}
	Utils.ResponseJSON(w, http.StatusOK, response)
}
