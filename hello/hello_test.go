package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ryan")
	want := "Hello, Ryan"

	if got != want {
		t.Errorf("got %q and wanted %q", got, want)
	}
}
