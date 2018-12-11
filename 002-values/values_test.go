package main

import "testing"

func testSumString(t *testing.T) {
	if "golang" != SumString("go", "lang") {
		t.Fatal("Test Fail")
	}
}

func testSumInt(t *testing.T) {
	if 2 != SumInt(1, 1) {
		t.Fatal("Test Fail")
	}

	if 10 != SumInt(4, 6) {
		t.Fatal("Test Fail")
	}
}

func testDivideInt(t *testing.T) {
	if 5 != DivideInt(25, 5) {
		t.Fatal("Test Fail")
	}
}

func testAnd(t *testing.T) {
	if !And(true, false) {
		t.Fatal("Test Fail")
	}
}

func testOr(t *testing.T) {
	if !Or(true, false) {
		t.Fatal("Test Fail")
	}
}

func testNot(t *testing.T) {
	if !Not(true) {
		t.Fatal("Test Fail")
	}
}
