package main

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("cannot run on windows")
	}

	assert.Equal(t, "hai", "hai", "pinjam dulu seratus")
}
