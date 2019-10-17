package main

import "testing"

func TestNumeric(t *testing.T) {
	for _, sut := range []struct {
		value       string
		description string
	}{
		{value: "1", description: "1です"},
		{value: "2", description: "2です"},
		{value: "3", description: "3です"},
		{value: "4", description: "4です"},
		{value: "5", description: "5です"},
	} {
		if !isNumeric(sut.value) {
			t.Errorf("expected number , but actual %s. test case :%s", sut.value, sut.description)
		}
	}
}

func TestNumericFail(t *testing.T) {
	for _, sut := range []struct {
		value       string
		description string
	}{
		{value: "one", description: "だめです"},
		{value: "two", description: "だめです"},
		{value: "three", description: "だめです"},
		{value: "four", description: "だめです"},
		{value: "five", description: "だめです"},
	} {
		if isNumeric(sut.value) {
			t.Errorf("expected not number , but actual %s. test case :%s", sut.value, sut.description)
		}
	}
}
