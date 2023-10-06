package main

import (
	"github.com/stretchr/testify/assert"
	go_unit_test "go-unit-test"
	"testing"
)

func TestSubTest(t *testing.T) {
	firstName, middleName, lastname := go_unit_test.CompleteName("saleh", "rashid", "babsel")

	/**
	untuk menjalankan sub test secara spesifik
	go test -v -run=TestSubTest/saleh
	*/

	//sebenarnya "saleh" adalah sebagai nama sub test ini
	t.Run("saleh", func(t *testing.T) {
		assert.Equal(t, "saleh", firstName, "harusnya saleh")
	})
	t.Run("rashid", func(t *testing.T) {
		assert.Equal(t, "rashid", middleName, "harusnya rashid")
	})
	t.Run("babsel", func(t *testing.T) {
		assert.Equal(t, "babsel", lastname, "harusnya babsel")
	})
}
