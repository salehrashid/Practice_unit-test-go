package main

import "testing"

func TestMain(m *testing.M) {
	//terserah kode nya apa aja
	for i := 1; i <= 3; i++ {
		println(i)
	}
	println("let's go to test")

	//untuk menjalankan semua test dalam package yang sama yaitu "helper"
	m.Run()

	//terserah kode nya apa aja
	println("stop the test dan pinjam dulu seratus")
}
