package main

import "testing"

func TestAdapterList(t *testing.T) {
	adapters := List()
	if len(adapters) != 2 {
		t.Fail()
	}
}
