package main

import "testing"

func TestNewD(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected a deck of 16 but got: %v", len(d))
	}
}
