package Models

type KategoriProduk struct {
	ID                 int    `gorm:"column:id_kategori_produk;primaryKey;autoIncrement" json:"id_kategori_produk"`
	NamaKategoriProduk string `gorm:"column:nama_kategori_produk;varchar(255)"           json:"nama_kategori_produk"`
}

type KategoriProdukInput struct {
	NamaKategoriProduk string `json:"nama_kategori_produk"`
}

type KategoriProdukResponse struct {
	ID                 int    `json:"id_kategori_produk"`
	NamaKategoriProduk string `json:"nama_kategori_produk"`
}

type KategoriProdukListResponse struct {
	KategoriProduks []KategoriProdukResponse `json:"kategori_produks"`
}