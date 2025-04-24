package main
import "fmt"

func main(){
	var n, p, l, harga int
	fmt.Scan(&n)
	luasAset(&n, &p, &l, &harga)
}
func luasAset(n, p, l, harga *int){
	var jumlah, totalAset, luasA, totalA int
	for *n > 0{
		fmt.Scan(p, l, harga)
		luasA = (*p)*(*l)
		jumlah = (*p)*(*l)*(*harga)
		totalAset += jumlah
		totalA += luasA
		*n-= 1
	}
	fmt.Println("Luas aset Bojongsoang Danau Permai ", totalA, " m2")
	fmt.Println("Nilai aset Bojong Danau Permai ", jumlah, " rupiah")
}