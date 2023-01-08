package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	a := tree.New(1)
	b := tree.New(2)
	got := Same(a, b)
	want := false

	if got != want {
		t.Fatalf("got %t, wanted %t", got, want)
	}

	got = Same(a, a)
	want = true

	if got != want {
		t.Fatalf("got %t, wanted %t", got, want)
	}
}
