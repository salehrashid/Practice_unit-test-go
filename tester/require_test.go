package main

import (
	"github.com/stretchr/testify/require"
	"go-unit-test"
	"testing"
)

func TestRequire(t *testing.T) {
	result := go_unit_test.FuncHello("saleh")

	/**
	require = t.FailNow()
	*/
	require.Equal(t, 123, 123, "harusnya sama")
	println("pinjam dulu seratus") //tidak dijalankan
	require.NotEqual(t, 123, 13, "harusnya beda")

	require.Equal(t, "Hello saleh", result, "hellonya salah")
}
