package main

import "testing"

func Test_sayHello(t *testing.T) {
	actual := sayHello()
	want := "Hello"
	if actual != want {
		t.Errorf("rude")
	}
}
