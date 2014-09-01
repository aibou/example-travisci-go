package main

import "testing"

func TestFib_2(t *testing.T) {
	actual := fib(2)
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant%v", actual, expected)
	}
}

func TestFib_3(t *testing.T) {
	actual := fib(3)
	expected := 2
	if actual != expected {
		t.Errorf("got %v\nwant%v", actual, expected)
	}
}

func TestFib_4(t *testing.T) {
	actual := fib(4)
	expected := 3
	if actual != expected {
		t.Errorf("got %v\nwant%v", actual, expected)
	}
}

func TestFib_5(t *testing.T) {
	actual := fib(5)
	expected := 5
	if actual != expected {
		t.Errorf("got %v\nwant%v", actual, expected)
	}
}

func TestFib_6(t *testing.T) {
	actual := fib(6)
	expected := 8
	if actual != expected {
		t.Errorf("got %v\nwant%v", actual, expected)
	}
}
