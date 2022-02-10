package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Worl"
	got := hello()

	if got != want {
		t.Fatalf("error")
	}
}
