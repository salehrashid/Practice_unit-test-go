package main

import (
	"testing"
)

func TestHelloSaleh(t *testing.T) {
	result := FuncHello("saleh")
	if result != "Hello saleh" {
		/**
		tidak disarankan menggunakan ini karena ketika gagal
		maka akan menggagalkan func test selanjutnya
		*/
		panic("panic")
	}
	println("print panic")
}

func TestHelloRizky(t *testing.T) {
	result := FuncHello("rizky")
	if result != "Hello rizky" {
		/**
		t.Fail(), saat terjadi error maka func test gagal dan
		akan melanjutkan kode dibawahnya
		*/
		t.Fail()
	}
	//dijalanin
	println("print t.Fail")
}

func TestHelloUmar(t *testing.T) {
	result := FuncHello("umar")
	if result != "Hello umar" {
		/**
		t.FailNow(), saat terjadi error maka func test gagal dan
		tidak akan melanjutkan kode dibawahnya
		*/
		t.FailNow()
	}
	//tidak dijalanin
	println("print t.FailNow")
}

func TestHelloAlip(t *testing.T) {
	result := FuncHello("alip")
	if result != "Hello alip" {
		/**
		t.Error(), saat terjadi error maka func test gagal dan
		mengeprint log error akan melanjutkan kode
		dibawahnya
		*/
		t.Error("pinjam dulu seratus")
	}
	//dijalanin
	println("print t.Error")
}

func TestHelloRapip(t *testing.T) {
	result := FuncHello("rapip")
	if result != "Hello rapip" {
		/**
		t.Fatal(), saat terjadi error maka func test gagal dan
		mengeprint log error lalu tidak akan melanjutkan kode
		dibawahnya
		*/
		t.Fatal("pinjam dulu seratus")
	}
	//tidak dijalanin
	println("print t.Fatal")
}

func TestOne(t *testing.T) {
	a := 2
	b := 2
	if a == b {
		println("benar")
	} else {
		t.Error("yyy")
	}
}
