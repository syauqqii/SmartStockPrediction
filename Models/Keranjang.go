package Models

type Keranjang struct {
	ID            int `gorm:"column:id_keranjang;primaryKey;autoIncrement" json:"id_keranjang"`
	IDPelanggan   int `gorm:"column:id_pelanggan"                          json:"id_pelanggan"`
	IDProduk      int `gorm:"column:id_produk"                             json:"id_produk"`
	JumlahProduk  int `gorm:"column:jumlah_produk"                         json:"jumlah_produk"`
}

type KeranjangInput struct {
	IDPelanggan int `json:"id_pelanggan"`
	IDProduk    int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}

type KeranjangResponse struct {
	ID          int `json:"id"`
	IDPelanggan int `json:"id_pelanggan"`
	IDProduk    int `json:"id_produk"`
	JumlahProduk int `json:"jumlah_produk"`
}

type KeranjangListResponse struct {
	Keranjangs []KeranjangResponse `json:"keranjangs"`
}