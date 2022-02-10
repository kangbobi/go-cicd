package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello World"
	got := hello()

	if got != want {
		t.Fatalf("error")
	}
}
