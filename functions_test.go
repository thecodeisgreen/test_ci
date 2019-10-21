package main

import "testing"

func TestF1(t *testing.T) {
	if F1() != "F1" {
		t.Error("F1() must return F1")
	}
}

func TestF2(t *testing.T) {
	if F2() != "F2" {
		t.Error("F2() must return F2")
	}
}
