package Models

type Pelanggan struct {
	ID            int    `gorm:"column:id_pelanggan;primaryKey;autoIncrement" json:"id_pelanggan"`
	NamaPelanggan string `gorm:"column:nama_pelanggan;type:varchar(255)"      json:"nama_pelanggan"`
	NomorHP       string `gorm:"column:nomor_hp;type:varchar(255)"            json:"nomor_hp"`
}

type PelangganInput struct {
	NamaPelanggan string `json:"nama_pelanggan"`
	NomorHP       string `json:"nomor_hp"`
}

type PelangganResponse struct {
	ID            int    `json:"id_pelanggan"`
	NamaPelanggan string `json:"nama_pelanggan"`
	NomorHP       string `json:"nomor_hp"`
}

type PelangganListResponse struct {
	Pelanggans []PelangganResponse `json:"pelanggans"`
}