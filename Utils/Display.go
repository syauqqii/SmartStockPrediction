package Utils

import "fmt"

func Display() {

	fmt.Printf("\n   > Request Methods :")
	fmt.Printf("\n     - GET    : Read, Read {id}  | - PUT    : Update {id}")
	fmt.Printf("\n     - POST   : Create           | - DELETE : Delete {id}")

	fmt.Printf("\n\n    ===================================================================================================")

	fmt.Printf("\n\n   > Endpoint Admin [Prefix: Admin] :")
	fmt.Printf("\n     - /users                -> GET, POST         |  - /pelanggan      -> GET, POST       ")
	fmt.Printf("\n     - /users/{id}           -> GET, PUT, DELETE  |  - /pelanggan/{id} -> GET, PUT, DELETE")

	fmt.Printf("\n\n     - /kategori-produk      -> GET, POST         |  - /produk         -> GET, POST")
	fmt.Printf("\n     - /kategori-produk/{id} -> GET, PUT, DELETE  |  - /produk/{id}    -> GET, PUT, DELETE")

	fmt.Printf("\n\n     - /rekomendasi-dari-stok      -> GET")
	fmt.Printf("\n     - /rekomendasi-dari-penjualan -> GET")

	fmt.Printf("\n\n    ===================================================================================================")

	fmt.Printf("\n\n   > Endpoint Kasir [Prefix: Kasir] :")
	fmt.Printf("\n     - /pelanggan      -> GET, POST    |  - /kategori-produk/      -> GET")
	fmt.Printf("\n     - /pelanggan/{id} -> GET, PUT     |  - /kategori-produk/{id}  -> GET")

	fmt.Printf("\n\n     - /produk         -> GET          |  - /keranjang/            -> GET, POST     ")
	fmt.Printf("\n     - /produk/{id}    -> GET          |  - /keranjang/{id}        -> GET, PUT, DELETE")

	fmt.Printf("\n\n     - /transaksi      -> GET, POST    |  - /detail-transaksi      -> GET, POST")
	fmt.Printf("\n     - /transaksi/{id} -> GET          |  - /detail-transaksi/{id} -> GET")
}