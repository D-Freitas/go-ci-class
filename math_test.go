package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)
	if total != 30 {
		t.Error("O resultado da soma deve ser 30")
	}
}
