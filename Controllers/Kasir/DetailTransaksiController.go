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
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> ListDetailTransaksi() - 1")
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
	Utils.Logger(3, "Kasir/DetailTransaksiController.go -> ListDetailTransaksi()")
}

func GetDetailTransaksiByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	detailTransaksiID, err := strconv.Atoi(params["id"])

	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> GetDetailTransaksiByID() - 1")
		return
	}

	var detailTransaksi Models.DetailTransaksi

	if err := Database.DB.First(&detailTransaksi, detailTransaksiID).Error; err != nil {
		response := map[string]string{"message": "detail transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> GetDetailTransaksiByID() - 2")
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
	Utils.Logger(3, "Kasir/DetailTransaksiController.go -> GetDetailTransaksiByID()")
}

func CreateDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	var detailTransaksiInput Models.DetailTransaksiInput

	if err := Utils.DecodeJSONBody(w, r, &detailTransaksiInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> CreateDetailTransaksi() - 1")
		return
	}

	var transaksi Models.Transaksi

	if err := Database.DB.First(&transaksi, detailTransaksiInput.IDTransaksi).Error; err != nil {
		response := map[string]string{"message": "transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> CreateDetailTransaksi() - 2")
		return
	}

	var produk Models.Produk

	if err := Database.DB.First(&produk, detailTransaksiInput.IDProduk).Error; err != nil {
		response := map[string]string{"message": "produk tidak ditemukan"}
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> CreateDetailTransaksi() - 3")
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if detailTransaksiInput.JumlahProduk > produk.StokProduk {
		response := map[string]interface{}{
			"message":      "stok produk tidak mencukupi",
			"sisa_produk":  produk.StokProduk,
		}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> CreateDetailTransaksi() - 4")
		return
	}

	detailTransaksi := Models.DetailTransaksi{
		IDTransaksi:  detailTransaksiInput.IDTransaksi,
		IDProduk:     detailTransaksiInput.IDProduk,
		JumlahProduk: detailTransaksiInput.JumlahProduk,
		HargaProduk:  produk.HargaProduk,
	}

	produk.StokProduk -= detailTransaksiInput.JumlahProduk
	if err := Database.DB.Save(&produk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> CreateDetailTransaksi() - 5")
		return
	}

	if err := Database.DB.Create(&detailTransaksi).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Kasir/DetailTransaksiController.go -> CreateDetailTransaksi() - 6")
		return
	}

	response := map[string]string{"message": "berhasil membuat detail transaksi"}
	Utils.ResponseJSON(w, http.StatusCreated, response)
	Utils.Logger(3, "Kasir/DetailTransaksiController.go -> CreateDetailTransaksi()")
}

