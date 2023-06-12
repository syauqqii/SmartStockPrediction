package Admin

import (
	"fmt"
	"net/http"
	"SmartStockPrediction/Models"
	"SmartStockPrediction/Database"
	"SmartStockPrediction/Utils"
)

func GetRecommendation(w http.ResponseWriter, r *http.Request) {
	recommendations := GetProductRecommendations()

	response := make(map[string]float64)
	for namaProduk, presentase := range recommendations {
		response[namaProduk] = presentase
	}

	Utils.ResponseJSON(w, http.StatusOK, response)
	fmt.Printf("Hasil rekomendasi: %v", response)

	return
}

func GetProductRecommendations() map[string]float64 {
	var produkList []Models.Produk
	if err := Database.DB.Find(&produkList).Error; err != nil {
		fmt.Println("Gagal mengambil data produk:", err)
		return nil
	}

	var detailTransaksiList []Models.DetailTransaksi
	if err := Database.DB.Find(&detailTransaksiList).Error; err != nil {
		fmt.Println("Gagal mengambil data detail transaksi:", err)
		return nil
	}

	// Menghitung regresi linear
	var sumX, sumY, sumXY, sumXSquare float64
	n := float64(len(detailTransaksiList)) // Mengubah tipe data n menjadi float64

	for _, dt := range detailTransaksiList {
		var produk Models.Produk
		if err := Database.DB.Where("id_produk = ?", dt.IDProduk).First(&produk).Error; err != nil {
			fmt.Println("Gagal mengambil data produk dengan ID:", dt.IDProduk)
			continue
		}

		sumX += float64(dt.JumlahProduk)
		sumY += float64(produk.StokProduk)
		sumXY += float64(dt.JumlahProduk) * float64(produk.StokProduk)
		sumXSquare += float64(dt.JumlahProduk) * float64(dt.JumlahProduk)
	}

	// Menghitung koefisien regresi
	// alpha: intercept, beta: slope
	if sumXSquare-sumX*sumX == 0 {
		fmt.Println("Tidak dapat menghitung koefisien regresi, pembagi adalah 0")
		return nil
	}
	beta := (n*sumXY - sumX*sumY) / (n*sumXSquare - sumX*sumX)
	alpha := (sumY/n) - beta*(sumX/n)

	// Menghitung presentase penambahan stok
	recommendations := make(map[string]float64)
	for _, dt := range detailTransaksiList {
		var produk Models.Produk
		if err := Database.DB.Where("id_produk = ?", dt.IDProduk).First(&produk).Error; err != nil {
			fmt.Println("Gagal mengambil data produk dengan ID:", dt.IDProduk)
			continue
		}

		presentase := alpha + beta*float64(dt.JumlahProduk)
		recommendations[produk.NamaProduk] = presentase
	}

	return recommendations
}
