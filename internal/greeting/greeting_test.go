package greeting

import "testing"

func TestMessage(t *testing.T) {
	got := Message("estudantes")
	want := "Olá mundo"

	if got != want {
		t.Fatalf("Message() = %q, want %q", got, want)
	}
}
