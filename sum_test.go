package main

import "testing"

func TestSum(t *testing.T) {

	resultSum := Sum(5, 5)

	if resultSum != 10 {
		t.Error("O resultado está diferente de 10")
	}
}
