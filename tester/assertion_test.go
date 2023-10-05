package main

import (
	"github.com/stretchr/testify/assert"
	"go-unit-test"
	"testing"
)

func TestAssertion(t *testing.T) {
	result := go_unit_test.FuncHello("saleh")

	/**
	assert = t.Fail()
	*/
	assert.Equal(t, 123, 123, "harusnya sama")
	println("pinjam dulu seratus") //dijalankan
	assert.NotEqual(t, 123, 13, "harusnya beda")

	assert.Equal(t, "Hello saleh", result)
}
