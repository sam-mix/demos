package main

import "testing"

func TestIsHu(t *testing.T) {
	// Test case 1: simple hu with man, pin, sou
	pais1 := []uint8{Man1, Man2, Man3, Man3, Man4, Man5, Man6, Man7, Man8, Man9, Pin1, Pin1, Pin1}
	if !IsHu(pais1) {
		t.Errorf("Expected IsHu(%v) to be true, but got false", pais1)
	}

	// Test case 2: simple hu with feng and hua
	pais2 := []uint8{Dong, Dong, Dong, Nan, Nan, Nan, Xi, Xi, Xi, Bei, Bei, Bei, Zhong, Zhong}
	if !IsHu(pais2) {
		t.Errorf("Expected IsHu(%v) to be true, but got false", pais2)
	}

	// Test case 3: not a hu
	pais3 := []uint8{Man1, Man2, Man3, Man4, Man5, Man6, Man7, Man8, Man9, Pin1, Pin2, Pin3, Sou1, Sou2}
	if IsHu(pais3) {
		t.Errorf("Expected IsHu(%v) to be false, but got true", pais3)
	}
}
