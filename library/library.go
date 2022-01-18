package library

import "fmt"

func namamahasiswa() {
	var s1 pelajar
	var s2 = pelajar{"Yosep", 18}
	s1.nama = "Aldo"
	s1.umur = 19
	s1.namasaya()
	s2.namasaya()
}
func Namamahasiswa() {
	namamahasiswa()
}

type pelajar struct {
	nama string
	umur int
}

func (s pelajar) namasaya() {
	fmt.Println("nama :", s.nama)
	fmt.Println("umur :", s.umur)

}
