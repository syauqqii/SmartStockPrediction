package Models

type Transaksi struct {
	ID                 	int       	`gorm:"column:id_transaksi;primaryKey;autoIncrement" json:"id_transaksi"`
	IDPelanggan        	int       	`gorm:"column:id_pelanggan"                          json:"id_pelanggan"`
	TanggalTransaksi   	string 	`gorm:"column:tanggal_transaksi"                     json:"tanggal_transaksi"`
	TotalHargaTransaksi float64  	`gorm:"column:total_harga_transaksi"                 json:"total_harga_transaksi"`
	Pelanggan     		Pelanggan   `gorm:"foreignKey:IDPelanggan" 						 json:"pelanggan"`
}

type TransaksiInput struct {
	IDPelanggan        int       `json:"id_pelanggan"`
	TanggalTransaksi   string `json:"tanggal_transaksi"`
	TotalHargaTransaksi float64   `json:"total_harga_transaksi"`
}

type TransaksiResponse struct {
	ID                 int       `json:"id"`
	IDPelanggan        int       `json:"id_pelanggan"`
	TanggalTransaksi   string `json:"tanggal_transaksi"`
	TotalHargaTransaksi float64   `json:"total_harga_transaksi"`
}

type TransaksiListResponse struct {
	Transaksis []TransaksiResponse `json:"transaksis"`
}