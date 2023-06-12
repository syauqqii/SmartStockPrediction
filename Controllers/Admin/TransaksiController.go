package Admin

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"SmartStockPrediction/Utils"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
)

func CreateTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksiInput Models.TransaksiInput

	if err := Utils.DecodeJSONBody(w, r, &transaksiInput); err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/TransaksiController.go -> CreateTransaksi() - 1")
		return
	}

	if transaksiInput.IDPelanggan == 0 || transaksiInput.TanggalTransaksi == "" || transaksiInput.TotalHargaTransaksi == 0 {
		response := map[string]string{"message": "request body tidak boleh kosong"}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		Utils.Logger(2, "Admin/TransaksiController.go -> CreateTransaksi() - 2")
		return
	}

	var pelanggan Models.Pelanggan

	if err := Database.DB.First(&pelanggan, transaksiInput.IDPelanggan).Error; err != nil {
		response := map[string]string{"message": "id pelanggan tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		Utils.Logger(2, "Admin/TransaksiController.go -> CreateTransaksi() - 2")
		return
	}

	var totalHarga float64

	var keranjang []Models.Keranjang
	
	if err := Database.DB.Where("id_pelanggan = ?", transaksiInput.IDPelanggan).Find(&keranjang).Error; err != nil {
		response := map[string]string{"message": "gagal mengambil keranjang"}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/TransaksiController.go -> CreateTransaksi() - 3")
		return
	}

	for _, item := range keranjang {
		var produk Models.Produk

		if err := Database.DB.First(&produk, item.IDProduk).Error; err != nil {
			response := map[string]string{"message": "produk tidak ditemukan"}
			Utils.ResponseJSON(w, http.StatusNotFound, response)
			Utils.Logger(2, "Admin/TransaksiController.go -> CreateTransaksi() - 4")
			return
		}

		if item.JumlahProduk > produk.StokProduk {
			response := map[string]string{"message": "stok produk tidak cukup"}
			Utils.ResponseJSON(w, http.StatusBadRequest, response)
			Utils.Logger(2, "Admin/TransaksiController.go -> CreateTransaksi() - 5")
			return
		}

		totalHarga += produk.HargaProduk * float64(item.JumlahProduk)
	}

	transaksi := Models.Transaksi{
		IDPelanggan:        transaksiInput.IDPelanggan,
		TanggalTransaksi:   time.Now().Format("2006-01-02 15:04:05"),
		TotalHargaTransaksi: totalHarga,
	}

	if err := Database.DB.Create(&transaksi).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		Utils.Logger(2, "Admin/TransaksiController.go -> CreateTransaksi() - 6")
		return
	}

	fmt.Println(&transaksi)

	response := map[string]string{"message": "berhasil menambahkan transaksi"}

	Utils.ResponseJSON(w, http.StatusCreated, response)
	Utils.Logger(3, "Admin/TransaksiController.go -> CreateTransaksi()")
}


func GetTransaksiByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transaksiID, err := strconv.Atoi(params["id"])

	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var transaksi Models.Transaksi

	if err := Database.DB.Preload("Pelanggan").First(&transaksi, transaksiID).Error; err != nil {
		response := map[string]string{"message": "transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	transaksiResponse := Models.TransaksiResponse{
		ID:                 transaksi.ID,
		IDPelanggan:        transaksi.IDPelanggan,
		TanggalTransaksi:   transaksi.TanggalTransaksi,
		TotalHargaTransaksi: transaksi.TotalHargaTransaksi,
	}

	response := map[string]interface{}{"transaksi": transaksiResponse}

	Utils.ResponseJSON(w, http.StatusOK, response)
}

func GetAllTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksis []Models.Transaksi

	if err := Database.DB.Preload("Pelanggan").Find(&transaksis).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	transaksiResponses := make([]Models.TransaksiResponse, 0, len(transaksis))

	for _, transaksi := range transaksis {
		transaksiResponse := Models.TransaksiResponse{
			ID:                 transaksi.ID,
			IDPelanggan:        transaksi.IDPelanggan,
			TanggalTransaksi:   transaksi.TanggalTransaksi,
			TotalHargaTransaksi: transaksi.TotalHargaTransaksi,
		}
		transaksiResponses = append(transaksiResponses, transaksiResponse)
	}

	response := map[string]interface{}{"transaksis": transaksiResponses}
	Utils.ResponseJSON(w, http.StatusOK, response)
}

func DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transaksiID, err := strconv.Atoi(params["id"])

	if err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var transaksi Models.Transaksi

	if err := Database.DB.First(&transaksi, transaksiID).Error; err != nil {
		response := map[string]string{"message": "transaksi tidak ditemukan"}
		Utils.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := Database.DB.Delete(&transaksi).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		Utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "berhasil menghapus transaksi"}

	Utils.ResponseJSON(w, http.StatusOK, response)
}
