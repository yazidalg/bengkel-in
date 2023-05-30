package sparepart

import "fmt"

type sparepart struct {
	nama                      string
	id, harga, stok, sold_out int
}

type arrSparepart [Nmax]sparepart

const Nmax = 5

func AddSparepart() {
	var sparepart arrSparepart
	var i int
	fmt.Println("Tambah Data Sparepart")
	fmt.Println("Ketik `exit` jika ingin membatalkan")
	fmt.Println()
	inputSparepart(&sparepart, &i)
	printSparepart(sparepart, i)
}

func inputSparepart(T *arrSparepart, n *int) {
	fmt.Println("Masukan data sparepart")
	fmt.Println("1 baris data yang dipisahkan oleh spasi untuk mengisi-nya")
	for T[*n].nama != "exit" {
		*n++
		fmt.Scanln(&T[*n].nama, &T[*n].harga, &T[*n].stok)
		T[*n].id = *n
	}
}

func printSparepart(T arrSparepart, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i].nama, T[i].harga, T[i].stok)
	}
}
