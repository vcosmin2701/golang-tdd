package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Cosmin")
	want := "Hello, Cosmin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
