package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hello Separador de Times")
	want := "Hello Separador de Times"

	if want != got {
		t.Errorf("want %s got %s", want, got)
	}
}
