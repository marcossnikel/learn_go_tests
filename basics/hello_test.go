package main

import "testing" 

func TestHello(t *testing.T){
	got := Hello("Marcos")
	want := "Hello Marcos"

	if got != want{
		t.Errorf("got %q want %q",got,want)
	}
}