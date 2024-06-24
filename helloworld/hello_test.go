package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Henry", "")
		want := "Hello, Henry"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elena", "Spanish")
		want := "Hola, Elena"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Adam", "French")
		want := "Bonjour, Adam"
		
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Vietnamese", func(t *testing.T) {
		got := Hello("Hoang", "Vietnamese")
		want := "Chao, Hoang"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// testing.TB is an interface implemented by both *testing.T and *testing.B for test and benchmark
	t.Helper()	// t.Helper() indicates that this function is the helper function, and if an error is caught it should return the line error in the caller function
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
