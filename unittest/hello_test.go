package main

import "testing"

func TestHello(t *testing.T) {
	if err := hello(); err != nil {
		t.Error(err)
	}
}
