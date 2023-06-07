package Kasir

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func ListDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	var detailTransaksis []Models.DetailTransaksi

	if err := Database.DB.Find(&detailTransaksis).Error; err != nil {
		response := map[string]string{"message": err.Error()}
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

	response := Models.DetailTransaksiListResponse{DetailTransaksis: detailTransaksiResponses}
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

	detailTransaksiResponse := Models.DetailTransaksiResponse{
		ID:              detailTransaksi.ID,
		IDTransaksi:     detailTransaksi.IDTransaksi,
		IDProduk:        detailTransaksi.IDProduk,
		JumlahProduk:    detailTransaksi.JumlahProduk,
		HargaProduk:     detailTransaksi.HargaProduk,
	}

	Utils.ResponseJSON(w, http.StatusOK, detailTransaksiResponse)
}

func CreateDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	var detailTransaksiInput Models.DetailTransaksiInput

	if err := Utils.DecodeJSONBody(w, r, &detailTransaksiInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var transaksi Models.Transaksi

	if err := Database.DB.First(&transaksi, detailTransaksiInput.IDTransaksi).Error; err != nil {
		response := map[string]string{"message": "transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	var produk Models.Produk

	if err := Database.DB.First(&produk, detailTransaksiInput.IDProduk).Error; err != nil {
		response := map[string]string{"message": "produk tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	detailTransaksi := Models.DetailTransaksi{
		IDTransaksi:     detailTransaksiInput.IDTransaksi,
		IDProduk:        detailTransaksiInput.IDProduk,
		JumlahProduk:    detailTransaksiInput.JumlahProduk,
		HargaProduk:     detailTransaksiInput.HargaProduk,
	}

	if err := Database.DB.Create(&detailTransaksi).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil membuat detail transaksi"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
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
		response := map[string]string{"message": "detail transaksi tidak ditemukan"}
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

	response := map[string]string{"message": "berhasil mengupdate detail transaksi"}
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